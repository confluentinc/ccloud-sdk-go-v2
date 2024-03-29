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
Network API

Network API

API version: 0.0.1-alpha1
Contact: cire-traffic@confluent.io
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

// NetworkingV1ZonesInfo Cloud provider zones metadata.
type NetworkingV1ZonesInfo struct {
	Items []NetworkingV1ZoneInfo
}

// NewNetworkingV1ZonesInfo instantiates a new NetworkingV1ZonesInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1ZonesInfo() *NetworkingV1ZonesInfo {
	this := NetworkingV1ZonesInfo{}
	return &this
}

// NewNetworkingV1ZonesInfoWithDefaults instantiates a new NetworkingV1ZonesInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1ZonesInfoWithDefaults() *NetworkingV1ZonesInfo {
	this := NetworkingV1ZonesInfo{}
	return &this
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1ZonesInfo) Redact() {
}

func (o *NetworkingV1ZonesInfo) recurseRedact(v interface{}) {
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

func (o NetworkingV1ZonesInfo) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1ZonesInfo) MarshalJSON() ([]byte, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

func (o *NetworkingV1ZonesInfo) UnmarshalJSON(bytes []byte) (err error) {
	return json.Unmarshal(bytes, &o.Items)
}

type NullableNetworkingV1ZonesInfo struct {
	value *NetworkingV1ZonesInfo
	isSet bool
}

func (v NullableNetworkingV1ZonesInfo) Get() *NetworkingV1ZonesInfo {
	return v.value
}

func (v *NullableNetworkingV1ZonesInfo) Set(val *NetworkingV1ZonesInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1ZonesInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1ZonesInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1ZonesInfo(val *NetworkingV1ZonesInfo) *NullableNetworkingV1ZonesInfo {
	return &NullableNetworkingV1ZonesInfo{value: val, isSet: true}
}

func (v NullableNetworkingV1ZonesInfo) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableNetworkingV1ZonesInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
