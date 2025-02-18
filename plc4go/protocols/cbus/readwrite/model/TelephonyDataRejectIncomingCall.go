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

// TelephonyDataRejectIncomingCall is the corresponding interface of TelephonyDataRejectIncomingCall
type TelephonyDataRejectIncomingCall interface {
	utils.LengthAware
	utils.Serializable
	TelephonyData
}

// TelephonyDataRejectIncomingCallExactly can be used when we want exactly this type and not a type which fulfills TelephonyDataRejectIncomingCall.
// This is useful for switch cases.
type TelephonyDataRejectIncomingCallExactly interface {
	TelephonyDataRejectIncomingCall
	isTelephonyDataRejectIncomingCall() bool
}

// _TelephonyDataRejectIncomingCall is the data-structure of this message
type _TelephonyDataRejectIncomingCall struct {
	*_TelephonyData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TelephonyDataRejectIncomingCall) InitializeParent(parent TelephonyData, commandTypeContainer TelephonyCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_TelephonyDataRejectIncomingCall) GetParent() TelephonyData {
	return m._TelephonyData
}

// NewTelephonyDataRejectIncomingCall factory function for _TelephonyDataRejectIncomingCall
func NewTelephonyDataRejectIncomingCall(commandTypeContainer TelephonyCommandTypeContainer, argument byte) *_TelephonyDataRejectIncomingCall {
	_result := &_TelephonyDataRejectIncomingCall{
		_TelephonyData: NewTelephonyData(commandTypeContainer, argument),
	}
	_result._TelephonyData._TelephonyDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTelephonyDataRejectIncomingCall(structType interface{}) TelephonyDataRejectIncomingCall {
	if casted, ok := structType.(TelephonyDataRejectIncomingCall); ok {
		return casted
	}
	if casted, ok := structType.(*TelephonyDataRejectIncomingCall); ok {
		return *casted
	}
	return nil
}

func (m *_TelephonyDataRejectIncomingCall) GetTypeName() string {
	return "TelephonyDataRejectIncomingCall"
}

func (m *_TelephonyDataRejectIncomingCall) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_TelephonyDataRejectIncomingCall) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_TelephonyDataRejectIncomingCall) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func TelephonyDataRejectIncomingCallParse(readBuffer utils.ReadBuffer) (TelephonyDataRejectIncomingCall, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TelephonyDataRejectIncomingCall"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TelephonyDataRejectIncomingCall")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("TelephonyDataRejectIncomingCall"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TelephonyDataRejectIncomingCall")
	}

	// Create a partially initialized instance
	_child := &_TelephonyDataRejectIncomingCall{
		_TelephonyData: &_TelephonyData{},
	}
	_child._TelephonyData._TelephonyDataChildRequirements = _child
	return _child, nil
}

func (m *_TelephonyDataRejectIncomingCall) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TelephonyDataRejectIncomingCall"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TelephonyDataRejectIncomingCall")
		}

		if popErr := writeBuffer.PopContext("TelephonyDataRejectIncomingCall"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TelephonyDataRejectIncomingCall")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_TelephonyDataRejectIncomingCall) isTelephonyDataRejectIncomingCall() bool {
	return true
}

func (m *_TelephonyDataRejectIncomingCall) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
