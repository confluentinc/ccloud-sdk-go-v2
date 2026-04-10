# StreamsGroupMemberDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** | The unique identifier of the Kafka cluster. | 
**GroupId** | **string** | The unique identifier of the Streams group. | 
**MemberId** | **string** | The unique identifier of the Streams group member. | 
**ProcessId** | **string** | The process identifier of the Streams group member. | 
**ClientId** | **string** | The client identifier of the Streams group member. | 
**InstanceId** | **string** | The instance identifier of the Streams group member. | 
**MemberEpoch** | **int32** | The epoch of the Streams group member. | 
**TopologyEpoch** | **int32** | The epoch of the Streams topology for the member. | 
**IsClassic** | **bool** | The flag indicating if the member is a classic consumer. | 
**Assignments** | [**Relationship**](Relationship.md) |  | 
**TargetAssignment** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewStreamsGroupMemberDataAllOf

`func NewStreamsGroupMemberDataAllOf(clusterId string, groupId string, memberId string, processId string, clientId string, instanceId string, memberEpoch int32, topologyEpoch int32, isClassic bool, assignments Relationship, targetAssignment Relationship, ) *StreamsGroupMemberDataAllOf`

NewStreamsGroupMemberDataAllOf instantiates a new StreamsGroupMemberDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsGroupMemberDataAllOfWithDefaults

`func NewStreamsGroupMemberDataAllOfWithDefaults() *StreamsGroupMemberDataAllOf`

NewStreamsGroupMemberDataAllOfWithDefaults instantiates a new StreamsGroupMemberDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *StreamsGroupMemberDataAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *StreamsGroupMemberDataAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *StreamsGroupMemberDataAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetGroupId

`func (o *StreamsGroupMemberDataAllOf) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *StreamsGroupMemberDataAllOf) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *StreamsGroupMemberDataAllOf) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetMemberId

`func (o *StreamsGroupMemberDataAllOf) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *StreamsGroupMemberDataAllOf) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *StreamsGroupMemberDataAllOf) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.


### GetProcessId

`func (o *StreamsGroupMemberDataAllOf) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *StreamsGroupMemberDataAllOf) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *StreamsGroupMemberDataAllOf) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.


### GetClientId

`func (o *StreamsGroupMemberDataAllOf) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *StreamsGroupMemberDataAllOf) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *StreamsGroupMemberDataAllOf) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetInstanceId

`func (o *StreamsGroupMemberDataAllOf) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *StreamsGroupMemberDataAllOf) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *StreamsGroupMemberDataAllOf) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetMemberEpoch

`func (o *StreamsGroupMemberDataAllOf) GetMemberEpoch() int32`

GetMemberEpoch returns the MemberEpoch field if non-nil, zero value otherwise.

### GetMemberEpochOk

`func (o *StreamsGroupMemberDataAllOf) GetMemberEpochOk() (*int32, bool)`

GetMemberEpochOk returns a tuple with the MemberEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberEpoch

`func (o *StreamsGroupMemberDataAllOf) SetMemberEpoch(v int32)`

SetMemberEpoch sets MemberEpoch field to given value.


### GetTopologyEpoch

`func (o *StreamsGroupMemberDataAllOf) GetTopologyEpoch() int32`

GetTopologyEpoch returns the TopologyEpoch field if non-nil, zero value otherwise.

### GetTopologyEpochOk

`func (o *StreamsGroupMemberDataAllOf) GetTopologyEpochOk() (*int32, bool)`

GetTopologyEpochOk returns a tuple with the TopologyEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyEpoch

`func (o *StreamsGroupMemberDataAllOf) SetTopologyEpoch(v int32)`

SetTopologyEpoch sets TopologyEpoch field to given value.


### GetIsClassic

`func (o *StreamsGroupMemberDataAllOf) GetIsClassic() bool`

GetIsClassic returns the IsClassic field if non-nil, zero value otherwise.

### GetIsClassicOk

`func (o *StreamsGroupMemberDataAllOf) GetIsClassicOk() (*bool, bool)`

GetIsClassicOk returns a tuple with the IsClassic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClassic

`func (o *StreamsGroupMemberDataAllOf) SetIsClassic(v bool)`

SetIsClassic sets IsClassic field to given value.


### GetAssignments

`func (o *StreamsGroupMemberDataAllOf) GetAssignments() Relationship`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *StreamsGroupMemberDataAllOf) GetAssignmentsOk() (*Relationship, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *StreamsGroupMemberDataAllOf) SetAssignments(v Relationship)`

SetAssignments sets Assignments field to given value.


### GetTargetAssignment

`func (o *StreamsGroupMemberDataAllOf) GetTargetAssignment() Relationship`

GetTargetAssignment returns the TargetAssignment field if non-nil, zero value otherwise.

### GetTargetAssignmentOk

`func (o *StreamsGroupMemberDataAllOf) GetTargetAssignmentOk() (*Relationship, bool)`

GetTargetAssignmentOk returns a tuple with the TargetAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAssignment

`func (o *StreamsGroupMemberDataAllOf) SetTargetAssignment(v Relationship)`

SetTargetAssignment sets TargetAssignment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


