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
Cluster Management for Apache Kafka API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1-alpha1
Contact: cpk@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// CmkV2ClusterSpecUpdate The desired state of the Cluster
type CmkV2ClusterSpecUpdate struct {
	// The name of the cluster.
	DisplayName *string `json:"display_name,omitempty"`
	// The availability zone configuration of the cluster
	Availability *string `json:"availability,omitempty"`
	// The configuration of the Kafka cluster.  Note: Clusters can be upgraded from Basic to Standard, but cannot be downgraded from Standard to Basic.
	Config *CmkV2ClusterSpecUpdateConfigOneOf `json:"config,omitempty"`
	// A map of endpoints for connecting to the Kafka cluster, keyed by access_point_id. Access Point ID 'public' and 'privatelink' are reserved. These can be used for different network access methods or regions.
	Endpoints *ModelMap `json:"endpoints,omitempty"`
	// The environment to which this belongs.
	Environment *EnvScopedObjectReference `json:"environment,omitempty"`
}

// NewCmkV2ClusterSpecUpdate instantiates a new CmkV2ClusterSpecUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCmkV2ClusterSpecUpdate() *CmkV2ClusterSpecUpdate {
	this := CmkV2ClusterSpecUpdate{}
	return &this
}

// NewCmkV2ClusterSpecUpdateWithDefaults instantiates a new CmkV2ClusterSpecUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCmkV2ClusterSpecUpdateWithDefaults() *CmkV2ClusterSpecUpdate {
	this := CmkV2ClusterSpecUpdate{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CmkV2ClusterSpecUpdate) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmkV2ClusterSpecUpdate) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CmkV2ClusterSpecUpdate) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CmkV2ClusterSpecUpdate) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetAvailability returns the Availability field value if set, zero value otherwise.
func (o *CmkV2ClusterSpecUpdate) GetAvailability() string {
	if o == nil || o.Availability == nil {
		var ret string
		return ret
	}
	return *o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmkV2ClusterSpecUpdate) GetAvailabilityOk() (*string, bool) {
	if o == nil || o.Availability == nil {
		return nil, false
	}
	return o.Availability, true
}

// HasAvailability returns a boolean if a field has been set.
func (o *CmkV2ClusterSpecUpdate) HasAvailability() bool {
	if o != nil && o.Availability != nil {
		return true
	}

	return false
}

// SetAvailability gets a reference to the given string and assigns it to the Availability field.
func (o *CmkV2ClusterSpecUpdate) SetAvailability(v string) {
	o.Availability = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *CmkV2ClusterSpecUpdate) GetConfig() CmkV2ClusterSpecUpdateConfigOneOf {
	if o == nil || o.Config == nil {
		var ret CmkV2ClusterSpecUpdateConfigOneOf
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmkV2ClusterSpecUpdate) GetConfigOk() (*CmkV2ClusterSpecUpdateConfigOneOf, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *CmkV2ClusterSpecUpdate) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given CmkV2ClusterSpecUpdateConfigOneOf and assigns it to the Config field.
func (o *CmkV2ClusterSpecUpdate) SetConfig(v CmkV2ClusterSpecUpdateConfigOneOf) {
	o.Config = &v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *CmkV2ClusterSpecUpdate) GetEndpoints() ModelMap {
	if o == nil || o.Endpoints == nil {
		var ret ModelMap
		return ret
	}
	return *o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmkV2ClusterSpecUpdate) GetEndpointsOk() (*ModelMap, bool) {
	if o == nil || o.Endpoints == nil {
		return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *CmkV2ClusterSpecUpdate) HasEndpoints() bool {
	if o != nil && o.Endpoints != nil {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given ModelMap and assigns it to the Endpoints field.
func (o *CmkV2ClusterSpecUpdate) SetEndpoints(v ModelMap) {
	o.Endpoints = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *CmkV2ClusterSpecUpdate) GetEnvironment() EnvScopedObjectReference {
	if o == nil || o.Environment == nil {
		var ret EnvScopedObjectReference
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmkV2ClusterSpecUpdate) GetEnvironmentOk() (*EnvScopedObjectReference, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *CmkV2ClusterSpecUpdate) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given EnvScopedObjectReference and assigns it to the Environment field.
func (o *CmkV2ClusterSpecUpdate) SetEnvironment(v EnvScopedObjectReference) {
	o.Environment = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *CmkV2ClusterSpecUpdate) Redact() {
	o.recurseRedact(o.DisplayName)
	o.recurseRedact(o.Availability)
	o.recurseRedact(o.Config)
	o.recurseRedact(o.Endpoints)
	o.recurseRedact(o.Environment)
}

func (o *CmkV2ClusterSpecUpdate) recurseRedact(v interface{}) {
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

func (o CmkV2ClusterSpecUpdate) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o CmkV2ClusterSpecUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Availability != nil {
		toSerialize["availability"] = o.Availability
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Endpoints != nil {
		toSerialize["endpoints"] = o.Endpoints
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

type NullableCmkV2ClusterSpecUpdate struct {
	value *CmkV2ClusterSpecUpdate
	isSet bool
}

func (v NullableCmkV2ClusterSpecUpdate) Get() *CmkV2ClusterSpecUpdate {
	return v.value
}

func (v *NullableCmkV2ClusterSpecUpdate) Set(val *CmkV2ClusterSpecUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableCmkV2ClusterSpecUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableCmkV2ClusterSpecUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmkV2ClusterSpecUpdate(val *CmkV2ClusterSpecUpdate) *NullableCmkV2ClusterSpecUpdate {
	return &NullableCmkV2ClusterSpecUpdate{value: val, isSet: true}
}

func (v NullableCmkV2ClusterSpecUpdate) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableCmkV2ClusterSpecUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
