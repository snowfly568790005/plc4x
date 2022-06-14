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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataInstantaneousPower is the data-structure of this message
type BACnetConstructedDataInstantaneousPower struct {
	*BACnetConstructedData
	InstantaneousPower *BACnetApplicationTagReal

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataInstantaneousPower is the corresponding interface of BACnetConstructedDataInstantaneousPower
type IBACnetConstructedDataInstantaneousPower interface {
	IBACnetConstructedData
	// GetInstantaneousPower returns InstantaneousPower (property field)
	GetInstantaneousPower() *BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetApplicationTagReal
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataInstantaneousPower) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataInstantaneousPower) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INSTANTANEOUS_POWER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataInstantaneousPower) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataInstantaneousPower) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataInstantaneousPower) GetInstantaneousPower() *BACnetApplicationTagReal {
	return m.InstantaneousPower
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataInstantaneousPower) GetActualValue() *BACnetApplicationTagReal {
	return CastBACnetApplicationTagReal(m.GetInstantaneousPower())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataInstantaneousPower factory function for BACnetConstructedDataInstantaneousPower
func NewBACnetConstructedDataInstantaneousPower(instantaneousPower *BACnetApplicationTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataInstantaneousPower {
	_result := &BACnetConstructedDataInstantaneousPower{
		InstantaneousPower:    instantaneousPower,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataInstantaneousPower(structType interface{}) *BACnetConstructedDataInstantaneousPower {
	if casted, ok := structType.(BACnetConstructedDataInstantaneousPower); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataInstantaneousPower); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataInstantaneousPower(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataInstantaneousPower(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataInstantaneousPower) GetTypeName() string {
	return "BACnetConstructedDataInstantaneousPower"
}

func (m *BACnetConstructedDataInstantaneousPower) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataInstantaneousPower) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (instantaneousPower)
	lengthInBits += m.InstantaneousPower.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataInstantaneousPower) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataInstantaneousPowerParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataInstantaneousPower, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataInstantaneousPower"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataInstantaneousPower")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (instantaneousPower)
	if pullErr := readBuffer.PullContext("instantaneousPower"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for instantaneousPower")
	}
	_instantaneousPower, _instantaneousPowerErr := BACnetApplicationTagParse(readBuffer)
	if _instantaneousPowerErr != nil {
		return nil, errors.Wrap(_instantaneousPowerErr, "Error parsing 'instantaneousPower' field")
	}
	instantaneousPower := CastBACnetApplicationTagReal(_instantaneousPower)
	if closeErr := readBuffer.CloseContext("instantaneousPower"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for instantaneousPower")
	}

	// Virtual field
	_actualValue := instantaneousPower
	actualValue := CastBACnetApplicationTagReal(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataInstantaneousPower"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataInstantaneousPower")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataInstantaneousPower{
		InstantaneousPower:    CastBACnetApplicationTagReal(instantaneousPower),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataInstantaneousPower) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataInstantaneousPower"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataInstantaneousPower")
		}

		// Simple Field (instantaneousPower)
		if pushErr := writeBuffer.PushContext("instantaneousPower"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for instantaneousPower")
		}
		_instantaneousPowerErr := writeBuffer.WriteSerializable(m.InstantaneousPower)
		if popErr := writeBuffer.PopContext("instantaneousPower"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for instantaneousPower")
		}
		if _instantaneousPowerErr != nil {
			return errors.Wrap(_instantaneousPowerErr, "Error serializing 'instantaneousPower' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataInstantaneousPower"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataInstantaneousPower")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataInstantaneousPower) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
