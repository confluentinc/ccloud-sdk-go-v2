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
Workspaces API v1beta1

The Workspace API for Flink

API version: 0.0.1
Contact: stream-designer@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1beta1

import (
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// WsV1beta1Block A block within a workspace.
type WsV1beta1Block struct {
	// The type of block.
	Type        *string               `json:"type,omitempty"`
	CodeOptions *WsV1beta1CodeOptions `json:"code_options,omitempty"`
}

// NewWsV1beta1Block instantiates a new WsV1beta1Block object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWsV1beta1Block() *WsV1beta1Block {
	this := WsV1beta1Block{}
	return &this
}

// NewWsV1beta1BlockWithDefaults instantiates a new WsV1beta1Block object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWsV1beta1BlockWithDefaults() *WsV1beta1Block {
	this := WsV1beta1Block{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WsV1beta1Block) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsV1beta1Block) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WsV1beta1Block) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WsV1beta1Block) SetType(v string) {
	o.Type = &v
}

// GetCodeOptions returns the CodeOptions field value if set, zero value otherwise.
func (o *WsV1beta1Block) GetCodeOptions() WsV1beta1CodeOptions {
	if o == nil || o.CodeOptions == nil {
		var ret WsV1beta1CodeOptions
		return ret
	}
	return *o.CodeOptions
}

// GetCodeOptionsOk returns a tuple with the CodeOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsV1beta1Block) GetCodeOptionsOk() (*WsV1beta1CodeOptions, bool) {
	if o == nil || o.CodeOptions == nil {
		return nil, false
	}
	return o.CodeOptions, true
}

// HasCodeOptions returns a boolean if a field has been set.
func (o *WsV1beta1Block) HasCodeOptions() bool {
	if o != nil && o.CodeOptions != nil {
		return true
	}

	return false
}

// SetCodeOptions gets a reference to the given WsV1beta1CodeOptions and assigns it to the CodeOptions field.
func (o *WsV1beta1Block) SetCodeOptions(v WsV1beta1CodeOptions) {
	o.CodeOptions = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *WsV1beta1Block) Redact() {
	o.recurseRedact(o.Type)
	o.recurseRedact(o.CodeOptions)
}

func (o *WsV1beta1Block) recurseRedact(v interface{}) {
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

func (o WsV1beta1Block) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o WsV1beta1Block) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.CodeOptions != nil {
		toSerialize["code_options"] = o.CodeOptions
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableWsV1beta1Block struct {
	value *WsV1beta1Block
	isSet bool
}

func (v NullableWsV1beta1Block) Get() *WsV1beta1Block {
	return v.value
}

func (v *NullableWsV1beta1Block) Set(val *WsV1beta1Block) {
	v.value = val
	v.isSet = true
}

func (v NullableWsV1beta1Block) IsSet() bool {
	return v.isSet
}

func (v *NullableWsV1beta1Block) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWsV1beta1Block(val *WsV1beta1Block) *NullableWsV1beta1Block {
	return &NullableWsV1beta1Block{value: val, isSet: true}
}

func (v NullableWsV1beta1Block) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableWsV1beta1Block) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}