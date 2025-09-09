# ShareGroupTopicPartitionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**TopicName** | **string** | The name of the topic | 
**PartitionId** | **int32** | The partition ID | 
**Partition** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewShareGroupTopicPartitionData

`func NewShareGroupTopicPartitionData(kind string, metadata ResourceMetadata, topicName string, partitionId int32, partition Relationship, ) *ShareGroupTopicPartitionData`

NewShareGroupTopicPartitionData instantiates a new ShareGroupTopicPartitionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareGroupTopicPartitionDataWithDefaults

`func NewShareGroupTopicPartitionDataWithDefaults() *ShareGroupTopicPartitionData`

NewShareGroupTopicPartitionDataWithDefaults instantiates a new ShareGroupTopicPartitionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ShareGroupTopicPartitionData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ShareGroupTopicPartitionData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ShareGroupTopicPartitionData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ShareGroupTopicPartitionData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ShareGroupTopicPartitionData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ShareGroupTopicPartitionData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetTopicName

`func (o *ShareGroupTopicPartitionData) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *ShareGroupTopicPartitionData) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *ShareGroupTopicPartitionData) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.


### GetPartitionId

`func (o *ShareGroupTopicPartitionData) GetPartitionId() int32`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *ShareGroupTopicPartitionData) GetPartitionIdOk() (*int32, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *ShareGroupTopicPartitionData) SetPartitionId(v int32)`

SetPartitionId sets PartitionId field to given value.


### GetPartition

`func (o *ShareGroupTopicPartitionData) GetPartition() Relationship`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *ShareGroupTopicPartitionData) GetPartitionOk() (*Relationship, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *ShareGroupTopicPartitionData) SetPartition(v Relationship)`

SetPartition sets Partition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


