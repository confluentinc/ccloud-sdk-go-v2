# PartitionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpecId** | Pointer to **int32** |  | [optional] [readonly] 
**Fields** | [**[]PartitionField**](PartitionField.md) |  | 

## Methods

### NewPartitionSpec

`func NewPartitionSpec(fields []PartitionField, ) *PartitionSpec`

NewPartitionSpec instantiates a new PartitionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionSpecWithDefaults

`func NewPartitionSpecWithDefaults() *PartitionSpec`

NewPartitionSpecWithDefaults instantiates a new PartitionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpecId

`func (o *PartitionSpec) GetSpecId() int32`

GetSpecId returns the SpecId field if non-nil, zero value otherwise.

### GetSpecIdOk

`func (o *PartitionSpec) GetSpecIdOk() (*int32, bool)`

GetSpecIdOk returns a tuple with the SpecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecId

`func (o *PartitionSpec) SetSpecId(v int32)`

SetSpecId sets SpecId field to given value.

### HasSpecId

`func (o *PartitionSpec) HasSpecId() bool`

HasSpecId returns a boolean if a field has been set.

### GetFields

`func (o *PartitionSpec) GetFields() []PartitionField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PartitionSpec) GetFieldsOk() (*[]PartitionField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PartitionSpec) SetFields(v []PartitionField)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


