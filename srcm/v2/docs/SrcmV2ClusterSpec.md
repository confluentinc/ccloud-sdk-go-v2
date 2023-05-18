# SrcmV2ClusterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The cluster name. | [optional] [readonly] 
**Package** | Pointer to **string** | The billing package.  Note: Clusters can be upgraded from ESSENTIALS to ADVANCED, but cannot be downgraded from ADVANCED to ESSENTIALS.  | [optional] 
**NetworkType** | Pointer to **string** | The network access type for the cluster.  | [optional] [default to "PUBLIC"]
**HttpEndpoint** | Pointer to **string** | The cluster HTTP request URL. | [optional] [readonly] 
**Environment** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The environment to which this belongs. | [optional] 
**Region** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The region to which this belongs. | [optional] 

## Methods

### NewSrcmV2ClusterSpec

`func NewSrcmV2ClusterSpec() *SrcmV2ClusterSpec`

NewSrcmV2ClusterSpec instantiates a new SrcmV2ClusterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSrcmV2ClusterSpecWithDefaults

`func NewSrcmV2ClusterSpecWithDefaults() *SrcmV2ClusterSpec`

NewSrcmV2ClusterSpecWithDefaults instantiates a new SrcmV2ClusterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *SrcmV2ClusterSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SrcmV2ClusterSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SrcmV2ClusterSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SrcmV2ClusterSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPackage

`func (o *SrcmV2ClusterSpec) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *SrcmV2ClusterSpec) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *SrcmV2ClusterSpec) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *SrcmV2ClusterSpec) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetNetworkType

`func (o *SrcmV2ClusterSpec) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *SrcmV2ClusterSpec) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *SrcmV2ClusterSpec) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *SrcmV2ClusterSpec) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetHttpEndpoint

`func (o *SrcmV2ClusterSpec) GetHttpEndpoint() string`

GetHttpEndpoint returns the HttpEndpoint field if non-nil, zero value otherwise.

### GetHttpEndpointOk

`func (o *SrcmV2ClusterSpec) GetHttpEndpointOk() (*string, bool)`

GetHttpEndpointOk returns a tuple with the HttpEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpEndpoint

`func (o *SrcmV2ClusterSpec) SetHttpEndpoint(v string)`

SetHttpEndpoint sets HttpEndpoint field to given value.

### HasHttpEndpoint

`func (o *SrcmV2ClusterSpec) HasHttpEndpoint() bool`

HasHttpEndpoint returns a boolean if a field has been set.

### GetEnvironment

`func (o *SrcmV2ClusterSpec) GetEnvironment() GlobalObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SrcmV2ClusterSpec) GetEnvironmentOk() (*GlobalObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SrcmV2ClusterSpec) SetEnvironment(v GlobalObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SrcmV2ClusterSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetRegion

`func (o *SrcmV2ClusterSpec) GetRegion() GlobalObjectReference`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *SrcmV2ClusterSpec) GetRegionOk() (*GlobalObjectReference, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *SrcmV2ClusterSpec) SetRegion(v GlobalObjectReference)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *SrcmV2ClusterSpec) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


