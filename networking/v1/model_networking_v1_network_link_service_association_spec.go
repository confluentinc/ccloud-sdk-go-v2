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

// NetworkingV1NetworkLinkServiceAssociationSpec The desired state of the Network Link Service Association
type NetworkingV1NetworkLinkServiceAssociationSpec struct {
	// The name of the network link endpoint
	DisplayName *string `json:"display_name,omitempty"`
	// The description of the network link endpoint
	Description *string `json:"description,omitempty"`
	// ID of the Network link endpoint.
	NetworkLinkEndpoint *string `json:"network_link_endpoint,omitempty"`
	// The network_link_service to which this belongs.
	NetworkLinkService *EnvScopedObjectReference `json:"network_link_service,omitempty"`
	// The environment to which this belongs.
	Environment *GlobalObjectReference `json:"environment,omitempty"`
}

// NewNetworkingV1NetworkLinkServiceAssociationSpec instantiates a new NetworkingV1NetworkLinkServiceAssociationSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1NetworkLinkServiceAssociationSpec() *NetworkingV1NetworkLinkServiceAssociationSpec {
	this := NetworkingV1NetworkLinkServiceAssociationSpec{}
	return &this
}

// NewNetworkingV1NetworkLinkServiceAssociationSpecWithDefaults instantiates a new NetworkingV1NetworkLinkServiceAssociationSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1NetworkLinkServiceAssociationSpecWithDefaults() *NetworkingV1NetworkLinkServiceAssociationSpec {
	this := NetworkingV1NetworkLinkServiceAssociationSpec{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkServiceAssociationSpec) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkServiceAssociationSpec) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkServiceAssociationSpec) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *NetworkingV1NetworkLinkServiceAssociationSpec) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkServiceAssociationSpec) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkServiceAssociationSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkServiceAssociationSpec) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NetworkingV1NetworkLinkServiceAssociationSpec) SetDescription(v string) {
	o.Description = &v
}

// GetNetworkLinkEndpoint returns the NetworkLinkEndpoint field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkServiceAssociationSpec) GetNetworkLinkEndpoint() string {
	if o == nil || o.NetworkLinkEndpoint == nil {
		var ret string
		return ret
	}
	return *o.NetworkLinkEndpoint
}

// GetNetworkLinkEndpointOk returns a tuple with the NetworkLinkEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkServiceAssociationSpec) GetNetworkLinkEndpointOk() (*string, bool) {
	if o == nil || o.NetworkLinkEndpoint == nil {
		return nil, false
	}
	return o.NetworkLinkEndpoint, true
}

// HasNetworkLinkEndpoint returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkServiceAssociationSpec) HasNetworkLinkEndpoint() bool {
	if o != nil && o.NetworkLinkEndpoint != nil {
		return true
	}

	return false
}

// SetNetworkLinkEndpoint gets a reference to the given string and assigns it to the NetworkLinkEndpoint field.
func (o *NetworkingV1NetworkLinkServiceAssociationSpec) SetNetworkLinkEndpoint(v string) {
	o.NetworkLinkEndpoint = &v
}

// GetNetworkLinkService returns the NetworkLinkService field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkServiceAssociationSpec) GetNetworkLinkService() EnvScopedObjectReference {
	if o == nil || o.NetworkLinkService == nil {
		var ret EnvScopedObjectReference
		return ret
	}
	return *o.NetworkLinkService
}

// GetNetworkLinkServiceOk returns a tuple with the NetworkLinkService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkServiceAssociationSpec) GetNetworkLinkServiceOk() (*EnvScopedObjectReference, bool) {
	if o == nil || o.NetworkLinkService == nil {
		return nil, false
	}
	return o.NetworkLinkService, true
}

// HasNetworkLinkService returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkServiceAssociationSpec) HasNetworkLinkService() bool {
	if o != nil && o.NetworkLinkService != nil {
		return true
	}

	return false
}

// SetNetworkLinkService gets a reference to the given EnvScopedObjectReference and assigns it to the NetworkLinkService field.
func (o *NetworkingV1NetworkLinkServiceAssociationSpec) SetNetworkLinkService(v EnvScopedObjectReference) {
	o.NetworkLinkService = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkServiceAssociationSpec) GetEnvironment() GlobalObjectReference {
	if o == nil || o.Environment == nil {
		var ret GlobalObjectReference
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkServiceAssociationSpec) GetEnvironmentOk() (*GlobalObjectReference, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkServiceAssociationSpec) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given GlobalObjectReference and assigns it to the Environment field.
func (o *NetworkingV1NetworkLinkServiceAssociationSpec) SetEnvironment(v GlobalObjectReference) {
	o.Environment = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1NetworkLinkServiceAssociationSpec) Redact() {
	o.recurseRedact(o.DisplayName)
	o.recurseRedact(o.Description)
	o.recurseRedact(o.NetworkLinkEndpoint)
	o.recurseRedact(o.NetworkLinkService)
	o.recurseRedact(o.Environment)
}

func (o *NetworkingV1NetworkLinkServiceAssociationSpec) recurseRedact(v interface{}) {
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

func (o NetworkingV1NetworkLinkServiceAssociationSpec) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1NetworkLinkServiceAssociationSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.NetworkLinkEndpoint != nil {
		toSerialize["network_link_endpoint"] = o.NetworkLinkEndpoint
	}
	if o.NetworkLinkService != nil {
		toSerialize["network_link_service"] = o.NetworkLinkService
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

type NullableNetworkingV1NetworkLinkServiceAssociationSpec struct {
	value *NetworkingV1NetworkLinkServiceAssociationSpec
	isSet bool
}

func (v NullableNetworkingV1NetworkLinkServiceAssociationSpec) Get() *NetworkingV1NetworkLinkServiceAssociationSpec {
	return v.value
}

func (v *NullableNetworkingV1NetworkLinkServiceAssociationSpec) Set(val *NetworkingV1NetworkLinkServiceAssociationSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1NetworkLinkServiceAssociationSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1NetworkLinkServiceAssociationSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1NetworkLinkServiceAssociationSpec(val *NetworkingV1NetworkLinkServiceAssociationSpec) *NullableNetworkingV1NetworkLinkServiceAssociationSpec {
	return &NullableNetworkingV1NetworkLinkServiceAssociationSpec{value: val, isSet: true}
}

func (v NullableNetworkingV1NetworkLinkServiceAssociationSpec) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableNetworkingV1NetworkLinkServiceAssociationSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
