# FcpmV2ComputePoolStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | Status of the Flink compute pool. | [readonly] 
**CurrentCfu** | Pointer to **int32** | The number of CSUs (Confluent Streaming Units) currently allocated to this Flink compute pool. | [optional] [readonly] 

## Methods

### NewFcpmV2ComputePoolStatus

`func NewFcpmV2ComputePoolStatus(phase string, ) *FcpmV2ComputePoolStatus`

NewFcpmV2ComputePoolStatus instantiates a new FcpmV2ComputePoolStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpmV2ComputePoolStatusWithDefaults

`func NewFcpmV2ComputePoolStatusWithDefaults() *FcpmV2ComputePoolStatus`

NewFcpmV2ComputePoolStatusWithDefaults instantiates a new FcpmV2ComputePoolStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *FcpmV2ComputePoolStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *FcpmV2ComputePoolStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *FcpmV2ComputePoolStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetCurrentCfu

`func (o *FcpmV2ComputePoolStatus) GetCurrentCfu() int32`

GetCurrentCfu returns the CurrentCfu field if non-nil, zero value otherwise.

### GetCurrentCfuOk

`func (o *FcpmV2ComputePoolStatus) GetCurrentCfuOk() (*int32, bool)`

GetCurrentCfuOk returns a tuple with the CurrentCfu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCfu

`func (o *FcpmV2ComputePoolStatus) SetCurrentCfu(v int32)`

SetCurrentCfu sets CurrentCfu field to given value.

### HasCurrentCfu

`func (o *FcpmV2ComputePoolStatus) HasCurrentCfu() bool`

HasCurrentCfu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


