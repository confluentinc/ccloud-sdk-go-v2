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
Organization Management API

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

// OrgV2Environment `Environment` objects represent an isolated namespace for your Confluent resources for organizational purposes.  The API allows you to create, delete, and update your environments. You can retrieve individual environments as well as a list of all your environments.   Related guide: [Environments in Confluent Cloud](https://docs.confluent.io/cloud/current/access-management/environments.html).  ## The Environments Model <SchemaDefinition schemaRef=\"#/components/schemas/org.v2.Environment\" />  ## Quotas and Limits This resource is subject to the [following quotas](https://docs.confluent.io/cloud/current/quotas/overview.html):  | Quota | Description | | --- | --- | | `environments_per_org` | Environments in one Confluent Cloud organization |
type OrgV2Environment struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind *string `json:"kind,omitempty"`
	// ID is the \"natural identifier\" for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\"time\"); however, it may collide with IDs for other object `kinds` or objects of the same `kind` within a different scope/namespace (\"space\").
	Id       *string     `json:"id,omitempty"`
	Metadata *ObjectMeta `json:"metadata,omitempty"`
	// A human-readable name for the Environment
	DisplayName *string `json:"display_name,omitempty"`
	// Stream Governance configurations for the environment
	StreamGovernanceConfig *OrgV2StreamGovernanceConfig `json:"stream_governance_config,omitempty"`
}

// NewOrgV2Environment instantiates a new OrgV2Environment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgV2Environment() *OrgV2Environment {
	this := OrgV2Environment{}
	return &this
}

// NewOrgV2EnvironmentWithDefaults instantiates a new OrgV2Environment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgV2EnvironmentWithDefaults() *OrgV2Environment {
	this := OrgV2Environment{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *OrgV2Environment) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgV2Environment) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *OrgV2Environment) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *OrgV2Environment) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *OrgV2Environment) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgV2Environment) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *OrgV2Environment) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *OrgV2Environment) SetKind(v string) {
	o.Kind = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrgV2Environment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgV2Environment) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrgV2Environment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrgV2Environment) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *OrgV2Environment) GetMetadata() ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgV2Environment) GetMetadataOk() (*ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *OrgV2Environment) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ObjectMeta and assigns it to the Metadata field.
func (o *OrgV2Environment) SetMetadata(v ObjectMeta) {
	o.Metadata = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *OrgV2Environment) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgV2Environment) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *OrgV2Environment) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *OrgV2Environment) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetStreamGovernanceConfig returns the StreamGovernanceConfig field value if set, zero value otherwise.
func (o *OrgV2Environment) GetStreamGovernanceConfig() OrgV2StreamGovernanceConfig {
	if o == nil || o.StreamGovernanceConfig == nil {
		var ret OrgV2StreamGovernanceConfig
		return ret
	}
	return *o.StreamGovernanceConfig
}

// GetStreamGovernanceConfigOk returns a tuple with the StreamGovernanceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgV2Environment) GetStreamGovernanceConfigOk() (*OrgV2StreamGovernanceConfig, bool) {
	if o == nil || o.StreamGovernanceConfig == nil {
		return nil, false
	}
	return o.StreamGovernanceConfig, true
}

// HasStreamGovernanceConfig returns a boolean if a field has been set.
func (o *OrgV2Environment) HasStreamGovernanceConfig() bool {
	if o != nil && o.StreamGovernanceConfig != nil {
		return true
	}

	return false
}

// SetStreamGovernanceConfig gets a reference to the given OrgV2StreamGovernanceConfig and assigns it to the StreamGovernanceConfig field.
func (o *OrgV2Environment) SetStreamGovernanceConfig(v OrgV2StreamGovernanceConfig) {
	o.StreamGovernanceConfig = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *OrgV2Environment) Redact() {
	o.recurseRedact(o.ApiVersion)
	o.recurseRedact(o.Kind)
	o.recurseRedact(o.Id)
	o.recurseRedact(o.Metadata)
	o.recurseRedact(o.DisplayName)
	o.recurseRedact(o.StreamGovernanceConfig)
}

func (o *OrgV2Environment) recurseRedact(v interface{}) {
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

func (o OrgV2Environment) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o OrgV2Environment) MarshalJSON() ([]byte, error) {
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
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.StreamGovernanceConfig != nil {
		toSerialize["stream_governance_config"] = o.StreamGovernanceConfig
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableOrgV2Environment struct {
	value *OrgV2Environment
	isSet bool
}

func (v NullableOrgV2Environment) Get() *OrgV2Environment {
	return v.value
}

func (v *NullableOrgV2Environment) Set(val *OrgV2Environment) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgV2Environment) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgV2Environment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgV2Environment(val *OrgV2Environment) *NullableOrgV2Environment {
	return &NullableOrgV2Environment{value: val, isSet: true}
}

func (v NullableOrgV2Environment) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableOrgV2Environment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
