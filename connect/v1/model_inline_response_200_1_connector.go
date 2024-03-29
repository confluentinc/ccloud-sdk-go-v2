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
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// InlineResponse2001Connector The map containing connector status.
type InlineResponse2001Connector struct {
	// The state of the connector.
	State string `json:"state,omitempty"`
	// The worker ID of the connector.
	WorkerId string `json:"worker_id,omitempty"`
	// The exception name in case of error.
	Trace *string `json:"trace,omitempty"`
}

// NewInlineResponse2001Connector instantiates a new InlineResponse2001Connector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2001Connector(state string, workerId string) *InlineResponse2001Connector {
	this := InlineResponse2001Connector{}
	this.State = state
	this.WorkerId = workerId
	return &this
}

// NewInlineResponse2001ConnectorWithDefaults instantiates a new InlineResponse2001Connector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2001ConnectorWithDefaults() *InlineResponse2001Connector {
	this := InlineResponse2001Connector{}
	return &this
}

// GetState returns the State field value
func (o *InlineResponse2001Connector) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001Connector) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *InlineResponse2001Connector) SetState(v string) {
	o.State = v
}

// GetWorkerId returns the WorkerId field value
func (o *InlineResponse2001Connector) GetWorkerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkerId
}

// GetWorkerIdOk returns a tuple with the WorkerId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001Connector) GetWorkerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkerId, true
}

// SetWorkerId sets field value
func (o *InlineResponse2001Connector) SetWorkerId(v string) {
	o.WorkerId = v
}

// GetTrace returns the Trace field value if set, zero value otherwise.
func (o *InlineResponse2001Connector) GetTrace() string {
	if o == nil || o.Trace == nil {
		var ret string
		return ret
	}
	return *o.Trace
}

// GetTraceOk returns a tuple with the Trace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001Connector) GetTraceOk() (*string, bool) {
	if o == nil || o.Trace == nil {
		return nil, false
	}
	return o.Trace, true
}

// HasTrace returns a boolean if a field has been set.
func (o *InlineResponse2001Connector) HasTrace() bool {
	if o != nil && o.Trace != nil {
		return true
	}

	return false
}

// SetTrace gets a reference to the given string and assigns it to the Trace field.
func (o *InlineResponse2001Connector) SetTrace(v string) {
	o.Trace = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *InlineResponse2001Connector) Redact() {
	o.recurseRedact(&o.State)
	o.recurseRedact(&o.WorkerId)
	o.recurseRedact(o.Trace)
}

func (o *InlineResponse2001Connector) recurseRedact(v interface{}) {
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

func (o InlineResponse2001Connector) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o InlineResponse2001Connector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["worker_id"] = o.WorkerId
	}
	if o.Trace != nil {
		toSerialize["trace"] = o.Trace
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableInlineResponse2001Connector struct {
	value *InlineResponse2001Connector
	isSet bool
}

func (v NullableInlineResponse2001Connector) Get() *InlineResponse2001Connector {
	return v.value
}

func (v *NullableInlineResponse2001Connector) Set(val *InlineResponse2001Connector) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2001Connector) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2001Connector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2001Connector(val *InlineResponse2001Connector) *NullableInlineResponse2001Connector {
	return &NullableInlineResponse2001Connector{value: val, isSet: true}
}

func (v NullableInlineResponse2001Connector) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableInlineResponse2001Connector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
