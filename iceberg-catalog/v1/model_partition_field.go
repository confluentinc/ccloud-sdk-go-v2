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

// checks if the PartitionField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartitionField{}

// PartitionField struct for PartitionField
type PartitionField struct {
	FieldId   *int32 `json:"field-id,omitempty"`
	SourceId  int32  `json:"source-id"`
	Name      string `json:"name"`
	Transform string `json:"transform"`
}

type _PartitionField PartitionField

// NewPartitionField instantiates a new PartitionField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartitionField(sourceId int32, name string, transform string) *PartitionField {
	this := PartitionField{}
	this.SourceId = sourceId
	this.Name = name
	this.Transform = transform
	return &this
}

// NewPartitionFieldWithDefaults instantiates a new PartitionField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartitionFieldWithDefaults() *PartitionField {
	this := PartitionField{}
	return &this
}

// GetFieldId returns the FieldId field value if set, zero value otherwise.
func (o *PartitionField) GetFieldId() int32 {
	if o == nil || IsNil(o.FieldId) {
		var ret int32
		return ret
	}
	return *o.FieldId
}

// GetFieldIdOk returns a tuple with the FieldId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartitionField) GetFieldIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FieldId) {
		return nil, false
	}
	return o.FieldId, true
}

// HasFieldId returns a boolean if a field has been set.
func (o *PartitionField) HasFieldId() bool {
	if o != nil && !IsNil(o.FieldId) {
		return true
	}

	return false
}

// SetFieldId gets a reference to the given int32 and assigns it to the FieldId field.
func (o *PartitionField) SetFieldId(v int32) {
	o.FieldId = &v
}

// GetSourceId returns the SourceId field value
func (o *PartitionField) GetSourceId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *PartitionField) GetSourceIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *PartitionField) SetSourceId(v int32) {
	o.SourceId = v
}

// GetName returns the Name field value
func (o *PartitionField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PartitionField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PartitionField) SetName(v string) {
	o.Name = v
}

// GetTransform returns the Transform field value
func (o *PartitionField) GetTransform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Transform
}

// GetTransformOk returns a tuple with the Transform field value
// and a boolean to check if the value has been set.
func (o *PartitionField) GetTransformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transform, true
}

// SetTransform sets field value
func (o *PartitionField) SetTransform(v string) {
	o.Transform = v
}

func (o PartitionField) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartitionField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FieldId) {
		toSerialize["field-id"] = o.FieldId
	}
	toSerialize["source-id"] = o.SourceId
	toSerialize["name"] = o.Name
	toSerialize["transform"] = o.Transform
	return toSerialize, nil
}

func (o *PartitionField) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source-id",
		"name",
		"transform",
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

	varPartitionField := _PartitionField{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varPartitionField)

	if err != nil {
		return err
	}

	*o = PartitionField(varPartitionField)

	return err
}

type NullablePartitionField struct {
	value *PartitionField
	isSet bool
}

func (v NullablePartitionField) Get() *PartitionField {
	return v.value
}

func (v *NullablePartitionField) Set(val *PartitionField) {
	v.value = val
	v.isSet = true
}

func (v NullablePartitionField) IsSet() bool {
	return v.isSet
}

func (v *NullablePartitionField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartitionField(val *PartitionField) *NullablePartitionField {
	return &NullablePartitionField{value: val, isSet: true}
}

func (v NullablePartitionField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartitionField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}