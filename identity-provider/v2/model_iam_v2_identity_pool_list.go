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
OAuth Identity Management API

OAuth Public API

API version: 0.0.1-alpha1
Contact: oauth-eng@confluent.io
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

// IamV2IdentityPoolList `IdentityPool` objects represent groups of identities tied to a given a `IdentityProvider` that authorizes them to Confluent Cloud resources.  It provides a mapping functionality of your `Identity Provider` user to a Confluent identity pool that is then used to provide access to Confluent Resources.   Related guide: [Use identity pools with your OAuth provider](https://docs.confluent.io/cloud/current/access-management/authenticate/oauth/identity-pools.html).  ## The Identity Pools Model <SchemaDefinition schemaRef=\"#/components/schemas/iam.v2.IdentityPool\" />  ## Quotas and Limits This resource is subject to the [following quotas](https://docs.confluent.io/cloud/current/quotas/overview.html):  | Quota | Description | | --- | --- | | `identity_pools_per_provider` | Number of Identity Pools per Identity Provider |
type IamV2IdentityPoolList struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind     string   `json:"kind,omitempty"`
	Metadata ListMeta `json:"metadata,omitempty"`
	// A data property that contains an array of resource items. Each entry in the array is a separate resource.
	Data []IamV2IdentityPool `json:"data,omitempty"`
}

// NewIamV2IdentityPoolList instantiates a new IamV2IdentityPoolList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamV2IdentityPoolList(apiVersion string, kind string, metadata ListMeta, data []IamV2IdentityPool) *IamV2IdentityPoolList {
	this := IamV2IdentityPoolList{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Data = data
	return &this
}

// NewIamV2IdentityPoolListWithDefaults instantiates a new IamV2IdentityPoolList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamV2IdentityPoolListWithDefaults() *IamV2IdentityPoolList {
	this := IamV2IdentityPoolList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *IamV2IdentityPoolList) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *IamV2IdentityPoolList) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *IamV2IdentityPoolList) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *IamV2IdentityPoolList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *IamV2IdentityPoolList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *IamV2IdentityPoolList) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *IamV2IdentityPoolList) GetMetadata() ListMeta {
	if o == nil {
		var ret ListMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *IamV2IdentityPoolList) GetMetadataOk() (*ListMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *IamV2IdentityPoolList) SetMetadata(v ListMeta) {
	o.Metadata = v
}

// GetData returns the Data field value
func (o *IamV2IdentityPoolList) GetData() []IamV2IdentityPool {
	if o == nil {
		var ret []IamV2IdentityPool
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IamV2IdentityPoolList) GetDataOk() (*[]IamV2IdentityPool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *IamV2IdentityPoolList) SetData(v []IamV2IdentityPool) {
	o.Data = v
}

// Redact resets all sensitive fields to their zero value.
func (o *IamV2IdentityPoolList) Redact() {
	o.recurseRedact(&o.ApiVersion)
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Metadata)
	o.recurseRedact(&o.Data)
}

func (o *IamV2IdentityPoolList) recurseRedact(v interface{}) {
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

func (o IamV2IdentityPoolList) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o IamV2IdentityPoolList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["api_version"] = o.ApiVersion
	}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["data"] = o.Data
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableIamV2IdentityPoolList struct {
	value *IamV2IdentityPoolList
	isSet bool
}

func (v NullableIamV2IdentityPoolList) Get() *IamV2IdentityPoolList {
	return v.value
}

func (v *NullableIamV2IdentityPoolList) Set(val *IamV2IdentityPoolList) {
	v.value = val
	v.isSet = true
}

func (v NullableIamV2IdentityPoolList) IsSet() bool {
	return v.isSet
}

func (v *NullableIamV2IdentityPoolList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamV2IdentityPoolList(val *IamV2IdentityPoolList) *NullableIamV2IdentityPoolList {
	return &NullableIamV2IdentityPoolList{value: val, isSet: true}
}

func (v NullableIamV2IdentityPoolList) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableIamV2IdentityPoolList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
