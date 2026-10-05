# SqlV1IntervalExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | **int32** | Numeric value of the time interval. | 
**TimeUnit** | **string** | Unit of time for the interval. | 

## Methods

### NewSqlV1IntervalExpression

`func NewSqlV1IntervalExpression(interval int32, timeUnit string, ) *SqlV1IntervalExpression`

NewSqlV1IntervalExpression instantiates a new SqlV1IntervalExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1IntervalExpressionWithDefaults

`func NewSqlV1IntervalExpressionWithDefaults() *SqlV1IntervalExpression`

NewSqlV1IntervalExpressionWithDefaults instantiates a new SqlV1IntervalExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *SqlV1IntervalExpression) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *SqlV1IntervalExpression) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *SqlV1IntervalExpression) SetInterval(v int32)`

SetInterval sets Interval field to given value.


### GetTimeUnit

`func (o *SqlV1IntervalExpression) GetTimeUnit() string`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *SqlV1IntervalExpression) GetTimeUnitOk() (*string, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *SqlV1IntervalExpression) SetTimeUnit(v string)`

SetTimeUnit sets TimeUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


