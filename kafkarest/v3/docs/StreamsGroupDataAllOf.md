# StreamsGroupDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** | The unique identifier of the Kafka cluster. | 
**GroupId** | **string** | The unique identifier of the Streams group. | 
**State** | **string** | The state of the Streams group. | 
**MemberCount** | **int32** | The number of members in the Streams group. | 
**SubtopologyCount** | **int32** | The number of subtopologies in the Streams group. | 
**GroupEpoch** | **int32** | The epoch of the Streams group. | 
**TopologyEpoch** | **int32** | The epoch of the Streams topology. | 
**TargetAssignmentEpoch** | **int32** | The epoch of the target assignment. | 
**Members** | [**Relationship**](Relationship.md) |  | 
**Subtopologies** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewStreamsGroupDataAllOf

`func NewStreamsGroupDataAllOf(clusterId string, groupId string, state string, memberCount int32, subtopologyCount int32, groupEpoch int32, topologyEpoch int32, targetAssignmentEpoch int32, members Relationship, subtopologies Relationship, ) *StreamsGroupDataAllOf`

NewStreamsGroupDataAllOf instantiates a new StreamsGroupDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsGroupDataAllOfWithDefaults

`func NewStreamsGroupDataAllOfWithDefaults() *StreamsGroupDataAllOf`

NewStreamsGroupDataAllOfWithDefaults instantiates a new StreamsGroupDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *StreamsGroupDataAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *StreamsGroupDataAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *StreamsGroupDataAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetGroupId

`func (o *StreamsGroupDataAllOf) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *StreamsGroupDataAllOf) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *StreamsGroupDataAllOf) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetState

`func (o *StreamsGroupDataAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StreamsGroupDataAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StreamsGroupDataAllOf) SetState(v string)`

SetState sets State field to given value.


### GetMemberCount

`func (o *StreamsGroupDataAllOf) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *StreamsGroupDataAllOf) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *StreamsGroupDataAllOf) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.


### GetSubtopologyCount

`func (o *StreamsGroupDataAllOf) GetSubtopologyCount() int32`

GetSubtopologyCount returns the SubtopologyCount field if non-nil, zero value otherwise.

### GetSubtopologyCountOk

`func (o *StreamsGroupDataAllOf) GetSubtopologyCountOk() (*int32, bool)`

GetSubtopologyCountOk returns a tuple with the SubtopologyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtopologyCount

`func (o *StreamsGroupDataAllOf) SetSubtopologyCount(v int32)`

SetSubtopologyCount sets SubtopologyCount field to given value.


### GetGroupEpoch

`func (o *StreamsGroupDataAllOf) GetGroupEpoch() int32`

GetGroupEpoch returns the GroupEpoch field if non-nil, zero value otherwise.

### GetGroupEpochOk

`func (o *StreamsGroupDataAllOf) GetGroupEpochOk() (*int32, bool)`

GetGroupEpochOk returns a tuple with the GroupEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupEpoch

`func (o *StreamsGroupDataAllOf) SetGroupEpoch(v int32)`

SetGroupEpoch sets GroupEpoch field to given value.


### GetTopologyEpoch

`func (o *StreamsGroupDataAllOf) GetTopologyEpoch() int32`

GetTopologyEpoch returns the TopologyEpoch field if non-nil, zero value otherwise.

### GetTopologyEpochOk

`func (o *StreamsGroupDataAllOf) GetTopologyEpochOk() (*int32, bool)`

GetTopologyEpochOk returns a tuple with the TopologyEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyEpoch

`func (o *StreamsGroupDataAllOf) SetTopologyEpoch(v int32)`

SetTopologyEpoch sets TopologyEpoch field to given value.


### GetTargetAssignmentEpoch

`func (o *StreamsGroupDataAllOf) GetTargetAssignmentEpoch() int32`

GetTargetAssignmentEpoch returns the TargetAssignmentEpoch field if non-nil, zero value otherwise.

### GetTargetAssignmentEpochOk

`func (o *StreamsGroupDataAllOf) GetTargetAssignmentEpochOk() (*int32, bool)`

GetTargetAssignmentEpochOk returns a tuple with the TargetAssignmentEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAssignmentEpoch

`func (o *StreamsGroupDataAllOf) SetTargetAssignmentEpoch(v int32)`

SetTargetAssignmentEpoch sets TargetAssignmentEpoch field to given value.


### GetMembers

`func (o *StreamsGroupDataAllOf) GetMembers() Relationship`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *StreamsGroupDataAllOf) GetMembersOk() (*Relationship, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *StreamsGroupDataAllOf) SetMembers(v Relationship)`

SetMembers sets Members field to given value.


### GetSubtopologies

`func (o *StreamsGroupDataAllOf) GetSubtopologies() Relationship`

GetSubtopologies returns the Subtopologies field if non-nil, zero value otherwise.

### GetSubtopologiesOk

`func (o *StreamsGroupDataAllOf) GetSubtopologiesOk() (*Relationship, bool)`

GetSubtopologiesOk returns a tuple with the Subtopologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtopologies

`func (o *StreamsGroupDataAllOf) SetSubtopologies(v Relationship)`

SetSubtopologies sets Subtopologies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


