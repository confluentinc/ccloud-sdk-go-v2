# ConsumerLagDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewConsumerLagDataAllOf

`func NewConsumerLagDataAllOf(clusterId string, consumerGroupId string, topicName string, partitionId int32, currentOffset int64, logEndOffset int64, lag int64, consumerId string, clientId string, ) *ConsumerLagDataAllOf`

NewConsumerLagDataAllOf instantiates a new ConsumerLagDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerLagDataAllOfWithDefaults

`func NewConsumerLagDataAllOfWithDefaults() *ConsumerLagDataAllOf`

NewConsumerLagDataAllOfWithDefaults instantiates a new ConsumerLagDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ConsumerLagDataAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ConsumerLagDataAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ConsumerLagDataAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetConsumerGroupId

`func (o *ConsumerLagDataAllOf) GetConsumerGroupId() string`

GetConsumerGroupId returns the ConsumerGroupId field if non-nil, zero value otherwise.

### GetConsumerGroupIdOk

`func (o *ConsumerLagDataAllOf) GetConsumerGroupIdOk() (*string, bool)`

GetConsumerGroupIdOk returns a tuple with the ConsumerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerGroupId

`func (o *ConsumerLagDataAllOf) SetConsumerGroupId(v string)`

SetConsumerGroupId sets ConsumerGroupId field to given value.


### GetTopicName

`func (o *ConsumerLagDataAllOf) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *ConsumerLagDataAllOf) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *ConsumerLagDataAllOf) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.


### GetPartitionId

`func (o *ConsumerLagDataAllOf) GetPartitionId() int32`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *ConsumerLagDataAllOf) GetPartitionIdOk() (*int32, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *ConsumerLagDataAllOf) SetPartitionId(v int32)`

SetPartitionId sets PartitionId field to given value.


### GetCurrentOffset

`func (o *ConsumerLagDataAllOf) GetCurrentOffset() int64`

GetCurrentOffset returns the CurrentOffset field if non-nil, zero value otherwise.

### GetCurrentOffsetOk

`func (o *ConsumerLagDataAllOf) GetCurrentOffsetOk() (*int64, bool)`

GetCurrentOffsetOk returns a tuple with the CurrentOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentOffset

`func (o *ConsumerLagDataAllOf) SetCurrentOffset(v int64)`

SetCurrentOffset sets CurrentOffset field to given value.


### GetLogEndOffset

`func (o *ConsumerLagDataAllOf) GetLogEndOffset() int64`

GetLogEndOffset returns the LogEndOffset field if non-nil, zero value otherwise.

### GetLogEndOffsetOk

`func (o *ConsumerLagDataAllOf) GetLogEndOffsetOk() (*int64, bool)`

GetLogEndOffsetOk returns a tuple with the LogEndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogEndOffset

`func (o *ConsumerLagDataAllOf) SetLogEndOffset(v int64)`

SetLogEndOffset sets LogEndOffset field to given value.


### GetLag

`func (o *ConsumerLagDataAllOf) GetLag() int64`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *ConsumerLagDataAllOf) GetLagOk() (*int64, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *ConsumerLagDataAllOf) SetLag(v int64)`

SetLag sets Lag field to given value.


### GetConsumerId

`func (o *ConsumerLagDataAllOf) GetConsumerId() string`

GetConsumerId returns the ConsumerId field if non-nil, zero value otherwise.

### GetConsumerIdOk

`func (o *ConsumerLagDataAllOf) GetConsumerIdOk() (*string, bool)`

GetConsumerIdOk returns a tuple with the ConsumerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerId

`func (o *ConsumerLagDataAllOf) SetConsumerId(v string)`

SetConsumerId sets ConsumerId field to given value.


### GetInstanceId

`func (o *ConsumerLagDataAllOf) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ConsumerLagDataAllOf) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ConsumerLagDataAllOf) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ConsumerLagDataAllOf) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *ConsumerLagDataAllOf) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *ConsumerLagDataAllOf) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetClientId

`func (o *ConsumerLagDataAllOf) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ConsumerLagDataAllOf) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ConsumerLagDataAllOf) SetClientId(v string)`

SetClientId sets ClientId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


