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
Go Template Service API

Go Template Service API for Contrast Testing

API version: 0.0.1-alpha0
Contact: apifengineers@confluent.io
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

// TemplateV1WidgetSpecUpdate The desired state of the Widget
type TemplateV1WidgetSpecUpdate struct {
	// The size of the widget
	Size *string `json:"size,omitempty"`
	// The color of the widget
	Color *string `json:"color,omitempty"`
}

// NewTemplateV1WidgetSpecUpdate instantiates a new TemplateV1WidgetSpecUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateV1WidgetSpecUpdate() *TemplateV1WidgetSpecUpdate {
	this := TemplateV1WidgetSpecUpdate{}
	return &this
}

// NewTemplateV1WidgetSpecUpdateWithDefaults instantiates a new TemplateV1WidgetSpecUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateV1WidgetSpecUpdateWithDefaults() *TemplateV1WidgetSpecUpdate {
	this := TemplateV1WidgetSpecUpdate{}
	return &this
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *TemplateV1WidgetSpecUpdate) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateV1WidgetSpecUpdate) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *TemplateV1WidgetSpecUpdate) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *TemplateV1WidgetSpecUpdate) SetSize(v string) {
	o.Size = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *TemplateV1WidgetSpecUpdate) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateV1WidgetSpecUpdate) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *TemplateV1WidgetSpecUpdate) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *TemplateV1WidgetSpecUpdate) SetColor(v string) {
	o.Color = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *TemplateV1WidgetSpecUpdate) Redact() {
	o.recurseRedact(o.Size)
	o.recurseRedact(o.Color)
}

func (o *TemplateV1WidgetSpecUpdate) recurseRedact(v interface{}) {
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

func (o TemplateV1WidgetSpecUpdate) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o TemplateV1WidgetSpecUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableTemplateV1WidgetSpecUpdate struct {
	value *TemplateV1WidgetSpecUpdate
	isSet bool
}

func (v NullableTemplateV1WidgetSpecUpdate) Get() *TemplateV1WidgetSpecUpdate {
	return v.value
}

func (v *NullableTemplateV1WidgetSpecUpdate) Set(val *TemplateV1WidgetSpecUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateV1WidgetSpecUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateV1WidgetSpecUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateV1WidgetSpecUpdate(val *TemplateV1WidgetSpecUpdate) *NullableTemplateV1WidgetSpecUpdate {
	return &NullableTemplateV1WidgetSpecUpdate{value: val, isSet: true}
}

func (v NullableTemplateV1WidgetSpecUpdate) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableTemplateV1WidgetSpecUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}