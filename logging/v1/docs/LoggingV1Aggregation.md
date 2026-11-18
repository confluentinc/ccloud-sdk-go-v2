# LoggingV1Aggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Field to bucket by. The service accepts only &#x60;timestamp&#x60;. | 
**Interval** | **string** | Bucket width, a number followed by s, m, h, d, w, M or y, for example &#x60;1h&#x60;. | 
**GroupBy** | **string** | A search key the resource type accepts, other than &#x60;search_text&#x60;, to sub-aggregate each bucket by. The service rejects an aggregation without a valid &#x60;group_by&#x60; with 422. | 

## Methods

### NewLoggingV1Aggregation

`func NewLoggingV1Aggregation(field string, interval string, groupBy string, ) *LoggingV1Aggregation`

NewLoggingV1Aggregation instantiates a new LoggingV1Aggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingV1AggregationWithDefaults

`func NewLoggingV1AggregationWithDefaults() *LoggingV1Aggregation`

NewLoggingV1AggregationWithDefaults instantiates a new LoggingV1Aggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *LoggingV1Aggregation) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *LoggingV1Aggregation) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *LoggingV1Aggregation) SetField(v string)`

SetField sets Field field to given value.


### GetInterval

`func (o *LoggingV1Aggregation) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *LoggingV1Aggregation) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *LoggingV1Aggregation) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetGroupBy

`func (o *LoggingV1Aggregation) GetGroupBy() string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *LoggingV1Aggregation) GetGroupByOk() (*string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *LoggingV1Aggregation) SetGroupBy(v string)`

SetGroupBy sets GroupBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


