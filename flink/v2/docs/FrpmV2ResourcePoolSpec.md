# FrpmV2ResourcePoolSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the Flink resource pool. | [optional] 
**Cloud** | Pointer to **string** | The cloud service provider that runs the resource pool. | [optional] 
**HttpEndpoint** | Pointer to **string** | The API endpoint of the Flink resource pool. | [optional] 
**Region** | Pointer to **string** | The cloud service provider region that hosts the Flink resource pool. | [optional] 
**CurrentCsu** | Pointer to **int32** | The number of CSUs (Confluent Streaming Units) currently allocated to this Flink resource pool. | [optional] [readonly] 
**MaxCsu** | Pointer to **int32** | Maximum number of Confluent Streaming Units (CSUs) that the Flink resource pool should auto-scale to. If not specified, the value defaults to 8.  | [optional] 
**Config** | Pointer to [**FrpmV2ResourcePoolSpecConfigOneOf**](FrpmV2ResourcePoolSpecConfigOneOf.md) | The type of of the Flink resource pool. Note: Clusters can be upgraded from Basic to Standard, but cannot be downgraded from Standard to Basic.  | [optional] 
**Environment** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The environment to which this belongs. | [optional] 
**Network** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The network to which this belongs. | [optional] 

## Methods

### NewFrpmV2ResourcePoolSpec

`func NewFrpmV2ResourcePoolSpec() *FrpmV2ResourcePoolSpec`

NewFrpmV2ResourcePoolSpec instantiates a new FrpmV2ResourcePoolSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrpmV2ResourcePoolSpecWithDefaults

`func NewFrpmV2ResourcePoolSpecWithDefaults() *FrpmV2ResourcePoolSpec`

NewFrpmV2ResourcePoolSpecWithDefaults instantiates a new FrpmV2ResourcePoolSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *FrpmV2ResourcePoolSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FrpmV2ResourcePoolSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FrpmV2ResourcePoolSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FrpmV2ResourcePoolSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCloud

`func (o *FrpmV2ResourcePoolSpec) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *FrpmV2ResourcePoolSpec) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *FrpmV2ResourcePoolSpec) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *FrpmV2ResourcePoolSpec) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetHttpEndpoint

`func (o *FrpmV2ResourcePoolSpec) GetHttpEndpoint() string`

GetHttpEndpoint returns the HttpEndpoint field if non-nil, zero value otherwise.

### GetHttpEndpointOk

`func (o *FrpmV2ResourcePoolSpec) GetHttpEndpointOk() (*string, bool)`

GetHttpEndpointOk returns a tuple with the HttpEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpEndpoint

`func (o *FrpmV2ResourcePoolSpec) SetHttpEndpoint(v string)`

SetHttpEndpoint sets HttpEndpoint field to given value.

### HasHttpEndpoint

`func (o *FrpmV2ResourcePoolSpec) HasHttpEndpoint() bool`

HasHttpEndpoint returns a boolean if a field has been set.

### GetRegion

`func (o *FrpmV2ResourcePoolSpec) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FrpmV2ResourcePoolSpec) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FrpmV2ResourcePoolSpec) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *FrpmV2ResourcePoolSpec) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCurrentCsu

`func (o *FrpmV2ResourcePoolSpec) GetCurrentCsu() int32`

GetCurrentCsu returns the CurrentCsu field if non-nil, zero value otherwise.

### GetCurrentCsuOk

`func (o *FrpmV2ResourcePoolSpec) GetCurrentCsuOk() (*int32, bool)`

GetCurrentCsuOk returns a tuple with the CurrentCsu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCsu

`func (o *FrpmV2ResourcePoolSpec) SetCurrentCsu(v int32)`

SetCurrentCsu sets CurrentCsu field to given value.

### HasCurrentCsu

`func (o *FrpmV2ResourcePoolSpec) HasCurrentCsu() bool`

HasCurrentCsu returns a boolean if a field has been set.

### GetMaxCsu

`func (o *FrpmV2ResourcePoolSpec) GetMaxCsu() int32`

GetMaxCsu returns the MaxCsu field if non-nil, zero value otherwise.

### GetMaxCsuOk

`func (o *FrpmV2ResourcePoolSpec) GetMaxCsuOk() (*int32, bool)`

GetMaxCsuOk returns a tuple with the MaxCsu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCsu

`func (o *FrpmV2ResourcePoolSpec) SetMaxCsu(v int32)`

SetMaxCsu sets MaxCsu field to given value.

### HasMaxCsu

`func (o *FrpmV2ResourcePoolSpec) HasMaxCsu() bool`

HasMaxCsu returns a boolean if a field has been set.

### GetConfig

`func (o *FrpmV2ResourcePoolSpec) GetConfig() FrpmV2ResourcePoolSpecConfigOneOf`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *FrpmV2ResourcePoolSpec) GetConfigOk() (*FrpmV2ResourcePoolSpecConfigOneOf, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *FrpmV2ResourcePoolSpec) SetConfig(v FrpmV2ResourcePoolSpecConfigOneOf)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *FrpmV2ResourcePoolSpec) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnvironment

`func (o *FrpmV2ResourcePoolSpec) GetEnvironment() GlobalObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *FrpmV2ResourcePoolSpec) GetEnvironmentOk() (*GlobalObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *FrpmV2ResourcePoolSpec) SetEnvironment(v GlobalObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *FrpmV2ResourcePoolSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetNetwork

`func (o *FrpmV2ResourcePoolSpec) GetNetwork() EnvScopedObjectReference`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *FrpmV2ResourcePoolSpec) GetNetworkOk() (*EnvScopedObjectReference, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *FrpmV2ResourcePoolSpec) SetNetwork(v EnvScopedObjectReference)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *FrpmV2ResourcePoolSpec) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


