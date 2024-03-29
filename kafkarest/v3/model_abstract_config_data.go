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
REST Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0.0
Contact: kafka-clients-proxy-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// AbstractConfigData struct for AbstractConfigData
type AbstractConfigData struct {
	Kind        string              `json:"kind,omitempty"`
	Metadata    ResourceMetadata    `json:"metadata,omitempty"`
	ClusterId   string              `json:"cluster_id,omitempty"`
	Name        string              `json:"name,omitempty"`
	Value       NullableString      `json:"value,omitempty"`
	IsDefault   bool                `json:"is_default,omitempty"`
	IsReadOnly  bool                `json:"is_read_only,omitempty"`
	IsSensitive bool                `json:"is_sensitive,omitempty"`
	Source      string              `json:"source,omitempty"`
	Synonyms    []ConfigSynonymData `json:"synonyms,omitempty"`
}

// NewAbstractConfigData instantiates a new AbstractConfigData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAbstractConfigData(kind string, metadata ResourceMetadata, clusterId string, name string, isDefault bool, isReadOnly bool, isSensitive bool, source string, synonyms []ConfigSynonymData) *AbstractConfigData {
	this := AbstractConfigData{}
	this.Kind = kind
	this.Metadata = metadata
	this.ClusterId = clusterId
	this.Name = name
	this.IsDefault = isDefault
	this.IsReadOnly = isReadOnly
	this.IsSensitive = isSensitive
	this.Source = source
	this.Synonyms = synonyms
	return &this
}

// NewAbstractConfigDataWithDefaults instantiates a new AbstractConfigData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAbstractConfigDataWithDefaults() *AbstractConfigData {
	this := AbstractConfigData{}
	return &this
}

// GetKind returns the Kind field value
func (o *AbstractConfigData) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *AbstractConfigData) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *AbstractConfigData) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *AbstractConfigData) GetMetadata() ResourceMetadata {
	if o == nil {
		var ret ResourceMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *AbstractConfigData) GetMetadataOk() (*ResourceMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *AbstractConfigData) SetMetadata(v ResourceMetadata) {
	o.Metadata = v
}

// GetClusterId returns the ClusterId field value
func (o *AbstractConfigData) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *AbstractConfigData) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *AbstractConfigData) SetClusterId(v string) {
	o.ClusterId = v
}

// GetName returns the Name field value
func (o *AbstractConfigData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AbstractConfigData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AbstractConfigData) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AbstractConfigData) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AbstractConfigData) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *AbstractConfigData) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *AbstractConfigData) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *AbstractConfigData) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *AbstractConfigData) UnsetValue() {
	o.Value.Unset()
}

// GetIsDefault returns the IsDefault field value
func (o *AbstractConfigData) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *AbstractConfigData) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value
func (o *AbstractConfigData) SetIsDefault(v bool) {
	o.IsDefault = v
}

// GetIsReadOnly returns the IsReadOnly field value
func (o *AbstractConfigData) GetIsReadOnly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsReadOnly
}

// GetIsReadOnlyOk returns a tuple with the IsReadOnly field value
// and a boolean to check if the value has been set.
func (o *AbstractConfigData) GetIsReadOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsReadOnly, true
}

// SetIsReadOnly sets field value
func (o *AbstractConfigData) SetIsReadOnly(v bool) {
	o.IsReadOnly = v
}

// GetIsSensitive returns the IsSensitive field value
func (o *AbstractConfigData) GetIsSensitive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSensitive
}

// GetIsSensitiveOk returns a tuple with the IsSensitive field value
// and a boolean to check if the value has been set.
func (o *AbstractConfigData) GetIsSensitiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSensitive, true
}

// SetIsSensitive sets field value
func (o *AbstractConfigData) SetIsSensitive(v bool) {
	o.IsSensitive = v
}

// GetSource returns the Source field value
func (o *AbstractConfigData) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *AbstractConfigData) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *AbstractConfigData) SetSource(v string) {
	o.Source = v
}

// GetSynonyms returns the Synonyms field value
func (o *AbstractConfigData) GetSynonyms() []ConfigSynonymData {
	if o == nil {
		var ret []ConfigSynonymData
		return ret
	}

	return o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value
// and a boolean to check if the value has been set.
func (o *AbstractConfigData) GetSynonymsOk() (*[]ConfigSynonymData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Synonyms, true
}

// SetSynonyms sets field value
func (o *AbstractConfigData) SetSynonyms(v []ConfigSynonymData) {
	o.Synonyms = v
}

// Redact resets all sensitive fields to their zero value.
func (o *AbstractConfigData) Redact() {
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Metadata)
	o.recurseRedact(&o.ClusterId)
	o.recurseRedact(&o.Name)
	o.recurseRedact(o.Value)
	o.recurseRedact(&o.IsDefault)
	o.recurseRedact(&o.IsReadOnly)
	o.recurseRedact(&o.IsSensitive)
	o.recurseRedact(&o.Source)
	o.recurseRedact(&o.Synonyms)
}

func (o *AbstractConfigData) recurseRedact(v interface{}) {
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

func (o AbstractConfigData) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o AbstractConfigData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if true {
		toSerialize["is_default"] = o.IsDefault
	}
	if true {
		toSerialize["is_read_only"] = o.IsReadOnly
	}
	if true {
		toSerialize["is_sensitive"] = o.IsSensitive
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["synonyms"] = o.Synonyms
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableAbstractConfigData struct {
	value *AbstractConfigData
	isSet bool
}

func (v NullableAbstractConfigData) Get() *AbstractConfigData {
	return v.value
}

func (v *NullableAbstractConfigData) Set(val *AbstractConfigData) {
	v.value = val
	v.isSet = true
}

func (v NullableAbstractConfigData) IsSet() bool {
	return v.isSet
}

func (v *NullableAbstractConfigData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAbstractConfigData(val *AbstractConfigData) *NullableAbstractConfigData {
	return &NullableAbstractConfigData{value: val, isSet: true}
}

func (v NullableAbstractConfigData) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableAbstractConfigData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
