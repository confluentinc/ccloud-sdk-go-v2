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

Network API

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

// NetworkingV1AccessPointSpecUpdateConfigOneOf - struct for NetworkingV1AccessPointSpecUpdateConfigOneOf
type NetworkingV1AccessPointSpecUpdateConfigOneOf struct {
	NetworkingV1AwsEgressPrivateLinkEndpoint           *NetworkingV1AwsEgressPrivateLinkEndpoint
	NetworkingV1AwsPrivateNetworkInterface             *NetworkingV1AwsPrivateNetworkInterface
	NetworkingV1AzureEgressPrivateLinkEndpoint         *NetworkingV1AzureEgressPrivateLinkEndpoint
	NetworkingV1GcpEgressPrivateServiceConnectEndpoint *NetworkingV1GcpEgressPrivateServiceConnectEndpoint
}

// NetworkingV1AwsEgressPrivateLinkEndpointAsNetworkingV1AccessPointSpecUpdateConfigOneOf is a convenience function that returns NetworkingV1AwsEgressPrivateLinkEndpoint wrapped in NetworkingV1AccessPointSpecUpdateConfigOneOf
func NetworkingV1AwsEgressPrivateLinkEndpointAsNetworkingV1AccessPointSpecUpdateConfigOneOf(v *NetworkingV1AwsEgressPrivateLinkEndpoint) NetworkingV1AccessPointSpecUpdateConfigOneOf {
	return NetworkingV1AccessPointSpecUpdateConfigOneOf{NetworkingV1AwsEgressPrivateLinkEndpoint: v}
}

// NetworkingV1AwsPrivateNetworkInterfaceAsNetworkingV1AccessPointSpecUpdateConfigOneOf is a convenience function that returns NetworkingV1AwsPrivateNetworkInterface wrapped in NetworkingV1AccessPointSpecUpdateConfigOneOf
func NetworkingV1AwsPrivateNetworkInterfaceAsNetworkingV1AccessPointSpecUpdateConfigOneOf(v *NetworkingV1AwsPrivateNetworkInterface) NetworkingV1AccessPointSpecUpdateConfigOneOf {
	return NetworkingV1AccessPointSpecUpdateConfigOneOf{NetworkingV1AwsPrivateNetworkInterface: v}
}

// NetworkingV1AzureEgressPrivateLinkEndpointAsNetworkingV1AccessPointSpecUpdateConfigOneOf is a convenience function that returns NetworkingV1AzureEgressPrivateLinkEndpoint wrapped in NetworkingV1AccessPointSpecUpdateConfigOneOf
func NetworkingV1AzureEgressPrivateLinkEndpointAsNetworkingV1AccessPointSpecUpdateConfigOneOf(v *NetworkingV1AzureEgressPrivateLinkEndpoint) NetworkingV1AccessPointSpecUpdateConfigOneOf {
	return NetworkingV1AccessPointSpecUpdateConfigOneOf{NetworkingV1AzureEgressPrivateLinkEndpoint: v}
}

// NetworkingV1GcpEgressPrivateServiceConnectEndpointAsNetworkingV1AccessPointSpecUpdateConfigOneOf is a convenience function that returns NetworkingV1GcpEgressPrivateServiceConnectEndpoint wrapped in NetworkingV1AccessPointSpecUpdateConfigOneOf
func NetworkingV1GcpEgressPrivateServiceConnectEndpointAsNetworkingV1AccessPointSpecUpdateConfigOneOf(v *NetworkingV1GcpEgressPrivateServiceConnectEndpoint) NetworkingV1AccessPointSpecUpdateConfigOneOf {
	return NetworkingV1AccessPointSpecUpdateConfigOneOf{NetworkingV1GcpEgressPrivateServiceConnectEndpoint: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NetworkingV1AccessPointSpecUpdateConfigOneOf) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'AwsEgressPrivateLinkEndpoint'
	if jsonDict["kind"] == "AwsEgressPrivateLinkEndpoint" {
		// try to unmarshal JSON data into NetworkingV1AwsEgressPrivateLinkEndpoint
		err = json.Unmarshal(data, &dst.NetworkingV1AwsEgressPrivateLinkEndpoint)
		if err == nil {
			return nil // data stored in dst.NetworkingV1AwsEgressPrivateLinkEndpoint, return on the first match
		} else {
			dst.NetworkingV1AwsEgressPrivateLinkEndpoint = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1AccessPointSpecUpdateConfigOneOf as NetworkingV1AwsEgressPrivateLinkEndpoint: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AwsPrivateNetworkInterface'
	if jsonDict["kind"] == "AwsPrivateNetworkInterface" {
		// try to unmarshal JSON data into NetworkingV1AwsPrivateNetworkInterface
		err = json.Unmarshal(data, &dst.NetworkingV1AwsPrivateNetworkInterface)
		if err == nil {
			return nil // data stored in dst.NetworkingV1AwsPrivateNetworkInterface, return on the first match
		} else {
			dst.NetworkingV1AwsPrivateNetworkInterface = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1AccessPointSpecUpdateConfigOneOf as NetworkingV1AwsPrivateNetworkInterface: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AzureEgressPrivateLinkEndpoint'
	if jsonDict["kind"] == "AzureEgressPrivateLinkEndpoint" {
		// try to unmarshal JSON data into NetworkingV1AzureEgressPrivateLinkEndpoint
		err = json.Unmarshal(data, &dst.NetworkingV1AzureEgressPrivateLinkEndpoint)
		if err == nil {
			return nil // data stored in dst.NetworkingV1AzureEgressPrivateLinkEndpoint, return on the first match
		} else {
			dst.NetworkingV1AzureEgressPrivateLinkEndpoint = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1AccessPointSpecUpdateConfigOneOf as NetworkingV1AzureEgressPrivateLinkEndpoint: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GcpEgressPrivateServiceConnectEndpoint'
	if jsonDict["kind"] == "GcpEgressPrivateServiceConnectEndpoint" {
		// try to unmarshal JSON data into NetworkingV1GcpEgressPrivateServiceConnectEndpoint
		err = json.Unmarshal(data, &dst.NetworkingV1GcpEgressPrivateServiceConnectEndpoint)
		if err == nil {
			return nil // data stored in dst.NetworkingV1GcpEgressPrivateServiceConnectEndpoint, return on the first match
		} else {
			dst.NetworkingV1GcpEgressPrivateServiceConnectEndpoint = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1AccessPointSpecUpdateConfigOneOf as NetworkingV1GcpEgressPrivateServiceConnectEndpoint: %s", err.Error())
		}
	}

	// check if the discriminator value is 'networking.v1.AwsEgressPrivateLinkEndpoint'
	if jsonDict["kind"] == "networking.v1.AwsEgressPrivateLinkEndpoint" {
		// try to unmarshal JSON data into NetworkingV1AwsEgressPrivateLinkEndpoint
		err = json.Unmarshal(data, &dst.NetworkingV1AwsEgressPrivateLinkEndpoint)
		if err == nil {
			return nil // data stored in dst.NetworkingV1AwsEgressPrivateLinkEndpoint, return on the first match
		} else {
			dst.NetworkingV1AwsEgressPrivateLinkEndpoint = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1AccessPointSpecUpdateConfigOneOf as NetworkingV1AwsEgressPrivateLinkEndpoint: %s", err.Error())
		}
	}

	// check if the discriminator value is 'networking.v1.AwsPrivateNetworkInterface'
	if jsonDict["kind"] == "networking.v1.AwsPrivateNetworkInterface" {
		// try to unmarshal JSON data into NetworkingV1AwsPrivateNetworkInterface
		err = json.Unmarshal(data, &dst.NetworkingV1AwsPrivateNetworkInterface)
		if err == nil {
			return nil // data stored in dst.NetworkingV1AwsPrivateNetworkInterface, return on the first match
		} else {
			dst.NetworkingV1AwsPrivateNetworkInterface = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1AccessPointSpecUpdateConfigOneOf as NetworkingV1AwsPrivateNetworkInterface: %s", err.Error())
		}
	}

	// check if the discriminator value is 'networking.v1.AzureEgressPrivateLinkEndpoint'
	if jsonDict["kind"] == "networking.v1.AzureEgressPrivateLinkEndpoint" {
		// try to unmarshal JSON data into NetworkingV1AzureEgressPrivateLinkEndpoint
		err = json.Unmarshal(data, &dst.NetworkingV1AzureEgressPrivateLinkEndpoint)
		if err == nil {
			return nil // data stored in dst.NetworkingV1AzureEgressPrivateLinkEndpoint, return on the first match
		} else {
			dst.NetworkingV1AzureEgressPrivateLinkEndpoint = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1AccessPointSpecUpdateConfigOneOf as NetworkingV1AzureEgressPrivateLinkEndpoint: %s", err.Error())
		}
	}

	// check if the discriminator value is 'networking.v1.GcpEgressPrivateServiceConnectEndpoint'
	if jsonDict["kind"] == "networking.v1.GcpEgressPrivateServiceConnectEndpoint" {
		// try to unmarshal JSON data into NetworkingV1GcpEgressPrivateServiceConnectEndpoint
		err = json.Unmarshal(data, &dst.NetworkingV1GcpEgressPrivateServiceConnectEndpoint)
		if err == nil {
			return nil // data stored in dst.NetworkingV1GcpEgressPrivateServiceConnectEndpoint, return on the first match
		} else {
			dst.NetworkingV1GcpEgressPrivateServiceConnectEndpoint = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1AccessPointSpecUpdateConfigOneOf as NetworkingV1GcpEgressPrivateServiceConnectEndpoint: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NetworkingV1AccessPointSpecUpdateConfigOneOf) MarshalJSON() ([]byte, error) {
	if src.NetworkingV1AwsEgressPrivateLinkEndpoint != nil {
		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(&src.NetworkingV1AwsEgressPrivateLinkEndpoint)
		return buffer.Bytes(), err
	}

	if src.NetworkingV1AwsPrivateNetworkInterface != nil {
		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(&src.NetworkingV1AwsPrivateNetworkInterface)
		return buffer.Bytes(), err
	}

	if src.NetworkingV1AzureEgressPrivateLinkEndpoint != nil {
		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(&src.NetworkingV1AzureEgressPrivateLinkEndpoint)
		return buffer.Bytes(), err
	}

	if src.NetworkingV1GcpEgressPrivateServiceConnectEndpoint != nil {
		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(&src.NetworkingV1GcpEgressPrivateServiceConnectEndpoint)
		return buffer.Bytes(), err
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NetworkingV1AccessPointSpecUpdateConfigOneOf) GetActualInstance() interface{} {
	if obj.NetworkingV1AwsEgressPrivateLinkEndpoint != nil {
		return obj.NetworkingV1AwsEgressPrivateLinkEndpoint
	}

	if obj.NetworkingV1AwsPrivateNetworkInterface != nil {
		return obj.NetworkingV1AwsPrivateNetworkInterface
	}

	if obj.NetworkingV1AzureEgressPrivateLinkEndpoint != nil {
		return obj.NetworkingV1AzureEgressPrivateLinkEndpoint
	}

	if obj.NetworkingV1GcpEgressPrivateServiceConnectEndpoint != nil {
		return obj.NetworkingV1GcpEgressPrivateServiceConnectEndpoint
	}

	// all schemas are nil
	return nil
}

type NullableNetworkingV1AccessPointSpecUpdateConfigOneOf struct {
	value *NetworkingV1AccessPointSpecUpdateConfigOneOf
	isSet bool
}

func (v NullableNetworkingV1AccessPointSpecUpdateConfigOneOf) Get() *NetworkingV1AccessPointSpecUpdateConfigOneOf {
	return v.value
}

func (v *NullableNetworkingV1AccessPointSpecUpdateConfigOneOf) Set(val *NetworkingV1AccessPointSpecUpdateConfigOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1AccessPointSpecUpdateConfigOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1AccessPointSpecUpdateConfigOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1AccessPointSpecUpdateConfigOneOf(val *NetworkingV1AccessPointSpecUpdateConfigOneOf) *NullableNetworkingV1AccessPointSpecUpdateConfigOneOf {
	return &NullableNetworkingV1AccessPointSpecUpdateConfigOneOf{value: val, isSet: true}
}

func (v NullableNetworkingV1AccessPointSpecUpdateConfigOneOf) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableNetworkingV1AccessPointSpecUpdateConfigOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}