# MultisetTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The data type of the column. | [optional] 
**MultisetElementType** | [**DataType**](DataType.md) | The type of the elements in the multiset. | 

## Methods

### NewMultisetTypeAllOf

`func NewMultisetTypeAllOf(multisetElementType DataType, ) *MultisetTypeAllOf`

NewMultisetTypeAllOf instantiates a new MultisetTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultisetTypeAllOfWithDefaults

`func NewMultisetTypeAllOfWithDefaults() *MultisetTypeAllOf`

NewMultisetTypeAllOfWithDefaults instantiates a new MultisetTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MultisetTypeAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MultisetTypeAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MultisetTypeAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MultisetTypeAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMultisetElementType

`func (o *MultisetTypeAllOf) GetMultisetElementType() DataType`

GetMultisetElementType returns the MultisetElementType field if non-nil, zero value otherwise.

### GetMultisetElementTypeOk

`func (o *MultisetTypeAllOf) GetMultisetElementTypeOk() (*DataType, bool)`

GetMultisetElementTypeOk returns a tuple with the MultisetElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultisetElementType

`func (o *MultisetTypeAllOf) SetMultisetElementType(v DataType)`

SetMultisetElementType sets MultisetElementType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


