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

package model

import (
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetServiceAckReadPropertyConditional is the corresponding interface of BACnetServiceAckReadPropertyConditional
type BACnetServiceAckReadPropertyConditional interface {
	utils.LengthAware
	utils.Serializable
	BACnetServiceAck
	// GetBytesOfRemovedService returns BytesOfRemovedService (property field)
	GetBytesOfRemovedService() []byte
}

// BACnetServiceAckReadPropertyConditionalExactly can be used when we want exactly this type and not a type which fulfills BACnetServiceAckReadPropertyConditional.
// This is useful for switch cases.
type BACnetServiceAckReadPropertyConditionalExactly interface {
	BACnetServiceAckReadPropertyConditional
	isBACnetServiceAckReadPropertyConditional() bool
}

// _BACnetServiceAckReadPropertyConditional is the data-structure of this message
type _BACnetServiceAckReadPropertyConditional struct {
	*_BACnetServiceAck
	BytesOfRemovedService []byte

	// Arguments.
	ServiceAckPayloadLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckReadPropertyConditional) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_READ_PROPERTY_CONDITIONAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckReadPropertyConditional) InitializeParent(parent BACnetServiceAck) {}

func (m *_BACnetServiceAckReadPropertyConditional) GetParent() BACnetServiceAck {
	return m._BACnetServiceAck
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckReadPropertyConditional) GetBytesOfRemovedService() []byte {
	return m.BytesOfRemovedService
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckReadPropertyConditional factory function for _BACnetServiceAckReadPropertyConditional
func NewBACnetServiceAckReadPropertyConditional(bytesOfRemovedService []byte, serviceAckLength uint16, serviceAckPayloadLength uint16) *_BACnetServiceAckReadPropertyConditional {
	_result := &_BACnetServiceAckReadPropertyConditional{
		BytesOfRemovedService: bytesOfRemovedService,
		_BACnetServiceAck:     NewBACnetServiceAck(serviceAckLength),
	}
	_result._BACnetServiceAck._BACnetServiceAckChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckReadPropertyConditional(structType interface{}) BACnetServiceAckReadPropertyConditional {
	if casted, ok := structType.(BACnetServiceAckReadPropertyConditional); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckReadPropertyConditional); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckReadPropertyConditional) GetTypeName() string {
	return "BACnetServiceAckReadPropertyConditional"
}

func (m *_BACnetServiceAckReadPropertyConditional) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetServiceAckReadPropertyConditional) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.BytesOfRemovedService) > 0 {
		lengthInBits += 8 * uint16(len(m.BytesOfRemovedService))
	}

	return lengthInBits
}

func (m *_BACnetServiceAckReadPropertyConditional) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServiceAckReadPropertyConditionalParse(readBuffer utils.ReadBuffer, serviceAckLength uint16, serviceAckPayloadLength uint16) (BACnetServiceAckReadPropertyConditional, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckReadPropertyConditional"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckReadPropertyConditional")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (bytesOfRemovedService)
	numberOfBytesbytesOfRemovedService := int(serviceAckPayloadLength)
	bytesOfRemovedService, _readArrayErr := readBuffer.ReadByteArray("bytesOfRemovedService", numberOfBytesbytesOfRemovedService)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'bytesOfRemovedService' field of BACnetServiceAckReadPropertyConditional")
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckReadPropertyConditional"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckReadPropertyConditional")
	}

	// Create a partially initialized instance
	_child := &_BACnetServiceAckReadPropertyConditional{
		_BACnetServiceAck: &_BACnetServiceAck{
			ServiceAckLength: serviceAckLength,
		},
		BytesOfRemovedService: bytesOfRemovedService,
	}
	_child._BACnetServiceAck._BACnetServiceAckChildRequirements = _child
	return _child, nil
}

func (m *_BACnetServiceAckReadPropertyConditional) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckReadPropertyConditional"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckReadPropertyConditional")
		}

		// Array Field (bytesOfRemovedService)
		// Byte Array field (bytesOfRemovedService)
		if err := writeBuffer.WriteByteArray("bytesOfRemovedService", m.GetBytesOfRemovedService()); err != nil {
			return errors.Wrap(err, "Error serializing 'bytesOfRemovedService' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckReadPropertyConditional"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckReadPropertyConditional")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BACnetServiceAckReadPropertyConditional) GetServiceAckPayloadLength() uint16 {
	return m.ServiceAckPayloadLength
}

//
////

func (m *_BACnetServiceAckReadPropertyConditional) isBACnetServiceAckReadPropertyConditional() bool {
	return true
}

func (m *_BACnetServiceAckReadPropertyConditional) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
