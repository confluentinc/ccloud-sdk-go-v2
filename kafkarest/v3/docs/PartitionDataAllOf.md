# PartitionDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**TopicName** | **string** |  | 
**PartitionId** | **int32** |  | 
**Leader** | Pointer to [**Relationship**](Relationship.md) |  | [optional] 
**Replicas** | [**Relationship**](Relationship.md) |  | 
**Reassignment** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewPartitionDataAllOf

`func NewPartitionDataAllOf(clusterId string, topicName string, partitionId int32, replicas Relationship, reassignment Relationship, ) *PartitionDataAllOf`

NewPartitionDataAllOf instantiates a new PartitionDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionDataAllOfWithDefaults

`func NewPartitionDataAllOfWithDefaults() *PartitionDataAllOf`

NewPartitionDataAllOfWithDefaults instantiates a new PartitionDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *PartitionDataAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *PartitionDataAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *PartitionDataAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetTopicName

`func (o *PartitionDataAllOf) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *PartitionDataAllOf) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *PartitionDataAllOf) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.


### GetPartitionId

`func (o *PartitionDataAllOf) GetPartitionId() int32`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *PartitionDataAllOf) GetPartitionIdOk() (*int32, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *PartitionDataAllOf) SetPartitionId(v int32)`

SetPartitionId sets PartitionId field to given value.


### GetLeader

`func (o *PartitionDataAllOf) GetLeader() Relationship`

GetLeader returns the Leader field if non-nil, zero value otherwise.

### GetLeaderOk

`func (o *PartitionDataAllOf) GetLeaderOk() (*Relationship, bool)`

GetLeaderOk returns a tuple with the Leader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeader

`func (o *PartitionDataAllOf) SetLeader(v Relationship)`

SetLeader sets Leader field to given value.

### HasLeader

`func (o *PartitionDataAllOf) HasLeader() bool`

HasLeader returns a boolean if a field has been set.

### GetReplicas

`func (o *PartitionDataAllOf) GetReplicas() Relationship`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *PartitionDataAllOf) GetReplicasOk() (*Relationship, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *PartitionDataAllOf) SetReplicas(v Relationship)`

SetReplicas sets Replicas field to given value.


### GetReassignment

`func (o *PartitionDataAllOf) GetReassignment() Relationship`

GetReassignment returns the Reassignment field if non-nil, zero value otherwise.

### GetReassignmentOk

`func (o *PartitionDataAllOf) GetReassignmentOk() (*Relationship, bool)`

GetReassignmentOk returns a tuple with the Reassignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReassignment

`func (o *PartitionDataAllOf) SetReassignment(v Relationship)`

SetReassignment sets Reassignment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


