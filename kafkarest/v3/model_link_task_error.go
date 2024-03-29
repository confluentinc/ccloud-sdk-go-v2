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

// LinkTaskError struct for LinkTaskError
type LinkTaskError struct {
	ErrorCode    string `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

// NewLinkTaskError instantiates a new LinkTaskError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTaskError(errorCode string, errorMessage string) *LinkTaskError {
	this := LinkTaskError{}
	this.ErrorCode = errorCode
	this.ErrorMessage = errorMessage
	return &this
}

// NewLinkTaskErrorWithDefaults instantiates a new LinkTaskError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTaskErrorWithDefaults() *LinkTaskError {
	this := LinkTaskError{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *LinkTaskError) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *LinkTaskError) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *LinkTaskError) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetErrorMessage returns the ErrorMessage field value
func (o *LinkTaskError) GetErrorMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value
// and a boolean to check if the value has been set.
func (o *LinkTaskError) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorMessage, true
}

// SetErrorMessage sets field value
func (o *LinkTaskError) SetErrorMessage(v string) {
	o.ErrorMessage = v
}

// Redact resets all sensitive fields to their zero value.
func (o *LinkTaskError) Redact() {
	o.recurseRedact(&o.ErrorCode)
	o.recurseRedact(&o.ErrorMessage)
}

func (o *LinkTaskError) recurseRedact(v interface{}) {
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

func (o LinkTaskError) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o LinkTaskError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["error_code"] = o.ErrorCode
	}
	if true {
		toSerialize["error_message"] = o.ErrorMessage
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableLinkTaskError struct {
	value *LinkTaskError
	isSet bool
}

func (v NullableLinkTaskError) Get() *LinkTaskError {
	return v.value
}

func (v *NullableLinkTaskError) Set(val *LinkTaskError) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTaskError) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTaskError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTaskError(val *LinkTaskError) *NullableLinkTaskError {
	return &NullableLinkTaskError{value: val, isSet: true}
}

func (v NullableLinkTaskError) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableLinkTaskError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
