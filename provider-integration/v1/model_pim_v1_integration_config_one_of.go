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
Provider Integration Management API

REST API for the Provider Integration

API version: 0.1.0
Contact: connect@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// PimV1IntegrationConfigOneOf - struct for PimV1IntegrationConfigOneOf
type PimV1IntegrationConfigOneOf struct {
	PimV1AwsIntegrationConfig *PimV1AwsIntegrationConfig
}

// PimV1AwsIntegrationConfigAsPimV1IntegrationConfigOneOf is a convenience function that returns PimV1AwsIntegrationConfig wrapped in PimV1IntegrationConfigOneOf
func PimV1AwsIntegrationConfigAsPimV1IntegrationConfigOneOf(v *PimV1AwsIntegrationConfig) PimV1IntegrationConfigOneOf {
	return PimV1IntegrationConfigOneOf{PimV1AwsIntegrationConfig: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PimV1IntegrationConfigOneOf) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'AwsIntegrationConfig'
	if jsonDict["kind"] == "AwsIntegrationConfig" {
		// try to unmarshal JSON data into PimV1AwsIntegrationConfig
		err = json.Unmarshal(data, &dst.PimV1AwsIntegrationConfig)
		if err == nil {
			return nil // data stored in dst.PimV1AwsIntegrationConfig, return on the first match
		} else {
			dst.PimV1AwsIntegrationConfig = nil
			return fmt.Errorf("Failed to unmarshal PimV1IntegrationConfigOneOf as PimV1AwsIntegrationConfig: %s", err.Error())
		}
	}

	// check if the discriminator value is 'pim.v1.AwsIntegrationConfig'
	if jsonDict["kind"] == "pim.v1.AwsIntegrationConfig" {
		// try to unmarshal JSON data into PimV1AwsIntegrationConfig
		err = json.Unmarshal(data, &dst.PimV1AwsIntegrationConfig)
		if err == nil {
			return nil // data stored in dst.PimV1AwsIntegrationConfig, return on the first match
		} else {
			dst.PimV1AwsIntegrationConfig = nil
			return fmt.Errorf("Failed to unmarshal PimV1IntegrationConfigOneOf as PimV1AwsIntegrationConfig: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PimV1IntegrationConfigOneOf) MarshalJSON() ([]byte, error) {
	if src.PimV1AwsIntegrationConfig != nil {
		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(&src.PimV1AwsIntegrationConfig)
		return buffer.Bytes(), err
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PimV1IntegrationConfigOneOf) GetActualInstance() interface{} {
	if obj.PimV1AwsIntegrationConfig != nil {
		return obj.PimV1AwsIntegrationConfig
	}

	// all schemas are nil
	return nil
}

type NullablePimV1IntegrationConfigOneOf struct {
	value *PimV1IntegrationConfigOneOf
	isSet bool
}

func (v NullablePimV1IntegrationConfigOneOf) Get() *PimV1IntegrationConfigOneOf {
	return v.value
}

func (v *NullablePimV1IntegrationConfigOneOf) Set(val *PimV1IntegrationConfigOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePimV1IntegrationConfigOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePimV1IntegrationConfigOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePimV1IntegrationConfigOneOf(val *PimV1IntegrationConfigOneOf) *NullablePimV1IntegrationConfigOneOf {
	return &NullablePimV1IntegrationConfigOneOf{value: val, isSet: true}
}

func (v NullablePimV1IntegrationConfigOneOf) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullablePimV1IntegrationConfigOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
