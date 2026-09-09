# SwitchoverV1SwitchoverEndpointStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The lifecycle phase of the switchover endpoint:   PROVISIONING:    the endpoint is being created and reconciled;   READY:           the endpoint is reconciled and serving traffic;   FAILED:          the endpoint encountered an error;   DEPROVISIONING:  the endpoint is being deleted.  | [readonly] 
**Conditions** | Pointer to [**[]SwitchoverV1SwitchoverEndpointCondition**](SwitchoverV1SwitchoverEndpointCondition.md) | Status conditions providing detailed information about the switchover endpoint&#39;s current state. | [optional] [readonly] 

## Methods

### NewSwitchoverV1SwitchoverEndpointStatus

`func NewSwitchoverV1SwitchoverEndpointStatus(phase string, ) *SwitchoverV1SwitchoverEndpointStatus`

NewSwitchoverV1SwitchoverEndpointStatus instantiates a new SwitchoverV1SwitchoverEndpointStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchoverV1SwitchoverEndpointStatusWithDefaults

`func NewSwitchoverV1SwitchoverEndpointStatusWithDefaults() *SwitchoverV1SwitchoverEndpointStatus`

NewSwitchoverV1SwitchoverEndpointStatusWithDefaults instantiates a new SwitchoverV1SwitchoverEndpointStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *SwitchoverV1SwitchoverEndpointStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *SwitchoverV1SwitchoverEndpointStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *SwitchoverV1SwitchoverEndpointStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetConditions

`func (o *SwitchoverV1SwitchoverEndpointStatus) GetConditions() []SwitchoverV1SwitchoverEndpointCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *SwitchoverV1SwitchoverEndpointStatus) GetConditionsOk() (*[]SwitchoverV1SwitchoverEndpointCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *SwitchoverV1SwitchoverEndpointStatus) SetConditions(v []SwitchoverV1SwitchoverEndpointCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *SwitchoverV1SwitchoverEndpointStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


