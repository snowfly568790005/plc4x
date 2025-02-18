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

// LRawReq is the corresponding interface of LRawReq
type LRawReq interface {
	utils.LengthAware
	utils.Serializable
	CEMI
}

// LRawReqExactly can be used when we want exactly this type and not a type which fulfills LRawReq.
// This is useful for switch cases.
type LRawReqExactly interface {
	LRawReq
	isLRawReq() bool
}

// _LRawReq is the data-structure of this message
type _LRawReq struct {
	*_CEMI
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LRawReq) GetMessageCode() uint8 {
	return 0x10
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LRawReq) InitializeParent(parent CEMI) {}

func (m *_LRawReq) GetParent() CEMI {
	return m._CEMI
}

// NewLRawReq factory function for _LRawReq
func NewLRawReq(size uint16) *_LRawReq {
	_result := &_LRawReq{
		_CEMI: NewCEMI(size),
	}
	_result._CEMI._CEMIChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastLRawReq(structType interface{}) LRawReq {
	if casted, ok := structType.(LRawReq); ok {
		return casted
	}
	if casted, ok := structType.(*LRawReq); ok {
		return *casted
	}
	return nil
}

func (m *_LRawReq) GetTypeName() string {
	return "LRawReq"
}

func (m *_LRawReq) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_LRawReq) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_LRawReq) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func LRawReqParse(readBuffer utils.ReadBuffer, size uint16) (LRawReq, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LRawReq"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LRawReq")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("LRawReq"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LRawReq")
	}

	// Create a partially initialized instance
	_child := &_LRawReq{
		_CEMI: &_CEMI{
			Size: size,
		},
	}
	_child._CEMI._CEMIChildRequirements = _child
	return _child, nil
}

func (m *_LRawReq) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LRawReq"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LRawReq")
		}

		if popErr := writeBuffer.PopContext("LRawReq"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LRawReq")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_LRawReq) isLRawReq() bool {
	return true
}

func (m *_LRawReq) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
