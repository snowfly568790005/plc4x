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

// BACnetOpeningTag is the corresponding interface of BACnetOpeningTag
type BACnetOpeningTag interface {
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
}

// BACnetOpeningTagExactly can be used when we want exactly this type and not a type which fulfills BACnetOpeningTag.
// This is useful for switch cases.
type BACnetOpeningTagExactly interface {
	BACnetOpeningTag
	isBACnetOpeningTag() bool
}

// _BACnetOpeningTag is the data-structure of this message
type _BACnetOpeningTag struct {
	Header BACnetTagHeader

	// Arguments.
	TagNumberArgument uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetOpeningTag) GetHeader() BACnetTagHeader {
	return m.Header
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetOpeningTag factory function for _BACnetOpeningTag
func NewBACnetOpeningTag(header BACnetTagHeader, tagNumberArgument uint8) *_BACnetOpeningTag {
	return &_BACnetOpeningTag{Header: header, TagNumberArgument: tagNumberArgument}
}

// Deprecated: use the interface for direct cast
func CastBACnetOpeningTag(structType interface{}) BACnetOpeningTag {
	if casted, ok := structType.(BACnetOpeningTag); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetOpeningTag); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetOpeningTag) GetTypeName() string {
	return "BACnetOpeningTag"
}

func (m *_BACnetOpeningTag) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetOpeningTag) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetOpeningTag) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetOpeningTagParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8) (BACnetOpeningTag, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetOpeningTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOpeningTag")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParse(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetOpeningTag")
	}
	header := _header.(BACnetTagHeader)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if !(bool((header.GetActualTagNumber()) == (tagNumberArgument))) {
		return nil, errors.WithStack(utils.ParseAssertError{"tagnumber doesn't match"})
	}

	// Validation
	if !(bool((header.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS))) {
		return nil, errors.WithStack(utils.ParseValidationError{"should be a context tag"})
	}

	// Validation
	if !(bool((header.GetLengthValueType()) == (6))) {
		return nil, errors.WithStack(utils.ParseValidationError{"opening tag should have a value of 6"})
	}

	if closeErr := readBuffer.CloseContext("BACnetOpeningTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOpeningTag")
	}

	// Create the instance
	return &_BACnetOpeningTag{
		TagNumberArgument: tagNumberArgument,
		Header:            header,
	}, nil
}

func (m *_BACnetOpeningTag) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetOpeningTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetOpeningTag")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	if popErr := writeBuffer.PopContext("BACnetOpeningTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetOpeningTag")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetOpeningTag) GetTagNumberArgument() uint8 {
	return m.TagNumberArgument
}

//
////

func (m *_BACnetOpeningTag) isBACnetOpeningTag() bool {
	return true
}

func (m *_BACnetOpeningTag) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
