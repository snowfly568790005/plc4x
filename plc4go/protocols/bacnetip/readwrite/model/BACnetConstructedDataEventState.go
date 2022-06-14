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

// BACnetConstructedDataEventState is the data-structure of this message
type BACnetConstructedDataEventState struct {
	*BACnetConstructedData
	EventState *BACnetEventStateTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataEventState is the corresponding interface of BACnetConstructedDataEventState
type IBACnetConstructedDataEventState interface {
	IBACnetConstructedData
	// GetEventState returns EventState (property field)
	GetEventState() *BACnetEventStateTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetEventStateTagged
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

func (m *BACnetConstructedDataEventState) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataEventState) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_STATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataEventState) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataEventState) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataEventState) GetEventState() *BACnetEventStateTagged {
	return m.EventState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataEventState) GetActualValue() *BACnetEventStateTagged {
	return CastBACnetEventStateTagged(m.GetEventState())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEventState factory function for BACnetConstructedDataEventState
func NewBACnetConstructedDataEventState(eventState *BACnetEventStateTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataEventState {
	_result := &BACnetConstructedDataEventState{
		EventState:            eventState,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataEventState(structType interface{}) *BACnetConstructedDataEventState {
	if casted, ok := structType.(BACnetConstructedDataEventState); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventState); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataEventState(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataEventState(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataEventState) GetTypeName() string {
	return "BACnetConstructedDataEventState"
}

func (m *BACnetConstructedDataEventState) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataEventState) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (eventState)
	lengthInBits += m.EventState.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataEventState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataEventStateParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataEventState, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (eventState)
	if pullErr := readBuffer.PullContext("eventState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventState")
	}
	_eventState, _eventStateErr := BACnetEventStateTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _eventStateErr != nil {
		return nil, errors.Wrap(_eventStateErr, "Error parsing 'eventState' field")
	}
	eventState := CastBACnetEventStateTagged(_eventState)
	if closeErr := readBuffer.CloseContext("eventState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventState")
	}

	// Virtual field
	_actualValue := eventState
	actualValue := CastBACnetEventStateTagged(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventState")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataEventState{
		EventState:            CastBACnetEventStateTagged(eventState),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataEventState) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventState")
		}

		// Simple Field (eventState)
		if pushErr := writeBuffer.PushContext("eventState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for eventState")
		}
		_eventStateErr := writeBuffer.WriteSerializable(m.EventState)
		if popErr := writeBuffer.PopContext("eventState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for eventState")
		}
		if _eventStateErr != nil {
			return errors.Wrap(_eventStateErr, "Error serializing 'eventState' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventState")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataEventState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
