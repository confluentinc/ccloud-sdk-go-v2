# ClusterDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**Controller** | Pointer to [**Relationship**](Relationship.md) |  | [optional] 
**Acls** | [**Relationship**](Relationship.md) |  | 
**Brokers** | [**Relationship**](Relationship.md) |  | 
**BrokerConfigs** | [**Relationship**](Relationship.md) |  | 
**ConsumerGroups** | [**Relationship**](Relationship.md) |  | 
**Topics** | [**Relationship**](Relationship.md) |  | 
**PartitionReassignments** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewClusterDataAllOf

`func NewClusterDataAllOf(clusterId string, acls Relationship, brokers Relationship, brokerConfigs Relationship, consumerGroups Relationship, topics Relationship, partitionReassignments Relationship, ) *ClusterDataAllOf`

NewClusterDataAllOf instantiates a new ClusterDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDataAllOfWithDefaults

`func NewClusterDataAllOfWithDefaults() *ClusterDataAllOf`

NewClusterDataAllOfWithDefaults instantiates a new ClusterDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ClusterDataAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterDataAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterDataAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetController

`func (o *ClusterDataAllOf) GetController() Relationship`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *ClusterDataAllOf) GetControllerOk() (*Relationship, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *ClusterDataAllOf) SetController(v Relationship)`

SetController sets Controller field to given value.

### HasController

`func (o *ClusterDataAllOf) HasController() bool`

HasController returns a boolean if a field has been set.

### GetAcls

`func (o *ClusterDataAllOf) GetAcls() Relationship`

GetAcls returns the Acls field if non-nil, zero value otherwise.

### GetAclsOk

`func (o *ClusterDataAllOf) GetAclsOk() (*Relationship, bool)`

GetAclsOk returns a tuple with the Acls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcls

`func (o *ClusterDataAllOf) SetAcls(v Relationship)`

SetAcls sets Acls field to given value.


### GetBrokers

`func (o *ClusterDataAllOf) GetBrokers() Relationship`

GetBrokers returns the Brokers field if non-nil, zero value otherwise.

### GetBrokersOk

`func (o *ClusterDataAllOf) GetBrokersOk() (*Relationship, bool)`

GetBrokersOk returns a tuple with the Brokers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokers

`func (o *ClusterDataAllOf) SetBrokers(v Relationship)`

SetBrokers sets Brokers field to given value.


### GetBrokerConfigs

`func (o *ClusterDataAllOf) GetBrokerConfigs() Relationship`

GetBrokerConfigs returns the BrokerConfigs field if non-nil, zero value otherwise.

### GetBrokerConfigsOk

`func (o *ClusterDataAllOf) GetBrokerConfigsOk() (*Relationship, bool)`

GetBrokerConfigsOk returns a tuple with the BrokerConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerConfigs

`func (o *ClusterDataAllOf) SetBrokerConfigs(v Relationship)`

SetBrokerConfigs sets BrokerConfigs field to given value.


### GetConsumerGroups

`func (o *ClusterDataAllOf) GetConsumerGroups() Relationship`

GetConsumerGroups returns the ConsumerGroups field if non-nil, zero value otherwise.

### GetConsumerGroupsOk

`func (o *ClusterDataAllOf) GetConsumerGroupsOk() (*Relationship, bool)`

GetConsumerGroupsOk returns a tuple with the ConsumerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerGroups

`func (o *ClusterDataAllOf) SetConsumerGroups(v Relationship)`

SetConsumerGroups sets ConsumerGroups field to given value.


### GetTopics

`func (o *ClusterDataAllOf) GetTopics() Relationship`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *ClusterDataAllOf) GetTopicsOk() (*Relationship, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *ClusterDataAllOf) SetTopics(v Relationship)`

SetTopics sets Topics field to given value.


### GetPartitionReassignments

`func (o *ClusterDataAllOf) GetPartitionReassignments() Relationship`

GetPartitionReassignments returns the PartitionReassignments field if non-nil, zero value otherwise.

### GetPartitionReassignmentsOk

`func (o *ClusterDataAllOf) GetPartitionReassignmentsOk() (*Relationship, bool)`

GetPartitionReassignmentsOk returns a tuple with the PartitionReassignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionReassignments

`func (o *ClusterDataAllOf) SetPartitionReassignments(v Relationship)`

SetPartitionReassignments sets PartitionReassignments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


