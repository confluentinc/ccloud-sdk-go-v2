# MapType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**KeyId** | **int32** |  | 
**Key** | [**Type**](Type.md) |  | 
**ValueId** | **int32** |  | 
**Value** | [**Type**](Type.md) |  | 
**ValueRequired** | **bool** |  | 

## Methods

### NewMapType

`func NewMapType(type_ string, keyId int32, key Type, valueId int32, value Type, valueRequired bool, ) *MapType`

NewMapType instantiates a new MapType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapTypeWithDefaults

`func NewMapTypeWithDefaults() *MapType`

NewMapTypeWithDefaults instantiates a new MapType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetKeyId

`func (o *MapType) GetKeyId() int32`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *MapType) GetKeyIdOk() (*int32, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *MapType) SetKeyId(v int32)`

SetKeyId sets KeyId field to given value.


### GetKey

`func (o *MapType) GetKey() Type`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MapType) GetKeyOk() (*Type, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MapType) SetKey(v Type)`

SetKey sets Key field to given value.


### GetValueId

`func (o *MapType) GetValueId() int32`

GetValueId returns the ValueId field if non-nil, zero value otherwise.

### GetValueIdOk

`func (o *MapType) GetValueIdOk() (*int32, bool)`

GetValueIdOk returns a tuple with the ValueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueId

`func (o *MapType) SetValueId(v int32)`

SetValueId sets ValueId field to given value.


### GetValue

`func (o *MapType) GetValue() Type`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MapType) GetValueOk() (*Type, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MapType) SetValue(v Type)`

SetValue sets Value field to given value.


### GetValueRequired

`func (o *MapType) GetValueRequired() bool`

GetValueRequired returns the ValueRequired field if non-nil, zero value otherwise.

### GetValueRequiredOk

`func (o *MapType) GetValueRequiredOk() (*bool, bool)`

GetValueRequiredOk returns a tuple with the ValueRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueRequired

`func (o *MapType) SetValueRequired(v bool)`

SetValueRequired sets ValueRequired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


