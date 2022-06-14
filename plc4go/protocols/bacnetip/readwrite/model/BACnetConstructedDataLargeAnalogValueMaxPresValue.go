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

// BACnetConstructedDataLargeAnalogValueMaxPresValue is the data-structure of this message
type BACnetConstructedDataLargeAnalogValueMaxPresValue struct {
	*BACnetConstructedData
	MaxPresValue *BACnetApplicationTagDouble

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataLargeAnalogValueMaxPresValue is the corresponding interface of BACnetConstructedDataLargeAnalogValueMaxPresValue
type IBACnetConstructedDataLargeAnalogValueMaxPresValue interface {
	IBACnetConstructedData
	// GetMaxPresValue returns MaxPresValue (property field)
	GetMaxPresValue() *BACnetApplicationTagDouble
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetApplicationTagDouble
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

func (m *BACnetConstructedDataLargeAnalogValueMaxPresValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LARGE_ANALOG_VALUE
}

func (m *BACnetConstructedDataLargeAnalogValueMaxPresValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_PRES_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLargeAnalogValueMaxPresValue) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLargeAnalogValueMaxPresValue) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLargeAnalogValueMaxPresValue) GetMaxPresValue() *BACnetApplicationTagDouble {
	return m.MaxPresValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataLargeAnalogValueMaxPresValue) GetActualValue() *BACnetApplicationTagDouble {
	return CastBACnetApplicationTagDouble(m.GetMaxPresValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLargeAnalogValueMaxPresValue factory function for BACnetConstructedDataLargeAnalogValueMaxPresValue
func NewBACnetConstructedDataLargeAnalogValueMaxPresValue(maxPresValue *BACnetApplicationTagDouble, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataLargeAnalogValueMaxPresValue {
	_result := &BACnetConstructedDataLargeAnalogValueMaxPresValue{
		MaxPresValue:          maxPresValue,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLargeAnalogValueMaxPresValue(structType interface{}) *BACnetConstructedDataLargeAnalogValueMaxPresValue {
	if casted, ok := structType.(BACnetConstructedDataLargeAnalogValueMaxPresValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLargeAnalogValueMaxPresValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLargeAnalogValueMaxPresValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLargeAnalogValueMaxPresValue(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLargeAnalogValueMaxPresValue) GetTypeName() string {
	return "BACnetConstructedDataLargeAnalogValueMaxPresValue"
}

func (m *BACnetConstructedDataLargeAnalogValueMaxPresValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLargeAnalogValueMaxPresValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (maxPresValue)
	lengthInBits += m.MaxPresValue.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataLargeAnalogValueMaxPresValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLargeAnalogValueMaxPresValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataLargeAnalogValueMaxPresValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLargeAnalogValueMaxPresValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLargeAnalogValueMaxPresValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maxPresValue)
	if pullErr := readBuffer.PullContext("maxPresValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maxPresValue")
	}
	_maxPresValue, _maxPresValueErr := BACnetApplicationTagParse(readBuffer)
	if _maxPresValueErr != nil {
		return nil, errors.Wrap(_maxPresValueErr, "Error parsing 'maxPresValue' field")
	}
	maxPresValue := CastBACnetApplicationTagDouble(_maxPresValue)
	if closeErr := readBuffer.CloseContext("maxPresValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maxPresValue")
	}

	// Virtual field
	_actualValue := maxPresValue
	actualValue := CastBACnetApplicationTagDouble(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLargeAnalogValueMaxPresValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLargeAnalogValueMaxPresValue")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLargeAnalogValueMaxPresValue{
		MaxPresValue:          CastBACnetApplicationTagDouble(maxPresValue),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLargeAnalogValueMaxPresValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLargeAnalogValueMaxPresValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLargeAnalogValueMaxPresValue")
		}

		// Simple Field (maxPresValue)
		if pushErr := writeBuffer.PushContext("maxPresValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maxPresValue")
		}
		_maxPresValueErr := writeBuffer.WriteSerializable(m.MaxPresValue)
		if popErr := writeBuffer.PopContext("maxPresValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maxPresValue")
		}
		if _maxPresValueErr != nil {
			return errors.Wrap(_maxPresValueErr, "Error serializing 'maxPresValue' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLargeAnalogValueMaxPresValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLargeAnalogValueMaxPresValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLargeAnalogValueMaxPresValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
