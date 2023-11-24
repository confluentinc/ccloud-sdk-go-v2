# ConsumerLagData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**ClusterId** | **string** |  | 
**ConsumerGroupId** | **string** |  | 
**TopicName** | **string** |  | 
**PartitionId** | **int32** |  | 
**CurrentOffset** | **int64** |  | 
**LogEndOffset** | **int64** |  | 
**Lag** | **int64** |  | 
**ConsumerId** | **string** |  | 
**InstanceId** | Pointer to **NullableString** |  | [optional] 
**ClientId** | **string** |  | 

## Methods

### NewConsumerLagData

`func NewConsumerLagData(kind string, metadata ResourceMetadata, clusterId string, consumerGroupId string, topicName string, partitionId int32, currentOffset int64, logEndOffset int64, lag int64, consumerId string, clientId string, ) *ConsumerLagData`

NewConsumerLagData instantiates a new ConsumerLagData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerLagDataWithDefaults

`func NewConsumerLagDataWithDefaults() *ConsumerLagData`

NewConsumerLagDataWithDefaults instantiates a new ConsumerLagData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ConsumerLagData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConsumerLagData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConsumerLagData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ConsumerLagData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConsumerLagData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConsumerLagData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *ConsumerLagData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ConsumerLagData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ConsumerLagData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetConsumerGroupId

`func (o *ConsumerLagData) GetConsumerGroupId() string`

GetConsumerGroupId returns the ConsumerGroupId field if non-nil, zero value otherwise.

### GetConsumerGroupIdOk

`func (o *ConsumerLagData) GetConsumerGroupIdOk() (*string, bool)`

GetConsumerGroupIdOk returns a tuple with the ConsumerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerGroupId

`func (o *ConsumerLagData) SetConsumerGroupId(v string)`

SetConsumerGroupId sets ConsumerGroupId field to given value.


### GetTopicName

`func (o *ConsumerLagData) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *ConsumerLagData) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *ConsumerLagData) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.


### GetPartitionId

`func (o *ConsumerLagData) GetPartitionId() int32`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *ConsumerLagData) GetPartitionIdOk() (*int32, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *ConsumerLagData) SetPartitionId(v int32)`

SetPartitionId sets PartitionId field to given value.


### GetCurrentOffset

`func (o *ConsumerLagData) GetCurrentOffset() int64`

GetCurrentOffset returns the CurrentOffset field if non-nil, zero value otherwise.

### GetCurrentOffsetOk

`func (o *ConsumerLagData) GetCurrentOffsetOk() (*int64, bool)`

GetCurrentOffsetOk returns a tuple with the CurrentOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentOffset

`func (o *ConsumerLagData) SetCurrentOffset(v int64)`

SetCurrentOffset sets CurrentOffset field to given value.


### GetLogEndOffset

`func (o *ConsumerLagData) GetLogEndOffset() int64`

GetLogEndOffset returns the LogEndOffset field if non-nil, zero value otherwise.

### GetLogEndOffsetOk

`func (o *ConsumerLagData) GetLogEndOffsetOk() (*int64, bool)`

GetLogEndOffsetOk returns a tuple with the LogEndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogEndOffset

`func (o *ConsumerLagData) SetLogEndOffset(v int64)`

SetLogEndOffset sets LogEndOffset field to given value.


### GetLag

`func (o *ConsumerLagData) GetLag() int64`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *ConsumerLagData) GetLagOk() (*int64, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *ConsumerLagData) SetLag(v int64)`

SetLag sets Lag field to given value.


### GetConsumerId

`func (o *ConsumerLagData) GetConsumerId() string`

GetConsumerId returns the ConsumerId field if non-nil, zero value otherwise.

### GetConsumerIdOk

`func (o *ConsumerLagData) GetConsumerIdOk() (*string, bool)`

GetConsumerIdOk returns a tuple with the ConsumerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerId

`func (o *ConsumerLagData) SetConsumerId(v string)`

SetConsumerId sets ConsumerId field to given value.


### GetInstanceId

`func (o *ConsumerLagData) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ConsumerLagData) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ConsumerLagData) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ConsumerLagData) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *ConsumerLagData) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *ConsumerLagData) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetClientId

`func (o *ConsumerLagData) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ConsumerLagData) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ConsumerLagData) SetClientId(v string)`

SetClientId sets ClientId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


