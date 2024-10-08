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
Cluster Management Access for Schema Registry API

Cluster Management Access for Schema Registry API

API version: 0.0.1-alpha1
Contact: data-governance@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// SrcmV3AccessSpec The desired state of the Access
type SrcmV3AccessSpec struct {
	// Alowed
	Allowed *bool `json:"allowed,omitempty"`
}

// NewSrcmV3AccessSpec instantiates a new SrcmV3AccessSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSrcmV3AccessSpec() *SrcmV3AccessSpec {
	this := SrcmV3AccessSpec{}
	return &this
}

// NewSrcmV3AccessSpecWithDefaults instantiates a new SrcmV3AccessSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSrcmV3AccessSpecWithDefaults() *SrcmV3AccessSpec {
	this := SrcmV3AccessSpec{}
	return &this
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *SrcmV3AccessSpec) GetAllowed() bool {
	if o == nil || o.Allowed == nil {
		var ret bool
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SrcmV3AccessSpec) GetAllowedOk() (*bool, bool) {
	if o == nil || o.Allowed == nil {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *SrcmV3AccessSpec) HasAllowed() bool {
	if o != nil && o.Allowed != nil {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given bool and assigns it to the Allowed field.
func (o *SrcmV3AccessSpec) SetAllowed(v bool) {
	o.Allowed = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *SrcmV3AccessSpec) Redact() {
    o.recurseRedact(o.Allowed)
}

func (o *SrcmV3AccessSpec) recurseRedact(v interface{}) {
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

func (o SrcmV3AccessSpec) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o SrcmV3AccessSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Allowed != nil {
		toSerialize["allowed"] = o.Allowed
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableSrcmV3AccessSpec struct {
	value *SrcmV3AccessSpec
	isSet bool
}

func (v NullableSrcmV3AccessSpec) Get() *SrcmV3AccessSpec {
	return v.value
}

func (v *NullableSrcmV3AccessSpec) Set(val *SrcmV3AccessSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableSrcmV3AccessSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSrcmV3AccessSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSrcmV3AccessSpec(val *SrcmV3AccessSpec) *NullableSrcmV3AccessSpec {
	return &NullableSrcmV3AccessSpec{value: val, isSet: true}
}

func (v NullableSrcmV3AccessSpec) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableSrcmV3AccessSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


