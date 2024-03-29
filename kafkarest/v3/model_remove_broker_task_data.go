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

// RemoveBrokerTaskData struct for RemoveBrokerTaskData
type RemoveBrokerTaskData struct {
	Kind                         string           `json:"kind,omitempty"`
	Metadata                     ResourceMetadata `json:"metadata,omitempty"`
	ClusterId                    string           `json:"cluster_id,omitempty"`
	BrokerId                     int32            `json:"broker_id,omitempty"`
	ShutdownScheduled            bool             `json:"shutdown_scheduled,omitempty"`
	BrokerReplicaExclusionStatus string           `json:"broker_replica_exclusion_status,omitempty"`
	PartitionReassignmentStatus  string           `json:"partition_reassignment_status,omitempty"`
	BrokerShutdownStatus         string           `json:"broker_shutdown_status,omitempty"`
	ErrorCode                    NullableInt32    `json:"error_code,omitempty"`
	ErrorMessage                 NullableString   `json:"error_message,omitempty"`
	Broker                       Relationship     `json:"broker,omitempty"`
}

// NewRemoveBrokerTaskData instantiates a new RemoveBrokerTaskData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveBrokerTaskData(kind string, metadata ResourceMetadata, clusterId string, brokerId int32, shutdownScheduled bool, brokerReplicaExclusionStatus string, partitionReassignmentStatus string, brokerShutdownStatus string, broker Relationship) *RemoveBrokerTaskData {
	this := RemoveBrokerTaskData{}
	this.Kind = kind
	this.Metadata = metadata
	this.ClusterId = clusterId
	this.BrokerId = brokerId
	this.ShutdownScheduled = shutdownScheduled
	this.BrokerReplicaExclusionStatus = brokerReplicaExclusionStatus
	this.PartitionReassignmentStatus = partitionReassignmentStatus
	this.BrokerShutdownStatus = brokerShutdownStatus
	this.Broker = broker
	return &this
}

// NewRemoveBrokerTaskDataWithDefaults instantiates a new RemoveBrokerTaskData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveBrokerTaskDataWithDefaults() *RemoveBrokerTaskData {
	this := RemoveBrokerTaskData{}
	return &this
}

// GetKind returns the Kind field value
func (o *RemoveBrokerTaskData) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *RemoveBrokerTaskData) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *RemoveBrokerTaskData) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *RemoveBrokerTaskData) GetMetadata() ResourceMetadata {
	if o == nil {
		var ret ResourceMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *RemoveBrokerTaskData) GetMetadataOk() (*ResourceMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *RemoveBrokerTaskData) SetMetadata(v ResourceMetadata) {
	o.Metadata = v
}

// GetClusterId returns the ClusterId field value
func (o *RemoveBrokerTaskData) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *RemoveBrokerTaskData) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *RemoveBrokerTaskData) SetClusterId(v string) {
	o.ClusterId = v
}

// GetBrokerId returns the BrokerId field value
func (o *RemoveBrokerTaskData) GetBrokerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BrokerId
}

// GetBrokerIdOk returns a tuple with the BrokerId field value
// and a boolean to check if the value has been set.
func (o *RemoveBrokerTaskData) GetBrokerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrokerId, true
}

// SetBrokerId sets field value
func (o *RemoveBrokerTaskData) SetBrokerId(v int32) {
	o.BrokerId = v
}

// GetShutdownScheduled returns the ShutdownScheduled field value
func (o *RemoveBrokerTaskData) GetShutdownScheduled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShutdownScheduled
}

// GetShutdownScheduledOk returns a tuple with the ShutdownScheduled field value
// and a boolean to check if the value has been set.
func (o *RemoveBrokerTaskData) GetShutdownScheduledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShutdownScheduled, true
}

// SetShutdownScheduled sets field value
func (o *RemoveBrokerTaskData) SetShutdownScheduled(v bool) {
	o.ShutdownScheduled = v
}

// GetBrokerReplicaExclusionStatus returns the BrokerReplicaExclusionStatus field value
func (o *RemoveBrokerTaskData) GetBrokerReplicaExclusionStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrokerReplicaExclusionStatus
}

// GetBrokerReplicaExclusionStatusOk returns a tuple with the BrokerReplicaExclusionStatus field value
// and a boolean to check if the value has been set.
func (o *RemoveBrokerTaskData) GetBrokerReplicaExclusionStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrokerReplicaExclusionStatus, true
}

// SetBrokerReplicaExclusionStatus sets field value
func (o *RemoveBrokerTaskData) SetBrokerReplicaExclusionStatus(v string) {
	o.BrokerReplicaExclusionStatus = v
}

// GetPartitionReassignmentStatus returns the PartitionReassignmentStatus field value
func (o *RemoveBrokerTaskData) GetPartitionReassignmentStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PartitionReassignmentStatus
}

// GetPartitionReassignmentStatusOk returns a tuple with the PartitionReassignmentStatus field value
// and a boolean to check if the value has been set.
func (o *RemoveBrokerTaskData) GetPartitionReassignmentStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartitionReassignmentStatus, true
}

// SetPartitionReassignmentStatus sets field value
func (o *RemoveBrokerTaskData) SetPartitionReassignmentStatus(v string) {
	o.PartitionReassignmentStatus = v
}

// GetBrokerShutdownStatus returns the BrokerShutdownStatus field value
func (o *RemoveBrokerTaskData) GetBrokerShutdownStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrokerShutdownStatus
}

// GetBrokerShutdownStatusOk returns a tuple with the BrokerShutdownStatus field value
// and a boolean to check if the value has been set.
func (o *RemoveBrokerTaskData) GetBrokerShutdownStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrokerShutdownStatus, true
}

// SetBrokerShutdownStatus sets field value
func (o *RemoveBrokerTaskData) SetBrokerShutdownStatus(v string) {
	o.BrokerShutdownStatus = v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoveBrokerTaskData) GetErrorCode() int32 {
	if o == nil || o.ErrorCode.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCode.Get()
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoveBrokerTaskData) GetErrorCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorCode.Get(), o.ErrorCode.IsSet()
}

// HasErrorCode returns a boolean if a field has been set.
func (o *RemoveBrokerTaskData) HasErrorCode() bool {
	if o != nil && o.ErrorCode.IsSet() {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullableInt32 and assigns it to the ErrorCode field.
func (o *RemoveBrokerTaskData) SetErrorCode(v int32) {
	o.ErrorCode.Set(&v)
}

// SetErrorCodeNil sets the value for ErrorCode to be an explicit nil
func (o *RemoveBrokerTaskData) SetErrorCodeNil() {
	o.ErrorCode.Set(nil)
}

// UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
func (o *RemoveBrokerTaskData) UnsetErrorCode() {
	o.ErrorCode.Unset()
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoveBrokerTaskData) GetErrorMessage() string {
	if o == nil || o.ErrorMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoveBrokerTaskData) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *RemoveBrokerTaskData) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *RemoveBrokerTaskData) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}

// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *RemoveBrokerTaskData) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *RemoveBrokerTaskData) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetBroker returns the Broker field value
func (o *RemoveBrokerTaskData) GetBroker() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.Broker
}

// GetBrokerOk returns a tuple with the Broker field value
// and a boolean to check if the value has been set.
func (o *RemoveBrokerTaskData) GetBrokerOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Broker, true
}

// SetBroker sets field value
func (o *RemoveBrokerTaskData) SetBroker(v Relationship) {
	o.Broker = v
}

// Redact resets all sensitive fields to their zero value.
func (o *RemoveBrokerTaskData) Redact() {
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Metadata)
	o.recurseRedact(&o.ClusterId)
	o.recurseRedact(&o.BrokerId)
	o.recurseRedact(&o.ShutdownScheduled)
	o.recurseRedact(&o.BrokerReplicaExclusionStatus)
	o.recurseRedact(&o.PartitionReassignmentStatus)
	o.recurseRedact(&o.BrokerShutdownStatus)
	o.recurseRedact(o.ErrorCode)
	o.recurseRedact(o.ErrorMessage)
	o.recurseRedact(&o.Broker)
}

func (o *RemoveBrokerTaskData) recurseRedact(v interface{}) {
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

func (o RemoveBrokerTaskData) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o RemoveBrokerTaskData) MarshalJSON() ([]byte, error) {
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
		toSerialize["shutdown_scheduled"] = o.ShutdownScheduled
	}
	if true {
		toSerialize["broker_replica_exclusion_status"] = o.BrokerReplicaExclusionStatus
	}
	if true {
		toSerialize["partition_reassignment_status"] = o.PartitionReassignmentStatus
	}
	if true {
		toSerialize["broker_shutdown_status"] = o.BrokerShutdownStatus
	}
	if o.ErrorCode.IsSet() {
		toSerialize["error_code"] = o.ErrorCode.Get()
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["error_message"] = o.ErrorMessage.Get()
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

type NullableRemoveBrokerTaskData struct {
	value *RemoveBrokerTaskData
	isSet bool
}

func (v NullableRemoveBrokerTaskData) Get() *RemoveBrokerTaskData {
	return v.value
}

func (v *NullableRemoveBrokerTaskData) Set(val *RemoveBrokerTaskData) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveBrokerTaskData) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveBrokerTaskData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveBrokerTaskData(val *RemoveBrokerTaskData) *NullableRemoveBrokerTaskData {
	return &NullableRemoveBrokerTaskData{value: val, isSet: true}
}

func (v NullableRemoveBrokerTaskData) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableRemoveBrokerTaskData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
