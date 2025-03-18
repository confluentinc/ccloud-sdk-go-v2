# TableflowV1TableflowTopicSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suspended** | Pointer to **bool** | Indicates whether the Tableflow should be suspended. The API allows setting it only to &#x60;false&#x60; i.e., to resume the Tableflow. Pausing the Tableflow on-demand is not currently supported. | [optional] 
**Config** | Pointer to [**TableflowV1TableFlowTopicConfigsSpec**](tableflow.v1.TableFlowTopicConfigsSpec.md) | The config for the Tableflow enabled topic | [optional] 
**TableFormats** | Pointer to **[]string** | The supported table formats for the Tableflow-enabled topic.  | [optional] [default to ["ICEBERG"]]
**Environment** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The environment to which the target Kafka cluster belongs. | [optional] 
**KafkaCluster** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The kafka cluster of the topic for which Tableflow is enabled | [optional] 

## Methods

### NewTableflowV1TableflowTopicSpecUpdate

`func NewTableflowV1TableflowTopicSpecUpdate() *TableflowV1TableflowTopicSpecUpdate`

NewTableflowV1TableflowTopicSpecUpdate instantiates a new TableflowV1TableflowTopicSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1TableflowTopicSpecUpdateWithDefaults

`func NewTableflowV1TableflowTopicSpecUpdateWithDefaults() *TableflowV1TableflowTopicSpecUpdate`

NewTableflowV1TableflowTopicSpecUpdateWithDefaults instantiates a new TableflowV1TableflowTopicSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuspended

`func (o *TableflowV1TableflowTopicSpecUpdate) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *TableflowV1TableflowTopicSpecUpdate) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *TableflowV1TableflowTopicSpecUpdate) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *TableflowV1TableflowTopicSpecUpdate) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetConfig

`func (o *TableflowV1TableflowTopicSpecUpdate) GetConfig() TableflowV1TableFlowTopicConfigsSpec`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *TableflowV1TableflowTopicSpecUpdate) GetConfigOk() (*TableflowV1TableFlowTopicConfigsSpec, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *TableflowV1TableflowTopicSpecUpdate) SetConfig(v TableflowV1TableFlowTopicConfigsSpec)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *TableflowV1TableflowTopicSpecUpdate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTableFormats

`func (o *TableflowV1TableflowTopicSpecUpdate) GetTableFormats() []string`

GetTableFormats returns the TableFormats field if non-nil, zero value otherwise.

### GetTableFormatsOk

`func (o *TableflowV1TableflowTopicSpecUpdate) GetTableFormatsOk() (*[]string, bool)`

GetTableFormatsOk returns a tuple with the TableFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableFormats

`func (o *TableflowV1TableflowTopicSpecUpdate) SetTableFormats(v []string)`

SetTableFormats sets TableFormats field to given value.

### HasTableFormats

`func (o *TableflowV1TableflowTopicSpecUpdate) HasTableFormats() bool`

HasTableFormats returns a boolean if a field has been set.

### GetEnvironment

`func (o *TableflowV1TableflowTopicSpecUpdate) GetEnvironment() GlobalObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *TableflowV1TableflowTopicSpecUpdate) GetEnvironmentOk() (*GlobalObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *TableflowV1TableflowTopicSpecUpdate) SetEnvironment(v GlobalObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *TableflowV1TableflowTopicSpecUpdate) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetKafkaCluster

`func (o *TableflowV1TableflowTopicSpecUpdate) GetKafkaCluster() EnvScopedObjectReference`

GetKafkaCluster returns the KafkaCluster field if non-nil, zero value otherwise.

### GetKafkaClusterOk

`func (o *TableflowV1TableflowTopicSpecUpdate) GetKafkaClusterOk() (*EnvScopedObjectReference, bool)`

GetKafkaClusterOk returns a tuple with the KafkaCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaCluster

`func (o *TableflowV1TableflowTopicSpecUpdate) SetKafkaCluster(v EnvScopedObjectReference)`

SetKafkaCluster sets KafkaCluster field to given value.

### HasKafkaCluster

`func (o *TableflowV1TableflowTopicSpecUpdate) HasKafkaCluster() bool`

HasKafkaCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


