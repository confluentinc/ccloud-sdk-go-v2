# MultisetType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nullable** | **bool** | Indicates whether values in this column can be null. | 
**Type** | **string** | The data type of the column. | 
**MultisetElementType** | [**DataType**](DataType.md) | The type of the elements in the multiset. | 

## Methods

### NewMultisetType

`func NewMultisetType(nullable bool, type_ string, multisetElementType DataType, ) *MultisetType`

NewMultisetType instantiates a new MultisetType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultisetTypeWithDefaults

`func NewMultisetTypeWithDefaults() *MultisetType`

NewMultisetTypeWithDefaults instantiates a new MultisetType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNullable

`func (o *MultisetType) GetNullable() bool`

GetNullable returns the Nullable field if non-nil, zero value otherwise.

### GetNullableOk

`func (o *MultisetType) GetNullableOk() (*bool, bool)`

GetNullableOk returns a tuple with the Nullable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullable

`func (o *MultisetType) SetNullable(v bool)`

SetNullable sets Nullable field to given value.


### GetType

`func (o *MultisetType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MultisetType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MultisetType) SetType(v string)`

SetType sets Type field to given value.


### GetMultisetElementType

`func (o *MultisetType) GetMultisetElementType() DataType`

GetMultisetElementType returns the MultisetElementType field if non-nil, zero value otherwise.

### GetMultisetElementTypeOk

`func (o *MultisetType) GetMultisetElementTypeOk() (*DataType, bool)`

GetMultisetElementTypeOk returns a tuple with the MultisetElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultisetElementType

`func (o *MultisetType) SetMultisetElementType(v DataType)`

SetMultisetElementType sets MultisetElementType field to given value.



### AsDataType

`func (s *MultisetType) AsDataType() DataType`

Convenience method to wrap this instance of MultisetType in DataType

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


