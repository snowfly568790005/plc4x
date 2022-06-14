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

// BACnetConstructedDataLightingOutputFeedbackValue is the data-structure of this message
type BACnetConstructedDataLightingOutputFeedbackValue struct {
	*BACnetConstructedData
	FeedbackValue *BACnetApplicationTagReal

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataLightingOutputFeedbackValue is the corresponding interface of BACnetConstructedDataLightingOutputFeedbackValue
type IBACnetConstructedDataLightingOutputFeedbackValue interface {
	IBACnetConstructedData
	// GetFeedbackValue returns FeedbackValue (property field)
	GetFeedbackValue() *BACnetApplicationTagReal
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

func (m *BACnetConstructedDataLightingOutputFeedbackValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LIGHTING_OUTPUT
}

func (m *BACnetConstructedDataLightingOutputFeedbackValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FEEDBACK_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLightingOutputFeedbackValue) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLightingOutputFeedbackValue) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLightingOutputFeedbackValue) GetFeedbackValue() *BACnetApplicationTagReal {
	return m.FeedbackValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataLightingOutputFeedbackValue) GetActualValue() *BACnetApplicationTagReal {
	return CastBACnetApplicationTagReal(m.GetFeedbackValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLightingOutputFeedbackValue factory function for BACnetConstructedDataLightingOutputFeedbackValue
func NewBACnetConstructedDataLightingOutputFeedbackValue(feedbackValue *BACnetApplicationTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataLightingOutputFeedbackValue {
	_result := &BACnetConstructedDataLightingOutputFeedbackValue{
		FeedbackValue:         feedbackValue,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLightingOutputFeedbackValue(structType interface{}) *BACnetConstructedDataLightingOutputFeedbackValue {
	if casted, ok := structType.(BACnetConstructedDataLightingOutputFeedbackValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLightingOutputFeedbackValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLightingOutputFeedbackValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLightingOutputFeedbackValue(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLightingOutputFeedbackValue) GetTypeName() string {
	return "BACnetConstructedDataLightingOutputFeedbackValue"
}

func (m *BACnetConstructedDataLightingOutputFeedbackValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLightingOutputFeedbackValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (feedbackValue)
	lengthInBits += m.FeedbackValue.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataLightingOutputFeedbackValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLightingOutputFeedbackValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataLightingOutputFeedbackValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLightingOutputFeedbackValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLightingOutputFeedbackValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (feedbackValue)
	if pullErr := readBuffer.PullContext("feedbackValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for feedbackValue")
	}
	_feedbackValue, _feedbackValueErr := BACnetApplicationTagParse(readBuffer)
	if _feedbackValueErr != nil {
		return nil, errors.Wrap(_feedbackValueErr, "Error parsing 'feedbackValue' field")
	}
	feedbackValue := CastBACnetApplicationTagReal(_feedbackValue)
	if closeErr := readBuffer.CloseContext("feedbackValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for feedbackValue")
	}

	// Virtual field
	_actualValue := feedbackValue
	actualValue := CastBACnetApplicationTagReal(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLightingOutputFeedbackValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLightingOutputFeedbackValue")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLightingOutputFeedbackValue{
		FeedbackValue:         CastBACnetApplicationTagReal(feedbackValue),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLightingOutputFeedbackValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLightingOutputFeedbackValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLightingOutputFeedbackValue")
		}

		// Simple Field (feedbackValue)
		if pushErr := writeBuffer.PushContext("feedbackValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for feedbackValue")
		}
		_feedbackValueErr := writeBuffer.WriteSerializable(m.FeedbackValue)
		if popErr := writeBuffer.PopContext("feedbackValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for feedbackValue")
		}
		if _feedbackValueErr != nil {
			return errors.Wrap(_feedbackValueErr, "Error serializing 'feedbackValue' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLightingOutputFeedbackValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLightingOutputFeedbackValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLightingOutputFeedbackValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
