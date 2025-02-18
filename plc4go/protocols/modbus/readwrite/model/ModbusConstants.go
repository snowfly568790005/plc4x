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
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const ModbusConstants_MODBUSTCPDEFAULTPORT uint16 = uint16(502)

// ModbusConstants is the corresponding interface of ModbusConstants
type ModbusConstants interface {
	utils.LengthAware
	utils.Serializable
}

// ModbusConstantsExactly can be used when we want exactly this type and not a type which fulfills ModbusConstants.
// This is useful for switch cases.
type ModbusConstantsExactly interface {
	ModbusConstants
	isModbusConstants() bool
}

// _ModbusConstants is the data-structure of this message
type _ModbusConstants struct {
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_ModbusConstants) GetModbusTcpDefaultPort() uint16 {
	return ModbusConstants_MODBUSTCPDEFAULTPORT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusConstants factory function for _ModbusConstants
func NewModbusConstants() *_ModbusConstants {
	return &_ModbusConstants{}
}

// Deprecated: use the interface for direct cast
func CastModbusConstants(structType interface{}) ModbusConstants {
	if casted, ok := structType.(ModbusConstants); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusConstants); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusConstants) GetTypeName() string {
	return "ModbusConstants"
}

func (m *_ModbusConstants) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ModbusConstants) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Const Field (modbusTcpDefaultPort)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusConstants) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusConstantsParse(readBuffer utils.ReadBuffer) (ModbusConstants, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusConstants"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusConstants")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (modbusTcpDefaultPort)
	modbusTcpDefaultPort, _modbusTcpDefaultPortErr := readBuffer.ReadUint16("modbusTcpDefaultPort", 16)
	if _modbusTcpDefaultPortErr != nil {
		return nil, errors.Wrap(_modbusTcpDefaultPortErr, "Error parsing 'modbusTcpDefaultPort' field of ModbusConstants")
	}
	if modbusTcpDefaultPort != ModbusConstants_MODBUSTCPDEFAULTPORT {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", ModbusConstants_MODBUSTCPDEFAULTPORT) + " but got " + fmt.Sprintf("%d", modbusTcpDefaultPort))
	}

	if closeErr := readBuffer.CloseContext("ModbusConstants"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusConstants")
	}

	// Create the instance
	return &_ModbusConstants{}, nil
}

func (m *_ModbusConstants) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("ModbusConstants"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ModbusConstants")
	}

	// Const Field (modbusTcpDefaultPort)
	_modbusTcpDefaultPortErr := writeBuffer.WriteUint16("modbusTcpDefaultPort", 16, 502)
	if _modbusTcpDefaultPortErr != nil {
		return errors.Wrap(_modbusTcpDefaultPortErr, "Error serializing 'modbusTcpDefaultPort' field")
	}

	if popErr := writeBuffer.PopContext("ModbusConstants"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ModbusConstants")
	}
	return nil
}

func (m *_ModbusConstants) isModbusConstants() bool {
	return true
}

func (m *_ModbusConstants) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
