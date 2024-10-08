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

Network Gateway API

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

// NetworkingV1GatewaySpecUpdate The desired state of the Gateway
type NetworkingV1GatewaySpecUpdate struct {
	// The name of the gateway.
	DisplayName *string `json:"display_name,omitempty"`
	// The environment to which this belongs.
	Environment *ObjectReference `json:"environment,omitempty"`
}

// NewNetworkingV1GatewaySpecUpdate instantiates a new NetworkingV1GatewaySpecUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1GatewaySpecUpdate() *NetworkingV1GatewaySpecUpdate {
	this := NetworkingV1GatewaySpecUpdate{}
	return &this
}

// NewNetworkingV1GatewaySpecUpdateWithDefaults instantiates a new NetworkingV1GatewaySpecUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1GatewaySpecUpdateWithDefaults() *NetworkingV1GatewaySpecUpdate {
	this := NetworkingV1GatewaySpecUpdate{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *NetworkingV1GatewaySpecUpdate) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1GatewaySpecUpdate) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *NetworkingV1GatewaySpecUpdate) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *NetworkingV1GatewaySpecUpdate) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *NetworkingV1GatewaySpecUpdate) GetEnvironment() ObjectReference {
	if o == nil || o.Environment == nil {
		var ret ObjectReference
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1GatewaySpecUpdate) GetEnvironmentOk() (*ObjectReference, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *NetworkingV1GatewaySpecUpdate) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectReference and assigns it to the Environment field.
func (o *NetworkingV1GatewaySpecUpdate) SetEnvironment(v ObjectReference) {
	o.Environment = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1GatewaySpecUpdate) Redact() {
	o.recurseRedact(o.DisplayName)
	o.recurseRedact(o.Environment)
}

func (o *NetworkingV1GatewaySpecUpdate) recurseRedact(v interface{}) {
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

func (o NetworkingV1GatewaySpecUpdate) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1GatewaySpecUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
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

type NullableNetworkingV1GatewaySpecUpdate struct {
	value *NetworkingV1GatewaySpecUpdate
	isSet bool
}

func (v NullableNetworkingV1GatewaySpecUpdate) Get() *NetworkingV1GatewaySpecUpdate {
	return v.value
}

func (v *NullableNetworkingV1GatewaySpecUpdate) Set(val *NetworkingV1GatewaySpecUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1GatewaySpecUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1GatewaySpecUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1GatewaySpecUpdate(val *NetworkingV1GatewaySpecUpdate) *NullableNetworkingV1GatewaySpecUpdate {
	return &NullableNetworkingV1GatewaySpecUpdate{value: val, isSet: true}
}

func (v NullableNetworkingV1GatewaySpecUpdate) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableNetworkingV1GatewaySpecUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
