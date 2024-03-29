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
)

import (
	"reflect"
)

// CdxV1ProviderShareStatus The status of the Provider Share
type CdxV1ProviderShareStatus struct {
	// Status of share
	Phase string `json:"phase"`
}

// NewCdxV1ProviderShareStatus instantiates a new CdxV1ProviderShareStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdxV1ProviderShareStatus(phase string) *CdxV1ProviderShareStatus {
	this := CdxV1ProviderShareStatus{}
	this.Phase = phase
	return &this
}

// NewCdxV1ProviderShareStatusWithDefaults instantiates a new CdxV1ProviderShareStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdxV1ProviderShareStatusWithDefaults() *CdxV1ProviderShareStatus {
	this := CdxV1ProviderShareStatus{}
	return &this
}

// GetPhase returns the Phase field value
func (o *CdxV1ProviderShareStatus) GetPhase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value
// and a boolean to check if the value has been set.
func (o *CdxV1ProviderShareStatus) GetPhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phase, true
}

// SetPhase sets field value
func (o *CdxV1ProviderShareStatus) SetPhase(v string) {
	o.Phase = v
}

// Redact resets all sensitive fields to their zero value.
func (o *CdxV1ProviderShareStatus) Redact() {
	o.recurseRedact(&o.Phase)
}

func (o *CdxV1ProviderShareStatus) recurseRedact(v interface{}) {
	type redactor interface {
		Redact()
	}
	if r, ok := v.(redactor); ok {
		r.Redact()
	} else {
		val := reflect.ValueOf(v)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		switch val.Kind() {
		case reflect.Slice, reflect.Array:
			for i := 0; i < val.Len(); i++ {
				// support data types declared without pointers
				o.recurseRedact(val.Index(i).Interface())
				// ... and data types that were declared without but need pointers (for Redact)
				if val.Index(i).CanAddr() {
					o.recurseRedact(val.Index(i).Addr().Interface())
				}
			}
		}
	}
}

func (o CdxV1ProviderShareStatus) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o CdxV1ProviderShareStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["phase"] = o.Phase
	}
	return json.Marshal(toSerialize)
}

type NullableCdxV1ProviderShareStatus struct {
	value *CdxV1ProviderShareStatus
	isSet bool
}

func (v NullableCdxV1ProviderShareStatus) Get() *CdxV1ProviderShareStatus {
	return v.value
}

func (v *NullableCdxV1ProviderShareStatus) Set(val *CdxV1ProviderShareStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCdxV1ProviderShareStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCdxV1ProviderShareStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdxV1ProviderShareStatus(val *CdxV1ProviderShareStatus) *NullableCdxV1ProviderShareStatus {
	return &NullableCdxV1ProviderShareStatus{value: val, isSet: true}
}

func (v NullableCdxV1ProviderShareStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdxV1ProviderShareStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
