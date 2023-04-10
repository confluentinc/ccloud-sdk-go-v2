# IntervalDayTimeType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nullable** | **bool** | Indicates whether values in this column can be null. | 
**Type** | **string** | The data type of the column. | 
**Resolution** | Pointer to **string** | The resolution for the interval | [optional] 

## Methods

### NewIntervalDayTimeType

`func NewIntervalDayTimeType(nullable bool, type_ string, ) *IntervalDayTimeType`

NewIntervalDayTimeType instantiates a new IntervalDayTimeType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntervalDayTimeTypeWithDefaults

`func NewIntervalDayTimeTypeWithDefaults() *IntervalDayTimeType`

NewIntervalDayTimeTypeWithDefaults instantiates a new IntervalDayTimeType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNullable

`func (o *IntervalDayTimeType) GetNullable() bool`

GetNullable returns the Nullable field if non-nil, zero value otherwise.

### GetNullableOk

`func (o *IntervalDayTimeType) GetNullableOk() (*bool, bool)`

GetNullableOk returns a tuple with the Nullable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullable

`func (o *IntervalDayTimeType) SetNullable(v bool)`

SetNullable sets Nullable field to given value.


### GetType

`func (o *IntervalDayTimeType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntervalDayTimeType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntervalDayTimeType) SetType(v string)`

SetType sets Type field to given value.


### GetResolution

`func (o *IntervalDayTimeType) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *IntervalDayTimeType) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *IntervalDayTimeType) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *IntervalDayTimeType) HasResolution() bool`

HasResolution returns a boolean if a field has been set.


### AsDataType

`func (s *IntervalDayTimeType) AsDataType() DataType`

Convenience method to wrap this instance of IntervalDayTimeType in DataType

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


