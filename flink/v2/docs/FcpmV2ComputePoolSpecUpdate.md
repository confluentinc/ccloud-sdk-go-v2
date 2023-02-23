# FcpmV2ComputePoolSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpEndpoint** | Pointer to **string** | The API endpoint of the Flink compute pool. | [optional] 
**CurrentCsu** | Pointer to **int32** | The number of CSUs (Confluent Streaming Units) currently allocated to this Flink compute pool. | [optional] [readonly] 
**MaxCsu** | Pointer to **int32** | Maximum number of Confluent Streaming Units (CSUs) that the Flink compute pool should auto-scale to. If not specified, the value defaults to 8.  | [optional] 
**Config** | Pointer to [**FcpmV2ComputePoolSpecUpdateConfigOneOf**](FcpmV2ComputePoolSpecUpdateConfigOneOf.md) | The type of of the Flink compute pool.  | [optional] 

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

### GetHttpEndpoint

`func (o *FcpmV2ComputePoolSpecUpdate) GetHttpEndpoint() string`

GetHttpEndpoint returns the HttpEndpoint field if non-nil, zero value otherwise.

### GetHttpEndpointOk

`func (o *FcpmV2ComputePoolSpecUpdate) GetHttpEndpointOk() (*string, bool)`

GetHttpEndpointOk returns a tuple with the HttpEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpEndpoint

`func (o *FcpmV2ComputePoolSpecUpdate) SetHttpEndpoint(v string)`

SetHttpEndpoint sets HttpEndpoint field to given value.

### HasHttpEndpoint

`func (o *FcpmV2ComputePoolSpecUpdate) HasHttpEndpoint() bool`

HasHttpEndpoint returns a boolean if a field has been set.

### GetCurrentCsu

`func (o *FcpmV2ComputePoolSpecUpdate) GetCurrentCsu() int32`

GetCurrentCsu returns the CurrentCsu field if non-nil, zero value otherwise.

### GetCurrentCsuOk

`func (o *FcpmV2ComputePoolSpecUpdate) GetCurrentCsuOk() (*int32, bool)`

GetCurrentCsuOk returns a tuple with the CurrentCsu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCsu

`func (o *FcpmV2ComputePoolSpecUpdate) SetCurrentCsu(v int32)`

SetCurrentCsu sets CurrentCsu field to given value.

### HasCurrentCsu

`func (o *FcpmV2ComputePoolSpecUpdate) HasCurrentCsu() bool`

HasCurrentCsu returns a boolean if a field has been set.

### GetMaxCsu

`func (o *FcpmV2ComputePoolSpecUpdate) GetMaxCsu() int32`

GetMaxCsu returns the MaxCsu field if non-nil, zero value otherwise.

### GetMaxCsuOk

`func (o *FcpmV2ComputePoolSpecUpdate) GetMaxCsuOk() (*int32, bool)`

GetMaxCsuOk returns a tuple with the MaxCsu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCsu

`func (o *FcpmV2ComputePoolSpecUpdate) SetMaxCsu(v int32)`

SetMaxCsu sets MaxCsu field to given value.

### HasMaxCsu

`func (o *FcpmV2ComputePoolSpecUpdate) HasMaxCsu() bool`

HasMaxCsu returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


