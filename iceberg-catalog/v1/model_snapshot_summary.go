/*
Apache Iceberg REST Catalog API

Defines the specification for the first version of the REST Catalog API. Implementations should ideally support both Iceberg table specs v1 and v2, with priority given to v2.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SnapshotSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnapshotSummary{}

// SnapshotSummary struct for SnapshotSummary
type SnapshotSummary struct {
	Operation                 string  `json:"operation"`
	AdditionalPropertiesField *string `json:"additionalProperties,omitempty"`
}

type _SnapshotSummary SnapshotSummary

// NewSnapshotSummary instantiates a new SnapshotSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotSummary(operation string) *SnapshotSummary {
	this := SnapshotSummary{}
	this.Operation = operation
	return &this
}

// NewSnapshotSummaryWithDefaults instantiates a new SnapshotSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotSummaryWithDefaults() *SnapshotSummary {
	this := SnapshotSummary{}
	return &this
}

// GetOperation returns the Operation field value
func (o *SnapshotSummary) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *SnapshotSummary) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *SnapshotSummary) SetOperation(v string) {
	o.Operation = v
}

// GetAdditionalPropertiesField returns the AdditionalPropertiesField field value if set, zero value otherwise.
func (o *SnapshotSummary) GetAdditionalPropertiesField() string {
	if o == nil || IsNil(o.AdditionalPropertiesField) {
		var ret string
		return ret
	}
	return *o.AdditionalPropertiesField
}

// GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotSummary) GetAdditionalPropertiesFieldOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalPropertiesField) {
		return nil, false
	}
	return o.AdditionalPropertiesField, true
}

// HasAdditionalPropertiesField returns a boolean if a field has been set.
func (o *SnapshotSummary) HasAdditionalPropertiesField() bool {
	if o != nil && !IsNil(o.AdditionalPropertiesField) {
		return true
	}

	return false
}

// SetAdditionalPropertiesField gets a reference to the given string and assigns it to the AdditionalPropertiesField field.
func (o *SnapshotSummary) SetAdditionalPropertiesField(v string) {
	o.AdditionalPropertiesField = &v
}

func (o SnapshotSummary) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnapshotSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operation"] = o.Operation
	if !IsNil(o.AdditionalPropertiesField) {
		toSerialize["additionalProperties"] = o.AdditionalPropertiesField
	}
	return toSerialize, nil
}

func (o *SnapshotSummary) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"operation",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSnapshotSummary := _SnapshotSummary{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSnapshotSummary)

	if err != nil {
		return err
	}

	*o = SnapshotSummary(varSnapshotSummary)

	return err
}

type NullableSnapshotSummary struct {
	value *SnapshotSummary
	isSet bool
}

func (v NullableSnapshotSummary) Get() *SnapshotSummary {
	return v.value
}

func (v *NullableSnapshotSummary) Set(val *SnapshotSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotSummary(val *SnapshotSummary) *NullableSnapshotSummary {
	return &NullableSnapshotSummary{value: val, isSet: true}
}

func (v NullableSnapshotSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}