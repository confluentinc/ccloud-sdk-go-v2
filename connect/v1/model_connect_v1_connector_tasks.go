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
Kafka Connect APIs

REST API for managing connectors

API version: 1.0
Contact: connect@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

import (
	"reflect"
)

// ConnectV1ConnectorTasks struct for ConnectV1ConnectorTasks
type ConnectV1ConnectorTasks struct {
	// The name of the connector the task belongs to
	Connector string `json:"connector"`
	// Task ID within the connector
	Task int32 `json:"task"`
}

// NewConnectV1ConnectorTasks instantiates a new ConnectV1ConnectorTasks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectV1ConnectorTasks(connector string, task int32) *ConnectV1ConnectorTasks {
	this := ConnectV1ConnectorTasks{}
	this.Connector = connector
	this.Task = task
	return &this
}

// NewConnectV1ConnectorTasksWithDefaults instantiates a new ConnectV1ConnectorTasks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectV1ConnectorTasksWithDefaults() *ConnectV1ConnectorTasks {
	this := ConnectV1ConnectorTasks{}
	return &this
}

// GetConnector returns the Connector field value
func (o *ConnectV1ConnectorTasks) GetConnector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value
// and a boolean to check if the value has been set.
func (o *ConnectV1ConnectorTasks) GetConnectorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Connector, true
}

// SetConnector sets field value
func (o *ConnectV1ConnectorTasks) SetConnector(v string) {
	o.Connector = v
}

// GetTask returns the Task field value
func (o *ConnectV1ConnectorTasks) GetTask() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Task
}

// GetTaskOk returns a tuple with the Task field value
// and a boolean to check if the value has been set.
func (o *ConnectV1ConnectorTasks) GetTaskOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Task, true
}

// SetTask sets field value
func (o *ConnectV1ConnectorTasks) SetTask(v int32) {
	o.Task = v
}

// Redact resets all sensitive fields to their zero value.
func (o *ConnectV1ConnectorTasks) Redact() {
    o.recurseRedact(&o.Connector)
    o.recurseRedact(&o.Task)
}

func (o *ConnectV1ConnectorTasks) recurseRedact(v interface{}) {
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

func (o ConnectV1ConnectorTasks) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o ConnectV1ConnectorTasks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["connector"] = o.Connector
	}
	if true {
		toSerialize["task"] = o.Task
	}
	return json.Marshal(toSerialize)
}

type NullableConnectV1ConnectorTasks struct {
	value *ConnectV1ConnectorTasks
	isSet bool
}

func (v NullableConnectV1ConnectorTasks) Get() *ConnectV1ConnectorTasks {
	return v.value
}

func (v *NullableConnectV1ConnectorTasks) Set(val *ConnectV1ConnectorTasks) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectV1ConnectorTasks) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectV1ConnectorTasks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectV1ConnectorTasks(val *ConnectV1ConnectorTasks) *NullableConnectV1ConnectorTasks {
	return &NullableConnectV1ConnectorTasks{value: val, isSet: true}
}

func (v NullableConnectV1ConnectorTasks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectV1ConnectorTasks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

