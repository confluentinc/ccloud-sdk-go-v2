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

// NetworkingV1AwsTransitGatewayAttachment AWS Transit Gateway Attachment.
type NetworkingV1AwsTransitGatewayAttachment struct {
	// AWS Transit Gateway Attachment kind type.
	Kind string `json:"kind,omitempty"`
	// The full AWS Resource Name (ARN) for the AWS Resource Access Manager (RAM) Share of the Transit Gateways that you want Confluent Cloud to be attached to.
	RamShareArn string `json:"ram_share_arn,omitempty"`
	// The ID of the AWS Transit Gateway that you want Confluent CLoud to be attached to.
	TransitGatewayId string `json:"transit_gateway_id,omitempty"`
	// List of destination routes.
	Routes []string `json:"routes,omitempty"`
}

// NewNetworkingV1AwsTransitGatewayAttachment instantiates a new NetworkingV1AwsTransitGatewayAttachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1AwsTransitGatewayAttachment(kind string, ramShareArn string, transitGatewayId string, routes []string) *NetworkingV1AwsTransitGatewayAttachment {
	this := NetworkingV1AwsTransitGatewayAttachment{}
	this.Kind = kind
	this.RamShareArn = ramShareArn
	this.TransitGatewayId = transitGatewayId
	this.Routes = routes
	return &this
}

// NewNetworkingV1AwsTransitGatewayAttachmentWithDefaults instantiates a new NetworkingV1AwsTransitGatewayAttachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1AwsTransitGatewayAttachmentWithDefaults() *NetworkingV1AwsTransitGatewayAttachment {
	this := NetworkingV1AwsTransitGatewayAttachment{}
	return &this
}

// GetKind returns the Kind field value
func (o *NetworkingV1AwsTransitGatewayAttachment) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AwsTransitGatewayAttachment) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *NetworkingV1AwsTransitGatewayAttachment) SetKind(v string) {
	o.Kind = v
}

// GetRamShareArn returns the RamShareArn field value
func (o *NetworkingV1AwsTransitGatewayAttachment) GetRamShareArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RamShareArn
}

// GetRamShareArnOk returns a tuple with the RamShareArn field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AwsTransitGatewayAttachment) GetRamShareArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RamShareArn, true
}

// SetRamShareArn sets field value
func (o *NetworkingV1AwsTransitGatewayAttachment) SetRamShareArn(v string) {
	o.RamShareArn = v
}

// GetTransitGatewayId returns the TransitGatewayId field value
func (o *NetworkingV1AwsTransitGatewayAttachment) GetTransitGatewayId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransitGatewayId
}

// GetTransitGatewayIdOk returns a tuple with the TransitGatewayId field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AwsTransitGatewayAttachment) GetTransitGatewayIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransitGatewayId, true
}

// SetTransitGatewayId sets field value
func (o *NetworkingV1AwsTransitGatewayAttachment) SetTransitGatewayId(v string) {
	o.TransitGatewayId = v
}

// GetRoutes returns the Routes field value
func (o *NetworkingV1AwsTransitGatewayAttachment) GetRoutes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AwsTransitGatewayAttachment) GetRoutesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Routes, true
}

// SetRoutes sets field value
func (o *NetworkingV1AwsTransitGatewayAttachment) SetRoutes(v []string) {
	o.Routes = v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1AwsTransitGatewayAttachment) Redact() {
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.RamShareArn)
	o.recurseRedact(&o.TransitGatewayId)
	o.recurseRedact(&o.Routes)
}

func (o *NetworkingV1AwsTransitGatewayAttachment) recurseRedact(v interface{}) {
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

func (o NetworkingV1AwsTransitGatewayAttachment) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1AwsTransitGatewayAttachment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["ram_share_arn"] = o.RamShareArn
	}
	if true {
		toSerialize["transit_gateway_id"] = o.TransitGatewayId
	}
	if true {
		toSerialize["routes"] = o.Routes
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableNetworkingV1AwsTransitGatewayAttachment struct {
	value *NetworkingV1AwsTransitGatewayAttachment
	isSet bool
}

func (v NullableNetworkingV1AwsTransitGatewayAttachment) Get() *NetworkingV1AwsTransitGatewayAttachment {
	return v.value
}

func (v *NullableNetworkingV1AwsTransitGatewayAttachment) Set(val *NetworkingV1AwsTransitGatewayAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1AwsTransitGatewayAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1AwsTransitGatewayAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1AwsTransitGatewayAttachment(val *NetworkingV1AwsTransitGatewayAttachment) *NullableNetworkingV1AwsTransitGatewayAttachment {
	return &NullableNetworkingV1AwsTransitGatewayAttachment{value: val, isSet: true}
}

func (v NullableNetworkingV1AwsTransitGatewayAttachment) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableNetworkingV1AwsTransitGatewayAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
