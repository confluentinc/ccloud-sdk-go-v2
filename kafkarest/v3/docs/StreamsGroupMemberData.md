# StreamsGroupMemberData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
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

### NewStreamsGroupMemberData

`func NewStreamsGroupMemberData(kind string, metadata ResourceMetadata, clusterId string, groupId string, memberId string, processId string, clientId string, instanceId string, memberEpoch int32, topologyEpoch int32, isClassic bool, assignments Relationship, targetAssignment Relationship, ) *StreamsGroupMemberData`

NewStreamsGroupMemberData instantiates a new StreamsGroupMemberData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsGroupMemberDataWithDefaults

`func NewStreamsGroupMemberDataWithDefaults() *StreamsGroupMemberData`

NewStreamsGroupMemberDataWithDefaults instantiates a new StreamsGroupMemberData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *StreamsGroupMemberData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StreamsGroupMemberData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StreamsGroupMemberData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *StreamsGroupMemberData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StreamsGroupMemberData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StreamsGroupMemberData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *StreamsGroupMemberData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *StreamsGroupMemberData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *StreamsGroupMemberData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetGroupId

`func (o *StreamsGroupMemberData) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *StreamsGroupMemberData) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *StreamsGroupMemberData) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetMemberId

`func (o *StreamsGroupMemberData) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *StreamsGroupMemberData) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *StreamsGroupMemberData) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.


### GetProcessId

`func (o *StreamsGroupMemberData) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *StreamsGroupMemberData) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *StreamsGroupMemberData) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.


### GetClientId

`func (o *StreamsGroupMemberData) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *StreamsGroupMemberData) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *StreamsGroupMemberData) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetInstanceId

`func (o *StreamsGroupMemberData) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *StreamsGroupMemberData) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *StreamsGroupMemberData) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetMemberEpoch

`func (o *StreamsGroupMemberData) GetMemberEpoch() int32`

GetMemberEpoch returns the MemberEpoch field if non-nil, zero value otherwise.

### GetMemberEpochOk

`func (o *StreamsGroupMemberData) GetMemberEpochOk() (*int32, bool)`

GetMemberEpochOk returns a tuple with the MemberEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberEpoch

`func (o *StreamsGroupMemberData) SetMemberEpoch(v int32)`

SetMemberEpoch sets MemberEpoch field to given value.


### GetTopologyEpoch

`func (o *StreamsGroupMemberData) GetTopologyEpoch() int32`

GetTopologyEpoch returns the TopologyEpoch field if non-nil, zero value otherwise.

### GetTopologyEpochOk

`func (o *StreamsGroupMemberData) GetTopologyEpochOk() (*int32, bool)`

GetTopologyEpochOk returns a tuple with the TopologyEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyEpoch

`func (o *StreamsGroupMemberData) SetTopologyEpoch(v int32)`

SetTopologyEpoch sets TopologyEpoch field to given value.


### GetIsClassic

`func (o *StreamsGroupMemberData) GetIsClassic() bool`

GetIsClassic returns the IsClassic field if non-nil, zero value otherwise.

### GetIsClassicOk

`func (o *StreamsGroupMemberData) GetIsClassicOk() (*bool, bool)`

GetIsClassicOk returns a tuple with the IsClassic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClassic

`func (o *StreamsGroupMemberData) SetIsClassic(v bool)`

SetIsClassic sets IsClassic field to given value.


### GetAssignments

`func (o *StreamsGroupMemberData) GetAssignments() Relationship`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *StreamsGroupMemberData) GetAssignmentsOk() (*Relationship, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *StreamsGroupMemberData) SetAssignments(v Relationship)`

SetAssignments sets Assignments field to given value.


### GetTargetAssignment

`func (o *StreamsGroupMemberData) GetTargetAssignment() Relationship`

GetTargetAssignment returns the TargetAssignment field if non-nil, zero value otherwise.

### GetTargetAssignmentOk

`func (o *StreamsGroupMemberData) GetTargetAssignmentOk() (*Relationship, bool)`

GetTargetAssignmentOk returns a tuple with the TargetAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAssignment

`func (o *StreamsGroupMemberData) SetTargetAssignment(v Relationship)`

SetTargetAssignment sets TargetAssignment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


