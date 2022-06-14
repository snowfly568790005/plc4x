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

// BACnetConstructedDataActiveText is the data-structure of this message
type BACnetConstructedDataActiveText struct {
	*BACnetConstructedData
	ActiveText *BACnetApplicationTagCharacterString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataActiveText is the corresponding interface of BACnetConstructedDataActiveText
type IBACnetConstructedDataActiveText interface {
	IBACnetConstructedData
	// GetActiveText returns ActiveText (property field)
	GetActiveText() *BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetApplicationTagCharacterString
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

func (m *BACnetConstructedDataActiveText) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataActiveText) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACTIVE_TEXT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataActiveText) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataActiveText) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataActiveText) GetActiveText() *BACnetApplicationTagCharacterString {
	return m.ActiveText
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataActiveText) GetActualValue() *BACnetApplicationTagCharacterString {
	return CastBACnetApplicationTagCharacterString(m.GetActiveText())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataActiveText factory function for BACnetConstructedDataActiveText
func NewBACnetConstructedDataActiveText(activeText *BACnetApplicationTagCharacterString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataActiveText {
	_result := &BACnetConstructedDataActiveText{
		ActiveText:            activeText,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataActiveText(structType interface{}) *BACnetConstructedDataActiveText {
	if casted, ok := structType.(BACnetConstructedDataActiveText); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataActiveText); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataActiveText(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataActiveText(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataActiveText) GetTypeName() string {
	return "BACnetConstructedDataActiveText"
}

func (m *BACnetConstructedDataActiveText) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataActiveText) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (activeText)
	lengthInBits += m.ActiveText.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataActiveText) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataActiveTextParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataActiveText, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataActiveText"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataActiveText")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (activeText)
	if pullErr := readBuffer.PullContext("activeText"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for activeText")
	}
	_activeText, _activeTextErr := BACnetApplicationTagParse(readBuffer)
	if _activeTextErr != nil {
		return nil, errors.Wrap(_activeTextErr, "Error parsing 'activeText' field")
	}
	activeText := CastBACnetApplicationTagCharacterString(_activeText)
	if closeErr := readBuffer.CloseContext("activeText"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for activeText")
	}

	// Virtual field
	_actualValue := activeText
	actualValue := CastBACnetApplicationTagCharacterString(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataActiveText"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataActiveText")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataActiveText{
		ActiveText:            CastBACnetApplicationTagCharacterString(activeText),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataActiveText) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataActiveText"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataActiveText")
		}

		// Simple Field (activeText)
		if pushErr := writeBuffer.PushContext("activeText"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for activeText")
		}
		_activeTextErr := writeBuffer.WriteSerializable(m.ActiveText)
		if popErr := writeBuffer.PopContext("activeText"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for activeText")
		}
		if _activeTextErr != nil {
			return errors.Wrap(_activeTextErr, "Error serializing 'activeText' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataActiveText"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataActiveText")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataActiveText) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
