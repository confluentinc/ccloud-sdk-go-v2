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
Stream Share APIs

# Introduction

API version: 0.1.0-alpha0
Contact: cdx@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

import (
	"reflect"
)

// CdxV1EmailConsumerRestriction Consumer restrictions limits by authenticated user's email
type CdxV1EmailConsumerRestriction struct {
	// The resource kind
	Kind string `json:"kind"`
	// Email based matching for the consumers
	Email string `json:"email"`
}

// NewCdxV1EmailConsumerRestriction instantiates a new CdxV1EmailConsumerRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdxV1EmailConsumerRestriction(kind string, email string) *CdxV1EmailConsumerRestriction {
	this := CdxV1EmailConsumerRestriction{}
	this.Kind = kind
	this.Email = email
	return &this
}

// NewCdxV1EmailConsumerRestrictionWithDefaults instantiates a new CdxV1EmailConsumerRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdxV1EmailConsumerRestrictionWithDefaults() *CdxV1EmailConsumerRestriction {
	this := CdxV1EmailConsumerRestriction{}
	return &this
}

// GetKind returns the Kind field value
func (o *CdxV1EmailConsumerRestriction) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *CdxV1EmailConsumerRestriction) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *CdxV1EmailConsumerRestriction) SetKind(v string) {
	o.Kind = v
}

// GetEmail returns the Email field value
func (o *CdxV1EmailConsumerRestriction) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CdxV1EmailConsumerRestriction) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CdxV1EmailConsumerRestriction) SetEmail(v string) {
	o.Email = v
}

// Redact resets all sensitive fields to their zero value.
func (o *CdxV1EmailConsumerRestriction) Redact() {
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Email)
}

func (o *CdxV1EmailConsumerRestriction) recurseRedact(v interface{}) {
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

func (o CdxV1EmailConsumerRestriction) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o CdxV1EmailConsumerRestriction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableCdxV1EmailConsumerRestriction struct {
	value *CdxV1EmailConsumerRestriction
	isSet bool
}

func (v NullableCdxV1EmailConsumerRestriction) Get() *CdxV1EmailConsumerRestriction {
	return v.value
}

func (v *NullableCdxV1EmailConsumerRestriction) Set(val *CdxV1EmailConsumerRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableCdxV1EmailConsumerRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableCdxV1EmailConsumerRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdxV1EmailConsumerRestriction(val *CdxV1EmailConsumerRestriction) *NullableCdxV1EmailConsumerRestriction {
	return &NullableCdxV1EmailConsumerRestriction{value: val, isSet: true}
}

func (v NullableCdxV1EmailConsumerRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdxV1EmailConsumerRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
