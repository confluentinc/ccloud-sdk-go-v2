# FcpmV2QuickstartComputePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | The environment for which quickstart a compute pool.  | 
**Region** | **string** | The region for which quickstart a compute pool.  | 
**Cloud** | **string** | The cloud for which quickstart a compute pool.  | 
**ComputePoolId** | Pointer to **string** | The id of the compute pool.  | [optional] [readonly] 
**DisplayName** | Pointer to **string** | The display_name of the compute pool.  | [optional] [readonly] 
**MaxCfu** | Pointer to **int32** | The max_cfu of the compute pool.  | [optional] [readonly] 
**Error** | Pointer to [**FcpmV2Error**](fcpm.V2.Error.md) | The error encountered while creating the compute pool.  | [optional] [readonly] 

## Methods

### NewFcpmV2QuickstartComputePool

`func NewFcpmV2QuickstartComputePool(environmentId string, region string, cloud string, ) *FcpmV2QuickstartComputePool`

NewFcpmV2QuickstartComputePool instantiates a new FcpmV2QuickstartComputePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpmV2QuickstartComputePoolWithDefaults

`func NewFcpmV2QuickstartComputePoolWithDefaults() *FcpmV2QuickstartComputePool`

NewFcpmV2QuickstartComputePoolWithDefaults instantiates a new FcpmV2QuickstartComputePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *FcpmV2QuickstartComputePool) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FcpmV2QuickstartComputePool) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FcpmV2QuickstartComputePool) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetRegion

`func (o *FcpmV2QuickstartComputePool) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FcpmV2QuickstartComputePool) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FcpmV2QuickstartComputePool) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetCloud

`func (o *FcpmV2QuickstartComputePool) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *FcpmV2QuickstartComputePool) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *FcpmV2QuickstartComputePool) SetCloud(v string)`

SetCloud sets Cloud field to given value.


### GetComputePoolId

`func (o *FcpmV2QuickstartComputePool) GetComputePoolId() string`

GetComputePoolId returns the ComputePoolId field if non-nil, zero value otherwise.

### GetComputePoolIdOk

`func (o *FcpmV2QuickstartComputePool) GetComputePoolIdOk() (*string, bool)`

GetComputePoolIdOk returns a tuple with the ComputePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePoolId

`func (o *FcpmV2QuickstartComputePool) SetComputePoolId(v string)`

SetComputePoolId sets ComputePoolId field to given value.

### HasComputePoolId

`func (o *FcpmV2QuickstartComputePool) HasComputePoolId() bool`

HasComputePoolId returns a boolean if a field has been set.

### GetDisplayName

`func (o *FcpmV2QuickstartComputePool) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FcpmV2QuickstartComputePool) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FcpmV2QuickstartComputePool) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FcpmV2QuickstartComputePool) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMaxCfu

`func (o *FcpmV2QuickstartComputePool) GetMaxCfu() int32`

GetMaxCfu returns the MaxCfu field if non-nil, zero value otherwise.

### GetMaxCfuOk

`func (o *FcpmV2QuickstartComputePool) GetMaxCfuOk() (*int32, bool)`

GetMaxCfuOk returns a tuple with the MaxCfu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCfu

`func (o *FcpmV2QuickstartComputePool) SetMaxCfu(v int32)`

SetMaxCfu sets MaxCfu field to given value.

### HasMaxCfu

`func (o *FcpmV2QuickstartComputePool) HasMaxCfu() bool`

HasMaxCfu returns a boolean if a field has been set.

### GetError

`func (o *FcpmV2QuickstartComputePool) GetError() FcpmV2Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *FcpmV2QuickstartComputePool) GetErrorOk() (*FcpmV2Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *FcpmV2QuickstartComputePool) SetError(v FcpmV2Error)`

SetError sets Error field to given value.

### HasError

`func (o *FcpmV2QuickstartComputePool) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


