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

// ConnectV1ConnectorExpansionNameId The ID of connector.
type ConnectV1ConnectorExpansionNameId struct {
	// The ID of the connector.
	Id *string `json:"id,omitempty"`
	// Type of the value in the `id` property.
	IdType *string `json:"id_type,omitempty"`
}

// NewConnectV1ConnectorExpansionNameId instantiates a new ConnectV1ConnectorExpansionNameId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectV1ConnectorExpansionNameId() *ConnectV1ConnectorExpansionNameId {
	this := ConnectV1ConnectorExpansionNameId{}
	return &this
}

// NewConnectV1ConnectorExpansionNameIdWithDefaults instantiates a new ConnectV1ConnectorExpansionNameId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectV1ConnectorExpansionNameIdWithDefaults() *ConnectV1ConnectorExpansionNameId {
	this := ConnectV1ConnectorExpansionNameId{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConnectV1ConnectorExpansionNameId) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1ConnectorExpansionNameId) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConnectV1ConnectorExpansionNameId) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConnectV1ConnectorExpansionNameId) SetId(v string) {
	o.Id = &v
}

// GetIdType returns the IdType field value if set, zero value otherwise.
func (o *ConnectV1ConnectorExpansionNameId) GetIdType() string {
	if o == nil || o.IdType == nil {
		var ret string
		return ret
	}
	return *o.IdType
}

// GetIdTypeOk returns a tuple with the IdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1ConnectorExpansionNameId) GetIdTypeOk() (*string, bool) {
	if o == nil || o.IdType == nil {
		return nil, false
	}
	return o.IdType, true
}

// HasIdType returns a boolean if a field has been set.
func (o *ConnectV1ConnectorExpansionNameId) HasIdType() bool {
	if o != nil && o.IdType != nil {
		return true
	}

	return false
}

// SetIdType gets a reference to the given string and assigns it to the IdType field.
func (o *ConnectV1ConnectorExpansionNameId) SetIdType(v string) {
	o.IdType = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *ConnectV1ConnectorExpansionNameId) Redact() {
    o.recurseRedact(o.Id)
    o.recurseRedact(o.IdType)
}

func (o *ConnectV1ConnectorExpansionNameId) recurseRedact(v interface{}) {
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

func (o ConnectV1ConnectorExpansionNameId) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o ConnectV1ConnectorExpansionNameId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IdType != nil {
		toSerialize["id_type"] = o.IdType
	}
	return json.Marshal(toSerialize)
}

type NullableConnectV1ConnectorExpansionNameId struct {
	value *ConnectV1ConnectorExpansionNameId
	isSet bool
}

func (v NullableConnectV1ConnectorExpansionNameId) Get() *ConnectV1ConnectorExpansionNameId {
	return v.value
}

func (v *NullableConnectV1ConnectorExpansionNameId) Set(val *ConnectV1ConnectorExpansionNameId) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectV1ConnectorExpansionNameId) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectV1ConnectorExpansionNameId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectV1ConnectorExpansionNameId(val *ConnectV1ConnectorExpansionNameId) *NullableConnectV1ConnectorExpansionNameId {
	return &NullableConnectV1ConnectorExpansionNameId{value: val, isSet: true}
}

func (v NullableConnectV1ConnectorExpansionNameId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectV1ConnectorExpansionNameId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

