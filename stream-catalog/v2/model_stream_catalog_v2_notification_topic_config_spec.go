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
Data Catalog Management API

Data Catalog Management API

API version: 0.0.1-alpha1
Contact: data-governance@confluent.io
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

// StreamCatalogV2NotificationTopicConfigSpec The desired state of the Notification Topic Config
type StreamCatalogV2NotificationTopicConfigSpec struct {
	// The kafka topic to write catalog notifications
	Topic *string `json:"topic,omitempty"`
	// Catalog notifications enablement
	Enabled *bool `json:"enabled,omitempty"`
	// The environment to which this belongs.
	Environment *EnvScopedObjectReference `json:"environment,omitempty"`
	// The kafka_cluster associated with this object.
	KafkaCluster *EnvScopedObjectReference `json:"kafka_cluster,omitempty"`
}

// NewStreamCatalogV2NotificationTopicConfigSpec instantiates a new StreamCatalogV2NotificationTopicConfigSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamCatalogV2NotificationTopicConfigSpec() *StreamCatalogV2NotificationTopicConfigSpec {
	this := StreamCatalogV2NotificationTopicConfigSpec{}
	return &this
}

// NewStreamCatalogV2NotificationTopicConfigSpecWithDefaults instantiates a new StreamCatalogV2NotificationTopicConfigSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamCatalogV2NotificationTopicConfigSpecWithDefaults() *StreamCatalogV2NotificationTopicConfigSpec {
	this := StreamCatalogV2NotificationTopicConfigSpec{}
	return &this
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *StreamCatalogV2NotificationTopicConfigSpec) GetTopic() string {
	if o == nil || o.Topic == nil {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamCatalogV2NotificationTopicConfigSpec) GetTopicOk() (*string, bool) {
	if o == nil || o.Topic == nil {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *StreamCatalogV2NotificationTopicConfigSpec) HasTopic() bool {
	if o != nil && o.Topic != nil {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *StreamCatalogV2NotificationTopicConfigSpec) SetTopic(v string) {
	o.Topic = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *StreamCatalogV2NotificationTopicConfigSpec) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamCatalogV2NotificationTopicConfigSpec) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *StreamCatalogV2NotificationTopicConfigSpec) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *StreamCatalogV2NotificationTopicConfigSpec) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *StreamCatalogV2NotificationTopicConfigSpec) GetEnvironment() EnvScopedObjectReference {
	if o == nil || o.Environment == nil {
		var ret EnvScopedObjectReference
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamCatalogV2NotificationTopicConfigSpec) GetEnvironmentOk() (*EnvScopedObjectReference, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *StreamCatalogV2NotificationTopicConfigSpec) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given EnvScopedObjectReference and assigns it to the Environment field.
func (o *StreamCatalogV2NotificationTopicConfigSpec) SetEnvironment(v EnvScopedObjectReference) {
	o.Environment = &v
}

// GetKafkaCluster returns the KafkaCluster field value if set, zero value otherwise.
func (o *StreamCatalogV2NotificationTopicConfigSpec) GetKafkaCluster() EnvScopedObjectReference {
	if o == nil || o.KafkaCluster == nil {
		var ret EnvScopedObjectReference
		return ret
	}
	return *o.KafkaCluster
}

// GetKafkaClusterOk returns a tuple with the KafkaCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamCatalogV2NotificationTopicConfigSpec) GetKafkaClusterOk() (*EnvScopedObjectReference, bool) {
	if o == nil || o.KafkaCluster == nil {
		return nil, false
	}
	return o.KafkaCluster, true
}

// HasKafkaCluster returns a boolean if a field has been set.
func (o *StreamCatalogV2NotificationTopicConfigSpec) HasKafkaCluster() bool {
	if o != nil && o.KafkaCluster != nil {
		return true
	}

	return false
}

// SetKafkaCluster gets a reference to the given EnvScopedObjectReference and assigns it to the KafkaCluster field.
func (o *StreamCatalogV2NotificationTopicConfigSpec) SetKafkaCluster(v EnvScopedObjectReference) {
	o.KafkaCluster = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *StreamCatalogV2NotificationTopicConfigSpec) Redact() {
	o.recurseRedact(o.Topic)
	o.recurseRedact(o.Enabled)
	o.recurseRedact(o.Environment)
	o.recurseRedact(o.KafkaCluster)
}

func (o *StreamCatalogV2NotificationTopicConfigSpec) recurseRedact(v interface{}) {
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

func (o StreamCatalogV2NotificationTopicConfigSpec) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o StreamCatalogV2NotificationTopicConfigSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Topic != nil {
		toSerialize["topic"] = o.Topic
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.KafkaCluster != nil {
		toSerialize["kafka_cluster"] = o.KafkaCluster
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableStreamCatalogV2NotificationTopicConfigSpec struct {
	value *StreamCatalogV2NotificationTopicConfigSpec
	isSet bool
}

func (v NullableStreamCatalogV2NotificationTopicConfigSpec) Get() *StreamCatalogV2NotificationTopicConfigSpec {
	return v.value
}

func (v *NullableStreamCatalogV2NotificationTopicConfigSpec) Set(val *StreamCatalogV2NotificationTopicConfigSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamCatalogV2NotificationTopicConfigSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamCatalogV2NotificationTopicConfigSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamCatalogV2NotificationTopicConfigSpec(val *StreamCatalogV2NotificationTopicConfigSpec) *NullableStreamCatalogV2NotificationTopicConfigSpec {
	return &NullableStreamCatalogV2NotificationTopicConfigSpec{value: val, isSet: true}
}

func (v NullableStreamCatalogV2NotificationTopicConfigSpec) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableStreamCatalogV2NotificationTopicConfigSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}