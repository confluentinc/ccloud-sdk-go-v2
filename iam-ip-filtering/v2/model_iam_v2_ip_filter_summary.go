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
Identity & Access Management API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1-alpha0
Contact: paas-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// IamV2IpFilterSummary The IP Filter Summary endpoint returns an aggregation of the IP Filters across the system. This API can be queried in the context of an organization or an environment. It returns a summary of every operation group in the system grouped with a higher summary by operation group category. 
type IamV2IpFilterSummary struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind *string `json:"kind,omitempty"`
	// The scope associated with this object.
	Scope *ObjectReference `json:"scope,omitempty"`
	// Summary of the operation groups and IP filters created in those operation groups. 
	Categories *[]IamV2IpFilterSummaryCategories `json:"categories,omitempty"`
}

// NewIamV2IpFilterSummary instantiates a new IamV2IpFilterSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamV2IpFilterSummary() *IamV2IpFilterSummary {
	this := IamV2IpFilterSummary{}
	return &this
}

// NewIamV2IpFilterSummaryWithDefaults instantiates a new IamV2IpFilterSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamV2IpFilterSummaryWithDefaults() *IamV2IpFilterSummary {
	this := IamV2IpFilterSummary{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *IamV2IpFilterSummary) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2IpFilterSummary) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *IamV2IpFilterSummary) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *IamV2IpFilterSummary) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *IamV2IpFilterSummary) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2IpFilterSummary) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *IamV2IpFilterSummary) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *IamV2IpFilterSummary) SetKind(v string) {
	o.Kind = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *IamV2IpFilterSummary) GetScope() ObjectReference {
	if o == nil || o.Scope == nil {
		var ret ObjectReference
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2IpFilterSummary) GetScopeOk() (*ObjectReference, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *IamV2IpFilterSummary) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given ObjectReference and assigns it to the Scope field.
func (o *IamV2IpFilterSummary) SetScope(v ObjectReference) {
	o.Scope = &v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *IamV2IpFilterSummary) GetCategories() []IamV2IpFilterSummaryCategories {
	if o == nil || o.Categories == nil {
		var ret []IamV2IpFilterSummaryCategories
		return ret
	}
	return *o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2IpFilterSummary) GetCategoriesOk() (*[]IamV2IpFilterSummaryCategories, bool) {
	if o == nil || o.Categories == nil {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *IamV2IpFilterSummary) HasCategories() bool {
	if o != nil && o.Categories != nil {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []IamV2IpFilterSummaryCategories and assigns it to the Categories field.
func (o *IamV2IpFilterSummary) SetCategories(v []IamV2IpFilterSummaryCategories) {
	o.Categories = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *IamV2IpFilterSummary) Redact() {
    o.recurseRedact(o.ApiVersion)
    o.recurseRedact(o.Kind)
    o.recurseRedact(o.Scope)
    o.recurseRedact(o.Categories)
}

func (o *IamV2IpFilterSummary) recurseRedact(v interface{}) {
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

func (o IamV2IpFilterSummary) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o IamV2IpFilterSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api_version"] = o.ApiVersion
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableIamV2IpFilterSummary struct {
	value *IamV2IpFilterSummary
	isSet bool
}

func (v NullableIamV2IpFilterSummary) Get() *IamV2IpFilterSummary {
	return v.value
}

func (v *NullableIamV2IpFilterSummary) Set(val *IamV2IpFilterSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableIamV2IpFilterSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableIamV2IpFilterSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamV2IpFilterSummary(val *IamV2IpFilterSummary) *NullableIamV2IpFilterSummary {
	return &NullableIamV2IpFilterSummary{value: val, isSet: true}
}

func (v NullableIamV2IpFilterSummary) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableIamV2IpFilterSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

