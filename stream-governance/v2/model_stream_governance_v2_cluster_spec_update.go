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
Cluster Management for Stream Governance API

Cluster Management for Stream Governance API

API version: 0.0.1-alpha1
Contact: data-governance@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

import (
	"reflect"
)

// StreamGovernanceV2ClusterSpecUpdate The desired state of the Cluster
type StreamGovernanceV2ClusterSpecUpdate struct {
	// The billing package.
	Package *string `json:"package,omitempty"`
	// The environment to which this belongs.
	Environment *ObjectReference `json:"environment,omitempty"`
}

// NewStreamGovernanceV2ClusterSpecUpdate instantiates a new StreamGovernanceV2ClusterSpecUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamGovernanceV2ClusterSpecUpdate() *StreamGovernanceV2ClusterSpecUpdate {
	this := StreamGovernanceV2ClusterSpecUpdate{}
	return &this
}

// NewStreamGovernanceV2ClusterSpecUpdateWithDefaults instantiates a new StreamGovernanceV2ClusterSpecUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamGovernanceV2ClusterSpecUpdateWithDefaults() *StreamGovernanceV2ClusterSpecUpdate {
	this := StreamGovernanceV2ClusterSpecUpdate{}
	return &this
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *StreamGovernanceV2ClusterSpecUpdate) GetPackage() string {
	if o == nil || o.Package == nil {
		var ret string
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamGovernanceV2ClusterSpecUpdate) GetPackageOk() (*string, bool) {
	if o == nil || o.Package == nil {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *StreamGovernanceV2ClusterSpecUpdate) HasPackage() bool {
	if o != nil && o.Package != nil {
		return true
	}

	return false
}

// SetPackage gets a reference to the given string and assigns it to the Package field.
func (o *StreamGovernanceV2ClusterSpecUpdate) SetPackage(v string) {
	o.Package = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *StreamGovernanceV2ClusterSpecUpdate) GetEnvironment() ObjectReference {
	if o == nil || o.Environment == nil {
		var ret ObjectReference
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamGovernanceV2ClusterSpecUpdate) GetEnvironmentOk() (*ObjectReference, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *StreamGovernanceV2ClusterSpecUpdate) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectReference and assigns it to the Environment field.
func (o *StreamGovernanceV2ClusterSpecUpdate) SetEnvironment(v ObjectReference) {
	o.Environment = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *StreamGovernanceV2ClusterSpecUpdate) Redact() {
    o.recurseRedact(o.Package)
    o.recurseRedact(o.Environment)
}

func (o *StreamGovernanceV2ClusterSpecUpdate) recurseRedact(v interface{}) {
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

func (o StreamGovernanceV2ClusterSpecUpdate) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o StreamGovernanceV2ClusterSpecUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Package != nil {
		toSerialize["package"] = o.Package
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	return json.Marshal(toSerialize)
}

type NullableStreamGovernanceV2ClusterSpecUpdate struct {
	value *StreamGovernanceV2ClusterSpecUpdate
	isSet bool
}

func (v NullableStreamGovernanceV2ClusterSpecUpdate) Get() *StreamGovernanceV2ClusterSpecUpdate {
	return v.value
}

func (v *NullableStreamGovernanceV2ClusterSpecUpdate) Set(val *StreamGovernanceV2ClusterSpecUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamGovernanceV2ClusterSpecUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamGovernanceV2ClusterSpecUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamGovernanceV2ClusterSpecUpdate(val *StreamGovernanceV2ClusterSpecUpdate) *NullableStreamGovernanceV2ClusterSpecUpdate {
	return &NullableStreamGovernanceV2ClusterSpecUpdate{value: val, isSet: true}
}

func (v NullableStreamGovernanceV2ClusterSpecUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamGovernanceV2ClusterSpecUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

