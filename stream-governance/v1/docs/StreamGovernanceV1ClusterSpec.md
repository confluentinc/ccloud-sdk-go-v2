# StreamGovernanceV1ClusterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The cluster name. | [optional] [readonly] 
**Package** | Pointer to **string** | The billing package. | [optional] 
**HttpEndpoint** | Pointer to **string** | The cluster HTTP request URL. | [optional] [readonly] 
**Cloud** | Pointer to **string** | The cloud service provider in which the cluster is running. | [optional] 
**Region** | Pointer to **string** | The cloud region where the cluster is running. | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewStreamGovernanceV1ClusterSpec

`func NewStreamGovernanceV1ClusterSpec() *StreamGovernanceV1ClusterSpec`

NewStreamGovernanceV1ClusterSpec instantiates a new StreamGovernanceV1ClusterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamGovernanceV1ClusterSpecWithDefaults

`func NewStreamGovernanceV1ClusterSpecWithDefaults() *StreamGovernanceV1ClusterSpec`

NewStreamGovernanceV1ClusterSpecWithDefaults instantiates a new StreamGovernanceV1ClusterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *StreamGovernanceV1ClusterSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *StreamGovernanceV1ClusterSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *StreamGovernanceV1ClusterSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *StreamGovernanceV1ClusterSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPackage

`func (o *StreamGovernanceV1ClusterSpec) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *StreamGovernanceV1ClusterSpec) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *StreamGovernanceV1ClusterSpec) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *StreamGovernanceV1ClusterSpec) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetHttpEndpoint

`func (o *StreamGovernanceV1ClusterSpec) GetHttpEndpoint() string`

GetHttpEndpoint returns the HttpEndpoint field if non-nil, zero value otherwise.

### GetHttpEndpointOk

`func (o *StreamGovernanceV1ClusterSpec) GetHttpEndpointOk() (*string, bool)`

GetHttpEndpointOk returns a tuple with the HttpEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpEndpoint

`func (o *StreamGovernanceV1ClusterSpec) SetHttpEndpoint(v string)`

SetHttpEndpoint sets HttpEndpoint field to given value.

### HasHttpEndpoint

`func (o *StreamGovernanceV1ClusterSpec) HasHttpEndpoint() bool`

HasHttpEndpoint returns a boolean if a field has been set.

### GetCloud

`func (o *StreamGovernanceV1ClusterSpec) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *StreamGovernanceV1ClusterSpec) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *StreamGovernanceV1ClusterSpec) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *StreamGovernanceV1ClusterSpec) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetRegion

`func (o *StreamGovernanceV1ClusterSpec) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StreamGovernanceV1ClusterSpec) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StreamGovernanceV1ClusterSpec) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *StreamGovernanceV1ClusterSpec) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetEnvironment

`func (o *StreamGovernanceV1ClusterSpec) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *StreamGovernanceV1ClusterSpec) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *StreamGovernanceV1ClusterSpec) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *StreamGovernanceV1ClusterSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


