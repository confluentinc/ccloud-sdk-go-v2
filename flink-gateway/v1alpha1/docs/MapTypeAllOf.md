# MapTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The data type of the column. | [optional] 
**KeyType** | [**DataType**](DataType.md) | The type of the key in the map. | 
**ValueType** | [**DataType**](DataType.md) | The type of the value in the map. | 

## Methods

### NewMapTypeAllOf

`func NewMapTypeAllOf(keyType DataType, valueType DataType, ) *MapTypeAllOf`

NewMapTypeAllOf instantiates a new MapTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapTypeAllOfWithDefaults

`func NewMapTypeAllOfWithDefaults() *MapTypeAllOf`

NewMapTypeAllOfWithDefaults instantiates a new MapTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MapTypeAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MapTypeAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MapTypeAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MapTypeAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetKeyType

`func (o *MapTypeAllOf) GetKeyType() DataType`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *MapTypeAllOf) GetKeyTypeOk() (*DataType, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *MapTypeAllOf) SetKeyType(v DataType)`

SetKeyType sets KeyType field to given value.


### GetValueType

`func (o *MapTypeAllOf) GetValueType() DataType`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *MapTypeAllOf) GetValueTypeOk() (*DataType, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *MapTypeAllOf) SetValueType(v DataType)`

SetValueType sets ValueType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


