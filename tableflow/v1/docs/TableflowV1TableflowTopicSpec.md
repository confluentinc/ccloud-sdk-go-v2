# TableflowV1TableflowTopicSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the Kafka topic for which Tableflow is enabled. | [optional] 
**Suspended** | Pointer to **bool** | Indicates whether the Tableflow should be suspended. The API allows setting it only to &#x60;false&#x60; i.e., to resume the Tableflow. Pausing the Tableflow on-demand is not currently supported. | [optional] 
**Config** | Pointer to [**TableflowV1TableFlowTopicConfigsSpec**](tableflow.v1.TableFlowTopicConfigsSpec.md) | The config for the Tableflow enabled topic | [optional] 
**Storage** | Pointer to [**TableflowV1TableflowTopicSpecStorageOneOf**](TableflowV1TableflowTopicSpecStorageOneOf.md) | The storage config | [optional] 
**TableFormats** | Pointer to **[]string** | The supported table formats for the Tableflow-enabled topic.  | [optional] [default to ["ICEBERG"]]
**Environment** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The environment to which the target Kafka cluster belongs. | [optional] 
**KafkaCluster** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The kafka cluster of the topic for which Tableflow is enabled | [optional] 

## Methods

### NewTableflowV1TableflowTopicSpec

`func NewTableflowV1TableflowTopicSpec() *TableflowV1TableflowTopicSpec`

NewTableflowV1TableflowTopicSpec instantiates a new TableflowV1TableflowTopicSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1TableflowTopicSpecWithDefaults

`func NewTableflowV1TableflowTopicSpecWithDefaults() *TableflowV1TableflowTopicSpec`

NewTableflowV1TableflowTopicSpecWithDefaults instantiates a new TableflowV1TableflowTopicSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *TableflowV1TableflowTopicSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TableflowV1TableflowTopicSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TableflowV1TableflowTopicSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TableflowV1TableflowTopicSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSuspended

`func (o *TableflowV1TableflowTopicSpec) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *TableflowV1TableflowTopicSpec) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *TableflowV1TableflowTopicSpec) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *TableflowV1TableflowTopicSpec) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetConfig

`func (o *TableflowV1TableflowTopicSpec) GetConfig() TableflowV1TableFlowTopicConfigsSpec`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *TableflowV1TableflowTopicSpec) GetConfigOk() (*TableflowV1TableFlowTopicConfigsSpec, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *TableflowV1TableflowTopicSpec) SetConfig(v TableflowV1TableFlowTopicConfigsSpec)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *TableflowV1TableflowTopicSpec) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStorage

`func (o *TableflowV1TableflowTopicSpec) GetStorage() TableflowV1TableflowTopicSpecStorageOneOf`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *TableflowV1TableflowTopicSpec) GetStorageOk() (*TableflowV1TableflowTopicSpecStorageOneOf, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *TableflowV1TableflowTopicSpec) SetStorage(v TableflowV1TableflowTopicSpecStorageOneOf)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *TableflowV1TableflowTopicSpec) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetTableFormats

`func (o *TableflowV1TableflowTopicSpec) GetTableFormats() []string`

GetTableFormats returns the TableFormats field if non-nil, zero value otherwise.

### GetTableFormatsOk

`func (o *TableflowV1TableflowTopicSpec) GetTableFormatsOk() (*[]string, bool)`

GetTableFormatsOk returns a tuple with the TableFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableFormats

`func (o *TableflowV1TableflowTopicSpec) SetTableFormats(v []string)`

SetTableFormats sets TableFormats field to given value.

### HasTableFormats

`func (o *TableflowV1TableflowTopicSpec) HasTableFormats() bool`

HasTableFormats returns a boolean if a field has been set.

### GetEnvironment

`func (o *TableflowV1TableflowTopicSpec) GetEnvironment() GlobalObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *TableflowV1TableflowTopicSpec) GetEnvironmentOk() (*GlobalObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *TableflowV1TableflowTopicSpec) SetEnvironment(v GlobalObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *TableflowV1TableflowTopicSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetKafkaCluster

`func (o *TableflowV1TableflowTopicSpec) GetKafkaCluster() EnvScopedObjectReference`

GetKafkaCluster returns the KafkaCluster field if non-nil, zero value otherwise.

### GetKafkaClusterOk

`func (o *TableflowV1TableflowTopicSpec) GetKafkaClusterOk() (*EnvScopedObjectReference, bool)`

GetKafkaClusterOk returns a tuple with the KafkaCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaCluster

`func (o *TableflowV1TableflowTopicSpec) SetKafkaCluster(v EnvScopedObjectReference)`

SetKafkaCluster sets KafkaCluster field to given value.

### HasKafkaCluster

`func (o *TableflowV1TableflowTopicSpec) HasKafkaCluster() bool`

HasKafkaCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


