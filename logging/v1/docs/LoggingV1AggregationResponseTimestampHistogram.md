# LoggingV1AggregationResponseTimestampHistogram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | Pointer to [**[]LoggingV1AggregationResponseTimestampHistogramBuckets**](LoggingV1AggregationResponseTimestampHistogramBuckets.md) | One entry per time bucket, in bucket order. Empty buckets are omitted. | [optional] 

## Methods

### NewLoggingV1AggregationResponseTimestampHistogram

`func NewLoggingV1AggregationResponseTimestampHistogram() *LoggingV1AggregationResponseTimestampHistogram`

NewLoggingV1AggregationResponseTimestampHistogram instantiates a new LoggingV1AggregationResponseTimestampHistogram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingV1AggregationResponseTimestampHistogramWithDefaults

`func NewLoggingV1AggregationResponseTimestampHistogramWithDefaults() *LoggingV1AggregationResponseTimestampHistogram`

NewLoggingV1AggregationResponseTimestampHistogramWithDefaults instantiates a new LoggingV1AggregationResponseTimestampHistogram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *LoggingV1AggregationResponseTimestampHistogram) GetBuckets() []LoggingV1AggregationResponseTimestampHistogramBuckets`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *LoggingV1AggregationResponseTimestampHistogram) GetBucketsOk() (*[]LoggingV1AggregationResponseTimestampHistogramBuckets, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *LoggingV1AggregationResponseTimestampHistogram) SetBuckets(v []LoggingV1AggregationResponseTimestampHistogramBuckets)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *LoggingV1AggregationResponseTimestampHistogram) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


