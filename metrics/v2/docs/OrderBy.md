# OrderBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **string** |  | 
**Agg** | Pointer to [**AggregationFunction**](AggregationFunction.md) |  | [optional] 
**Order** | Pointer to **string** |  | [optional] [default to "DESCENDING"]

## Methods

### NewOrderBy

`func NewOrderBy(metric string, ) *OrderBy`

NewOrderBy instantiates a new OrderBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderByWithDefaults

`func NewOrderByWithDefaults() *OrderBy`

NewOrderByWithDefaults instantiates a new OrderBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *OrderBy) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *OrderBy) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *OrderBy) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetAgg

`func (o *OrderBy) GetAgg() AggregationFunction`

GetAgg returns the Agg field if non-nil, zero value otherwise.

### GetAggOk

`func (o *OrderBy) GetAggOk() (*AggregationFunction, bool)`

GetAggOk returns a tuple with the Agg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgg

`func (o *OrderBy) SetAgg(v AggregationFunction)`

SetAgg sets Agg field to given value.

### HasAgg

`func (o *OrderBy) HasAgg() bool`

HasAgg returns a boolean if a field has been set.

### GetOrder

`func (o *OrderBy) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *OrderBy) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *OrderBy) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *OrderBy) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


