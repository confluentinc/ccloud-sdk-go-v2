# StreamsGroupMemberAssignmentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**ClusterId** | **string** | The unique identifier of the Kafka cluster. | 
**GroupId** | **string** | The unique identifier of the Streams group. | 
**MemberId** | **string** | The unique identifier of the Streams group member. | 
**ActiveTasks** | [**Relationship**](Relationship.md) |  | 
**StandbyTasks** | [**Relationship**](Relationship.md) |  | 
**WarmupTasks** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewStreamsGroupMemberAssignmentData

`func NewStreamsGroupMemberAssignmentData(kind string, metadata ResourceMetadata, clusterId string, groupId string, memberId string, activeTasks Relationship, standbyTasks Relationship, warmupTasks Relationship, ) *StreamsGroupMemberAssignmentData`

NewStreamsGroupMemberAssignmentData instantiates a new StreamsGroupMemberAssignmentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsGroupMemberAssignmentDataWithDefaults

`func NewStreamsGroupMemberAssignmentDataWithDefaults() *StreamsGroupMemberAssignmentData`

NewStreamsGroupMemberAssignmentDataWithDefaults instantiates a new StreamsGroupMemberAssignmentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *StreamsGroupMemberAssignmentData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StreamsGroupMemberAssignmentData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StreamsGroupMemberAssignmentData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *StreamsGroupMemberAssignmentData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StreamsGroupMemberAssignmentData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StreamsGroupMemberAssignmentData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *StreamsGroupMemberAssignmentData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *StreamsGroupMemberAssignmentData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *StreamsGroupMemberAssignmentData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetGroupId

`func (o *StreamsGroupMemberAssignmentData) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *StreamsGroupMemberAssignmentData) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *StreamsGroupMemberAssignmentData) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetMemberId

`func (o *StreamsGroupMemberAssignmentData) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *StreamsGroupMemberAssignmentData) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *StreamsGroupMemberAssignmentData) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.


### GetActiveTasks

`func (o *StreamsGroupMemberAssignmentData) GetActiveTasks() Relationship`

GetActiveTasks returns the ActiveTasks field if non-nil, zero value otherwise.

### GetActiveTasksOk

`func (o *StreamsGroupMemberAssignmentData) GetActiveTasksOk() (*Relationship, bool)`

GetActiveTasksOk returns a tuple with the ActiveTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTasks

`func (o *StreamsGroupMemberAssignmentData) SetActiveTasks(v Relationship)`

SetActiveTasks sets ActiveTasks field to given value.


### GetStandbyTasks

`func (o *StreamsGroupMemberAssignmentData) GetStandbyTasks() Relationship`

GetStandbyTasks returns the StandbyTasks field if non-nil, zero value otherwise.

### GetStandbyTasksOk

`func (o *StreamsGroupMemberAssignmentData) GetStandbyTasksOk() (*Relationship, bool)`

GetStandbyTasksOk returns a tuple with the StandbyTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyTasks

`func (o *StreamsGroupMemberAssignmentData) SetStandbyTasks(v Relationship)`

SetStandbyTasks sets StandbyTasks field to given value.


### GetWarmupTasks

`func (o *StreamsGroupMemberAssignmentData) GetWarmupTasks() Relationship`

GetWarmupTasks returns the WarmupTasks field if non-nil, zero value otherwise.

### GetWarmupTasksOk

`func (o *StreamsGroupMemberAssignmentData) GetWarmupTasksOk() (*Relationship, bool)`

GetWarmupTasksOk returns a tuple with the WarmupTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarmupTasks

`func (o *StreamsGroupMemberAssignmentData) SetWarmupTasks(v Relationship)`

SetWarmupTasks sets WarmupTasks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


