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
Tableflow Management API

Tableflow Management API

API version: 0.0.1
Contact: cts-team@confluent.io
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

// TableflowV1CatalogIntegrationStatus The status of the Catalog Integration
type TableflowV1CatalogIntegrationStatus struct {
	// The lifecycle phase of the catalog integration:    PENDING: sync to catalog integration is pending;    CONNECTED: catalog integration is connected and syncing;    FAILED: catalog integration failed.
	Phase *string `json:"phase,omitempty"`
	// Displayable error message if catalog integration is in a failed state.
	ErrorMessage *string `json:"error_message,omitempty"`
	// The date and time at which the catalog was last synced. It is represented in RFC3339 format and is in UTC.
	LastSyncAt *string `json:"last_sync_at,omitempty"`
}

// NewTableflowV1CatalogIntegrationStatus instantiates a new TableflowV1CatalogIntegrationStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTableflowV1CatalogIntegrationStatus() *TableflowV1CatalogIntegrationStatus {
	this := TableflowV1CatalogIntegrationStatus{}
	return &this
}

// NewTableflowV1CatalogIntegrationStatusWithDefaults instantiates a new TableflowV1CatalogIntegrationStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTableflowV1CatalogIntegrationStatusWithDefaults() *TableflowV1CatalogIntegrationStatus {
	this := TableflowV1CatalogIntegrationStatus{}
	return &this
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *TableflowV1CatalogIntegrationStatus) GetPhase() string {
	if o == nil || o.Phase == nil {
		var ret string
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableflowV1CatalogIntegrationStatus) GetPhaseOk() (*string, bool) {
	if o == nil || o.Phase == nil {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *TableflowV1CatalogIntegrationStatus) HasPhase() bool {
	if o != nil && o.Phase != nil {
		return true
	}

	return false
}

// SetPhase gets a reference to the given string and assigns it to the Phase field.
func (o *TableflowV1CatalogIntegrationStatus) SetPhase(v string) {
	o.Phase = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *TableflowV1CatalogIntegrationStatus) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableflowV1CatalogIntegrationStatus) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *TableflowV1CatalogIntegrationStatus) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *TableflowV1CatalogIntegrationStatus) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetLastSyncAt returns the LastSyncAt field value if set, zero value otherwise.
func (o *TableflowV1CatalogIntegrationStatus) GetLastSyncAt() string {
	if o == nil || o.LastSyncAt == nil {
		var ret string
		return ret
	}
	return *o.LastSyncAt
}

// GetLastSyncAtOk returns a tuple with the LastSyncAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableflowV1CatalogIntegrationStatus) GetLastSyncAtOk() (*string, bool) {
	if o == nil || o.LastSyncAt == nil {
		return nil, false
	}
	return o.LastSyncAt, true
}

// HasLastSyncAt returns a boolean if a field has been set.
func (o *TableflowV1CatalogIntegrationStatus) HasLastSyncAt() bool {
	if o != nil && o.LastSyncAt != nil {
		return true
	}

	return false
}

// SetLastSyncAt gets a reference to the given string and assigns it to the LastSyncAt field.
func (o *TableflowV1CatalogIntegrationStatus) SetLastSyncAt(v string) {
	o.LastSyncAt = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *TableflowV1CatalogIntegrationStatus) Redact() {
	o.recurseRedact(o.Phase)
	o.recurseRedact(o.ErrorMessage)
	o.recurseRedact(o.LastSyncAt)
}

func (o *TableflowV1CatalogIntegrationStatus) recurseRedact(v interface{}) {
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

func (o TableflowV1CatalogIntegrationStatus) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o TableflowV1CatalogIntegrationStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Phase != nil {
		toSerialize["phase"] = o.Phase
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.LastSyncAt != nil {
		toSerialize["last_sync_at"] = o.LastSyncAt
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableTableflowV1CatalogIntegrationStatus struct {
	value *TableflowV1CatalogIntegrationStatus
	isSet bool
}

func (v NullableTableflowV1CatalogIntegrationStatus) Get() *TableflowV1CatalogIntegrationStatus {
	return v.value
}

func (v *NullableTableflowV1CatalogIntegrationStatus) Set(val *TableflowV1CatalogIntegrationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTableflowV1CatalogIntegrationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTableflowV1CatalogIntegrationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTableflowV1CatalogIntegrationStatus(val *TableflowV1CatalogIntegrationStatus) *NullableTableflowV1CatalogIntegrationStatus {
	return &NullableTableflowV1CatalogIntegrationStatus{value: val, isSet: true}
}

func (v NullableTableflowV1CatalogIntegrationStatus) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableTableflowV1CatalogIntegrationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
