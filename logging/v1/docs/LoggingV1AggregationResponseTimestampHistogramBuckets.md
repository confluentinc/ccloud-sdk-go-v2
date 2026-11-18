# LoggingV1AggregationResponseTimestampHistogramBuckets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Bucket start, RFC 3339 UTC. | [optional] 
**Count** | Pointer to **int64** | Number of matching documents in this bucket. | [optional] 
**AggregatedTerms** | Pointer to [**LoggingV1AggregationResponseTimestampHistogramAggregatedTerms**](LoggingV1AggregationResponseTimestampHistogramAggregatedTerms.md) |  | [optional] 

## Methods

### NewLoggingV1AggregationResponseTimestampHistogramBuckets

`func NewLoggingV1AggregationResponseTimestampHistogramBuckets() *LoggingV1AggregationResponseTimestampHistogramBuckets`

NewLoggingV1AggregationResponseTimestampHistogramBuckets instantiates a new LoggingV1AggregationResponseTimestampHistogramBuckets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingV1AggregationResponseTimestampHistogramBucketsWithDefaults

`func NewLoggingV1AggregationResponseTimestampHistogramBucketsWithDefaults() *LoggingV1AggregationResponseTimestampHistogramBuckets`

NewLoggingV1AggregationResponseTimestampHistogramBucketsWithDefaults instantiates a new LoggingV1AggregationResponseTimestampHistogramBuckets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetCount

`func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetAggregatedTerms

`func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) GetAggregatedTerms() LoggingV1AggregationResponseTimestampHistogramAggregatedTerms`

GetAggregatedTerms returns the AggregatedTerms field if non-nil, zero value otherwise.

### GetAggregatedTermsOk

`func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) GetAggregatedTermsOk() (*LoggingV1AggregationResponseTimestampHistogramAggregatedTerms, bool)`

GetAggregatedTermsOk returns a tuple with the AggregatedTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatedTerms

`func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) SetAggregatedTerms(v LoggingV1AggregationResponseTimestampHistogramAggregatedTerms)`

SetAggregatedTerms sets AggregatedTerms field to given value.

### HasAggregatedTerms

`func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) HasAggregatedTerms() bool`

HasAggregatedTerms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


