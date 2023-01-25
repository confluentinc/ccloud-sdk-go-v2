# ClusterData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**ClusterId** | **string** |  | 
**Controller** | Pointer to [**Relationship**](Relationship.md) |  | [optional] 
**Acls** | [**Relationship**](Relationship.md) |  | 
**Brokers** | [**Relationship**](Relationship.md) |  | 
**BrokerConfigs** | [**Relationship**](Relationship.md) |  | 
**ConsumerGroups** | [**Relationship**](Relationship.md) |  | 
**Topics** | [**Relationship**](Relationship.md) |  | 
**PartitionReassignments** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewClusterData

`func NewClusterData(kind string, metadata ResourceMetadata, clusterId string, acls Relationship, brokers Relationship, brokerConfigs Relationship, consumerGroups Relationship, topics Relationship, partitionReassignments Relationship, ) *ClusterData`

NewClusterData instantiates a new ClusterData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDataWithDefaults

`func NewClusterDataWithDefaults() *ClusterData`

NewClusterDataWithDefaults instantiates a new ClusterData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ClusterData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ClusterData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ClusterData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ClusterData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *ClusterData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetController

`func (o *ClusterData) GetController() Relationship`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *ClusterData) GetControllerOk() (*Relationship, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *ClusterData) SetController(v Relationship)`

SetController sets Controller field to given value.

### HasController

`func (o *ClusterData) HasController() bool`

HasController returns a boolean if a field has been set.

### GetAcls

`func (o *ClusterData) GetAcls() Relationship`

GetAcls returns the Acls field if non-nil, zero value otherwise.

### GetAclsOk

`func (o *ClusterData) GetAclsOk() (*Relationship, bool)`

GetAclsOk returns a tuple with the Acls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcls

`func (o *ClusterData) SetAcls(v Relationship)`

SetAcls sets Acls field to given value.


### GetBrokers

`func (o *ClusterData) GetBrokers() Relationship`

GetBrokers returns the Brokers field if non-nil, zero value otherwise.

### GetBrokersOk

`func (o *ClusterData) GetBrokersOk() (*Relationship, bool)`

GetBrokersOk returns a tuple with the Brokers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokers

`func (o *ClusterData) SetBrokers(v Relationship)`

SetBrokers sets Brokers field to given value.


### GetBrokerConfigs

`func (o *ClusterData) GetBrokerConfigs() Relationship`

GetBrokerConfigs returns the BrokerConfigs field if non-nil, zero value otherwise.

### GetBrokerConfigsOk

`func (o *ClusterData) GetBrokerConfigsOk() (*Relationship, bool)`

GetBrokerConfigsOk returns a tuple with the BrokerConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerConfigs

`func (o *ClusterData) SetBrokerConfigs(v Relationship)`

SetBrokerConfigs sets BrokerConfigs field to given value.


### GetConsumerGroups

`func (o *ClusterData) GetConsumerGroups() Relationship`

GetConsumerGroups returns the ConsumerGroups field if non-nil, zero value otherwise.

### GetConsumerGroupsOk

`func (o *ClusterData) GetConsumerGroupsOk() (*Relationship, bool)`

GetConsumerGroupsOk returns a tuple with the ConsumerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerGroups

`func (o *ClusterData) SetConsumerGroups(v Relationship)`

SetConsumerGroups sets ConsumerGroups field to given value.


### GetTopics

`func (o *ClusterData) GetTopics() Relationship`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *ClusterData) GetTopicsOk() (*Relationship, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *ClusterData) SetTopics(v Relationship)`

SetTopics sets Topics field to given value.


### GetPartitionReassignments

`func (o *ClusterData) GetPartitionReassignments() Relationship`

GetPartitionReassignments returns the PartitionReassignments field if non-nil, zero value otherwise.

### GetPartitionReassignmentsOk

`func (o *ClusterData) GetPartitionReassignmentsOk() (*Relationship, bool)`

GetPartitionReassignmentsOk returns a tuple with the PartitionReassignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionReassignments

`func (o *ClusterData) SetPartitionReassignments(v Relationship)`

SetPartitionReassignments sets PartitionReassignments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


