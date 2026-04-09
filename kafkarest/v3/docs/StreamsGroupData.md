# StreamsGroupData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
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

### NewStreamsGroupData

`func NewStreamsGroupData(kind string, metadata ResourceMetadata, clusterId string, groupId string, state string, memberCount int32, subtopologyCount int32, groupEpoch int32, topologyEpoch int32, targetAssignmentEpoch int32, members Relationship, subtopologies Relationship, ) *StreamsGroupData`

NewStreamsGroupData instantiates a new StreamsGroupData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsGroupDataWithDefaults

`func NewStreamsGroupDataWithDefaults() *StreamsGroupData`

NewStreamsGroupDataWithDefaults instantiates a new StreamsGroupData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *StreamsGroupData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StreamsGroupData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StreamsGroupData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *StreamsGroupData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StreamsGroupData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StreamsGroupData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *StreamsGroupData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *StreamsGroupData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *StreamsGroupData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetGroupId

`func (o *StreamsGroupData) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *StreamsGroupData) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *StreamsGroupData) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetState

`func (o *StreamsGroupData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StreamsGroupData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StreamsGroupData) SetState(v string)`

SetState sets State field to given value.


### GetMemberCount

`func (o *StreamsGroupData) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *StreamsGroupData) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *StreamsGroupData) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.


### GetSubtopologyCount

`func (o *StreamsGroupData) GetSubtopologyCount() int32`

GetSubtopologyCount returns the SubtopologyCount field if non-nil, zero value otherwise.

### GetSubtopologyCountOk

`func (o *StreamsGroupData) GetSubtopologyCountOk() (*int32, bool)`

GetSubtopologyCountOk returns a tuple with the SubtopologyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtopologyCount

`func (o *StreamsGroupData) SetSubtopologyCount(v int32)`

SetSubtopologyCount sets SubtopologyCount field to given value.


### GetGroupEpoch

`func (o *StreamsGroupData) GetGroupEpoch() int32`

GetGroupEpoch returns the GroupEpoch field if non-nil, zero value otherwise.

### GetGroupEpochOk

`func (o *StreamsGroupData) GetGroupEpochOk() (*int32, bool)`

GetGroupEpochOk returns a tuple with the GroupEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupEpoch

`func (o *StreamsGroupData) SetGroupEpoch(v int32)`

SetGroupEpoch sets GroupEpoch field to given value.


### GetTopologyEpoch

`func (o *StreamsGroupData) GetTopologyEpoch() int32`

GetTopologyEpoch returns the TopologyEpoch field if non-nil, zero value otherwise.

### GetTopologyEpochOk

`func (o *StreamsGroupData) GetTopologyEpochOk() (*int32, bool)`

GetTopologyEpochOk returns a tuple with the TopologyEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyEpoch

`func (o *StreamsGroupData) SetTopologyEpoch(v int32)`

SetTopologyEpoch sets TopologyEpoch field to given value.


### GetTargetAssignmentEpoch

`func (o *StreamsGroupData) GetTargetAssignmentEpoch() int32`

GetTargetAssignmentEpoch returns the TargetAssignmentEpoch field if non-nil, zero value otherwise.

### GetTargetAssignmentEpochOk

`func (o *StreamsGroupData) GetTargetAssignmentEpochOk() (*int32, bool)`

GetTargetAssignmentEpochOk returns a tuple with the TargetAssignmentEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAssignmentEpoch

`func (o *StreamsGroupData) SetTargetAssignmentEpoch(v int32)`

SetTargetAssignmentEpoch sets TargetAssignmentEpoch field to given value.


### GetMembers

`func (o *StreamsGroupData) GetMembers() Relationship`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *StreamsGroupData) GetMembersOk() (*Relationship, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *StreamsGroupData) SetMembers(v Relationship)`

SetMembers sets Members field to given value.


### GetSubtopologies

`func (o *StreamsGroupData) GetSubtopologies() Relationship`

GetSubtopologies returns the Subtopologies field if non-nil, zero value otherwise.

### GetSubtopologiesOk

`func (o *StreamsGroupData) GetSubtopologiesOk() (*Relationship, bool)`

GetSubtopologiesOk returns a tuple with the Subtopologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtopologies

`func (o *StreamsGroupData) SetSubtopologies(v Relationship)`

SetSubtopologies sets Subtopologies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


