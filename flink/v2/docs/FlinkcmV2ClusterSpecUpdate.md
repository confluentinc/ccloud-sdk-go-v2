# FlinkcmV2ClusterSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | The cloud provider that the Flink cluster is in. | [optional] 
**Region** | Pointer to **string** | The region that the Flink cluster is in. | [optional] 
**Csu** | Pointer to **int32** | The number of CSUs (Confluent Streaming Units) in a Flink cluster. | [optional] 
**NetworkType** | Pointer to **string** | The Type of network to use. | [optional] 

## Methods

### NewFlinkcmV2ClusterSpecUpdate

`func NewFlinkcmV2ClusterSpecUpdate() *FlinkcmV2ClusterSpecUpdate`

NewFlinkcmV2ClusterSpecUpdate instantiates a new FlinkcmV2ClusterSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlinkcmV2ClusterSpecUpdateWithDefaults

`func NewFlinkcmV2ClusterSpecUpdateWithDefaults() *FlinkcmV2ClusterSpecUpdate`

NewFlinkcmV2ClusterSpecUpdateWithDefaults instantiates a new FlinkcmV2ClusterSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *FlinkcmV2ClusterSpecUpdate) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *FlinkcmV2ClusterSpecUpdate) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *FlinkcmV2ClusterSpecUpdate) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *FlinkcmV2ClusterSpecUpdate) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetRegion

`func (o *FlinkcmV2ClusterSpecUpdate) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FlinkcmV2ClusterSpecUpdate) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FlinkcmV2ClusterSpecUpdate) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *FlinkcmV2ClusterSpecUpdate) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCsu

`func (o *FlinkcmV2ClusterSpecUpdate) GetCsu() int32`

GetCsu returns the Csu field if non-nil, zero value otherwise.

### GetCsuOk

`func (o *FlinkcmV2ClusterSpecUpdate) GetCsuOk() (*int32, bool)`

GetCsuOk returns a tuple with the Csu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsu

`func (o *FlinkcmV2ClusterSpecUpdate) SetCsu(v int32)`

SetCsu sets Csu field to given value.

### HasCsu

`func (o *FlinkcmV2ClusterSpecUpdate) HasCsu() bool`

HasCsu returns a boolean if a field has been set.

### GetNetworkType

`func (o *FlinkcmV2ClusterSpecUpdate) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *FlinkcmV2ClusterSpecUpdate) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *FlinkcmV2ClusterSpecUpdate) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *FlinkcmV2ClusterSpecUpdate) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


