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

// BACnetConstructedDataMaximumValue is the data-structure of this message
type BACnetConstructedDataMaximumValue struct {
	*BACnetConstructedData
	MaximumValue *BACnetApplicationTagReal

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataMaximumValue is the corresponding interface of BACnetConstructedDataMaximumValue
type IBACnetConstructedDataMaximumValue interface {
	IBACnetConstructedData
	// GetMaximumValue returns MaximumValue (property field)
	GetMaximumValue() *BACnetApplicationTagReal
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

func (m *BACnetConstructedDataMaximumValue) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataMaximumValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAXIMUM_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataMaximumValue) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataMaximumValue) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataMaximumValue) GetMaximumValue() *BACnetApplicationTagReal {
	return m.MaximumValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataMaximumValue) GetActualValue() *BACnetApplicationTagReal {
	return CastBACnetApplicationTagReal(m.GetMaximumValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMaximumValue factory function for BACnetConstructedDataMaximumValue
func NewBACnetConstructedDataMaximumValue(maximumValue *BACnetApplicationTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataMaximumValue {
	_result := &BACnetConstructedDataMaximumValue{
		MaximumValue:          maximumValue,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataMaximumValue(structType interface{}) *BACnetConstructedDataMaximumValue {
	if casted, ok := structType.(BACnetConstructedDataMaximumValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaximumValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataMaximumValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataMaximumValue(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataMaximumValue) GetTypeName() string {
	return "BACnetConstructedDataMaximumValue"
}

func (m *BACnetConstructedDataMaximumValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataMaximumValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (maximumValue)
	lengthInBits += m.MaximumValue.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataMaximumValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMaximumValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataMaximumValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaximumValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaximumValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maximumValue)
	if pullErr := readBuffer.PullContext("maximumValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maximumValue")
	}
	_maximumValue, _maximumValueErr := BACnetApplicationTagParse(readBuffer)
	if _maximumValueErr != nil {
		return nil, errors.Wrap(_maximumValueErr, "Error parsing 'maximumValue' field")
	}
	maximumValue := CastBACnetApplicationTagReal(_maximumValue)
	if closeErr := readBuffer.CloseContext("maximumValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maximumValue")
	}

	// Virtual field
	_actualValue := maximumValue
	actualValue := CastBACnetApplicationTagReal(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaximumValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaximumValue")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataMaximumValue{
		MaximumValue:          CastBACnetApplicationTagReal(maximumValue),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataMaximumValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaximumValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaximumValue")
		}

		// Simple Field (maximumValue)
		if pushErr := writeBuffer.PushContext("maximumValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maximumValue")
		}
		_maximumValueErr := writeBuffer.WriteSerializable(m.MaximumValue)
		if popErr := writeBuffer.PopContext("maximumValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maximumValue")
		}
		if _maximumValueErr != nil {
			return errors.Wrap(_maximumValueErr, "Error serializing 'maximumValue' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaximumValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaximumValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataMaximumValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
