# FrpmV2ResourcePoolSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpEndpoint** | Pointer to **string** | The API endpoint of the Flink resource pool. | [optional] 
**CurrentCsu** | Pointer to **int32** | The number of CSUs (Confluent Streaming Units) currently allocated to this Flink resource pool. | [optional] [readonly] 
**MaxCsu** | Pointer to **int32** | Maximum number of Confluent Streaming Units (CSUs) that the Flink resource pool should auto-scale to. If not specified, the value defaults to 8.  | [optional] 
**Config** | Pointer to [**FrpmV2ResourcePoolSpecUpdateConfigOneOf**](FrpmV2ResourcePoolSpecUpdateConfigOneOf.md) | The type of of the Flink resource pool. Note: Clusters can be upgraded from Basic to Standard, but cannot be downgraded from Standard to Basic.  | [optional] 

## Methods

### NewFrpmV2ResourcePoolSpecUpdate

`func NewFrpmV2ResourcePoolSpecUpdate() *FrpmV2ResourcePoolSpecUpdate`

NewFrpmV2ResourcePoolSpecUpdate instantiates a new FrpmV2ResourcePoolSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrpmV2ResourcePoolSpecUpdateWithDefaults

`func NewFrpmV2ResourcePoolSpecUpdateWithDefaults() *FrpmV2ResourcePoolSpecUpdate`

NewFrpmV2ResourcePoolSpecUpdateWithDefaults instantiates a new FrpmV2ResourcePoolSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpEndpoint

`func (o *FrpmV2ResourcePoolSpecUpdate) GetHttpEndpoint() string`

GetHttpEndpoint returns the HttpEndpoint field if non-nil, zero value otherwise.

### GetHttpEndpointOk

`func (o *FrpmV2ResourcePoolSpecUpdate) GetHttpEndpointOk() (*string, bool)`

GetHttpEndpointOk returns a tuple with the HttpEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpEndpoint

`func (o *FrpmV2ResourcePoolSpecUpdate) SetHttpEndpoint(v string)`

SetHttpEndpoint sets HttpEndpoint field to given value.

### HasHttpEndpoint

`func (o *FrpmV2ResourcePoolSpecUpdate) HasHttpEndpoint() bool`

HasHttpEndpoint returns a boolean if a field has been set.

### GetCurrentCsu

`func (o *FrpmV2ResourcePoolSpecUpdate) GetCurrentCsu() int32`

GetCurrentCsu returns the CurrentCsu field if non-nil, zero value otherwise.

### GetCurrentCsuOk

`func (o *FrpmV2ResourcePoolSpecUpdate) GetCurrentCsuOk() (*int32, bool)`

GetCurrentCsuOk returns a tuple with the CurrentCsu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCsu

`func (o *FrpmV2ResourcePoolSpecUpdate) SetCurrentCsu(v int32)`

SetCurrentCsu sets CurrentCsu field to given value.

### HasCurrentCsu

`func (o *FrpmV2ResourcePoolSpecUpdate) HasCurrentCsu() bool`

HasCurrentCsu returns a boolean if a field has been set.

### GetMaxCsu

`func (o *FrpmV2ResourcePoolSpecUpdate) GetMaxCsu() int32`

GetMaxCsu returns the MaxCsu field if non-nil, zero value otherwise.

### GetMaxCsuOk

`func (o *FrpmV2ResourcePoolSpecUpdate) GetMaxCsuOk() (*int32, bool)`

GetMaxCsuOk returns a tuple with the MaxCsu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCsu

`func (o *FrpmV2ResourcePoolSpecUpdate) SetMaxCsu(v int32)`

SetMaxCsu sets MaxCsu field to given value.

### HasMaxCsu

`func (o *FrpmV2ResourcePoolSpecUpdate) HasMaxCsu() bool`

HasMaxCsu returns a boolean if a field has been set.

### GetConfig

`func (o *FrpmV2ResourcePoolSpecUpdate) GetConfig() FrpmV2ResourcePoolSpecUpdateConfigOneOf`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *FrpmV2ResourcePoolSpecUpdate) GetConfigOk() (*FrpmV2ResourcePoolSpecUpdateConfigOneOf, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *FrpmV2ResourcePoolSpecUpdate) SetConfig(v FrpmV2ResourcePoolSpecUpdateConfigOneOf)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *FrpmV2ResourcePoolSpecUpdate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


