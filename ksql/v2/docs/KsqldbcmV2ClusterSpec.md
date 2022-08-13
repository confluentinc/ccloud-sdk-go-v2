# KsqldbcmV2ClusterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the ksqlDB cluster. | [optional] 
**UseDetailedProcessingLog** | Pointer to **bool** | This flag controls whether you want to include the row data in the processing log topic. Turn it off if you don&#39;t want to emit sensitive information to the processing log  | [optional] [default to true]
**Csu** | Pointer to **int32** | The number of CSUs (Confluent Streaming Units) in a ksqlDB cluster. | [optional] 
**KafkaCluster** | Pointer to [**ObjectReference**](ObjectReference.md) | The kafka_cluster to which this belongs. | [optional] 
**CredentialIdentity** | Pointer to [**ObjectReference**](ObjectReference.md) | The credential_identity to which this belongs. The credential_identity can be one of iam.v2.User, iam.v2.ServiceAccount. | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewKsqldbcmV2ClusterSpec

`func NewKsqldbcmV2ClusterSpec() *KsqldbcmV2ClusterSpec`

NewKsqldbcmV2ClusterSpec instantiates a new KsqldbcmV2ClusterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKsqldbcmV2ClusterSpecWithDefaults

`func NewKsqldbcmV2ClusterSpecWithDefaults() *KsqldbcmV2ClusterSpec`

NewKsqldbcmV2ClusterSpecWithDefaults instantiates a new KsqldbcmV2ClusterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *KsqldbcmV2ClusterSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KsqldbcmV2ClusterSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KsqldbcmV2ClusterSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KsqldbcmV2ClusterSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetUseDetailedProcessingLog

`func (o *KsqldbcmV2ClusterSpec) GetUseDetailedProcessingLog() bool`

GetUseDetailedProcessingLog returns the UseDetailedProcessingLog field if non-nil, zero value otherwise.

### GetUseDetailedProcessingLogOk

`func (o *KsqldbcmV2ClusterSpec) GetUseDetailedProcessingLogOk() (*bool, bool)`

GetUseDetailedProcessingLogOk returns a tuple with the UseDetailedProcessingLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDetailedProcessingLog

`func (o *KsqldbcmV2ClusterSpec) SetUseDetailedProcessingLog(v bool)`

SetUseDetailedProcessingLog sets UseDetailedProcessingLog field to given value.

### HasUseDetailedProcessingLog

`func (o *KsqldbcmV2ClusterSpec) HasUseDetailedProcessingLog() bool`

HasUseDetailedProcessingLog returns a boolean if a field has been set.

### GetCsu

`func (o *KsqldbcmV2ClusterSpec) GetCsu() int32`

GetCsu returns the Csu field if non-nil, zero value otherwise.

### GetCsuOk

`func (o *KsqldbcmV2ClusterSpec) GetCsuOk() (*int32, bool)`

GetCsuOk returns a tuple with the Csu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsu

`func (o *KsqldbcmV2ClusterSpec) SetCsu(v int32)`

SetCsu sets Csu field to given value.

### HasCsu

`func (o *KsqldbcmV2ClusterSpec) HasCsu() bool`

HasCsu returns a boolean if a field has been set.

### GetKafkaCluster

`func (o *KsqldbcmV2ClusterSpec) GetKafkaCluster() ObjectReference`

GetKafkaCluster returns the KafkaCluster field if non-nil, zero value otherwise.

### GetKafkaClusterOk

`func (o *KsqldbcmV2ClusterSpec) GetKafkaClusterOk() (*ObjectReference, bool)`

GetKafkaClusterOk returns a tuple with the KafkaCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaCluster

`func (o *KsqldbcmV2ClusterSpec) SetKafkaCluster(v ObjectReference)`

SetKafkaCluster sets KafkaCluster field to given value.

### HasKafkaCluster

`func (o *KsqldbcmV2ClusterSpec) HasKafkaCluster() bool`

HasKafkaCluster returns a boolean if a field has been set.

### GetCredentialIdentity

`func (o *KsqldbcmV2ClusterSpec) GetCredentialIdentity() ObjectReference`

GetCredentialIdentity returns the CredentialIdentity field if non-nil, zero value otherwise.

### GetCredentialIdentityOk

`func (o *KsqldbcmV2ClusterSpec) GetCredentialIdentityOk() (*ObjectReference, bool)`

GetCredentialIdentityOk returns a tuple with the CredentialIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialIdentity

`func (o *KsqldbcmV2ClusterSpec) SetCredentialIdentity(v ObjectReference)`

SetCredentialIdentity sets CredentialIdentity field to given value.

### HasCredentialIdentity

`func (o *KsqldbcmV2ClusterSpec) HasCredentialIdentity() bool`

HasCredentialIdentity returns a boolean if a field has been set.

### GetEnvironment

`func (o *KsqldbcmV2ClusterSpec) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *KsqldbcmV2ClusterSpec) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *KsqldbcmV2ClusterSpec) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *KsqldbcmV2ClusterSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


