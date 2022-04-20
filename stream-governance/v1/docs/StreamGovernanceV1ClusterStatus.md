# StreamGovernanceV1ClusterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The lifecyle phase of the cluster:   PROVISIONED:  cluster is provisioned;   PROVISIONING:  cluster provisioning is in progress;   FAILED:  provisioning failed  | [readonly] 

## Methods

### NewStreamGovernanceV1ClusterStatus

`func NewStreamGovernanceV1ClusterStatus(phase string, ) *StreamGovernanceV1ClusterStatus`

NewStreamGovernanceV1ClusterStatus instantiates a new StreamGovernanceV1ClusterStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamGovernanceV1ClusterStatusWithDefaults

`func NewStreamGovernanceV1ClusterStatusWithDefaults() *StreamGovernanceV1ClusterStatus`

NewStreamGovernanceV1ClusterStatusWithDefaults instantiates a new StreamGovernanceV1ClusterStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *StreamGovernanceV1ClusterStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *StreamGovernanceV1ClusterStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *StreamGovernanceV1ClusterStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


