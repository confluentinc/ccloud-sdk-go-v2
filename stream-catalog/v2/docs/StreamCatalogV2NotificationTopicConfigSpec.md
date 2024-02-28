# StreamCatalogV2NotificationTopicConfigSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | Pointer to **string** | The kafka topic to write catalog notifications | [optional] [readonly] 
**Enabled** | Pointer to **bool** | Catalog notifications enablement | [optional] [readonly] 
**Environment** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The environment to which this belongs. | [optional] 
**KafkaCluster** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The kafka_cluster associated with this object. | [optional] 

## Methods

### NewStreamCatalogV2NotificationTopicConfigSpec

`func NewStreamCatalogV2NotificationTopicConfigSpec() *StreamCatalogV2NotificationTopicConfigSpec`

NewStreamCatalogV2NotificationTopicConfigSpec instantiates a new StreamCatalogV2NotificationTopicConfigSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamCatalogV2NotificationTopicConfigSpecWithDefaults

`func NewStreamCatalogV2NotificationTopicConfigSpecWithDefaults() *StreamCatalogV2NotificationTopicConfigSpec`

NewStreamCatalogV2NotificationTopicConfigSpecWithDefaults instantiates a new StreamCatalogV2NotificationTopicConfigSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *StreamCatalogV2NotificationTopicConfigSpec) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *StreamCatalogV2NotificationTopicConfigSpec) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *StreamCatalogV2NotificationTopicConfigSpec) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *StreamCatalogV2NotificationTopicConfigSpec) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetEnabled

`func (o *StreamCatalogV2NotificationTopicConfigSpec) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StreamCatalogV2NotificationTopicConfigSpec) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StreamCatalogV2NotificationTopicConfigSpec) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StreamCatalogV2NotificationTopicConfigSpec) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEnvironment

`func (o *StreamCatalogV2NotificationTopicConfigSpec) GetEnvironment() EnvScopedObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *StreamCatalogV2NotificationTopicConfigSpec) GetEnvironmentOk() (*EnvScopedObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *StreamCatalogV2NotificationTopicConfigSpec) SetEnvironment(v EnvScopedObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *StreamCatalogV2NotificationTopicConfigSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetKafkaCluster

`func (o *StreamCatalogV2NotificationTopicConfigSpec) GetKafkaCluster() EnvScopedObjectReference`

GetKafkaCluster returns the KafkaCluster field if non-nil, zero value otherwise.

### GetKafkaClusterOk

`func (o *StreamCatalogV2NotificationTopicConfigSpec) GetKafkaClusterOk() (*EnvScopedObjectReference, bool)`

GetKafkaClusterOk returns a tuple with the KafkaCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaCluster

`func (o *StreamCatalogV2NotificationTopicConfigSpec) SetKafkaCluster(v EnvScopedObjectReference)`

SetKafkaCluster sets KafkaCluster field to given value.

### HasKafkaCluster

`func (o *StreamCatalogV2NotificationTopicConfigSpec) HasKafkaCluster() bool`

HasKafkaCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


