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

// ConsumerLagDataAllOf struct for ConsumerLagDataAllOf
type ConsumerLagDataAllOf struct {
	ClusterId       string         `json:"cluster_id,omitempty"`
	ConsumerGroupId string         `json:"consumer_group_id,omitempty"`
	TopicName       string         `json:"topic_name,omitempty"`
	PartitionId     int32          `json:"partition_id,omitempty"`
	CurrentOffset   int64          `json:"current_offset,omitempty"`
	LogEndOffset    int64          `json:"log_end_offset,omitempty"`
	Lag             int64          `json:"lag,omitempty"`
	ConsumerId      string         `json:"consumer_id,omitempty"`
	InstanceId      NullableString `json:"instance_id,omitempty"`
	ClientId        string         `json:"client_id,omitempty"`
}

// NewConsumerLagDataAllOf instantiates a new ConsumerLagDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumerLagDataAllOf(clusterId string, consumerGroupId string, topicName string, partitionId int32, currentOffset int64, logEndOffset int64, lag int64, consumerId string, clientId string) *ConsumerLagDataAllOf {
	this := ConsumerLagDataAllOf{}
	this.ClusterId = clusterId
	this.ConsumerGroupId = consumerGroupId
	this.TopicName = topicName
	this.PartitionId = partitionId
	this.CurrentOffset = currentOffset
	this.LogEndOffset = logEndOffset
	this.Lag = lag
	this.ConsumerId = consumerId
	this.ClientId = clientId
	return &this
}

// NewConsumerLagDataAllOfWithDefaults instantiates a new ConsumerLagDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumerLagDataAllOfWithDefaults() *ConsumerLagDataAllOf {
	this := ConsumerLagDataAllOf{}
	return &this
}

// GetClusterId returns the ClusterId field value
func (o *ConsumerLagDataAllOf) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *ConsumerLagDataAllOf) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *ConsumerLagDataAllOf) SetClusterId(v string) {
	o.ClusterId = v
}

// GetConsumerGroupId returns the ConsumerGroupId field value
func (o *ConsumerLagDataAllOf) GetConsumerGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsumerGroupId
}

// GetConsumerGroupIdOk returns a tuple with the ConsumerGroupId field value
// and a boolean to check if the value has been set.
func (o *ConsumerLagDataAllOf) GetConsumerGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumerGroupId, true
}

// SetConsumerGroupId sets field value
func (o *ConsumerLagDataAllOf) SetConsumerGroupId(v string) {
	o.ConsumerGroupId = v
}

// GetTopicName returns the TopicName field value
func (o *ConsumerLagDataAllOf) GetTopicName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TopicName
}

// GetTopicNameOk returns a tuple with the TopicName field value
// and a boolean to check if the value has been set.
func (o *ConsumerLagDataAllOf) GetTopicNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopicName, true
}

// SetTopicName sets field value
func (o *ConsumerLagDataAllOf) SetTopicName(v string) {
	o.TopicName = v
}

// GetPartitionId returns the PartitionId field value
func (o *ConsumerLagDataAllOf) GetPartitionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PartitionId
}

// GetPartitionIdOk returns a tuple with the PartitionId field value
// and a boolean to check if the value has been set.
func (o *ConsumerLagDataAllOf) GetPartitionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartitionId, true
}

// SetPartitionId sets field value
func (o *ConsumerLagDataAllOf) SetPartitionId(v int32) {
	o.PartitionId = v
}

// GetCurrentOffset returns the CurrentOffset field value
func (o *ConsumerLagDataAllOf) GetCurrentOffset() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CurrentOffset
}

// GetCurrentOffsetOk returns a tuple with the CurrentOffset field value
// and a boolean to check if the value has been set.
func (o *ConsumerLagDataAllOf) GetCurrentOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentOffset, true
}

// SetCurrentOffset sets field value
func (o *ConsumerLagDataAllOf) SetCurrentOffset(v int64) {
	o.CurrentOffset = v
}

// GetLogEndOffset returns the LogEndOffset field value
func (o *ConsumerLagDataAllOf) GetLogEndOffset() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LogEndOffset
}

// GetLogEndOffsetOk returns a tuple with the LogEndOffset field value
// and a boolean to check if the value has been set.
func (o *ConsumerLagDataAllOf) GetLogEndOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogEndOffset, true
}

// SetLogEndOffset sets field value
func (o *ConsumerLagDataAllOf) SetLogEndOffset(v int64) {
	o.LogEndOffset = v
}

// GetLag returns the Lag field value
func (o *ConsumerLagDataAllOf) GetLag() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Lag
}

// GetLagOk returns a tuple with the Lag field value
// and a boolean to check if the value has been set.
func (o *ConsumerLagDataAllOf) GetLagOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lag, true
}

// SetLag sets field value
func (o *ConsumerLagDataAllOf) SetLag(v int64) {
	o.Lag = v
}

// GetConsumerId returns the ConsumerId field value
func (o *ConsumerLagDataAllOf) GetConsumerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsumerId
}

// GetConsumerIdOk returns a tuple with the ConsumerId field value
// and a boolean to check if the value has been set.
func (o *ConsumerLagDataAllOf) GetConsumerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumerId, true
}

// SetConsumerId sets field value
func (o *ConsumerLagDataAllOf) SetConsumerId(v string) {
	o.ConsumerId = v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsumerLagDataAllOf) GetInstanceId() string {
	if o == nil || o.InstanceId.Get() == nil {
		var ret string
		return ret
	}
	return *o.InstanceId.Get()
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsumerLagDataAllOf) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstanceId.Get(), o.InstanceId.IsSet()
}

// HasInstanceId returns a boolean if a field has been set.
func (o *ConsumerLagDataAllOf) HasInstanceId() bool {
	if o != nil && o.InstanceId.IsSet() {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given NullableString and assigns it to the InstanceId field.
func (o *ConsumerLagDataAllOf) SetInstanceId(v string) {
	o.InstanceId.Set(&v)
}

// SetInstanceIdNil sets the value for InstanceId to be an explicit nil
func (o *ConsumerLagDataAllOf) SetInstanceIdNil() {
	o.InstanceId.Set(nil)
}

// UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
func (o *ConsumerLagDataAllOf) UnsetInstanceId() {
	o.InstanceId.Unset()
}

// GetClientId returns the ClientId field value
func (o *ConsumerLagDataAllOf) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *ConsumerLagDataAllOf) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *ConsumerLagDataAllOf) SetClientId(v string) {
	o.ClientId = v
}

// Redact resets all sensitive fields to their zero value.
func (o *ConsumerLagDataAllOf) Redact() {
	o.recurseRedact(&o.ClusterId)
	o.recurseRedact(&o.ConsumerGroupId)
	o.recurseRedact(&o.TopicName)
	o.recurseRedact(&o.PartitionId)
	o.recurseRedact(&o.CurrentOffset)
	o.recurseRedact(&o.LogEndOffset)
	o.recurseRedact(&o.Lag)
	o.recurseRedact(&o.ConsumerId)
	o.recurseRedact(o.InstanceId)
	o.recurseRedact(&o.ClientId)
}

func (o *ConsumerLagDataAllOf) recurseRedact(v interface{}) {
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

func (o ConsumerLagDataAllOf) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ConsumerLagDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if true {
		toSerialize["consumer_group_id"] = o.ConsumerGroupId
	}
	if true {
		toSerialize["topic_name"] = o.TopicName
	}
	if true {
		toSerialize["partition_id"] = o.PartitionId
	}
	if true {
		toSerialize["current_offset"] = o.CurrentOffset
	}
	if true {
		toSerialize["log_end_offset"] = o.LogEndOffset
	}
	if true {
		toSerialize["lag"] = o.Lag
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
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableConsumerLagDataAllOf struct {
	value *ConsumerLagDataAllOf
	isSet bool
}

func (v NullableConsumerLagDataAllOf) Get() *ConsumerLagDataAllOf {
	return v.value
}

func (v *NullableConsumerLagDataAllOf) Set(val *ConsumerLagDataAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumerLagDataAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumerLagDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumerLagDataAllOf(val *ConsumerLagDataAllOf) *NullableConsumerLagDataAllOf {
	return &NullableConsumerLagDataAllOf{value: val, isSet: true}
}

func (v NullableConsumerLagDataAllOf) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableConsumerLagDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
