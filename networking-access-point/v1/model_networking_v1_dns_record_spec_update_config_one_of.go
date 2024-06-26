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

// NetworkingV1DnsRecordSpecUpdateConfigOneOf - struct for NetworkingV1DnsRecordSpecUpdateConfigOneOf
type NetworkingV1DnsRecordSpecUpdateConfigOneOf struct {
	NetworkingV1PrivateLinkAccessPoint *NetworkingV1PrivateLinkAccessPoint
}

// NetworkingV1PrivateLinkAccessPointAsNetworkingV1DnsRecordSpecUpdateConfigOneOf is a convenience function that returns NetworkingV1PrivateLinkAccessPoint wrapped in NetworkingV1DnsRecordSpecUpdateConfigOneOf
func NetworkingV1PrivateLinkAccessPointAsNetworkingV1DnsRecordSpecUpdateConfigOneOf(v *NetworkingV1PrivateLinkAccessPoint) NetworkingV1DnsRecordSpecUpdateConfigOneOf {
	return NetworkingV1DnsRecordSpecUpdateConfigOneOf{NetworkingV1PrivateLinkAccessPoint: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NetworkingV1DnsRecordSpecUpdateConfigOneOf) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'PrivateLinkAccessPoint'
	if jsonDict["kind"] == "PrivateLinkAccessPoint" {
		// try to unmarshal JSON data into NetworkingV1PrivateLinkAccessPoint
		err = json.Unmarshal(data, &dst.NetworkingV1PrivateLinkAccessPoint)
		if err == nil {
			return nil // data stored in dst.NetworkingV1PrivateLinkAccessPoint, return on the first match
		} else {
			dst.NetworkingV1PrivateLinkAccessPoint = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1DnsRecordSpecUpdateConfigOneOf as NetworkingV1PrivateLinkAccessPoint: %s", err.Error())
		}
	}

	// check if the discriminator value is 'networking.v1.PrivateLinkAccessPoint'
	if jsonDict["kind"] == "networking.v1.PrivateLinkAccessPoint" {
		// try to unmarshal JSON data into NetworkingV1PrivateLinkAccessPoint
		err = json.Unmarshal(data, &dst.NetworkingV1PrivateLinkAccessPoint)
		if err == nil {
			return nil // data stored in dst.NetworkingV1PrivateLinkAccessPoint, return on the first match
		} else {
			dst.NetworkingV1PrivateLinkAccessPoint = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1DnsRecordSpecUpdateConfigOneOf as NetworkingV1PrivateLinkAccessPoint: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NetworkingV1DnsRecordSpecUpdateConfigOneOf) MarshalJSON() ([]byte, error) {
	if src.NetworkingV1PrivateLinkAccessPoint != nil {
		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(&src.NetworkingV1PrivateLinkAccessPoint)
		return buffer.Bytes(), err
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NetworkingV1DnsRecordSpecUpdateConfigOneOf) GetActualInstance() interface{} {
	if obj.NetworkingV1PrivateLinkAccessPoint != nil {
		return obj.NetworkingV1PrivateLinkAccessPoint
	}

	// all schemas are nil
	return nil
}

type NullableNetworkingV1DnsRecordSpecUpdateConfigOneOf struct {
	value *NetworkingV1DnsRecordSpecUpdateConfigOneOf
	isSet bool
}

func (v NullableNetworkingV1DnsRecordSpecUpdateConfigOneOf) Get() *NetworkingV1DnsRecordSpecUpdateConfigOneOf {
	return v.value
}

func (v *NullableNetworkingV1DnsRecordSpecUpdateConfigOneOf) Set(val *NetworkingV1DnsRecordSpecUpdateConfigOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1DnsRecordSpecUpdateConfigOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1DnsRecordSpecUpdateConfigOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1DnsRecordSpecUpdateConfigOneOf(val *NetworkingV1DnsRecordSpecUpdateConfigOneOf) *NullableNetworkingV1DnsRecordSpecUpdateConfigOneOf {
	return &NullableNetworkingV1DnsRecordSpecUpdateConfigOneOf{value: val, isSet: true}
}

func (v NullableNetworkingV1DnsRecordSpecUpdateConfigOneOf) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableNetworkingV1DnsRecordSpecUpdateConfigOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
