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
Flink Compute Pool Management API

This is the Flink Compute Pool management API.

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

// FiamV2FlinkIdentityPoolEnvRegionBindingList A FlinkIdentityPoolEnvRegionBinding represents the binding of identity pools to all Flink resource pools in an environment and region.  This is needed for authentication purposes.   ## The Flink Identity Pool Env Region Bindings Model <SchemaDefinition schemaRef=\"#/components/schemas/fiam.v2.FlinkIdentityPoolEnvRegionBinding\" />
type FiamV2FlinkIdentityPoolEnvRegionBindingList struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion string `json:"api_version"`
	// Kind defines the object this REST resource represents.
	Kind string `json:"kind"`
	Metadata ListMeta `json:"metadata"`
	// A data property that contains an array of resource items. Each entry in the array is a separate resource.
	Data []FiamV2FlinkIdentityPoolEnvRegionBinding `json:"data"`
}

// NewFiamV2FlinkIdentityPoolEnvRegionBindingList instantiates a new FiamV2FlinkIdentityPoolEnvRegionBindingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiamV2FlinkIdentityPoolEnvRegionBindingList(apiVersion string, kind string, metadata ListMeta, data []FiamV2FlinkIdentityPoolEnvRegionBinding) *FiamV2FlinkIdentityPoolEnvRegionBindingList {
	this := FiamV2FlinkIdentityPoolEnvRegionBindingList{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Data = data
	return &this
}

// NewFiamV2FlinkIdentityPoolEnvRegionBindingListWithDefaults instantiates a new FiamV2FlinkIdentityPoolEnvRegionBindingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiamV2FlinkIdentityPoolEnvRegionBindingListWithDefaults() *FiamV2FlinkIdentityPoolEnvRegionBindingList {
	this := FiamV2FlinkIdentityPoolEnvRegionBindingList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) GetApiVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) GetMetadata() ListMeta {
	if o == nil {
		var ret ListMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) GetMetadataOk() (*ListMeta, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) SetMetadata(v ListMeta) {
	o.Metadata = v
}

// GetData returns the Data field value
func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) GetData() []FiamV2FlinkIdentityPoolEnvRegionBinding {
	if o == nil {
		var ret []FiamV2FlinkIdentityPoolEnvRegionBinding
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) GetDataOk() (*[]FiamV2FlinkIdentityPoolEnvRegionBinding, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) SetData(v []FiamV2FlinkIdentityPoolEnvRegionBinding) {
	o.Data = v
}

// Redact resets all sensitive fields to their zero value.
func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) Redact() {
    o.recurseRedact(&o.ApiVersion)
    o.recurseRedact(&o.Kind)
    o.recurseRedact(&o.Metadata)
    o.recurseRedact(&o.Data)
}

func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) recurseRedact(v interface{}) {
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

func (o FiamV2FlinkIdentityPoolEnvRegionBindingList) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o FiamV2FlinkIdentityPoolEnvRegionBindingList) MarshalJSON() ([]byte, error) {
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

type NullableFiamV2FlinkIdentityPoolEnvRegionBindingList struct {
	value *FiamV2FlinkIdentityPoolEnvRegionBindingList
	isSet bool
}

func (v NullableFiamV2FlinkIdentityPoolEnvRegionBindingList) Get() *FiamV2FlinkIdentityPoolEnvRegionBindingList {
	return v.value
}

func (v *NullableFiamV2FlinkIdentityPoolEnvRegionBindingList) Set(val *FiamV2FlinkIdentityPoolEnvRegionBindingList) {
	v.value = val
	v.isSet = true
}

func (v NullableFiamV2FlinkIdentityPoolEnvRegionBindingList) IsSet() bool {
	return v.isSet
}

func (v *NullableFiamV2FlinkIdentityPoolEnvRegionBindingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiamV2FlinkIdentityPoolEnvRegionBindingList(val *FiamV2FlinkIdentityPoolEnvRegionBindingList) *NullableFiamV2FlinkIdentityPoolEnvRegionBindingList {
	return &NullableFiamV2FlinkIdentityPoolEnvRegionBindingList{value: val, isSet: true}
}

func (v NullableFiamV2FlinkIdentityPoolEnvRegionBindingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiamV2FlinkIdentityPoolEnvRegionBindingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

