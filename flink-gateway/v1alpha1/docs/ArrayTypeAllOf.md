# ArrayTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The data type of the column. | [optional] 
**ArrayElementType** | [**DataType**](DataType.md) | The type of the elements in the array. | 

## Methods

### NewArrayTypeAllOf

`func NewArrayTypeAllOf(arrayElementType DataType, ) *ArrayTypeAllOf`

NewArrayTypeAllOf instantiates a new ArrayTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArrayTypeAllOfWithDefaults

`func NewArrayTypeAllOfWithDefaults() *ArrayTypeAllOf`

NewArrayTypeAllOfWithDefaults instantiates a new ArrayTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ArrayTypeAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArrayTypeAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArrayTypeAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ArrayTypeAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetArrayElementType

`func (o *ArrayTypeAllOf) GetArrayElementType() DataType`

GetArrayElementType returns the ArrayElementType field if non-nil, zero value otherwise.

### GetArrayElementTypeOk

`func (o *ArrayTypeAllOf) GetArrayElementTypeOk() (*DataType, bool)`

GetArrayElementTypeOk returns a tuple with the ArrayElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayElementType

`func (o *ArrayTypeAllOf) SetArrayElementType(v DataType)`

SetArrayElementType sets ArrayElementType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


