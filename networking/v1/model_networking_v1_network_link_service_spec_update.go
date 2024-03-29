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

// NetworkingV1NetworkLinkServiceSpecUpdate The desired state of the Network Link Service
type NetworkingV1NetworkLinkServiceSpecUpdate struct {
	// The name of the network link service
	DisplayName *string `json:"display_name,omitempty"`
	// The description of the network link service
	Description *string `json:"description,omitempty"`
	// Network Link Service Accept policy
	Accept *NetworkingV1NetworkLinkServiceAcceptPolicy `json:"accept,omitempty"`
	// The environment to which this belongs.
	Environment *GlobalObjectReference `json:"environment,omitempty"`
}

// NewNetworkingV1NetworkLinkServiceSpecUpdate instantiates a new NetworkingV1NetworkLinkServiceSpecUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1NetworkLinkServiceSpecUpdate() *NetworkingV1NetworkLinkServiceSpecUpdate {
	this := NetworkingV1NetworkLinkServiceSpecUpdate{}
	return &this
}

// NewNetworkingV1NetworkLinkServiceSpecUpdateWithDefaults instantiates a new NetworkingV1NetworkLinkServiceSpecUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1NetworkLinkServiceSpecUpdateWithDefaults() *NetworkingV1NetworkLinkServiceSpecUpdate {
	this := NetworkingV1NetworkLinkServiceSpecUpdate{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkServiceSpecUpdate) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkServiceSpecUpdate) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkServiceSpecUpdate) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *NetworkingV1NetworkLinkServiceSpecUpdate) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkServiceSpecUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkServiceSpecUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkServiceSpecUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NetworkingV1NetworkLinkServiceSpecUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetAccept returns the Accept field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkServiceSpecUpdate) GetAccept() NetworkingV1NetworkLinkServiceAcceptPolicy {
	if o == nil || o.Accept == nil {
		var ret NetworkingV1NetworkLinkServiceAcceptPolicy
		return ret
	}
	return *o.Accept
}

// GetAcceptOk returns a tuple with the Accept field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkServiceSpecUpdate) GetAcceptOk() (*NetworkingV1NetworkLinkServiceAcceptPolicy, bool) {
	if o == nil || o.Accept == nil {
		return nil, false
	}
	return o.Accept, true
}

// HasAccept returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkServiceSpecUpdate) HasAccept() bool {
	if o != nil && o.Accept != nil {
		return true
	}

	return false
}

// SetAccept gets a reference to the given NetworkingV1NetworkLinkServiceAcceptPolicy and assigns it to the Accept field.
func (o *NetworkingV1NetworkLinkServiceSpecUpdate) SetAccept(v NetworkingV1NetworkLinkServiceAcceptPolicy) {
	o.Accept = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkServiceSpecUpdate) GetEnvironment() GlobalObjectReference {
	if o == nil || o.Environment == nil {
		var ret GlobalObjectReference
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkServiceSpecUpdate) GetEnvironmentOk() (*GlobalObjectReference, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkServiceSpecUpdate) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given GlobalObjectReference and assigns it to the Environment field.
func (o *NetworkingV1NetworkLinkServiceSpecUpdate) SetEnvironment(v GlobalObjectReference) {
	o.Environment = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1NetworkLinkServiceSpecUpdate) Redact() {
	o.recurseRedact(o.DisplayName)
	o.recurseRedact(o.Description)
	o.recurseRedact(o.Accept)
	o.recurseRedact(o.Environment)
}

func (o *NetworkingV1NetworkLinkServiceSpecUpdate) recurseRedact(v interface{}) {
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

func (o NetworkingV1NetworkLinkServiceSpecUpdate) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1NetworkLinkServiceSpecUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Accept != nil {
		toSerialize["accept"] = o.Accept
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableNetworkingV1NetworkLinkServiceSpecUpdate struct {
	value *NetworkingV1NetworkLinkServiceSpecUpdate
	isSet bool
}

func (v NullableNetworkingV1NetworkLinkServiceSpecUpdate) Get() *NetworkingV1NetworkLinkServiceSpecUpdate {
	return v.value
}

func (v *NullableNetworkingV1NetworkLinkServiceSpecUpdate) Set(val *NetworkingV1NetworkLinkServiceSpecUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1NetworkLinkServiceSpecUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1NetworkLinkServiceSpecUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1NetworkLinkServiceSpecUpdate(val *NetworkingV1NetworkLinkServiceSpecUpdate) *NullableNetworkingV1NetworkLinkServiceSpecUpdate {
	return &NullableNetworkingV1NetworkLinkServiceSpecUpdate{value: val, isSet: true}
}

func (v NullableNetworkingV1NetworkLinkServiceSpecUpdate) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableNetworkingV1NetworkLinkServiceSpecUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
