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
Custom Code Logging API

This is Custom Code Logging API.

API version: 1.0.0
Contact: compute-platform-team@confluent.io
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

// CclV1CustomCodeLogging CustomCodeLogging objects represent Custom Code Logging on Confluent Cloud. The API allows you to list, create, read, update, and delete your Custom Code Logging.   ## The Custom Code Loggings Model <SchemaDefinition schemaRef=\"#/components/schemas/ccl.v1.CustomCodeLogging\" />
type CclV1CustomCodeLogging struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind *string `json:"kind,omitempty"`
	// ID is the \"natural identifier\" for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\"time\"); however, it may collide with IDs for other object `kinds` or objects of the same `kind` within a different scope/namespace (\"space\").
	Id       *string     `json:"id,omitempty"`
	Metadata *ObjectMeta `json:"metadata,omitempty"`
	// Cloud provider where the Custom Code Logging is sent.
	Cloud *string `json:"cloud,omitempty"`
	// The Cloud provider region the Custom Code Logging is sent.
	Region *string `json:"region,omitempty"`
	// Destination Settings of the Custom Code Logging.
	DestinationSettings *CclV1CustomCodeLoggingDestinationSettingsOneOf `json:"destination_settings,omitempty"`
	// The environment to which this belongs.
	Environment *EnvScopedObjectReference `json:"environment,omitempty"`
}

// NewCclV1CustomCodeLogging instantiates a new CclV1CustomCodeLogging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCclV1CustomCodeLogging() *CclV1CustomCodeLogging {
	this := CclV1CustomCodeLogging{}
	return &this
}

// NewCclV1CustomCodeLoggingWithDefaults instantiates a new CclV1CustomCodeLogging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCclV1CustomCodeLoggingWithDefaults() *CclV1CustomCodeLogging {
	this := CclV1CustomCodeLogging{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *CclV1CustomCodeLogging) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CclV1CustomCodeLogging) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *CclV1CustomCodeLogging) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *CclV1CustomCodeLogging) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *CclV1CustomCodeLogging) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CclV1CustomCodeLogging) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *CclV1CustomCodeLogging) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *CclV1CustomCodeLogging) SetKind(v string) {
	o.Kind = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CclV1CustomCodeLogging) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CclV1CustomCodeLogging) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CclV1CustomCodeLogging) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CclV1CustomCodeLogging) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CclV1CustomCodeLogging) GetMetadata() ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CclV1CustomCodeLogging) GetMetadataOk() (*ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CclV1CustomCodeLogging) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ObjectMeta and assigns it to the Metadata field.
func (o *CclV1CustomCodeLogging) SetMetadata(v ObjectMeta) {
	o.Metadata = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *CclV1CustomCodeLogging) GetCloud() string {
	if o == nil || o.Cloud == nil {
		var ret string
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CclV1CustomCodeLogging) GetCloudOk() (*string, bool) {
	if o == nil || o.Cloud == nil {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *CclV1CustomCodeLogging) HasCloud() bool {
	if o != nil && o.Cloud != nil {
		return true
	}

	return false
}

// SetCloud gets a reference to the given string and assigns it to the Cloud field.
func (o *CclV1CustomCodeLogging) SetCloud(v string) {
	o.Cloud = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CclV1CustomCodeLogging) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CclV1CustomCodeLogging) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CclV1CustomCodeLogging) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CclV1CustomCodeLogging) SetRegion(v string) {
	o.Region = &v
}

// GetDestinationSettings returns the DestinationSettings field value if set, zero value otherwise.
func (o *CclV1CustomCodeLogging) GetDestinationSettings() CclV1CustomCodeLoggingDestinationSettingsOneOf {
	if o == nil || o.DestinationSettings == nil {
		var ret CclV1CustomCodeLoggingDestinationSettingsOneOf
		return ret
	}
	return *o.DestinationSettings
}

// GetDestinationSettingsOk returns a tuple with the DestinationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CclV1CustomCodeLogging) GetDestinationSettingsOk() (*CclV1CustomCodeLoggingDestinationSettingsOneOf, bool) {
	if o == nil || o.DestinationSettings == nil {
		return nil, false
	}
	return o.DestinationSettings, true
}

// HasDestinationSettings returns a boolean if a field has been set.
func (o *CclV1CustomCodeLogging) HasDestinationSettings() bool {
	if o != nil && o.DestinationSettings != nil {
		return true
	}

	return false
}

// SetDestinationSettings gets a reference to the given CclV1CustomCodeLoggingDestinationSettingsOneOf and assigns it to the DestinationSettings field.
func (o *CclV1CustomCodeLogging) SetDestinationSettings(v CclV1CustomCodeLoggingDestinationSettingsOneOf) {
	o.DestinationSettings = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *CclV1CustomCodeLogging) GetEnvironment() EnvScopedObjectReference {
	if o == nil || o.Environment == nil {
		var ret EnvScopedObjectReference
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CclV1CustomCodeLogging) GetEnvironmentOk() (*EnvScopedObjectReference, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *CclV1CustomCodeLogging) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given EnvScopedObjectReference and assigns it to the Environment field.
func (o *CclV1CustomCodeLogging) SetEnvironment(v EnvScopedObjectReference) {
	o.Environment = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *CclV1CustomCodeLogging) Redact() {
	o.recurseRedact(o.ApiVersion)
	o.recurseRedact(o.Kind)
	o.recurseRedact(o.Id)
	o.recurseRedact(o.Metadata)
	o.recurseRedact(o.Cloud)
	o.recurseRedact(o.Region)
	o.recurseRedact(o.DestinationSettings)
	o.recurseRedact(o.Environment)
}

func (o *CclV1CustomCodeLogging) recurseRedact(v interface{}) {
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

func (o CclV1CustomCodeLogging) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o CclV1CustomCodeLogging) MarshalJSON() ([]byte, error) {
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
	if o.Cloud != nil {
		toSerialize["cloud"] = o.Cloud
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.DestinationSettings != nil {
		toSerialize["destination_settings"] = o.DestinationSettings
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

type NullableCclV1CustomCodeLogging struct {
	value *CclV1CustomCodeLogging
	isSet bool
}

func (v NullableCclV1CustomCodeLogging) Get() *CclV1CustomCodeLogging {
	return v.value
}

func (v *NullableCclV1CustomCodeLogging) Set(val *CclV1CustomCodeLogging) {
	v.value = val
	v.isSet = true
}

func (v NullableCclV1CustomCodeLogging) IsSet() bool {
	return v.isSet
}

func (v *NullableCclV1CustomCodeLogging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCclV1CustomCodeLogging(val *CclV1CustomCodeLogging) *NullableCclV1CustomCodeLogging {
	return &NullableCclV1CustomCodeLogging{value: val, isSet: true}
}

func (v NullableCclV1CustomCodeLogging) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableCclV1CustomCodeLogging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}