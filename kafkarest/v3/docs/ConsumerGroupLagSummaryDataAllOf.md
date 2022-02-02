# ConsumerGroupLagSummaryDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**ConsumerGroupId** | **string** |  | 
**MaxLagConsumerId** | **string** |  | 
**MaxLagInstanceId** | Pointer to **NullableString** |  | [optional] 
**MaxLagClientId** | **string** |  | 
**MaxLagTopicName** | **string** |  | 
**MaxLagPartitionId** | **int32** |  | 
**MaxLag** | **int64** |  | 
**TotalLag** | **int64** |  | 
**MaxLagConsumer** | [**Relationship**](Relationship.md) |  | 
**MaxLagPartition** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewConsumerGroupLagSummaryDataAllOf

`func NewConsumerGroupLagSummaryDataAllOf(clusterId string, consumerGroupId string, maxLagConsumerId string, maxLagClientId string, maxLagTopicName string, maxLagPartitionId int32, maxLag int64, totalLag int64, maxLagConsumer Relationship, maxLagPartition Relationship, ) *ConsumerGroupLagSummaryDataAllOf`

NewConsumerGroupLagSummaryDataAllOf instantiates a new ConsumerGroupLagSummaryDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerGroupLagSummaryDataAllOfWithDefaults

`func NewConsumerGroupLagSummaryDataAllOfWithDefaults() *ConsumerGroupLagSummaryDataAllOf`

NewConsumerGroupLagSummaryDataAllOfWithDefaults instantiates a new ConsumerGroupLagSummaryDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ConsumerGroupLagSummaryDataAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ConsumerGroupLagSummaryDataAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ConsumerGroupLagSummaryDataAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetConsumerGroupId

`func (o *ConsumerGroupLagSummaryDataAllOf) GetConsumerGroupId() string`

GetConsumerGroupId returns the ConsumerGroupId field if non-nil, zero value otherwise.

### GetConsumerGroupIdOk

`func (o *ConsumerGroupLagSummaryDataAllOf) GetConsumerGroupIdOk() (*string, bool)`

GetConsumerGroupIdOk returns a tuple with the ConsumerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerGroupId

`func (o *ConsumerGroupLagSummaryDataAllOf) SetConsumerGroupId(v string)`

SetConsumerGroupId sets ConsumerGroupId field to given value.


### GetMaxLagConsumerId

`func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagConsumerId() string`

GetMaxLagConsumerId returns the MaxLagConsumerId field if non-nil, zero value otherwise.

### GetMaxLagConsumerIdOk

`func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagConsumerIdOk() (*string, bool)`

GetMaxLagConsumerIdOk returns a tuple with the MaxLagConsumerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLagConsumerId

`func (o *ConsumerGroupLagSummaryDataAllOf) SetMaxLagConsumerId(v string)`

SetMaxLagConsumerId sets MaxLagConsumerId field to given value.


### GetMaxLagInstanceId

`func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagInstanceId() string`

GetMaxLagInstanceId returns the MaxLagInstanceId field if non-nil, zero value otherwise.

### GetMaxLagInstanceIdOk

`func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagInstanceIdOk() (*string, bool)`

GetMaxLagInstanceIdOk returns a tuple with the MaxLagInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLagInstanceId

`func (o *ConsumerGroupLagSummaryDataAllOf) SetMaxLagInstanceId(v string)`

SetMaxLagInstanceId sets MaxLagInstanceId field to given value.

### HasMaxLagInstanceId

`func (o *ConsumerGroupLagSummaryDataAllOf) HasMaxLagInstanceId() bool`

HasMaxLagInstanceId returns a boolean if a field has been set.

### SetMaxLagInstanceIdNil

`func (o *ConsumerGroupLagSummaryDataAllOf) SetMaxLagInstanceIdNil(b bool)`

 SetMaxLagInstanceIdNil sets the value for MaxLagInstanceId to be an explicit nil

### UnsetMaxLagInstanceId
`func (o *ConsumerGroupLagSummaryDataAllOf) UnsetMaxLagInstanceId()`

UnsetMaxLagInstanceId ensures that no value is present for MaxLagInstanceId, not even an explicit nil
### GetMaxLagClientId

`func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagClientId() string`

GetMaxLagClientId returns the MaxLagClientId field if non-nil, zero value otherwise.

### GetMaxLagClientIdOk

`func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagClientIdOk() (*string, bool)`

GetMaxLagClientIdOk returns a tuple with the MaxLagClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLagClientId

`func (o *ConsumerGroupLagSummaryDataAllOf) SetMaxLagClientId(v string)`

SetMaxLagClientId sets MaxLagClientId field to given value.


### GetMaxLagTopicName

`func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagTopicName() string`

GetMaxLagTopicName returns the MaxLagTopicName field if non-nil, zero value otherwise.

### GetMaxLagTopicNameOk

`func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagTopicNameOk() (*string, bool)`

GetMaxLagTopicNameOk returns a tuple with the MaxLagTopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLagTopicName

`func (o *ConsumerGroupLagSummaryDataAllOf) SetMaxLagTopicName(v string)`

SetMaxLagTopicName sets MaxLagTopicName field to given value.


### GetMaxLagPartitionId

`func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagPartitionId() int32`

GetMaxLagPartitionId returns the MaxLagPartitionId field if non-nil, zero value otherwise.

### GetMaxLagPartitionIdOk

`func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagPartitionIdOk() (*int32, bool)`

GetMaxLagPartitionIdOk returns a tuple with the MaxLagPartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLagPartitionId

`func (o *ConsumerGroupLagSummaryDataAllOf) SetMaxLagPartitionId(v int32)`

SetMaxLagPartitionId sets MaxLagPartitionId field to given value.


### GetMaxLag

`func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLag() int64`

GetMaxLag returns the MaxLag field if non-nil, zero value otherwise.

### GetMaxLagOk

`func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagOk() (*int64, bool)`

GetMaxLagOk returns a tuple with the MaxLag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLag

`func (o *ConsumerGroupLagSummaryDataAllOf) SetMaxLag(v int64)`

SetMaxLag sets MaxLag field to given value.


### GetTotalLag

`func (o *ConsumerGroupLagSummaryDataAllOf) GetTotalLag() int64`

GetTotalLag returns the TotalLag field if non-nil, zero value otherwise.

### GetTotalLagOk

`func (o *ConsumerGroupLagSummaryDataAllOf) GetTotalLagOk() (*int64, bool)`

GetTotalLagOk returns a tuple with the TotalLag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLag

`func (o *ConsumerGroupLagSummaryDataAllOf) SetTotalLag(v int64)`

SetTotalLag sets TotalLag field to given value.


### GetMaxLagConsumer

`func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagConsumer() Relationship`

GetMaxLagConsumer returns the MaxLagConsumer field if non-nil, zero value otherwise.

### GetMaxLagConsumerOk

`func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagConsumerOk() (*Relationship, bool)`

GetMaxLagConsumerOk returns a tuple with the MaxLagConsumer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLagConsumer

`func (o *ConsumerGroupLagSummaryDataAllOf) SetMaxLagConsumer(v Relationship)`

SetMaxLagConsumer sets MaxLagConsumer field to given value.


### GetMaxLagPartition

`func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagPartition() Relationship`

GetMaxLagPartition returns the MaxLagPartition field if non-nil, zero value otherwise.

### GetMaxLagPartitionOk

`func (o *ConsumerGroupLagSummaryDataAllOf) GetMaxLagPartitionOk() (*Relationship, bool)`

GetMaxLagPartitionOk returns a tuple with the MaxLagPartition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLagPartition

`func (o *ConsumerGroupLagSummaryDataAllOf) SetMaxLagPartition(v Relationship)`

SetMaxLagPartition sets MaxLagPartition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


