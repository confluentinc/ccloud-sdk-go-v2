# AttributesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | Pointer to **string** | The metric that the label values are enumerated for. | [optional] 
**GroupBy** | **[]string** | The label(s) that the values are enumerated for. | 
**Filter** | Pointer to [**Filter**](Filter.md) |  | [optional] 
**Intervals** | Pointer to **[]string** | Defines the time range(s) for which available metrics will be listed. A time range is an ISO-8601 interval. When unspecified, the value defaults to the last hour before the request was made | [optional] 
**Limit** | Pointer to **int32** |  | [optional] [default to 100]

## Methods

### NewAttributesRequest

`func NewAttributesRequest(groupBy []string, ) *AttributesRequest`

NewAttributesRequest instantiates a new AttributesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributesRequestWithDefaults

`func NewAttributesRequestWithDefaults() *AttributesRequest`

NewAttributesRequestWithDefaults instantiates a new AttributesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *AttributesRequest) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *AttributesRequest) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *AttributesRequest) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *AttributesRequest) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetGroupBy

`func (o *AttributesRequest) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *AttributesRequest) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *AttributesRequest) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.


### GetFilter

`func (o *AttributesRequest) GetFilter() Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AttributesRequest) GetFilterOk() (*Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AttributesRequest) SetFilter(v Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AttributesRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetIntervals

`func (o *AttributesRequest) GetIntervals() []string`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *AttributesRequest) GetIntervalsOk() (*[]string, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *AttributesRequest) SetIntervals(v []string)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *AttributesRequest) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.

### GetLimit

`func (o *AttributesRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AttributesRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AttributesRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AttributesRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


