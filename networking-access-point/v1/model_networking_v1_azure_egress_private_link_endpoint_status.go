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

// NetworkingV1AzureEgressPrivateLinkEndpointStatus Status of an Azure Private Endpoint.
type NetworkingV1AzureEgressPrivateLinkEndpointStatus struct {
	// AzureEgressPrivateLinkEndpointStatus kind.
	Kind string `json:"kind,omitempty"`
	// Resource ID of the Private Endpoint (if any) that is connected to the Private Link service.
	PrivateEndpointResourceId string `json:"private_endpoint_resource_id,omitempty"`
	// Domain of the Private Endpoint (if any) that is connected to the Private Link service.
	PrivateEndpointDomain *string `json:"private_endpoint_domain,omitempty"`
	// IP address of the Private Endpoint (if any) that is connected to the Private Link service.
	PrivateEndpointIpAddress string `json:"private_endpoint_ip_address,omitempty"`
	// Domains of the Private Endpoint (if any) based off FQDNs in Azure custom DNS configs, which are required in your private DNS setup.
	PrivateEndpointCustomDnsConfigDomains *[]string `json:"private_endpoint_custom_dns_config_domains,omitempty"`
}

// NewNetworkingV1AzureEgressPrivateLinkEndpointStatus instantiates a new NetworkingV1AzureEgressPrivateLinkEndpointStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1AzureEgressPrivateLinkEndpointStatus(kind string, privateEndpointResourceId string, privateEndpointIpAddress string) *NetworkingV1AzureEgressPrivateLinkEndpointStatus {
	this := NetworkingV1AzureEgressPrivateLinkEndpointStatus{}
	this.Kind = kind
	this.PrivateEndpointResourceId = privateEndpointResourceId
	this.PrivateEndpointIpAddress = privateEndpointIpAddress
	return &this
}

// NewNetworkingV1AzureEgressPrivateLinkEndpointStatusWithDefaults instantiates a new NetworkingV1AzureEgressPrivateLinkEndpointStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1AzureEgressPrivateLinkEndpointStatusWithDefaults() *NetworkingV1AzureEgressPrivateLinkEndpointStatus {
	this := NetworkingV1AzureEgressPrivateLinkEndpointStatus{}
	return &this
}

// GetKind returns the Kind field value
func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) SetKind(v string) {
	o.Kind = v
}

// GetPrivateEndpointResourceId returns the PrivateEndpointResourceId field value
func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) GetPrivateEndpointResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateEndpointResourceId
}

// GetPrivateEndpointResourceIdOk returns a tuple with the PrivateEndpointResourceId field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) GetPrivateEndpointResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateEndpointResourceId, true
}

// SetPrivateEndpointResourceId sets field value
func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) SetPrivateEndpointResourceId(v string) {
	o.PrivateEndpointResourceId = v
}

// GetPrivateEndpointDomain returns the PrivateEndpointDomain field value if set, zero value otherwise.
func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) GetPrivateEndpointDomain() string {
	if o == nil || o.PrivateEndpointDomain == nil {
		var ret string
		return ret
	}
	return *o.PrivateEndpointDomain
}

// GetPrivateEndpointDomainOk returns a tuple with the PrivateEndpointDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) GetPrivateEndpointDomainOk() (*string, bool) {
	if o == nil || o.PrivateEndpointDomain == nil {
		return nil, false
	}
	return o.PrivateEndpointDomain, true
}

// HasPrivateEndpointDomain returns a boolean if a field has been set.
func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) HasPrivateEndpointDomain() bool {
	if o != nil && o.PrivateEndpointDomain != nil {
		return true
	}

	return false
}

// SetPrivateEndpointDomain gets a reference to the given string and assigns it to the PrivateEndpointDomain field.
func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) SetPrivateEndpointDomain(v string) {
	o.PrivateEndpointDomain = &v
}

// GetPrivateEndpointIpAddress returns the PrivateEndpointIpAddress field value
func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) GetPrivateEndpointIpAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateEndpointIpAddress
}

// GetPrivateEndpointIpAddressOk returns a tuple with the PrivateEndpointIpAddress field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) GetPrivateEndpointIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateEndpointIpAddress, true
}

// SetPrivateEndpointIpAddress sets field value
func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) SetPrivateEndpointIpAddress(v string) {
	o.PrivateEndpointIpAddress = v
}

// GetPrivateEndpointCustomDnsConfigDomains returns the PrivateEndpointCustomDnsConfigDomains field value if set, zero value otherwise.
func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) GetPrivateEndpointCustomDnsConfigDomains() []string {
	if o == nil || o.PrivateEndpointCustomDnsConfigDomains == nil {
		var ret []string
		return ret
	}
	return *o.PrivateEndpointCustomDnsConfigDomains
}

// GetPrivateEndpointCustomDnsConfigDomainsOk returns a tuple with the PrivateEndpointCustomDnsConfigDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) GetPrivateEndpointCustomDnsConfigDomainsOk() (*[]string, bool) {
	if o == nil || o.PrivateEndpointCustomDnsConfigDomains == nil {
		return nil, false
	}
	return o.PrivateEndpointCustomDnsConfigDomains, true
}

// HasPrivateEndpointCustomDnsConfigDomains returns a boolean if a field has been set.
func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) HasPrivateEndpointCustomDnsConfigDomains() bool {
	if o != nil && o.PrivateEndpointCustomDnsConfigDomains != nil {
		return true
	}

	return false
}

// SetPrivateEndpointCustomDnsConfigDomains gets a reference to the given []string and assigns it to the PrivateEndpointCustomDnsConfigDomains field.
func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) SetPrivateEndpointCustomDnsConfigDomains(v []string) {
	o.PrivateEndpointCustomDnsConfigDomains = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) Redact() {
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.PrivateEndpointResourceId)
	o.recurseRedact(o.PrivateEndpointDomain)
	o.recurseRedact(&o.PrivateEndpointIpAddress)
	o.recurseRedact(o.PrivateEndpointCustomDnsConfigDomains)
}

func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) recurseRedact(v interface{}) {
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

func (o NetworkingV1AzureEgressPrivateLinkEndpointStatus) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1AzureEgressPrivateLinkEndpointStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["private_endpoint_resource_id"] = o.PrivateEndpointResourceId
	}
	if o.PrivateEndpointDomain != nil {
		toSerialize["private_endpoint_domain"] = o.PrivateEndpointDomain
	}
	if true {
		toSerialize["private_endpoint_ip_address"] = o.PrivateEndpointIpAddress
	}
	if o.PrivateEndpointCustomDnsConfigDomains != nil {
		toSerialize["private_endpoint_custom_dns_config_domains"] = o.PrivateEndpointCustomDnsConfigDomains
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableNetworkingV1AzureEgressPrivateLinkEndpointStatus struct {
	value *NetworkingV1AzureEgressPrivateLinkEndpointStatus
	isSet bool
}

func (v NullableNetworkingV1AzureEgressPrivateLinkEndpointStatus) Get() *NetworkingV1AzureEgressPrivateLinkEndpointStatus {
	return v.value
}

func (v *NullableNetworkingV1AzureEgressPrivateLinkEndpointStatus) Set(val *NetworkingV1AzureEgressPrivateLinkEndpointStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1AzureEgressPrivateLinkEndpointStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1AzureEgressPrivateLinkEndpointStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1AzureEgressPrivateLinkEndpointStatus(val *NetworkingV1AzureEgressPrivateLinkEndpointStatus) *NullableNetworkingV1AzureEgressPrivateLinkEndpointStatus {
	return &NullableNetworkingV1AzureEgressPrivateLinkEndpointStatus{value: val, isSet: true}
}

func (v NullableNetworkingV1AzureEgressPrivateLinkEndpointStatus) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableNetworkingV1AzureEgressPrivateLinkEndpointStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
