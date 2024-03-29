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
SQL API v1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
Contact: flink-control-plane@confluent.io
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

// SqlV1StatementList `Statement` represents a core resource used to model SQL statements for execution. A statement generalizes DDL, DML, DQL, etc., but doesn’t attempt to handle session management or any higher-level functionality. The API allows you to list, create, read, and delete your statements. ## The Statements Model <SchemaDefinition schemaRef=\"#/components/schemas/sql.v1.Statement\" />
type SqlV1StatementList struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind     string   `json:"kind,omitempty"`
	Metadata ListMeta `json:"metadata,omitempty"`
	// A data property that contains an array of resource items. Each entry in the array is a separate resource.
	Data []SqlV1Statement `json:"data,omitempty"`
}

// NewSqlV1StatementList instantiates a new SqlV1StatementList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSqlV1StatementList(apiVersion string, kind string, metadata ListMeta, data []SqlV1Statement) *SqlV1StatementList {
	this := SqlV1StatementList{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Data = data
	return &this
}

// NewSqlV1StatementListWithDefaults instantiates a new SqlV1StatementList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSqlV1StatementListWithDefaults() *SqlV1StatementList {
	this := SqlV1StatementList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *SqlV1StatementList) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *SqlV1StatementList) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *SqlV1StatementList) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *SqlV1StatementList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *SqlV1StatementList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *SqlV1StatementList) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *SqlV1StatementList) GetMetadata() ListMeta {
	if o == nil {
		var ret ListMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *SqlV1StatementList) GetMetadataOk() (*ListMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *SqlV1StatementList) SetMetadata(v ListMeta) {
	o.Metadata = v
}

// GetData returns the Data field value
func (o *SqlV1StatementList) GetData() []SqlV1Statement {
	if o == nil {
		var ret []SqlV1Statement
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SqlV1StatementList) GetDataOk() (*[]SqlV1Statement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SqlV1StatementList) SetData(v []SqlV1Statement) {
	o.Data = v
}

// Redact resets all sensitive fields to their zero value.
func (o *SqlV1StatementList) Redact() {
	o.recurseRedact(&o.ApiVersion)
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Metadata)
	o.recurseRedact(&o.Data)
}

func (o *SqlV1StatementList) recurseRedact(v interface{}) {
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

func (o SqlV1StatementList) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o SqlV1StatementList) MarshalJSON() ([]byte, error) {
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

type NullableSqlV1StatementList struct {
	value *SqlV1StatementList
	isSet bool
}

func (v NullableSqlV1StatementList) Get() *SqlV1StatementList {
	return v.value
}

func (v *NullableSqlV1StatementList) Set(val *SqlV1StatementList) {
	v.value = val
	v.isSet = true
}

func (v NullableSqlV1StatementList) IsSet() bool {
	return v.isSet
}

func (v *NullableSqlV1StatementList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSqlV1StatementList(val *SqlV1StatementList) *NullableSqlV1StatementList {
	return &NullableSqlV1StatementList{value: val, isSet: true}
}

func (v NullableSqlV1StatementList) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableSqlV1StatementList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
