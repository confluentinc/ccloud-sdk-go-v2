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

// NetworkingV1IpAddress IP Addresses  Related guide: [Use Public Egress IP addresses on Confluent Cloud](https://docs.confluent.io/cloud/current/networking/static-egress-ip-addresses.html)  ## The IP Addresses Model <SchemaDefinition schemaRef=\"#/components/schemas/networking.v1.IpAddress\" />
type NetworkingV1IpAddress struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind *string `json:"kind,omitempty"`
	// The IP Address range.
	IpPrefix *string `json:"ip_prefix,omitempty"`
	// The cloud service provider in which the address exists.
	Cloud *string `json:"cloud,omitempty"`
	// The region/location where the IP Address is in use.
	Region   *string `json:"region,omitempty"`
	Services *Set    `json:"services,omitempty"`
	// Whether the address is used for egress or ingress.
	AddressType *string `json:"address_type,omitempty"`
}

// NewNetworkingV1IpAddress instantiates a new NetworkingV1IpAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1IpAddress() *NetworkingV1IpAddress {
	this := NetworkingV1IpAddress{}
	return &this
}

// NewNetworkingV1IpAddressWithDefaults instantiates a new NetworkingV1IpAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1IpAddressWithDefaults() *NetworkingV1IpAddress {
	this := NetworkingV1IpAddress{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *NetworkingV1IpAddress) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1IpAddress) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *NetworkingV1IpAddress) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *NetworkingV1IpAddress) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *NetworkingV1IpAddress) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1IpAddress) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *NetworkingV1IpAddress) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *NetworkingV1IpAddress) SetKind(v string) {
	o.Kind = &v
}

// GetIpPrefix returns the IpPrefix field value if set, zero value otherwise.
func (o *NetworkingV1IpAddress) GetIpPrefix() string {
	if o == nil || o.IpPrefix == nil {
		var ret string
		return ret
	}
	return *o.IpPrefix
}

// GetIpPrefixOk returns a tuple with the IpPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1IpAddress) GetIpPrefixOk() (*string, bool) {
	if o == nil || o.IpPrefix == nil {
		return nil, false
	}
	return o.IpPrefix, true
}

// HasIpPrefix returns a boolean if a field has been set.
func (o *NetworkingV1IpAddress) HasIpPrefix() bool {
	if o != nil && o.IpPrefix != nil {
		return true
	}

	return false
}

// SetIpPrefix gets a reference to the given string and assigns it to the IpPrefix field.
func (o *NetworkingV1IpAddress) SetIpPrefix(v string) {
	o.IpPrefix = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *NetworkingV1IpAddress) GetCloud() string {
	if o == nil || o.Cloud == nil {
		var ret string
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1IpAddress) GetCloudOk() (*string, bool) {
	if o == nil || o.Cloud == nil {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *NetworkingV1IpAddress) HasCloud() bool {
	if o != nil && o.Cloud != nil {
		return true
	}

	return false
}

// SetCloud gets a reference to the given string and assigns it to the Cloud field.
func (o *NetworkingV1IpAddress) SetCloud(v string) {
	o.Cloud = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *NetworkingV1IpAddress) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1IpAddress) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *NetworkingV1IpAddress) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *NetworkingV1IpAddress) SetRegion(v string) {
	o.Region = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *NetworkingV1IpAddress) GetServices() Set {
	if o == nil || o.Services == nil {
		var ret Set
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1IpAddress) GetServicesOk() (*Set, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *NetworkingV1IpAddress) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given Set and assigns it to the Services field.
func (o *NetworkingV1IpAddress) SetServices(v Set) {
	o.Services = &v
}

// GetAddressType returns the AddressType field value if set, zero value otherwise.
func (o *NetworkingV1IpAddress) GetAddressType() string {
	if o == nil || o.AddressType == nil {
		var ret string
		return ret
	}
	return *o.AddressType
}

// GetAddressTypeOk returns a tuple with the AddressType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1IpAddress) GetAddressTypeOk() (*string, bool) {
	if o == nil || o.AddressType == nil {
		return nil, false
	}
	return o.AddressType, true
}

// HasAddressType returns a boolean if a field has been set.
func (o *NetworkingV1IpAddress) HasAddressType() bool {
	if o != nil && o.AddressType != nil {
		return true
	}

	return false
}

// SetAddressType gets a reference to the given string and assigns it to the AddressType field.
func (o *NetworkingV1IpAddress) SetAddressType(v string) {
	o.AddressType = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1IpAddress) Redact() {
	o.recurseRedact(o.ApiVersion)
	o.recurseRedact(o.Kind)
	o.recurseRedact(o.IpPrefix)
	o.recurseRedact(o.Cloud)
	o.recurseRedact(o.Region)
	o.recurseRedact(o.Services)
	o.recurseRedact(o.AddressType)
}

func (o *NetworkingV1IpAddress) recurseRedact(v interface{}) {
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

func (o NetworkingV1IpAddress) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1IpAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api_version"] = o.ApiVersion
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.IpPrefix != nil {
		toSerialize["ip_prefix"] = o.IpPrefix
	}
	if o.Cloud != nil {
		toSerialize["cloud"] = o.Cloud
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	if o.AddressType != nil {
		toSerialize["address_type"] = o.AddressType
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableNetworkingV1IpAddress struct {
	value *NetworkingV1IpAddress
	isSet bool
}

func (v NullableNetworkingV1IpAddress) Get() *NetworkingV1IpAddress {
	return v.value
}

func (v *NullableNetworkingV1IpAddress) Set(val *NetworkingV1IpAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1IpAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1IpAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1IpAddress(val *NetworkingV1IpAddress) *NullableNetworkingV1IpAddress {
	return &NullableNetworkingV1IpAddress{value: val, isSet: true}
}

func (v NullableNetworkingV1IpAddress) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableNetworkingV1IpAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
