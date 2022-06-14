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

// BACnetConstructedDataAckRequired is the data-structure of this message
type BACnetConstructedDataAckRequired struct {
	*BACnetConstructedData
	AckRequired *BACnetEventTransitionBitsTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataAckRequired is the corresponding interface of BACnetConstructedDataAckRequired
type IBACnetConstructedDataAckRequired interface {
	IBACnetConstructedData
	// GetAckRequired returns AckRequired (property field)
	GetAckRequired() *BACnetEventTransitionBitsTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetEventTransitionBitsTagged
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

func (m *BACnetConstructedDataAckRequired) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataAckRequired) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACK_REQUIRED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataAckRequired) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataAckRequired) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataAckRequired) GetAckRequired() *BACnetEventTransitionBitsTagged {
	return m.AckRequired
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataAckRequired) GetActualValue() *BACnetEventTransitionBitsTagged {
	return CastBACnetEventTransitionBitsTagged(m.GetAckRequired())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAckRequired factory function for BACnetConstructedDataAckRequired
func NewBACnetConstructedDataAckRequired(ackRequired *BACnetEventTransitionBitsTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataAckRequired {
	_result := &BACnetConstructedDataAckRequired{
		AckRequired:           ackRequired,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataAckRequired(structType interface{}) *BACnetConstructedDataAckRequired {
	if casted, ok := structType.(BACnetConstructedDataAckRequired); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAckRequired); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataAckRequired(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataAckRequired(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataAckRequired) GetTypeName() string {
	return "BACnetConstructedDataAckRequired"
}

func (m *BACnetConstructedDataAckRequired) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataAckRequired) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (ackRequired)
	lengthInBits += m.AckRequired.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataAckRequired) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAckRequiredParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataAckRequired, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAckRequired"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAckRequired")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ackRequired)
	if pullErr := readBuffer.PullContext("ackRequired"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ackRequired")
	}
	_ackRequired, _ackRequiredErr := BACnetEventTransitionBitsTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _ackRequiredErr != nil {
		return nil, errors.Wrap(_ackRequiredErr, "Error parsing 'ackRequired' field")
	}
	ackRequired := CastBACnetEventTransitionBitsTagged(_ackRequired)
	if closeErr := readBuffer.CloseContext("ackRequired"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ackRequired")
	}

	// Virtual field
	_actualValue := ackRequired
	actualValue := CastBACnetEventTransitionBitsTagged(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAckRequired"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAckRequired")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataAckRequired{
		AckRequired:           CastBACnetEventTransitionBitsTagged(ackRequired),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataAckRequired) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAckRequired"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAckRequired")
		}

		// Simple Field (ackRequired)
		if pushErr := writeBuffer.PushContext("ackRequired"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ackRequired")
		}
		_ackRequiredErr := writeBuffer.WriteSerializable(m.AckRequired)
		if popErr := writeBuffer.PopContext("ackRequired"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ackRequired")
		}
		if _ackRequiredErr != nil {
			return errors.Wrap(_ackRequiredErr, "Error serializing 'ackRequired' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAckRequired"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAckRequired")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataAckRequired) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
