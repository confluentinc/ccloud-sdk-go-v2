# FcpmV2ComputePoolSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxCfu** | Pointer to **int32** | Maximum number of Confluent Flink Units (CFUs) that the Flink compute pool should auto-scale to.  | [optional] 
**Config** | Pointer to [**FcpmV2ComputePoolSpecUpdateConfigOneOf**](FcpmV2ComputePoolSpecUpdateConfigOneOf.md) | The type of of the Flink compute pool.  | [optional] 
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

### GetConfig

`func (o *FcpmV2ComputePoolSpecUpdate) GetConfig() FcpmV2ComputePoolSpecUpdateConfigOneOf`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *FcpmV2ComputePoolSpecUpdate) GetConfigOk() (*FcpmV2ComputePoolSpecUpdateConfigOneOf, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *FcpmV2ComputePoolSpecUpdate) SetConfig(v FcpmV2ComputePoolSpecUpdateConfigOneOf)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *FcpmV2ComputePoolSpecUpdate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

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


