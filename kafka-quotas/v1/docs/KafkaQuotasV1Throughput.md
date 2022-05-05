# KafkaQuotasV1Throughput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IngressByteRate** | Pointer to **string** | ingress throughput limit in bytes per sec | [optional] 
**EgressByteRate** | Pointer to **string** | egress throughput limit in bytes per sec | [optional] 

## Methods

### NewKafkaQuotasV1Throughput

`func NewKafkaQuotasV1Throughput() *KafkaQuotasV1Throughput`

NewKafkaQuotasV1Throughput instantiates a new KafkaQuotasV1Throughput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaQuotasV1ThroughputWithDefaults

`func NewKafkaQuotasV1ThroughputWithDefaults() *KafkaQuotasV1Throughput`

NewKafkaQuotasV1ThroughputWithDefaults instantiates a new KafkaQuotasV1Throughput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngressByteRate

`func (o *KafkaQuotasV1Throughput) GetIngressByteRate() string`

GetIngressByteRate returns the IngressByteRate field if non-nil, zero value otherwise.

### GetIngressByteRateOk

`func (o *KafkaQuotasV1Throughput) GetIngressByteRateOk() (*string, bool)`

GetIngressByteRateOk returns a tuple with the IngressByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressByteRate

`func (o *KafkaQuotasV1Throughput) SetIngressByteRate(v string)`

SetIngressByteRate sets IngressByteRate field to given value.

### HasIngressByteRate

`func (o *KafkaQuotasV1Throughput) HasIngressByteRate() bool`

HasIngressByteRate returns a boolean if a field has been set.

### GetEgressByteRate

`func (o *KafkaQuotasV1Throughput) GetEgressByteRate() string`

GetEgressByteRate returns the EgressByteRate field if non-nil, zero value otherwise.

### GetEgressByteRateOk

`func (o *KafkaQuotasV1Throughput) GetEgressByteRateOk() (*string, bool)`

GetEgressByteRateOk returns a tuple with the EgressByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressByteRate

`func (o *KafkaQuotasV1Throughput) SetEgressByteRate(v string)`

SetEgressByteRate sets EgressByteRate field to given value.

### HasEgressByteRate

`func (o *KafkaQuotasV1Throughput) HasEgressByteRate() bool`

HasEgressByteRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


