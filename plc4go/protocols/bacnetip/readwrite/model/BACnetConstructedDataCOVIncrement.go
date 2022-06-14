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

// BACnetConstructedDataCOVIncrement is the data-structure of this message
type BACnetConstructedDataCOVIncrement struct {
	*BACnetConstructedData
	CovIncrement *BACnetApplicationTagReal

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataCOVIncrement is the corresponding interface of BACnetConstructedDataCOVIncrement
type IBACnetConstructedDataCOVIncrement interface {
	IBACnetConstructedData
	// GetCovIncrement returns CovIncrement (property field)
	GetCovIncrement() *BACnetApplicationTagReal
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

func (m *BACnetConstructedDataCOVIncrement) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataCOVIncrement) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_COV_INCREMENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataCOVIncrement) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataCOVIncrement) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataCOVIncrement) GetCovIncrement() *BACnetApplicationTagReal {
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

func (m *BACnetConstructedDataCOVIncrement) GetActualValue() *BACnetApplicationTagReal {
	return CastBACnetApplicationTagReal(m.GetCovIncrement())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCOVIncrement factory function for BACnetConstructedDataCOVIncrement
func NewBACnetConstructedDataCOVIncrement(covIncrement *BACnetApplicationTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataCOVIncrement {
	_result := &BACnetConstructedDataCOVIncrement{
		CovIncrement:          covIncrement,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataCOVIncrement(structType interface{}) *BACnetConstructedDataCOVIncrement {
	if casted, ok := structType.(BACnetConstructedDataCOVIncrement); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCOVIncrement); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataCOVIncrement(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataCOVIncrement(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataCOVIncrement) GetTypeName() string {
	return "BACnetConstructedDataCOVIncrement"
}

func (m *BACnetConstructedDataCOVIncrement) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataCOVIncrement) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (covIncrement)
	lengthInBits += m.CovIncrement.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataCOVIncrement) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataCOVIncrementParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataCOVIncrement, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCOVIncrement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCOVIncrement")
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
	covIncrement := CastBACnetApplicationTagReal(_covIncrement)
	if closeErr := readBuffer.CloseContext("covIncrement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for covIncrement")
	}

	// Virtual field
	_actualValue := covIncrement
	actualValue := CastBACnetApplicationTagReal(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCOVIncrement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCOVIncrement")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataCOVIncrement{
		CovIncrement:          CastBACnetApplicationTagReal(covIncrement),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataCOVIncrement) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCOVIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCOVIncrement")
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

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCOVIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCOVIncrement")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataCOVIncrement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
