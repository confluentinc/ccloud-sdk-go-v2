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
Custom Connect Plugin Management API

This is Custom Connect Plugin Management API.

API version: 0.1.0
Contact: connect-team@confluent.io
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

// CcpmV1PresignedUrl Request a presigned upload URL for new Custom Connect Plugin. Note that the URL policy expires in one hour. If the policy expires, you can request a new presigned upload URL.  Related guide: [Custom Connect Plugin API](https://docs.confluent.io/cloud/current/connectors/connect-api-section.html).   ## The Presigned Urls Model <SchemaDefinition schemaRef=\"#/components/schemas/ccpm.v1.PresignedUrl\" />
type CcpmV1PresignedUrl struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind *string `json:"kind,omitempty"`
	// Content format of the Custom Connect Plugin archive.
	ContentFormat *string `json:"content_format,omitempty"`
	// Cloud provider where the Custom Connect Plugin archive is uploaded.
	Cloud *string `json:"cloud,omitempty"`
	// Unique identifier of this upload.
	UploadId *string `json:"upload_id,omitempty"`
	// Upload URL for the Custom Connect Plugin archive.
	UploadUrl *string `json:"upload_url,omitempty"`
	// Upload form data of the Custom Connect Plugin. All values should be strings.
	UploadFormData *map[string]interface{} `json:"upload_form_data,omitempty"`
	// The environment to which this belongs.
	Environment *EnvScopedObjectReference `json:"environment,omitempty"`
}

// NewCcpmV1PresignedUrl instantiates a new CcpmV1PresignedUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCcpmV1PresignedUrl() *CcpmV1PresignedUrl {
	this := CcpmV1PresignedUrl{}
	return &this
}

// NewCcpmV1PresignedUrlWithDefaults instantiates a new CcpmV1PresignedUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCcpmV1PresignedUrlWithDefaults() *CcpmV1PresignedUrl {
	this := CcpmV1PresignedUrl{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *CcpmV1PresignedUrl) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CcpmV1PresignedUrl) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *CcpmV1PresignedUrl) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *CcpmV1PresignedUrl) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *CcpmV1PresignedUrl) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CcpmV1PresignedUrl) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *CcpmV1PresignedUrl) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *CcpmV1PresignedUrl) SetKind(v string) {
	o.Kind = &v
}

// GetContentFormat returns the ContentFormat field value if set, zero value otherwise.
func (o *CcpmV1PresignedUrl) GetContentFormat() string {
	if o == nil || o.ContentFormat == nil {
		var ret string
		return ret
	}
	return *o.ContentFormat
}

// GetContentFormatOk returns a tuple with the ContentFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CcpmV1PresignedUrl) GetContentFormatOk() (*string, bool) {
	if o == nil || o.ContentFormat == nil {
		return nil, false
	}
	return o.ContentFormat, true
}

// HasContentFormat returns a boolean if a field has been set.
func (o *CcpmV1PresignedUrl) HasContentFormat() bool {
	if o != nil && o.ContentFormat != nil {
		return true
	}

	return false
}

// SetContentFormat gets a reference to the given string and assigns it to the ContentFormat field.
func (o *CcpmV1PresignedUrl) SetContentFormat(v string) {
	o.ContentFormat = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *CcpmV1PresignedUrl) GetCloud() string {
	if o == nil || o.Cloud == nil {
		var ret string
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CcpmV1PresignedUrl) GetCloudOk() (*string, bool) {
	if o == nil || o.Cloud == nil {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *CcpmV1PresignedUrl) HasCloud() bool {
	if o != nil && o.Cloud != nil {
		return true
	}

	return false
}

// SetCloud gets a reference to the given string and assigns it to the Cloud field.
func (o *CcpmV1PresignedUrl) SetCloud(v string) {
	o.Cloud = &v
}

// GetUploadId returns the UploadId field value if set, zero value otherwise.
func (o *CcpmV1PresignedUrl) GetUploadId() string {
	if o == nil || o.UploadId == nil {
		var ret string
		return ret
	}
	return *o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CcpmV1PresignedUrl) GetUploadIdOk() (*string, bool) {
	if o == nil || o.UploadId == nil {
		return nil, false
	}
	return o.UploadId, true
}

// HasUploadId returns a boolean if a field has been set.
func (o *CcpmV1PresignedUrl) HasUploadId() bool {
	if o != nil && o.UploadId != nil {
		return true
	}

	return false
}

// SetUploadId gets a reference to the given string and assigns it to the UploadId field.
func (o *CcpmV1PresignedUrl) SetUploadId(v string) {
	o.UploadId = &v
}

// GetUploadUrl returns the UploadUrl field value if set, zero value otherwise.
func (o *CcpmV1PresignedUrl) GetUploadUrl() string {
	if o == nil || o.UploadUrl == nil {
		var ret string
		return ret
	}
	return *o.UploadUrl
}

// GetUploadUrlOk returns a tuple with the UploadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CcpmV1PresignedUrl) GetUploadUrlOk() (*string, bool) {
	if o == nil || o.UploadUrl == nil {
		return nil, false
	}
	return o.UploadUrl, true
}

// HasUploadUrl returns a boolean if a field has been set.
func (o *CcpmV1PresignedUrl) HasUploadUrl() bool {
	if o != nil && o.UploadUrl != nil {
		return true
	}

	return false
}

// SetUploadUrl gets a reference to the given string and assigns it to the UploadUrl field.
func (o *CcpmV1PresignedUrl) SetUploadUrl(v string) {
	o.UploadUrl = &v
}

// GetUploadFormData returns the UploadFormData field value if set, zero value otherwise.
func (o *CcpmV1PresignedUrl) GetUploadFormData() map[string]interface{} {
	if o == nil || o.UploadFormData == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.UploadFormData
}

// GetUploadFormDataOk returns a tuple with the UploadFormData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CcpmV1PresignedUrl) GetUploadFormDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.UploadFormData == nil {
		return nil, false
	}
	return o.UploadFormData, true
}

// HasUploadFormData returns a boolean if a field has been set.
func (o *CcpmV1PresignedUrl) HasUploadFormData() bool {
	if o != nil && o.UploadFormData != nil {
		return true
	}

	return false
}

// SetUploadFormData gets a reference to the given map[string]interface{} and assigns it to the UploadFormData field.
func (o *CcpmV1PresignedUrl) SetUploadFormData(v map[string]interface{}) {
	o.UploadFormData = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *CcpmV1PresignedUrl) GetEnvironment() EnvScopedObjectReference {
	if o == nil || o.Environment == nil {
		var ret EnvScopedObjectReference
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CcpmV1PresignedUrl) GetEnvironmentOk() (*EnvScopedObjectReference, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *CcpmV1PresignedUrl) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given EnvScopedObjectReference and assigns it to the Environment field.
func (o *CcpmV1PresignedUrl) SetEnvironment(v EnvScopedObjectReference) {
	o.Environment = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *CcpmV1PresignedUrl) Redact() {
	o.recurseRedact(o.ApiVersion)
	o.recurseRedact(o.Kind)
	o.recurseRedact(o.ContentFormat)
	o.recurseRedact(o.Cloud)
	o.recurseRedact(o.UploadId)
	o.recurseRedact(o.UploadUrl)
	o.recurseRedact(o.UploadFormData)
	o.recurseRedact(o.Environment)
}

func (o *CcpmV1PresignedUrl) recurseRedact(v interface{}) {
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

func (o CcpmV1PresignedUrl) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o CcpmV1PresignedUrl) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api_version"] = o.ApiVersion
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.ContentFormat != nil {
		toSerialize["content_format"] = o.ContentFormat
	}
	if o.Cloud != nil {
		toSerialize["cloud"] = o.Cloud
	}
	if o.UploadId != nil {
		toSerialize["upload_id"] = o.UploadId
	}
	if o.UploadUrl != nil {
		toSerialize["upload_url"] = o.UploadUrl
	}
	if o.UploadFormData != nil {
		toSerialize["upload_form_data"] = o.UploadFormData
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableCcpmV1PresignedUrl struct {
	value *CcpmV1PresignedUrl
	isSet bool
}

func (v NullableCcpmV1PresignedUrl) Get() *CcpmV1PresignedUrl {
	return v.value
}

func (v *NullableCcpmV1PresignedUrl) Set(val *CcpmV1PresignedUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableCcpmV1PresignedUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableCcpmV1PresignedUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCcpmV1PresignedUrl(val *CcpmV1PresignedUrl) *NullableCcpmV1PresignedUrl {
	return &NullableCcpmV1PresignedUrl{value: val, isSet: true}
}

func (v NullableCcpmV1PresignedUrl) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableCcpmV1PresignedUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
