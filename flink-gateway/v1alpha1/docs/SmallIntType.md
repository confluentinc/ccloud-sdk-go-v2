# SmallIntType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nullable** | **bool** | Indicates whether values in this column can be null. | 
**Type** | **string** | The data type of the column. | 

## Methods

### NewSmallIntType

`func NewSmallIntType(nullable bool, type_ string, ) *SmallIntType`

NewSmallIntType instantiates a new SmallIntType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmallIntTypeWithDefaults

`func NewSmallIntTypeWithDefaults() *SmallIntType`

NewSmallIntTypeWithDefaults instantiates a new SmallIntType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNullable

`func (o *SmallIntType) GetNullable() bool`

GetNullable returns the Nullable field if non-nil, zero value otherwise.

### GetNullableOk

`func (o *SmallIntType) GetNullableOk() (*bool, bool)`

GetNullableOk returns a tuple with the Nullable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullable

`func (o *SmallIntType) SetNullable(v bool)`

SetNullable sets Nullable field to given value.


### GetType

`func (o *SmallIntType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SmallIntType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SmallIntType) SetType(v string)`

SetType sets Type field to given value.



### AsDataType

`func (s *SmallIntType) AsDataType() DataType`

Convenience method to wrap this instance of SmallIntType in DataType

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


