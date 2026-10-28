# QueryV1RowFieldType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the field. | 
**FieldType** | [**QueryV1DataType**](QueryV1DataType.md) | The data type of the field. | 
**Description** | Pointer to **string** | Optional field comment from the type declaration. | [optional] 

## Methods

### NewQueryV1RowFieldType

`func NewQueryV1RowFieldType(name string, fieldType QueryV1DataType, ) *QueryV1RowFieldType`

NewQueryV1RowFieldType instantiates a new QueryV1RowFieldType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryV1RowFieldTypeWithDefaults

`func NewQueryV1RowFieldTypeWithDefaults() *QueryV1RowFieldType`

NewQueryV1RowFieldTypeWithDefaults instantiates a new QueryV1RowFieldType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QueryV1RowFieldType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueryV1RowFieldType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueryV1RowFieldType) SetName(v string)`

SetName sets Name field to given value.


### GetFieldType

`func (o *QueryV1RowFieldType) GetFieldType() QueryV1DataType`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *QueryV1RowFieldType) GetFieldTypeOk() (*QueryV1DataType, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *QueryV1RowFieldType) SetFieldType(v QueryV1DataType)`

SetFieldType sets FieldType field to given value.


### GetDescription

`func (o *QueryV1RowFieldType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QueryV1RowFieldType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QueryV1RowFieldType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QueryV1RowFieldType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


