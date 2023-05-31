# RowFieldType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the field. | 
**FieldType** | [**DataType**](DataType.md) | The data type of the field. | 
**Description** | Pointer to **string** | The description of the field. | [optional] 

## Methods

### NewRowFieldType

`func NewRowFieldType(name string, fieldType DataType, ) *RowFieldType`

NewRowFieldType instantiates a new RowFieldType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRowFieldTypeWithDefaults

`func NewRowFieldTypeWithDefaults() *RowFieldType`

NewRowFieldTypeWithDefaults instantiates a new RowFieldType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RowFieldType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RowFieldType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RowFieldType) SetName(v string)`

SetName sets Name field to given value.


### GetFieldType

`func (o *RowFieldType) GetFieldType() DataType`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *RowFieldType) GetFieldTypeOk() (*DataType, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *RowFieldType) SetFieldType(v DataType)`

SetFieldType sets FieldType field to given value.


### GetDescription

`func (o *RowFieldType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RowFieldType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RowFieldType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RowFieldType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


