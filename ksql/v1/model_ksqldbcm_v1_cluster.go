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

package v1

import (
	"encoding/json"
)

import (
	"reflect"
)

// KsqldbcmV1Cluster `Cluster` represents a ksqlDB runtime that you can issue queries to using its API endpoint. It executes SQL statements and queries which under the hood get built into corresponding Kafka Streams topologies. The API allows you to create, read, and delete your ksqlDB clusters.   Related guide: [ksqlDB in Confluent Cloud](https://docs.confluent.io/cloud/current/ksqldb/index.html).  ## The Clusters Model <SchemaDefinition schemaRef=\"#/components/schemas/ksqldbcm.v1.Cluster\" />  ## Quotas and Limits This resource is subject to the following quotas:  | Quota | Description | | --- | --- | | `ksql.limits.max_apps_per_cluster` | Clusters in one Confluent Cloud Kafka Cluster. |
type KsqldbcmV1Cluster struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind *string `json:"kind,omitempty"`
	// ID is the \"natural identifier\" for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\"time\"); however, it may collide with IDs for other object `kinds` or objects of the same `kind` within a different scope/namespace (\"space\").
	Id *string `json:"id,omitempty"`
	Metadata *ObjectMeta `json:"metadata,omitempty"`
	Spec *KsqldbcmV1ClusterSpec `json:"spec,omitempty"`
	Status *KsqldbcmV1ClusterStatus `json:"status,omitempty"`
}

// NewKsqldbcmV1Cluster instantiates a new KsqldbcmV1Cluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKsqldbcmV1Cluster() *KsqldbcmV1Cluster {
	this := KsqldbcmV1Cluster{}
	return &this
}

// NewKsqldbcmV1ClusterWithDefaults instantiates a new KsqldbcmV1Cluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKsqldbcmV1ClusterWithDefaults() *KsqldbcmV1Cluster {
	this := KsqldbcmV1Cluster{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *KsqldbcmV1Cluster) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsqldbcmV1Cluster) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *KsqldbcmV1Cluster) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *KsqldbcmV1Cluster) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *KsqldbcmV1Cluster) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsqldbcmV1Cluster) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *KsqldbcmV1Cluster) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *KsqldbcmV1Cluster) SetKind(v string) {
	o.Kind = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KsqldbcmV1Cluster) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsqldbcmV1Cluster) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KsqldbcmV1Cluster) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KsqldbcmV1Cluster) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *KsqldbcmV1Cluster) GetMetadata() ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsqldbcmV1Cluster) GetMetadataOk() (*ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *KsqldbcmV1Cluster) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ObjectMeta and assigns it to the Metadata field.
func (o *KsqldbcmV1Cluster) SetMetadata(v ObjectMeta) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *KsqldbcmV1Cluster) GetSpec() KsqldbcmV1ClusterSpec {
	if o == nil || o.Spec == nil {
		var ret KsqldbcmV1ClusterSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsqldbcmV1Cluster) GetSpecOk() (*KsqldbcmV1ClusterSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *KsqldbcmV1Cluster) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given KsqldbcmV1ClusterSpec and assigns it to the Spec field.
func (o *KsqldbcmV1Cluster) SetSpec(v KsqldbcmV1ClusterSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KsqldbcmV1Cluster) GetStatus() KsqldbcmV1ClusterStatus {
	if o == nil || o.Status == nil {
		var ret KsqldbcmV1ClusterStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsqldbcmV1Cluster) GetStatusOk() (*KsqldbcmV1ClusterStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KsqldbcmV1Cluster) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given KsqldbcmV1ClusterStatus and assigns it to the Status field.
func (o *KsqldbcmV1Cluster) SetStatus(v KsqldbcmV1ClusterStatus) {
	o.Status = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *KsqldbcmV1Cluster) Redact() {
    o.recurseRedact(o.ApiVersion)
    o.recurseRedact(o.Kind)
    o.recurseRedact(o.Id)
    o.recurseRedact(o.Metadata)
    o.recurseRedact(o.Spec)
    o.recurseRedact(o.Status)
}

func (o *KsqldbcmV1Cluster) recurseRedact(v interface{}) {
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

func (o KsqldbcmV1Cluster) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o KsqldbcmV1Cluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api_version"] = o.ApiVersion
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableKsqldbcmV1Cluster struct {
	value *KsqldbcmV1Cluster
	isSet bool
}

func (v NullableKsqldbcmV1Cluster) Get() *KsqldbcmV1Cluster {
	return v.value
}

func (v *NullableKsqldbcmV1Cluster) Set(val *KsqldbcmV1Cluster) {
	v.value = val
	v.isSet = true
}

func (v NullableKsqldbcmV1Cluster) IsSet() bool {
	return v.isSet
}

func (v *NullableKsqldbcmV1Cluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKsqldbcmV1Cluster(val *KsqldbcmV1Cluster) *NullableKsqldbcmV1Cluster {
	return &NullableKsqldbcmV1Cluster{value: val, isSet: true}
}

func (v NullableKsqldbcmV1Cluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKsqldbcmV1Cluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

