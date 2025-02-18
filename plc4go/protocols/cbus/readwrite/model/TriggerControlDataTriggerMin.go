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

// TriggerControlDataTriggerMin is the corresponding interface of TriggerControlDataTriggerMin
type TriggerControlDataTriggerMin interface {
	utils.LengthAware
	utils.Serializable
	TriggerControlData
}

// TriggerControlDataTriggerMinExactly can be used when we want exactly this type and not a type which fulfills TriggerControlDataTriggerMin.
// This is useful for switch cases.
type TriggerControlDataTriggerMinExactly interface {
	TriggerControlDataTriggerMin
	isTriggerControlDataTriggerMin() bool
}

// _TriggerControlDataTriggerMin is the data-structure of this message
type _TriggerControlDataTriggerMin struct {
	*_TriggerControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TriggerControlDataTriggerMin) InitializeParent(parent TriggerControlData, commandTypeContainer TriggerControlCommandTypeContainer, triggerGroup byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.TriggerGroup = triggerGroup
}

func (m *_TriggerControlDataTriggerMin) GetParent() TriggerControlData {
	return m._TriggerControlData
}

// NewTriggerControlDataTriggerMin factory function for _TriggerControlDataTriggerMin
func NewTriggerControlDataTriggerMin(commandTypeContainer TriggerControlCommandTypeContainer, triggerGroup byte) *_TriggerControlDataTriggerMin {
	_result := &_TriggerControlDataTriggerMin{
		_TriggerControlData: NewTriggerControlData(commandTypeContainer, triggerGroup),
	}
	_result._TriggerControlData._TriggerControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTriggerControlDataTriggerMin(structType interface{}) TriggerControlDataTriggerMin {
	if casted, ok := structType.(TriggerControlDataTriggerMin); ok {
		return casted
	}
	if casted, ok := structType.(*TriggerControlDataTriggerMin); ok {
		return *casted
	}
	return nil
}

func (m *_TriggerControlDataTriggerMin) GetTypeName() string {
	return "TriggerControlDataTriggerMin"
}

func (m *_TriggerControlDataTriggerMin) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_TriggerControlDataTriggerMin) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_TriggerControlDataTriggerMin) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func TriggerControlDataTriggerMinParse(readBuffer utils.ReadBuffer) (TriggerControlDataTriggerMin, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TriggerControlDataTriggerMin"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TriggerControlDataTriggerMin")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("TriggerControlDataTriggerMin"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TriggerControlDataTriggerMin")
	}

	// Create a partially initialized instance
	_child := &_TriggerControlDataTriggerMin{
		_TriggerControlData: &_TriggerControlData{},
	}
	_child._TriggerControlData._TriggerControlDataChildRequirements = _child
	return _child, nil
}

func (m *_TriggerControlDataTriggerMin) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TriggerControlDataTriggerMin"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TriggerControlDataTriggerMin")
		}

		if popErr := writeBuffer.PopContext("TriggerControlDataTriggerMin"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TriggerControlDataTriggerMin")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_TriggerControlDataTriggerMin) isTriggerControlDataTriggerMin() bool {
	return true
}

func (m *_TriggerControlDataTriggerMin) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
