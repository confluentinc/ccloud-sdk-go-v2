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
Confluent Schema Registry APIs

REST API for the Schema Registry

API version: 1.0.0
Contact: data-governance@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// ConfigOverrideMetadata Override value for the metadata to be used during schema registration.
type ConfigOverrideMetadata struct {
	// The metadata properties and their new values
	Properties *map[string]interface{} `json:"properties,omitempty"`
}

// NewConfigOverrideMetadata instantiates a new ConfigOverrideMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigOverrideMetadata() *ConfigOverrideMetadata {
	this := ConfigOverrideMetadata{}
	return &this
}

// NewConfigOverrideMetadataWithDefaults instantiates a new ConfigOverrideMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigOverrideMetadataWithDefaults() *ConfigOverrideMetadata {
	this := ConfigOverrideMetadata{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ConfigOverrideMetadata) GetProperties() map[string]interface{} {
	if o == nil || o.Properties == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigOverrideMetadata) GetPropertiesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ConfigOverrideMetadata) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *ConfigOverrideMetadata) SetProperties(v map[string]interface{}) {
	o.Properties = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *ConfigOverrideMetadata) Redact() {
	o.recurseRedact(o.Properties)
}

func (o *ConfigOverrideMetadata) recurseRedact(v interface{}) {
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

func (o ConfigOverrideMetadata) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ConfigOverrideMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableConfigOverrideMetadata struct {
	value *ConfigOverrideMetadata
	isSet bool
}

func (v NullableConfigOverrideMetadata) Get() *ConfigOverrideMetadata {
	return v.value
}

func (v *NullableConfigOverrideMetadata) Set(val *ConfigOverrideMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigOverrideMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigOverrideMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigOverrideMetadata(val *ConfigOverrideMetadata) *NullableConfigOverrideMetadata {
	return &NullableConfigOverrideMetadata{value: val, isSet: true}
}

func (v NullableConfigOverrideMetadata) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableConfigOverrideMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
