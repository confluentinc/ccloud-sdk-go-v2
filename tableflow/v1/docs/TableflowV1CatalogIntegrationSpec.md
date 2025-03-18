# TableflowV1CatalogIntegrationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the catalog integration | [optional] 
**Suspended** | Pointer to **bool** | Indicates whether the Catalog Integration should be suspended. The API allows setting it only to &#x60;false&#x60; i.e., to resume the Catalog Integration. Pausing the Catalog Integration on-demand is not currently supported. | [optional] 
**Config** | Pointer to [**TableflowV1CatalogIntegrationSpecConfigOneOf**](TableflowV1CatalogIntegrationSpecConfigOneOf.md) | The integration config | [optional] 
**Environment** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The environment to which the target Kafka cluster belongs. | [optional] 
**KafkaCluster** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The kafka cluster of the topic for which Tableflow is enabled | [optional] 

## Methods

### NewTableflowV1CatalogIntegrationSpec

`func NewTableflowV1CatalogIntegrationSpec() *TableflowV1CatalogIntegrationSpec`

NewTableflowV1CatalogIntegrationSpec instantiates a new TableflowV1CatalogIntegrationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1CatalogIntegrationSpecWithDefaults

`func NewTableflowV1CatalogIntegrationSpecWithDefaults() *TableflowV1CatalogIntegrationSpec`

NewTableflowV1CatalogIntegrationSpecWithDefaults instantiates a new TableflowV1CatalogIntegrationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *TableflowV1CatalogIntegrationSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TableflowV1CatalogIntegrationSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TableflowV1CatalogIntegrationSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TableflowV1CatalogIntegrationSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSuspended

`func (o *TableflowV1CatalogIntegrationSpec) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *TableflowV1CatalogIntegrationSpec) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *TableflowV1CatalogIntegrationSpec) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *TableflowV1CatalogIntegrationSpec) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetConfig

`func (o *TableflowV1CatalogIntegrationSpec) GetConfig() TableflowV1CatalogIntegrationSpecConfigOneOf`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *TableflowV1CatalogIntegrationSpec) GetConfigOk() (*TableflowV1CatalogIntegrationSpecConfigOneOf, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *TableflowV1CatalogIntegrationSpec) SetConfig(v TableflowV1CatalogIntegrationSpecConfigOneOf)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *TableflowV1CatalogIntegrationSpec) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnvironment

`func (o *TableflowV1CatalogIntegrationSpec) GetEnvironment() GlobalObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *TableflowV1CatalogIntegrationSpec) GetEnvironmentOk() (*GlobalObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *TableflowV1CatalogIntegrationSpec) SetEnvironment(v GlobalObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *TableflowV1CatalogIntegrationSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetKafkaCluster

`func (o *TableflowV1CatalogIntegrationSpec) GetKafkaCluster() EnvScopedObjectReference`

GetKafkaCluster returns the KafkaCluster field if non-nil, zero value otherwise.

### GetKafkaClusterOk

`func (o *TableflowV1CatalogIntegrationSpec) GetKafkaClusterOk() (*EnvScopedObjectReference, bool)`

GetKafkaClusterOk returns a tuple with the KafkaCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaCluster

`func (o *TableflowV1CatalogIntegrationSpec) SetKafkaCluster(v EnvScopedObjectReference)`

SetKafkaCluster sets KafkaCluster field to given value.

### HasKafkaCluster

`func (o *TableflowV1CatalogIntegrationSpec) HasKafkaCluster() bool`

HasKafkaCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


