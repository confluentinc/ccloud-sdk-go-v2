# KafkaQuotasV1ClientQuotaSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the client quota. | [optional] 
**Description** | Pointer to **string** | A human readable description for the client quota. | [optional] 
**Throughput** | Pointer to [**KafkaQuotasV1Throughput**](kafka-quotas.v1.Throughput.md) | Throughput for the client quota. | [optional] 
**Cluster** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The ID of the Dedicated Kafka cluster where the client quota is applied.  | [optional] 
**Principals** | Pointer to [**[]GlobalObjectReference**](GlobalObjectReference.md) | A list of principals to apply a client quota to. Use \&quot;&lt;default&gt;\&quot; to apply a client quota to all service accounts (see [Control application usage with Client Quotas](https://docs.confluent.io/cloud/current/clusters/client-quotas.html#control-application-usage-with-client-quotas) for more details).  | [optional] 
**Environment** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewKafkaQuotasV1ClientQuotaSpec

`func NewKafkaQuotasV1ClientQuotaSpec() *KafkaQuotasV1ClientQuotaSpec`

NewKafkaQuotasV1ClientQuotaSpec instantiates a new KafkaQuotasV1ClientQuotaSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaQuotasV1ClientQuotaSpecWithDefaults

`func NewKafkaQuotasV1ClientQuotaSpecWithDefaults() *KafkaQuotasV1ClientQuotaSpec`

NewKafkaQuotasV1ClientQuotaSpecWithDefaults instantiates a new KafkaQuotasV1ClientQuotaSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *KafkaQuotasV1ClientQuotaSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KafkaQuotasV1ClientQuotaSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KafkaQuotasV1ClientQuotaSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KafkaQuotasV1ClientQuotaSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *KafkaQuotasV1ClientQuotaSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KafkaQuotasV1ClientQuotaSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KafkaQuotasV1ClientQuotaSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KafkaQuotasV1ClientQuotaSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetThroughput

`func (o *KafkaQuotasV1ClientQuotaSpec) GetThroughput() KafkaQuotasV1Throughput`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *KafkaQuotasV1ClientQuotaSpec) GetThroughputOk() (*KafkaQuotasV1Throughput, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *KafkaQuotasV1ClientQuotaSpec) SetThroughput(v KafkaQuotasV1Throughput)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *KafkaQuotasV1ClientQuotaSpec) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.

### GetCluster

`func (o *KafkaQuotasV1ClientQuotaSpec) GetCluster() EnvScopedObjectReference`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *KafkaQuotasV1ClientQuotaSpec) GetClusterOk() (*EnvScopedObjectReference, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *KafkaQuotasV1ClientQuotaSpec) SetCluster(v EnvScopedObjectReference)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *KafkaQuotasV1ClientQuotaSpec) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetPrincipals

`func (o *KafkaQuotasV1ClientQuotaSpec) GetPrincipals() []GlobalObjectReference`

GetPrincipals returns the Principals field if non-nil, zero value otherwise.

### GetPrincipalsOk

`func (o *KafkaQuotasV1ClientQuotaSpec) GetPrincipalsOk() (*[]GlobalObjectReference, bool)`

GetPrincipalsOk returns a tuple with the Principals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipals

`func (o *KafkaQuotasV1ClientQuotaSpec) SetPrincipals(v []GlobalObjectReference)`

SetPrincipals sets Principals field to given value.

### HasPrincipals

`func (o *KafkaQuotasV1ClientQuotaSpec) HasPrincipals() bool`

HasPrincipals returns a boolean if a field has been set.

### GetEnvironment

`func (o *KafkaQuotasV1ClientQuotaSpec) GetEnvironment() GlobalObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *KafkaQuotasV1ClientQuotaSpec) GetEnvironmentOk() (*GlobalObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *KafkaQuotasV1ClientQuotaSpec) SetEnvironment(v GlobalObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *KafkaQuotasV1ClientQuotaSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


