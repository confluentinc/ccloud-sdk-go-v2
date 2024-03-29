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

// NetworkingV1NetworkLinkEndpointUpdate A Network Link Enpoint is associated with a Private Link Confluent Cloud Network at the origin and a Network Link Service (associated with another Private Link Confluent Cloud Network) at the target. It enables connectivity between the origin network and the target network. It can only be associated with a Private Link network.   Related guide: [Network Linking Overview](https://docs.confluent.io/cloud/current/networking/network-linking.html).  ## The Network Link Endpoints Model <SchemaDefinition schemaRef=\"#/components/schemas/networking.v1.NetworkLinkEndpoint\" />  ## Quotas and Limits This resource is subject to the [following quotas](https://docs.confluent.io/cloud/current/quotas/overview.html):  | Quota | Description | | --- | --- | | `network_link_endpoints_per_network` | Number of network link endpoints per network |
type NetworkingV1NetworkLinkEndpointUpdate struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind *string `json:"kind,omitempty"`
	// ID is the \"natural identifier\" for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\"time\"); however, it may collide with IDs for other object `kinds` or objects of the same `kind` within a different scope/namespace (\"space\").
	Id       *string                                    `json:"id,omitempty"`
	Metadata *ObjectMeta                                `json:"metadata,omitempty"`
	Spec     *NetworkingV1NetworkLinkEndpointSpecUpdate `json:"spec,omitempty"`
	Status   *NetworkingV1NetworkLinkEndpointStatus     `json:"status,omitempty"`
}

// NewNetworkingV1NetworkLinkEndpointUpdate instantiates a new NetworkingV1NetworkLinkEndpointUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1NetworkLinkEndpointUpdate() *NetworkingV1NetworkLinkEndpointUpdate {
	this := NetworkingV1NetworkLinkEndpointUpdate{}
	return &this
}

// NewNetworkingV1NetworkLinkEndpointUpdateWithDefaults instantiates a new NetworkingV1NetworkLinkEndpointUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1NetworkLinkEndpointUpdateWithDefaults() *NetworkingV1NetworkLinkEndpointUpdate {
	this := NetworkingV1NetworkLinkEndpointUpdate{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkEndpointUpdate) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkEndpointUpdate) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkEndpointUpdate) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *NetworkingV1NetworkLinkEndpointUpdate) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkEndpointUpdate) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkEndpointUpdate) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkEndpointUpdate) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *NetworkingV1NetworkLinkEndpointUpdate) SetKind(v string) {
	o.Kind = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkEndpointUpdate) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkEndpointUpdate) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkEndpointUpdate) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NetworkingV1NetworkLinkEndpointUpdate) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkEndpointUpdate) GetMetadata() ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkEndpointUpdate) GetMetadataOk() (*ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkEndpointUpdate) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ObjectMeta and assigns it to the Metadata field.
func (o *NetworkingV1NetworkLinkEndpointUpdate) SetMetadata(v ObjectMeta) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkEndpointUpdate) GetSpec() NetworkingV1NetworkLinkEndpointSpecUpdate {
	if o == nil || o.Spec == nil {
		var ret NetworkingV1NetworkLinkEndpointSpecUpdate
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkEndpointUpdate) GetSpecOk() (*NetworkingV1NetworkLinkEndpointSpecUpdate, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkEndpointUpdate) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given NetworkingV1NetworkLinkEndpointSpecUpdate and assigns it to the Spec field.
func (o *NetworkingV1NetworkLinkEndpointUpdate) SetSpec(v NetworkingV1NetworkLinkEndpointSpecUpdate) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkEndpointUpdate) GetStatus() NetworkingV1NetworkLinkEndpointStatus {
	if o == nil || o.Status == nil {
		var ret NetworkingV1NetworkLinkEndpointStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkEndpointUpdate) GetStatusOk() (*NetworkingV1NetworkLinkEndpointStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkEndpointUpdate) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NetworkingV1NetworkLinkEndpointStatus and assigns it to the Status field.
func (o *NetworkingV1NetworkLinkEndpointUpdate) SetStatus(v NetworkingV1NetworkLinkEndpointStatus) {
	o.Status = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1NetworkLinkEndpointUpdate) Redact() {
	o.recurseRedact(o.ApiVersion)
	o.recurseRedact(o.Kind)
	o.recurseRedact(o.Id)
	o.recurseRedact(o.Metadata)
	o.recurseRedact(o.Spec)
	o.recurseRedact(o.Status)
}

func (o *NetworkingV1NetworkLinkEndpointUpdate) recurseRedact(v interface{}) {
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

func (o NetworkingV1NetworkLinkEndpointUpdate) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1NetworkLinkEndpointUpdate) MarshalJSON() ([]byte, error) {
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
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableNetworkingV1NetworkLinkEndpointUpdate struct {
	value *NetworkingV1NetworkLinkEndpointUpdate
	isSet bool
}

func (v NullableNetworkingV1NetworkLinkEndpointUpdate) Get() *NetworkingV1NetworkLinkEndpointUpdate {
	return v.value
}

func (v *NullableNetworkingV1NetworkLinkEndpointUpdate) Set(val *NetworkingV1NetworkLinkEndpointUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1NetworkLinkEndpointUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1NetworkLinkEndpointUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1NetworkLinkEndpointUpdate(val *NetworkingV1NetworkLinkEndpointUpdate) *NullableNetworkingV1NetworkLinkEndpointUpdate {
	return &NullableNetworkingV1NetworkLinkEndpointUpdate{value: val, isSet: true}
}

func (v NullableNetworkingV1NetworkLinkEndpointUpdate) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableNetworkingV1NetworkLinkEndpointUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
