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

// InlineResponse200 struct for InlineResponse200
type InlineResponse200 struct {
	ConsumerResources *[]CdxV1ConsumerSharedResource `json:"consumer_resources,omitempty"`
}

// NewInlineResponse200 instantiates a new InlineResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200() *InlineResponse200 {
	this := InlineResponse200{}
	return &this
}

// NewInlineResponse200WithDefaults instantiates a new InlineResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200WithDefaults() *InlineResponse200 {
	this := InlineResponse200{}
	return &this
}

// GetConsumerResources returns the ConsumerResources field value if set, zero value otherwise.
func (o *InlineResponse200) GetConsumerResources() []CdxV1ConsumerSharedResource {
	if o == nil || o.ConsumerResources == nil {
		var ret []CdxV1ConsumerSharedResource
		return ret
	}
	return *o.ConsumerResources
}

// GetConsumerResourcesOk returns a tuple with the ConsumerResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetConsumerResourcesOk() (*[]CdxV1ConsumerSharedResource, bool) {
	if o == nil || o.ConsumerResources == nil {
		return nil, false
	}
	return o.ConsumerResources, true
}

// HasConsumerResources returns a boolean if a field has been set.
func (o *InlineResponse200) HasConsumerResources() bool {
	if o != nil && o.ConsumerResources != nil {
		return true
	}

	return false
}

// SetConsumerResources gets a reference to the given []CdxV1ConsumerSharedResource and assigns it to the ConsumerResources field.
func (o *InlineResponse200) SetConsumerResources(v []CdxV1ConsumerSharedResource) {
	o.ConsumerResources = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *InlineResponse200) Redact() {
    o.recurseRedact(o.ConsumerResources)
}

func (o *InlineResponse200) recurseRedact(v interface{}) {
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

func (o InlineResponse200) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o InlineResponse200) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConsumerResources != nil {
		toSerialize["consumer_resources"] = o.ConsumerResources
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200 struct {
	value *InlineResponse200
	isSet bool
}

func (v NullableInlineResponse200) Get() *InlineResponse200 {
	return v.value
}

func (v *NullableInlineResponse200) Set(val *InlineResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200(val *InlineResponse200) *NullableInlineResponse200 {
	return &NullableInlineResponse200{value: val, isSet: true}
}

func (v NullableInlineResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

