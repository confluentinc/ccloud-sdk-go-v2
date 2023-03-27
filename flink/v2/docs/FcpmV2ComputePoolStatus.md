# FcpmV2ComputePoolStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | Status of the Flink compute pool. | [readonly] 
**CurrentCsu** | **int32** | The number of CSUs (Confluent Streaming Units) currently allocated to this Flink compute pool. | [readonly] 

## Methods

### NewFcpmV2ComputePoolStatus

`func NewFcpmV2ComputePoolStatus(phase string, currentCsu int32, ) *FcpmV2ComputePoolStatus`

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


### GetCurrentCsu

`func (o *FcpmV2ComputePoolStatus) GetCurrentCsu() int32`

GetCurrentCsu returns the CurrentCsu field if non-nil, zero value otherwise.

### GetCurrentCsuOk

`func (o *FcpmV2ComputePoolStatus) GetCurrentCsuOk() (*int32, bool)`

GetCurrentCsuOk returns a tuple with the CurrentCsu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCsu

`func (o *FcpmV2ComputePoolStatus) SetCurrentCsu(v int32)`

SetCurrentCsu sets CurrentCsu field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


