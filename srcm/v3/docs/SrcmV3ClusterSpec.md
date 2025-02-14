# SrcmV3ClusterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The cluster name. | [optional] [readonly] 
**Package** | Pointer to **string** | The billing package.  Note: Clusters can be upgraded from ESSENTIALS to ADVANCED, but cannot be downgraded from ADVANCED to ESSENTIALS.  | [optional] 
**HttpEndpoint** | Pointer to **string** | The cluster HTTP request URL. | [optional] [readonly] 
**CatalogHttpEndpoint** | Pointer to **string** | The cluster&#39;s catalog HTTP request URL. | [optional] [readonly] 
**PrivateHttpEndpoint** | Pointer to **string** | The cluster&#39;s private HTTP request URL. | [optional] [readonly] 
**PrivateNetworkingConfig** | Pointer to [**SrcmV3ClusterSpecPrivateNetworkingConfig**](SrcmV3ClusterSpecPrivateNetworkingConfig.md) |  | [optional] 
**Cloud** | Pointer to **string** | The cloud service provider in which the cluster is running. | [optional] 
**Region** | Pointer to **string** | The cloud service provider region where the cluster is running. | [optional] 
**Environment** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewSrcmV3ClusterSpec

`func NewSrcmV3ClusterSpec() *SrcmV3ClusterSpec`

NewSrcmV3ClusterSpec instantiates a new SrcmV3ClusterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSrcmV3ClusterSpecWithDefaults

`func NewSrcmV3ClusterSpecWithDefaults() *SrcmV3ClusterSpec`

NewSrcmV3ClusterSpecWithDefaults instantiates a new SrcmV3ClusterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *SrcmV3ClusterSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SrcmV3ClusterSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SrcmV3ClusterSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SrcmV3ClusterSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPackage

`func (o *SrcmV3ClusterSpec) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *SrcmV3ClusterSpec) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *SrcmV3ClusterSpec) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *SrcmV3ClusterSpec) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetHttpEndpoint

`func (o *SrcmV3ClusterSpec) GetHttpEndpoint() string`

GetHttpEndpoint returns the HttpEndpoint field if non-nil, zero value otherwise.

### GetHttpEndpointOk

`func (o *SrcmV3ClusterSpec) GetHttpEndpointOk() (*string, bool)`

GetHttpEndpointOk returns a tuple with the HttpEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpEndpoint

`func (o *SrcmV3ClusterSpec) SetHttpEndpoint(v string)`

SetHttpEndpoint sets HttpEndpoint field to given value.

### HasHttpEndpoint

`func (o *SrcmV3ClusterSpec) HasHttpEndpoint() bool`

HasHttpEndpoint returns a boolean if a field has been set.

### GetCatalogHttpEndpoint

`func (o *SrcmV3ClusterSpec) GetCatalogHttpEndpoint() string`

GetCatalogHttpEndpoint returns the CatalogHttpEndpoint field if non-nil, zero value otherwise.

### GetCatalogHttpEndpointOk

`func (o *SrcmV3ClusterSpec) GetCatalogHttpEndpointOk() (*string, bool)`

GetCatalogHttpEndpointOk returns a tuple with the CatalogHttpEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogHttpEndpoint

`func (o *SrcmV3ClusterSpec) SetCatalogHttpEndpoint(v string)`

SetCatalogHttpEndpoint sets CatalogHttpEndpoint field to given value.

### HasCatalogHttpEndpoint

`func (o *SrcmV3ClusterSpec) HasCatalogHttpEndpoint() bool`

HasCatalogHttpEndpoint returns a boolean if a field has been set.

### GetPrivateHttpEndpoint

`func (o *SrcmV3ClusterSpec) GetPrivateHttpEndpoint() string`

GetPrivateHttpEndpoint returns the PrivateHttpEndpoint field if non-nil, zero value otherwise.

### GetPrivateHttpEndpointOk

`func (o *SrcmV3ClusterSpec) GetPrivateHttpEndpointOk() (*string, bool)`

GetPrivateHttpEndpointOk returns a tuple with the PrivateHttpEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateHttpEndpoint

`func (o *SrcmV3ClusterSpec) SetPrivateHttpEndpoint(v string)`

SetPrivateHttpEndpoint sets PrivateHttpEndpoint field to given value.

### HasPrivateHttpEndpoint

`func (o *SrcmV3ClusterSpec) HasPrivateHttpEndpoint() bool`

HasPrivateHttpEndpoint returns a boolean if a field has been set.

### GetPrivateNetworkingConfig

`func (o *SrcmV3ClusterSpec) GetPrivateNetworkingConfig() SrcmV3ClusterSpecPrivateNetworkingConfig`

GetPrivateNetworkingConfig returns the PrivateNetworkingConfig field if non-nil, zero value otherwise.

### GetPrivateNetworkingConfigOk

`func (o *SrcmV3ClusterSpec) GetPrivateNetworkingConfigOk() (*SrcmV3ClusterSpecPrivateNetworkingConfig, bool)`

GetPrivateNetworkingConfigOk returns a tuple with the PrivateNetworkingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworkingConfig

`func (o *SrcmV3ClusterSpec) SetPrivateNetworkingConfig(v SrcmV3ClusterSpecPrivateNetworkingConfig)`

SetPrivateNetworkingConfig sets PrivateNetworkingConfig field to given value.

### HasPrivateNetworkingConfig

`func (o *SrcmV3ClusterSpec) HasPrivateNetworkingConfig() bool`

HasPrivateNetworkingConfig returns a boolean if a field has been set.

### GetCloud

`func (o *SrcmV3ClusterSpec) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *SrcmV3ClusterSpec) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *SrcmV3ClusterSpec) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *SrcmV3ClusterSpec) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetRegion

`func (o *SrcmV3ClusterSpec) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *SrcmV3ClusterSpec) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *SrcmV3ClusterSpec) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *SrcmV3ClusterSpec) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetEnvironment

`func (o *SrcmV3ClusterSpec) GetEnvironment() GlobalObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SrcmV3ClusterSpec) GetEnvironmentOk() (*GlobalObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SrcmV3ClusterSpec) SetEnvironment(v GlobalObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SrcmV3ClusterSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


