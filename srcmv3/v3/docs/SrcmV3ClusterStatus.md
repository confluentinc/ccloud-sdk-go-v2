# SrcmV3ClusterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The lifecyle phase of the cluster:    PROVISIONED:  cluster is provisioned;    PROVISIONING:  cluster provisioning is in progress;    FAILED:  provisioning failed  | [readonly] 

## Methods

### NewSrcmV3ClusterStatus

`func NewSrcmV3ClusterStatus(phase string, ) *SrcmV3ClusterStatus`

NewSrcmV3ClusterStatus instantiates a new SrcmV3ClusterStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSrcmV3ClusterStatusWithDefaults

`func NewSrcmV3ClusterStatusWithDefaults() *SrcmV3ClusterStatus`

NewSrcmV3ClusterStatusWithDefaults instantiates a new SrcmV3ClusterStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *SrcmV3ClusterStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *SrcmV3ClusterStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *SrcmV3ClusterStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


