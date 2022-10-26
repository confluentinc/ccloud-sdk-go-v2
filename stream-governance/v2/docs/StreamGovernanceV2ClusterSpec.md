# StreamGovernanceV2ClusterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The cluster name. | [optional] [readonly] 
**Package** | Pointer to **string** | The billing package. | [optional] 
**HttpEndpoint** | Pointer to **string** | The cluster HTTP request URL. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 
**Region** | Pointer to [**ObjectReference**](ObjectReference.md) | The region to which this belongs. | [optional] 

## Methods

### NewStreamGovernanceV2ClusterSpec

`func NewStreamGovernanceV2ClusterSpec() *StreamGovernanceV2ClusterSpec`

NewStreamGovernanceV2ClusterSpec instantiates a new StreamGovernanceV2ClusterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamGovernanceV2ClusterSpecWithDefaults

`func NewStreamGovernanceV2ClusterSpecWithDefaults() *StreamGovernanceV2ClusterSpec`

NewStreamGovernanceV2ClusterSpecWithDefaults instantiates a new StreamGovernanceV2ClusterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *StreamGovernanceV2ClusterSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *StreamGovernanceV2ClusterSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *StreamGovernanceV2ClusterSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *StreamGovernanceV2ClusterSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPackage

`func (o *StreamGovernanceV2ClusterSpec) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *StreamGovernanceV2ClusterSpec) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *StreamGovernanceV2ClusterSpec) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *StreamGovernanceV2ClusterSpec) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetHttpEndpoint

`func (o *StreamGovernanceV2ClusterSpec) GetHttpEndpoint() string`

GetHttpEndpoint returns the HttpEndpoint field if non-nil, zero value otherwise.

### GetHttpEndpointOk

`func (o *StreamGovernanceV2ClusterSpec) GetHttpEndpointOk() (*string, bool)`

GetHttpEndpointOk returns a tuple with the HttpEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpEndpoint

`func (o *StreamGovernanceV2ClusterSpec) SetHttpEndpoint(v string)`

SetHttpEndpoint sets HttpEndpoint field to given value.

### HasHttpEndpoint

`func (o *StreamGovernanceV2ClusterSpec) HasHttpEndpoint() bool`

HasHttpEndpoint returns a boolean if a field has been set.

### GetEnvironment

`func (o *StreamGovernanceV2ClusterSpec) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *StreamGovernanceV2ClusterSpec) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *StreamGovernanceV2ClusterSpec) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *StreamGovernanceV2ClusterSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetRegion

`func (o *StreamGovernanceV2ClusterSpec) GetRegion() ObjectReference`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StreamGovernanceV2ClusterSpec) GetRegionOk() (*ObjectReference, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StreamGovernanceV2ClusterSpec) SetRegion(v ObjectReference)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *StreamGovernanceV2ClusterSpec) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


