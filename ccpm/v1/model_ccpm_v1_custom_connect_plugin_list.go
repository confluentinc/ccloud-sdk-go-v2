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
Custom Connect Plugin Management API

This is Custom Connect Plugin Management API.

API version: 0.1.0
Contact: connect-team@confluent.io
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

// CcpmV1CustomConnectPluginList CustomConnectPlugins objects represent Custom Connect artifacts containing connector, and SMT jars on Confluent  Cloud. The API allows you to list, create, read, update, and delete your Custom Connect Plugins. Related guide: [Custom Connect Plugin API](https://docs.confluent.io/cloud/current/connectors/connect-api-section.html).   ## The Custom Connect Plugins Model <SchemaDefinition schemaRef=\"#/components/schemas/ccpm.v1.CustomConnectPlugin\" />
type CcpmV1CustomConnectPluginList struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind     string   `json:"kind,omitempty"`
	Metadata ListMeta `json:"metadata,omitempty"`
	// A data property that contains an array of resource items. Each entry in the array is a separate resource.
	Data []CcpmV1CustomConnectPlugin `json:"data,omitempty"`
}

// NewCcpmV1CustomConnectPluginList instantiates a new CcpmV1CustomConnectPluginList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCcpmV1CustomConnectPluginList(apiVersion string, kind string, metadata ListMeta, data []CcpmV1CustomConnectPlugin) *CcpmV1CustomConnectPluginList {
	this := CcpmV1CustomConnectPluginList{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Data = data
	return &this
}

// NewCcpmV1CustomConnectPluginListWithDefaults instantiates a new CcpmV1CustomConnectPluginList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCcpmV1CustomConnectPluginListWithDefaults() *CcpmV1CustomConnectPluginList {
	this := CcpmV1CustomConnectPluginList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *CcpmV1CustomConnectPluginList) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *CcpmV1CustomConnectPluginList) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *CcpmV1CustomConnectPluginList) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *CcpmV1CustomConnectPluginList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *CcpmV1CustomConnectPluginList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *CcpmV1CustomConnectPluginList) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *CcpmV1CustomConnectPluginList) GetMetadata() ListMeta {
	if o == nil {
		var ret ListMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *CcpmV1CustomConnectPluginList) GetMetadataOk() (*ListMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *CcpmV1CustomConnectPluginList) SetMetadata(v ListMeta) {
	o.Metadata = v
}

// GetData returns the Data field value
func (o *CcpmV1CustomConnectPluginList) GetData() []CcpmV1CustomConnectPlugin {
	if o == nil {
		var ret []CcpmV1CustomConnectPlugin
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CcpmV1CustomConnectPluginList) GetDataOk() (*[]CcpmV1CustomConnectPlugin, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CcpmV1CustomConnectPluginList) SetData(v []CcpmV1CustomConnectPlugin) {
	o.Data = v
}

// Redact resets all sensitive fields to their zero value.
func (o *CcpmV1CustomConnectPluginList) Redact() {
	o.recurseRedact(&o.ApiVersion)
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Metadata)
	o.recurseRedact(&o.Data)
}

func (o *CcpmV1CustomConnectPluginList) recurseRedact(v interface{}) {
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

func (o CcpmV1CustomConnectPluginList) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o CcpmV1CustomConnectPluginList) MarshalJSON() ([]byte, error) {
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

type NullableCcpmV1CustomConnectPluginList struct {
	value *CcpmV1CustomConnectPluginList
	isSet bool
}

func (v NullableCcpmV1CustomConnectPluginList) Get() *CcpmV1CustomConnectPluginList {
	return v.value
}

func (v *NullableCcpmV1CustomConnectPluginList) Set(val *CcpmV1CustomConnectPluginList) {
	v.value = val
	v.isSet = true
}

func (v NullableCcpmV1CustomConnectPluginList) IsSet() bool {
	return v.isSet
}

func (v *NullableCcpmV1CustomConnectPluginList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCcpmV1CustomConnectPluginList(val *CcpmV1CustomConnectPluginList) *NullableCcpmV1CustomConnectPluginList {
	return &NullableCcpmV1CustomConnectPluginList{value: val, isSet: true}
}

func (v NullableCcpmV1CustomConnectPluginList) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableCcpmV1CustomConnectPluginList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
