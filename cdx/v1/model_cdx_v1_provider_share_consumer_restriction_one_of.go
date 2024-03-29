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
Stream Share APIs

# Introduction

API version: 0.1.0-alpha0
Contact: cdx@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"fmt"
)

// CdxV1ProviderShareConsumerRestrictionOneOf - struct for CdxV1ProviderShareConsumerRestrictionOneOf
type CdxV1ProviderShareConsumerRestrictionOneOf struct {
	CdxV1EmailConsumerRestriction *CdxV1EmailConsumerRestriction
}

// CdxV1EmailConsumerRestrictionAsCdxV1ProviderShareConsumerRestrictionOneOf is a convenience function that returns CdxV1EmailConsumerRestriction wrapped in CdxV1ProviderShareConsumerRestrictionOneOf
func CdxV1EmailConsumerRestrictionAsCdxV1ProviderShareConsumerRestrictionOneOf(v *CdxV1EmailConsumerRestriction) CdxV1ProviderShareConsumerRestrictionOneOf {
	return CdxV1ProviderShareConsumerRestrictionOneOf{CdxV1EmailConsumerRestriction: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CdxV1ProviderShareConsumerRestrictionOneOf) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'Email'
	if jsonDict["kind"] == "Email" {
		// try to unmarshal JSON data into CdxV1EmailConsumerRestriction
		err = json.Unmarshal(data, &dst.CdxV1EmailConsumerRestriction)
		if err == nil {
			return nil // data stored in dst.CdxV1EmailConsumerRestriction, return on the first match
		} else {
			dst.CdxV1EmailConsumerRestriction = nil
			return fmt.Errorf("Failed to unmarshal CdxV1ProviderShareConsumerRestrictionOneOf as CdxV1EmailConsumerRestriction: %s", err.Error())
		}
	}

	// check if the discriminator value is 'cdx.v1.EmailConsumerRestriction'
	if jsonDict["kind"] == "cdx.v1.EmailConsumerRestriction" {
		// try to unmarshal JSON data into CdxV1EmailConsumerRestriction
		err = json.Unmarshal(data, &dst.CdxV1EmailConsumerRestriction)
		if err == nil {
			return nil // data stored in dst.CdxV1EmailConsumerRestriction, return on the first match
		} else {
			dst.CdxV1EmailConsumerRestriction = nil
			return fmt.Errorf("Failed to unmarshal CdxV1ProviderShareConsumerRestrictionOneOf as CdxV1EmailConsumerRestriction: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CdxV1ProviderShareConsumerRestrictionOneOf) MarshalJSON() ([]byte, error) {
	if src.CdxV1EmailConsumerRestriction != nil {
		return json.Marshal(&src.CdxV1EmailConsumerRestriction)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CdxV1ProviderShareConsumerRestrictionOneOf) GetActualInstance() interface{} {
	if obj.CdxV1EmailConsumerRestriction != nil {
		return obj.CdxV1EmailConsumerRestriction
	}

	// all schemas are nil
	return nil
}

type NullableCdxV1ProviderShareConsumerRestrictionOneOf struct {
	value *CdxV1ProviderShareConsumerRestrictionOneOf
	isSet bool
}

func (v NullableCdxV1ProviderShareConsumerRestrictionOneOf) Get() *CdxV1ProviderShareConsumerRestrictionOneOf {
	return v.value
}

func (v *NullableCdxV1ProviderShareConsumerRestrictionOneOf) Set(val *CdxV1ProviderShareConsumerRestrictionOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCdxV1ProviderShareConsumerRestrictionOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCdxV1ProviderShareConsumerRestrictionOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdxV1ProviderShareConsumerRestrictionOneOf(val *CdxV1ProviderShareConsumerRestrictionOneOf) *NullableCdxV1ProviderShareConsumerRestrictionOneOf {
	return &NullableCdxV1ProviderShareConsumerRestrictionOneOf{value: val, isSet: true}
}

func (v NullableCdxV1ProviderShareConsumerRestrictionOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdxV1ProviderShareConsumerRestrictionOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
