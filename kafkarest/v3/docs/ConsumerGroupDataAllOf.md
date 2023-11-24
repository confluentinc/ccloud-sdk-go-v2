# ConsumerGroupDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**ConsumerGroupId** | **string** |  | 
**IsSimple** | **bool** |  | 
**PartitionAssignor** | **string** |  | 
**State** | **string** |  | 
**Coordinator** | [**Relationship**](Relationship.md) |  | 
**Consumer** | Pointer to [**Relationship**](Relationship.md) |  | [optional] 
**LagSummary** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewConsumerGroupDataAllOf

`func NewConsumerGroupDataAllOf(clusterId string, consumerGroupId string, isSimple bool, partitionAssignor string, state string, coordinator Relationship, lagSummary Relationship, ) *ConsumerGroupDataAllOf`

NewConsumerGroupDataAllOf instantiates a new ConsumerGroupDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerGroupDataAllOfWithDefaults

`func NewConsumerGroupDataAllOfWithDefaults() *ConsumerGroupDataAllOf`

NewConsumerGroupDataAllOfWithDefaults instantiates a new ConsumerGroupDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ConsumerGroupDataAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ConsumerGroupDataAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ConsumerGroupDataAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetConsumerGroupId

`func (o *ConsumerGroupDataAllOf) GetConsumerGroupId() string`

GetConsumerGroupId returns the ConsumerGroupId field if non-nil, zero value otherwise.

### GetConsumerGroupIdOk

`func (o *ConsumerGroupDataAllOf) GetConsumerGroupIdOk() (*string, bool)`

GetConsumerGroupIdOk returns a tuple with the ConsumerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerGroupId

`func (o *ConsumerGroupDataAllOf) SetConsumerGroupId(v string)`

SetConsumerGroupId sets ConsumerGroupId field to given value.


### GetIsSimple

`func (o *ConsumerGroupDataAllOf) GetIsSimple() bool`

GetIsSimple returns the IsSimple field if non-nil, zero value otherwise.

### GetIsSimpleOk

`func (o *ConsumerGroupDataAllOf) GetIsSimpleOk() (*bool, bool)`

GetIsSimpleOk returns a tuple with the IsSimple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSimple

`func (o *ConsumerGroupDataAllOf) SetIsSimple(v bool)`

SetIsSimple sets IsSimple field to given value.


### GetPartitionAssignor

`func (o *ConsumerGroupDataAllOf) GetPartitionAssignor() string`

GetPartitionAssignor returns the PartitionAssignor field if non-nil, zero value otherwise.

### GetPartitionAssignorOk

`func (o *ConsumerGroupDataAllOf) GetPartitionAssignorOk() (*string, bool)`

GetPartitionAssignorOk returns a tuple with the PartitionAssignor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionAssignor

`func (o *ConsumerGroupDataAllOf) SetPartitionAssignor(v string)`

SetPartitionAssignor sets PartitionAssignor field to given value.


### GetState

`func (o *ConsumerGroupDataAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConsumerGroupDataAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConsumerGroupDataAllOf) SetState(v string)`

SetState sets State field to given value.


### GetCoordinator

`func (o *ConsumerGroupDataAllOf) GetCoordinator() Relationship`

GetCoordinator returns the Coordinator field if non-nil, zero value otherwise.

### GetCoordinatorOk

`func (o *ConsumerGroupDataAllOf) GetCoordinatorOk() (*Relationship, bool)`

GetCoordinatorOk returns a tuple with the Coordinator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinator

`func (o *ConsumerGroupDataAllOf) SetCoordinator(v Relationship)`

SetCoordinator sets Coordinator field to given value.


### GetConsumer

`func (o *ConsumerGroupDataAllOf) GetConsumer() Relationship`

GetConsumer returns the Consumer field if non-nil, zero value otherwise.

### GetConsumerOk

`func (o *ConsumerGroupDataAllOf) GetConsumerOk() (*Relationship, bool)`

GetConsumerOk returns a tuple with the Consumer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumer

`func (o *ConsumerGroupDataAllOf) SetConsumer(v Relationship)`

SetConsumer sets Consumer field to given value.

### HasConsumer

`func (o *ConsumerGroupDataAllOf) HasConsumer() bool`

HasConsumer returns a boolean if a field has been set.

### GetLagSummary

`func (o *ConsumerGroupDataAllOf) GetLagSummary() Relationship`

GetLagSummary returns the LagSummary field if non-nil, zero value otherwise.

### GetLagSummaryOk

`func (o *ConsumerGroupDataAllOf) GetLagSummaryOk() (*Relationship, bool)`

GetLagSummaryOk returns a tuple with the LagSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagSummary

`func (o *ConsumerGroupDataAllOf) SetLagSummary(v Relationship)`

SetLagSummary sets LagSummary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


