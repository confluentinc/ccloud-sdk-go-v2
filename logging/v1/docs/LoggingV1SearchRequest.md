# LoggingV1SearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Crn** | **string** | Confluent Resource Name of the resource whose logs are requested. Its elements must match one of the shapes in the Resource types section above, and its organization must match the caller&#39;s. | 
**Source** | Pointer to **string** | The log source to query, for example cluster-health (see GET /logs/v1/sources). Optional; the service infers it from the crn&#39;s trailing element. Required when the crn&#39;s resource type has more than one log source. | [optional] 
**Search** | Pointer to **map[string]interface{}** | Field filters. Keys must be search keys the resource type accepts. A string value is a phrase match; an array of strings matches any value (OR). The reserved key &#x60;search_text&#x60; runs a phrase or wildcard match over the resource type&#39;s text fields. Keys and string values are limited to 250 characters; array elements are not length-checked. | [optional] 
**Sort** | Pointer to **string** | Sort order by event timestamp. Omitted sorts ascending. Any value other than &#x60;asc&#x60; or &#x60;desc&#x60; is rejected with 422. | [optional] [default to "asc"]
**Aggregation** | Pointer to [**LoggingV1Aggregation**](LoggingV1Aggregation.md) |  | [optional] 
**StartTime** | **string** | Inclusive window start, RFC 3339 in UTC with whole seconds and a &#x60;Z&#x60; suffix, as in the example. This is a string with a pattern, not &#x60;date-time&#x60;, so a generated client sends exactly this form and not a fractional second or a local offset. The service parses &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ssX&#x60; and rejects a fractional second or an offset with minutes such as &#x60;+00:00&#x60; with 422. It accepts an hour-only offset such as &#x60;+05&#x60;, and reads it as that offset, not as UTC, so always send &#x60;Z&#x60;. Starts older than the lookback are clamped to it. The generated client does not validate this format, so callers must send whole seconds and &#x60;Z&#x60; themselves. | 
**EndTime** | **string** | Inclusive window end, RFC 3339 in UTC with whole seconds and a &#x60;Z&#x60; suffix, with the same rules as &#x60;start_time&#x60;. An end older than the lookback returns an empty 200. The generated client does not validate this format, so callers must send whole seconds and &#x60;Z&#x60; themselves. | 

## Methods

### NewLoggingV1SearchRequest

`func NewLoggingV1SearchRequest(crn string, startTime string, endTime string, ) *LoggingV1SearchRequest`

NewLoggingV1SearchRequest instantiates a new LoggingV1SearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingV1SearchRequestWithDefaults

`func NewLoggingV1SearchRequestWithDefaults() *LoggingV1SearchRequest`

NewLoggingV1SearchRequestWithDefaults instantiates a new LoggingV1SearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrn

`func (o *LoggingV1SearchRequest) GetCrn() string`

GetCrn returns the Crn field if non-nil, zero value otherwise.

### GetCrnOk

`func (o *LoggingV1SearchRequest) GetCrnOk() (*string, bool)`

GetCrnOk returns a tuple with the Crn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrn

`func (o *LoggingV1SearchRequest) SetCrn(v string)`

SetCrn sets Crn field to given value.


### GetSource

`func (o *LoggingV1SearchRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LoggingV1SearchRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LoggingV1SearchRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *LoggingV1SearchRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSearch

`func (o *LoggingV1SearchRequest) GetSearch() map[string]interface{}`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *LoggingV1SearchRequest) GetSearchOk() (*map[string]interface{}, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *LoggingV1SearchRequest) SetSearch(v map[string]interface{})`

SetSearch sets Search field to given value.

### HasSearch

`func (o *LoggingV1SearchRequest) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetSort

`func (o *LoggingV1SearchRequest) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *LoggingV1SearchRequest) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *LoggingV1SearchRequest) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *LoggingV1SearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetAggregation

`func (o *LoggingV1SearchRequest) GetAggregation() LoggingV1Aggregation`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *LoggingV1SearchRequest) GetAggregationOk() (*LoggingV1Aggregation, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *LoggingV1SearchRequest) SetAggregation(v LoggingV1Aggregation)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *LoggingV1SearchRequest) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetStartTime

`func (o *LoggingV1SearchRequest) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *LoggingV1SearchRequest) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *LoggingV1SearchRequest) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *LoggingV1SearchRequest) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *LoggingV1SearchRequest) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *LoggingV1SearchRequest) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


