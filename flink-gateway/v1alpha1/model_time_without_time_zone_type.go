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

// TimeWithoutTimeZoneType struct for TimeWithoutTimeZoneType
type TimeWithoutTimeZoneType struct {
	// Indicates whether values in this column can be null.
	Nullable bool `json:"nullable"`
	// The data type of the column.
	Type string `json:"type"`
	// The scale of the time type.
	Scale *int32 `json:"scale,omitempty"`
	// The precision of the time type.
	Precision *int32 `json:"precision,omitempty"`
}

// NewTimeWithoutTimeZoneType instantiates a new TimeWithoutTimeZoneType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeWithoutTimeZoneType(nullable bool, type_ string) *TimeWithoutTimeZoneType {
	this := TimeWithoutTimeZoneType{}
	this.Nullable = nullable
	this.Type = type_
	return &this
}

// NewTimeWithoutTimeZoneTypeWithDefaults instantiates a new TimeWithoutTimeZoneType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeWithoutTimeZoneTypeWithDefaults() *TimeWithoutTimeZoneType {
	this := TimeWithoutTimeZoneType{}
	return &this
}

// GetNullable returns the Nullable field value
func (o *TimeWithoutTimeZoneType) GetNullable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Nullable
}

// GetNullableOk returns a tuple with the Nullable field value
// and a boolean to check if the value has been set.
func (o *TimeWithoutTimeZoneType) GetNullableOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Nullable, true
}

// SetNullable sets field value
func (o *TimeWithoutTimeZoneType) SetNullable(v bool) {
	o.Nullable = v
}

// GetType returns the Type field value
func (o *TimeWithoutTimeZoneType) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TimeWithoutTimeZoneType) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TimeWithoutTimeZoneType) SetType(v string) {
	o.Type = v
}

// GetScale returns the Scale field value if set, zero value otherwise.
func (o *TimeWithoutTimeZoneType) GetScale() int32 {
	if o == nil || o.Scale == nil {
		var ret int32
		return ret
	}
	return *o.Scale
}

// GetScaleOk returns a tuple with the Scale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeWithoutTimeZoneType) GetScaleOk() (*int32, bool) {
	if o == nil || o.Scale == nil {
		return nil, false
	}
	return o.Scale, true
}

// HasScale returns a boolean if a field has been set.
func (o *TimeWithoutTimeZoneType) HasScale() bool {
	if o != nil && o.Scale != nil {
		return true
	}

	return false
}

// SetScale gets a reference to the given int32 and assigns it to the Scale field.
func (o *TimeWithoutTimeZoneType) SetScale(v int32) {
	o.Scale = &v
}

// GetPrecision returns the Precision field value if set, zero value otherwise.
func (o *TimeWithoutTimeZoneType) GetPrecision() int32 {
	if o == nil || o.Precision == nil {
		var ret int32
		return ret
	}
	return *o.Precision
}

// GetPrecisionOk returns a tuple with the Precision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeWithoutTimeZoneType) GetPrecisionOk() (*int32, bool) {
	if o == nil || o.Precision == nil {
		return nil, false
	}
	return o.Precision, true
}

// HasPrecision returns a boolean if a field has been set.
func (o *TimeWithoutTimeZoneType) HasPrecision() bool {
	if o != nil && o.Precision != nil {
		return true
	}

	return false
}

// SetPrecision gets a reference to the given int32 and assigns it to the Precision field.
func (o *TimeWithoutTimeZoneType) SetPrecision(v int32) {
	o.Precision = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *TimeWithoutTimeZoneType) Redact() {
    o.recurseRedact(&o.Nullable)
    o.recurseRedact(&o.Type)
    o.recurseRedact(o.Scale)
    o.recurseRedact(o.Precision)
}

func (o *TimeWithoutTimeZoneType) recurseRedact(v interface{}) {
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

func (o TimeWithoutTimeZoneType) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o TimeWithoutTimeZoneType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nullable"] = o.Nullable
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Scale != nil {
		toSerialize["scale"] = o.Scale
	}
	if o.Precision != nil {
		toSerialize["precision"] = o.Precision
	}
	return json.Marshal(toSerialize)
}

type NullableTimeWithoutTimeZoneType struct {
	value *TimeWithoutTimeZoneType
	isSet bool
}

func (v NullableTimeWithoutTimeZoneType) Get() *TimeWithoutTimeZoneType {
	return v.value
}

func (v *NullableTimeWithoutTimeZoneType) Set(val *TimeWithoutTimeZoneType) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeWithoutTimeZoneType) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeWithoutTimeZoneType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeWithoutTimeZoneType(val *TimeWithoutTimeZoneType) *NullableTimeWithoutTimeZoneType {
	return &NullableTimeWithoutTimeZoneType{value: val, isSet: true}
}

func (v NullableTimeWithoutTimeZoneType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeWithoutTimeZoneType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

