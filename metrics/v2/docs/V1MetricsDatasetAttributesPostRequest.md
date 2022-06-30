# V1MetricsDatasetAttributesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **string** | The metric that the label values are enumerated for. | 
**GroupBy** | **[]string** | The label(s) that the values are enumerated for. | 
**Filter** | Pointer to [**Filter**](Filter.md) |  | [optional] 
**Intervals** | Pointer to **[]string** | Defines the time range(s) for which available metrics will be listed. A time range is an ISO-8601 interval. When unspecified, the value defaults to the last hour before the request was made | [optional] 
**Limit** | Pointer to **int32** |  | [optional] [default to 100]

## Methods

### NewV1MetricsDatasetAttributesPostRequest

`func NewV1MetricsDatasetAttributesPostRequest(metric string, groupBy []string, ) *V1MetricsDatasetAttributesPostRequest`

NewV1MetricsDatasetAttributesPostRequest instantiates a new V1MetricsDatasetAttributesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1MetricsDatasetAttributesPostRequestWithDefaults

`func NewV1MetricsDatasetAttributesPostRequestWithDefaults() *V1MetricsDatasetAttributesPostRequest`

NewV1MetricsDatasetAttributesPostRequestWithDefaults instantiates a new V1MetricsDatasetAttributesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *V1MetricsDatasetAttributesPostRequest) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *V1MetricsDatasetAttributesPostRequest) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *V1MetricsDatasetAttributesPostRequest) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetGroupBy

`func (o *V1MetricsDatasetAttributesPostRequest) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *V1MetricsDatasetAttributesPostRequest) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *V1MetricsDatasetAttributesPostRequest) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.


### GetFilter

`func (o *V1MetricsDatasetAttributesPostRequest) GetFilter() Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *V1MetricsDatasetAttributesPostRequest) GetFilterOk() (*Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *V1MetricsDatasetAttributesPostRequest) SetFilter(v Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *V1MetricsDatasetAttributesPostRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetIntervals

`func (o *V1MetricsDatasetAttributesPostRequest) GetIntervals() []string`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *V1MetricsDatasetAttributesPostRequest) GetIntervalsOk() (*[]string, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *V1MetricsDatasetAttributesPostRequest) SetIntervals(v []string)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *V1MetricsDatasetAttributesPostRequest) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.

### GetLimit

`func (o *V1MetricsDatasetAttributesPostRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *V1MetricsDatasetAttributesPostRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *V1MetricsDatasetAttributesPostRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *V1MetricsDatasetAttributesPostRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


