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

// EipDisconnectRequest is the corresponding interface of EipDisconnectRequest
type EipDisconnectRequest interface {
	utils.LengthAware
	utils.Serializable
	EipPacket
}

// EipDisconnectRequestExactly can be used when we want exactly this type and not a type which fulfills EipDisconnectRequest.
// This is useful for switch cases.
type EipDisconnectRequestExactly interface {
	EipDisconnectRequest
	isEipDisconnectRequest() bool
}

// _EipDisconnectRequest is the data-structure of this message
type _EipDisconnectRequest struct {
	*_EipPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EipDisconnectRequest) GetCommand() uint16 {
	return 0x0066
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EipDisconnectRequest) InitializeParent(parent EipPacket, sessionHandle uint32, status uint32, senderContext []uint8, options uint32) {
	m.SessionHandle = sessionHandle
	m.Status = status
	m.SenderContext = senderContext
	m.Options = options
}

func (m *_EipDisconnectRequest) GetParent() EipPacket {
	return m._EipPacket
}

// NewEipDisconnectRequest factory function for _EipDisconnectRequest
func NewEipDisconnectRequest(sessionHandle uint32, status uint32, senderContext []uint8, options uint32) *_EipDisconnectRequest {
	_result := &_EipDisconnectRequest{
		_EipPacket: NewEipPacket(sessionHandle, status, senderContext, options),
	}
	_result._EipPacket._EipPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastEipDisconnectRequest(structType interface{}) EipDisconnectRequest {
	if casted, ok := structType.(EipDisconnectRequest); ok {
		return casted
	}
	if casted, ok := structType.(*EipDisconnectRequest); ok {
		return *casted
	}
	return nil
}

func (m *_EipDisconnectRequest) GetTypeName() string {
	return "EipDisconnectRequest"
}

func (m *_EipDisconnectRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_EipDisconnectRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_EipDisconnectRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func EipDisconnectRequestParse(readBuffer utils.ReadBuffer) (EipDisconnectRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EipDisconnectRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EipDisconnectRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("EipDisconnectRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EipDisconnectRequest")
	}

	// Create a partially initialized instance
	_child := &_EipDisconnectRequest{
		_EipPacket: &_EipPacket{},
	}
	_child._EipPacket._EipPacketChildRequirements = _child
	return _child, nil
}

func (m *_EipDisconnectRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EipDisconnectRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EipDisconnectRequest")
		}

		if popErr := writeBuffer.PopContext("EipDisconnectRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EipDisconnectRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_EipDisconnectRequest) isEipDisconnectRequest() bool {
	return true
}

func (m *_EipDisconnectRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
