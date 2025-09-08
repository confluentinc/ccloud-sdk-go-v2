# ShareGroupTopicPartitionDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopicName** | **string** | The name of the topic | 
**PartitionId** | **int32** | The partition ID | 
**Partition** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewShareGroupTopicPartitionDataAllOf

`func NewShareGroupTopicPartitionDataAllOf(topicName string, partitionId int32, partition Relationship, ) *ShareGroupTopicPartitionDataAllOf`

NewShareGroupTopicPartitionDataAllOf instantiates a new ShareGroupTopicPartitionDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareGroupTopicPartitionDataAllOfWithDefaults

`func NewShareGroupTopicPartitionDataAllOfWithDefaults() *ShareGroupTopicPartitionDataAllOf`

NewShareGroupTopicPartitionDataAllOfWithDefaults instantiates a new ShareGroupTopicPartitionDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopicName

`func (o *ShareGroupTopicPartitionDataAllOf) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *ShareGroupTopicPartitionDataAllOf) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *ShareGroupTopicPartitionDataAllOf) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.


### GetPartitionId

`func (o *ShareGroupTopicPartitionDataAllOf) GetPartitionId() int32`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *ShareGroupTopicPartitionDataAllOf) GetPartitionIdOk() (*int32, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *ShareGroupTopicPartitionDataAllOf) SetPartitionId(v int32)`

SetPartitionId sets PartitionId field to given value.


### GetPartition

`func (o *ShareGroupTopicPartitionDataAllOf) GetPartition() Relationship`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *ShareGroupTopicPartitionDataAllOf) GetPartitionOk() (*Relationship, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *ShareGroupTopicPartitionDataAllOf) SetPartition(v Relationship)`

SetPartition sets Partition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


