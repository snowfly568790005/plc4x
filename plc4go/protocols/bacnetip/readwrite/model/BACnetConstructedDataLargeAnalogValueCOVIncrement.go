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

// BACnetConstructedDataLargeAnalogValueCOVIncrement is the data-structure of this message
type BACnetConstructedDataLargeAnalogValueCOVIncrement struct {
	*BACnetConstructedData
	CovIncrement *BACnetApplicationTagDouble

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataLargeAnalogValueCOVIncrement is the corresponding interface of BACnetConstructedDataLargeAnalogValueCOVIncrement
type IBACnetConstructedDataLargeAnalogValueCOVIncrement interface {
	IBACnetConstructedData
	// GetCovIncrement returns CovIncrement (property field)
	GetCovIncrement() *BACnetApplicationTagDouble
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

func (m *BACnetConstructedDataLargeAnalogValueCOVIncrement) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LARGE_ANALOG_VALUE
}

func (m *BACnetConstructedDataLargeAnalogValueCOVIncrement) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_COV_INCREMENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLargeAnalogValueCOVIncrement) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLargeAnalogValueCOVIncrement) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLargeAnalogValueCOVIncrement) GetCovIncrement() *BACnetApplicationTagDouble {
	return m.CovIncrement
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataLargeAnalogValueCOVIncrement) GetActualValue() *BACnetApplicationTagDouble {
	return CastBACnetApplicationTagDouble(m.GetCovIncrement())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLargeAnalogValueCOVIncrement factory function for BACnetConstructedDataLargeAnalogValueCOVIncrement
func NewBACnetConstructedDataLargeAnalogValueCOVIncrement(covIncrement *BACnetApplicationTagDouble, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataLargeAnalogValueCOVIncrement {
	_result := &BACnetConstructedDataLargeAnalogValueCOVIncrement{
		CovIncrement:          covIncrement,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLargeAnalogValueCOVIncrement(structType interface{}) *BACnetConstructedDataLargeAnalogValueCOVIncrement {
	if casted, ok := structType.(BACnetConstructedDataLargeAnalogValueCOVIncrement); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLargeAnalogValueCOVIncrement); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLargeAnalogValueCOVIncrement(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLargeAnalogValueCOVIncrement(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLargeAnalogValueCOVIncrement) GetTypeName() string {
	return "BACnetConstructedDataLargeAnalogValueCOVIncrement"
}

func (m *BACnetConstructedDataLargeAnalogValueCOVIncrement) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLargeAnalogValueCOVIncrement) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (covIncrement)
	lengthInBits += m.CovIncrement.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataLargeAnalogValueCOVIncrement) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLargeAnalogValueCOVIncrementParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataLargeAnalogValueCOVIncrement, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLargeAnalogValueCOVIncrement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLargeAnalogValueCOVIncrement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (covIncrement)
	if pullErr := readBuffer.PullContext("covIncrement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for covIncrement")
	}
	_covIncrement, _covIncrementErr := BACnetApplicationTagParse(readBuffer)
	if _covIncrementErr != nil {
		return nil, errors.Wrap(_covIncrementErr, "Error parsing 'covIncrement' field")
	}
	covIncrement := CastBACnetApplicationTagDouble(_covIncrement)
	if closeErr := readBuffer.CloseContext("covIncrement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for covIncrement")
	}

	// Virtual field
	_actualValue := covIncrement
	actualValue := CastBACnetApplicationTagDouble(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLargeAnalogValueCOVIncrement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLargeAnalogValueCOVIncrement")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLargeAnalogValueCOVIncrement{
		CovIncrement:          CastBACnetApplicationTagDouble(covIncrement),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLargeAnalogValueCOVIncrement) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLargeAnalogValueCOVIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLargeAnalogValueCOVIncrement")
		}

		// Simple Field (covIncrement)
		if pushErr := writeBuffer.PushContext("covIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for covIncrement")
		}
		_covIncrementErr := writeBuffer.WriteSerializable(m.CovIncrement)
		if popErr := writeBuffer.PopContext("covIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for covIncrement")
		}
		if _covIncrementErr != nil {
			return errors.Wrap(_covIncrementErr, "Error serializing 'covIncrement' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLargeAnalogValueCOVIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLargeAnalogValueCOVIncrement")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLargeAnalogValueCOVIncrement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
