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

// TableflowV1CatalogIntegrationUpdateRequest The desired state of the Catalog Integration
type TableflowV1CatalogIntegrationUpdateRequest struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind *string `json:"kind,omitempty"`
	// ID is the \"natural identifier\" for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\"time\"); however, it may collide with IDs for other object `kinds` or objects of the same `kind` within a different scope/namespace (\"space\").
	Id       *string                                  `json:"id,omitempty"`
	Metadata *ObjectMeta                              `json:"metadata,omitempty"`
	Spec     *TableflowV1CatalogIntegrationUpdateSpec `json:"spec,omitempty"`
}

// NewTableflowV1CatalogIntegrationUpdateRequest instantiates a new TableflowV1CatalogIntegrationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTableflowV1CatalogIntegrationUpdateRequest() *TableflowV1CatalogIntegrationUpdateRequest {
	this := TableflowV1CatalogIntegrationUpdateRequest{}
	return &this
}

// NewTableflowV1CatalogIntegrationUpdateRequestWithDefaults instantiates a new TableflowV1CatalogIntegrationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTableflowV1CatalogIntegrationUpdateRequestWithDefaults() *TableflowV1CatalogIntegrationUpdateRequest {
	this := TableflowV1CatalogIntegrationUpdateRequest{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *TableflowV1CatalogIntegrationUpdateRequest) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableflowV1CatalogIntegrationUpdateRequest) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *TableflowV1CatalogIntegrationUpdateRequest) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *TableflowV1CatalogIntegrationUpdateRequest) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *TableflowV1CatalogIntegrationUpdateRequest) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableflowV1CatalogIntegrationUpdateRequest) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *TableflowV1CatalogIntegrationUpdateRequest) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *TableflowV1CatalogIntegrationUpdateRequest) SetKind(v string) {
	o.Kind = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TableflowV1CatalogIntegrationUpdateRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableflowV1CatalogIntegrationUpdateRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TableflowV1CatalogIntegrationUpdateRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TableflowV1CatalogIntegrationUpdateRequest) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *TableflowV1CatalogIntegrationUpdateRequest) GetMetadata() ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableflowV1CatalogIntegrationUpdateRequest) GetMetadataOk() (*ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TableflowV1CatalogIntegrationUpdateRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ObjectMeta and assigns it to the Metadata field.
func (o *TableflowV1CatalogIntegrationUpdateRequest) SetMetadata(v ObjectMeta) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *TableflowV1CatalogIntegrationUpdateRequest) GetSpec() TableflowV1CatalogIntegrationUpdateSpec {
	if o == nil || o.Spec == nil {
		var ret TableflowV1CatalogIntegrationUpdateSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableflowV1CatalogIntegrationUpdateRequest) GetSpecOk() (*TableflowV1CatalogIntegrationUpdateSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *TableflowV1CatalogIntegrationUpdateRequest) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given TableflowV1CatalogIntegrationUpdateSpec and assigns it to the Spec field.
func (o *TableflowV1CatalogIntegrationUpdateRequest) SetSpec(v TableflowV1CatalogIntegrationUpdateSpec) {
	o.Spec = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *TableflowV1CatalogIntegrationUpdateRequest) Redact() {
	o.recurseRedact(o.ApiVersion)
	o.recurseRedact(o.Kind)
	o.recurseRedact(o.Id)
	o.recurseRedact(o.Metadata)
	o.recurseRedact(o.Spec)
}

func (o *TableflowV1CatalogIntegrationUpdateRequest) recurseRedact(v interface{}) {
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

func (o TableflowV1CatalogIntegrationUpdateRequest) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o TableflowV1CatalogIntegrationUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api_version"] = o.ApiVersion
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableTableflowV1CatalogIntegrationUpdateRequest struct {
	value *TableflowV1CatalogIntegrationUpdateRequest
	isSet bool
}

func (v NullableTableflowV1CatalogIntegrationUpdateRequest) Get() *TableflowV1CatalogIntegrationUpdateRequest {
	return v.value
}

func (v *NullableTableflowV1CatalogIntegrationUpdateRequest) Set(val *TableflowV1CatalogIntegrationUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTableflowV1CatalogIntegrationUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTableflowV1CatalogIntegrationUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTableflowV1CatalogIntegrationUpdateRequest(val *TableflowV1CatalogIntegrationUpdateRequest) *NullableTableflowV1CatalogIntegrationUpdateRequest {
	return &NullableTableflowV1CatalogIntegrationUpdateRequest{value: val, isSet: true}
}

func (v NullableTableflowV1CatalogIntegrationUpdateRequest) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableTableflowV1CatalogIntegrationUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
