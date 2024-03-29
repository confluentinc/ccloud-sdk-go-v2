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

// TopicData struct for TopicData
type TopicData struct {
	Kind                   string                `json:"kind,omitempty"`
	Metadata               ResourceMetadata      `json:"metadata,omitempty"`
	ClusterId              string                `json:"cluster_id,omitempty"`
	TopicName              string                `json:"topic_name,omitempty"`
	IsInternal             bool                  `json:"is_internal,omitempty"`
	ReplicationFactor      int32                 `json:"replication_factor,omitempty"`
	PartitionsCount        int32                 `json:"partitions_count,omitempty"`
	Partitions             Relationship          `json:"partitions,omitempty"`
	Configs                Relationship          `json:"configs,omitempty"`
	PartitionReassignments Relationship          `json:"partition_reassignments,omitempty"`
	AuthorizedOperations   *AuthorizedOperations `json:"authorized_operations,omitempty"`
}

// NewTopicData instantiates a new TopicData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopicData(kind string, metadata ResourceMetadata, clusterId string, topicName string, isInternal bool, replicationFactor int32, partitionsCount int32, partitions Relationship, configs Relationship, partitionReassignments Relationship) *TopicData {
	this := TopicData{}
	this.Kind = kind
	this.Metadata = metadata
	this.ClusterId = clusterId
	this.TopicName = topicName
	this.IsInternal = isInternal
	this.ReplicationFactor = replicationFactor
	this.PartitionsCount = partitionsCount
	this.Partitions = partitions
	this.Configs = configs
	this.PartitionReassignments = partitionReassignments
	return &this
}

// NewTopicDataWithDefaults instantiates a new TopicData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopicDataWithDefaults() *TopicData {
	this := TopicData{}
	return &this
}

// GetKind returns the Kind field value
func (o *TopicData) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *TopicData) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *TopicData) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *TopicData) GetMetadata() ResourceMetadata {
	if o == nil {
		var ret ResourceMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *TopicData) GetMetadataOk() (*ResourceMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *TopicData) SetMetadata(v ResourceMetadata) {
	o.Metadata = v
}

// GetClusterId returns the ClusterId field value
func (o *TopicData) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *TopicData) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *TopicData) SetClusterId(v string) {
	o.ClusterId = v
}

// GetTopicName returns the TopicName field value
func (o *TopicData) GetTopicName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TopicName
}

// GetTopicNameOk returns a tuple with the TopicName field value
// and a boolean to check if the value has been set.
func (o *TopicData) GetTopicNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopicName, true
}

// SetTopicName sets field value
func (o *TopicData) SetTopicName(v string) {
	o.TopicName = v
}

// GetIsInternal returns the IsInternal field value
func (o *TopicData) GetIsInternal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsInternal
}

// GetIsInternalOk returns a tuple with the IsInternal field value
// and a boolean to check if the value has been set.
func (o *TopicData) GetIsInternalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsInternal, true
}

// SetIsInternal sets field value
func (o *TopicData) SetIsInternal(v bool) {
	o.IsInternal = v
}

// GetReplicationFactor returns the ReplicationFactor field value
func (o *TopicData) GetReplicationFactor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReplicationFactor
}

// GetReplicationFactorOk returns a tuple with the ReplicationFactor field value
// and a boolean to check if the value has been set.
func (o *TopicData) GetReplicationFactorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReplicationFactor, true
}

// SetReplicationFactor sets field value
func (o *TopicData) SetReplicationFactor(v int32) {
	o.ReplicationFactor = v
}

// GetPartitionsCount returns the PartitionsCount field value
func (o *TopicData) GetPartitionsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PartitionsCount
}

// GetPartitionsCountOk returns a tuple with the PartitionsCount field value
// and a boolean to check if the value has been set.
func (o *TopicData) GetPartitionsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartitionsCount, true
}

// SetPartitionsCount sets field value
func (o *TopicData) SetPartitionsCount(v int32) {
	o.PartitionsCount = v
}

// GetPartitions returns the Partitions field value
func (o *TopicData) GetPartitions() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value
// and a boolean to check if the value has been set.
func (o *TopicData) GetPartitionsOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Partitions, true
}

// SetPartitions sets field value
func (o *TopicData) SetPartitions(v Relationship) {
	o.Partitions = v
}

// GetConfigs returns the Configs field value
func (o *TopicData) GetConfigs() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.Configs
}

// GetConfigsOk returns a tuple with the Configs field value
// and a boolean to check if the value has been set.
func (o *TopicData) GetConfigsOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Configs, true
}

// SetConfigs sets field value
func (o *TopicData) SetConfigs(v Relationship) {
	o.Configs = v
}

// GetPartitionReassignments returns the PartitionReassignments field value
func (o *TopicData) GetPartitionReassignments() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.PartitionReassignments
}

// GetPartitionReassignmentsOk returns a tuple with the PartitionReassignments field value
// and a boolean to check if the value has been set.
func (o *TopicData) GetPartitionReassignmentsOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartitionReassignments, true
}

// SetPartitionReassignments sets field value
func (o *TopicData) SetPartitionReassignments(v Relationship) {
	o.PartitionReassignments = v
}

// GetAuthorizedOperations returns the AuthorizedOperations field value if set, zero value otherwise.
func (o *TopicData) GetAuthorizedOperations() AuthorizedOperations {
	if o == nil || o.AuthorizedOperations == nil {
		var ret AuthorizedOperations
		return ret
	}
	return *o.AuthorizedOperations
}

// GetAuthorizedOperationsOk returns a tuple with the AuthorizedOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicData) GetAuthorizedOperationsOk() (*AuthorizedOperations, bool) {
	if o == nil || o.AuthorizedOperations == nil {
		return nil, false
	}
	return o.AuthorizedOperations, true
}

// HasAuthorizedOperations returns a boolean if a field has been set.
func (o *TopicData) HasAuthorizedOperations() bool {
	if o != nil && o.AuthorizedOperations != nil {
		return true
	}

	return false
}

// SetAuthorizedOperations gets a reference to the given AuthorizedOperations and assigns it to the AuthorizedOperations field.
func (o *TopicData) SetAuthorizedOperations(v AuthorizedOperations) {
	o.AuthorizedOperations = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *TopicData) Redact() {
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Metadata)
	o.recurseRedact(&o.ClusterId)
	o.recurseRedact(&o.TopicName)
	o.recurseRedact(&o.IsInternal)
	o.recurseRedact(&o.ReplicationFactor)
	o.recurseRedact(&o.PartitionsCount)
	o.recurseRedact(&o.Partitions)
	o.recurseRedact(&o.Configs)
	o.recurseRedact(&o.PartitionReassignments)
	o.recurseRedact(o.AuthorizedOperations)
}

func (o *TopicData) recurseRedact(v interface{}) {
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

func (o TopicData) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o TopicData) MarshalJSON() ([]byte, error) {
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
		toSerialize["topic_name"] = o.TopicName
	}
	if true {
		toSerialize["is_internal"] = o.IsInternal
	}
	if true {
		toSerialize["replication_factor"] = o.ReplicationFactor
	}
	if true {
		toSerialize["partitions_count"] = o.PartitionsCount
	}
	if true {
		toSerialize["partitions"] = o.Partitions
	}
	if true {
		toSerialize["configs"] = o.Configs
	}
	if true {
		toSerialize["partition_reassignments"] = o.PartitionReassignments
	}
	if o.AuthorizedOperations != nil {
		toSerialize["authorized_operations"] = o.AuthorizedOperations
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableTopicData struct {
	value *TopicData
	isSet bool
}

func (v NullableTopicData) Get() *TopicData {
	return v.value
}

func (v *NullableTopicData) Set(val *TopicData) {
	v.value = val
	v.isSet = true
}

func (v NullableTopicData) IsSet() bool {
	return v.isSet
}

func (v *NullableTopicData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopicData(val *TopicData) *NullableTopicData {
	return &NullableTopicData{value: val, isSet: true}
}

func (v NullableTopicData) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableTopicData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
