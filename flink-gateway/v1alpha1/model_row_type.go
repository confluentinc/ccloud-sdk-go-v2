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
SQL API v1alpha1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
Contact: flink-control-plane@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha1

import (
	"encoding/json"
)

import (
	"reflect"
)

// RowType struct for RowType
type RowType struct {
	// Indicates whether values in this column can be null.
	Nullable bool `json:"nullable"`
	// The data type of the column.
	Type string `json:"type"`
	// The fields of the row. Can be of type [INTEGER, DECIMAL, CHARACTER, ROW, ARRAY, TIMESTAMP, MAP<INT, STRING>]
	Fields []RowFieldType `json:"fields"`
}

// NewRowType instantiates a new RowType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRowType(nullable bool, type_ string, fields []RowFieldType) *RowType {
	this := RowType{}
	this.Nullable = nullable
	this.Type = type_
	this.Fields = fields
	return &this
}

// NewRowTypeWithDefaults instantiates a new RowType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRowTypeWithDefaults() *RowType {
	this := RowType{}
	return &this
}

// GetNullable returns the Nullable field value
func (o *RowType) GetNullable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Nullable
}

// GetNullableOk returns a tuple with the Nullable field value
// and a boolean to check if the value has been set.
func (o *RowType) GetNullableOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Nullable, true
}

// SetNullable sets field value
func (o *RowType) SetNullable(v bool) {
	o.Nullable = v
}

// GetType returns the Type field value
func (o *RowType) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RowType) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RowType) SetType(v string) {
	o.Type = v
}

// GetFields returns the Fields field value
func (o *RowType) GetFields() []RowFieldType {
	if o == nil {
		var ret []RowFieldType
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *RowType) GetFieldsOk() (*[]RowFieldType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *RowType) SetFields(v []RowFieldType) {
	o.Fields = v
}

// Redact resets all sensitive fields to their zero value.
func (o *RowType) Redact() {
    o.recurseRedact(&o.Nullable)
    o.recurseRedact(&o.Type)
    o.recurseRedact(&o.Fields)
}

func (o *RowType) recurseRedact(v interface{}) {
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

func (o RowType) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o RowType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nullable"] = o.Nullable
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["fields"] = o.Fields
	}
	return json.Marshal(toSerialize)
}

type NullableRowType struct {
	value *RowType
	isSet bool
}

func (v NullableRowType) Get() *RowType {
	return v.value
}

func (v *NullableRowType) Set(val *RowType) {
	v.value = val
	v.isSet = true
}

func (v NullableRowType) IsSet() bool {
	return v.isSet
}

func (v *NullableRowType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRowType(val *RowType) *NullableRowType {
	return &NullableRowType{value: val, isSet: true}
}

func (v NullableRowType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRowType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

