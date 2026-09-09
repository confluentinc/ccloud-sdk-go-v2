# SwitchoverV1SwitchoverPairFailoverRequestSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentCrn** | **string** | The CRN of the switchover pair&#39;s environment. Required; unlike the other operations, this endpoint takes it from the request body rather than a query parameter.  | 
**ActiveMember** | Pointer to **string** | The name of the member to promote to active. If type is PLANNED/UNPLANNED, then &#x60;active_member&#x60; must be provided.  | [optional] 
**FailoverType** | **string** | The failover semantics to apply. Required — this field decides whether the operation can lose data, so it must be stated on every call rather than defaulted.   PLANNED:    graceful failover after replication lag reaches zero (cluster link is reversed);   UNPLANNED:  immediate failover without waiting for replication (cluster link is stopped);   RESTORE:    re-establish the cluster link after an unplanned failover.  | 

## Methods

### NewSwitchoverV1SwitchoverPairFailoverRequestSpec

`func NewSwitchoverV1SwitchoverPairFailoverRequestSpec(environmentCrn string, failoverType string, ) *SwitchoverV1SwitchoverPairFailoverRequestSpec`

NewSwitchoverV1SwitchoverPairFailoverRequestSpec instantiates a new SwitchoverV1SwitchoverPairFailoverRequestSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchoverV1SwitchoverPairFailoverRequestSpecWithDefaults

`func NewSwitchoverV1SwitchoverPairFailoverRequestSpecWithDefaults() *SwitchoverV1SwitchoverPairFailoverRequestSpec`

NewSwitchoverV1SwitchoverPairFailoverRequestSpecWithDefaults instantiates a new SwitchoverV1SwitchoverPairFailoverRequestSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentCrn

`func (o *SwitchoverV1SwitchoverPairFailoverRequestSpec) GetEnvironmentCrn() string`

GetEnvironmentCrn returns the EnvironmentCrn field if non-nil, zero value otherwise.

### GetEnvironmentCrnOk

`func (o *SwitchoverV1SwitchoverPairFailoverRequestSpec) GetEnvironmentCrnOk() (*string, bool)`

GetEnvironmentCrnOk returns a tuple with the EnvironmentCrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentCrn

`func (o *SwitchoverV1SwitchoverPairFailoverRequestSpec) SetEnvironmentCrn(v string)`

SetEnvironmentCrn sets EnvironmentCrn field to given value.


### GetActiveMember

`func (o *SwitchoverV1SwitchoverPairFailoverRequestSpec) GetActiveMember() string`

GetActiveMember returns the ActiveMember field if non-nil, zero value otherwise.

### GetActiveMemberOk

`func (o *SwitchoverV1SwitchoverPairFailoverRequestSpec) GetActiveMemberOk() (*string, bool)`

GetActiveMemberOk returns a tuple with the ActiveMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveMember

`func (o *SwitchoverV1SwitchoverPairFailoverRequestSpec) SetActiveMember(v string)`

SetActiveMember sets ActiveMember field to given value.

### HasActiveMember

`func (o *SwitchoverV1SwitchoverPairFailoverRequestSpec) HasActiveMember() bool`

HasActiveMember returns a boolean if a field has been set.

### GetFailoverType

`func (o *SwitchoverV1SwitchoverPairFailoverRequestSpec) GetFailoverType() string`

GetFailoverType returns the FailoverType field if non-nil, zero value otherwise.

### GetFailoverTypeOk

`func (o *SwitchoverV1SwitchoverPairFailoverRequestSpec) GetFailoverTypeOk() (*string, bool)`

GetFailoverTypeOk returns a tuple with the FailoverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverType

`func (o *SwitchoverV1SwitchoverPairFailoverRequestSpec) SetFailoverType(v string)`

SetFailoverType sets FailoverType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


