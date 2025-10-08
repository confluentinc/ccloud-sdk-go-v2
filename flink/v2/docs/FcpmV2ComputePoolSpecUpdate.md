# FcpmV2ComputePoolSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the Flink compute pool. | [optional] 
**MaxCfu** | Pointer to **int32** | Maximum number of Confluent Flink Units (CFUs) that the Flink compute pool should auto-scale to.  | [optional] 
**DefaultPool** | Pointer to **bool** | The flag to indicate whether the Flink compute pool is a default compute pool or not. Only one default compute pool per environment and region is allowed.  | [optional] [default to false]
**Environment** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewFcpmV2ComputePoolSpecUpdate

`func NewFcpmV2ComputePoolSpecUpdate() *FcpmV2ComputePoolSpecUpdate`

NewFcpmV2ComputePoolSpecUpdate instantiates a new FcpmV2ComputePoolSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpmV2ComputePoolSpecUpdateWithDefaults

`func NewFcpmV2ComputePoolSpecUpdateWithDefaults() *FcpmV2ComputePoolSpecUpdate`

NewFcpmV2ComputePoolSpecUpdateWithDefaults instantiates a new FcpmV2ComputePoolSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *FcpmV2ComputePoolSpecUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FcpmV2ComputePoolSpecUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FcpmV2ComputePoolSpecUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FcpmV2ComputePoolSpecUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMaxCfu

`func (o *FcpmV2ComputePoolSpecUpdate) GetMaxCfu() int32`

GetMaxCfu returns the MaxCfu field if non-nil, zero value otherwise.

### GetMaxCfuOk

`func (o *FcpmV2ComputePoolSpecUpdate) GetMaxCfuOk() (*int32, bool)`

GetMaxCfuOk returns a tuple with the MaxCfu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCfu

`func (o *FcpmV2ComputePoolSpecUpdate) SetMaxCfu(v int32)`

SetMaxCfu sets MaxCfu field to given value.

### HasMaxCfu

`func (o *FcpmV2ComputePoolSpecUpdate) HasMaxCfu() bool`

HasMaxCfu returns a boolean if a field has been set.

### GetDefaultPool

`func (o *FcpmV2ComputePoolSpecUpdate) GetDefaultPool() bool`

GetDefaultPool returns the DefaultPool field if non-nil, zero value otherwise.

### GetDefaultPoolOk

`func (o *FcpmV2ComputePoolSpecUpdate) GetDefaultPoolOk() (*bool, bool)`

GetDefaultPoolOk returns a tuple with the DefaultPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPool

`func (o *FcpmV2ComputePoolSpecUpdate) SetDefaultPool(v bool)`

SetDefaultPool sets DefaultPool field to given value.

### HasDefaultPool

`func (o *FcpmV2ComputePoolSpecUpdate) HasDefaultPool() bool`

HasDefaultPool returns a boolean if a field has been set.

### GetEnvironment

`func (o *FcpmV2ComputePoolSpecUpdate) GetEnvironment() GlobalObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *FcpmV2ComputePoolSpecUpdate) GetEnvironmentOk() (*GlobalObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *FcpmV2ComputePoolSpecUpdate) SetEnvironment(v GlobalObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *FcpmV2ComputePoolSpecUpdate) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


