# ConsumerGroupLagSummaryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
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

### NewConsumerGroupLagSummaryData

`func NewConsumerGroupLagSummaryData(kind string, metadata ResourceMetadata, clusterId string, consumerGroupId string, maxLagConsumerId string, maxLagClientId string, maxLagTopicName string, maxLagPartitionId int32, maxLag int64, totalLag int64, maxLagConsumer Relationship, maxLagPartition Relationship, ) *ConsumerGroupLagSummaryData`

NewConsumerGroupLagSummaryData instantiates a new ConsumerGroupLagSummaryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerGroupLagSummaryDataWithDefaults

`func NewConsumerGroupLagSummaryDataWithDefaults() *ConsumerGroupLagSummaryData`

NewConsumerGroupLagSummaryDataWithDefaults instantiates a new ConsumerGroupLagSummaryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ConsumerGroupLagSummaryData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConsumerGroupLagSummaryData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConsumerGroupLagSummaryData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ConsumerGroupLagSummaryData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConsumerGroupLagSummaryData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConsumerGroupLagSummaryData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *ConsumerGroupLagSummaryData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ConsumerGroupLagSummaryData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ConsumerGroupLagSummaryData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetConsumerGroupId

`func (o *ConsumerGroupLagSummaryData) GetConsumerGroupId() string`

GetConsumerGroupId returns the ConsumerGroupId field if non-nil, zero value otherwise.

### GetConsumerGroupIdOk

`func (o *ConsumerGroupLagSummaryData) GetConsumerGroupIdOk() (*string, bool)`

GetConsumerGroupIdOk returns a tuple with the ConsumerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerGroupId

`func (o *ConsumerGroupLagSummaryData) SetConsumerGroupId(v string)`

SetConsumerGroupId sets ConsumerGroupId field to given value.


### GetMaxLagConsumerId

`func (o *ConsumerGroupLagSummaryData) GetMaxLagConsumerId() string`

GetMaxLagConsumerId returns the MaxLagConsumerId field if non-nil, zero value otherwise.

### GetMaxLagConsumerIdOk

`func (o *ConsumerGroupLagSummaryData) GetMaxLagConsumerIdOk() (*string, bool)`

GetMaxLagConsumerIdOk returns a tuple with the MaxLagConsumerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLagConsumerId

`func (o *ConsumerGroupLagSummaryData) SetMaxLagConsumerId(v string)`

SetMaxLagConsumerId sets MaxLagConsumerId field to given value.


### GetMaxLagInstanceId

`func (o *ConsumerGroupLagSummaryData) GetMaxLagInstanceId() string`

GetMaxLagInstanceId returns the MaxLagInstanceId field if non-nil, zero value otherwise.

### GetMaxLagInstanceIdOk

`func (o *ConsumerGroupLagSummaryData) GetMaxLagInstanceIdOk() (*string, bool)`

GetMaxLagInstanceIdOk returns a tuple with the MaxLagInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLagInstanceId

`func (o *ConsumerGroupLagSummaryData) SetMaxLagInstanceId(v string)`

SetMaxLagInstanceId sets MaxLagInstanceId field to given value.

### HasMaxLagInstanceId

`func (o *ConsumerGroupLagSummaryData) HasMaxLagInstanceId() bool`

HasMaxLagInstanceId returns a boolean if a field has been set.

### SetMaxLagInstanceIdNil

`func (o *ConsumerGroupLagSummaryData) SetMaxLagInstanceIdNil(b bool)`

 SetMaxLagInstanceIdNil sets the value for MaxLagInstanceId to be an explicit nil

### UnsetMaxLagInstanceId
`func (o *ConsumerGroupLagSummaryData) UnsetMaxLagInstanceId()`

UnsetMaxLagInstanceId ensures that no value is present for MaxLagInstanceId, not even an explicit nil
### GetMaxLagClientId

`func (o *ConsumerGroupLagSummaryData) GetMaxLagClientId() string`

GetMaxLagClientId returns the MaxLagClientId field if non-nil, zero value otherwise.

### GetMaxLagClientIdOk

`func (o *ConsumerGroupLagSummaryData) GetMaxLagClientIdOk() (*string, bool)`

GetMaxLagClientIdOk returns a tuple with the MaxLagClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLagClientId

`func (o *ConsumerGroupLagSummaryData) SetMaxLagClientId(v string)`

SetMaxLagClientId sets MaxLagClientId field to given value.


### GetMaxLagTopicName

`func (o *ConsumerGroupLagSummaryData) GetMaxLagTopicName() string`

GetMaxLagTopicName returns the MaxLagTopicName field if non-nil, zero value otherwise.

### GetMaxLagTopicNameOk

`func (o *ConsumerGroupLagSummaryData) GetMaxLagTopicNameOk() (*string, bool)`

GetMaxLagTopicNameOk returns a tuple with the MaxLagTopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLagTopicName

`func (o *ConsumerGroupLagSummaryData) SetMaxLagTopicName(v string)`

SetMaxLagTopicName sets MaxLagTopicName field to given value.


### GetMaxLagPartitionId

`func (o *ConsumerGroupLagSummaryData) GetMaxLagPartitionId() int32`

GetMaxLagPartitionId returns the MaxLagPartitionId field if non-nil, zero value otherwise.

### GetMaxLagPartitionIdOk

`func (o *ConsumerGroupLagSummaryData) GetMaxLagPartitionIdOk() (*int32, bool)`

GetMaxLagPartitionIdOk returns a tuple with the MaxLagPartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLagPartitionId

`func (o *ConsumerGroupLagSummaryData) SetMaxLagPartitionId(v int32)`

SetMaxLagPartitionId sets MaxLagPartitionId field to given value.


### GetMaxLag

`func (o *ConsumerGroupLagSummaryData) GetMaxLag() int64`

GetMaxLag returns the MaxLag field if non-nil, zero value otherwise.

### GetMaxLagOk

`func (o *ConsumerGroupLagSummaryData) GetMaxLagOk() (*int64, bool)`

GetMaxLagOk returns a tuple with the MaxLag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLag

`func (o *ConsumerGroupLagSummaryData) SetMaxLag(v int64)`

SetMaxLag sets MaxLag field to given value.


### GetTotalLag

`func (o *ConsumerGroupLagSummaryData) GetTotalLag() int64`

GetTotalLag returns the TotalLag field if non-nil, zero value otherwise.

### GetTotalLagOk

`func (o *ConsumerGroupLagSummaryData) GetTotalLagOk() (*int64, bool)`

GetTotalLagOk returns a tuple with the TotalLag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLag

`func (o *ConsumerGroupLagSummaryData) SetTotalLag(v int64)`

SetTotalLag sets TotalLag field to given value.


### GetMaxLagConsumer

`func (o *ConsumerGroupLagSummaryData) GetMaxLagConsumer() Relationship`

GetMaxLagConsumer returns the MaxLagConsumer field if non-nil, zero value otherwise.

### GetMaxLagConsumerOk

`func (o *ConsumerGroupLagSummaryData) GetMaxLagConsumerOk() (*Relationship, bool)`

GetMaxLagConsumerOk returns a tuple with the MaxLagConsumer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLagConsumer

`func (o *ConsumerGroupLagSummaryData) SetMaxLagConsumer(v Relationship)`

SetMaxLagConsumer sets MaxLagConsumer field to given value.


### GetMaxLagPartition

`func (o *ConsumerGroupLagSummaryData) GetMaxLagPartition() Relationship`

GetMaxLagPartition returns the MaxLagPartition field if non-nil, zero value otherwise.

### GetMaxLagPartitionOk

`func (o *ConsumerGroupLagSummaryData) GetMaxLagPartitionOk() (*Relationship, bool)`

GetMaxLagPartitionOk returns a tuple with the MaxLagPartition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLagPartition

`func (o *ConsumerGroupLagSummaryData) SetMaxLagPartition(v Relationship)`

SetMaxLagPartition sets MaxLagPartition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


