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
Flink Cluster Management API

This is the new Flink cluster management API.

API version: 0.0.1
Contact: ksql-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

import (
	"reflect"
)

// FlinkcmV2ClusterSpec The desired state of the Cluster
type FlinkcmV2ClusterSpec struct {
	// The name of the Flink cluster.
	DisplayName *string `json:"display_name,omitempty"`
	// The cloud provider that the Flink cluster is in.
	CloudProvider *string `json:"cloud_provider,omitempty"`
	// The region that the Flink cluster is in.
	Region *string `json:"region,omitempty"`
	// The number of CSUs (Confluent Streaming Units) in a Flink cluster.
	Csu *int32 `json:"csu,omitempty"`
	// The Type of network to use.
	NetworkType *string `json:"network_type,omitempty"`
	// The principal to which this belongs. The principal can be one of iam.v2.User, iam.v2.ServiceAccount.
	Principal *TypedGlobalObjectReference `json:"principal,omitempty"`
	// The environment to which this belongs.
	Environment *GlobalObjectReference `json:"environment,omitempty"`
	// The network to which this belongs.
	Network *ObjectReference `json:"network,omitempty"`
}

// NewFlinkcmV2ClusterSpec instantiates a new FlinkcmV2ClusterSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlinkcmV2ClusterSpec() *FlinkcmV2ClusterSpec {
	this := FlinkcmV2ClusterSpec{}
	return &this
}

// NewFlinkcmV2ClusterSpecWithDefaults instantiates a new FlinkcmV2ClusterSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlinkcmV2ClusterSpecWithDefaults() *FlinkcmV2ClusterSpec {
	this := FlinkcmV2ClusterSpec{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *FlinkcmV2ClusterSpec) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkcmV2ClusterSpec) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *FlinkcmV2ClusterSpec) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *FlinkcmV2ClusterSpec) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *FlinkcmV2ClusterSpec) GetCloudProvider() string {
	if o == nil || o.CloudProvider == nil {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkcmV2ClusterSpec) GetCloudProviderOk() (*string, bool) {
	if o == nil || o.CloudProvider == nil {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *FlinkcmV2ClusterSpec) HasCloudProvider() bool {
	if o != nil && o.CloudProvider != nil {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *FlinkcmV2ClusterSpec) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *FlinkcmV2ClusterSpec) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkcmV2ClusterSpec) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *FlinkcmV2ClusterSpec) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *FlinkcmV2ClusterSpec) SetRegion(v string) {
	o.Region = &v
}

// GetCsu returns the Csu field value if set, zero value otherwise.
func (o *FlinkcmV2ClusterSpec) GetCsu() int32 {
	if o == nil || o.Csu == nil {
		var ret int32
		return ret
	}
	return *o.Csu
}

// GetCsuOk returns a tuple with the Csu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkcmV2ClusterSpec) GetCsuOk() (*int32, bool) {
	if o == nil || o.Csu == nil {
		return nil, false
	}
	return o.Csu, true
}

// HasCsu returns a boolean if a field has been set.
func (o *FlinkcmV2ClusterSpec) HasCsu() bool {
	if o != nil && o.Csu != nil {
		return true
	}

	return false
}

// SetCsu gets a reference to the given int32 and assigns it to the Csu field.
func (o *FlinkcmV2ClusterSpec) SetCsu(v int32) {
	o.Csu = &v
}

// GetNetworkType returns the NetworkType field value if set, zero value otherwise.
func (o *FlinkcmV2ClusterSpec) GetNetworkType() string {
	if o == nil || o.NetworkType == nil {
		var ret string
		return ret
	}
	return *o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkcmV2ClusterSpec) GetNetworkTypeOk() (*string, bool) {
	if o == nil || o.NetworkType == nil {
		return nil, false
	}
	return o.NetworkType, true
}

// HasNetworkType returns a boolean if a field has been set.
func (o *FlinkcmV2ClusterSpec) HasNetworkType() bool {
	if o != nil && o.NetworkType != nil {
		return true
	}

	return false
}

// SetNetworkType gets a reference to the given string and assigns it to the NetworkType field.
func (o *FlinkcmV2ClusterSpec) SetNetworkType(v string) {
	o.NetworkType = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *FlinkcmV2ClusterSpec) GetPrincipal() TypedGlobalObjectReference {
	if o == nil || o.Principal == nil {
		var ret TypedGlobalObjectReference
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkcmV2ClusterSpec) GetPrincipalOk() (*TypedGlobalObjectReference, bool) {
	if o == nil || o.Principal == nil {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *FlinkcmV2ClusterSpec) HasPrincipal() bool {
	if o != nil && o.Principal != nil {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given TypedGlobalObjectReference and assigns it to the Principal field.
func (o *FlinkcmV2ClusterSpec) SetPrincipal(v TypedGlobalObjectReference) {
	o.Principal = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *FlinkcmV2ClusterSpec) GetEnvironment() GlobalObjectReference {
	if o == nil || o.Environment == nil {
		var ret GlobalObjectReference
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkcmV2ClusterSpec) GetEnvironmentOk() (*GlobalObjectReference, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *FlinkcmV2ClusterSpec) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given GlobalObjectReference and assigns it to the Environment field.
func (o *FlinkcmV2ClusterSpec) SetEnvironment(v GlobalObjectReference) {
	o.Environment = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *FlinkcmV2ClusterSpec) GetNetwork() ObjectReference {
	if o == nil || o.Network == nil {
		var ret ObjectReference
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkcmV2ClusterSpec) GetNetworkOk() (*ObjectReference, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *FlinkcmV2ClusterSpec) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given ObjectReference and assigns it to the Network field.
func (o *FlinkcmV2ClusterSpec) SetNetwork(v ObjectReference) {
	o.Network = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *FlinkcmV2ClusterSpec) Redact() {
    o.recurseRedact(o.DisplayName)
    o.recurseRedact(o.CloudProvider)
    o.recurseRedact(o.Region)
    o.recurseRedact(o.Csu)
    o.recurseRedact(o.NetworkType)
    o.recurseRedact(o.Principal)
    o.recurseRedact(o.Environment)
    o.recurseRedact(o.Network)
}

func (o *FlinkcmV2ClusterSpec) recurseRedact(v interface{}) {
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

func (o FlinkcmV2ClusterSpec) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o FlinkcmV2ClusterSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.CloudProvider != nil {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Csu != nil {
		toSerialize["csu"] = o.Csu
	}
	if o.NetworkType != nil {
		toSerialize["network_type"] = o.NetworkType
	}
	if o.Principal != nil {
		toSerialize["principal"] = o.Principal
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	return json.Marshal(toSerialize)
}

type NullableFlinkcmV2ClusterSpec struct {
	value *FlinkcmV2ClusterSpec
	isSet bool
}

func (v NullableFlinkcmV2ClusterSpec) Get() *FlinkcmV2ClusterSpec {
	return v.value
}

func (v *NullableFlinkcmV2ClusterSpec) Set(val *FlinkcmV2ClusterSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableFlinkcmV2ClusterSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableFlinkcmV2ClusterSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlinkcmV2ClusterSpec(val *FlinkcmV2ClusterSpec) *NullableFlinkcmV2ClusterSpec {
	return &NullableFlinkcmV2ClusterSpec{value: val, isSet: true}
}

func (v NullableFlinkcmV2ClusterSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlinkcmV2ClusterSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

