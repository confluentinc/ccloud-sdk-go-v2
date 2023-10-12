# ConsumerAssignmentDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**ConsumerGroupId** | **string** |  | 
**ConsumerId** | **string** |  | 
**TopicName** | **string** |  | 
**PartitionId** | **int32** |  | 
**Partition** | [**Relationship**](Relationship.md) |  | 
**Lag** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewConsumerAssignmentDataAllOf

`func NewConsumerAssignmentDataAllOf(clusterId string, consumerGroupId string, consumerId string, topicName string, partitionId int32, partition Relationship, lag Relationship, ) *ConsumerAssignmentDataAllOf`

NewConsumerAssignmentDataAllOf instantiates a new ConsumerAssignmentDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerAssignmentDataAllOfWithDefaults

`func NewConsumerAssignmentDataAllOfWithDefaults() *ConsumerAssignmentDataAllOf`

NewConsumerAssignmentDataAllOfWithDefaults instantiates a new ConsumerAssignmentDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ConsumerAssignmentDataAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ConsumerAssignmentDataAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ConsumerAssignmentDataAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetConsumerGroupId

`func (o *ConsumerAssignmentDataAllOf) GetConsumerGroupId() string`

GetConsumerGroupId returns the ConsumerGroupId field if non-nil, zero value otherwise.

### GetConsumerGroupIdOk

`func (o *ConsumerAssignmentDataAllOf) GetConsumerGroupIdOk() (*string, bool)`

GetConsumerGroupIdOk returns a tuple with the ConsumerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerGroupId

`func (o *ConsumerAssignmentDataAllOf) SetConsumerGroupId(v string)`

SetConsumerGroupId sets ConsumerGroupId field to given value.


### GetConsumerId

`func (o *ConsumerAssignmentDataAllOf) GetConsumerId() string`

GetConsumerId returns the ConsumerId field if non-nil, zero value otherwise.

### GetConsumerIdOk

`func (o *ConsumerAssignmentDataAllOf) GetConsumerIdOk() (*string, bool)`

GetConsumerIdOk returns a tuple with the ConsumerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerId

`func (o *ConsumerAssignmentDataAllOf) SetConsumerId(v string)`

SetConsumerId sets ConsumerId field to given value.


### GetTopicName

`func (o *ConsumerAssignmentDataAllOf) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *ConsumerAssignmentDataAllOf) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *ConsumerAssignmentDataAllOf) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.


### GetPartitionId

`func (o *ConsumerAssignmentDataAllOf) GetPartitionId() int32`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *ConsumerAssignmentDataAllOf) GetPartitionIdOk() (*int32, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *ConsumerAssignmentDataAllOf) SetPartitionId(v int32)`

SetPartitionId sets PartitionId field to given value.


### GetPartition

`func (o *ConsumerAssignmentDataAllOf) GetPartition() Relationship`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *ConsumerAssignmentDataAllOf) GetPartitionOk() (*Relationship, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *ConsumerAssignmentDataAllOf) SetPartition(v Relationship)`

SetPartition sets Partition field to given value.


### GetLag

`func (o *ConsumerAssignmentDataAllOf) GetLag() Relationship`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *ConsumerAssignmentDataAllOf) GetLagOk() (*Relationship, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *ConsumerAssignmentDataAllOf) SetLag(v Relationship)`

SetLag sets Lag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


