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
Flink Resource Pool Management API

This is the Flink Resource Pool management API.

API version: 0.0.1
Contact: ksql-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"fmt"
)

// FrpmV2ResourcePoolSpecUpdateConfigOneOf - struct for FrpmV2ResourcePoolSpecUpdateConfigOneOf
type FrpmV2ResourcePoolSpecUpdateConfigOneOf struct {
	FrpmV2Basic *FrpmV2Basic
	FrpmV2Standard *FrpmV2Standard
}

// FrpmV2BasicAsFrpmV2ResourcePoolSpecUpdateConfigOneOf is a convenience function that returns FrpmV2Basic wrapped in FrpmV2ResourcePoolSpecUpdateConfigOneOf
func FrpmV2BasicAsFrpmV2ResourcePoolSpecUpdateConfigOneOf(v *FrpmV2Basic) FrpmV2ResourcePoolSpecUpdateConfigOneOf {
	return FrpmV2ResourcePoolSpecUpdateConfigOneOf{ FrpmV2Basic: v}
}

// FrpmV2StandardAsFrpmV2ResourcePoolSpecUpdateConfigOneOf is a convenience function that returns FrpmV2Standard wrapped in FrpmV2ResourcePoolSpecUpdateConfigOneOf
func FrpmV2StandardAsFrpmV2ResourcePoolSpecUpdateConfigOneOf(v *FrpmV2Standard) FrpmV2ResourcePoolSpecUpdateConfigOneOf {
	return FrpmV2ResourcePoolSpecUpdateConfigOneOf{ FrpmV2Standard: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *FrpmV2ResourcePoolSpecUpdateConfigOneOf) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'Basic'
	if jsonDict["kind"] == "Basic" {
		// try to unmarshal JSON data into FrpmV2Basic
		err = json.Unmarshal(data, &dst.FrpmV2Basic)
		if err == nil {
			return nil // data stored in dst.FrpmV2Basic, return on the first match
		} else {
			dst.FrpmV2Basic = nil
			return fmt.Errorf("Failed to unmarshal FrpmV2ResourcePoolSpecUpdateConfigOneOf as FrpmV2Basic: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Standard'
	if jsonDict["kind"] == "Standard" {
		// try to unmarshal JSON data into FrpmV2Standard
		err = json.Unmarshal(data, &dst.FrpmV2Standard)
		if err == nil {
			return nil // data stored in dst.FrpmV2Standard, return on the first match
		} else {
			dst.FrpmV2Standard = nil
			return fmt.Errorf("Failed to unmarshal FrpmV2ResourcePoolSpecUpdateConfigOneOf as FrpmV2Standard: %s", err.Error())
		}
	}

	// check if the discriminator value is 'frpm.v2.Basic'
	if jsonDict["kind"] == "frpm.v2.Basic" {
		// try to unmarshal JSON data into FrpmV2Basic
		err = json.Unmarshal(data, &dst.FrpmV2Basic)
		if err == nil {
			return nil // data stored in dst.FrpmV2Basic, return on the first match
		} else {
			dst.FrpmV2Basic = nil
			return fmt.Errorf("Failed to unmarshal FrpmV2ResourcePoolSpecUpdateConfigOneOf as FrpmV2Basic: %s", err.Error())
		}
	}

	// check if the discriminator value is 'frpm.v2.Standard'
	if jsonDict["kind"] == "frpm.v2.Standard" {
		// try to unmarshal JSON data into FrpmV2Standard
		err = json.Unmarshal(data, &dst.FrpmV2Standard)
		if err == nil {
			return nil // data stored in dst.FrpmV2Standard, return on the first match
		} else {
			dst.FrpmV2Standard = nil
			return fmt.Errorf("Failed to unmarshal FrpmV2ResourcePoolSpecUpdateConfigOneOf as FrpmV2Standard: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FrpmV2ResourcePoolSpecUpdateConfigOneOf) MarshalJSON() ([]byte, error) {
	if src.FrpmV2Basic != nil {
		return json.Marshal(&src.FrpmV2Basic)
	}

	if src.FrpmV2Standard != nil {
		return json.Marshal(&src.FrpmV2Standard)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *FrpmV2ResourcePoolSpecUpdateConfigOneOf) GetActualInstance() (interface{}) {
	if obj.FrpmV2Basic != nil {
		return obj.FrpmV2Basic
	}

	if obj.FrpmV2Standard != nil {
		return obj.FrpmV2Standard
	}

	// all schemas are nil
	return nil
}

type NullableFrpmV2ResourcePoolSpecUpdateConfigOneOf struct {
	value *FrpmV2ResourcePoolSpecUpdateConfigOneOf
	isSet bool
}

func (v NullableFrpmV2ResourcePoolSpecUpdateConfigOneOf) Get() *FrpmV2ResourcePoolSpecUpdateConfigOneOf {
	return v.value
}

func (v *NullableFrpmV2ResourcePoolSpecUpdateConfigOneOf) Set(val *FrpmV2ResourcePoolSpecUpdateConfigOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFrpmV2ResourcePoolSpecUpdateConfigOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFrpmV2ResourcePoolSpecUpdateConfigOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrpmV2ResourcePoolSpecUpdateConfigOneOf(val *FrpmV2ResourcePoolSpecUpdateConfigOneOf) *NullableFrpmV2ResourcePoolSpecUpdateConfigOneOf {
	return &NullableFrpmV2ResourcePoolSpecUpdateConfigOneOf{value: val, isSet: true}
}

func (v NullableFrpmV2ResourcePoolSpecUpdateConfigOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrpmV2ResourcePoolSpecUpdateConfigOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

