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
Service Quota API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1-alpha0
Contact: api-framework-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

import (
	"reflect"
)

// ServiceQuotaV1AppliedQuota A `quota` object represents a quota configuration for a specific Confluent Cloud resource. Use this API to retrieve an individual quota or list of quotas for a given scope.   Related guide: [Service Quotas for Confluent Cloud](https://docs.confluent.io/cloud/current/quotas/index.html).  ## The Applied Quotas Model <SchemaDefinition schemaRef=\"#/components/schemas/service-quota.v1.AppliedQuota\" />
type ServiceQuotaV1AppliedQuota struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind *string `json:"kind,omitempty"`
	// ID is the \"natural identifier\" for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\"time\"); however, it may collide with IDs for other object `kinds` or objects of the same `kind` within a different scope/namespace (\"space\").
	Id *string `json:"id,omitempty"`
	Metadata *ObjectMeta `json:"metadata,omitempty"`
	// The applied scope that this quota belongs to.
	Scope *string `json:"scope,omitempty"`
	// A human-readable name for the quota type name.
	DisplayName *string `json:"display_name,omitempty"`
	// The default service quota value. 
	DefaultLimit *int32 `json:"default_limit,omitempty"`
	// The latest applied service quota value, taking into account any limit adjustments. 
	AppliedLimit *int32 `json:"applied_limit,omitempty"`
	// Show the quota usage value if the quota usage is available for this quota. 
	Usage *int32 `json:"usage,omitempty"`
	// The user associated with this object.
	User *ObjectReference `json:"user,omitempty"`
	// A unique organization id to associate a specific organization to this quota. May be `null` if not associated with a organization.
	Organization *ObjectReference `json:"organization,omitempty"`
	// The environment ID the quota is associated with.  May be `null` if not associated with a environment.
	Environment *ObjectReference `json:"environment,omitempty"`
	// The network ID the quota is associated with.  May be `null` if not associated with a network.
	Network *ObjectReference `json:"network,omitempty"`
	// The kafka cluster ID the quota is associated with.  May be `null` if not associated with a kafka_cluster.
	KafkaCluster *ObjectReference `json:"kafka_cluster,omitempty"`
}

// NewServiceQuotaV1AppliedQuota instantiates a new ServiceQuotaV1AppliedQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceQuotaV1AppliedQuota() *ServiceQuotaV1AppliedQuota {
	this := ServiceQuotaV1AppliedQuota{}
	return &this
}

// NewServiceQuotaV1AppliedQuotaWithDefaults instantiates a new ServiceQuotaV1AppliedQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceQuotaV1AppliedQuotaWithDefaults() *ServiceQuotaV1AppliedQuota {
	this := ServiceQuotaV1AppliedQuota{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ServiceQuotaV1AppliedQuota) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceQuotaV1AppliedQuota) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ServiceQuotaV1AppliedQuota) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ServiceQuotaV1AppliedQuota) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ServiceQuotaV1AppliedQuota) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceQuotaV1AppliedQuota) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ServiceQuotaV1AppliedQuota) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ServiceQuotaV1AppliedQuota) SetKind(v string) {
	o.Kind = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceQuotaV1AppliedQuota) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceQuotaV1AppliedQuota) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceQuotaV1AppliedQuota) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceQuotaV1AppliedQuota) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ServiceQuotaV1AppliedQuota) GetMetadata() ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceQuotaV1AppliedQuota) GetMetadataOk() (*ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ServiceQuotaV1AppliedQuota) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ObjectMeta and assigns it to the Metadata field.
func (o *ServiceQuotaV1AppliedQuota) SetMetadata(v ObjectMeta) {
	o.Metadata = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ServiceQuotaV1AppliedQuota) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceQuotaV1AppliedQuota) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ServiceQuotaV1AppliedQuota) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *ServiceQuotaV1AppliedQuota) SetScope(v string) {
	o.Scope = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ServiceQuotaV1AppliedQuota) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceQuotaV1AppliedQuota) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ServiceQuotaV1AppliedQuota) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ServiceQuotaV1AppliedQuota) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDefaultLimit returns the DefaultLimit field value if set, zero value otherwise.
func (o *ServiceQuotaV1AppliedQuota) GetDefaultLimit() int32 {
	if o == nil || o.DefaultLimit == nil {
		var ret int32
		return ret
	}
	return *o.DefaultLimit
}

// GetDefaultLimitOk returns a tuple with the DefaultLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceQuotaV1AppliedQuota) GetDefaultLimitOk() (*int32, bool) {
	if o == nil || o.DefaultLimit == nil {
		return nil, false
	}
	return o.DefaultLimit, true
}

// HasDefaultLimit returns a boolean if a field has been set.
func (o *ServiceQuotaV1AppliedQuota) HasDefaultLimit() bool {
	if o != nil && o.DefaultLimit != nil {
		return true
	}

	return false
}

// SetDefaultLimit gets a reference to the given int32 and assigns it to the DefaultLimit field.
func (o *ServiceQuotaV1AppliedQuota) SetDefaultLimit(v int32) {
	o.DefaultLimit = &v
}

// GetAppliedLimit returns the AppliedLimit field value if set, zero value otherwise.
func (o *ServiceQuotaV1AppliedQuota) GetAppliedLimit() int32 {
	if o == nil || o.AppliedLimit == nil {
		var ret int32
		return ret
	}
	return *o.AppliedLimit
}

// GetAppliedLimitOk returns a tuple with the AppliedLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceQuotaV1AppliedQuota) GetAppliedLimitOk() (*int32, bool) {
	if o == nil || o.AppliedLimit == nil {
		return nil, false
	}
	return o.AppliedLimit, true
}

// HasAppliedLimit returns a boolean if a field has been set.
func (o *ServiceQuotaV1AppliedQuota) HasAppliedLimit() bool {
	if o != nil && o.AppliedLimit != nil {
		return true
	}

	return false
}

// SetAppliedLimit gets a reference to the given int32 and assigns it to the AppliedLimit field.
func (o *ServiceQuotaV1AppliedQuota) SetAppliedLimit(v int32) {
	o.AppliedLimit = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *ServiceQuotaV1AppliedQuota) GetUsage() int32 {
	if o == nil || o.Usage == nil {
		var ret int32
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceQuotaV1AppliedQuota) GetUsageOk() (*int32, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *ServiceQuotaV1AppliedQuota) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given int32 and assigns it to the Usage field.
func (o *ServiceQuotaV1AppliedQuota) SetUsage(v int32) {
	o.Usage = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ServiceQuotaV1AppliedQuota) GetUser() ObjectReference {
	if o == nil || o.User == nil {
		var ret ObjectReference
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceQuotaV1AppliedQuota) GetUserOk() (*ObjectReference, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ServiceQuotaV1AppliedQuota) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given ObjectReference and assigns it to the User field.
func (o *ServiceQuotaV1AppliedQuota) SetUser(v ObjectReference) {
	o.User = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ServiceQuotaV1AppliedQuota) GetOrganization() ObjectReference {
	if o == nil || o.Organization == nil {
		var ret ObjectReference
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceQuotaV1AppliedQuota) GetOrganizationOk() (*ObjectReference, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ServiceQuotaV1AppliedQuota) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given ObjectReference and assigns it to the Organization field.
func (o *ServiceQuotaV1AppliedQuota) SetOrganization(v ObjectReference) {
	o.Organization = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ServiceQuotaV1AppliedQuota) GetEnvironment() ObjectReference {
	if o == nil || o.Environment == nil {
		var ret ObjectReference
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceQuotaV1AppliedQuota) GetEnvironmentOk() (*ObjectReference, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ServiceQuotaV1AppliedQuota) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectReference and assigns it to the Environment field.
func (o *ServiceQuotaV1AppliedQuota) SetEnvironment(v ObjectReference) {
	o.Environment = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *ServiceQuotaV1AppliedQuota) GetNetwork() ObjectReference {
	if o == nil || o.Network == nil {
		var ret ObjectReference
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceQuotaV1AppliedQuota) GetNetworkOk() (*ObjectReference, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *ServiceQuotaV1AppliedQuota) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given ObjectReference and assigns it to the Network field.
func (o *ServiceQuotaV1AppliedQuota) SetNetwork(v ObjectReference) {
	o.Network = &v
}

// GetKafkaCluster returns the KafkaCluster field value if set, zero value otherwise.
func (o *ServiceQuotaV1AppliedQuota) GetKafkaCluster() ObjectReference {
	if o == nil || o.KafkaCluster == nil {
		var ret ObjectReference
		return ret
	}
	return *o.KafkaCluster
}

// GetKafkaClusterOk returns a tuple with the KafkaCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceQuotaV1AppliedQuota) GetKafkaClusterOk() (*ObjectReference, bool) {
	if o == nil || o.KafkaCluster == nil {
		return nil, false
	}
	return o.KafkaCluster, true
}

// HasKafkaCluster returns a boolean if a field has been set.
func (o *ServiceQuotaV1AppliedQuota) HasKafkaCluster() bool {
	if o != nil && o.KafkaCluster != nil {
		return true
	}

	return false
}

// SetKafkaCluster gets a reference to the given ObjectReference and assigns it to the KafkaCluster field.
func (o *ServiceQuotaV1AppliedQuota) SetKafkaCluster(v ObjectReference) {
	o.KafkaCluster = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *ServiceQuotaV1AppliedQuota) Redact() {
    o.recurseRedact(o.ApiVersion)
    o.recurseRedact(o.Kind)
    o.recurseRedact(o.Id)
    o.recurseRedact(o.Metadata)
    o.recurseRedact(o.Scope)
    o.recurseRedact(o.DisplayName)
    o.recurseRedact(o.DefaultLimit)
    o.recurseRedact(o.AppliedLimit)
    o.recurseRedact(o.Usage)
    o.recurseRedact(o.User)
    o.recurseRedact(o.Organization)
    o.recurseRedact(o.Environment)
    o.recurseRedact(o.Network)
    o.recurseRedact(o.KafkaCluster)
}

func (o *ServiceQuotaV1AppliedQuota) recurseRedact(v interface{}) {
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

func (o ServiceQuotaV1AppliedQuota) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o ServiceQuotaV1AppliedQuota) MarshalJSON() ([]byte, error) {
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
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.DefaultLimit != nil {
		toSerialize["default_limit"] = o.DefaultLimit
	}
	if o.AppliedLimit != nil {
		toSerialize["applied_limit"] = o.AppliedLimit
	}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.KafkaCluster != nil {
		toSerialize["kafka_cluster"] = o.KafkaCluster
	}
	return json.Marshal(toSerialize)
}

type NullableServiceQuotaV1AppliedQuota struct {
	value *ServiceQuotaV1AppliedQuota
	isSet bool
}

func (v NullableServiceQuotaV1AppliedQuota) Get() *ServiceQuotaV1AppliedQuota {
	return v.value
}

func (v *NullableServiceQuotaV1AppliedQuota) Set(val *ServiceQuotaV1AppliedQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceQuotaV1AppliedQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceQuotaV1AppliedQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceQuotaV1AppliedQuota(val *ServiceQuotaV1AppliedQuota) *NullableServiceQuotaV1AppliedQuota {
	return &NullableServiceQuotaV1AppliedQuota{value: val, isSet: true}
}

func (v NullableServiceQuotaV1AppliedQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceQuotaV1AppliedQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


