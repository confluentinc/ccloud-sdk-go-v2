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
Certificate Authority Management API

mTLS Public API spec

API version: 0.0.1-alpha1
Contact: cloud-authn-platform-team@confluent.io
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

// IamV2CertificateIdentityPool `Identitypool` objects represent workload identities in Confluent Cloud.  The API allows you to list, create, read, update, and delete your identity pools associated with Certificate Authorities   Related guide: [Manage Certificate Identity Pools for Granular Client Access Management](https://docs.confluent.io/cloud/current/access-management/authenticate/mtls/configure.html#step-2-create-certificate-identity-pools-for-granular-access-control).  ## The Certificate Identity Pools Model <SchemaDefinition schemaRef=\"#/components/schemas/iam.v2.CertificateIdentityPool\" />  ## Quotas and Limits This resource is subject to the [following quotas](https://docs.confluent.io/cloud/current/quotas/overview.html):  | Quota | Description | | --- | --- | | `identity_pools_per_certificate_authority` | Number of Identity Pools per Certificate Authority |
type IamV2CertificateIdentityPool struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind *string `json:"kind,omitempty"`
	// ID is the \"natural identifier\" for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\"time\"); however, it may collide with IDs for other object `kinds` or objects of the same `kind` within a different scope/namespace (\"space\").
	Id       *string     `json:"id,omitempty"`
	Metadata *ObjectMeta `json:"metadata,omitempty"`
	// The name of the `IdentityPool`.
	DisplayName *string `json:"display_name,omitempty"`
	// A description of how this `IdentityPool` is used
	Description *string `json:"description,omitempty"`
	// The certificate field that will be used to represent the pool's external identifier for audit logging.
	ExternalIdentifier *string `json:"external_identifier,omitempty"`
	// A filter expression in [Supported Common Expression Language (CEL)](https://docs.confluent.io/cloud/current/access-management/authenticate/mtls/cel-filters.html) that specifies which identities can authenticate using your certificate identity pool (see [CEL filter for mTLS](https://docs.confluent.io/cloud/current/access-management/authenticate/mtls/cel-filters.html) for more details).
	Filter *string `json:"filter,omitempty"`
	// Represents the federated identity associated with this pool.
	Principal *string `json:"principal,omitempty"`
	// The current state of the identity pool
	State *string `json:"state,omitempty"`
}

// NewIamV2CertificateIdentityPool instantiates a new IamV2CertificateIdentityPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamV2CertificateIdentityPool() *IamV2CertificateIdentityPool {
	this := IamV2CertificateIdentityPool{}
	return &this
}

// NewIamV2CertificateIdentityPoolWithDefaults instantiates a new IamV2CertificateIdentityPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamV2CertificateIdentityPoolWithDefaults() *IamV2CertificateIdentityPool {
	this := IamV2CertificateIdentityPool{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *IamV2CertificateIdentityPool) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2CertificateIdentityPool) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *IamV2CertificateIdentityPool) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *IamV2CertificateIdentityPool) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *IamV2CertificateIdentityPool) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2CertificateIdentityPool) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *IamV2CertificateIdentityPool) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *IamV2CertificateIdentityPool) SetKind(v string) {
	o.Kind = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IamV2CertificateIdentityPool) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2CertificateIdentityPool) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IamV2CertificateIdentityPool) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IamV2CertificateIdentityPool) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *IamV2CertificateIdentityPool) GetMetadata() ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2CertificateIdentityPool) GetMetadataOk() (*ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *IamV2CertificateIdentityPool) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ObjectMeta and assigns it to the Metadata field.
func (o *IamV2CertificateIdentityPool) SetMetadata(v ObjectMeta) {
	o.Metadata = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *IamV2CertificateIdentityPool) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2CertificateIdentityPool) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *IamV2CertificateIdentityPool) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *IamV2CertificateIdentityPool) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IamV2CertificateIdentityPool) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2CertificateIdentityPool) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IamV2CertificateIdentityPool) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IamV2CertificateIdentityPool) SetDescription(v string) {
	o.Description = &v
}

// GetExternalIdentifier returns the ExternalIdentifier field value if set, zero value otherwise.
func (o *IamV2CertificateIdentityPool) GetExternalIdentifier() string {
	if o == nil || o.ExternalIdentifier == nil {
		var ret string
		return ret
	}
	return *o.ExternalIdentifier
}

// GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2CertificateIdentityPool) GetExternalIdentifierOk() (*string, bool) {
	if o == nil || o.ExternalIdentifier == nil {
		return nil, false
	}
	return o.ExternalIdentifier, true
}

// HasExternalIdentifier returns a boolean if a field has been set.
func (o *IamV2CertificateIdentityPool) HasExternalIdentifier() bool {
	if o != nil && o.ExternalIdentifier != nil {
		return true
	}

	return false
}

// SetExternalIdentifier gets a reference to the given string and assigns it to the ExternalIdentifier field.
func (o *IamV2CertificateIdentityPool) SetExternalIdentifier(v string) {
	o.ExternalIdentifier = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *IamV2CertificateIdentityPool) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2CertificateIdentityPool) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *IamV2CertificateIdentityPool) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *IamV2CertificateIdentityPool) SetFilter(v string) {
	o.Filter = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *IamV2CertificateIdentityPool) GetPrincipal() string {
	if o == nil || o.Principal == nil {
		var ret string
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2CertificateIdentityPool) GetPrincipalOk() (*string, bool) {
	if o == nil || o.Principal == nil {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *IamV2CertificateIdentityPool) HasPrincipal() bool {
	if o != nil && o.Principal != nil {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given string and assigns it to the Principal field.
func (o *IamV2CertificateIdentityPool) SetPrincipal(v string) {
	o.Principal = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *IamV2CertificateIdentityPool) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2CertificateIdentityPool) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *IamV2CertificateIdentityPool) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *IamV2CertificateIdentityPool) SetState(v string) {
	o.State = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *IamV2CertificateIdentityPool) Redact() {
	o.recurseRedact(o.ApiVersion)
	o.recurseRedact(o.Kind)
	o.recurseRedact(o.Id)
	o.recurseRedact(o.Metadata)
	o.recurseRedact(o.DisplayName)
	o.recurseRedact(o.Description)
	o.recurseRedact(o.ExternalIdentifier)
	o.recurseRedact(o.Filter)
	o.recurseRedact(o.Principal)
	o.recurseRedact(o.State)
}

func (o *IamV2CertificateIdentityPool) recurseRedact(v interface{}) {
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

func (o IamV2CertificateIdentityPool) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o IamV2CertificateIdentityPool) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ExternalIdentifier != nil {
		toSerialize["external_identifier"] = o.ExternalIdentifier
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Principal != nil {
		toSerialize["principal"] = o.Principal
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableIamV2CertificateIdentityPool struct {
	value *IamV2CertificateIdentityPool
	isSet bool
}

func (v NullableIamV2CertificateIdentityPool) Get() *IamV2CertificateIdentityPool {
	return v.value
}

func (v *NullableIamV2CertificateIdentityPool) Set(val *IamV2CertificateIdentityPool) {
	v.value = val
	v.isSet = true
}

func (v NullableIamV2CertificateIdentityPool) IsSet() bool {
	return v.isSet
}

func (v *NullableIamV2CertificateIdentityPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamV2CertificateIdentityPool(val *IamV2CertificateIdentityPool) *NullableIamV2CertificateIdentityPool {
	return &NullableIamV2CertificateIdentityPool{value: val, isSet: true}
}

func (v NullableIamV2CertificateIdentityPool) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableIamV2CertificateIdentityPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}