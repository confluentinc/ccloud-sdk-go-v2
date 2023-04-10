# MapType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nullable** | **bool** | Indicates whether values in this column can be null. | 
**Type** | **string** | The data type of the column. | 
**KeyType** | [**DataType**](DataType.md) | The type of the key in the map. | 
**ValueType** | [**DataType**](DataType.md) | The type of the value in the map. | 

## Methods

### NewMapType

`func NewMapType(nullable bool, type_ string, keyType DataType, valueType DataType, ) *MapType`

NewMapType instantiates a new MapType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapTypeWithDefaults

`func NewMapTypeWithDefaults() *MapType`

NewMapTypeWithDefaults instantiates a new MapType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNullable

`func (o *MapType) GetNullable() bool`

GetNullable returns the Nullable field if non-nil, zero value otherwise.

### GetNullableOk

`func (o *MapType) GetNullableOk() (*bool, bool)`

GetNullableOk returns a tuple with the Nullable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullable

`func (o *MapType) SetNullable(v bool)`

SetNullable sets Nullable field to given value.


### GetType

`func (o *MapType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MapType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MapType) SetType(v string)`

SetType sets Type field to given value.


### GetKeyType

`func (o *MapType) GetKeyType() DataType`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *MapType) GetKeyTypeOk() (*DataType, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *MapType) SetKeyType(v DataType)`

SetKeyType sets KeyType field to given value.


### GetValueType

`func (o *MapType) GetValueType() DataType`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *MapType) GetValueTypeOk() (*DataType, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *MapType) SetValueType(v DataType)`

SetValueType sets ValueType field to given value.



### AsDataType

`func (s *MapType) AsDataType() DataType`

Convenience method to wrap this instance of MapType in DataType

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


