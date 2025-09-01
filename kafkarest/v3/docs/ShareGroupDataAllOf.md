# ShareGroupDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**ShareGroupId** | **string** |  | 
**State** | **string** |  | 
**Coordinator** | [**Relationship**](Relationship.md) |  | 
**Consumers** | [**Relationship**](Relationship.md) |  | 
**ConsumerCount** | **int32** | Number of consumers in this share group | 
**PartitionCount** | **int32** | Total number of partitions assigned to this share group across all consumers | 
**AssignedTopicPartitions** | Pointer to [**[]ShareGroupTopicPartitionData**](ShareGroupTopicPartitionData.md) | List of topic-partitions assigned to this share group, including those from empty groups | [optional] 

## Methods

### NewShareGroupDataAllOf

`func NewShareGroupDataAllOf(clusterId string, shareGroupId string, state string, coordinator Relationship, consumers Relationship, consumerCount int32, partitionCount int32, ) *ShareGroupDataAllOf`

NewShareGroupDataAllOf instantiates a new ShareGroupDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareGroupDataAllOfWithDefaults

`func NewShareGroupDataAllOfWithDefaults() *ShareGroupDataAllOf`

NewShareGroupDataAllOfWithDefaults instantiates a new ShareGroupDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ShareGroupDataAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ShareGroupDataAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ShareGroupDataAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetShareGroupId

`func (o *ShareGroupDataAllOf) GetShareGroupId() string`

GetShareGroupId returns the ShareGroupId field if non-nil, zero value otherwise.

### GetShareGroupIdOk

`func (o *ShareGroupDataAllOf) GetShareGroupIdOk() (*string, bool)`

GetShareGroupIdOk returns a tuple with the ShareGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareGroupId

`func (o *ShareGroupDataAllOf) SetShareGroupId(v string)`

SetShareGroupId sets ShareGroupId field to given value.


### GetState

`func (o *ShareGroupDataAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ShareGroupDataAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ShareGroupDataAllOf) SetState(v string)`

SetState sets State field to given value.


### GetCoordinator

`func (o *ShareGroupDataAllOf) GetCoordinator() Relationship`

GetCoordinator returns the Coordinator field if non-nil, zero value otherwise.

### GetCoordinatorOk

`func (o *ShareGroupDataAllOf) GetCoordinatorOk() (*Relationship, bool)`

GetCoordinatorOk returns a tuple with the Coordinator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinator

`func (o *ShareGroupDataAllOf) SetCoordinator(v Relationship)`

SetCoordinator sets Coordinator field to given value.


### GetConsumers

`func (o *ShareGroupDataAllOf) GetConsumers() Relationship`

GetConsumers returns the Consumers field if non-nil, zero value otherwise.

### GetConsumersOk

`func (o *ShareGroupDataAllOf) GetConsumersOk() (*Relationship, bool)`

GetConsumersOk returns a tuple with the Consumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumers

`func (o *ShareGroupDataAllOf) SetConsumers(v Relationship)`

SetConsumers sets Consumers field to given value.


### GetConsumerCount

`func (o *ShareGroupDataAllOf) GetConsumerCount() int32`

GetConsumerCount returns the ConsumerCount field if non-nil, zero value otherwise.

### GetConsumerCountOk

`func (o *ShareGroupDataAllOf) GetConsumerCountOk() (*int32, bool)`

GetConsumerCountOk returns a tuple with the ConsumerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerCount

`func (o *ShareGroupDataAllOf) SetConsumerCount(v int32)`

SetConsumerCount sets ConsumerCount field to given value.


### GetPartitionCount

`func (o *ShareGroupDataAllOf) GetPartitionCount() int32`

GetPartitionCount returns the PartitionCount field if non-nil, zero value otherwise.

### GetPartitionCountOk

`func (o *ShareGroupDataAllOf) GetPartitionCountOk() (*int32, bool)`

GetPartitionCountOk returns a tuple with the PartitionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionCount

`func (o *ShareGroupDataAllOf) SetPartitionCount(v int32)`

SetPartitionCount sets PartitionCount field to given value.


### GetAssignedTopicPartitions

`func (o *ShareGroupDataAllOf) GetAssignedTopicPartitions() []ShareGroupTopicPartitionData`

GetAssignedTopicPartitions returns the AssignedTopicPartitions field if non-nil, zero value otherwise.

### GetAssignedTopicPartitionsOk

`func (o *ShareGroupDataAllOf) GetAssignedTopicPartitionsOk() (*[]ShareGroupTopicPartitionData, bool)`

GetAssignedTopicPartitionsOk returns a tuple with the AssignedTopicPartitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTopicPartitions

`func (o *ShareGroupDataAllOf) SetAssignedTopicPartitions(v []ShareGroupTopicPartitionData)`

SetAssignedTopicPartitions sets AssignedTopicPartitions field to given value.

### HasAssignedTopicPartitions

`func (o *ShareGroupDataAllOf) HasAssignedTopicPartitions() bool`

HasAssignedTopicPartitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


