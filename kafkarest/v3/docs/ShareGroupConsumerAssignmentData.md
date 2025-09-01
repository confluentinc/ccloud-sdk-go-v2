# ShareGroupConsumerAssignmentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**ClusterId** | **string** |  | 
**GroupId** | **string** |  | 
**ConsumerId** | **string** |  | 
**TopicName** | **string** |  | 
**PartitionId** | **int32** |  | 
**Partition** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewShareGroupConsumerAssignmentData

`func NewShareGroupConsumerAssignmentData(kind string, metadata ResourceMetadata, clusterId string, groupId string, consumerId string, topicName string, partitionId int32, partition Relationship, ) *ShareGroupConsumerAssignmentData`

NewShareGroupConsumerAssignmentData instantiates a new ShareGroupConsumerAssignmentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareGroupConsumerAssignmentDataWithDefaults

`func NewShareGroupConsumerAssignmentDataWithDefaults() *ShareGroupConsumerAssignmentData`

NewShareGroupConsumerAssignmentDataWithDefaults instantiates a new ShareGroupConsumerAssignmentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ShareGroupConsumerAssignmentData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ShareGroupConsumerAssignmentData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ShareGroupConsumerAssignmentData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ShareGroupConsumerAssignmentData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ShareGroupConsumerAssignmentData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ShareGroupConsumerAssignmentData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *ShareGroupConsumerAssignmentData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ShareGroupConsumerAssignmentData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ShareGroupConsumerAssignmentData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetGroupId

`func (o *ShareGroupConsumerAssignmentData) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ShareGroupConsumerAssignmentData) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ShareGroupConsumerAssignmentData) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetConsumerId

`func (o *ShareGroupConsumerAssignmentData) GetConsumerId() string`

GetConsumerId returns the ConsumerId field if non-nil, zero value otherwise.

### GetConsumerIdOk

`func (o *ShareGroupConsumerAssignmentData) GetConsumerIdOk() (*string, bool)`

GetConsumerIdOk returns a tuple with the ConsumerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerId

`func (o *ShareGroupConsumerAssignmentData) SetConsumerId(v string)`

SetConsumerId sets ConsumerId field to given value.


### GetTopicName

`func (o *ShareGroupConsumerAssignmentData) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *ShareGroupConsumerAssignmentData) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *ShareGroupConsumerAssignmentData) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.


### GetPartitionId

`func (o *ShareGroupConsumerAssignmentData) GetPartitionId() int32`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *ShareGroupConsumerAssignmentData) GetPartitionIdOk() (*int32, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *ShareGroupConsumerAssignmentData) SetPartitionId(v int32)`

SetPartitionId sets PartitionId field to given value.


### GetPartition

`func (o *ShareGroupConsumerAssignmentData) GetPartition() Relationship`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *ShareGroupConsumerAssignmentData) GetPartitionOk() (*Relationship, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *ShareGroupConsumerAssignmentData) SetPartition(v Relationship)`

SetPartition sets Partition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


