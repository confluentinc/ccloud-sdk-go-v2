# ConsumerGroupData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**ClusterId** | **string** |  | 
**ConsumerGroupId** | **string** |  | 
**IsSimple** | **bool** |  | 
**PartitionAssignor** | **string** |  | 
**State** | **string** |  | 
**Type** | **string** |  | 
**IsMixedConsumerGroup** | **bool** |  | 
**Coordinator** | [**Relationship**](Relationship.md) |  | 
**Consumers** | [**Relationship**](Relationship.md) |  | 
**LagSummary** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewConsumerGroupData

`func NewConsumerGroupData(kind string, metadata ResourceMetadata, clusterId string, consumerGroupId string, isSimple bool, partitionAssignor string, state string, type_ string, isMixedConsumerGroup bool, coordinator Relationship, consumers Relationship, lagSummary Relationship, ) *ConsumerGroupData`

NewConsumerGroupData instantiates a new ConsumerGroupData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerGroupDataWithDefaults

`func NewConsumerGroupDataWithDefaults() *ConsumerGroupData`

NewConsumerGroupDataWithDefaults instantiates a new ConsumerGroupData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ConsumerGroupData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConsumerGroupData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConsumerGroupData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ConsumerGroupData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConsumerGroupData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConsumerGroupData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *ConsumerGroupData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ConsumerGroupData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ConsumerGroupData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetConsumerGroupId

`func (o *ConsumerGroupData) GetConsumerGroupId() string`

GetConsumerGroupId returns the ConsumerGroupId field if non-nil, zero value otherwise.

### GetConsumerGroupIdOk

`func (o *ConsumerGroupData) GetConsumerGroupIdOk() (*string, bool)`

GetConsumerGroupIdOk returns a tuple with the ConsumerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerGroupId

`func (o *ConsumerGroupData) SetConsumerGroupId(v string)`

SetConsumerGroupId sets ConsumerGroupId field to given value.


### GetIsSimple

`func (o *ConsumerGroupData) GetIsSimple() bool`

GetIsSimple returns the IsSimple field if non-nil, zero value otherwise.

### GetIsSimpleOk

`func (o *ConsumerGroupData) GetIsSimpleOk() (*bool, bool)`

GetIsSimpleOk returns a tuple with the IsSimple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSimple

`func (o *ConsumerGroupData) SetIsSimple(v bool)`

SetIsSimple sets IsSimple field to given value.


### GetPartitionAssignor

`func (o *ConsumerGroupData) GetPartitionAssignor() string`

GetPartitionAssignor returns the PartitionAssignor field if non-nil, zero value otherwise.

### GetPartitionAssignorOk

`func (o *ConsumerGroupData) GetPartitionAssignorOk() (*string, bool)`

GetPartitionAssignorOk returns a tuple with the PartitionAssignor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionAssignor

`func (o *ConsumerGroupData) SetPartitionAssignor(v string)`

SetPartitionAssignor sets PartitionAssignor field to given value.


### GetState

`func (o *ConsumerGroupData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConsumerGroupData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConsumerGroupData) SetState(v string)`

SetState sets State field to given value.


### GetType

`func (o *ConsumerGroupData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConsumerGroupData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConsumerGroupData) SetType(v string)`

SetType sets Type field to given value.


### GetIsMixedConsumerGroup

`func (o *ConsumerGroupData) GetIsMixedConsumerGroup() bool`

GetIsMixedConsumerGroup returns the IsMixedConsumerGroup field if non-nil, zero value otherwise.

### GetIsMixedConsumerGroupOk

`func (o *ConsumerGroupData) GetIsMixedConsumerGroupOk() (*bool, bool)`

GetIsMixedConsumerGroupOk returns a tuple with the IsMixedConsumerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMixedConsumerGroup

`func (o *ConsumerGroupData) SetIsMixedConsumerGroup(v bool)`

SetIsMixedConsumerGroup sets IsMixedConsumerGroup field to given value.


### GetCoordinator

`func (o *ConsumerGroupData) GetCoordinator() Relationship`

GetCoordinator returns the Coordinator field if non-nil, zero value otherwise.

### GetCoordinatorOk

`func (o *ConsumerGroupData) GetCoordinatorOk() (*Relationship, bool)`

GetCoordinatorOk returns a tuple with the Coordinator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinator

`func (o *ConsumerGroupData) SetCoordinator(v Relationship)`

SetCoordinator sets Coordinator field to given value.


### GetConsumers

`func (o *ConsumerGroupData) GetConsumers() Relationship`

GetConsumers returns the Consumers field if non-nil, zero value otherwise.

### GetConsumersOk

`func (o *ConsumerGroupData) GetConsumersOk() (*Relationship, bool)`

GetConsumersOk returns a tuple with the Consumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumers

`func (o *ConsumerGroupData) SetConsumers(v Relationship)`

SetConsumers sets Consumers field to given value.


### GetLagSummary

`func (o *ConsumerGroupData) GetLagSummary() Relationship`

GetLagSummary returns the LagSummary field if non-nil, zero value otherwise.

### GetLagSummaryOk

`func (o *ConsumerGroupData) GetLagSummaryOk() (*Relationship, bool)`

GetLagSummaryOk returns a tuple with the LagSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagSummary

`func (o *ConsumerGroupData) SetLagSummary(v Relationship)`

SetLagSummary sets LagSummary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


