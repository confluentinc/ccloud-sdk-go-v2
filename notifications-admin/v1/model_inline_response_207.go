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

// InlineResponse207 struct for InlineResponse207
type InlineResponse207 struct {
	// Data for individual response of overall multi-status response.
	Results *[]InlineResponse207Results `json:"results,omitempty"`
	Metadata *InlineResponse207Metadata `json:"metadata,omitempty"`
}

// NewInlineResponse207 instantiates a new InlineResponse207 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse207() *InlineResponse207 {
	this := InlineResponse207{}
	return &this
}

// NewInlineResponse207WithDefaults instantiates a new InlineResponse207 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse207WithDefaults() *InlineResponse207 {
	this := InlineResponse207{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *InlineResponse207) GetResults() []InlineResponse207Results {
	if o == nil || o.Results == nil {
		var ret []InlineResponse207Results
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse207) GetResultsOk() (*[]InlineResponse207Results, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *InlineResponse207) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []InlineResponse207Results and assigns it to the Results field.
func (o *InlineResponse207) SetResults(v []InlineResponse207Results) {
	o.Results = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *InlineResponse207) GetMetadata() InlineResponse207Metadata {
	if o == nil || o.Metadata == nil {
		var ret InlineResponse207Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse207) GetMetadataOk() (*InlineResponse207Metadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *InlineResponse207) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given InlineResponse207Metadata and assigns it to the Metadata field.
func (o *InlineResponse207) SetMetadata(v InlineResponse207Metadata) {
	o.Metadata = &v
}

func (o InlineResponse207) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse207 struct {
	value *InlineResponse207
	isSet bool
}

func (v NullableInlineResponse207) Get() *InlineResponse207 {
	return v.value
}

func (v *NullableInlineResponse207) Set(val *InlineResponse207) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse207) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse207) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse207(val *InlineResponse207) *NullableInlineResponse207 {
	return &NullableInlineResponse207{value: val, isSet: true}
}

func (v NullableInlineResponse207) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse207) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

