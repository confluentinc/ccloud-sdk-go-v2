# FloatType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nullable** | **bool** | Indicates whether values in this column can be null. | 
**Type** | **string** | The data type of the column. | 

## Methods

### NewFloatType

`func NewFloatType(nullable bool, type_ string, ) *FloatType`

NewFloatType instantiates a new FloatType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloatTypeWithDefaults

`func NewFloatTypeWithDefaults() *FloatType`

NewFloatTypeWithDefaults instantiates a new FloatType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNullable

`func (o *FloatType) GetNullable() bool`

GetNullable returns the Nullable field if non-nil, zero value otherwise.

### GetNullableOk

`func (o *FloatType) GetNullableOk() (*bool, bool)`

GetNullableOk returns a tuple with the Nullable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullable

`func (o *FloatType) SetNullable(v bool)`

SetNullable sets Nullable field to given value.


### GetType

`func (o *FloatType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FloatType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FloatType) SetType(v string)`

SetType sets Type field to given value.



### AsDataType

`func (s *FloatType) AsDataType() DataType`

Convenience method to wrap this instance of FloatType in DataType

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


