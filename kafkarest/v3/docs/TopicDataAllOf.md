# TopicDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**TopicName** | **string** |  | 
**IsInternal** | **bool** |  | 
**ReplicationFactor** | **int32** |  | 
**PartitionsCount** | **int32** |  | 
**Partitions** | [**Relationship**](Relationship.md) |  | 
**Configs** | [**Relationship**](Relationship.md) |  | 
**PartitionReassignments** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewTopicDataAllOf

`func NewTopicDataAllOf(clusterId string, topicName string, isInternal bool, replicationFactor int32, partitionsCount int32, partitions Relationship, configs Relationship, partitionReassignments Relationship, ) *TopicDataAllOf`

NewTopicDataAllOf instantiates a new TopicDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicDataAllOfWithDefaults

`func NewTopicDataAllOfWithDefaults() *TopicDataAllOf`

NewTopicDataAllOfWithDefaults instantiates a new TopicDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *TopicDataAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *TopicDataAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *TopicDataAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetTopicName

`func (o *TopicDataAllOf) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *TopicDataAllOf) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *TopicDataAllOf) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.


### GetIsInternal

`func (o *TopicDataAllOf) GetIsInternal() bool`

GetIsInternal returns the IsInternal field if non-nil, zero value otherwise.

### GetIsInternalOk

`func (o *TopicDataAllOf) GetIsInternalOk() (*bool, bool)`

GetIsInternalOk returns a tuple with the IsInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternal

`func (o *TopicDataAllOf) SetIsInternal(v bool)`

SetIsInternal sets IsInternal field to given value.


### GetReplicationFactor

`func (o *TopicDataAllOf) GetReplicationFactor() int32`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *TopicDataAllOf) GetReplicationFactorOk() (*int32, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *TopicDataAllOf) SetReplicationFactor(v int32)`

SetReplicationFactor sets ReplicationFactor field to given value.


### GetPartitionsCount

`func (o *TopicDataAllOf) GetPartitionsCount() int32`

GetPartitionsCount returns the PartitionsCount field if non-nil, zero value otherwise.

### GetPartitionsCountOk

`func (o *TopicDataAllOf) GetPartitionsCountOk() (*int32, bool)`

GetPartitionsCountOk returns a tuple with the PartitionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionsCount

`func (o *TopicDataAllOf) SetPartitionsCount(v int32)`

SetPartitionsCount sets PartitionsCount field to given value.


### GetPartitions

`func (o *TopicDataAllOf) GetPartitions() Relationship`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *TopicDataAllOf) GetPartitionsOk() (*Relationship, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *TopicDataAllOf) SetPartitions(v Relationship)`

SetPartitions sets Partitions field to given value.


### GetConfigs

`func (o *TopicDataAllOf) GetConfigs() Relationship`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *TopicDataAllOf) GetConfigsOk() (*Relationship, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *TopicDataAllOf) SetConfigs(v Relationship)`

SetConfigs sets Configs field to given value.


### GetPartitionReassignments

`func (o *TopicDataAllOf) GetPartitionReassignments() Relationship`

GetPartitionReassignments returns the PartitionReassignments field if non-nil, zero value otherwise.

### GetPartitionReassignmentsOk

`func (o *TopicDataAllOf) GetPartitionReassignmentsOk() (*Relationship, bool)`

GetPartitionReassignmentsOk returns a tuple with the PartitionReassignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionReassignments

`func (o *TopicDataAllOf) SetPartitionReassignments(v Relationship)`

SetPartitionReassignments sets PartitionReassignments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


