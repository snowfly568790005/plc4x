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

// NLMICouldBeRouterToNetwork is the corresponding interface of NLMICouldBeRouterToNetwork
type NLMICouldBeRouterToNetwork interface {
	utils.LengthAware
	utils.Serializable
	NLM
	// GetDestinationNetworkAddress returns DestinationNetworkAddress (property field)
	GetDestinationNetworkAddress() uint16
	// GetPerformanceIndex returns PerformanceIndex (property field)
	GetPerformanceIndex() uint8
}

// NLMICouldBeRouterToNetworkExactly can be used when we want exactly this type and not a type which fulfills NLMICouldBeRouterToNetwork.
// This is useful for switch cases.
type NLMICouldBeRouterToNetworkExactly interface {
	NLMICouldBeRouterToNetwork
	isNLMICouldBeRouterToNetwork() bool
}

// _NLMICouldBeRouterToNetwork is the data-structure of this message
type _NLMICouldBeRouterToNetwork struct {
	*_NLM
	DestinationNetworkAddress uint16
	PerformanceIndex          uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMICouldBeRouterToNetwork) GetMessageType() uint8 {
	return 0x02
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMICouldBeRouterToNetwork) InitializeParent(parent NLM, vendorId *BACnetVendorId) {
	m.VendorId = vendorId
}

func (m *_NLMICouldBeRouterToNetwork) GetParent() NLM {
	return m._NLM
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMICouldBeRouterToNetwork) GetDestinationNetworkAddress() uint16 {
	return m.DestinationNetworkAddress
}

func (m *_NLMICouldBeRouterToNetwork) GetPerformanceIndex() uint8 {
	return m.PerformanceIndex
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMICouldBeRouterToNetwork factory function for _NLMICouldBeRouterToNetwork
func NewNLMICouldBeRouterToNetwork(destinationNetworkAddress uint16, performanceIndex uint8, vendorId *BACnetVendorId, apduLength uint16) *_NLMICouldBeRouterToNetwork {
	_result := &_NLMICouldBeRouterToNetwork{
		DestinationNetworkAddress: destinationNetworkAddress,
		PerformanceIndex:          performanceIndex,
		_NLM:                      NewNLM(vendorId, apduLength),
	}
	_result._NLM._NLMChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNLMICouldBeRouterToNetwork(structType interface{}) NLMICouldBeRouterToNetwork {
	if casted, ok := structType.(NLMICouldBeRouterToNetwork); ok {
		return casted
	}
	if casted, ok := structType.(*NLMICouldBeRouterToNetwork); ok {
		return *casted
	}
	return nil
}

func (m *_NLMICouldBeRouterToNetwork) GetTypeName() string {
	return "NLMICouldBeRouterToNetwork"
}

func (m *_NLMICouldBeRouterToNetwork) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_NLMICouldBeRouterToNetwork) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (destinationNetworkAddress)
	lengthInBits += 16

	// Simple field (performanceIndex)
	lengthInBits += 8

	return lengthInBits
}

func (m *_NLMICouldBeRouterToNetwork) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NLMICouldBeRouterToNetworkParse(readBuffer utils.ReadBuffer, apduLength uint16, messageType uint8) (NLMICouldBeRouterToNetwork, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMICouldBeRouterToNetwork"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMICouldBeRouterToNetwork")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (destinationNetworkAddress)
	_destinationNetworkAddress, _destinationNetworkAddressErr := readBuffer.ReadUint16("destinationNetworkAddress", 16)
	if _destinationNetworkAddressErr != nil {
		return nil, errors.Wrap(_destinationNetworkAddressErr, "Error parsing 'destinationNetworkAddress' field of NLMICouldBeRouterToNetwork")
	}
	destinationNetworkAddress := _destinationNetworkAddress

	// Simple Field (performanceIndex)
	_performanceIndex, _performanceIndexErr := readBuffer.ReadUint8("performanceIndex", 8)
	if _performanceIndexErr != nil {
		return nil, errors.Wrap(_performanceIndexErr, "Error parsing 'performanceIndex' field of NLMICouldBeRouterToNetwork")
	}
	performanceIndex := _performanceIndex

	if closeErr := readBuffer.CloseContext("NLMICouldBeRouterToNetwork"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMICouldBeRouterToNetwork")
	}

	// Create a partially initialized instance
	_child := &_NLMICouldBeRouterToNetwork{
		_NLM: &_NLM{
			ApduLength: apduLength,
		},
		DestinationNetworkAddress: destinationNetworkAddress,
		PerformanceIndex:          performanceIndex,
	}
	_child._NLM._NLMChildRequirements = _child
	return _child, nil
}

func (m *_NLMICouldBeRouterToNetwork) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMICouldBeRouterToNetwork"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMICouldBeRouterToNetwork")
		}

		// Simple Field (destinationNetworkAddress)
		destinationNetworkAddress := uint16(m.GetDestinationNetworkAddress())
		_destinationNetworkAddressErr := writeBuffer.WriteUint16("destinationNetworkAddress", 16, (destinationNetworkAddress))
		if _destinationNetworkAddressErr != nil {
			return errors.Wrap(_destinationNetworkAddressErr, "Error serializing 'destinationNetworkAddress' field")
		}

		// Simple Field (performanceIndex)
		performanceIndex := uint8(m.GetPerformanceIndex())
		_performanceIndexErr := writeBuffer.WriteUint8("performanceIndex", 8, (performanceIndex))
		if _performanceIndexErr != nil {
			return errors.Wrap(_performanceIndexErr, "Error serializing 'performanceIndex' field")
		}

		if popErr := writeBuffer.PopContext("NLMICouldBeRouterToNetwork"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMICouldBeRouterToNetwork")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_NLMICouldBeRouterToNetwork) isNLMICouldBeRouterToNetwork() bool {
	return true
}

func (m *_NLMICouldBeRouterToNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
