/*
 * Network API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1-alpha1
 * Contact: cire-traffic@confluent.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// ErrorSource struct for ErrorSource
type ErrorSource struct {
	// A JSON Pointer [RFC6901] to the associated entity in the request document [e.g. \"/spec\" for a spec object, or \"/spec/title\" for a specific field].
	Pointer *string `json:"pointer,omitempty"`
	// A string indicating which query parameter caused the error.
	Parameter *string `json:"parameter,omitempty"`
}

// NewErrorSource instantiates a new ErrorSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorSource() *ErrorSource {
	this := ErrorSource{}
	return &this
}

// NewErrorSourceWithDefaults instantiates a new ErrorSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorSourceWithDefaults() *ErrorSource {
	this := ErrorSource{}
	return &this
}

// GetPointer returns the Pointer field value if set, zero value otherwise.
func (o *ErrorSource) GetPointer() string {
	if o == nil || o.Pointer == nil {
		var ret string
		return ret
	}
	return *o.Pointer
}

// GetPointerOk returns a tuple with the Pointer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorSource) GetPointerOk() (*string, bool) {
	if o == nil || o.Pointer == nil {
		return nil, false
	}
	return o.Pointer, true
}

// HasPointer returns a boolean if a field has been set.
func (o *ErrorSource) HasPointer() bool {
	if o != nil && o.Pointer != nil {
		return true
	}

	return false
}

// SetPointer gets a reference to the given string and assigns it to the Pointer field.
func (o *ErrorSource) SetPointer(v string) {
	o.Pointer = &v
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *ErrorSource) GetParameter() string {
	if o == nil || o.Parameter == nil {
		var ret string
		return ret
	}
	return *o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorSource) GetParameterOk() (*string, bool) {
	if o == nil || o.Parameter == nil {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *ErrorSource) HasParameter() bool {
	if o != nil && o.Parameter != nil {
		return true
	}

	return false
}

// SetParameter gets a reference to the given string and assigns it to the Parameter field.
func (o *ErrorSource) SetParameter(v string) {
	o.Parameter = &v
}

func (o ErrorSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pointer != nil {
		toSerialize["pointer"] = o.Pointer
	}
	if o.Parameter != nil {
		toSerialize["parameter"] = o.Parameter
	}
	return json.Marshal(toSerialize)
}

type NullableErrorSource struct {
	value *ErrorSource
	isSet bool
}

func (v NullableErrorSource) Get() *ErrorSource {
	return v.value
}

func (v *NullableErrorSource) Set(val *ErrorSource) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorSource) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorSource(val *ErrorSource) *NullableErrorSource {
	return &NullableErrorSource{value: val, isSet: true}
}

func (v NullableErrorSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

