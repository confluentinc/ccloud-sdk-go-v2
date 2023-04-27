# VarcharType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nullable** | **bool** | Indicates whether values in this column can be null. | 
**Type** | **string** | The data type of the column. | 
**Length** | Pointer to **int32** | The length of the column. | [optional] 

## Methods

### NewVarcharType

`func NewVarcharType(nullable bool, type_ string, ) *VarcharType`

NewVarcharType instantiates a new VarcharType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVarcharTypeWithDefaults

`func NewVarcharTypeWithDefaults() *VarcharType`

NewVarcharTypeWithDefaults instantiates a new VarcharType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNullable

`func (o *VarcharType) GetNullable() bool`

GetNullable returns the Nullable field if non-nil, zero value otherwise.

### GetNullableOk

`func (o *VarcharType) GetNullableOk() (*bool, bool)`

GetNullableOk returns a tuple with the Nullable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullable

`func (o *VarcharType) SetNullable(v bool)`

SetNullable sets Nullable field to given value.


### GetType

`func (o *VarcharType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VarcharType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VarcharType) SetType(v string)`

SetType sets Type field to given value.


### GetLength

`func (o *VarcharType) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *VarcharType) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *VarcharType) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *VarcharType) HasLength() bool`

HasLength returns a boolean if a field has been set.


### AsDataType

`func (s *VarcharType) AsDataType() DataType`

Convenience method to wrap this instance of VarcharType in DataType

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


