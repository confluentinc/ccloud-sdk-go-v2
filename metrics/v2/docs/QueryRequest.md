# QueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregations** | [**[]Aggregation**](Aggregation.md) | Specifies which metrics to query and the aggregation operator to apply across the &#x60;group_by&#x60; labels. **Currently, only one aggregation per request is supported.** | 
**GroupBy** | Pointer to **[]string** | Specifies how data gets bucketed by label(s); see [here](#section/Object-Model/Labels) on using labels for grouping query results.  | [optional] 
**Granularity** | [**Granularity**](Granularity.md) |  | 
**Filter** | Pointer to [**Filter**](Filter.md) |  | [optional] 
**OrderBy** | Pointer to [**[]OrderBy**](OrderBy.md) | Sort ordering for result groups. **Only valid for &#x60;granularity: \&quot;ALL\&quot;&#x60;.** If not specified, defaults to the first &#x60;aggregation&#x60; in descending order.  Note that this ordering applies to the groups. Within a group (or for ungrouped results), data points are always ordered by &#x60;timestamp&#x60; in descending order.  | [optional] 
**Intervals** | **[]string** | Defines the time range(s) that the query runs over. A time range is an ISO-8601 interval.  The keyword &#x60;now&#x60; can be used in place of a timestamp to refer to the current time. Offset and truncation modifiers can be also be applied to the &#x60;now&#x60; expression:  | Modifier | Syntax | Examples | | --- | --- | --- | | Offset | &#x60;(+\\|-)&lt;amount&gt;(m\\|h\\|d)&#x60; | &#x60;-2m&#x60; (minus 2 minutes)&lt;br/&gt;&#x60;-1h&#x60; (minus 1 hour) | | Truncation | &#x60;\\|(m\\|h\\|d)&#x60; | &#x60;\\|m&#x60; (round down to start of minute)&lt;br/&gt;&#x60;\\|h&#x60; (round down to start of hour) |  All hour/day truncation is performed against the UTC timezone.  If &#x60;now&#x60; is &#x60;2020-01-01T02:13:27Z&#x60;, some examples are: * &#x60;now-2m|m&#x60;: &#x60;now&#x60; minus 2 minutes, truncated to start of minute. &lt;br/&gt;Resolves to &#x60;2020-01-01T02:11:00Z&#x60; * &#x60;now|h&#x60;: &#x60;now&#x60; truncated to start of hour. &lt;br/&gt;Resolves to &#x60;2020-01-01T02:00:00Z&#x60; * &#x60;now-1d|d&#x60;: &#x60;now&#x60; minus 1 day, truncated to start of day. &lt;br/&gt;Resolves to &#x60;2019-12-31T00:00:00Z&#x60;  When using &#x60;now&#x60;, it is recommended to apply a negative offset to avoid incomplete data (see [metric availability delays](#section/Client-Considerations-and-Best-Practices/Metric-Data-Latency)) and align to minute boundaries (e.g. &#x60;now-2m|m&#x60;).  | 
**Limit** | Pointer to **int32** | The maximum number of _groups_ to return. The maximum number of data points in the response is equal to &#x60;limit * (interval / granularity)&#x60;. For example, with an interval of 1 day, granularity of &#x60;PT1H&#x60;, and limit of &#x60;2&#x60; there will be a maximum of 48 data points in the response (24 for each group). | [optional] [default to 100]
**Format** | Pointer to [**ResponseFormat**](ResponseFormat.md) |  | [optional] [default to FLAT]

## Methods

### NewQueryRequest

`func NewQueryRequest(aggregations []Aggregation, granularity Granularity, intervals []string, ) *QueryRequest`

NewQueryRequest instantiates a new QueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryRequestWithDefaults

`func NewQueryRequestWithDefaults() *QueryRequest`

NewQueryRequestWithDefaults instantiates a new QueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregations

`func (o *QueryRequest) GetAggregations() []Aggregation`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *QueryRequest) GetAggregationsOk() (*[]Aggregation, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *QueryRequest) SetAggregations(v []Aggregation)`

SetAggregations sets Aggregations field to given value.


### GetGroupBy

`func (o *QueryRequest) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *QueryRequest) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *QueryRequest) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *QueryRequest) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetGranularity

`func (o *QueryRequest) GetGranularity() Granularity`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *QueryRequest) GetGranularityOk() (*Granularity, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *QueryRequest) SetGranularity(v Granularity)`

SetGranularity sets Granularity field to given value.


### GetFilter

`func (o *QueryRequest) GetFilter() Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *QueryRequest) GetFilterOk() (*Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *QueryRequest) SetFilter(v Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *QueryRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetOrderBy

`func (o *QueryRequest) GetOrderBy() []OrderBy`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *QueryRequest) GetOrderByOk() (*[]OrderBy, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *QueryRequest) SetOrderBy(v []OrderBy)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *QueryRequest) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetIntervals

`func (o *QueryRequest) GetIntervals() []string`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *QueryRequest) GetIntervalsOk() (*[]string, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *QueryRequest) SetIntervals(v []string)`

SetIntervals sets Intervals field to given value.


### GetLimit

`func (o *QueryRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QueryRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QueryRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *QueryRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetFormat

`func (o *QueryRequest) GetFormat() ResponseFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *QueryRequest) GetFormatOk() (*ResponseFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *QueryRequest) SetFormat(v ResponseFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *QueryRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


