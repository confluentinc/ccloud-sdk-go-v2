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
Stream Designer API

# Introduction  Stream Designer API provides resources/API for defining stream processing pipelines. Each pipeline describes a set of stream processing components, including connectors, topics, streams, tables, queries and schemas. The components in a pipeline need not exist as Confluent Cloud resources until the pipeline is activated.  This API defines operations to create, list, modify, manage and delete pipelines. 

API version: 0.0.1-alpha0
Contact: stream-designer@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

import (
	"reflect"
)

// SdV1PipelineList `Pipeline` objects represent information about a user-defined pipeline of Confluent Cloud components. The pipeline's content is available separately.  The API allows you to create, retrieve, update, and delete your pipelines, as well as list all of your pipelines for the particular environment and Kafka cluster.   Related guide: [Pipelines in Confluent Cloud](https://docs.confluent.io/cloud/current/stream-designer/).  ## The Pipelines Model <SchemaDefinition schemaRef=\"#/components/schemas/sd.v1.Pipeline\" />  ## Quotas and Limits This resource is subject to the following quotas:  | Quota | Description | | --- | --- | | `pipelines_per_org` | Pipelines in one Confluent Cloud organization | | `pipelines_per_cluster` | Pipelines in one Confluent Cloud Kafka cluster |
type SdV1PipelineList struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion string `json:"api_version"`
	// Kind defines the object this REST resource represents.
	Kind string `json:"kind"`
	Metadata ListMeta `json:"metadata"`
	// A data property that contains an array of resource items. Each entry in the array is a separate resource.
	Data []SdV1Pipeline `json:"data"`
}

// NewSdV1PipelineList instantiates a new SdV1PipelineList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdV1PipelineList(apiVersion string, kind string, metadata ListMeta, data []SdV1Pipeline) *SdV1PipelineList {
	this := SdV1PipelineList{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Data = data
	return &this
}

// NewSdV1PipelineListWithDefaults instantiates a new SdV1PipelineList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdV1PipelineListWithDefaults() *SdV1PipelineList {
	this := SdV1PipelineList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *SdV1PipelineList) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *SdV1PipelineList) GetApiVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *SdV1PipelineList) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *SdV1PipelineList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *SdV1PipelineList) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *SdV1PipelineList) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *SdV1PipelineList) GetMetadata() ListMeta {
	if o == nil {
		var ret ListMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *SdV1PipelineList) GetMetadataOk() (*ListMeta, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *SdV1PipelineList) SetMetadata(v ListMeta) {
	o.Metadata = v
}

// GetData returns the Data field value
func (o *SdV1PipelineList) GetData() []SdV1Pipeline {
	if o == nil {
		var ret []SdV1Pipeline
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SdV1PipelineList) GetDataOk() (*[]SdV1Pipeline, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SdV1PipelineList) SetData(v []SdV1Pipeline) {
	o.Data = v
}

// Redact resets all sensitive fields to their zero value.
func (o *SdV1PipelineList) Redact() {
    o.recurseRedact(&o.ApiVersion)
    o.recurseRedact(&o.Kind)
    o.recurseRedact(&o.Metadata)
    o.recurseRedact(&o.Data)
}

func (o *SdV1PipelineList) recurseRedact(v interface{}) {
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

func (o SdV1PipelineList) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o SdV1PipelineList) MarshalJSON() ([]byte, error) {
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

type NullableSdV1PipelineList struct {
	value *SdV1PipelineList
	isSet bool
}

func (v NullableSdV1PipelineList) Get() *SdV1PipelineList {
	return v.value
}

func (v *NullableSdV1PipelineList) Set(val *SdV1PipelineList) {
	v.value = val
	v.isSet = true
}

func (v NullableSdV1PipelineList) IsSet() bool {
	return v.isSet
}

func (v *NullableSdV1PipelineList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdV1PipelineList(val *SdV1PipelineList) *NullableSdV1PipelineList {
	return &NullableSdV1PipelineList{value: val, isSet: true}
}

func (v NullableSdV1PipelineList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdV1PipelineList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


