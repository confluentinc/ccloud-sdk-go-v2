# StreamGovernanceV2RegionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The display name. | [optional] [readonly] 
**Cloud** | Pointer to **string** | The cloud service provider that hosts the region. | [optional] [readonly] 
**RegionName** | Pointer to **string** | The region name. | [optional] [readonly] 
**Packages** | Pointer to **[]string** | List of Stream Governance packages allowing placement in this region. | [optional] [readonly] 

## Methods

### NewStreamGovernanceV2RegionSpec

`func NewStreamGovernanceV2RegionSpec() *StreamGovernanceV2RegionSpec`

NewStreamGovernanceV2RegionSpec instantiates a new StreamGovernanceV2RegionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamGovernanceV2RegionSpecWithDefaults

`func NewStreamGovernanceV2RegionSpecWithDefaults() *StreamGovernanceV2RegionSpec`

NewStreamGovernanceV2RegionSpecWithDefaults instantiates a new StreamGovernanceV2RegionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *StreamGovernanceV2RegionSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *StreamGovernanceV2RegionSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *StreamGovernanceV2RegionSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *StreamGovernanceV2RegionSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCloud

`func (o *StreamGovernanceV2RegionSpec) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *StreamGovernanceV2RegionSpec) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *StreamGovernanceV2RegionSpec) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *StreamGovernanceV2RegionSpec) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetRegionName

`func (o *StreamGovernanceV2RegionSpec) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *StreamGovernanceV2RegionSpec) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *StreamGovernanceV2RegionSpec) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *StreamGovernanceV2RegionSpec) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetPackages

`func (o *StreamGovernanceV2RegionSpec) GetPackages() []string`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *StreamGovernanceV2RegionSpec) GetPackagesOk() (*[]string, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *StreamGovernanceV2RegionSpec) SetPackages(v []string)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *StreamGovernanceV2RegionSpec) HasPackages() bool`

HasPackages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


