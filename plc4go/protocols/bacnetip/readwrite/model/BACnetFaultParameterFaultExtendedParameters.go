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

// BACnetFaultParameterFaultExtendedParameters is the corresponding interface of BACnetFaultParameterFaultExtendedParameters
type BACnetFaultParameterFaultExtendedParameters interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetParameters returns Parameters (property field)
	GetParameters() []BACnetFaultParameterFaultExtendedParametersEntry
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetFaultParameterFaultExtendedParametersExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultExtendedParameters.
// This is useful for switch cases.
type BACnetFaultParameterFaultExtendedParametersExactly interface {
	BACnetFaultParameterFaultExtendedParameters
	isBACnetFaultParameterFaultExtendedParameters() bool
}

// _BACnetFaultParameterFaultExtendedParameters is the data-structure of this message
type _BACnetFaultParameterFaultExtendedParameters struct {
	OpeningTag BACnetOpeningTag
	Parameters []BACnetFaultParameterFaultExtendedParametersEntry
	ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtendedParameters) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultExtendedParameters) GetParameters() []BACnetFaultParameterFaultExtendedParametersEntry {
	return m.Parameters
}

func (m *_BACnetFaultParameterFaultExtendedParameters) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultExtendedParameters factory function for _BACnetFaultParameterFaultExtendedParameters
func NewBACnetFaultParameterFaultExtendedParameters(openingTag BACnetOpeningTag, parameters []BACnetFaultParameterFaultExtendedParametersEntry, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetFaultParameterFaultExtendedParameters {
	return &_BACnetFaultParameterFaultExtendedParameters{OpeningTag: openingTag, Parameters: parameters, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtendedParameters(structType interface{}) BACnetFaultParameterFaultExtendedParameters {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParameters); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParameters); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParameters) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParameters"
}

func (m *_BACnetFaultParameterFaultExtendedParameters) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetFaultParameterFaultExtendedParameters) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Array field
	if len(m.Parameters) > 0 {
		for _, element := range m.Parameters {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultExtendedParameters) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultExtendedParametersParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetFaultParameterFaultExtendedParameters, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParameters"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParameters")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetFaultParameterFaultExtendedParameters")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Array field (parameters)
	if pullErr := readBuffer.PullContext("parameters", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for parameters")
	}
	// Terminated array
	var parameters []BACnetFaultParameterFaultExtendedParametersEntry
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetFaultParameterFaultExtendedParametersEntryParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'parameters' field of BACnetFaultParameterFaultExtendedParameters")
			}
			parameters = append(parameters, _item.(BACnetFaultParameterFaultExtendedParametersEntry))

		}
	}
	if closeErr := readBuffer.CloseContext("parameters", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for parameters")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetFaultParameterFaultExtendedParameters")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParameters"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParameters")
	}

	// Create the instance
	return &_BACnetFaultParameterFaultExtendedParameters{
		TagNumber:  tagNumber,
		OpeningTag: openingTag,
		Parameters: parameters,
		ClosingTag: closingTag,
	}, nil
}

func (m *_BACnetFaultParameterFaultExtendedParameters) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParameters"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParameters")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Array Field (parameters)
	if pushErr := writeBuffer.PushContext("parameters", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for parameters")
	}
	for _, _element := range m.GetParameters() {
		_elementErr := writeBuffer.WriteSerializable(_element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'parameters' field")
		}
	}
	if popErr := writeBuffer.PopContext("parameters", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for parameters")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParameters"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParameters")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetFaultParameterFaultExtendedParameters) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetFaultParameterFaultExtendedParameters) isBACnetFaultParameterFaultExtendedParameters() bool {
	return true
}

func (m *_BACnetFaultParameterFaultExtendedParameters) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
