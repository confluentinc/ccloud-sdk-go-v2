# FrpmV2ResourcePoolStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | Status of the Flink resource pool. | [readonly] 

## Methods

### NewFrpmV2ResourcePoolStatus

`func NewFrpmV2ResourcePoolStatus(phase string, ) *FrpmV2ResourcePoolStatus`

NewFrpmV2ResourcePoolStatus instantiates a new FrpmV2ResourcePoolStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrpmV2ResourcePoolStatusWithDefaults

`func NewFrpmV2ResourcePoolStatusWithDefaults() *FrpmV2ResourcePoolStatus`

NewFrpmV2ResourcePoolStatusWithDefaults instantiates a new FrpmV2ResourcePoolStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *FrpmV2ResourcePoolStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *FrpmV2ResourcePoolStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *FrpmV2ResourcePoolStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


