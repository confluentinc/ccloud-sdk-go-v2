// Copyright 2021 Confluent Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Network API

Network Gateway API

API version: 0.0.1-alpha1
Contact: cire-traffic@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// NetworkingV1GatewayStatusCloudGatewayOneOf - struct for NetworkingV1GatewayStatusCloudGatewayOneOf
type NetworkingV1GatewayStatusCloudGatewayOneOf struct {
	NetworkingV1AwsEgressPrivateLinkGatewayStatus           *NetworkingV1AwsEgressPrivateLinkGatewayStatus
	NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus     *NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus
	NetworkingV1AzureEgressPrivateLinkGatewayStatus         *NetworkingV1AzureEgressPrivateLinkGatewayStatus
	NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus *NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus
	NetworkingV1GcpPeeringGatewayStatus                     *NetworkingV1GcpPeeringGatewayStatus
}

// NetworkingV1AwsEgressPrivateLinkGatewayStatusAsNetworkingV1GatewayStatusCloudGatewayOneOf is a convenience function that returns NetworkingV1AwsEgressPrivateLinkGatewayStatus wrapped in NetworkingV1GatewayStatusCloudGatewayOneOf
func NetworkingV1AwsEgressPrivateLinkGatewayStatusAsNetworkingV1GatewayStatusCloudGatewayOneOf(v *NetworkingV1AwsEgressPrivateLinkGatewayStatus) NetworkingV1GatewayStatusCloudGatewayOneOf {
	return NetworkingV1GatewayStatusCloudGatewayOneOf{NetworkingV1AwsEgressPrivateLinkGatewayStatus: v}
}

// NetworkingV1AwsPrivateNetworkInterfaceGatewayStatusAsNetworkingV1GatewayStatusCloudGatewayOneOf is a convenience function that returns NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus wrapped in NetworkingV1GatewayStatusCloudGatewayOneOf
func NetworkingV1AwsPrivateNetworkInterfaceGatewayStatusAsNetworkingV1GatewayStatusCloudGatewayOneOf(v *NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus) NetworkingV1GatewayStatusCloudGatewayOneOf {
	return NetworkingV1GatewayStatusCloudGatewayOneOf{NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus: v}
}

// NetworkingV1AzureEgressPrivateLinkGatewayStatusAsNetworkingV1GatewayStatusCloudGatewayOneOf is a convenience function that returns NetworkingV1AzureEgressPrivateLinkGatewayStatus wrapped in NetworkingV1GatewayStatusCloudGatewayOneOf
func NetworkingV1AzureEgressPrivateLinkGatewayStatusAsNetworkingV1GatewayStatusCloudGatewayOneOf(v *NetworkingV1AzureEgressPrivateLinkGatewayStatus) NetworkingV1GatewayStatusCloudGatewayOneOf {
	return NetworkingV1GatewayStatusCloudGatewayOneOf{NetworkingV1AzureEgressPrivateLinkGatewayStatus: v}
}

// NetworkingV1GcpEgressPrivateServiceConnectGatewayStatusAsNetworkingV1GatewayStatusCloudGatewayOneOf is a convenience function that returns NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus wrapped in NetworkingV1GatewayStatusCloudGatewayOneOf
func NetworkingV1GcpEgressPrivateServiceConnectGatewayStatusAsNetworkingV1GatewayStatusCloudGatewayOneOf(v *NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus) NetworkingV1GatewayStatusCloudGatewayOneOf {
	return NetworkingV1GatewayStatusCloudGatewayOneOf{NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus: v}
}

// NetworkingV1GcpPeeringGatewayStatusAsNetworkingV1GatewayStatusCloudGatewayOneOf is a convenience function that returns NetworkingV1GcpPeeringGatewayStatus wrapped in NetworkingV1GatewayStatusCloudGatewayOneOf
func NetworkingV1GcpPeeringGatewayStatusAsNetworkingV1GatewayStatusCloudGatewayOneOf(v *NetworkingV1GcpPeeringGatewayStatus) NetworkingV1GatewayStatusCloudGatewayOneOf {
	return NetworkingV1GatewayStatusCloudGatewayOneOf{NetworkingV1GcpPeeringGatewayStatus: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NetworkingV1GatewayStatusCloudGatewayOneOf) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'AwsEgressPrivateLinkGatewayStatus'
	if jsonDict["kind"] == "AwsEgressPrivateLinkGatewayStatus" {
		// try to unmarshal JSON data into NetworkingV1AwsEgressPrivateLinkGatewayStatus
		err = json.Unmarshal(data, &dst.NetworkingV1AwsEgressPrivateLinkGatewayStatus)
		if err == nil {
			return nil // data stored in dst.NetworkingV1AwsEgressPrivateLinkGatewayStatus, return on the first match
		} else {
			dst.NetworkingV1AwsEgressPrivateLinkGatewayStatus = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1GatewayStatusCloudGatewayOneOf as NetworkingV1AwsEgressPrivateLinkGatewayStatus: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AwsPrivateNetworkInterfaceGatewayStatus'
	if jsonDict["kind"] == "AwsPrivateNetworkInterfaceGatewayStatus" {
		// try to unmarshal JSON data into NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus
		err = json.Unmarshal(data, &dst.NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus)
		if err == nil {
			return nil // data stored in dst.NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus, return on the first match
		} else {
			dst.NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1GatewayStatusCloudGatewayOneOf as NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AzureEgressPrivateLinkGatewayStatus'
	if jsonDict["kind"] == "AzureEgressPrivateLinkGatewayStatus" {
		// try to unmarshal JSON data into NetworkingV1AzureEgressPrivateLinkGatewayStatus
		err = json.Unmarshal(data, &dst.NetworkingV1AzureEgressPrivateLinkGatewayStatus)
		if err == nil {
			return nil // data stored in dst.NetworkingV1AzureEgressPrivateLinkGatewayStatus, return on the first match
		} else {
			dst.NetworkingV1AzureEgressPrivateLinkGatewayStatus = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1GatewayStatusCloudGatewayOneOf as NetworkingV1AzureEgressPrivateLinkGatewayStatus: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GcpEgressPrivateServiceConnectGatewayStatus'
	if jsonDict["kind"] == "GcpEgressPrivateServiceConnectGatewayStatus" {
		// try to unmarshal JSON data into NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus
		err = json.Unmarshal(data, &dst.NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus)
		if err == nil {
			return nil // data stored in dst.NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus, return on the first match
		} else {
			dst.NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1GatewayStatusCloudGatewayOneOf as NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GcpPeeringGatewayStatus'
	if jsonDict["kind"] == "GcpPeeringGatewayStatus" {
		// try to unmarshal JSON data into NetworkingV1GcpPeeringGatewayStatus
		err = json.Unmarshal(data, &dst.NetworkingV1GcpPeeringGatewayStatus)
		if err == nil {
			return nil // data stored in dst.NetworkingV1GcpPeeringGatewayStatus, return on the first match
		} else {
			dst.NetworkingV1GcpPeeringGatewayStatus = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1GatewayStatusCloudGatewayOneOf as NetworkingV1GcpPeeringGatewayStatus: %s", err.Error())
		}
	}

	// check if the discriminator value is 'networking.v1.AwsEgressPrivateLinkGatewayStatus'
	if jsonDict["kind"] == "networking.v1.AwsEgressPrivateLinkGatewayStatus" {
		// try to unmarshal JSON data into NetworkingV1AwsEgressPrivateLinkGatewayStatus
		err = json.Unmarshal(data, &dst.NetworkingV1AwsEgressPrivateLinkGatewayStatus)
		if err == nil {
			return nil // data stored in dst.NetworkingV1AwsEgressPrivateLinkGatewayStatus, return on the first match
		} else {
			dst.NetworkingV1AwsEgressPrivateLinkGatewayStatus = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1GatewayStatusCloudGatewayOneOf as NetworkingV1AwsEgressPrivateLinkGatewayStatus: %s", err.Error())
		}
	}

	// check if the discriminator value is 'networking.v1.AwsPrivateNetworkInterfaceGatewayStatus'
	if jsonDict["kind"] == "networking.v1.AwsPrivateNetworkInterfaceGatewayStatus" {
		// try to unmarshal JSON data into NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus
		err = json.Unmarshal(data, &dst.NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus)
		if err == nil {
			return nil // data stored in dst.NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus, return on the first match
		} else {
			dst.NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1GatewayStatusCloudGatewayOneOf as NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus: %s", err.Error())
		}
	}

	// check if the discriminator value is 'networking.v1.AzureEgressPrivateLinkGatewayStatus'
	if jsonDict["kind"] == "networking.v1.AzureEgressPrivateLinkGatewayStatus" {
		// try to unmarshal JSON data into NetworkingV1AzureEgressPrivateLinkGatewayStatus
		err = json.Unmarshal(data, &dst.NetworkingV1AzureEgressPrivateLinkGatewayStatus)
		if err == nil {
			return nil // data stored in dst.NetworkingV1AzureEgressPrivateLinkGatewayStatus, return on the first match
		} else {
			dst.NetworkingV1AzureEgressPrivateLinkGatewayStatus = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1GatewayStatusCloudGatewayOneOf as NetworkingV1AzureEgressPrivateLinkGatewayStatus: %s", err.Error())
		}
	}

	// check if the discriminator value is 'networking.v1.GcpEgressPrivateServiceConnectGatewayStatus'
	if jsonDict["kind"] == "networking.v1.GcpEgressPrivateServiceConnectGatewayStatus" {
		// try to unmarshal JSON data into NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus
		err = json.Unmarshal(data, &dst.NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus)
		if err == nil {
			return nil // data stored in dst.NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus, return on the first match
		} else {
			dst.NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1GatewayStatusCloudGatewayOneOf as NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus: %s", err.Error())
		}
	}

	// check if the discriminator value is 'networking.v1.GcpPeeringGatewayStatus'
	if jsonDict["kind"] == "networking.v1.GcpPeeringGatewayStatus" {
		// try to unmarshal JSON data into NetworkingV1GcpPeeringGatewayStatus
		err = json.Unmarshal(data, &dst.NetworkingV1GcpPeeringGatewayStatus)
		if err == nil {
			return nil // data stored in dst.NetworkingV1GcpPeeringGatewayStatus, return on the first match
		} else {
			dst.NetworkingV1GcpPeeringGatewayStatus = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1GatewayStatusCloudGatewayOneOf as NetworkingV1GcpPeeringGatewayStatus: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NetworkingV1GatewayStatusCloudGatewayOneOf) MarshalJSON() ([]byte, error) {
	if src.NetworkingV1AwsEgressPrivateLinkGatewayStatus != nil {
		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(&src.NetworkingV1AwsEgressPrivateLinkGatewayStatus)
		return buffer.Bytes(), err
	}

	if src.NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus != nil {
		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(&src.NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus)
		return buffer.Bytes(), err
	}

	if src.NetworkingV1AzureEgressPrivateLinkGatewayStatus != nil {
		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(&src.NetworkingV1AzureEgressPrivateLinkGatewayStatus)
		return buffer.Bytes(), err
	}

	if src.NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus != nil {
		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(&src.NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus)
		return buffer.Bytes(), err
	}

	if src.NetworkingV1GcpPeeringGatewayStatus != nil {
		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(&src.NetworkingV1GcpPeeringGatewayStatus)
		return buffer.Bytes(), err
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NetworkingV1GatewayStatusCloudGatewayOneOf) GetActualInstance() interface{} {
	if obj.NetworkingV1AwsEgressPrivateLinkGatewayStatus != nil {
		return obj.NetworkingV1AwsEgressPrivateLinkGatewayStatus
	}

	if obj.NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus != nil {
		return obj.NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus
	}

	if obj.NetworkingV1AzureEgressPrivateLinkGatewayStatus != nil {
		return obj.NetworkingV1AzureEgressPrivateLinkGatewayStatus
	}

	if obj.NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus != nil {
		return obj.NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus
	}

	if obj.NetworkingV1GcpPeeringGatewayStatus != nil {
		return obj.NetworkingV1GcpPeeringGatewayStatus
	}

	// all schemas are nil
	return nil
}

type NullableNetworkingV1GatewayStatusCloudGatewayOneOf struct {
	value *NetworkingV1GatewayStatusCloudGatewayOneOf
	isSet bool
}

func (v NullableNetworkingV1GatewayStatusCloudGatewayOneOf) Get() *NetworkingV1GatewayStatusCloudGatewayOneOf {
	return v.value
}

func (v *NullableNetworkingV1GatewayStatusCloudGatewayOneOf) Set(val *NetworkingV1GatewayStatusCloudGatewayOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1GatewayStatusCloudGatewayOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1GatewayStatusCloudGatewayOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1GatewayStatusCloudGatewayOneOf(val *NetworkingV1GatewayStatusCloudGatewayOneOf) *NullableNetworkingV1GatewayStatusCloudGatewayOneOf {
	return &NullableNetworkingV1GatewayStatusCloudGatewayOneOf{value: val, isSet: true}
}

func (v NullableNetworkingV1GatewayStatusCloudGatewayOneOf) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableNetworkingV1GatewayStatusCloudGatewayOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
