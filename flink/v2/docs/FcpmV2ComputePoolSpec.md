# FcpmV2ComputePoolSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the Flink compute pool. | [optional] 
**Cloud** | Pointer to **string** | The cloud service provider that runs the compute pool. | [optional] 
**Region** | Pointer to **string** | Flink compute pools in the region provided will be able to use this identity pool | [optional] 
**MaxCfu** | Pointer to **int32** | Maximum number of Confluent Flink Units (CFUs) that the Flink compute pool should auto-scale to.  | [optional] 
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

### GetMaxCfu

`func (o *FcpmV2ComputePoolSpec) GetMaxCfu() int32`

GetMaxCfu returns the MaxCfu field if non-nil, zero value otherwise.

### GetMaxCfuOk

`func (o *FcpmV2ComputePoolSpec) GetMaxCfuOk() (*int32, bool)`

GetMaxCfuOk returns a tuple with the MaxCfu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCfu

`func (o *FcpmV2ComputePoolSpec) SetMaxCfu(v int32)`

SetMaxCfu sets MaxCfu field to given value.

### HasMaxCfu

`func (o *FcpmV2ComputePoolSpec) HasMaxCfu() bool`

HasMaxCfu returns a boolean if a field has been set.

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


