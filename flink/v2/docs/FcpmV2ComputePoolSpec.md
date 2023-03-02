# FcpmV2ComputePoolSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the Flink compute pool. | [optional] 
**Cloud** | Pointer to **string** | The cloud service provider that runs the compute pool. | [optional] 
**HttpEndpoint** | Pointer to **string** | The API endpoint of the Flink compute pool. | [optional] [readonly] 
**Region** | Pointer to **string** | The cloud service provider region that hosts the Flink compute pool. | [optional] 
**CurrentCsu** | Pointer to **int32** | The number of CSUs (Confluent Streaming Units) currently allocated to this Flink compute pool. | [optional] [readonly] 
**MaxCsu** | Pointer to **int32** | Maximum number of Confluent Streaming Units (CSUs) that the Flink compute pool should auto-scale to. If not specified, the value defaults to 8.  | [optional] 
**Config** | Pointer to [**FcpmV2ComputePoolSpecConfigOneOf**](FcpmV2ComputePoolSpecConfigOneOf.md) | The type of of the Flink compute pool.  | [optional] 
**Environment** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The environment to which this belongs. | [optional] 
**Network** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The network to which this belongs. | [optional] 

## Methods

### NewFcpmV2ComputePoolSpec

`func NewFcpmV2ComputePoolSpec() *FcpmV2ComputePoolSpec`

NewFcpmV2ComputePoolSpec instantiates a new FcpmV2ComputePoolSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpmV2ComputePoolSpecWithDefaults

`func NewFcpmV2ComputePoolSpecWithDefaults() *FcpmV2ComputePoolSpec`

NewFcpmV2ComputePoolSpecWithDefaults instantiates a new FcpmV2ComputePoolSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *FcpmV2ComputePoolSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FcpmV2ComputePoolSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FcpmV2ComputePoolSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FcpmV2ComputePoolSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCloud

`func (o *FcpmV2ComputePoolSpec) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *FcpmV2ComputePoolSpec) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *FcpmV2ComputePoolSpec) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *FcpmV2ComputePoolSpec) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetHttpEndpoint

`func (o *FcpmV2ComputePoolSpec) GetHttpEndpoint() string`

GetHttpEndpoint returns the HttpEndpoint field if non-nil, zero value otherwise.

### GetHttpEndpointOk

`func (o *FcpmV2ComputePoolSpec) GetHttpEndpointOk() (*string, bool)`

GetHttpEndpointOk returns a tuple with the HttpEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpEndpoint

`func (o *FcpmV2ComputePoolSpec) SetHttpEndpoint(v string)`

SetHttpEndpoint sets HttpEndpoint field to given value.

### HasHttpEndpoint

`func (o *FcpmV2ComputePoolSpec) HasHttpEndpoint() bool`

HasHttpEndpoint returns a boolean if a field has been set.

### GetRegion

`func (o *FcpmV2ComputePoolSpec) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FcpmV2ComputePoolSpec) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FcpmV2ComputePoolSpec) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *FcpmV2ComputePoolSpec) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCurrentCsu

`func (o *FcpmV2ComputePoolSpec) GetCurrentCsu() int32`

GetCurrentCsu returns the CurrentCsu field if non-nil, zero value otherwise.

### GetCurrentCsuOk

`func (o *FcpmV2ComputePoolSpec) GetCurrentCsuOk() (*int32, bool)`

GetCurrentCsuOk returns a tuple with the CurrentCsu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCsu

`func (o *FcpmV2ComputePoolSpec) SetCurrentCsu(v int32)`

SetCurrentCsu sets CurrentCsu field to given value.

### HasCurrentCsu

`func (o *FcpmV2ComputePoolSpec) HasCurrentCsu() bool`

HasCurrentCsu returns a boolean if a field has been set.

### GetMaxCsu

`func (o *FcpmV2ComputePoolSpec) GetMaxCsu() int32`

GetMaxCsu returns the MaxCsu field if non-nil, zero value otherwise.

### GetMaxCsuOk

`func (o *FcpmV2ComputePoolSpec) GetMaxCsuOk() (*int32, bool)`

GetMaxCsuOk returns a tuple with the MaxCsu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCsu

`func (o *FcpmV2ComputePoolSpec) SetMaxCsu(v int32)`

SetMaxCsu sets MaxCsu field to given value.

### HasMaxCsu

`func (o *FcpmV2ComputePoolSpec) HasMaxCsu() bool`

HasMaxCsu returns a boolean if a field has been set.

### GetConfig

`func (o *FcpmV2ComputePoolSpec) GetConfig() FcpmV2ComputePoolSpecConfigOneOf`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *FcpmV2ComputePoolSpec) GetConfigOk() (*FcpmV2ComputePoolSpecConfigOneOf, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *FcpmV2ComputePoolSpec) SetConfig(v FcpmV2ComputePoolSpecConfigOneOf)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *FcpmV2ComputePoolSpec) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnvironment

`func (o *FcpmV2ComputePoolSpec) GetEnvironment() GlobalObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *FcpmV2ComputePoolSpec) GetEnvironmentOk() (*GlobalObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *FcpmV2ComputePoolSpec) SetEnvironment(v GlobalObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *FcpmV2ComputePoolSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetNetwork

`func (o *FcpmV2ComputePoolSpec) GetNetwork() EnvScopedObjectReference`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *FcpmV2ComputePoolSpec) GetNetworkOk() (*EnvScopedObjectReference, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *FcpmV2ComputePoolSpec) SetNetwork(v EnvScopedObjectReference)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *FcpmV2ComputePoolSpec) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


