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
Custom Connector Plugin Management API

This is Custom Connector Plugin Management API.

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

// ConnectV1CustomConnectorPluginVersion CustomConnectorPluginVersion objects represent Custom Connector Plugin Versions on Confluent Cloud. The API allows you to list, create, read, update, and delete your Custom Connector Plugin Versions.   ## The Custom Connector Plugin Versions Model <SchemaDefinition schemaRef=\"#/components/schemas/connect.v1.CustomConnectorPluginVersion\" />
type ConnectV1CustomConnectorPluginVersion struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind *string `json:"kind,omitempty"`
	// ID is the \"natural identifier\" for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\"time\"); however, it may collide with IDs for other object `kinds` or objects of the same `kind` within a different scope/namespace (\"space\").
	Id       *string     `json:"id,omitempty"`
	Metadata *ObjectMeta `json:"metadata,omitempty"`
	// Version of the Custom Connector Plugin.
	Version *string `json:"version,omitempty"`
	// Release Notes of the Custom Connector Plugin Version.
	ReleaseNotes *string `json:"release_notes,omitempty"`
	// A sensitive property is a connector configuration property that must be hidden after a user enters property value when setting up connector.
	SensitiveConfigProperties *[]string `json:"sensitive_config_properties,omitempty"`
	// Flag to specify stability of the version
	IsBeta *string `json:"is_beta,omitempty"`
	// Upload source of Custom Connector Plugin Version. Only required in `create` request, will be ignored in `read`, `update` or `list`.
	UploadSource *ConnectV1CustomConnectorPluginVersionUploadSourceOneOf `json:"upload_source,omitempty"`
}

// NewConnectV1CustomConnectorPluginVersion instantiates a new ConnectV1CustomConnectorPluginVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectV1CustomConnectorPluginVersion() *ConnectV1CustomConnectorPluginVersion {
	this := ConnectV1CustomConnectorPluginVersion{}
	return &this
}

// NewConnectV1CustomConnectorPluginVersionWithDefaults instantiates a new ConnectV1CustomConnectorPluginVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectV1CustomConnectorPluginVersionWithDefaults() *ConnectV1CustomConnectorPluginVersion {
	this := ConnectV1CustomConnectorPluginVersion{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ConnectV1CustomConnectorPluginVersion) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorPluginVersion) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ConnectV1CustomConnectorPluginVersion) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ConnectV1CustomConnectorPluginVersion) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ConnectV1CustomConnectorPluginVersion) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorPluginVersion) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ConnectV1CustomConnectorPluginVersion) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ConnectV1CustomConnectorPluginVersion) SetKind(v string) {
	o.Kind = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConnectV1CustomConnectorPluginVersion) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorPluginVersion) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConnectV1CustomConnectorPluginVersion) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConnectV1CustomConnectorPluginVersion) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ConnectV1CustomConnectorPluginVersion) GetMetadata() ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorPluginVersion) GetMetadataOk() (*ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ConnectV1CustomConnectorPluginVersion) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ObjectMeta and assigns it to the Metadata field.
func (o *ConnectV1CustomConnectorPluginVersion) SetMetadata(v ObjectMeta) {
	o.Metadata = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ConnectV1CustomConnectorPluginVersion) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorPluginVersion) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ConnectV1CustomConnectorPluginVersion) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ConnectV1CustomConnectorPluginVersion) SetVersion(v string) {
	o.Version = &v
}

// GetReleaseNotes returns the ReleaseNotes field value if set, zero value otherwise.
func (o *ConnectV1CustomConnectorPluginVersion) GetReleaseNotes() string {
	if o == nil || o.ReleaseNotes == nil {
		var ret string
		return ret
	}
	return *o.ReleaseNotes
}

// GetReleaseNotesOk returns a tuple with the ReleaseNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorPluginVersion) GetReleaseNotesOk() (*string, bool) {
	if o == nil || o.ReleaseNotes == nil {
		return nil, false
	}
	return o.ReleaseNotes, true
}

// HasReleaseNotes returns a boolean if a field has been set.
func (o *ConnectV1CustomConnectorPluginVersion) HasReleaseNotes() bool {
	if o != nil && o.ReleaseNotes != nil {
		return true
	}

	return false
}

// SetReleaseNotes gets a reference to the given string and assigns it to the ReleaseNotes field.
func (o *ConnectV1CustomConnectorPluginVersion) SetReleaseNotes(v string) {
	o.ReleaseNotes = &v
}

// GetSensitiveConfigProperties returns the SensitiveConfigProperties field value if set, zero value otherwise.
func (o *ConnectV1CustomConnectorPluginVersion) GetSensitiveConfigProperties() []string {
	if o == nil || o.SensitiveConfigProperties == nil {
		var ret []string
		return ret
	}
	return *o.SensitiveConfigProperties
}

// GetSensitiveConfigPropertiesOk returns a tuple with the SensitiveConfigProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorPluginVersion) GetSensitiveConfigPropertiesOk() (*[]string, bool) {
	if o == nil || o.SensitiveConfigProperties == nil {
		return nil, false
	}
	return o.SensitiveConfigProperties, true
}

// HasSensitiveConfigProperties returns a boolean if a field has been set.
func (o *ConnectV1CustomConnectorPluginVersion) HasSensitiveConfigProperties() bool {
	if o != nil && o.SensitiveConfigProperties != nil {
		return true
	}

	return false
}

// SetSensitiveConfigProperties gets a reference to the given []string and assigns it to the SensitiveConfigProperties field.
func (o *ConnectV1CustomConnectorPluginVersion) SetSensitiveConfigProperties(v []string) {
	o.SensitiveConfigProperties = &v
}

// GetIsBeta returns the IsBeta field value if set, zero value otherwise.
func (o *ConnectV1CustomConnectorPluginVersion) GetIsBeta() string {
	if o == nil || o.IsBeta == nil {
		var ret string
		return ret
	}
	return *o.IsBeta
}

// GetIsBetaOk returns a tuple with the IsBeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorPluginVersion) GetIsBetaOk() (*string, bool) {
	if o == nil || o.IsBeta == nil {
		return nil, false
	}
	return o.IsBeta, true
}

// HasIsBeta returns a boolean if a field has been set.
func (o *ConnectV1CustomConnectorPluginVersion) HasIsBeta() bool {
	if o != nil && o.IsBeta != nil {
		return true
	}

	return false
}

// SetIsBeta gets a reference to the given string and assigns it to the IsBeta field.
func (o *ConnectV1CustomConnectorPluginVersion) SetIsBeta(v string) {
	o.IsBeta = &v
}

// GetUploadSource returns the UploadSource field value if set, zero value otherwise.
func (o *ConnectV1CustomConnectorPluginVersion) GetUploadSource() ConnectV1CustomConnectorPluginVersionUploadSourceOneOf {
	if o == nil || o.UploadSource == nil {
		var ret ConnectV1CustomConnectorPluginVersionUploadSourceOneOf
		return ret
	}
	return *o.UploadSource
}

// GetUploadSourceOk returns a tuple with the UploadSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorPluginVersion) GetUploadSourceOk() (*ConnectV1CustomConnectorPluginVersionUploadSourceOneOf, bool) {
	if o == nil || o.UploadSource == nil {
		return nil, false
	}
	return o.UploadSource, true
}

// HasUploadSource returns a boolean if a field has been set.
func (o *ConnectV1CustomConnectorPluginVersion) HasUploadSource() bool {
	if o != nil && o.UploadSource != nil {
		return true
	}

	return false
}

// SetUploadSource gets a reference to the given ConnectV1CustomConnectorPluginVersionUploadSourceOneOf and assigns it to the UploadSource field.
func (o *ConnectV1CustomConnectorPluginVersion) SetUploadSource(v ConnectV1CustomConnectorPluginVersionUploadSourceOneOf) {
	o.UploadSource = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *ConnectV1CustomConnectorPluginVersion) Redact() {
	o.recurseRedact(o.ApiVersion)
	o.recurseRedact(o.Kind)
	o.recurseRedact(o.Id)
	o.recurseRedact(o.Metadata)
	o.recurseRedact(o.Version)
	o.recurseRedact(o.ReleaseNotes)
	o.recurseRedact(o.SensitiveConfigProperties)
	o.recurseRedact(o.IsBeta)
	o.recurseRedact(o.UploadSource)
}

func (o *ConnectV1CustomConnectorPluginVersion) recurseRedact(v interface{}) {
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

func (o ConnectV1CustomConnectorPluginVersion) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ConnectV1CustomConnectorPluginVersion) MarshalJSON() ([]byte, error) {
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
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.ReleaseNotes != nil {
		toSerialize["release_notes"] = o.ReleaseNotes
	}
	if o.SensitiveConfigProperties != nil {
		toSerialize["sensitive_config_properties"] = o.SensitiveConfigProperties
	}
	if o.IsBeta != nil {
		toSerialize["is_beta"] = o.IsBeta
	}
	if o.UploadSource != nil {
		toSerialize["upload_source"] = o.UploadSource
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableConnectV1CustomConnectorPluginVersion struct {
	value *ConnectV1CustomConnectorPluginVersion
	isSet bool
}

func (v NullableConnectV1CustomConnectorPluginVersion) Get() *ConnectV1CustomConnectorPluginVersion {
	return v.value
}

func (v *NullableConnectV1CustomConnectorPluginVersion) Set(val *ConnectV1CustomConnectorPluginVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectV1CustomConnectorPluginVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectV1CustomConnectorPluginVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectV1CustomConnectorPluginVersion(val *ConnectV1CustomConnectorPluginVersion) *NullableConnectV1CustomConnectorPluginVersion {
	return &NullableConnectV1CustomConnectorPluginVersion{value: val, isSet: true}
}

func (v NullableConnectV1CustomConnectorPluginVersion) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableConnectV1CustomConnectorPluginVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}