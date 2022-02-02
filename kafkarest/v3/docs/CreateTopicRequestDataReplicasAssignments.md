# CreateTopicRequestDataReplicasAssignments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartitionId** | **int32** |  | 
**BrokerIds** | **[]int32** |  | 

## Methods

### NewCreateTopicRequestDataReplicasAssignments

`func NewCreateTopicRequestDataReplicasAssignments(partitionId int32, brokerIds []int32, ) *CreateTopicRequestDataReplicasAssignments`

NewCreateTopicRequestDataReplicasAssignments instantiates a new CreateTopicRequestDataReplicasAssignments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTopicRequestDataReplicasAssignmentsWithDefaults

`func NewCreateTopicRequestDataReplicasAssignmentsWithDefaults() *CreateTopicRequestDataReplicasAssignments`

NewCreateTopicRequestDataReplicasAssignmentsWithDefaults instantiates a new CreateTopicRequestDataReplicasAssignments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartitionId

`func (o *CreateTopicRequestDataReplicasAssignments) GetPartitionId() int32`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *CreateTopicRequestDataReplicasAssignments) GetPartitionIdOk() (*int32, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *CreateTopicRequestDataReplicasAssignments) SetPartitionId(v int32)`

SetPartitionId sets PartitionId field to given value.


### GetBrokerIds

`func (o *CreateTopicRequestDataReplicasAssignments) GetBrokerIds() []int32`

GetBrokerIds returns the BrokerIds field if non-nil, zero value otherwise.

### GetBrokerIdsOk

`func (o *CreateTopicRequestDataReplicasAssignments) GetBrokerIdsOk() (*[]int32, bool)`

GetBrokerIdsOk returns a tuple with the BrokerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerIds

`func (o *CreateTopicRequestDataReplicasAssignments) SetBrokerIds(v []int32)`

SetBrokerIds sets BrokerIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


