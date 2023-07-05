# TopicData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**ClusterId** | **string** |  | 
**TopicName** | **string** |  | 
**IsInternal** | **bool** |  | 
**ReplicationFactor** | **int32** |  | 
**PartitionsCount** | **int32** |  | 
**Partitions** | [**Relationship**](Relationship.md) |  | 
**Configs** | [**Relationship**](Relationship.md) |  | 
**PartitionReassignments** | [**Relationship**](Relationship.md) |  | 
**AuthorizedOperations** | Pointer to [**AuthorizedOperations**](AuthorizedOperations.md) |  | [optional] 

## Methods

### NewTopicData

`func NewTopicData(kind string, metadata ResourceMetadata, clusterId string, topicName string, isInternal bool, replicationFactor int32, partitionsCount int32, partitions Relationship, configs Relationship, partitionReassignments Relationship, ) *TopicData`

NewTopicData instantiates a new TopicData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicDataWithDefaults

`func NewTopicDataWithDefaults() *TopicData`

NewTopicDataWithDefaults instantiates a new TopicData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *TopicData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TopicData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TopicData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *TopicData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TopicData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TopicData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *TopicData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *TopicData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *TopicData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetTopicName

`func (o *TopicData) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *TopicData) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *TopicData) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.


### GetIsInternal

`func (o *TopicData) GetIsInternal() bool`

GetIsInternal returns the IsInternal field if non-nil, zero value otherwise.

### GetIsInternalOk

`func (o *TopicData) GetIsInternalOk() (*bool, bool)`

GetIsInternalOk returns a tuple with the IsInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternal

`func (o *TopicData) SetIsInternal(v bool)`

SetIsInternal sets IsInternal field to given value.


### GetReplicationFactor

`func (o *TopicData) GetReplicationFactor() int32`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *TopicData) GetReplicationFactorOk() (*int32, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *TopicData) SetReplicationFactor(v int32)`

SetReplicationFactor sets ReplicationFactor field to given value.


### GetPartitionsCount

`func (o *TopicData) GetPartitionsCount() int32`

GetPartitionsCount returns the PartitionsCount field if non-nil, zero value otherwise.

### GetPartitionsCountOk

`func (o *TopicData) GetPartitionsCountOk() (*int32, bool)`

GetPartitionsCountOk returns a tuple with the PartitionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionsCount

`func (o *TopicData) SetPartitionsCount(v int32)`

SetPartitionsCount sets PartitionsCount field to given value.


### GetPartitions

`func (o *TopicData) GetPartitions() Relationship`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *TopicData) GetPartitionsOk() (*Relationship, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *TopicData) SetPartitions(v Relationship)`

SetPartitions sets Partitions field to given value.


### GetConfigs

`func (o *TopicData) GetConfigs() Relationship`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *TopicData) GetConfigsOk() (*Relationship, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *TopicData) SetConfigs(v Relationship)`

SetConfigs sets Configs field to given value.


### GetPartitionReassignments

`func (o *TopicData) GetPartitionReassignments() Relationship`

GetPartitionReassignments returns the PartitionReassignments field if non-nil, zero value otherwise.

### GetPartitionReassignmentsOk

`func (o *TopicData) GetPartitionReassignmentsOk() (*Relationship, bool)`

GetPartitionReassignmentsOk returns a tuple with the PartitionReassignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionReassignments

`func (o *TopicData) SetPartitionReassignments(v Relationship)`

SetPartitionReassignments sets PartitionReassignments field to given value.


### GetAuthorizedOperations

`func (o *TopicData) GetAuthorizedOperations() AuthorizedOperations`

GetAuthorizedOperations returns the AuthorizedOperations field if non-nil, zero value otherwise.

### GetAuthorizedOperationsOk

`func (o *TopicData) GetAuthorizedOperationsOk() (*AuthorizedOperations, bool)`

GetAuthorizedOperationsOk returns a tuple with the AuthorizedOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedOperations

`func (o *TopicData) SetAuthorizedOperations(v AuthorizedOperations)`

SetAuthorizedOperations sets AuthorizedOperations field to given value.

### HasAuthorizedOperations

`func (o *TopicData) HasAuthorizedOperations() bool`

HasAuthorizedOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


