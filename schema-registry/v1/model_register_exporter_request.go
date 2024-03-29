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

// RegisterExporterRequest Exporter register request
type RegisterExporterRequest struct {
	// References to other schemas
	References *[]ExporterReference `json:"references,omitempty"`
}

// NewRegisterExporterRequest instantiates a new RegisterExporterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterExporterRequest() *RegisterExporterRequest {
	this := RegisterExporterRequest{}
	return &this
}

// NewRegisterExporterRequestWithDefaults instantiates a new RegisterExporterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterExporterRequestWithDefaults() *RegisterExporterRequest {
	this := RegisterExporterRequest{}
	return &this
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *RegisterExporterRequest) GetReferences() []ExporterReference {
	if o == nil || o.References == nil {
		var ret []ExporterReference
		return ret
	}
	return *o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterExporterRequest) GetReferencesOk() (*[]ExporterReference, bool) {
	if o == nil || o.References == nil {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *RegisterExporterRequest) HasReferences() bool {
	if o != nil && o.References != nil {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []ExporterReference and assigns it to the References field.
func (o *RegisterExporterRequest) SetReferences(v []ExporterReference) {
	o.References = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *RegisterExporterRequest) Redact() {
	o.recurseRedact(o.References)
}

func (o *RegisterExporterRequest) recurseRedact(v interface{}) {
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

func (o RegisterExporterRequest) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o RegisterExporterRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.References != nil {
		toSerialize["references"] = o.References
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableRegisterExporterRequest struct {
	value *RegisterExporterRequest
	isSet bool
}

func (v NullableRegisterExporterRequest) Get() *RegisterExporterRequest {
	return v.value
}

func (v *NullableRegisterExporterRequest) Set(val *RegisterExporterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterExporterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterExporterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterExporterRequest(val *RegisterExporterRequest) *NullableRegisterExporterRequest {
	return &NullableRegisterExporterRequest{value: val, isSet: true}
}

func (v NullableRegisterExporterRequest) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableRegisterExporterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
