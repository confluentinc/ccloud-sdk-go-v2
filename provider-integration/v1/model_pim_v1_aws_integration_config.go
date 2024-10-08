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
Provider Integration Management API

REST API for the Provider Integration

API version: 0.1.0
Contact: connect@confluent.io
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

// PimV1AwsIntegrationConfig config schema for AWS cloud service provider.
type PimV1AwsIntegrationConfig struct {
	// Amazon Resource Name (ARN) that identifies the Amazon Web Services (AWS) Identity and Access Management (IAM) role that Confluent Cloud uses to assume customer IAM role when it accesses resources in your AWS account.
	IamRoleArn *string `json:"iam_role_arn,omitempty"`
	// Unique external ID that Confluent Cloud uses when it assumes the IAM role in your Amazon Web Services (AWS) account.
	ExternalId *string `json:"external_id,omitempty"`
	// Amazon Resource Name (ARN) that identifies the Amazon Web Services (AWS) Identity and Access Management (IAM) role that Confluent Cloud assumes when it accesses resources in your AWS account.
	CustomerIamRoleArn *string `json:"customer_iam_role_arn,omitempty"`
	// Cloud provider specific config to which access is provided through provider integration.
	Kind string `json:"kind,omitempty"`
}

// NewPimV1AwsIntegrationConfig instantiates a new PimV1AwsIntegrationConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPimV1AwsIntegrationConfig(kind string) *PimV1AwsIntegrationConfig {
	this := PimV1AwsIntegrationConfig{}
	this.Kind = kind
	return &this
}

// NewPimV1AwsIntegrationConfigWithDefaults instantiates a new PimV1AwsIntegrationConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPimV1AwsIntegrationConfigWithDefaults() *PimV1AwsIntegrationConfig {
	this := PimV1AwsIntegrationConfig{}
	return &this
}

// GetIamRoleArn returns the IamRoleArn field value if set, zero value otherwise.
func (o *PimV1AwsIntegrationConfig) GetIamRoleArn() string {
	if o == nil || o.IamRoleArn == nil {
		var ret string
		return ret
	}
	return *o.IamRoleArn
}

// GetIamRoleArnOk returns a tuple with the IamRoleArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimV1AwsIntegrationConfig) GetIamRoleArnOk() (*string, bool) {
	if o == nil || o.IamRoleArn == nil {
		return nil, false
	}
	return o.IamRoleArn, true
}

// HasIamRoleArn returns a boolean if a field has been set.
func (o *PimV1AwsIntegrationConfig) HasIamRoleArn() bool {
	if o != nil && o.IamRoleArn != nil {
		return true
	}

	return false
}

// SetIamRoleArn gets a reference to the given string and assigns it to the IamRoleArn field.
func (o *PimV1AwsIntegrationConfig) SetIamRoleArn(v string) {
	o.IamRoleArn = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *PimV1AwsIntegrationConfig) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimV1AwsIntegrationConfig) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *PimV1AwsIntegrationConfig) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *PimV1AwsIntegrationConfig) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetCustomerIamRoleArn returns the CustomerIamRoleArn field value if set, zero value otherwise.
func (o *PimV1AwsIntegrationConfig) GetCustomerIamRoleArn() string {
	if o == nil || o.CustomerIamRoleArn == nil {
		var ret string
		return ret
	}
	return *o.CustomerIamRoleArn
}

// GetCustomerIamRoleArnOk returns a tuple with the CustomerIamRoleArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimV1AwsIntegrationConfig) GetCustomerIamRoleArnOk() (*string, bool) {
	if o == nil || o.CustomerIamRoleArn == nil {
		return nil, false
	}
	return o.CustomerIamRoleArn, true
}

// HasCustomerIamRoleArn returns a boolean if a field has been set.
func (o *PimV1AwsIntegrationConfig) HasCustomerIamRoleArn() bool {
	if o != nil && o.CustomerIamRoleArn != nil {
		return true
	}

	return false
}

// SetCustomerIamRoleArn gets a reference to the given string and assigns it to the CustomerIamRoleArn field.
func (o *PimV1AwsIntegrationConfig) SetCustomerIamRoleArn(v string) {
	o.CustomerIamRoleArn = &v
}

// GetKind returns the Kind field value
func (o *PimV1AwsIntegrationConfig) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *PimV1AwsIntegrationConfig) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *PimV1AwsIntegrationConfig) SetKind(v string) {
	o.Kind = v
}

// Redact resets all sensitive fields to their zero value.
func (o *PimV1AwsIntegrationConfig) Redact() {
	o.recurseRedact(o.IamRoleArn)
	o.recurseRedact(o.ExternalId)
	o.recurseRedact(o.CustomerIamRoleArn)
	o.recurseRedact(&o.Kind)
}

func (o *PimV1AwsIntegrationConfig) recurseRedact(v interface{}) {
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

func (o PimV1AwsIntegrationConfig) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o PimV1AwsIntegrationConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IamRoleArn != nil {
		toSerialize["iam_role_arn"] = o.IamRoleArn
	}
	if o.ExternalId != nil {
		toSerialize["external_id"] = o.ExternalId
	}
	if o.CustomerIamRoleArn != nil {
		toSerialize["customer_iam_role_arn"] = o.CustomerIamRoleArn
	}
	if true {
		toSerialize["kind"] = o.Kind
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullablePimV1AwsIntegrationConfig struct {
	value *PimV1AwsIntegrationConfig
	isSet bool
}

func (v NullablePimV1AwsIntegrationConfig) Get() *PimV1AwsIntegrationConfig {
	return v.value
}

func (v *NullablePimV1AwsIntegrationConfig) Set(val *PimV1AwsIntegrationConfig) {
	v.value = val
	v.isSet = true
}

func (v NullablePimV1AwsIntegrationConfig) IsSet() bool {
	return v.isSet
}

func (v *NullablePimV1AwsIntegrationConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePimV1AwsIntegrationConfig(val *PimV1AwsIntegrationConfig) *NullablePimV1AwsIntegrationConfig {
	return &NullablePimV1AwsIntegrationConfig{value: val, isSet: true}
}

func (v NullablePimV1AwsIntegrationConfig) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullablePimV1AwsIntegrationConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
