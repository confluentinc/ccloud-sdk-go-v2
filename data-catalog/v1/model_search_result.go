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
Confluent Data Catalog

REST API for the Data Catalog

API version: 1.0.0
Contact: data-governance@confluent.io
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

// SearchResult struct for SearchResult
type SearchResult struct {
	SearchParameters *SearchParams `json:"searchParameters,omitempty"`
	// The types
	Types *[]string `json:"types,omitempty"`
	// The entities
	Entities *[]EntityHeader `json:"entities,omitempty"`
	// The referred entities
	ReferredEntities *map[string]EntityHeader `json:"referredEntities,omitempty"`
}

// NewSearchResult instantiates a new SearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchResult() *SearchResult {
	this := SearchResult{}
	return &this
}

// NewSearchResultWithDefaults instantiates a new SearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchResultWithDefaults() *SearchResult {
	this := SearchResult{}
	return &this
}

// GetSearchParameters returns the SearchParameters field value if set, zero value otherwise.
func (o *SearchResult) GetSearchParameters() SearchParams {
	if o == nil || o.SearchParameters == nil {
		var ret SearchParams
		return ret
	}
	return *o.SearchParameters
}

// GetSearchParametersOk returns a tuple with the SearchParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResult) GetSearchParametersOk() (*SearchParams, bool) {
	if o == nil || o.SearchParameters == nil {
		return nil, false
	}
	return o.SearchParameters, true
}

// HasSearchParameters returns a boolean if a field has been set.
func (o *SearchResult) HasSearchParameters() bool {
	if o != nil && o.SearchParameters != nil {
		return true
	}

	return false
}

// SetSearchParameters gets a reference to the given SearchParams and assigns it to the SearchParameters field.
func (o *SearchResult) SetSearchParameters(v SearchParams) {
	o.SearchParameters = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *SearchResult) GetTypes() []string {
	if o == nil || o.Types == nil {
		var ret []string
		return ret
	}
	return *o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResult) GetTypesOk() (*[]string, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *SearchResult) HasTypes() bool {
	if o != nil && o.Types != nil {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *SearchResult) SetTypes(v []string) {
	o.Types = &v
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *SearchResult) GetEntities() []EntityHeader {
	if o == nil || o.Entities == nil {
		var ret []EntityHeader
		return ret
	}
	return *o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResult) GetEntitiesOk() (*[]EntityHeader, bool) {
	if o == nil || o.Entities == nil {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *SearchResult) HasEntities() bool {
	if o != nil && o.Entities != nil {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []EntityHeader and assigns it to the Entities field.
func (o *SearchResult) SetEntities(v []EntityHeader) {
	o.Entities = &v
}

// GetReferredEntities returns the ReferredEntities field value if set, zero value otherwise.
func (o *SearchResult) GetReferredEntities() map[string]EntityHeader {
	if o == nil || o.ReferredEntities == nil {
		var ret map[string]EntityHeader
		return ret
	}
	return *o.ReferredEntities
}

// GetReferredEntitiesOk returns a tuple with the ReferredEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResult) GetReferredEntitiesOk() (*map[string]EntityHeader, bool) {
	if o == nil || o.ReferredEntities == nil {
		return nil, false
	}
	return o.ReferredEntities, true
}

// HasReferredEntities returns a boolean if a field has been set.
func (o *SearchResult) HasReferredEntities() bool {
	if o != nil && o.ReferredEntities != nil {
		return true
	}

	return false
}

// SetReferredEntities gets a reference to the given map[string]EntityHeader and assigns it to the ReferredEntities field.
func (o *SearchResult) SetReferredEntities(v map[string]EntityHeader) {
	o.ReferredEntities = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *SearchResult) Redact() {
	o.recurseRedact(o.SearchParameters)
	o.recurseRedact(o.Types)
	o.recurseRedact(o.Entities)
	o.recurseRedact(o.ReferredEntities)
}

func (o *SearchResult) recurseRedact(v interface{}) {
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

func (o SearchResult) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o SearchResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SearchParameters != nil {
		toSerialize["searchParameters"] = o.SearchParameters
	}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}
	if o.Entities != nil {
		toSerialize["entities"] = o.Entities
	}
	if o.ReferredEntities != nil {
		toSerialize["referredEntities"] = o.ReferredEntities
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableSearchResult struct {
	value *SearchResult
	isSet bool
}

func (v NullableSearchResult) Get() *SearchResult {
	return v.value
}

func (v *NullableSearchResult) Set(val *SearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchResult(val *SearchResult) *NullableSearchResult {
	return &NullableSearchResult{value: val, isSet: true}
}

func (v NullableSearchResult) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
