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

// BrokerConfigDataList struct for BrokerConfigDataList
type BrokerConfigDataList struct {
	Kind     string                     `json:"kind,omitempty"`
	Metadata ResourceCollectionMetadata `json:"metadata,omitempty"`
	Data     []BrokerConfigData         `json:"data,omitempty"`
}

// NewBrokerConfigDataList instantiates a new BrokerConfigDataList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrokerConfigDataList(kind string, metadata ResourceCollectionMetadata, data []BrokerConfigData) *BrokerConfigDataList {
	this := BrokerConfigDataList{}
	this.Kind = kind
	this.Metadata = metadata
	this.Data = data
	return &this
}

// NewBrokerConfigDataListWithDefaults instantiates a new BrokerConfigDataList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrokerConfigDataListWithDefaults() *BrokerConfigDataList {
	this := BrokerConfigDataList{}
	return &this
}

// GetKind returns the Kind field value
func (o *BrokerConfigDataList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *BrokerConfigDataList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *BrokerConfigDataList) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *BrokerConfigDataList) GetMetadata() ResourceCollectionMetadata {
	if o == nil {
		var ret ResourceCollectionMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *BrokerConfigDataList) GetMetadataOk() (*ResourceCollectionMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *BrokerConfigDataList) SetMetadata(v ResourceCollectionMetadata) {
	o.Metadata = v
}

// GetData returns the Data field value
func (o *BrokerConfigDataList) GetData() []BrokerConfigData {
	if o == nil {
		var ret []BrokerConfigData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BrokerConfigDataList) GetDataOk() (*[]BrokerConfigData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BrokerConfigDataList) SetData(v []BrokerConfigData) {
	o.Data = v
}

// Redact resets all sensitive fields to their zero value.
func (o *BrokerConfigDataList) Redact() {
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Metadata)
	o.recurseRedact(&o.Data)
}

func (o *BrokerConfigDataList) recurseRedact(v interface{}) {
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

func (o BrokerConfigDataList) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o BrokerConfigDataList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["data"] = o.Data
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableBrokerConfigDataList struct {
	value *BrokerConfigDataList
	isSet bool
}

func (v NullableBrokerConfigDataList) Get() *BrokerConfigDataList {
	return v.value
}

func (v *NullableBrokerConfigDataList) Set(val *BrokerConfigDataList) {
	v.value = val
	v.isSet = true
}

func (v NullableBrokerConfigDataList) IsSet() bool {
	return v.isSet
}

func (v *NullableBrokerConfigDataList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrokerConfigDataList(val *BrokerConfigDataList) *NullableBrokerConfigDataList {
	return &NullableBrokerConfigDataList{value: val, isSet: true}
}

func (v NullableBrokerConfigDataList) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableBrokerConfigDataList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
