/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package eip

import (
	"context"
	"github.com/apache/plc4x/plc4go/pkg/api/model"
	"github.com/apache/plc4x/plc4go/pkg/api/values"
	readWriteModel "github.com/apache/plc4x/plc4go/protocols/eip/readwrite/model"
	"github.com/apache/plc4x/plc4go/spi"
	plc4goModel "github.com/apache/plc4x/plc4go/spi/model"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"strings"
	"time"
)

type Writer struct {
	messageCodec  spi.MessageCodec
	tm            *spi.RequestTransactionManager
	configuration Configuration
	sessionHandle *uint32
	senderContext *[]uint8
}

func NewWriter(messageCodec spi.MessageCodec, tm *spi.RequestTransactionManager, configuration Configuration, sessionHandle *uint32, senderContext *[]uint8) Writer {
	return Writer{
		messageCodec:  messageCodec,
		tm:            tm,
		configuration: configuration,
		sessionHandle: sessionHandle,
		senderContext: senderContext,
	}
}

func (m Writer) Write(ctx context.Context, writeRequest model.PlcWriteRequest) <-chan model.PlcWriteRequestResult {
	// TODO: handle context
	result := make(chan model.PlcWriteRequestResult)
	go func() {
		items := make([]readWriteModel.CipService, len(writeRequest.GetFieldNames()))
		for i, fieldName := range writeRequest.GetFieldNames() {
			field := writeRequest.GetField(fieldName).(EIPPlcField)
			value := writeRequest.GetValue(fieldName)
			tag := field.GetTag()
			elements := uint16(1)
			if field.GetElementNb() > 1 {
				elements = field.GetElementNb()
			}
			// We need the size of the request in words (0x91, tagLength, ... tag + possible pad)
			// Taking half to get word size
			isArray := false
			tagIsolated := tag
			if strings.Contains(tag, "[") {
				isArray = true
				tagIsolated = tag[:strings.Index(tag, "[")]
			}
			dataLength := len(tagIsolated) + 2 + (len(tagIsolated) % 2)
			if isArray {
				dataLength += 2
			}
			requestPathSize := int8(dataLength / 2)
			data, err := encodeValue(value, field.GetType(), elements)
			if err != nil {
				result <- &plc4goModel.DefaultPlcWriteRequestResult{
					Request:  writeRequest,
					Response: nil,
					Err:      errors.Wrapf(err, "Error encoding value for field %s", fieldName),
				}
				return
			}
			ansi, err := toAnsi(tag)
			if err != nil {
				result <- &plc4goModel.DefaultPlcWriteRequestResult{
					Request:  writeRequest,
					Response: nil,
					Err:      errors.Wrapf(err, "Error encoding eip ansi for field %s", fieldName),
				}
				return
			}
			items[i] = readWriteModel.NewCipWriteRequest(requestPathSize, ansi, field.GetType(), elements, data, 0)
		}

		if len(items) == 1 {
			// Assemble the finished paket
			log.Trace().Msg("Assemble paket")
			pkt := readWriteModel.NewCipRRData(
				readWriteModel.NewCipExchange(
					readWriteModel.NewCipUnconnectedRequest(
						items[0],
						m.configuration.backplane,
						m.configuration.slot,
						0,
					),
					0,
				),
				*m.sessionHandle,
				0,
				*m.senderContext,
				0,
				0,
			)
			// Start a new request-transaction (Is ended in the response-handler)
			transaction := m.tm.StartTransaction()
			transaction.Submit(func() {
				// Send the  over the wire
				if err := m.messageCodec.SendRequest(ctx, pkt, func(message spi.Message) bool {
					eipPacket := message.(readWriteModel.EipPacket)
					if eipPacket == nil {
						return false
					}
					cipRRData := eipPacket.(readWriteModel.CipRRData)
					if cipRRData == nil {
						return false
					}
					if eipPacket.GetSessionHandle() != *m.sessionHandle {
						return false
					}
					cipWriteResponse := cipRRData.GetExchange().GetService().(readWriteModel.CipWriteResponse)
					if cipWriteResponse == nil {
						return false
					}
					return true
				}, func(message spi.Message) error {
					// Convert the response into an
					log.Trace().Msg("convert response to ")
					eipPacket := message.(readWriteModel.EipPacket)
					cipRRData := eipPacket.(readWriteModel.CipRRData)
					cipWriteResponse := cipRRData.GetExchange().GetService().(readWriteModel.CipWriteResponse)
					// Convert the eip response into a PLC4X response
					log.Trace().Msg("convert response to PLC4X response")
					readResponse, err := m.ToPlc4xWriteResponse(cipWriteResponse, writeRequest)

					if err != nil {
						result <- &plc4goModel.DefaultPlcWriteRequestResult{
							Request: writeRequest,
							Err:     errors.Wrap(err, "Error decoding response"),
						}
						return transaction.EndRequest()
					}
					result <- &plc4goModel.DefaultPlcWriteRequestResult{
						Request:  writeRequest,
						Response: readResponse,
					}
					return transaction.EndRequest()
				}, func(err error) error {
					result <- &plc4goModel.DefaultPlcWriteRequestResult{
						Request: writeRequest,
						Err:     errors.New("got timeout while waiting for response"),
					}
					return transaction.EndRequest()
				}, time.Second*1); err != nil {
					result <- &plc4goModel.DefaultPlcWriteRequestResult{
						Request:  writeRequest,
						Response: nil,
						Err:      errors.Wrap(err, "error sending message"),
					}
					_ = transaction.EndRequest()
				}
			})
		} else {
			nb := uint16(len(items))
			offsets := make([]uint16, nb)
			offset := 2 + nb*2
			for i := uint16(0); i < nb; i++ {
				offsets[i] = offset
				offset += items[i].GetLengthInBytes()
			}

			serviceArr := make([]readWriteModel.CipService, nb)
			for i := uint16(0); i < nb; i++ {
				serviceArr[i] = items[i]
			}

			data := readWriteModel.NewServices(nb, offsets, serviceArr, 0)

			// Assemble the finished paket
			log.Trace().Msg("Assemble paket")
			pkt := readWriteModel.NewCipRRData(
				readWriteModel.NewCipExchange(
					readWriteModel.NewCipUnconnectedRequest(
						readWriteModel.NewMultipleServiceRequest(data, 0),
						m.configuration.backplane,
						m.configuration.slot,
						0,
					),
					0,
				),
				*m.sessionHandle,
				0,
				*m.senderContext,
				0,
				0,
			)
			// Start a new request-transaction (Is ended in the response-handler)
			transaction := m.tm.StartTransaction()
			transaction.Submit(func() {
				// Send the  over the wire
				if err := m.messageCodec.SendRequest(ctx, pkt, func(message spi.Message) bool {
					eipPacket := message.(readWriteModel.EipPacket)
					if eipPacket == nil {
						return false
					}
					cipRRData := eipPacket.(readWriteModel.CipRRData)
					if cipRRData == nil {
						return false
					}
					if eipPacket.GetSessionHandle() != *m.sessionHandle {
						return false
					}
					multipleServiceResponse := cipRRData.GetExchange().GetService().(readWriteModel.MultipleServiceResponse)
					if multipleServiceResponse == nil {
						return false
					}
					if multipleServiceResponse.GetServiceNb() != nb {
						return false
					}
					return true
				}, func(message spi.Message) error {
					// Convert the response into an
					log.Trace().Msg("convert response to ")
					eipPacket := message.(readWriteModel.EipPacket)
					cipRRData := eipPacket.(readWriteModel.CipRRData)
					multipleServiceResponse := cipRRData.GetExchange().GetService().(readWriteModel.MultipleServiceResponse)
					// Convert the eip response into a PLC4X response
					log.Trace().Msg("convert response to PLC4X response")
					readResponse, err := m.ToPlc4xWriteResponse(multipleServiceResponse, writeRequest)

					if err != nil {
						result <- &plc4goModel.DefaultPlcWriteRequestResult{
							Request: writeRequest,
							Err:     errors.Wrap(err, "Error decoding response"),
						}
						return transaction.EndRequest()
					}
					result <- &plc4goModel.DefaultPlcWriteRequestResult{
						Request:  writeRequest,
						Response: readResponse,
					}
					return transaction.EndRequest()
				}, func(err error) error {
					result <- &plc4goModel.DefaultPlcWriteRequestResult{
						Request: writeRequest,
						Err:     errors.New("got timeout while waiting for response"),
					}
					return transaction.EndRequest()
				}, time.Second*1); err != nil {
					result <- &plc4goModel.DefaultPlcWriteRequestResult{
						Request:  writeRequest,
						Response: nil,
						Err:      errors.Wrap(err, "error sending message"),
					}
					_ = transaction.EndRequest()
				}
			})
		}
	}()
	return result
}

func encodeValue(value values.PlcValue, _type readWriteModel.CIPDataTypeCode, elements uint16) ([]byte, error) {
	buffer := utils.NewLittleEndianWriteBufferByteBased()
	switch _type {
	case readWriteModel.CIPDataTypeCode_SINT:
		err := buffer.WriteByte("", value.GetUint8())
		if err != nil {
			return nil, err
		}
	case readWriteModel.CIPDataTypeCode_INT:
		err := buffer.WriteInt16("", 16, value.GetInt16())
		if err != nil {
			return nil, err
		}
	case readWriteModel.CIPDataTypeCode_DINT:
		err := buffer.WriteInt32("", 32, value.GetInt32())
		if err != nil {
			return nil, err
		}
	case readWriteModel.CIPDataTypeCode_REAL:
		err := buffer.WriteFloat64("", 64, value.GetFloat64())
		if err != nil {
			return nil, err
		}
	default:
		// TODO: what is the default type? write nothing?
		//panic("unmapped type: " + strconv.Itoa(int(_type)))
	}
	return buffer.GetBytes(), nil
}

func (m Writer) ToPlc4xWriteResponse(response readWriteModel.CipService, writeRequest model.PlcWriteRequest) (model.PlcWriteResponse, error) {
	responseCodes := map[string]model.PlcResponseCode{}
	switch response := response.(type) {
	case readWriteModel.CipWriteResponse: // only 1 field
		cipReadResponse := response
		fieldName := writeRequest.GetFieldNames()[0]
		code := decodeResponseCode(cipReadResponse.GetStatus())
		responseCodes[fieldName] = code
	case readWriteModel.MultipleServiceResponse: //Multiple response
		multipleServiceResponse := response
		nb := multipleServiceResponse.GetServiceNb()
		arr := make([]readWriteModel.CipService, nb)
		read := utils.NewLittleEndianReadBufferByteBased(multipleServiceResponse.GetServicesData())
		total := read.GetTotalBytes()
		for i := uint16(0); i < nb; i++ {
			length := uint16(0)
			offset := multipleServiceResponse.GetOffsets()[i] - multipleServiceResponse.GetOffsets()[0] //Substract first offset as we only have the service in the buffer (not servicesNb and offsets)
			if i == nb-1 {
				length = uint16(total) - offset //Get the rest if last
			} else {
				length = multipleServiceResponse.GetOffsets()[i+1] - offset - multipleServiceResponse.GetOffsets()[0] //Calculate length with offsets (substracting first offset)
			}
			serviceBuf := utils.NewLittleEndianReadBufferByteBased(read.GetBytes()[offset : offset+length])
			var err error
			arr[i], err = readWriteModel.CipServiceParse(serviceBuf, length)
			if err != nil {
				return nil, err
			}
		}
		services := readWriteModel.NewServices(nb, multipleServiceResponse.GetOffsets(), arr, 0)
		for i, fieldName := range writeRequest.GetFieldNames() {
			if writeResponse, ok := services.Services[i].(readWriteModel.CipWriteResponse); ok {
				code := decodeResponseCode(writeResponse.GetStatus())
				responseCodes[fieldName] = code
			} else {
				responseCodes[fieldName] = model.PlcResponseCode_INTERNAL_ERROR
			}
		}
	default:
		return nil, errors.Errorf("unsupported response type %T", response)
	}

	// Return the response
	log.Trace().Msg("Returning the response")
	return plc4goModel.NewDefaultPlcWriteResponse(writeRequest, responseCodes), nil
}
