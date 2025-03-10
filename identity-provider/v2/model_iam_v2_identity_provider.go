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

// IamV2IdentityProvider `IdentityProvider` objects represent external OAuth-OIDC providers in Confluent Cloud.  The API allows you to list, create, read, update, and delete your Identity Provider.   Related guide: [OAuth for Confluent Cloud](https://docs.confluent.io/cloud/current/access-management/authenticate/oauth/overview.html).  ## The Identity Providers Model <SchemaDefinition schemaRef=\"#/components/schemas/iam.v2.IdentityProvider\" />  ## Quotas and Limits This resource is subject to the [following quotas](https://docs.confluent.io/cloud/current/quotas/overview.html):  | Quota | Description | | --- | --- | | `identity_providers_per_org` | Number of OAuth identity providers per organization | | `public_keys_per_provider` | Number of public keys saved per identity provider |
type IamV2IdentityProvider struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind *string `json:"kind,omitempty"`
	// ID is the \"natural identifier\" for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\"time\"); however, it may collide with IDs for other object `kinds` or objects of the same `kind` within a different scope/namespace (\"space\").
	Id       *string     `json:"id,omitempty"`
	Metadata *ObjectMeta `json:"metadata,omitempty"`
	// The human-readable name of the OAuth identity provider.
	DisplayName *string `json:"display_name,omitempty"`
	// A description of the identity provider.
	Description *string `json:"description,omitempty"`
	// The JSON Web Token (JWT) claim to extract the authenticating identity to Confluent resources from [Registered Claim Names](https://datatracker.ietf.org/doc/html/rfc7519#section-4.1). This appears in audit log records. Note: if the client specifies mapping to one identity pool ID, the identity claim configured with that pool will be used instead.
	IdentityClaim *string `json:"identity_claim,omitempty"`
	// The current state of the identity provider.
	State *string `json:"state,omitempty"`
	// A publicly accessible URL uniquely identifying the OAuth identity provider authorized to issue access tokens.
	Issuer *string `json:"issuer,omitempty"`
	// A publicly accessible JSON Web Key Set (JWKS) URI for the OAuth identity provider. JWKS provides a set of crypotgraphic keys used to verify the authenticity and integrity of JSON Web Tokens (JWTs) issued by the OAuth identity provider.
	JwksUri *string `json:"jwks_uri,omitempty"`
	// The JWKS issued by the OAuth identity provider. Only `kid` (key ID) and `alg` (algorithm) properties for each key set are included.
	Keys *[]IamV2JwksObject `json:"keys,omitempty"`
}

// NewIamV2IdentityProvider instantiates a new IamV2IdentityProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamV2IdentityProvider() *IamV2IdentityProvider {
	this := IamV2IdentityProvider{}
	return &this
}

// NewIamV2IdentityProviderWithDefaults instantiates a new IamV2IdentityProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamV2IdentityProviderWithDefaults() *IamV2IdentityProvider {
	this := IamV2IdentityProvider{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *IamV2IdentityProvider) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2IdentityProvider) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *IamV2IdentityProvider) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *IamV2IdentityProvider) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *IamV2IdentityProvider) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2IdentityProvider) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *IamV2IdentityProvider) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *IamV2IdentityProvider) SetKind(v string) {
	o.Kind = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IamV2IdentityProvider) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2IdentityProvider) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IamV2IdentityProvider) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IamV2IdentityProvider) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *IamV2IdentityProvider) GetMetadata() ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2IdentityProvider) GetMetadataOk() (*ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *IamV2IdentityProvider) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ObjectMeta and assigns it to the Metadata field.
func (o *IamV2IdentityProvider) SetMetadata(v ObjectMeta) {
	o.Metadata = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *IamV2IdentityProvider) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2IdentityProvider) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *IamV2IdentityProvider) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *IamV2IdentityProvider) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IamV2IdentityProvider) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2IdentityProvider) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IamV2IdentityProvider) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IamV2IdentityProvider) SetDescription(v string) {
	o.Description = &v
}

// GetIdentityClaim returns the IdentityClaim field value if set, zero value otherwise.
func (o *IamV2IdentityProvider) GetIdentityClaim() string {
	if o == nil || o.IdentityClaim == nil {
		var ret string
		return ret
	}
	return *o.IdentityClaim
}

// GetIdentityClaimOk returns a tuple with the IdentityClaim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2IdentityProvider) GetIdentityClaimOk() (*string, bool) {
	if o == nil || o.IdentityClaim == nil {
		return nil, false
	}
	return o.IdentityClaim, true
}

// HasIdentityClaim returns a boolean if a field has been set.
func (o *IamV2IdentityProvider) HasIdentityClaim() bool {
	if o != nil && o.IdentityClaim != nil {
		return true
	}

	return false
}

// SetIdentityClaim gets a reference to the given string and assigns it to the IdentityClaim field.
func (o *IamV2IdentityProvider) SetIdentityClaim(v string) {
	o.IdentityClaim = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *IamV2IdentityProvider) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2IdentityProvider) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *IamV2IdentityProvider) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *IamV2IdentityProvider) SetState(v string) {
	o.State = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *IamV2IdentityProvider) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2IdentityProvider) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *IamV2IdentityProvider) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *IamV2IdentityProvider) SetIssuer(v string) {
	o.Issuer = &v
}

// GetJwksUri returns the JwksUri field value if set, zero value otherwise.
func (o *IamV2IdentityProvider) GetJwksUri() string {
	if o == nil || o.JwksUri == nil {
		var ret string
		return ret
	}
	return *o.JwksUri
}

// GetJwksUriOk returns a tuple with the JwksUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2IdentityProvider) GetJwksUriOk() (*string, bool) {
	if o == nil || o.JwksUri == nil {
		return nil, false
	}
	return o.JwksUri, true
}

// HasJwksUri returns a boolean if a field has been set.
func (o *IamV2IdentityProvider) HasJwksUri() bool {
	if o != nil && o.JwksUri != nil {
		return true
	}

	return false
}

// SetJwksUri gets a reference to the given string and assigns it to the JwksUri field.
func (o *IamV2IdentityProvider) SetJwksUri(v string) {
	o.JwksUri = &v
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *IamV2IdentityProvider) GetKeys() []IamV2JwksObject {
	if o == nil || o.Keys == nil {
		var ret []IamV2JwksObject
		return ret
	}
	return *o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2IdentityProvider) GetKeysOk() (*[]IamV2JwksObject, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *IamV2IdentityProvider) HasKeys() bool {
	if o != nil && o.Keys != nil {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []IamV2JwksObject and assigns it to the Keys field.
func (o *IamV2IdentityProvider) SetKeys(v []IamV2JwksObject) {
	o.Keys = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *IamV2IdentityProvider) Redact() {
	o.recurseRedact(o.ApiVersion)
	o.recurseRedact(o.Kind)
	o.recurseRedact(o.Id)
	o.recurseRedact(o.Metadata)
	o.recurseRedact(o.DisplayName)
	o.recurseRedact(o.Description)
	o.recurseRedact(o.IdentityClaim)
	o.recurseRedact(o.State)
	o.recurseRedact(o.Issuer)
	o.recurseRedact(o.JwksUri)
	o.recurseRedact(o.Keys)
}

func (o *IamV2IdentityProvider) recurseRedact(v interface{}) {
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

func (o IamV2IdentityProvider) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o IamV2IdentityProvider) MarshalJSON() ([]byte, error) {
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
	if o.IdentityClaim != nil {
		toSerialize["identity_claim"] = o.IdentityClaim
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.JwksUri != nil {
		toSerialize["jwks_uri"] = o.JwksUri
	}
	if o.Keys != nil {
		toSerialize["keys"] = o.Keys
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableIamV2IdentityProvider struct {
	value *IamV2IdentityProvider
	isSet bool
}

func (v NullableIamV2IdentityProvider) Get() *IamV2IdentityProvider {
	return v.value
}

func (v *NullableIamV2IdentityProvider) Set(val *IamV2IdentityProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableIamV2IdentityProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableIamV2IdentityProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamV2IdentityProvider(val *IamV2IdentityProvider) *NullableIamV2IdentityProvider {
	return &NullableIamV2IdentityProvider{value: val, isSet: true}
}

func (v NullableIamV2IdentityProvider) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableIamV2IdentityProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
