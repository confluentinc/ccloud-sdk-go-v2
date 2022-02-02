# ReplicaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**ClusterId** | **string** |  | 
**TopicName** | **string** |  | 
**PartitionId** | **int32** |  | 
**BrokerId** | **int32** |  | 
**IsLeader** | **bool** |  | 
**IsInSync** | **bool** |  | 
**Broker** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewReplicaData

`func NewReplicaData(kind string, metadata ResourceMetadata, clusterId string, topicName string, partitionId int32, brokerId int32, isLeader bool, isInSync bool, broker Relationship, ) *ReplicaData`

NewReplicaData instantiates a new ReplicaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicaDataWithDefaults

`func NewReplicaDataWithDefaults() *ReplicaData`

NewReplicaDataWithDefaults instantiates a new ReplicaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ReplicaData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ReplicaData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ReplicaData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ReplicaData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ReplicaData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ReplicaData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *ReplicaData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ReplicaData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ReplicaData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetTopicName

`func (o *ReplicaData) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *ReplicaData) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *ReplicaData) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.


### GetPartitionId

`func (o *ReplicaData) GetPartitionId() int32`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *ReplicaData) GetPartitionIdOk() (*int32, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *ReplicaData) SetPartitionId(v int32)`

SetPartitionId sets PartitionId field to given value.


### GetBrokerId

`func (o *ReplicaData) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *ReplicaData) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *ReplicaData) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.


### GetIsLeader

`func (o *ReplicaData) GetIsLeader() bool`

GetIsLeader returns the IsLeader field if non-nil, zero value otherwise.

### GetIsLeaderOk

`func (o *ReplicaData) GetIsLeaderOk() (*bool, bool)`

GetIsLeaderOk returns a tuple with the IsLeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLeader

`func (o *ReplicaData) SetIsLeader(v bool)`

SetIsLeader sets IsLeader field to given value.


### GetIsInSync

`func (o *ReplicaData) GetIsInSync() bool`

GetIsInSync returns the IsInSync field if non-nil, zero value otherwise.

### GetIsInSyncOk

`func (o *ReplicaData) GetIsInSyncOk() (*bool, bool)`

GetIsInSyncOk returns a tuple with the IsInSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInSync

`func (o *ReplicaData) SetIsInSync(v bool)`

SetIsInSync sets IsInSync field to given value.


### GetBroker

`func (o *ReplicaData) GetBroker() Relationship`

GetBroker returns the Broker field if non-nil, zero value otherwise.

### GetBrokerOk

`func (o *ReplicaData) GetBrokerOk() (*Relationship, bool)`

GetBrokerOk returns a tuple with the Broker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroker

`func (o *ReplicaData) SetBroker(v Relationship)`

SetBroker sets Broker field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


