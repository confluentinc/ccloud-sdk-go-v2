# CmkV2ClusterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the cluster. | [optional] 
**Availability** | Pointer to **string** | The availability zone configuration of the cluster Note: The availability zone can be updated from Single to Multi-Zone for Basic and Standard clusters but cannot be downgraded from Multi-Zone to Single Zone.  | [optional] 
**Cloud** | Pointer to **string** | The cloud service provider in which the cluster is running. | [optional] 
**Region** | Pointer to **string** | The cloud service provider region where the cluster is running. | [optional] 
**Config** | Pointer to [**CmkV2ClusterSpecConfigOneOf**](CmkV2ClusterSpecConfigOneOf.md) | The configuration of the Kafka cluster.  Note: Clusters can be upgraded from Basic to Standard, but cannot be downgraded from Standard to Basic.  | [optional] 
**KafkaBootstrapEndpoint** | Pointer to **string** | The bootstrap endpoint used by Kafka clients to connect to the cluster. | [optional] [readonly] 
**HttpEndpoint** | Pointer to **string** | The cluster HTTP request URL. | [optional] [readonly] 
**ApiEndpoint** | Pointer to **string** | The Kafka API cluster endpoint used by Kafka clients to connect to the cluster. | [optional] [readonly] 
**Environment** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The environment to which this belongs. | [optional] 
**Network** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The network associated with this object. | [optional] 
**Byok** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The byok associated with this object. | [optional] 

## Methods

### NewCmkV2ClusterSpec

`func NewCmkV2ClusterSpec() *CmkV2ClusterSpec`

NewCmkV2ClusterSpec instantiates a new CmkV2ClusterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmkV2ClusterSpecWithDefaults

`func NewCmkV2ClusterSpecWithDefaults() *CmkV2ClusterSpec`

NewCmkV2ClusterSpecWithDefaults instantiates a new CmkV2ClusterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *CmkV2ClusterSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CmkV2ClusterSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CmkV2ClusterSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CmkV2ClusterSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAvailability

`func (o *CmkV2ClusterSpec) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *CmkV2ClusterSpec) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *CmkV2ClusterSpec) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *CmkV2ClusterSpec) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetCloud

`func (o *CmkV2ClusterSpec) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *CmkV2ClusterSpec) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *CmkV2ClusterSpec) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *CmkV2ClusterSpec) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetRegion

`func (o *CmkV2ClusterSpec) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CmkV2ClusterSpec) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CmkV2ClusterSpec) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CmkV2ClusterSpec) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetConfig

`func (o *CmkV2ClusterSpec) GetConfig() CmkV2ClusterSpecConfigOneOf`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CmkV2ClusterSpec) GetConfigOk() (*CmkV2ClusterSpecConfigOneOf, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CmkV2ClusterSpec) SetConfig(v CmkV2ClusterSpecConfigOneOf)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CmkV2ClusterSpec) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetKafkaBootstrapEndpoint

`func (o *CmkV2ClusterSpec) GetKafkaBootstrapEndpoint() string`

GetKafkaBootstrapEndpoint returns the KafkaBootstrapEndpoint field if non-nil, zero value otherwise.

### GetKafkaBootstrapEndpointOk

`func (o *CmkV2ClusterSpec) GetKafkaBootstrapEndpointOk() (*string, bool)`

GetKafkaBootstrapEndpointOk returns a tuple with the KafkaBootstrapEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaBootstrapEndpoint

`func (o *CmkV2ClusterSpec) SetKafkaBootstrapEndpoint(v string)`

SetKafkaBootstrapEndpoint sets KafkaBootstrapEndpoint field to given value.

### HasKafkaBootstrapEndpoint

`func (o *CmkV2ClusterSpec) HasKafkaBootstrapEndpoint() bool`

HasKafkaBootstrapEndpoint returns a boolean if a field has been set.

### GetHttpEndpoint

`func (o *CmkV2ClusterSpec) GetHttpEndpoint() string`

GetHttpEndpoint returns the HttpEndpoint field if non-nil, zero value otherwise.

### GetHttpEndpointOk

`func (o *CmkV2ClusterSpec) GetHttpEndpointOk() (*string, bool)`

GetHttpEndpointOk returns a tuple with the HttpEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpEndpoint

`func (o *CmkV2ClusterSpec) SetHttpEndpoint(v string)`

SetHttpEndpoint sets HttpEndpoint field to given value.

### HasHttpEndpoint

`func (o *CmkV2ClusterSpec) HasHttpEndpoint() bool`

HasHttpEndpoint returns a boolean if a field has been set.

### GetApiEndpoint

`func (o *CmkV2ClusterSpec) GetApiEndpoint() string`

GetApiEndpoint returns the ApiEndpoint field if non-nil, zero value otherwise.

### GetApiEndpointOk

`func (o *CmkV2ClusterSpec) GetApiEndpointOk() (*string, bool)`

GetApiEndpointOk returns a tuple with the ApiEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiEndpoint

`func (o *CmkV2ClusterSpec) SetApiEndpoint(v string)`

SetApiEndpoint sets ApiEndpoint field to given value.

### HasApiEndpoint

`func (o *CmkV2ClusterSpec) HasApiEndpoint() bool`

HasApiEndpoint returns a boolean if a field has been set.

### GetEnvironment

`func (o *CmkV2ClusterSpec) GetEnvironment() EnvScopedObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CmkV2ClusterSpec) GetEnvironmentOk() (*EnvScopedObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CmkV2ClusterSpec) SetEnvironment(v EnvScopedObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CmkV2ClusterSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetNetwork

`func (o *CmkV2ClusterSpec) GetNetwork() EnvScopedObjectReference`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CmkV2ClusterSpec) GetNetworkOk() (*EnvScopedObjectReference, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CmkV2ClusterSpec) SetNetwork(v EnvScopedObjectReference)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CmkV2ClusterSpec) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetByok

`func (o *CmkV2ClusterSpec) GetByok() GlobalObjectReference`

GetByok returns the Byok field if non-nil, zero value otherwise.

### GetByokOk

`func (o *CmkV2ClusterSpec) GetByokOk() (*GlobalObjectReference, bool)`

GetByokOk returns a tuple with the Byok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByok

`func (o *CmkV2ClusterSpec) SetByok(v GlobalObjectReference)`

SetByok sets Byok field to given value.

### HasByok

`func (o *CmkV2ClusterSpec) HasByok() bool`

HasByok returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


