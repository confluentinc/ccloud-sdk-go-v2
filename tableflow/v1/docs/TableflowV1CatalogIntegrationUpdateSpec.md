# TableflowV1CatalogIntegrationUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the catalog integration | [optional] 
**Suspended** | Pointer to **bool** | Indicates whether the Catalog Integration should be suspended. | [optional] 
**Config** | Pointer to [**TableflowV1CatalogIntegrationUpdateSpecConfigOneOf**](TableflowV1CatalogIntegrationUpdateSpecConfigOneOf.md) | The integration config | [optional] 
**Environment** | [**GlobalObjectReference**](GlobalObjectReference.md) | The environment to which the target Kafka cluster belongs. | 
**KafkaCluster** | [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The kafka cluster of the topic for which Tableflow is enabled | 

## Methods

### NewTableflowV1CatalogIntegrationUpdateSpec

`func NewTableflowV1CatalogIntegrationUpdateSpec(environment GlobalObjectReference, kafkaCluster EnvScopedObjectReference, ) *TableflowV1CatalogIntegrationUpdateSpec`

NewTableflowV1CatalogIntegrationUpdateSpec instantiates a new TableflowV1CatalogIntegrationUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1CatalogIntegrationUpdateSpecWithDefaults

`func NewTableflowV1CatalogIntegrationUpdateSpecWithDefaults() *TableflowV1CatalogIntegrationUpdateSpec`

NewTableflowV1CatalogIntegrationUpdateSpecWithDefaults instantiates a new TableflowV1CatalogIntegrationUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *TableflowV1CatalogIntegrationUpdateSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TableflowV1CatalogIntegrationUpdateSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TableflowV1CatalogIntegrationUpdateSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TableflowV1CatalogIntegrationUpdateSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSuspended

`func (o *TableflowV1CatalogIntegrationUpdateSpec) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *TableflowV1CatalogIntegrationUpdateSpec) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *TableflowV1CatalogIntegrationUpdateSpec) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *TableflowV1CatalogIntegrationUpdateSpec) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetConfig

`func (o *TableflowV1CatalogIntegrationUpdateSpec) GetConfig() TableflowV1CatalogIntegrationUpdateSpecConfigOneOf`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *TableflowV1CatalogIntegrationUpdateSpec) GetConfigOk() (*TableflowV1CatalogIntegrationUpdateSpecConfigOneOf, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *TableflowV1CatalogIntegrationUpdateSpec) SetConfig(v TableflowV1CatalogIntegrationUpdateSpecConfigOneOf)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *TableflowV1CatalogIntegrationUpdateSpec) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnvironment

`func (o *TableflowV1CatalogIntegrationUpdateSpec) GetEnvironment() GlobalObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *TableflowV1CatalogIntegrationUpdateSpec) GetEnvironmentOk() (*GlobalObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *TableflowV1CatalogIntegrationUpdateSpec) SetEnvironment(v GlobalObjectReference)`

SetEnvironment sets Environment field to given value.


### GetKafkaCluster

`func (o *TableflowV1CatalogIntegrationUpdateSpec) GetKafkaCluster() EnvScopedObjectReference`

GetKafkaCluster returns the KafkaCluster field if non-nil, zero value otherwise.

### GetKafkaClusterOk

`func (o *TableflowV1CatalogIntegrationUpdateSpec) GetKafkaClusterOk() (*EnvScopedObjectReference, bool)`

GetKafkaClusterOk returns a tuple with the KafkaCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaCluster

`func (o *TableflowV1CatalogIntegrationUpdateSpec) SetKafkaCluster(v EnvScopedObjectReference)`

SetKafkaCluster sets KafkaCluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


