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
Custom Plugin Management API

This is Custom Plugin Management API.

API version: 1.0.0
Contact: compute-platform-team@confluent.io
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

// CcpV1CustomPluginConnect Connect Custom plugin type config.
type CcpV1CustomPluginConnect struct {
	// [immutable] Plugin Type.
	PluginType string `json:"plugin_type,omitempty"`
	// [immutable] Java class or alias for connector. You can get connector class from connector documentation provided by developer.
	ConnectorClass string `json:"connector_class,omitempty"`
	// [immutable] Custom Connector type.
	ConnectorType string `json:"connector_type,omitempty"`
	// A sensitive property is a connector configuration property that must be hidden after a user enters property value when setting up connector.
	SensitiveConfigProperties *[]string `json:"sensitive_config_properties,omitempty"`
}

// NewCcpV1CustomPluginConnect instantiates a new CcpV1CustomPluginConnect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCcpV1CustomPluginConnect(pluginType string, connectorClass string, connectorType string) *CcpV1CustomPluginConnect {
	this := CcpV1CustomPluginConnect{}
	this.PluginType = pluginType
	this.ConnectorClass = connectorClass
	this.ConnectorType = connectorType
	return &this
}

// NewCcpV1CustomPluginConnectWithDefaults instantiates a new CcpV1CustomPluginConnect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCcpV1CustomPluginConnectWithDefaults() *CcpV1CustomPluginConnect {
	this := CcpV1CustomPluginConnect{}
	return &this
}

// GetPluginType returns the PluginType field value
func (o *CcpV1CustomPluginConnect) GetPluginType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PluginType
}

// GetPluginTypeOk returns a tuple with the PluginType field value
// and a boolean to check if the value has been set.
func (o *CcpV1CustomPluginConnect) GetPluginTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginType, true
}

// SetPluginType sets field value
func (o *CcpV1CustomPluginConnect) SetPluginType(v string) {
	o.PluginType = v
}

// GetConnectorClass returns the ConnectorClass field value
func (o *CcpV1CustomPluginConnect) GetConnectorClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorClass
}

// GetConnectorClassOk returns a tuple with the ConnectorClass field value
// and a boolean to check if the value has been set.
func (o *CcpV1CustomPluginConnect) GetConnectorClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorClass, true
}

// SetConnectorClass sets field value
func (o *CcpV1CustomPluginConnect) SetConnectorClass(v string) {
	o.ConnectorClass = v
}

// GetConnectorType returns the ConnectorType field value
func (o *CcpV1CustomPluginConnect) GetConnectorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorType
}

// GetConnectorTypeOk returns a tuple with the ConnectorType field value
// and a boolean to check if the value has been set.
func (o *CcpV1CustomPluginConnect) GetConnectorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorType, true
}

// SetConnectorType sets field value
func (o *CcpV1CustomPluginConnect) SetConnectorType(v string) {
	o.ConnectorType = v
}

// GetSensitiveConfigProperties returns the SensitiveConfigProperties field value if set, zero value otherwise.
func (o *CcpV1CustomPluginConnect) GetSensitiveConfigProperties() []string {
	if o == nil || o.SensitiveConfigProperties == nil {
		var ret []string
		return ret
	}
	return *o.SensitiveConfigProperties
}

// GetSensitiveConfigPropertiesOk returns a tuple with the SensitiveConfigProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CcpV1CustomPluginConnect) GetSensitiveConfigPropertiesOk() (*[]string, bool) {
	if o == nil || o.SensitiveConfigProperties == nil {
		return nil, false
	}
	return o.SensitiveConfigProperties, true
}

// HasSensitiveConfigProperties returns a boolean if a field has been set.
func (o *CcpV1CustomPluginConnect) HasSensitiveConfigProperties() bool {
	if o != nil && o.SensitiveConfigProperties != nil {
		return true
	}

	return false
}

// SetSensitiveConfigProperties gets a reference to the given []string and assigns it to the SensitiveConfigProperties field.
func (o *CcpV1CustomPluginConnect) SetSensitiveConfigProperties(v []string) {
	o.SensitiveConfigProperties = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *CcpV1CustomPluginConnect) Redact() {
	o.recurseRedact(&o.PluginType)
	o.recurseRedact(&o.ConnectorClass)
	o.recurseRedact(&o.ConnectorType)
	o.recurseRedact(o.SensitiveConfigProperties)
}

func (o *CcpV1CustomPluginConnect) recurseRedact(v interface{}) {
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

func (o CcpV1CustomPluginConnect) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o CcpV1CustomPluginConnect) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["plugin_type"] = o.PluginType
	}
	if true {
		toSerialize["connector_class"] = o.ConnectorClass
	}
	if true {
		toSerialize["connector_type"] = o.ConnectorType
	}
	if o.SensitiveConfigProperties != nil {
		toSerialize["sensitive_config_properties"] = o.SensitiveConfigProperties
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableCcpV1CustomPluginConnect struct {
	value *CcpV1CustomPluginConnect
	isSet bool
}

func (v NullableCcpV1CustomPluginConnect) Get() *CcpV1CustomPluginConnect {
	return v.value
}

func (v *NullableCcpV1CustomPluginConnect) Set(val *CcpV1CustomPluginConnect) {
	v.value = val
	v.isSet = true
}

func (v NullableCcpV1CustomPluginConnect) IsSet() bool {
	return v.isSet
}

func (v *NullableCcpV1CustomPluginConnect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCcpV1CustomPluginConnect(val *CcpV1CustomPluginConnect) *NullableCcpV1CustomPluginConnect {
	return &NullableCcpV1CustomPluginConnect{value: val, isSet: true}
}

func (v NullableCcpV1CustomPluginConnect) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableCcpV1CustomPluginConnect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}