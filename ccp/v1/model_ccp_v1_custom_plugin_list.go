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
Custom Plugin Management API

This is Custom Plugin Management API.

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

// CcpV1CustomPluginList `CustomPlugins` objects represent Custom Plugins on Confluent Cloud.  The API allows you to list, create, read, update, and delete your Custom Plugins.   ## The Custom Plugins Model <SchemaDefinition schemaRef=\"#/components/schemas/ccp.v1.CustomPlugin\" />
type CcpV1CustomPluginList struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind     string   `json:"kind,omitempty"`
	Metadata ListMeta `json:"metadata,omitempty"`
	// A data property that contains an array of resource items. Each entry in the array is a separate resource.
	Data []CcpV1CustomPlugin `json:"data,omitempty"`
}

// NewCcpV1CustomPluginList instantiates a new CcpV1CustomPluginList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCcpV1CustomPluginList(apiVersion string, kind string, metadata ListMeta, data []CcpV1CustomPlugin) *CcpV1CustomPluginList {
	this := CcpV1CustomPluginList{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Data = data
	return &this
}

// NewCcpV1CustomPluginListWithDefaults instantiates a new CcpV1CustomPluginList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCcpV1CustomPluginListWithDefaults() *CcpV1CustomPluginList {
	this := CcpV1CustomPluginList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *CcpV1CustomPluginList) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *CcpV1CustomPluginList) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *CcpV1CustomPluginList) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *CcpV1CustomPluginList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *CcpV1CustomPluginList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *CcpV1CustomPluginList) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *CcpV1CustomPluginList) GetMetadata() ListMeta {
	if o == nil {
		var ret ListMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *CcpV1CustomPluginList) GetMetadataOk() (*ListMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *CcpV1CustomPluginList) SetMetadata(v ListMeta) {
	o.Metadata = v
}

// GetData returns the Data field value
func (o *CcpV1CustomPluginList) GetData() []CcpV1CustomPlugin {
	if o == nil {
		var ret []CcpV1CustomPlugin
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CcpV1CustomPluginList) GetDataOk() (*[]CcpV1CustomPlugin, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CcpV1CustomPluginList) SetData(v []CcpV1CustomPlugin) {
	o.Data = v
}

// Redact resets all sensitive fields to their zero value.
func (o *CcpV1CustomPluginList) Redact() {
	o.recurseRedact(&o.ApiVersion)
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Metadata)
	o.recurseRedact(&o.Data)
}

func (o *CcpV1CustomPluginList) recurseRedact(v interface{}) {
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

func (o CcpV1CustomPluginList) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o CcpV1CustomPluginList) MarshalJSON() ([]byte, error) {
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

type NullableCcpV1CustomPluginList struct {
	value *CcpV1CustomPluginList
	isSet bool
}

func (v NullableCcpV1CustomPluginList) Get() *CcpV1CustomPluginList {
	return v.value
}

func (v *NullableCcpV1CustomPluginList) Set(val *CcpV1CustomPluginList) {
	v.value = val
	v.isSet = true
}

func (v NullableCcpV1CustomPluginList) IsSet() bool {
	return v.isSet
}

func (v *NullableCcpV1CustomPluginList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCcpV1CustomPluginList(val *CcpV1CustomPluginList) *NullableCcpV1CustomPluginList {
	return &NullableCcpV1CustomPluginList{value: val, isSet: true}
}

func (v NullableCcpV1CustomPluginList) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableCcpV1CustomPluginList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}