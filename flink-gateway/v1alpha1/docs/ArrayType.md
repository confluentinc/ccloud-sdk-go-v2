# ArrayType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nullable** | **bool** | Indicates whether values in this column can be null. | 
**Type** | **string** | The data type of the column. | 
**ArrayElementType** | [**DataType**](DataType.md) | The type of the elements in the array. | 

## Methods

### NewArrayType

`func NewArrayType(nullable bool, type_ string, arrayElementType DataType, ) *ArrayType`

NewArrayType instantiates a new ArrayType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArrayTypeWithDefaults

`func NewArrayTypeWithDefaults() *ArrayType`

NewArrayTypeWithDefaults instantiates a new ArrayType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNullable

`func (o *ArrayType) GetNullable() bool`

GetNullable returns the Nullable field if non-nil, zero value otherwise.

### GetNullableOk

`func (o *ArrayType) GetNullableOk() (*bool, bool)`

GetNullableOk returns a tuple with the Nullable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullable

`func (o *ArrayType) SetNullable(v bool)`

SetNullable sets Nullable field to given value.


### GetType

`func (o *ArrayType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArrayType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArrayType) SetType(v string)`

SetType sets Type field to given value.


### GetArrayElementType

`func (o *ArrayType) GetArrayElementType() DataType`

GetArrayElementType returns the ArrayElementType field if non-nil, zero value otherwise.

### GetArrayElementTypeOk

`func (o *ArrayType) GetArrayElementTypeOk() (*DataType, bool)`

GetArrayElementTypeOk returns a tuple with the ArrayElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayElementType

`func (o *ArrayType) SetArrayElementType(v DataType)`

SetArrayElementType sets ArrayElementType field to given value.



### AsDataType

`func (s *ArrayType) AsDataType() DataType`

Convenience method to wrap this instance of ArrayType in DataType

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


