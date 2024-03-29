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

// ConsumerAssignmentDataListAllOf struct for ConsumerAssignmentDataListAllOf
type ConsumerAssignmentDataListAllOf struct {
	Data []ConsumerAssignmentData `json:"data,omitempty"`
}

// NewConsumerAssignmentDataListAllOf instantiates a new ConsumerAssignmentDataListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumerAssignmentDataListAllOf(data []ConsumerAssignmentData) *ConsumerAssignmentDataListAllOf {
	this := ConsumerAssignmentDataListAllOf{}
	this.Data = data
	return &this
}

// NewConsumerAssignmentDataListAllOfWithDefaults instantiates a new ConsumerAssignmentDataListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumerAssignmentDataListAllOfWithDefaults() *ConsumerAssignmentDataListAllOf {
	this := ConsumerAssignmentDataListAllOf{}
	return &this
}

// GetData returns the Data field value
func (o *ConsumerAssignmentDataListAllOf) GetData() []ConsumerAssignmentData {
	if o == nil {
		var ret []ConsumerAssignmentData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ConsumerAssignmentDataListAllOf) GetDataOk() (*[]ConsumerAssignmentData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ConsumerAssignmentDataListAllOf) SetData(v []ConsumerAssignmentData) {
	o.Data = v
}

// Redact resets all sensitive fields to their zero value.
func (o *ConsumerAssignmentDataListAllOf) Redact() {
	o.recurseRedact(&o.Data)
}

func (o *ConsumerAssignmentDataListAllOf) recurseRedact(v interface{}) {
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

func (o ConsumerAssignmentDataListAllOf) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ConsumerAssignmentDataListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableConsumerAssignmentDataListAllOf struct {
	value *ConsumerAssignmentDataListAllOf
	isSet bool
}

func (v NullableConsumerAssignmentDataListAllOf) Get() *ConsumerAssignmentDataListAllOf {
	return v.value
}

func (v *NullableConsumerAssignmentDataListAllOf) Set(val *ConsumerAssignmentDataListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumerAssignmentDataListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumerAssignmentDataListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumerAssignmentDataListAllOf(val *ConsumerAssignmentDataListAllOf) *NullableConsumerAssignmentDataListAllOf {
	return &NullableConsumerAssignmentDataListAllOf{value: val, isSet: true}
}

func (v NullableConsumerAssignmentDataListAllOf) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableConsumerAssignmentDataListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
