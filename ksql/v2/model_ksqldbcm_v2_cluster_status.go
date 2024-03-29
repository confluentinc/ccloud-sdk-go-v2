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
ksqlDB Cluster Management API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
Contact: ksql-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

import (
	"reflect"
)

// KsqldbcmV2ClusterStatus The status of the Cluster
type KsqldbcmV2ClusterStatus struct {
	// The dataplane endpoint of the ksqlDB cluster.
	HttpEndpoint *string `json:"http_endpoint,omitempty"`
	// Status of the ksqlDB cluster.
	Phase string `json:"phase"`
	// Tells you if the cluster has been paused
	IsPaused bool `json:"is_paused"`
	// Amount of storage (in GB) provisioned to this cluster
	Storage int32 `json:"storage"`
	// Topic name prefix used by this ksqlDB cluster. Used to assign ACLs for this ksqlDB cluster to use.
	TopicPrefix *string `json:"topic_prefix,omitempty"`
}

// NewKsqldbcmV2ClusterStatus instantiates a new KsqldbcmV2ClusterStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKsqldbcmV2ClusterStatus(phase string, isPaused bool, storage int32) *KsqldbcmV2ClusterStatus {
	this := KsqldbcmV2ClusterStatus{}
	this.Phase = phase
	this.IsPaused = isPaused
	this.Storage = storage
	return &this
}

// NewKsqldbcmV2ClusterStatusWithDefaults instantiates a new KsqldbcmV2ClusterStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKsqldbcmV2ClusterStatusWithDefaults() *KsqldbcmV2ClusterStatus {
	this := KsqldbcmV2ClusterStatus{}
	return &this
}

// GetHttpEndpoint returns the HttpEndpoint field value if set, zero value otherwise.
func (o *KsqldbcmV2ClusterStatus) GetHttpEndpoint() string {
	if o == nil || o.HttpEndpoint == nil {
		var ret string
		return ret
	}
	return *o.HttpEndpoint
}

// GetHttpEndpointOk returns a tuple with the HttpEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsqldbcmV2ClusterStatus) GetHttpEndpointOk() (*string, bool) {
	if o == nil || o.HttpEndpoint == nil {
		return nil, false
	}
	return o.HttpEndpoint, true
}

// HasHttpEndpoint returns a boolean if a field has been set.
func (o *KsqldbcmV2ClusterStatus) HasHttpEndpoint() bool {
	if o != nil && o.HttpEndpoint != nil {
		return true
	}

	return false
}

// SetHttpEndpoint gets a reference to the given string and assigns it to the HttpEndpoint field.
func (o *KsqldbcmV2ClusterStatus) SetHttpEndpoint(v string) {
	o.HttpEndpoint = &v
}

// GetPhase returns the Phase field value
func (o *KsqldbcmV2ClusterStatus) GetPhase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value
// and a boolean to check if the value has been set.
func (o *KsqldbcmV2ClusterStatus) GetPhaseOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Phase, true
}

// SetPhase sets field value
func (o *KsqldbcmV2ClusterStatus) SetPhase(v string) {
	o.Phase = v
}

// GetIsPaused returns the IsPaused field value
func (o *KsqldbcmV2ClusterStatus) GetIsPaused() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPaused
}

// GetIsPausedOk returns a tuple with the IsPaused field value
// and a boolean to check if the value has been set.
func (o *KsqldbcmV2ClusterStatus) GetIsPausedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsPaused, true
}

// SetIsPaused sets field value
func (o *KsqldbcmV2ClusterStatus) SetIsPaused(v bool) {
	o.IsPaused = v
}

// GetStorage returns the Storage field value
func (o *KsqldbcmV2ClusterStatus) GetStorage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *KsqldbcmV2ClusterStatus) GetStorageOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *KsqldbcmV2ClusterStatus) SetStorage(v int32) {
	o.Storage = v
}

// GetTopicPrefix returns the TopicPrefix field value if set, zero value otherwise.
func (o *KsqldbcmV2ClusterStatus) GetTopicPrefix() string {
	if o == nil || o.TopicPrefix == nil {
		var ret string
		return ret
	}
	return *o.TopicPrefix
}

// GetTopicPrefixOk returns a tuple with the TopicPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsqldbcmV2ClusterStatus) GetTopicPrefixOk() (*string, bool) {
	if o == nil || o.TopicPrefix == nil {
		return nil, false
	}
	return o.TopicPrefix, true
}

// HasTopicPrefix returns a boolean if a field has been set.
func (o *KsqldbcmV2ClusterStatus) HasTopicPrefix() bool {
	if o != nil && o.TopicPrefix != nil {
		return true
	}

	return false
}

// SetTopicPrefix gets a reference to the given string and assigns it to the TopicPrefix field.
func (o *KsqldbcmV2ClusterStatus) SetTopicPrefix(v string) {
	o.TopicPrefix = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *KsqldbcmV2ClusterStatus) Redact() {
    o.recurseRedact(o.HttpEndpoint)
    o.recurseRedact(&o.Phase)
    o.recurseRedact(&o.IsPaused)
    o.recurseRedact(&o.Storage)
    o.recurseRedact(o.TopicPrefix)
}

func (o *KsqldbcmV2ClusterStatus) recurseRedact(v interface{}) {
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

func (o KsqldbcmV2ClusterStatus) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o KsqldbcmV2ClusterStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HttpEndpoint != nil {
		toSerialize["http_endpoint"] = o.HttpEndpoint
	}
	if true {
		toSerialize["phase"] = o.Phase
	}
	if true {
		toSerialize["is_paused"] = o.IsPaused
	}
	if true {
		toSerialize["storage"] = o.Storage
	}
	if o.TopicPrefix != nil {
		toSerialize["topic_prefix"] = o.TopicPrefix
	}
	return json.Marshal(toSerialize)
}

type NullableKsqldbcmV2ClusterStatus struct {
	value *KsqldbcmV2ClusterStatus
	isSet bool
}

func (v NullableKsqldbcmV2ClusterStatus) Get() *KsqldbcmV2ClusterStatus {
	return v.value
}

func (v *NullableKsqldbcmV2ClusterStatus) Set(val *KsqldbcmV2ClusterStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableKsqldbcmV2ClusterStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableKsqldbcmV2ClusterStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKsqldbcmV2ClusterStatus(val *KsqldbcmV2ClusterStatus) *NullableKsqldbcmV2ClusterStatus {
	return &NullableKsqldbcmV2ClusterStatus{value: val, isSet: true}
}

func (v NullableKsqldbcmV2ClusterStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKsqldbcmV2ClusterStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


