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
Confluent Data Catalog

REST API for the Data Catalog

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

// ConstraintDef struct for ConstraintDef
type ConstraintDef struct {
	// The type
	Type *string `json:"type,omitempty"`
	// The params
	Params *map[string]map[string]interface{} `json:"params,omitempty"`
}

// NewConstraintDef instantiates a new ConstraintDef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstraintDef() *ConstraintDef {
	this := ConstraintDef{}
	return &this
}

// NewConstraintDefWithDefaults instantiates a new ConstraintDef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstraintDefWithDefaults() *ConstraintDef {
	this := ConstraintDef{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConstraintDef) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstraintDef) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConstraintDef) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ConstraintDef) SetType(v string) {
	o.Type = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *ConstraintDef) GetParams() map[string]map[string]interface{} {
	if o == nil || o.Params == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstraintDef) GetParamsOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *ConstraintDef) HasParams() bool {
	if o != nil && o.Params != nil {
		return true
	}

	return false
}

// SetParams gets a reference to the given map[string]map[string]interface{} and assigns it to the Params field.
func (o *ConstraintDef) SetParams(v map[string]map[string]interface{}) {
	o.Params = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *ConstraintDef) Redact() {
	o.recurseRedact(o.Type)
	o.recurseRedact(o.Params)
}

func (o *ConstraintDef) recurseRedact(v interface{}) {
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

func (o ConstraintDef) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ConstraintDef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableConstraintDef struct {
	value *ConstraintDef
	isSet bool
}

func (v NullableConstraintDef) Get() *ConstraintDef {
	return v.value
}

func (v *NullableConstraintDef) Set(val *ConstraintDef) {
	v.value = val
	v.isSet = true
}

func (v NullableConstraintDef) IsSet() bool {
	return v.isSet
}

func (v *NullableConstraintDef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstraintDef(val *ConstraintDef) *NullableConstraintDef {
	return &NullableConstraintDef{value: val, isSet: true}
}

func (v NullableConstraintDef) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableConstraintDef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
