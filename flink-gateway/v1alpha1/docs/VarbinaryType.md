# VarbinaryType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nullable** | **bool** | Indicates whether values in this column can be null. | 
**Type** | **string** | The data type of the column. | 
**Length** | Pointer to **int32** | The length of the column. | [optional] 

## Methods

### NewVarbinaryType

`func NewVarbinaryType(nullable bool, type_ string, ) *VarbinaryType`

NewVarbinaryType instantiates a new VarbinaryType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVarbinaryTypeWithDefaults

`func NewVarbinaryTypeWithDefaults() *VarbinaryType`

NewVarbinaryTypeWithDefaults instantiates a new VarbinaryType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNullable

`func (o *VarbinaryType) GetNullable() bool`

GetNullable returns the Nullable field if non-nil, zero value otherwise.

### GetNullableOk

`func (o *VarbinaryType) GetNullableOk() (*bool, bool)`

GetNullableOk returns a tuple with the Nullable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullable

`func (o *VarbinaryType) SetNullable(v bool)`

SetNullable sets Nullable field to given value.


### GetType

`func (o *VarbinaryType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VarbinaryType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VarbinaryType) SetType(v string)`

SetType sets Type field to given value.


### GetLength

`func (o *VarbinaryType) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *VarbinaryType) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *VarbinaryType) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *VarbinaryType) HasLength() bool`

HasLength returns a boolean if a field has been set.


### AsDataType

`func (s *VarbinaryType) AsDataType() DataType`

Convenience method to wrap this instance of VarbinaryType in DataType

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


