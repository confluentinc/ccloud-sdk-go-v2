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
Flink Cluster Management API

This is the new Flink cluster management API.

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

// FlinkcmV2ClusterList `Cluster` represents a Flink runtime that you can issue queries to using its API endpoint. The API allows you to list, create, read, and delete your Flink clusters.   Related guide: [Flink in Confluent Cloud](https://docs.confluent.io/cloud/current/ksqldb/ksqldb-cluster-api.html).  ## The Clusters Model <SchemaDefinition schemaRef=\"#/components/schemas/flinkcm.v2.Cluster\" />  ## Quotas and Limits This resource is subject to the following quotas:  | Quota | Description | | --- | --- | | `ksql.limits.max_apps_per_cluster` | Clusters in one Confluent Cloud Kafka Cluster. |
type FlinkcmV2ClusterList struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion string `json:"api_version"`
	// Kind defines the object this REST resource represents.
	Kind string `json:"kind"`
	Metadata ListMeta `json:"metadata"`
	// A data property that contains an array of resource items. Each entry in the array is a separate resource.
	Data []FlinkcmV2Cluster `json:"data"`
}

// NewFlinkcmV2ClusterList instantiates a new FlinkcmV2ClusterList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlinkcmV2ClusterList(apiVersion string, kind string, metadata ListMeta, data []FlinkcmV2Cluster) *FlinkcmV2ClusterList {
	this := FlinkcmV2ClusterList{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Data = data
	return &this
}

// NewFlinkcmV2ClusterListWithDefaults instantiates a new FlinkcmV2ClusterList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlinkcmV2ClusterListWithDefaults() *FlinkcmV2ClusterList {
	this := FlinkcmV2ClusterList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *FlinkcmV2ClusterList) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *FlinkcmV2ClusterList) GetApiVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *FlinkcmV2ClusterList) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *FlinkcmV2ClusterList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *FlinkcmV2ClusterList) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *FlinkcmV2ClusterList) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *FlinkcmV2ClusterList) GetMetadata() ListMeta {
	if o == nil {
		var ret ListMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *FlinkcmV2ClusterList) GetMetadataOk() (*ListMeta, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *FlinkcmV2ClusterList) SetMetadata(v ListMeta) {
	o.Metadata = v
}

// GetData returns the Data field value
func (o *FlinkcmV2ClusterList) GetData() []FlinkcmV2Cluster {
	if o == nil {
		var ret []FlinkcmV2Cluster
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FlinkcmV2ClusterList) GetDataOk() (*[]FlinkcmV2Cluster, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *FlinkcmV2ClusterList) SetData(v []FlinkcmV2Cluster) {
	o.Data = v
}

// Redact resets all sensitive fields to their zero value.
func (o *FlinkcmV2ClusterList) Redact() {
    o.recurseRedact(&o.ApiVersion)
    o.recurseRedact(&o.Kind)
    o.recurseRedact(&o.Metadata)
    o.recurseRedact(&o.Data)
}

func (o *FlinkcmV2ClusterList) recurseRedact(v interface{}) {
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

func (o FlinkcmV2ClusterList) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o FlinkcmV2ClusterList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["api_version"] = o.ApiVersion
	}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableFlinkcmV2ClusterList struct {
	value *FlinkcmV2ClusterList
	isSet bool
}

func (v NullableFlinkcmV2ClusterList) Get() *FlinkcmV2ClusterList {
	return v.value
}

func (v *NullableFlinkcmV2ClusterList) Set(val *FlinkcmV2ClusterList) {
	v.value = val
	v.isSet = true
}

func (v NullableFlinkcmV2ClusterList) IsSet() bool {
	return v.isSet
}

func (v *NullableFlinkcmV2ClusterList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlinkcmV2ClusterList(val *FlinkcmV2ClusterList) *NullableFlinkcmV2ClusterList {
	return &NullableFlinkcmV2ClusterList{value: val, isSet: true}
}

func (v NullableFlinkcmV2ClusterList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlinkcmV2ClusterList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

