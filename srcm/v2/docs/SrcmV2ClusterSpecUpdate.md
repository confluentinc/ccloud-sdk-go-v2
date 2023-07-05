# SrcmV2ClusterSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Package** | Pointer to **string** | The billing package.  Note: Clusters can be upgraded from ESSENTIALS to ADVANCED, but cannot be downgraded from ADVANCED to ESSENTIALS.  | [optional] 
**NetworkType** | Pointer to **string** | The network access type for the cluster.  | [optional] [default to "PUBLIC"]
**Environment** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewSrcmV2ClusterSpecUpdate

`func NewSrcmV2ClusterSpecUpdate() *SrcmV2ClusterSpecUpdate`

NewSrcmV2ClusterSpecUpdate instantiates a new SrcmV2ClusterSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSrcmV2ClusterSpecUpdateWithDefaults

`func NewSrcmV2ClusterSpecUpdateWithDefaults() *SrcmV2ClusterSpecUpdate`

NewSrcmV2ClusterSpecUpdateWithDefaults instantiates a new SrcmV2ClusterSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackage

`func (o *SrcmV2ClusterSpecUpdate) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *SrcmV2ClusterSpecUpdate) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *SrcmV2ClusterSpecUpdate) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *SrcmV2ClusterSpecUpdate) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetNetworkType

`func (o *SrcmV2ClusterSpecUpdate) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *SrcmV2ClusterSpecUpdate) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *SrcmV2ClusterSpecUpdate) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *SrcmV2ClusterSpecUpdate) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetEnvironment

`func (o *SrcmV2ClusterSpecUpdate) GetEnvironment() GlobalObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SrcmV2ClusterSpecUpdate) GetEnvironmentOk() (*GlobalObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SrcmV2ClusterSpecUpdate) SetEnvironment(v GlobalObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SrcmV2ClusterSpecUpdate) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


