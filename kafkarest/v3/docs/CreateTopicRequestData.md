# CreateTopicRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopicName** | **string** |  | 
**PartitionsCount** | Pointer to **int32** |  | [optional] 
**ReplicationFactor** | Pointer to **int32** |  | [optional] 
**ReplicasAssignments** | Pointer to [**[]CreateTopicRequestDataReplicasAssignments**](CreateTopicRequestDataReplicasAssignments.md) |  | [optional] 
**Configs** | Pointer to [**[]CreateTopicRequestDataConfigs**](CreateTopicRequestDataConfigs.md) |  | [optional] 

## Methods

### NewCreateTopicRequestData

`func NewCreateTopicRequestData(topicName string, ) *CreateTopicRequestData`

NewCreateTopicRequestData instantiates a new CreateTopicRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTopicRequestDataWithDefaults

`func NewCreateTopicRequestDataWithDefaults() *CreateTopicRequestData`

NewCreateTopicRequestDataWithDefaults instantiates a new CreateTopicRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopicName

`func (o *CreateTopicRequestData) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *CreateTopicRequestData) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *CreateTopicRequestData) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.


### GetPartitionsCount

`func (o *CreateTopicRequestData) GetPartitionsCount() int32`

GetPartitionsCount returns the PartitionsCount field if non-nil, zero value otherwise.

### GetPartitionsCountOk

`func (o *CreateTopicRequestData) GetPartitionsCountOk() (*int32, bool)`

GetPartitionsCountOk returns a tuple with the PartitionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionsCount

`func (o *CreateTopicRequestData) SetPartitionsCount(v int32)`

SetPartitionsCount sets PartitionsCount field to given value.

### HasPartitionsCount

`func (o *CreateTopicRequestData) HasPartitionsCount() bool`

HasPartitionsCount returns a boolean if a field has been set.

### GetReplicationFactor

`func (o *CreateTopicRequestData) GetReplicationFactor() int32`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *CreateTopicRequestData) GetReplicationFactorOk() (*int32, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *CreateTopicRequestData) SetReplicationFactor(v int32)`

SetReplicationFactor sets ReplicationFactor field to given value.

### HasReplicationFactor

`func (o *CreateTopicRequestData) HasReplicationFactor() bool`

HasReplicationFactor returns a boolean if a field has been set.

### GetReplicasAssignments

`func (o *CreateTopicRequestData) GetReplicasAssignments() []CreateTopicRequestDataReplicasAssignments`

GetReplicasAssignments returns the ReplicasAssignments field if non-nil, zero value otherwise.

### GetReplicasAssignmentsOk

`func (o *CreateTopicRequestData) GetReplicasAssignmentsOk() (*[]CreateTopicRequestDataReplicasAssignments, bool)`

GetReplicasAssignmentsOk returns a tuple with the ReplicasAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicasAssignments

`func (o *CreateTopicRequestData) SetReplicasAssignments(v []CreateTopicRequestDataReplicasAssignments)`

SetReplicasAssignments sets ReplicasAssignments field to given value.

### HasReplicasAssignments

`func (o *CreateTopicRequestData) HasReplicasAssignments() bool`

HasReplicasAssignments returns a boolean if a field has been set.

### GetConfigs

`func (o *CreateTopicRequestData) GetConfigs() []CreateTopicRequestDataConfigs`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *CreateTopicRequestData) GetConfigsOk() (*[]CreateTopicRequestDataConfigs, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *CreateTopicRequestData) SetConfigs(v []CreateTopicRequestDataConfigs)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *CreateTopicRequestData) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


