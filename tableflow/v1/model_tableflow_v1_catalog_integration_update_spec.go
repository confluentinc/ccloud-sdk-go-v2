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
Tableflow Management API

Tableflow Management API

API version: 0.0.1
Contact: cts-team@confluent.io
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

// TableflowV1CatalogIntegrationUpdateSpec The desired state of the Catalog Integration
type TableflowV1CatalogIntegrationUpdateSpec struct {
	// The name of the catalog integration
	DisplayName *string `json:"display_name,omitempty"`
	// Indicates whether the Catalog Integration should be suspended.
	Suspended *bool `json:"suspended,omitempty"`
	// The integration config
	Config *TableflowV1CatalogIntegrationUpdateSpecConfigOneOf `json:"config,omitempty"`
	// The environment to which the target Kafka cluster belongs.
	Environment GlobalObjectReference `json:"environment,omitempty"`
	// The kafka cluster of the topic for which Tableflow is enabled
	KafkaCluster EnvScopedObjectReference `json:"kafka_cluster,omitempty"`
}

// NewTableflowV1CatalogIntegrationUpdateSpec instantiates a new TableflowV1CatalogIntegrationUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTableflowV1CatalogIntegrationUpdateSpec(environment GlobalObjectReference, kafkaCluster EnvScopedObjectReference) *TableflowV1CatalogIntegrationUpdateSpec {
	this := TableflowV1CatalogIntegrationUpdateSpec{}
	this.Environment = environment
	this.KafkaCluster = kafkaCluster
	return &this
}

// NewTableflowV1CatalogIntegrationUpdateSpecWithDefaults instantiates a new TableflowV1CatalogIntegrationUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTableflowV1CatalogIntegrationUpdateSpecWithDefaults() *TableflowV1CatalogIntegrationUpdateSpec {
	this := TableflowV1CatalogIntegrationUpdateSpec{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *TableflowV1CatalogIntegrationUpdateSpec) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableflowV1CatalogIntegrationUpdateSpec) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *TableflowV1CatalogIntegrationUpdateSpec) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *TableflowV1CatalogIntegrationUpdateSpec) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetSuspended returns the Suspended field value if set, zero value otherwise.
func (o *TableflowV1CatalogIntegrationUpdateSpec) GetSuspended() bool {
	if o == nil || o.Suspended == nil {
		var ret bool
		return ret
	}
	return *o.Suspended
}

// GetSuspendedOk returns a tuple with the Suspended field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableflowV1CatalogIntegrationUpdateSpec) GetSuspendedOk() (*bool, bool) {
	if o == nil || o.Suspended == nil {
		return nil, false
	}
	return o.Suspended, true
}

// HasSuspended returns a boolean if a field has been set.
func (o *TableflowV1CatalogIntegrationUpdateSpec) HasSuspended() bool {
	if o != nil && o.Suspended != nil {
		return true
	}

	return false
}

// SetSuspended gets a reference to the given bool and assigns it to the Suspended field.
func (o *TableflowV1CatalogIntegrationUpdateSpec) SetSuspended(v bool) {
	o.Suspended = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *TableflowV1CatalogIntegrationUpdateSpec) GetConfig() TableflowV1CatalogIntegrationUpdateSpecConfigOneOf {
	if o == nil || o.Config == nil {
		var ret TableflowV1CatalogIntegrationUpdateSpecConfigOneOf
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableflowV1CatalogIntegrationUpdateSpec) GetConfigOk() (*TableflowV1CatalogIntegrationUpdateSpecConfigOneOf, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *TableflowV1CatalogIntegrationUpdateSpec) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given TableflowV1CatalogIntegrationUpdateSpecConfigOneOf and assigns it to the Config field.
func (o *TableflowV1CatalogIntegrationUpdateSpec) SetConfig(v TableflowV1CatalogIntegrationUpdateSpecConfigOneOf) {
	o.Config = &v
}

// GetEnvironment returns the Environment field value
func (o *TableflowV1CatalogIntegrationUpdateSpec) GetEnvironment() GlobalObjectReference {
	if o == nil {
		var ret GlobalObjectReference
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *TableflowV1CatalogIntegrationUpdateSpec) GetEnvironmentOk() (*GlobalObjectReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *TableflowV1CatalogIntegrationUpdateSpec) SetEnvironment(v GlobalObjectReference) {
	o.Environment = v
}

// GetKafkaCluster returns the KafkaCluster field value
func (o *TableflowV1CatalogIntegrationUpdateSpec) GetKafkaCluster() EnvScopedObjectReference {
	if o == nil {
		var ret EnvScopedObjectReference
		return ret
	}

	return o.KafkaCluster
}

// GetKafkaClusterOk returns a tuple with the KafkaCluster field value
// and a boolean to check if the value has been set.
func (o *TableflowV1CatalogIntegrationUpdateSpec) GetKafkaClusterOk() (*EnvScopedObjectReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KafkaCluster, true
}

// SetKafkaCluster sets field value
func (o *TableflowV1CatalogIntegrationUpdateSpec) SetKafkaCluster(v EnvScopedObjectReference) {
	o.KafkaCluster = v
}

// Redact resets all sensitive fields to their zero value.
func (o *TableflowV1CatalogIntegrationUpdateSpec) Redact() {
	o.recurseRedact(o.DisplayName)
	o.recurseRedact(o.Suspended)
	o.recurseRedact(o.Config)
	o.recurseRedact(&o.Environment)
	o.recurseRedact(&o.KafkaCluster)
}

func (o *TableflowV1CatalogIntegrationUpdateSpec) recurseRedact(v interface{}) {
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

func (o TableflowV1CatalogIntegrationUpdateSpec) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o TableflowV1CatalogIntegrationUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Suspended != nil {
		toSerialize["suspended"] = o.Suspended
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if true {
		toSerialize["environment"] = o.Environment
	}
	if true {
		toSerialize["kafka_cluster"] = o.KafkaCluster
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableTableflowV1CatalogIntegrationUpdateSpec struct {
	value *TableflowV1CatalogIntegrationUpdateSpec
	isSet bool
}

func (v NullableTableflowV1CatalogIntegrationUpdateSpec) Get() *TableflowV1CatalogIntegrationUpdateSpec {
	return v.value
}

func (v *NullableTableflowV1CatalogIntegrationUpdateSpec) Set(val *TableflowV1CatalogIntegrationUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTableflowV1CatalogIntegrationUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTableflowV1CatalogIntegrationUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTableflowV1CatalogIntegrationUpdateSpec(val *TableflowV1CatalogIntegrationUpdateSpec) *NullableTableflowV1CatalogIntegrationUpdateSpec {
	return &NullableTableflowV1CatalogIntegrationUpdateSpec{value: val, isSet: true}
}

func (v NullableTableflowV1CatalogIntegrationUpdateSpec) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableTableflowV1CatalogIntegrationUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
