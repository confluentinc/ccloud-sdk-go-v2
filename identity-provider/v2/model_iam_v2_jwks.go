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

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1-alpha1
Contact: oauth-eng@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

import (
	"reflect"
)

// IamV2Jwks `JWKS` contains kid and alg for the published keys for the given OpenIDProvider
type IamV2Jwks struct {
	// Specifies the cryptographic algorithm family used with the key
	Kty string `json:"kty"`
	// Specifies the key-id issued by the OpenIDProvider for the particular tenant
	Kid string `json:"kid"`
	// Specifies the algorithm to be used to generate the public key
	Alg string `json:"alg"`
	// Specifies the intended use of the public key
	Use *string `json:"use,omitempty"`
	// Identifies the operation(s) for which the key is intended to be used
	KeyOps *string `json:"key_ops,omitempty"`
}

// NewIamV2Jwks instantiates a new IamV2Jwks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamV2Jwks(kty string, kid string, alg string) *IamV2Jwks {
	this := IamV2Jwks{}
	this.Kty = kty
	this.Kid = kid
	this.Alg = alg
	return &this
}

// NewIamV2JwksWithDefaults instantiates a new IamV2Jwks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamV2JwksWithDefaults() *IamV2Jwks {
	this := IamV2Jwks{}
	return &this
}

// GetKty returns the Kty field value
func (o *IamV2Jwks) GetKty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kty
}

// GetKtyOk returns a tuple with the Kty field value
// and a boolean to check if the value has been set.
func (o *IamV2Jwks) GetKtyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kty, true
}

// SetKty sets field value
func (o *IamV2Jwks) SetKty(v string) {
	o.Kty = v
}

// GetKid returns the Kid field value
func (o *IamV2Jwks) GetKid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kid
}

// GetKidOk returns a tuple with the Kid field value
// and a boolean to check if the value has been set.
func (o *IamV2Jwks) GetKidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kid, true
}

// SetKid sets field value
func (o *IamV2Jwks) SetKid(v string) {
	o.Kid = v
}

// GetAlg returns the Alg field value
func (o *IamV2Jwks) GetAlg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alg
}

// GetAlgOk returns a tuple with the Alg field value
// and a boolean to check if the value has been set.
func (o *IamV2Jwks) GetAlgOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Alg, true
}

// SetAlg sets field value
func (o *IamV2Jwks) SetAlg(v string) {
	o.Alg = v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *IamV2Jwks) GetUse() string {
	if o == nil || o.Use == nil {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2Jwks) GetUseOk() (*string, bool) {
	if o == nil || o.Use == nil {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *IamV2Jwks) HasUse() bool {
	if o != nil && o.Use != nil {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *IamV2Jwks) SetUse(v string) {
	o.Use = &v
}

// GetKeyOps returns the KeyOps field value if set, zero value otherwise.
func (o *IamV2Jwks) GetKeyOps() string {
	if o == nil || o.KeyOps == nil {
		var ret string
		return ret
	}
	return *o.KeyOps
}

// GetKeyOpsOk returns a tuple with the KeyOps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2Jwks) GetKeyOpsOk() (*string, bool) {
	if o == nil || o.KeyOps == nil {
		return nil, false
	}
	return o.KeyOps, true
}

// HasKeyOps returns a boolean if a field has been set.
func (o *IamV2Jwks) HasKeyOps() bool {
	if o != nil && o.KeyOps != nil {
		return true
	}

	return false
}

// SetKeyOps gets a reference to the given string and assigns it to the KeyOps field.
func (o *IamV2Jwks) SetKeyOps(v string) {
	o.KeyOps = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *IamV2Jwks) Redact() {
    o.recurseRedact(&o.Kty)
    o.recurseRedact(&o.Kid)
    o.recurseRedact(&o.Alg)
    o.recurseRedact(o.Use)
    o.recurseRedact(o.KeyOps)
}

func (o *IamV2Jwks) recurseRedact(v interface{}) {
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

func (o IamV2Jwks) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o IamV2Jwks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kty"] = o.Kty
	}
	if true {
		toSerialize["kid"] = o.Kid
	}
	if true {
		toSerialize["alg"] = o.Alg
	}
	if o.Use != nil {
		toSerialize["use"] = o.Use
	}
	if o.KeyOps != nil {
		toSerialize["key_ops"] = o.KeyOps
	}
	return json.Marshal(toSerialize)
}

type NullableIamV2Jwks struct {
	value *IamV2Jwks
	isSet bool
}

func (v NullableIamV2Jwks) Get() *IamV2Jwks {
	return v.value
}

func (v *NullableIamV2Jwks) Set(val *IamV2Jwks) {
	v.value = val
	v.isSet = true
}

func (v NullableIamV2Jwks) IsSet() bool {
	return v.isSet
}

func (v *NullableIamV2Jwks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamV2Jwks(val *IamV2Jwks) *NullableIamV2Jwks {
	return &NullableIamV2Jwks{value: val, isSet: true}
}

func (v NullableIamV2Jwks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamV2Jwks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


