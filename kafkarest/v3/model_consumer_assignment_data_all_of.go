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

// ConsumerAssignmentDataAllOf struct for ConsumerAssignmentDataAllOf
type ConsumerAssignmentDataAllOf struct {
	ClusterId       string       `json:"cluster_id,omitempty"`
	ConsumerGroupId string       `json:"consumer_group_id,omitempty"`
	ConsumerId      string       `json:"consumer_id,omitempty"`
	TopicName       string       `json:"topic_name,omitempty"`
	PartitionId     int32        `json:"partition_id,omitempty"`
	Partition       Relationship `json:"partition,omitempty"`
	Lag             Relationship `json:"lag,omitempty"`
}

// NewConsumerAssignmentDataAllOf instantiates a new ConsumerAssignmentDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumerAssignmentDataAllOf(clusterId string, consumerGroupId string, consumerId string, topicName string, partitionId int32, partition Relationship, lag Relationship) *ConsumerAssignmentDataAllOf {
	this := ConsumerAssignmentDataAllOf{}
	this.ClusterId = clusterId
	this.ConsumerGroupId = consumerGroupId
	this.ConsumerId = consumerId
	this.TopicName = topicName
	this.PartitionId = partitionId
	this.Partition = partition
	this.Lag = lag
	return &this
}

// NewConsumerAssignmentDataAllOfWithDefaults instantiates a new ConsumerAssignmentDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumerAssignmentDataAllOfWithDefaults() *ConsumerAssignmentDataAllOf {
	this := ConsumerAssignmentDataAllOf{}
	return &this
}

// GetClusterId returns the ClusterId field value
func (o *ConsumerAssignmentDataAllOf) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *ConsumerAssignmentDataAllOf) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *ConsumerAssignmentDataAllOf) SetClusterId(v string) {
	o.ClusterId = v
}

// GetConsumerGroupId returns the ConsumerGroupId field value
func (o *ConsumerAssignmentDataAllOf) GetConsumerGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsumerGroupId
}

// GetConsumerGroupIdOk returns a tuple with the ConsumerGroupId field value
// and a boolean to check if the value has been set.
func (o *ConsumerAssignmentDataAllOf) GetConsumerGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumerGroupId, true
}

// SetConsumerGroupId sets field value
func (o *ConsumerAssignmentDataAllOf) SetConsumerGroupId(v string) {
	o.ConsumerGroupId = v
}

// GetConsumerId returns the ConsumerId field value
func (o *ConsumerAssignmentDataAllOf) GetConsumerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsumerId
}

// GetConsumerIdOk returns a tuple with the ConsumerId field value
// and a boolean to check if the value has been set.
func (o *ConsumerAssignmentDataAllOf) GetConsumerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumerId, true
}

// SetConsumerId sets field value
func (o *ConsumerAssignmentDataAllOf) SetConsumerId(v string) {
	o.ConsumerId = v
}

// GetTopicName returns the TopicName field value
func (o *ConsumerAssignmentDataAllOf) GetTopicName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TopicName
}

// GetTopicNameOk returns a tuple with the TopicName field value
// and a boolean to check if the value has been set.
func (o *ConsumerAssignmentDataAllOf) GetTopicNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopicName, true
}

// SetTopicName sets field value
func (o *ConsumerAssignmentDataAllOf) SetTopicName(v string) {
	o.TopicName = v
}

// GetPartitionId returns the PartitionId field value
func (o *ConsumerAssignmentDataAllOf) GetPartitionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PartitionId
}

// GetPartitionIdOk returns a tuple with the PartitionId field value
// and a boolean to check if the value has been set.
func (o *ConsumerAssignmentDataAllOf) GetPartitionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartitionId, true
}

// SetPartitionId sets field value
func (o *ConsumerAssignmentDataAllOf) SetPartitionId(v int32) {
	o.PartitionId = v
}

// GetPartition returns the Partition field value
func (o *ConsumerAssignmentDataAllOf) GetPartition() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value
// and a boolean to check if the value has been set.
func (o *ConsumerAssignmentDataAllOf) GetPartitionOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Partition, true
}

// SetPartition sets field value
func (o *ConsumerAssignmentDataAllOf) SetPartition(v Relationship) {
	o.Partition = v
}

// GetLag returns the Lag field value
func (o *ConsumerAssignmentDataAllOf) GetLag() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.Lag
}

// GetLagOk returns a tuple with the Lag field value
// and a boolean to check if the value has been set.
func (o *ConsumerAssignmentDataAllOf) GetLagOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lag, true
}

// SetLag sets field value
func (o *ConsumerAssignmentDataAllOf) SetLag(v Relationship) {
	o.Lag = v
}

// Redact resets all sensitive fields to their zero value.
func (o *ConsumerAssignmentDataAllOf) Redact() {
	o.recurseRedact(&o.ClusterId)
	o.recurseRedact(&o.ConsumerGroupId)
	o.recurseRedact(&o.ConsumerId)
	o.recurseRedact(&o.TopicName)
	o.recurseRedact(&o.PartitionId)
	o.recurseRedact(&o.Partition)
	o.recurseRedact(&o.Lag)
}

func (o *ConsumerAssignmentDataAllOf) recurseRedact(v interface{}) {
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

func (o ConsumerAssignmentDataAllOf) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ConsumerAssignmentDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if true {
		toSerialize["consumer_group_id"] = o.ConsumerGroupId
	}
	if true {
		toSerialize["consumer_id"] = o.ConsumerId
	}
	if true {
		toSerialize["topic_name"] = o.TopicName
	}
	if true {
		toSerialize["partition_id"] = o.PartitionId
	}
	if true {
		toSerialize["partition"] = o.Partition
	}
	if true {
		toSerialize["lag"] = o.Lag
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableConsumerAssignmentDataAllOf struct {
	value *ConsumerAssignmentDataAllOf
	isSet bool
}

func (v NullableConsumerAssignmentDataAllOf) Get() *ConsumerAssignmentDataAllOf {
	return v.value
}

func (v *NullableConsumerAssignmentDataAllOf) Set(val *ConsumerAssignmentDataAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumerAssignmentDataAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumerAssignmentDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumerAssignmentDataAllOf(val *ConsumerAssignmentDataAllOf) *NullableConsumerAssignmentDataAllOf {
	return &NullableConsumerAssignmentDataAllOf{value: val, isSet: true}
}

func (v NullableConsumerAssignmentDataAllOf) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableConsumerAssignmentDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
