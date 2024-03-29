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
SQL API v1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
Contact: flink-control-plane@confluent.io
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

// RowFieldType struct for RowFieldType
type RowFieldType struct {
	// The name of the field.
	Name string `json:"name,omitempty"`
	// The data type of the field.
	FieldType DataType `json:"field_type,omitempty"`
	// The description of the field.
	Description *string `json:"description,omitempty"`
}

// NewRowFieldType instantiates a new RowFieldType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRowFieldType(name string, fieldType DataType) *RowFieldType {
	this := RowFieldType{}
	this.Name = name
	this.FieldType = fieldType
	return &this
}

// NewRowFieldTypeWithDefaults instantiates a new RowFieldType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRowFieldTypeWithDefaults() *RowFieldType {
	this := RowFieldType{}
	return &this
}

// GetName returns the Name field value
func (o *RowFieldType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RowFieldType) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RowFieldType) SetName(v string) {
	o.Name = v
}

// GetFieldType returns the FieldType field value
func (o *RowFieldType) GetFieldType() DataType {
	if o == nil {
		var ret DataType
		return ret
	}

	return o.FieldType
}

// GetFieldTypeOk returns a tuple with the FieldType field value
// and a boolean to check if the value has been set.
func (o *RowFieldType) GetFieldTypeOk() (*DataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldType, true
}

// SetFieldType sets field value
func (o *RowFieldType) SetFieldType(v DataType) {
	o.FieldType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RowFieldType) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RowFieldType) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RowFieldType) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RowFieldType) SetDescription(v string) {
	o.Description = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *RowFieldType) Redact() {
	o.recurseRedact(&o.Name)
	o.recurseRedact(&o.FieldType)
	o.recurseRedact(o.Description)
}

func (o *RowFieldType) recurseRedact(v interface{}) {
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

func (o RowFieldType) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o RowFieldType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["field_type"] = o.FieldType
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableRowFieldType struct {
	value *RowFieldType
	isSet bool
}

func (v NullableRowFieldType) Get() *RowFieldType {
	return v.value
}

func (v *NullableRowFieldType) Set(val *RowFieldType) {
	v.value = val
	v.isSet = true
}

func (v NullableRowFieldType) IsSet() bool {
	return v.isSet
}

func (v *NullableRowFieldType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRowFieldType(val *RowFieldType) *NullableRowFieldType {
	return &NullableRowFieldType{value: val, isSet: true}
}

func (v NullableRowFieldType) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableRowFieldType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
