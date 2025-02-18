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

// BACnetConstructedDataBaseDeviceSecurityPolicy is the corresponding interface of BACnetConstructedDataBaseDeviceSecurityPolicy
type BACnetConstructedDataBaseDeviceSecurityPolicy interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetBaseDeviceSecurityPolicy returns BaseDeviceSecurityPolicy (property field)
	GetBaseDeviceSecurityPolicy() BACnetSecurityLevelTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetSecurityLevelTagged
}

// BACnetConstructedDataBaseDeviceSecurityPolicyExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBaseDeviceSecurityPolicy.
// This is useful for switch cases.
type BACnetConstructedDataBaseDeviceSecurityPolicyExactly interface {
	BACnetConstructedDataBaseDeviceSecurityPolicy
	isBACnetConstructedDataBaseDeviceSecurityPolicy() bool
}

// _BACnetConstructedDataBaseDeviceSecurityPolicy is the data-structure of this message
type _BACnetConstructedDataBaseDeviceSecurityPolicy struct {
	*_BACnetConstructedData
	BaseDeviceSecurityPolicy BACnetSecurityLevelTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BASE_DEVICE_SECURITY_POLICY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) GetBaseDeviceSecurityPolicy() BACnetSecurityLevelTagged {
	return m.BaseDeviceSecurityPolicy
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) GetActualValue() BACnetSecurityLevelTagged {
	return CastBACnetSecurityLevelTagged(m.GetBaseDeviceSecurityPolicy())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBaseDeviceSecurityPolicy factory function for _BACnetConstructedDataBaseDeviceSecurityPolicy
func NewBACnetConstructedDataBaseDeviceSecurityPolicy(baseDeviceSecurityPolicy BACnetSecurityLevelTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBaseDeviceSecurityPolicy {
	_result := &_BACnetConstructedDataBaseDeviceSecurityPolicy{
		BaseDeviceSecurityPolicy: baseDeviceSecurityPolicy,
		_BACnetConstructedData:   NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBaseDeviceSecurityPolicy(structType interface{}) BACnetConstructedDataBaseDeviceSecurityPolicy {
	if casted, ok := structType.(BACnetConstructedDataBaseDeviceSecurityPolicy); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBaseDeviceSecurityPolicy); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) GetTypeName() string {
	return "BACnetConstructedDataBaseDeviceSecurityPolicy"
}

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (baseDeviceSecurityPolicy)
	lengthInBits += m.BaseDeviceSecurityPolicy.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataBaseDeviceSecurityPolicyParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBaseDeviceSecurityPolicy, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBaseDeviceSecurityPolicy"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBaseDeviceSecurityPolicy")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (baseDeviceSecurityPolicy)
	if pullErr := readBuffer.PullContext("baseDeviceSecurityPolicy"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for baseDeviceSecurityPolicy")
	}
	_baseDeviceSecurityPolicy, _baseDeviceSecurityPolicyErr := BACnetSecurityLevelTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _baseDeviceSecurityPolicyErr != nil {
		return nil, errors.Wrap(_baseDeviceSecurityPolicyErr, "Error parsing 'baseDeviceSecurityPolicy' field of BACnetConstructedDataBaseDeviceSecurityPolicy")
	}
	baseDeviceSecurityPolicy := _baseDeviceSecurityPolicy.(BACnetSecurityLevelTagged)
	if closeErr := readBuffer.CloseContext("baseDeviceSecurityPolicy"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for baseDeviceSecurityPolicy")
	}

	// Virtual field
	_actualValue := baseDeviceSecurityPolicy
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBaseDeviceSecurityPolicy"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBaseDeviceSecurityPolicy")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataBaseDeviceSecurityPolicy{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		BaseDeviceSecurityPolicy: baseDeviceSecurityPolicy,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBaseDeviceSecurityPolicy"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBaseDeviceSecurityPolicy")
		}

		// Simple Field (baseDeviceSecurityPolicy)
		if pushErr := writeBuffer.PushContext("baseDeviceSecurityPolicy"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for baseDeviceSecurityPolicy")
		}
		_baseDeviceSecurityPolicyErr := writeBuffer.WriteSerializable(m.GetBaseDeviceSecurityPolicy())
		if popErr := writeBuffer.PopContext("baseDeviceSecurityPolicy"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for baseDeviceSecurityPolicy")
		}
		if _baseDeviceSecurityPolicyErr != nil {
			return errors.Wrap(_baseDeviceSecurityPolicyErr, "Error serializing 'baseDeviceSecurityPolicy' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBaseDeviceSecurityPolicy"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBaseDeviceSecurityPolicy")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) isBACnetConstructedDataBaseDeviceSecurityPolicy() bool {
	return true
}

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
