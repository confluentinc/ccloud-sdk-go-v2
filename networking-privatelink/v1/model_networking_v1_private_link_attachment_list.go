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
Network API

Network API

API version: 0.0.1-alpha1
Contact: cire-traffic@confluent.io
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

// NetworkingV1PrivateLinkAttachmentList PrivateLink attachment objects represent reservations to establish PrivateLink connections to a cloud region in order to access resources that belong to a Confluent Cloud Environment. The API allows you to list, create, read update and delete your PrivateLink attachments.   ## The Private Link Attachments Model <SchemaDefinition schemaRef=\"#/components/schemas/networking.v1.PrivateLinkAttachment\" />  ## Quotas and Limits This resource is subject to the [following quotas](https://docs.confluent.io/cloud/current/quotas/overview.html):  | Quota | Description | | --- | --- | | `private_link_attachments_per_environment` | Number of PrivateLink Attachments per environment |
type NetworkingV1PrivateLinkAttachmentList struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind     string   `json:"kind,omitempty"`
	Metadata ListMeta `json:"metadata,omitempty"`
	// A data property that contains an array of resource items. Each entry in the array is a separate resource.
	Data []NetworkingV1PrivateLinkAttachment `json:"data,omitempty"`
}

// NewNetworkingV1PrivateLinkAttachmentList instantiates a new NetworkingV1PrivateLinkAttachmentList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1PrivateLinkAttachmentList(apiVersion string, kind string, metadata ListMeta, data []NetworkingV1PrivateLinkAttachment) *NetworkingV1PrivateLinkAttachmentList {
	this := NetworkingV1PrivateLinkAttachmentList{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Data = data
	return &this
}

// NewNetworkingV1PrivateLinkAttachmentListWithDefaults instantiates a new NetworkingV1PrivateLinkAttachmentList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1PrivateLinkAttachmentListWithDefaults() *NetworkingV1PrivateLinkAttachmentList {
	this := NetworkingV1PrivateLinkAttachmentList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *NetworkingV1PrivateLinkAttachmentList) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1PrivateLinkAttachmentList) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *NetworkingV1PrivateLinkAttachmentList) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *NetworkingV1PrivateLinkAttachmentList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1PrivateLinkAttachmentList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *NetworkingV1PrivateLinkAttachmentList) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *NetworkingV1PrivateLinkAttachmentList) GetMetadata() ListMeta {
	if o == nil {
		var ret ListMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1PrivateLinkAttachmentList) GetMetadataOk() (*ListMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *NetworkingV1PrivateLinkAttachmentList) SetMetadata(v ListMeta) {
	o.Metadata = v
}

// GetData returns the Data field value
func (o *NetworkingV1PrivateLinkAttachmentList) GetData() []NetworkingV1PrivateLinkAttachment {
	if o == nil {
		var ret []NetworkingV1PrivateLinkAttachment
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1PrivateLinkAttachmentList) GetDataOk() (*[]NetworkingV1PrivateLinkAttachment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *NetworkingV1PrivateLinkAttachmentList) SetData(v []NetworkingV1PrivateLinkAttachment) {
	o.Data = v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1PrivateLinkAttachmentList) Redact() {
	o.recurseRedact(&o.ApiVersion)
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Metadata)
	o.recurseRedact(&o.Data)
}

func (o *NetworkingV1PrivateLinkAttachmentList) recurseRedact(v interface{}) {
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

func (o NetworkingV1PrivateLinkAttachmentList) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1PrivateLinkAttachmentList) MarshalJSON() ([]byte, error) {
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

type NullableNetworkingV1PrivateLinkAttachmentList struct {
	value *NetworkingV1PrivateLinkAttachmentList
	isSet bool
}

func (v NullableNetworkingV1PrivateLinkAttachmentList) Get() *NetworkingV1PrivateLinkAttachmentList {
	return v.value
}

func (v *NullableNetworkingV1PrivateLinkAttachmentList) Set(val *NetworkingV1PrivateLinkAttachmentList) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1PrivateLinkAttachmentList) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1PrivateLinkAttachmentList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1PrivateLinkAttachmentList(val *NetworkingV1PrivateLinkAttachmentList) *NullableNetworkingV1PrivateLinkAttachmentList {
	return &NullableNetworkingV1PrivateLinkAttachmentList{value: val, isSet: true}
}

func (v NullableNetworkingV1PrivateLinkAttachmentList) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableNetworkingV1PrivateLinkAttachmentList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
