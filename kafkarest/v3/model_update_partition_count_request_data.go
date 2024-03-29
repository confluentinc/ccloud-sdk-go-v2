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
REST Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0.0
Contact: kafka-clients-proxy-team@confluent.io
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

// UpdatePartitionCountRequestData struct for UpdatePartitionCountRequestData
type UpdatePartitionCountRequestData struct {
	PartitionsCount int32 `json:"partitions_count,omitempty"`
}

// NewUpdatePartitionCountRequestData instantiates a new UpdatePartitionCountRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePartitionCountRequestData(partitionsCount int32) *UpdatePartitionCountRequestData {
	this := UpdatePartitionCountRequestData{}
	this.PartitionsCount = partitionsCount
	return &this
}

// NewUpdatePartitionCountRequestDataWithDefaults instantiates a new UpdatePartitionCountRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePartitionCountRequestDataWithDefaults() *UpdatePartitionCountRequestData {
	this := UpdatePartitionCountRequestData{}
	return &this
}

// GetPartitionsCount returns the PartitionsCount field value
func (o *UpdatePartitionCountRequestData) GetPartitionsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PartitionsCount
}

// GetPartitionsCountOk returns a tuple with the PartitionsCount field value
// and a boolean to check if the value has been set.
func (o *UpdatePartitionCountRequestData) GetPartitionsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartitionsCount, true
}

// SetPartitionsCount sets field value
func (o *UpdatePartitionCountRequestData) SetPartitionsCount(v int32) {
	o.PartitionsCount = v
}

// Redact resets all sensitive fields to their zero value.
func (o *UpdatePartitionCountRequestData) Redact() {
	o.recurseRedact(&o.PartitionsCount)
}

func (o *UpdatePartitionCountRequestData) recurseRedact(v interface{}) {
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

func (o UpdatePartitionCountRequestData) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o UpdatePartitionCountRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["partitions_count"] = o.PartitionsCount
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableUpdatePartitionCountRequestData struct {
	value *UpdatePartitionCountRequestData
	isSet bool
}

func (v NullableUpdatePartitionCountRequestData) Get() *UpdatePartitionCountRequestData {
	return v.value
}

func (v *NullableUpdatePartitionCountRequestData) Set(val *UpdatePartitionCountRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePartitionCountRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePartitionCountRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePartitionCountRequestData(val *UpdatePartitionCountRequestData) *NullableUpdatePartitionCountRequestData {
	return &NullableUpdatePartitionCountRequestData{value: val, isSet: true}
}

func (v NullableUpdatePartitionCountRequestData) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableUpdatePartitionCountRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
