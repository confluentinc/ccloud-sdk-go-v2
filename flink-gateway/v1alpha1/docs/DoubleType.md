# DoubleType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nullable** | **bool** | Indicates whether values in this column can be null. | 
**Type** | **string** | The data type of the column. | 

## Methods

### NewDoubleType

`func NewDoubleType(nullable bool, type_ string, ) *DoubleType`

NewDoubleType instantiates a new DoubleType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDoubleTypeWithDefaults

`func NewDoubleTypeWithDefaults() *DoubleType`

NewDoubleTypeWithDefaults instantiates a new DoubleType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNullable

`func (o *DoubleType) GetNullable() bool`

GetNullable returns the Nullable field if non-nil, zero value otherwise.

### GetNullableOk

`func (o *DoubleType) GetNullableOk() (*bool, bool)`

GetNullableOk returns a tuple with the Nullable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullable

`func (o *DoubleType) SetNullable(v bool)`

SetNullable sets Nullable field to given value.


### GetType

`func (o *DoubleType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DoubleType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DoubleType) SetType(v string)`

SetType sets Type field to given value.



### AsDataType

`func (s *DoubleType) AsDataType() DataType`

Convenience method to wrap this instance of DoubleType in DataType

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


