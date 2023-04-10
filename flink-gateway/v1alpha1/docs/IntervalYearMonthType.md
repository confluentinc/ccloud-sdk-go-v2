# IntervalYearMonthType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nullable** | **bool** | Indicates whether values in this column can be null. | 
**Type** | **string** | The data type of the column. | 
**Resolution** | Pointer to **string** | The resolution for the interval | [optional] 

## Methods

### NewIntervalYearMonthType

`func NewIntervalYearMonthType(nullable bool, type_ string, ) *IntervalYearMonthType`

NewIntervalYearMonthType instantiates a new IntervalYearMonthType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntervalYearMonthTypeWithDefaults

`func NewIntervalYearMonthTypeWithDefaults() *IntervalYearMonthType`

NewIntervalYearMonthTypeWithDefaults instantiates a new IntervalYearMonthType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNullable

`func (o *IntervalYearMonthType) GetNullable() bool`

GetNullable returns the Nullable field if non-nil, zero value otherwise.

### GetNullableOk

`func (o *IntervalYearMonthType) GetNullableOk() (*bool, bool)`

GetNullableOk returns a tuple with the Nullable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullable

`func (o *IntervalYearMonthType) SetNullable(v bool)`

SetNullable sets Nullable field to given value.


### GetType

`func (o *IntervalYearMonthType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntervalYearMonthType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntervalYearMonthType) SetType(v string)`

SetType sets Type field to given value.


### GetResolution

`func (o *IntervalYearMonthType) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *IntervalYearMonthType) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *IntervalYearMonthType) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *IntervalYearMonthType) HasResolution() bool`

HasResolution returns a boolean if a field has been set.


### AsDataType

`func (s *IntervalYearMonthType) AsDataType() DataType`

Convenience method to wrap this instance of IntervalYearMonthType in DataType

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


