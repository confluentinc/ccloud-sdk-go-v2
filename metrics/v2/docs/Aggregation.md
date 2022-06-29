# Aggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **string** | The metric to aggregate. | 
**Agg** | Pointer to [**NullableAggregationFunction**](AggregationFunction.md) | The aggregation function for the label buckets defined in &#x60;group_by&#x60;. | [optional] 

## Methods

### NewAggregation

`func NewAggregation(metric string, ) *Aggregation`

NewAggregation instantiates a new Aggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationWithDefaults

`func NewAggregationWithDefaults() *Aggregation`

NewAggregationWithDefaults instantiates a new Aggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *Aggregation) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *Aggregation) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *Aggregation) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetAgg

`func (o *Aggregation) GetAgg() AggregationFunction`

GetAgg returns the Agg field if non-nil, zero value otherwise.

### GetAggOk

`func (o *Aggregation) GetAggOk() (*AggregationFunction, bool)`

GetAggOk returns a tuple with the Agg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgg

`func (o *Aggregation) SetAgg(v AggregationFunction)`

SetAgg sets Agg field to given value.

### HasAgg

`func (o *Aggregation) HasAgg() bool`

HasAgg returns a boolean if a field has been set.

### SetAggNil

`func (o *Aggregation) SetAggNil(b bool)`

 SetAggNil sets the value for Agg to be an explicit nil

### UnsetAgg
`func (o *Aggregation) UnsetAgg()`

UnsetAgg ensures that no value is present for Agg, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


