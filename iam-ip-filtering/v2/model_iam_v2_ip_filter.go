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

// IamV2IpFilter `IP Filter` objects are bindings between IP Groups and Confluent resource(s). For example, a binding between \"CorpNet\" and \"Management APIs\" will enforce that access must come from one of the CIDR blocks associated with CorpNet. If there are multiple IP filters bound to a resource, a request matching any of the CIDR blocks for any of the IP Group will allow the request. If there are no IP Filters for a resource, then access will be granted to requests originating from any IP Address.   ## The IP Filters Model <SchemaDefinition schemaRef=\"#/components/schemas/iam.v2.IpFilter\" />
type IamV2IpFilter struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind *string `json:"kind,omitempty"`
	// ID is the \"natural identifier\" for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\"time\"); however, it may collide with IDs for other object `kinds` or objects of the same `kind` within a different scope/namespace (\"space\").
	Id       *string     `json:"id,omitempty"`
	Metadata *ObjectMeta `json:"metadata,omitempty"`
	// A human readable name for an IP Filter. Can contain any unicode letter or number, the ASCII space character, or any of the following special characters: `[`, `]`, `|`, `&`, `+`, `-`, `_`, `/`, `.`, `,`.
	FilterName *string `json:"filter_name,omitempty"`
	// Scope of resources covered by this IP filter. The only resource_group currently available is \"management\".
	ResourceGroup *string `json:"resource_group,omitempty"`
	// Scope of resources covered by this IP filter.
	OperationGroups *[]string `json:"operation_groups,omitempty"`
	// A list of IP Groups.
	IpGroups *[]GlobalObjectReference `json:"ip_groups,omitempty"`
}

// NewIamV2IpFilter instantiates a new IamV2IpFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamV2IpFilter() *IamV2IpFilter {
	this := IamV2IpFilter{}
	return &this
}

// NewIamV2IpFilterWithDefaults instantiates a new IamV2IpFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamV2IpFilterWithDefaults() *IamV2IpFilter {
	this := IamV2IpFilter{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *IamV2IpFilter) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2IpFilter) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *IamV2IpFilter) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *IamV2IpFilter) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *IamV2IpFilter) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2IpFilter) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *IamV2IpFilter) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *IamV2IpFilter) SetKind(v string) {
	o.Kind = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IamV2IpFilter) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2IpFilter) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IamV2IpFilter) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IamV2IpFilter) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *IamV2IpFilter) GetMetadata() ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2IpFilter) GetMetadataOk() (*ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *IamV2IpFilter) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ObjectMeta and assigns it to the Metadata field.
func (o *IamV2IpFilter) SetMetadata(v ObjectMeta) {
	o.Metadata = &v
}

// GetFilterName returns the FilterName field value if set, zero value otherwise.
func (o *IamV2IpFilter) GetFilterName() string {
	if o == nil || o.FilterName == nil {
		var ret string
		return ret
	}
	return *o.FilterName
}

// GetFilterNameOk returns a tuple with the FilterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2IpFilter) GetFilterNameOk() (*string, bool) {
	if o == nil || o.FilterName == nil {
		return nil, false
	}
	return o.FilterName, true
}

// HasFilterName returns a boolean if a field has been set.
func (o *IamV2IpFilter) HasFilterName() bool {
	if o != nil && o.FilterName != nil {
		return true
	}

	return false
}

// SetFilterName gets a reference to the given string and assigns it to the FilterName field.
func (o *IamV2IpFilter) SetFilterName(v string) {
	o.FilterName = &v
}

// GetResourceGroup returns the ResourceGroup field value if set, zero value otherwise.
func (o *IamV2IpFilter) GetResourceGroup() string {
	if o == nil || o.ResourceGroup == nil {
		var ret string
		return ret
	}
	return *o.ResourceGroup
}

// GetResourceGroupOk returns a tuple with the ResourceGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2IpFilter) GetResourceGroupOk() (*string, bool) {
	if o == nil || o.ResourceGroup == nil {
		return nil, false
	}
	return o.ResourceGroup, true
}

// HasResourceGroup returns a boolean if a field has been set.
func (o *IamV2IpFilter) HasResourceGroup() bool {
	if o != nil && o.ResourceGroup != nil {
		return true
	}

	return false
}

// SetResourceGroup gets a reference to the given string and assigns it to the ResourceGroup field.
func (o *IamV2IpFilter) SetResourceGroup(v string) {
	o.ResourceGroup = &v
}

// GetOperationGroups returns the OperationGroups field value if set, zero value otherwise.
func (o *IamV2IpFilter) GetOperationGroups() []string {
	if o == nil || o.OperationGroups == nil {
		var ret []string
		return ret
	}
	return *o.OperationGroups
}

// GetOperationGroupsOk returns a tuple with the OperationGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2IpFilter) GetOperationGroupsOk() (*[]string, bool) {
	if o == nil || o.OperationGroups == nil {
		return nil, false
	}
	return o.OperationGroups, true
}

// HasOperationGroups returns a boolean if a field has been set.
func (o *IamV2IpFilter) HasOperationGroups() bool {
	if o != nil && o.OperationGroups != nil {
		return true
	}

	return false
}

// SetOperationGroups gets a reference to the given []string and assigns it to the OperationGroups field.
func (o *IamV2IpFilter) SetOperationGroups(v []string) {
	o.OperationGroups = &v
}

// GetIpGroups returns the IpGroups field value if set, zero value otherwise.
func (o *IamV2IpFilter) GetIpGroups() []GlobalObjectReference {
	if o == nil || o.IpGroups == nil {
		var ret []GlobalObjectReference
		return ret
	}
	return *o.IpGroups
}

// GetIpGroupsOk returns a tuple with the IpGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2IpFilter) GetIpGroupsOk() (*[]GlobalObjectReference, bool) {
	if o == nil || o.IpGroups == nil {
		return nil, false
	}
	return o.IpGroups, true
}

// HasIpGroups returns a boolean if a field has been set.
func (o *IamV2IpFilter) HasIpGroups() bool {
	if o != nil && o.IpGroups != nil {
		return true
	}

	return false
}

// SetIpGroups gets a reference to the given []GlobalObjectReference and assigns it to the IpGroups field.
func (o *IamV2IpFilter) SetIpGroups(v []GlobalObjectReference) {
	o.IpGroups = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *IamV2IpFilter) Redact() {
	o.recurseRedact(o.ApiVersion)
	o.recurseRedact(o.Kind)
	o.recurseRedact(o.Id)
	o.recurseRedact(o.Metadata)
	o.recurseRedact(o.FilterName)
	o.recurseRedact(o.ResourceGroup)
	o.recurseRedact(o.OperationGroups)
	o.recurseRedact(o.IpGroups)
}

func (o *IamV2IpFilter) recurseRedact(v interface{}) {
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

func (o IamV2IpFilter) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o IamV2IpFilter) MarshalJSON() ([]byte, error) {
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
	if o.FilterName != nil {
		toSerialize["filter_name"] = o.FilterName
	}
	if o.ResourceGroup != nil {
		toSerialize["resource_group"] = o.ResourceGroup
	}
	if o.OperationGroups != nil {
		toSerialize["operation_groups"] = o.OperationGroups
	}
	if o.IpGroups != nil {
		toSerialize["ip_groups"] = o.IpGroups
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableIamV2IpFilter struct {
	value *IamV2IpFilter
	isSet bool
}

func (v NullableIamV2IpFilter) Get() *IamV2IpFilter {
	return v.value
}

func (v *NullableIamV2IpFilter) Set(val *IamV2IpFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableIamV2IpFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableIamV2IpFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamV2IpFilter(val *IamV2IpFilter) *NullableIamV2IpFilter {
	return &NullableIamV2IpFilter{value: val, isSet: true}
}

func (v NullableIamV2IpFilter) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableIamV2IpFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}