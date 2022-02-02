# ReassignmentDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**TopicName** | **string** |  | 
**PartitionId** | **int32** |  | 
**AddingReplicas** | **[]int32** |  | 
**RemovingReplicas** | **[]int32** |  | 
**Replicas** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewReassignmentDataAllOf

`func NewReassignmentDataAllOf(clusterId string, topicName string, partitionId int32, addingReplicas []int32, removingReplicas []int32, replicas Relationship, ) *ReassignmentDataAllOf`

NewReassignmentDataAllOf instantiates a new ReassignmentDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReassignmentDataAllOfWithDefaults

`func NewReassignmentDataAllOfWithDefaults() *ReassignmentDataAllOf`

NewReassignmentDataAllOfWithDefaults instantiates a new ReassignmentDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ReassignmentDataAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ReassignmentDataAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ReassignmentDataAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetTopicName

`func (o *ReassignmentDataAllOf) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *ReassignmentDataAllOf) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *ReassignmentDataAllOf) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.


### GetPartitionId

`func (o *ReassignmentDataAllOf) GetPartitionId() int32`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *ReassignmentDataAllOf) GetPartitionIdOk() (*int32, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *ReassignmentDataAllOf) SetPartitionId(v int32)`

SetPartitionId sets PartitionId field to given value.


### GetAddingReplicas

`func (o *ReassignmentDataAllOf) GetAddingReplicas() []int32`

GetAddingReplicas returns the AddingReplicas field if non-nil, zero value otherwise.

### GetAddingReplicasOk

`func (o *ReassignmentDataAllOf) GetAddingReplicasOk() (*[]int32, bool)`

GetAddingReplicasOk returns a tuple with the AddingReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddingReplicas

`func (o *ReassignmentDataAllOf) SetAddingReplicas(v []int32)`

SetAddingReplicas sets AddingReplicas field to given value.


### GetRemovingReplicas

`func (o *ReassignmentDataAllOf) GetRemovingReplicas() []int32`

GetRemovingReplicas returns the RemovingReplicas field if non-nil, zero value otherwise.

### GetRemovingReplicasOk

`func (o *ReassignmentDataAllOf) GetRemovingReplicasOk() (*[]int32, bool)`

GetRemovingReplicasOk returns a tuple with the RemovingReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovingReplicas

`func (o *ReassignmentDataAllOf) SetRemovingReplicas(v []int32)`

SetRemovingReplicas sets RemovingReplicas field to given value.


### GetReplicas

`func (o *ReassignmentDataAllOf) GetReplicas() Relationship`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *ReassignmentDataAllOf) GetReplicasOk() (*Relationship, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *ReassignmentDataAllOf) SetReplicas(v Relationship)`

SetReplicas sets Replicas field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


