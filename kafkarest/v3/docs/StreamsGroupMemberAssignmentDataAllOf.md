# StreamsGroupMemberAssignmentDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** | The unique identifier of the Kafka cluster. | 
**GroupId** | **string** | The unique identifier of the Streams group. | 
**MemberId** | **string** | The unique identifier of the Streams group member. | 
**ActiveTasks** | [**Relationship**](Relationship.md) |  | 
**StandbyTasks** | [**Relationship**](Relationship.md) |  | 
**WarmupTasks** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewStreamsGroupMemberAssignmentDataAllOf

`func NewStreamsGroupMemberAssignmentDataAllOf(clusterId string, groupId string, memberId string, activeTasks Relationship, standbyTasks Relationship, warmupTasks Relationship, ) *StreamsGroupMemberAssignmentDataAllOf`

NewStreamsGroupMemberAssignmentDataAllOf instantiates a new StreamsGroupMemberAssignmentDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsGroupMemberAssignmentDataAllOfWithDefaults

`func NewStreamsGroupMemberAssignmentDataAllOfWithDefaults() *StreamsGroupMemberAssignmentDataAllOf`

NewStreamsGroupMemberAssignmentDataAllOfWithDefaults instantiates a new StreamsGroupMemberAssignmentDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *StreamsGroupMemberAssignmentDataAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *StreamsGroupMemberAssignmentDataAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *StreamsGroupMemberAssignmentDataAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetGroupId

`func (o *StreamsGroupMemberAssignmentDataAllOf) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *StreamsGroupMemberAssignmentDataAllOf) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *StreamsGroupMemberAssignmentDataAllOf) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetMemberId

`func (o *StreamsGroupMemberAssignmentDataAllOf) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *StreamsGroupMemberAssignmentDataAllOf) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *StreamsGroupMemberAssignmentDataAllOf) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.


### GetActiveTasks

`func (o *StreamsGroupMemberAssignmentDataAllOf) GetActiveTasks() Relationship`

GetActiveTasks returns the ActiveTasks field if non-nil, zero value otherwise.

### GetActiveTasksOk

`func (o *StreamsGroupMemberAssignmentDataAllOf) GetActiveTasksOk() (*Relationship, bool)`

GetActiveTasksOk returns a tuple with the ActiveTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTasks

`func (o *StreamsGroupMemberAssignmentDataAllOf) SetActiveTasks(v Relationship)`

SetActiveTasks sets ActiveTasks field to given value.


### GetStandbyTasks

`func (o *StreamsGroupMemberAssignmentDataAllOf) GetStandbyTasks() Relationship`

GetStandbyTasks returns the StandbyTasks field if non-nil, zero value otherwise.

### GetStandbyTasksOk

`func (o *StreamsGroupMemberAssignmentDataAllOf) GetStandbyTasksOk() (*Relationship, bool)`

GetStandbyTasksOk returns a tuple with the StandbyTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyTasks

`func (o *StreamsGroupMemberAssignmentDataAllOf) SetStandbyTasks(v Relationship)`

SetStandbyTasks sets StandbyTasks field to given value.


### GetWarmupTasks

`func (o *StreamsGroupMemberAssignmentDataAllOf) GetWarmupTasks() Relationship`

GetWarmupTasks returns the WarmupTasks field if non-nil, zero value otherwise.

### GetWarmupTasksOk

`func (o *StreamsGroupMemberAssignmentDataAllOf) GetWarmupTasksOk() (*Relationship, bool)`

GetWarmupTasksOk returns a tuple with the WarmupTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarmupTasks

`func (o *StreamsGroupMemberAssignmentDataAllOf) SetWarmupTasks(v Relationship)`

SetWarmupTasks sets WarmupTasks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


