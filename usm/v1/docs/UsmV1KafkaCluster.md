# UsmV1KafkaCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**DisplayName** | Pointer to **string** | A human-readable name for the Confluent Platform Kafka cluster. | [optional] 
**ConfluentPlatformKafkaClusterId** | Pointer to **string** | The unique identifier of the Kafka cluster within the Confluent Platform environment. | [optional] 
**Cloud** | Pointer to **string** | The cloud service provider where the metadata for the Kafka Cluster should be stored. | [optional] 
**Region** | Pointer to **string** | The home region of the Confluent Platform Kafka cluster where the metadata should be stored. | [optional] 
**Environment** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewUsmV1KafkaCluster

`func NewUsmV1KafkaCluster() *UsmV1KafkaCluster`

NewUsmV1KafkaCluster instantiates a new UsmV1KafkaCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsmV1KafkaClusterWithDefaults

`func NewUsmV1KafkaClusterWithDefaults() *UsmV1KafkaCluster`

NewUsmV1KafkaClusterWithDefaults instantiates a new UsmV1KafkaCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *UsmV1KafkaCluster) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *UsmV1KafkaCluster) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *UsmV1KafkaCluster) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *UsmV1KafkaCluster) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *UsmV1KafkaCluster) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *UsmV1KafkaCluster) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *UsmV1KafkaCluster) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *UsmV1KafkaCluster) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *UsmV1KafkaCluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsmV1KafkaCluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsmV1KafkaCluster) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UsmV1KafkaCluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *UsmV1KafkaCluster) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UsmV1KafkaCluster) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UsmV1KafkaCluster) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UsmV1KafkaCluster) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetConfluentPlatformKafkaClusterId

`func (o *UsmV1KafkaCluster) GetConfluentPlatformKafkaClusterId() string`

GetConfluentPlatformKafkaClusterId returns the ConfluentPlatformKafkaClusterId field if non-nil, zero value otherwise.

### GetConfluentPlatformKafkaClusterIdOk

`func (o *UsmV1KafkaCluster) GetConfluentPlatformKafkaClusterIdOk() (*string, bool)`

GetConfluentPlatformKafkaClusterIdOk returns a tuple with the ConfluentPlatformKafkaClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfluentPlatformKafkaClusterId

`func (o *UsmV1KafkaCluster) SetConfluentPlatformKafkaClusterId(v string)`

SetConfluentPlatformKafkaClusterId sets ConfluentPlatformKafkaClusterId field to given value.

### HasConfluentPlatformKafkaClusterId

`func (o *UsmV1KafkaCluster) HasConfluentPlatformKafkaClusterId() bool`

HasConfluentPlatformKafkaClusterId returns a boolean if a field has been set.

### GetCloud

`func (o *UsmV1KafkaCluster) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *UsmV1KafkaCluster) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *UsmV1KafkaCluster) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *UsmV1KafkaCluster) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetRegion

`func (o *UsmV1KafkaCluster) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *UsmV1KafkaCluster) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *UsmV1KafkaCluster) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *UsmV1KafkaCluster) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetEnvironment

`func (o *UsmV1KafkaCluster) GetEnvironment() EnvScopedObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *UsmV1KafkaCluster) GetEnvironmentOk() (*EnvScopedObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *UsmV1KafkaCluster) SetEnvironment(v EnvScopedObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *UsmV1KafkaCluster) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


