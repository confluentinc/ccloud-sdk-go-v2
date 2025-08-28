# UsmV1ConnectCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**ConfluentPlatformConnectClusterId** | Pointer to **string** | The unique identifier of the Connect cluster within the Confluent Platform environment. | [optional] 
**KafkaClusterId** | Pointer to **string** | The unique identifier of the metadata Kafka cluster for the Connect Cluster.  | [optional] 
**Cloud** | Pointer to **string** | The cloud service provider where the metadata for the Connect Cluster should be stored. This field is optional. If provided, &#39;region&#39; must also be provided. If neither &#39;cloud&#39; nor &#39;region&#39; are provided, the cloud provider of the associated metadata Kafka cluster (identified by &#39;kafka_cluster_id&#39;) will be used as a fallback.  | [optional] 
**Region** | Pointer to **string** | The home region of the Confluent Platform Connect cluster where the metadata should be stored. This field is optional. If provided, &#39;cloud&#39; must also be provided. If neither &#39;cloud&#39; nor &#39;region&#39; are provided, the home region of the associated metadata Kafka cluster (identified by &#39;kafka_cluster_id&#39;) will be used as a fallback.  | [optional] 
**Environment** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewUsmV1ConnectCluster

`func NewUsmV1ConnectCluster() *UsmV1ConnectCluster`

NewUsmV1ConnectCluster instantiates a new UsmV1ConnectCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsmV1ConnectClusterWithDefaults

`func NewUsmV1ConnectClusterWithDefaults() *UsmV1ConnectCluster`

NewUsmV1ConnectClusterWithDefaults instantiates a new UsmV1ConnectCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *UsmV1ConnectCluster) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *UsmV1ConnectCluster) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *UsmV1ConnectCluster) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *UsmV1ConnectCluster) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *UsmV1ConnectCluster) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *UsmV1ConnectCluster) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *UsmV1ConnectCluster) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *UsmV1ConnectCluster) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *UsmV1ConnectCluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsmV1ConnectCluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsmV1ConnectCluster) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UsmV1ConnectCluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConfluentPlatformConnectClusterId

`func (o *UsmV1ConnectCluster) GetConfluentPlatformConnectClusterId() string`

GetConfluentPlatformConnectClusterId returns the ConfluentPlatformConnectClusterId field if non-nil, zero value otherwise.

### GetConfluentPlatformConnectClusterIdOk

`func (o *UsmV1ConnectCluster) GetConfluentPlatformConnectClusterIdOk() (*string, bool)`

GetConfluentPlatformConnectClusterIdOk returns a tuple with the ConfluentPlatformConnectClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfluentPlatformConnectClusterId

`func (o *UsmV1ConnectCluster) SetConfluentPlatformConnectClusterId(v string)`

SetConfluentPlatformConnectClusterId sets ConfluentPlatformConnectClusterId field to given value.

### HasConfluentPlatformConnectClusterId

`func (o *UsmV1ConnectCluster) HasConfluentPlatformConnectClusterId() bool`

HasConfluentPlatformConnectClusterId returns a boolean if a field has been set.

### GetKafkaClusterId

`func (o *UsmV1ConnectCluster) GetKafkaClusterId() string`

GetKafkaClusterId returns the KafkaClusterId field if non-nil, zero value otherwise.

### GetKafkaClusterIdOk

`func (o *UsmV1ConnectCluster) GetKafkaClusterIdOk() (*string, bool)`

GetKafkaClusterIdOk returns a tuple with the KafkaClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaClusterId

`func (o *UsmV1ConnectCluster) SetKafkaClusterId(v string)`

SetKafkaClusterId sets KafkaClusterId field to given value.

### HasKafkaClusterId

`func (o *UsmV1ConnectCluster) HasKafkaClusterId() bool`

HasKafkaClusterId returns a boolean if a field has been set.

### GetCloud

`func (o *UsmV1ConnectCluster) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *UsmV1ConnectCluster) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *UsmV1ConnectCluster) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *UsmV1ConnectCluster) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetRegion

`func (o *UsmV1ConnectCluster) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *UsmV1ConnectCluster) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *UsmV1ConnectCluster) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *UsmV1ConnectCluster) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetEnvironment

`func (o *UsmV1ConnectCluster) GetEnvironment() EnvScopedObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *UsmV1ConnectCluster) GetEnvironmentOk() (*EnvScopedObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *UsmV1ConnectCluster) SetEnvironment(v EnvScopedObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *UsmV1ConnectCluster) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


