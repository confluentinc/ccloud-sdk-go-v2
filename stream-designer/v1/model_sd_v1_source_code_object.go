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

// SdV1SourceCodeObject A object containing pipeline's source code definition.
type SdV1SourceCodeObject struct {
	// A list of KSQL statements that defines a pipeline.
	Sql string `json:"sql"`
}

// NewSdV1SourceCodeObject instantiates a new SdV1SourceCodeObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdV1SourceCodeObject(sql string) *SdV1SourceCodeObject {
	this := SdV1SourceCodeObject{}
	this.Sql = sql
	return &this
}

// NewSdV1SourceCodeObjectWithDefaults instantiates a new SdV1SourceCodeObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdV1SourceCodeObjectWithDefaults() *SdV1SourceCodeObject {
	this := SdV1SourceCodeObject{}
	return &this
}

// GetSql returns the Sql field value
func (o *SdV1SourceCodeObject) GetSql() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sql
}

// GetSqlOk returns a tuple with the Sql field value
// and a boolean to check if the value has been set.
func (o *SdV1SourceCodeObject) GetSqlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sql, true
}

// SetSql sets field value
func (o *SdV1SourceCodeObject) SetSql(v string) {
	o.Sql = v
}

// Redact resets all sensitive fields to their zero value.
func (o *SdV1SourceCodeObject) Redact() {
    o.recurseRedact(&o.Sql)
}

func (o *SdV1SourceCodeObject) recurseRedact(v interface{}) {
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

func (o SdV1SourceCodeObject) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o SdV1SourceCodeObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sql"] = o.Sql
	}
	return json.Marshal(toSerialize)
}

type NullableSdV1SourceCodeObject struct {
	value *SdV1SourceCodeObject
	isSet bool
}

func (v NullableSdV1SourceCodeObject) Get() *SdV1SourceCodeObject {
	return v.value
}

func (v *NullableSdV1SourceCodeObject) Set(val *SdV1SourceCodeObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSdV1SourceCodeObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSdV1SourceCodeObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdV1SourceCodeObject(val *SdV1SourceCodeObject) *NullableSdV1SourceCodeObject {
	return &NullableSdV1SourceCodeObject{value: val, isSet: true}
}

func (v NullableSdV1SourceCodeObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdV1SourceCodeObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


