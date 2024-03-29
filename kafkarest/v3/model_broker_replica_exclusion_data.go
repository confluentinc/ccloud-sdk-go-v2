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

// BrokerReplicaExclusionData struct for BrokerReplicaExclusionData
type BrokerReplicaExclusionData struct {
	Kind      string           `json:"kind,omitempty"`
	Metadata  ResourceMetadata `json:"metadata,omitempty"`
	ClusterId string           `json:"cluster_id,omitempty"`
	BrokerId  int32            `json:"broker_id,omitempty"`
	Reason    string           `json:"reason,omitempty"`
	Broker    Relationship     `json:"broker,omitempty"`
}

// NewBrokerReplicaExclusionData instantiates a new BrokerReplicaExclusionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrokerReplicaExclusionData(kind string, metadata ResourceMetadata, clusterId string, brokerId int32, reason string, broker Relationship) *BrokerReplicaExclusionData {
	this := BrokerReplicaExclusionData{}
	this.Kind = kind
	this.Metadata = metadata
	this.ClusterId = clusterId
	this.BrokerId = brokerId
	this.Reason = reason
	this.Broker = broker
	return &this
}

// NewBrokerReplicaExclusionDataWithDefaults instantiates a new BrokerReplicaExclusionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrokerReplicaExclusionDataWithDefaults() *BrokerReplicaExclusionData {
	this := BrokerReplicaExclusionData{}
	return &this
}

// GetKind returns the Kind field value
func (o *BrokerReplicaExclusionData) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *BrokerReplicaExclusionData) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *BrokerReplicaExclusionData) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *BrokerReplicaExclusionData) GetMetadata() ResourceMetadata {
	if o == nil {
		var ret ResourceMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *BrokerReplicaExclusionData) GetMetadataOk() (*ResourceMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *BrokerReplicaExclusionData) SetMetadata(v ResourceMetadata) {
	o.Metadata = v
}

// GetClusterId returns the ClusterId field value
func (o *BrokerReplicaExclusionData) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *BrokerReplicaExclusionData) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *BrokerReplicaExclusionData) SetClusterId(v string) {
	o.ClusterId = v
}

// GetBrokerId returns the BrokerId field value
func (o *BrokerReplicaExclusionData) GetBrokerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BrokerId
}

// GetBrokerIdOk returns a tuple with the BrokerId field value
// and a boolean to check if the value has been set.
func (o *BrokerReplicaExclusionData) GetBrokerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrokerId, true
}

// SetBrokerId sets field value
func (o *BrokerReplicaExclusionData) SetBrokerId(v int32) {
	o.BrokerId = v
}

// GetReason returns the Reason field value
func (o *BrokerReplicaExclusionData) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *BrokerReplicaExclusionData) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *BrokerReplicaExclusionData) SetReason(v string) {
	o.Reason = v
}

// GetBroker returns the Broker field value
func (o *BrokerReplicaExclusionData) GetBroker() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.Broker
}

// GetBrokerOk returns a tuple with the Broker field value
// and a boolean to check if the value has been set.
func (o *BrokerReplicaExclusionData) GetBrokerOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Broker, true
}

// SetBroker sets field value
func (o *BrokerReplicaExclusionData) SetBroker(v Relationship) {
	o.Broker = v
}

// Redact resets all sensitive fields to their zero value.
func (o *BrokerReplicaExclusionData) Redact() {
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Metadata)
	o.recurseRedact(&o.ClusterId)
	o.recurseRedact(&o.BrokerId)
	o.recurseRedact(&o.Reason)
	o.recurseRedact(&o.Broker)
}

func (o *BrokerReplicaExclusionData) recurseRedact(v interface{}) {
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

func (o BrokerReplicaExclusionData) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o BrokerReplicaExclusionData) MarshalJSON() ([]byte, error) {
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
		toSerialize["broker_id"] = o.BrokerId
	}
	if true {
		toSerialize["reason"] = o.Reason
	}
	if true {
		toSerialize["broker"] = o.Broker
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableBrokerReplicaExclusionData struct {
	value *BrokerReplicaExclusionData
	isSet bool
}

func (v NullableBrokerReplicaExclusionData) Get() *BrokerReplicaExclusionData {
	return v.value
}

func (v *NullableBrokerReplicaExclusionData) Set(val *BrokerReplicaExclusionData) {
	v.value = val
	v.isSet = true
}

func (v NullableBrokerReplicaExclusionData) IsSet() bool {
	return v.isSet
}

func (v *NullableBrokerReplicaExclusionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrokerReplicaExclusionData(val *BrokerReplicaExclusionData) *NullableBrokerReplicaExclusionData {
	return &NullableBrokerReplicaExclusionData{value: val, isSet: true}
}

func (v NullableBrokerReplicaExclusionData) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableBrokerReplicaExclusionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
