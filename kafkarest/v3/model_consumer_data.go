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
REST Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0.0
Contact: kafka-clients-proxy-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// ConsumerData struct for ConsumerData
type ConsumerData struct {
	Kind            string           `json:"kind,omitempty"`
	Metadata        ResourceMetadata `json:"metadata,omitempty"`
	ClusterId       string           `json:"cluster_id,omitempty"`
	ConsumerGroupId string           `json:"consumer_group_id,omitempty"`
	ConsumerId      string           `json:"consumer_id,omitempty"`
	InstanceId      NullableString   `json:"instance_id,omitempty"`
	ClientId        string           `json:"client_id,omitempty"`
	Assignments     Relationship     `json:"assignments,omitempty"`
}

// NewConsumerData instantiates a new ConsumerData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumerData(kind string, metadata ResourceMetadata, clusterId string, consumerGroupId string, consumerId string, clientId string, assignments Relationship) *ConsumerData {
	this := ConsumerData{}
	this.Kind = kind
	this.Metadata = metadata
	this.ClusterId = clusterId
	this.ConsumerGroupId = consumerGroupId
	this.ConsumerId = consumerId
	this.ClientId = clientId
	this.Assignments = assignments
	return &this
}

// NewConsumerDataWithDefaults instantiates a new ConsumerData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumerDataWithDefaults() *ConsumerData {
	this := ConsumerData{}
	return &this
}

// GetKind returns the Kind field value
func (o *ConsumerData) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ConsumerData) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ConsumerData) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *ConsumerData) GetMetadata() ResourceMetadata {
	if o == nil {
		var ret ResourceMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ConsumerData) GetMetadataOk() (*ResourceMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ConsumerData) SetMetadata(v ResourceMetadata) {
	o.Metadata = v
}

// GetClusterId returns the ClusterId field value
func (o *ConsumerData) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *ConsumerData) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *ConsumerData) SetClusterId(v string) {
	o.ClusterId = v
}

// GetConsumerGroupId returns the ConsumerGroupId field value
func (o *ConsumerData) GetConsumerGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsumerGroupId
}

// GetConsumerGroupIdOk returns a tuple with the ConsumerGroupId field value
// and a boolean to check if the value has been set.
func (o *ConsumerData) GetConsumerGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumerGroupId, true
}

// SetConsumerGroupId sets field value
func (o *ConsumerData) SetConsumerGroupId(v string) {
	o.ConsumerGroupId = v
}

// GetConsumerId returns the ConsumerId field value
func (o *ConsumerData) GetConsumerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsumerId
}

// GetConsumerIdOk returns a tuple with the ConsumerId field value
// and a boolean to check if the value has been set.
func (o *ConsumerData) GetConsumerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumerId, true
}

// SetConsumerId sets field value
func (o *ConsumerData) SetConsumerId(v string) {
	o.ConsumerId = v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsumerData) GetInstanceId() string {
	if o == nil || o.InstanceId.Get() == nil {
		var ret string
		return ret
	}
	return *o.InstanceId.Get()
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsumerData) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstanceId.Get(), o.InstanceId.IsSet()
}

// HasInstanceId returns a boolean if a field has been set.
func (o *ConsumerData) HasInstanceId() bool {
	if o != nil && o.InstanceId.IsSet() {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given NullableString and assigns it to the InstanceId field.
func (o *ConsumerData) SetInstanceId(v string) {
	o.InstanceId.Set(&v)
}

// SetInstanceIdNil sets the value for InstanceId to be an explicit nil
func (o *ConsumerData) SetInstanceIdNil() {
	o.InstanceId.Set(nil)
}

// UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
func (o *ConsumerData) UnsetInstanceId() {
	o.InstanceId.Unset()
}

// GetClientId returns the ClientId field value
func (o *ConsumerData) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *ConsumerData) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *ConsumerData) SetClientId(v string) {
	o.ClientId = v
}

// GetAssignments returns the Assignments field value
func (o *ConsumerData) GetAssignments() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value
// and a boolean to check if the value has been set.
func (o *ConsumerData) GetAssignmentsOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Assignments, true
}

// SetAssignments sets field value
func (o *ConsumerData) SetAssignments(v Relationship) {
	o.Assignments = v
}

// Redact resets all sensitive fields to their zero value.
func (o *ConsumerData) Redact() {
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Metadata)
	o.recurseRedact(&o.ClusterId)
	o.recurseRedact(&o.ConsumerGroupId)
	o.recurseRedact(&o.ConsumerId)
	o.recurseRedact(o.InstanceId)
	o.recurseRedact(&o.ClientId)
	o.recurseRedact(&o.Assignments)
}

func (o *ConsumerData) recurseRedact(v interface{}) {
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

func (o ConsumerData) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ConsumerData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if true {
		toSerialize["consumer_group_id"] = o.ConsumerGroupId
	}
	if true {
		toSerialize["consumer_id"] = o.ConsumerId
	}
	if o.InstanceId.IsSet() {
		toSerialize["instance_id"] = o.InstanceId.Get()
	}
	if true {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["assignments"] = o.Assignments
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableConsumerData struct {
	value *ConsumerData
	isSet bool
}

func (v NullableConsumerData) Get() *ConsumerData {
	return v.value
}

func (v *NullableConsumerData) Set(val *ConsumerData) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumerData) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumerData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumerData(val *ConsumerData) *NullableConsumerData {
	return &NullableConsumerData{value: val, isSet: true}
}

func (v NullableConsumerData) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableConsumerData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
