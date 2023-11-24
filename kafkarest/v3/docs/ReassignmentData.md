# ReassignmentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**ClusterId** | **string** |  | 
**TopicName** | **string** |  | 
**PartitionId** | **int32** |  | 
**AddingReplicas** | **[]int32** |  | 
**RemovingReplicas** | **[]int32** |  | 
**Replicas** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewReassignmentData

`func NewReassignmentData(kind string, metadata ResourceMetadata, clusterId string, topicName string, partitionId int32, addingReplicas []int32, removingReplicas []int32, replicas Relationship, ) *ReassignmentData`

NewReassignmentData instantiates a new ReassignmentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReassignmentDataWithDefaults

`func NewReassignmentDataWithDefaults() *ReassignmentData`

NewReassignmentDataWithDefaults instantiates a new ReassignmentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ReassignmentData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ReassignmentData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ReassignmentData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ReassignmentData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ReassignmentData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ReassignmentData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *ReassignmentData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ReassignmentData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ReassignmentData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetTopicName

`func (o *ReassignmentData) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *ReassignmentData) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *ReassignmentData) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.


### GetPartitionId

`func (o *ReassignmentData) GetPartitionId() int32`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *ReassignmentData) GetPartitionIdOk() (*int32, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *ReassignmentData) SetPartitionId(v int32)`

SetPartitionId sets PartitionId field to given value.


### GetAddingReplicas

`func (o *ReassignmentData) GetAddingReplicas() []int32`

GetAddingReplicas returns the AddingReplicas field if non-nil, zero value otherwise.

### GetAddingReplicasOk

`func (o *ReassignmentData) GetAddingReplicasOk() (*[]int32, bool)`

GetAddingReplicasOk returns a tuple with the AddingReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddingReplicas

`func (o *ReassignmentData) SetAddingReplicas(v []int32)`

SetAddingReplicas sets AddingReplicas field to given value.


### GetRemovingReplicas

`func (o *ReassignmentData) GetRemovingReplicas() []int32`

GetRemovingReplicas returns the RemovingReplicas field if non-nil, zero value otherwise.

### GetRemovingReplicasOk

`func (o *ReassignmentData) GetRemovingReplicasOk() (*[]int32, bool)`

GetRemovingReplicasOk returns a tuple with the RemovingReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovingReplicas

`func (o *ReassignmentData) SetRemovingReplicas(v []int32)`

SetRemovingReplicas sets RemovingReplicas field to given value.


### GetReplicas

`func (o *ReassignmentData) GetReplicas() Relationship`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *ReassignmentData) GetReplicasOk() (*Relationship, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *ReassignmentData) SetReplicas(v Relationship)`

SetReplicas sets Replicas field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


