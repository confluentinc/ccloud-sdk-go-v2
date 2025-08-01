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
Custom Connector Plugin Management API

This is Custom Connector Plugin Management API.

API version: 1.0.0
Contact: compute-platform-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// ConnectV1CustomConnectorRuntimeList List of supported runtime languages for Custom Connector Plugin. The list defines the supported entries for confluent.custom.connect.plugin.runtime attribute in CustomConnectorPlugin object. Each entry also defines the set of supported java versions for that runtime which can be specified during connector provisioning via the confluent.custom.connect.plugin.java.version attribute.   ## The Custom Connector Runtimes Model <SchemaDefinition schemaRef=\"#/components/schemas/connect.v1.CustomConnectorRuntime\" />
type ConnectV1CustomConnectorRuntimeList struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind     string   `json:"kind,omitempty"`
	Metadata ListMeta `json:"metadata,omitempty"`
	// A data property that contains an array of resource items. Each entry in the array is a separate resource.
	Data []ConnectV1CustomConnectorRuntime `json:"data,omitempty"`
}

// NewConnectV1CustomConnectorRuntimeList instantiates a new ConnectV1CustomConnectorRuntimeList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectV1CustomConnectorRuntimeList(apiVersion string, kind string, metadata ListMeta, data []ConnectV1CustomConnectorRuntime) *ConnectV1CustomConnectorRuntimeList {
	this := ConnectV1CustomConnectorRuntimeList{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Data = data
	return &this
}

// NewConnectV1CustomConnectorRuntimeListWithDefaults instantiates a new ConnectV1CustomConnectorRuntimeList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectV1CustomConnectorRuntimeListWithDefaults() *ConnectV1CustomConnectorRuntimeList {
	this := ConnectV1CustomConnectorRuntimeList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *ConnectV1CustomConnectorRuntimeList) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorRuntimeList) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *ConnectV1CustomConnectorRuntimeList) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *ConnectV1CustomConnectorRuntimeList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorRuntimeList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ConnectV1CustomConnectorRuntimeList) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *ConnectV1CustomConnectorRuntimeList) GetMetadata() ListMeta {
	if o == nil {
		var ret ListMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorRuntimeList) GetMetadataOk() (*ListMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ConnectV1CustomConnectorRuntimeList) SetMetadata(v ListMeta) {
	o.Metadata = v
}

// GetData returns the Data field value
func (o *ConnectV1CustomConnectorRuntimeList) GetData() []ConnectV1CustomConnectorRuntime {
	if o == nil {
		var ret []ConnectV1CustomConnectorRuntime
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorRuntimeList) GetDataOk() (*[]ConnectV1CustomConnectorRuntime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ConnectV1CustomConnectorRuntimeList) SetData(v []ConnectV1CustomConnectorRuntime) {
	o.Data = v
}

// Redact resets all sensitive fields to their zero value.
func (o *ConnectV1CustomConnectorRuntimeList) Redact() {
	o.recurseRedact(&o.ApiVersion)
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Metadata)
	o.recurseRedact(&o.Data)
}

func (o *ConnectV1CustomConnectorRuntimeList) recurseRedact(v interface{}) {
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

func (o ConnectV1CustomConnectorRuntimeList) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ConnectV1CustomConnectorRuntimeList) MarshalJSON() ([]byte, error) {
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
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableConnectV1CustomConnectorRuntimeList struct {
	value *ConnectV1CustomConnectorRuntimeList
	isSet bool
}

func (v NullableConnectV1CustomConnectorRuntimeList) Get() *ConnectV1CustomConnectorRuntimeList {
	return v.value
}

func (v *NullableConnectV1CustomConnectorRuntimeList) Set(val *ConnectV1CustomConnectorRuntimeList) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectV1CustomConnectorRuntimeList) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectV1CustomConnectorRuntimeList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectV1CustomConnectorRuntimeList(val *ConnectV1CustomConnectorRuntimeList) *NullableConnectV1CustomConnectorRuntimeList {
	return &NullableConnectV1CustomConnectorRuntimeList{value: val, isSet: true}
}

func (v NullableConnectV1CustomConnectorRuntimeList) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableConnectV1CustomConnectorRuntimeList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
