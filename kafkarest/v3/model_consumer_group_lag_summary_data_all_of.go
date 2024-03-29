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

// ConsumerGroupLagSummaryDataAllOf struct for ConsumerGroupLagSummaryDataAllOf
type ConsumerGroupLagSummaryDataAllOf struct {
	ClusterId         string         `json:"cluster_id,omitempty"`
	ConsumerGroupId   string         `json:"consumer_group_id,omitempty"`
	MaxLagConsumerId  string         `json:"max_lag_consumer_id,omitempty"`
	MaxLagInstanceId  NullableString `json:"max_lag_instance_id,omitempty"`
	MaxLagClientId    string         `json:"max_lag_client_id,omitempty"`
	MaxLagTopicName   string         `json:"max_lag_topic_name,omitempty"`
	MaxLagPartitionId int32          `json:"max_lag_partition_id,omitempty"`
	MaxLag            int64          `json:"max_lag,omitempty"`
	TotalLag          int64          `json:"total_lag,omitempty"`
	MaxLagConsumer    Relationship   `json:"max_lag_consumer,omitempty"`
	MaxLagPartition   Relationship   `json:"max_lag_partition,omitempty"`
}

// NewConsumerGroupLagSummaryDataAllOf instantiates a new ConsumerGroupLagSummaryDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumerGroupLagSummaryDataAllOf(clusterId string, consumerGroupId string, maxLagConsumerId string, maxLagClientId string, maxLagTopicName string, maxLagPartitionId int32, maxLag int64, totalLag int64, maxLagConsumer Relationship, maxLagPartition Relationship) *ConsumerGroupLagSummaryDataAllOf {
	this := ConsumerGroupLagSummaryDataAllOf{}
	this.ClusterId = clusterId
	this.ConsumerGroupId = consumerGroupId
	this.MaxLagConsumerId = maxLagConsumerId
	this.MaxLagClientId = maxLagClientId
	this.MaxLagTopicName = maxLagTopicName
	this.MaxLagPartitionId = maxLagPartitionId
	this.MaxLag = maxLag
	this.TotalLag = totalLag
	this.MaxLagConsumer = maxLagConsumer
	this.MaxLagPartition = maxLagPartition
	return &this
}

// NewConsumerGroupLagSummaryDataAllOfWithDefaults instantiates a new ConsumerGroupLagSummaryDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumerGroupLagSummaryDataAllOfWithDefaults() *ConsumerGroupLagSummaryDataAllOf {
	this := ConsumerGroupLagSummaryDataAllOf{}
	return &this
}

// GetClusterId returns the ClusterId field value
func (o *ConsumerGroupLagSummaryDataAllOf) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupLagSummaryDataAllOf) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *ConsumerGroupLagSummaryDataAllOf) SetClusterId(v string) {
	o.ClusterId = v
}

// GetConsumerGroupId returns the ConsumerGroupId field value
func (o *ConsumerGroupLagSummaryDataAllOf) GetConsumerGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsumerGroupId
}

// GetConsumerGroupIdOk returns a tuple with the ConsumerGroupId field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupLagSummaryDataAllOf) GetConsumerGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumerGroupId, true
}

// SetConsumerGroupId sets field value
func (o *ConsumerGroupLagSummaryDataAllOf) SetConsumerGroupId(v string) {
	o.ConsumerGroupId = v
}

// GetMaxLagConsumerId returns the MaxLagConsumerId field value
func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagConsumerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxLagConsumerId
}

// GetMaxLagConsumerIdOk returns a tuple with the MaxLagConsumerId field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagConsumerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxLagConsumerId, true
}

// SetMaxLagConsumerId sets field value
func (o *ConsumerGroupLagSummaryDataAllOf) SetMaxLagConsumerId(v string) {
	o.MaxLagConsumerId = v
}

// GetMaxLagInstanceId returns the MaxLagInstanceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagInstanceId() string {
	if o == nil || o.MaxLagInstanceId.Get() == nil {
		var ret string
		return ret
	}
	return *o.MaxLagInstanceId.Get()
}

// GetMaxLagInstanceIdOk returns a tuple with the MaxLagInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxLagInstanceId.Get(), o.MaxLagInstanceId.IsSet()
}

// HasMaxLagInstanceId returns a boolean if a field has been set.
func (o *ConsumerGroupLagSummaryDataAllOf) HasMaxLagInstanceId() bool {
	if o != nil && o.MaxLagInstanceId.IsSet() {
		return true
	}

	return false
}

// SetMaxLagInstanceId gets a reference to the given NullableString and assigns it to the MaxLagInstanceId field.
func (o *ConsumerGroupLagSummaryDataAllOf) SetMaxLagInstanceId(v string) {
	o.MaxLagInstanceId.Set(&v)
}

// SetMaxLagInstanceIdNil sets the value for MaxLagInstanceId to be an explicit nil
func (o *ConsumerGroupLagSummaryDataAllOf) SetMaxLagInstanceIdNil() {
	o.MaxLagInstanceId.Set(nil)
}

// UnsetMaxLagInstanceId ensures that no value is present for MaxLagInstanceId, not even an explicit nil
func (o *ConsumerGroupLagSummaryDataAllOf) UnsetMaxLagInstanceId() {
	o.MaxLagInstanceId.Unset()
}

// GetMaxLagClientId returns the MaxLagClientId field value
func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxLagClientId
}

// GetMaxLagClientIdOk returns a tuple with the MaxLagClientId field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxLagClientId, true
}

// SetMaxLagClientId sets field value
func (o *ConsumerGroupLagSummaryDataAllOf) SetMaxLagClientId(v string) {
	o.MaxLagClientId = v
}

// GetMaxLagTopicName returns the MaxLagTopicName field value
func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagTopicName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxLagTopicName
}

// GetMaxLagTopicNameOk returns a tuple with the MaxLagTopicName field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagTopicNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxLagTopicName, true
}

// SetMaxLagTopicName sets field value
func (o *ConsumerGroupLagSummaryDataAllOf) SetMaxLagTopicName(v string) {
	o.MaxLagTopicName = v
}

// GetMaxLagPartitionId returns the MaxLagPartitionId field value
func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagPartitionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxLagPartitionId
}

// GetMaxLagPartitionIdOk returns a tuple with the MaxLagPartitionId field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagPartitionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxLagPartitionId, true
}

// SetMaxLagPartitionId sets field value
func (o *ConsumerGroupLagSummaryDataAllOf) SetMaxLagPartitionId(v int32) {
	o.MaxLagPartitionId = v
}

// GetMaxLag returns the MaxLag field value
func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLag() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxLag
}

// GetMaxLagOk returns a tuple with the MaxLag field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxLag, true
}

// SetMaxLag sets field value
func (o *ConsumerGroupLagSummaryDataAllOf) SetMaxLag(v int64) {
	o.MaxLag = v
}

// GetTotalLag returns the TotalLag field value
func (o *ConsumerGroupLagSummaryDataAllOf) GetTotalLag() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalLag
}

// GetTotalLagOk returns a tuple with the TotalLag field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupLagSummaryDataAllOf) GetTotalLagOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalLag, true
}

// SetTotalLag sets field value
func (o *ConsumerGroupLagSummaryDataAllOf) SetTotalLag(v int64) {
	o.TotalLag = v
}

// GetMaxLagConsumer returns the MaxLagConsumer field value
func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagConsumer() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.MaxLagConsumer
}

// GetMaxLagConsumerOk returns a tuple with the MaxLagConsumer field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagConsumerOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxLagConsumer, true
}

// SetMaxLagConsumer sets field value
func (o *ConsumerGroupLagSummaryDataAllOf) SetMaxLagConsumer(v Relationship) {
	o.MaxLagConsumer = v
}

// GetMaxLagPartition returns the MaxLagPartition field value
func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagPartition() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.MaxLagPartition
}

// GetMaxLagPartitionOk returns a tuple with the MaxLagPartition field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagPartitionOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxLagPartition, true
}

// SetMaxLagPartition sets field value
func (o *ConsumerGroupLagSummaryDataAllOf) SetMaxLagPartition(v Relationship) {
	o.MaxLagPartition = v
}

// Redact resets all sensitive fields to their zero value.
func (o *ConsumerGroupLagSummaryDataAllOf) Redact() {
	o.recurseRedact(&o.ClusterId)
	o.recurseRedact(&o.ConsumerGroupId)
	o.recurseRedact(&o.MaxLagConsumerId)
	o.recurseRedact(o.MaxLagInstanceId)
	o.recurseRedact(&o.MaxLagClientId)
	o.recurseRedact(&o.MaxLagTopicName)
	o.recurseRedact(&o.MaxLagPartitionId)
	o.recurseRedact(&o.MaxLag)
	o.recurseRedact(&o.TotalLag)
	o.recurseRedact(&o.MaxLagConsumer)
	o.recurseRedact(&o.MaxLagPartition)
}

func (o *ConsumerGroupLagSummaryDataAllOf) recurseRedact(v interface{}) {
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

func (o ConsumerGroupLagSummaryDataAllOf) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ConsumerGroupLagSummaryDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if true {
		toSerialize["consumer_group_id"] = o.ConsumerGroupId
	}
	if true {
		toSerialize["max_lag_consumer_id"] = o.MaxLagConsumerId
	}
	if o.MaxLagInstanceId.IsSet() {
		toSerialize["max_lag_instance_id"] = o.MaxLagInstanceId.Get()
	}
	if true {
		toSerialize["max_lag_client_id"] = o.MaxLagClientId
	}
	if true {
		toSerialize["max_lag_topic_name"] = o.MaxLagTopicName
	}
	if true {
		toSerialize["max_lag_partition_id"] = o.MaxLagPartitionId
	}
	if true {
		toSerialize["max_lag"] = o.MaxLag
	}
	if true {
		toSerialize["total_lag"] = o.TotalLag
	}
	if true {
		toSerialize["max_lag_consumer"] = o.MaxLagConsumer
	}
	if true {
		toSerialize["max_lag_partition"] = o.MaxLagPartition
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableConsumerGroupLagSummaryDataAllOf struct {
	value *ConsumerGroupLagSummaryDataAllOf
	isSet bool
}

func (v NullableConsumerGroupLagSummaryDataAllOf) Get() *ConsumerGroupLagSummaryDataAllOf {
	return v.value
}

func (v *NullableConsumerGroupLagSummaryDataAllOf) Set(val *ConsumerGroupLagSummaryDataAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumerGroupLagSummaryDataAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumerGroupLagSummaryDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumerGroupLagSummaryDataAllOf(val *ConsumerGroupLagSummaryDataAllOf) *NullableConsumerGroupLagSummaryDataAllOf {
	return &NullableConsumerGroupLagSummaryDataAllOf{value: val, isSet: true}
}

func (v NullableConsumerGroupLagSummaryDataAllOf) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableConsumerGroupLagSummaryDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
