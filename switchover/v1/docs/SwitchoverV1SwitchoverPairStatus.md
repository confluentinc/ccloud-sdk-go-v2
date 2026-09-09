# SwitchoverV1SwitchoverPairStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The lifecycle phase of the switchover pair:   PROVISIONING:      the pair is being created and validated;   READY_TO_FAILOVER: the pair is validated and a failover (planned or unplanned) can be triggered;   UPDATING:          a failover, failback, or restore operation is in progress;   READY_TO_RESTORE:  an unplanned failover has completed; the cluster link can be restored;   FAILED:            resource validation failed for the pair; it must be deleted and recreated;   DEPROVISIONING:    the pair is being deleted.  | [readonly] 
**Conditions** | Pointer to [**[]SwitchoverV1Condition**](SwitchoverV1Condition.md) | Status conditions providing detailed information about the switchover pair&#39;s current state. | [optional] [readonly] 

## Methods

### NewSwitchoverV1SwitchoverPairStatus

`func NewSwitchoverV1SwitchoverPairStatus(phase string, ) *SwitchoverV1SwitchoverPairStatus`

NewSwitchoverV1SwitchoverPairStatus instantiates a new SwitchoverV1SwitchoverPairStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchoverV1SwitchoverPairStatusWithDefaults

`func NewSwitchoverV1SwitchoverPairStatusWithDefaults() *SwitchoverV1SwitchoverPairStatus`

NewSwitchoverV1SwitchoverPairStatusWithDefaults instantiates a new SwitchoverV1SwitchoverPairStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *SwitchoverV1SwitchoverPairStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *SwitchoverV1SwitchoverPairStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *SwitchoverV1SwitchoverPairStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetConditions

`func (o *SwitchoverV1SwitchoverPairStatus) GetConditions() []SwitchoverV1Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *SwitchoverV1SwitchoverPairStatus) GetConditionsOk() (*[]SwitchoverV1Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *SwitchoverV1SwitchoverPairStatus) SetConditions(v []SwitchoverV1Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *SwitchoverV1SwitchoverPairStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


