# RtceV1RtceTopicSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A model-readable description of the RTCE topic. | [optional] 
**Environment** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The environment to which the target Kafka cluster belongs. | [optional] 
**KafkaCluster** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The Kafka cluster containing the topic to be materialized. | [optional] 

## Methods

### NewRtceV1RtceTopicSpecUpdate

`func NewRtceV1RtceTopicSpecUpdate() *RtceV1RtceTopicSpecUpdate`

NewRtceV1RtceTopicSpecUpdate instantiates a new RtceV1RtceTopicSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRtceV1RtceTopicSpecUpdateWithDefaults

`func NewRtceV1RtceTopicSpecUpdateWithDefaults() *RtceV1RtceTopicSpecUpdate`

NewRtceV1RtceTopicSpecUpdateWithDefaults instantiates a new RtceV1RtceTopicSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RtceV1RtceTopicSpecUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RtceV1RtceTopicSpecUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RtceV1RtceTopicSpecUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RtceV1RtceTopicSpecUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *RtceV1RtceTopicSpecUpdate) GetEnvironment() EnvScopedObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *RtceV1RtceTopicSpecUpdate) GetEnvironmentOk() (*EnvScopedObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *RtceV1RtceTopicSpecUpdate) SetEnvironment(v EnvScopedObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *RtceV1RtceTopicSpecUpdate) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetKafkaCluster

`func (o *RtceV1RtceTopicSpecUpdate) GetKafkaCluster() EnvScopedObjectReference`

GetKafkaCluster returns the KafkaCluster field if non-nil, zero value otherwise.

### GetKafkaClusterOk

`func (o *RtceV1RtceTopicSpecUpdate) GetKafkaClusterOk() (*EnvScopedObjectReference, bool)`

GetKafkaClusterOk returns a tuple with the KafkaCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaCluster

`func (o *RtceV1RtceTopicSpecUpdate) SetKafkaCluster(v EnvScopedObjectReference)`

SetKafkaCluster sets KafkaCluster field to given value.

### HasKafkaCluster

`func (o *RtceV1RtceTopicSpecUpdate) HasKafkaCluster() bool`

HasKafkaCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


