# KsqldbcmV2Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**DisplayName** | Pointer to **string** | The name of the ksqlDB cluster. | [optional] 
**UseDetailedProcessingLog** | Pointer to **bool** | This flag controls whether you want to include the row data in the processing log topic. Turn it off if you don&#39;t want to emit sensitive information to the processing log  | [optional] [default to true]
**Csu** | Pointer to **int32** | The number of CSUs (Confluent Streaming Units) in a ksqlDB cluster. | [optional] 
**KafkaCluster** | Pointer to [**ObjectReference**](ObjectReference.md) | The kafka_cluster to which this belongs. | [optional] 
**CredentialIdentity** | Pointer to [**ObjectReference**](ObjectReference.md) | The credential_identity to which this belongs. The credential_identity can be one of iam.v2.User, iam.v2.ServiceAccount. | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 
**Status** | Pointer to [**KsqldbcmV2ClusterStatus**](KsqldbcmV2ClusterStatus.md) |  | [optional] 

## Methods

### NewKsqldbcmV2Cluster

`func NewKsqldbcmV2Cluster() *KsqldbcmV2Cluster`

NewKsqldbcmV2Cluster instantiates a new KsqldbcmV2Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKsqldbcmV2ClusterWithDefaults

`func NewKsqldbcmV2ClusterWithDefaults() *KsqldbcmV2Cluster`

NewKsqldbcmV2ClusterWithDefaults instantiates a new KsqldbcmV2Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *KsqldbcmV2Cluster) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KsqldbcmV2Cluster) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KsqldbcmV2Cluster) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *KsqldbcmV2Cluster) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *KsqldbcmV2Cluster) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KsqldbcmV2Cluster) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KsqldbcmV2Cluster) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KsqldbcmV2Cluster) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *KsqldbcmV2Cluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KsqldbcmV2Cluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KsqldbcmV2Cluster) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KsqldbcmV2Cluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *KsqldbcmV2Cluster) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KsqldbcmV2Cluster) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KsqldbcmV2Cluster) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KsqldbcmV2Cluster) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *KsqldbcmV2Cluster) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KsqldbcmV2Cluster) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KsqldbcmV2Cluster) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KsqldbcmV2Cluster) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetUseDetailedProcessingLog

`func (o *KsqldbcmV2Cluster) GetUseDetailedProcessingLog() bool`

GetUseDetailedProcessingLog returns the UseDetailedProcessingLog field if non-nil, zero value otherwise.

### GetUseDetailedProcessingLogOk

`func (o *KsqldbcmV2Cluster) GetUseDetailedProcessingLogOk() (*bool, bool)`

GetUseDetailedProcessingLogOk returns a tuple with the UseDetailedProcessingLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDetailedProcessingLog

`func (o *KsqldbcmV2Cluster) SetUseDetailedProcessingLog(v bool)`

SetUseDetailedProcessingLog sets UseDetailedProcessingLog field to given value.

### HasUseDetailedProcessingLog

`func (o *KsqldbcmV2Cluster) HasUseDetailedProcessingLog() bool`

HasUseDetailedProcessingLog returns a boolean if a field has been set.

### GetCsu

`func (o *KsqldbcmV2Cluster) GetCsu() int32`

GetCsu returns the Csu field if non-nil, zero value otherwise.

### GetCsuOk

`func (o *KsqldbcmV2Cluster) GetCsuOk() (*int32, bool)`

GetCsuOk returns a tuple with the Csu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsu

`func (o *KsqldbcmV2Cluster) SetCsu(v int32)`

SetCsu sets Csu field to given value.

### HasCsu

`func (o *KsqldbcmV2Cluster) HasCsu() bool`

HasCsu returns a boolean if a field has been set.

### GetKafkaCluster

`func (o *KsqldbcmV2Cluster) GetKafkaCluster() ObjectReference`

GetKafkaCluster returns the KafkaCluster field if non-nil, zero value otherwise.

### GetKafkaClusterOk

`func (o *KsqldbcmV2Cluster) GetKafkaClusterOk() (*ObjectReference, bool)`

GetKafkaClusterOk returns a tuple with the KafkaCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaCluster

`func (o *KsqldbcmV2Cluster) SetKafkaCluster(v ObjectReference)`

SetKafkaCluster sets KafkaCluster field to given value.

### HasKafkaCluster

`func (o *KsqldbcmV2Cluster) HasKafkaCluster() bool`

HasKafkaCluster returns a boolean if a field has been set.

### GetCredentialIdentity

`func (o *KsqldbcmV2Cluster) GetCredentialIdentity() ObjectReference`

GetCredentialIdentity returns the CredentialIdentity field if non-nil, zero value otherwise.

### GetCredentialIdentityOk

`func (o *KsqldbcmV2Cluster) GetCredentialIdentityOk() (*ObjectReference, bool)`

GetCredentialIdentityOk returns a tuple with the CredentialIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialIdentity

`func (o *KsqldbcmV2Cluster) SetCredentialIdentity(v ObjectReference)`

SetCredentialIdentity sets CredentialIdentity field to given value.

### HasCredentialIdentity

`func (o *KsqldbcmV2Cluster) HasCredentialIdentity() bool`

HasCredentialIdentity returns a boolean if a field has been set.

### GetEnvironment

`func (o *KsqldbcmV2Cluster) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *KsqldbcmV2Cluster) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *KsqldbcmV2Cluster) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *KsqldbcmV2Cluster) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetStatus

`func (o *KsqldbcmV2Cluster) GetStatus() KsqldbcmV2ClusterStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KsqldbcmV2Cluster) GetStatusOk() (*KsqldbcmV2ClusterStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KsqldbcmV2Cluster) SetStatus(v KsqldbcmV2ClusterStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KsqldbcmV2Cluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


