# RtceV1RtceTopicSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cloud** | Pointer to **string** | The cloud provider where the RTCE topic is deployed. | [optional] 
**Region** | Pointer to **string** | The cloud region where the RTCE topic is deployed. | [optional] 
**TopicName** | Pointer to **string** | The Kafka topic name containing the data for the RTCE topic. | [optional] 
**Description** | Pointer to **string** | A model-readable description of the RTCE topic. | [optional] 
**Environment** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The environment to which the target Kafka cluster belongs. | [optional] 
**KafkaCluster** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The Kafka cluster containing the topic to be materialized. | [optional] 

## Methods

### NewRtceV1RtceTopicSpec

`func NewRtceV1RtceTopicSpec() *RtceV1RtceTopicSpec`

NewRtceV1RtceTopicSpec instantiates a new RtceV1RtceTopicSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRtceV1RtceTopicSpecWithDefaults

`func NewRtceV1RtceTopicSpecWithDefaults() *RtceV1RtceTopicSpec`

NewRtceV1RtceTopicSpecWithDefaults instantiates a new RtceV1RtceTopicSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloud

`func (o *RtceV1RtceTopicSpec) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *RtceV1RtceTopicSpec) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *RtceV1RtceTopicSpec) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *RtceV1RtceTopicSpec) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetRegion

`func (o *RtceV1RtceTopicSpec) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RtceV1RtceTopicSpec) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RtceV1RtceTopicSpec) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *RtceV1RtceTopicSpec) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetTopicName

`func (o *RtceV1RtceTopicSpec) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *RtceV1RtceTopicSpec) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *RtceV1RtceTopicSpec) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.

### HasTopicName

`func (o *RtceV1RtceTopicSpec) HasTopicName() bool`

HasTopicName returns a boolean if a field has been set.

### GetDescription

`func (o *RtceV1RtceTopicSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RtceV1RtceTopicSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RtceV1RtceTopicSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RtceV1RtceTopicSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *RtceV1RtceTopicSpec) GetEnvironment() EnvScopedObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *RtceV1RtceTopicSpec) GetEnvironmentOk() (*EnvScopedObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *RtceV1RtceTopicSpec) SetEnvironment(v EnvScopedObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *RtceV1RtceTopicSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetKafkaCluster

`func (o *RtceV1RtceTopicSpec) GetKafkaCluster() EnvScopedObjectReference`

GetKafkaCluster returns the KafkaCluster field if non-nil, zero value otherwise.

### GetKafkaClusterOk

`func (o *RtceV1RtceTopicSpec) GetKafkaClusterOk() (*EnvScopedObjectReference, bool)`

GetKafkaClusterOk returns a tuple with the KafkaCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaCluster

`func (o *RtceV1RtceTopicSpec) SetKafkaCluster(v EnvScopedObjectReference)`

SetKafkaCluster sets KafkaCluster field to given value.

### HasKafkaCluster

`func (o *RtceV1RtceTopicSpec) HasKafkaCluster() bool`

HasKafkaCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


