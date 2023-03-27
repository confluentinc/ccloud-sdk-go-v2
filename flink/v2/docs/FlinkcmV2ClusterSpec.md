# FlinkcmV2ClusterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the Flink cluster. | [optional] 
**CloudProvider** | Pointer to **string** | The cloud provider that the Flink cluster is in. | [optional] 
**Region** | Pointer to **string** | The region that the Flink cluster is in. | [optional] 
**Csu** | Pointer to **int32** | The number of CSUs (Confluent Streaming Units) in a Flink cluster. | [optional] 
**NetworkType** | Pointer to **string** | The Type of network to use. | [optional] 
**Principal** | Pointer to [**TypedGlobalObjectReference**](TypedGlobalObjectReference.md) | The principal to which this belongs. The principal can be one of iam.v2.User, iam.v2.ServiceAccount. | [optional] 
**Environment** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The environment to which this belongs. | [optional] 
**Network** | Pointer to [**ObjectReference**](ObjectReference.md) | The network to which this belongs. | [optional] 

## Methods

### NewFlinkcmV2ClusterSpec

`func NewFlinkcmV2ClusterSpec() *FlinkcmV2ClusterSpec`

NewFlinkcmV2ClusterSpec instantiates a new FlinkcmV2ClusterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlinkcmV2ClusterSpecWithDefaults

`func NewFlinkcmV2ClusterSpecWithDefaults() *FlinkcmV2ClusterSpec`

NewFlinkcmV2ClusterSpecWithDefaults instantiates a new FlinkcmV2ClusterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *FlinkcmV2ClusterSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FlinkcmV2ClusterSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FlinkcmV2ClusterSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FlinkcmV2ClusterSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCloudProvider

`func (o *FlinkcmV2ClusterSpec) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *FlinkcmV2ClusterSpec) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *FlinkcmV2ClusterSpec) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *FlinkcmV2ClusterSpec) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetRegion

`func (o *FlinkcmV2ClusterSpec) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FlinkcmV2ClusterSpec) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FlinkcmV2ClusterSpec) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *FlinkcmV2ClusterSpec) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCsu

`func (o *FlinkcmV2ClusterSpec) GetCsu() int32`

GetCsu returns the Csu field if non-nil, zero value otherwise.

### GetCsuOk

`func (o *FlinkcmV2ClusterSpec) GetCsuOk() (*int32, bool)`

GetCsuOk returns a tuple with the Csu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsu

`func (o *FlinkcmV2ClusterSpec) SetCsu(v int32)`

SetCsu sets Csu field to given value.

### HasCsu

`func (o *FlinkcmV2ClusterSpec) HasCsu() bool`

HasCsu returns a boolean if a field has been set.

### GetNetworkType

`func (o *FlinkcmV2ClusterSpec) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *FlinkcmV2ClusterSpec) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *FlinkcmV2ClusterSpec) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *FlinkcmV2ClusterSpec) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetPrincipal

`func (o *FlinkcmV2ClusterSpec) GetPrincipal() TypedGlobalObjectReference`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *FlinkcmV2ClusterSpec) GetPrincipalOk() (*TypedGlobalObjectReference, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *FlinkcmV2ClusterSpec) SetPrincipal(v TypedGlobalObjectReference)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *FlinkcmV2ClusterSpec) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetEnvironment

`func (o *FlinkcmV2ClusterSpec) GetEnvironment() GlobalObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *FlinkcmV2ClusterSpec) GetEnvironmentOk() (*GlobalObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *FlinkcmV2ClusterSpec) SetEnvironment(v GlobalObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *FlinkcmV2ClusterSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetNetwork

`func (o *FlinkcmV2ClusterSpec) GetNetwork() ObjectReference`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *FlinkcmV2ClusterSpec) GetNetworkOk() (*ObjectReference, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *FlinkcmV2ClusterSpec) SetNetwork(v ObjectReference)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *FlinkcmV2ClusterSpec) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


