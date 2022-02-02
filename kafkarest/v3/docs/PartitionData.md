# PartitionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**ClusterId** | **string** |  | 
**TopicName** | **string** |  | 
**PartitionId** | **int32** |  | 
**Leader** | Pointer to [**Relationship**](Relationship.md) |  | [optional] 
**Replicas** | [**Relationship**](Relationship.md) |  | 
**Reassignment** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewPartitionData

`func NewPartitionData(kind string, metadata ResourceMetadata, clusterId string, topicName string, partitionId int32, replicas Relationship, reassignment Relationship, ) *PartitionData`

NewPartitionData instantiates a new PartitionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionDataWithDefaults

`func NewPartitionDataWithDefaults() *PartitionData`

NewPartitionDataWithDefaults instantiates a new PartitionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *PartitionData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PartitionData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PartitionData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *PartitionData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PartitionData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PartitionData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *PartitionData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *PartitionData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *PartitionData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetTopicName

`func (o *PartitionData) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *PartitionData) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *PartitionData) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.


### GetPartitionId

`func (o *PartitionData) GetPartitionId() int32`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *PartitionData) GetPartitionIdOk() (*int32, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *PartitionData) SetPartitionId(v int32)`

SetPartitionId sets PartitionId field to given value.


### GetLeader

`func (o *PartitionData) GetLeader() Relationship`

GetLeader returns the Leader field if non-nil, zero value otherwise.

### GetLeaderOk

`func (o *PartitionData) GetLeaderOk() (*Relationship, bool)`

GetLeaderOk returns a tuple with the Leader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeader

`func (o *PartitionData) SetLeader(v Relationship)`

SetLeader sets Leader field to given value.

### HasLeader

`func (o *PartitionData) HasLeader() bool`

HasLeader returns a boolean if a field has been set.

### GetReplicas

`func (o *PartitionData) GetReplicas() Relationship`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *PartitionData) GetReplicasOk() (*Relationship, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *PartitionData) SetReplicas(v Relationship)`

SetReplicas sets Replicas field to given value.


### GetReassignment

`func (o *PartitionData) GetReassignment() Relationship`

GetReassignment returns the Reassignment field if non-nil, zero value otherwise.

### GetReassignmentOk

`func (o *PartitionData) GetReassignmentOk() (*Relationship, bool)`

GetReassignmentOk returns a tuple with the Reassignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReassignment

`func (o *PartitionData) SetReassignment(v Relationship)`

SetReassignment sets Reassignment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


