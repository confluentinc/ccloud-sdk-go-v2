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
Notification System Admin APIs

# Introduction  Notifications Admin API provides resources/API which are generally not applicable for normal users, in order to access these API user must have super user access for your access token  # Authorization Currently only confluent okta uses are only able to access these API.  All API requests must be made over HTTPS. Calls made over plain HTTP will fail. API requests without authentication will also fail.  The broad categories of Admin API for notification-service are categorises in below sections. * **Templates** -    Internal users like billing team will be able to add/modify the templates for notifications    in self service manner. Individual API will describe these operations in more details. * **Manual Notifications** -    To allow internal users to manually send out a notification to a given list of orgs. 

API version: 0.1.0-alpha0
Contact: observability-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// InlineResponse207Results Notification Delivery Status for an Organization Id.
type InlineResponse207Results struct {
	// Status code of message sent to specified integration.
	Status *int32 `json:"status,omitempty"`
	// Organization Id for which the notification is sent.
	OrgId *int32 `json:"org_id,omitempty"`
	// The type of the target system where notification is delivered.
	IntegrationKind *string `json:"integration_kind,omitempty"`
}

// NewInlineResponse207Results instantiates a new InlineResponse207Results object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse207Results() *InlineResponse207Results {
	this := InlineResponse207Results{}
	return &this
}

// NewInlineResponse207ResultsWithDefaults instantiates a new InlineResponse207Results object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse207ResultsWithDefaults() *InlineResponse207Results {
	this := InlineResponse207Results{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse207Results) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse207Results) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse207Results) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *InlineResponse207Results) SetStatus(v int32) {
	o.Status = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *InlineResponse207Results) GetOrgId() int32 {
	if o == nil || o.OrgId == nil {
		var ret int32
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse207Results) GetOrgIdOk() (*int32, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *InlineResponse207Results) HasOrgId() bool {
	if o != nil && o.OrgId != nil {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given int32 and assigns it to the OrgId field.
func (o *InlineResponse207Results) SetOrgId(v int32) {
	o.OrgId = &v
}

// GetIntegrationKind returns the IntegrationKind field value if set, zero value otherwise.
func (o *InlineResponse207Results) GetIntegrationKind() string {
	if o == nil || o.IntegrationKind == nil {
		var ret string
		return ret
	}
	return *o.IntegrationKind
}

// GetIntegrationKindOk returns a tuple with the IntegrationKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse207Results) GetIntegrationKindOk() (*string, bool) {
	if o == nil || o.IntegrationKind == nil {
		return nil, false
	}
	return o.IntegrationKind, true
}

// HasIntegrationKind returns a boolean if a field has been set.
func (o *InlineResponse207Results) HasIntegrationKind() bool {
	if o != nil && o.IntegrationKind != nil {
		return true
	}

	return false
}

// SetIntegrationKind gets a reference to the given string and assigns it to the IntegrationKind field.
func (o *InlineResponse207Results) SetIntegrationKind(v string) {
	o.IntegrationKind = &v
}

func (o InlineResponse207Results) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.OrgId != nil {
		toSerialize["org_id"] = o.OrgId
	}
	if o.IntegrationKind != nil {
		toSerialize["integration_kind"] = o.IntegrationKind
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse207Results struct {
	value *InlineResponse207Results
	isSet bool
}

func (v NullableInlineResponse207Results) Get() *InlineResponse207Results {
	return v.value
}

func (v *NullableInlineResponse207Results) Set(val *InlineResponse207Results) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse207Results) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse207Results) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse207Results(val *InlineResponse207Results) *NullableInlineResponse207Results {
	return &NullableInlineResponse207Results{value: val, isSet: true}
}

func (v NullableInlineResponse207Results) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse207Results) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

