# ReplicaStatusDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**TopicName** | **string** |  | 
**BrokerId** | **int32** |  | 
**PartitionId** | **int32** |  | 
**Leader** | **bool** |  | 
**Observer** | **bool** |  | 
**IsrEligible** | **bool** |  | 
**InIsr** | **bool** |  | 
**CaughtUp** | **bool** |  | 
**LogStartOffset** | **int32** |  | 
**LogEndOffset** | **int32** |  | 
**LastCaughtUpTimeMs** | **int32** |  | 
**LastFetchTimeMs** | **int32** |  | 
**LinkName** | Pointer to **string** |  | [optional] 

## Methods

### NewReplicaStatusDataAllOf

`func NewReplicaStatusDataAllOf(clusterId string, topicName string, brokerId int32, partitionId int32, leader bool, observer bool, isrEligible bool, inIsr bool, caughtUp bool, logStartOffset int32, logEndOffset int32, lastCaughtUpTimeMs int32, lastFetchTimeMs int32, ) *ReplicaStatusDataAllOf`

NewReplicaStatusDataAllOf instantiates a new ReplicaStatusDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicaStatusDataAllOfWithDefaults

`func NewReplicaStatusDataAllOfWithDefaults() *ReplicaStatusDataAllOf`

NewReplicaStatusDataAllOfWithDefaults instantiates a new ReplicaStatusDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ReplicaStatusDataAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ReplicaStatusDataAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ReplicaStatusDataAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetTopicName

`func (o *ReplicaStatusDataAllOf) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *ReplicaStatusDataAllOf) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *ReplicaStatusDataAllOf) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.


### GetBrokerId

`func (o *ReplicaStatusDataAllOf) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *ReplicaStatusDataAllOf) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *ReplicaStatusDataAllOf) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.


### GetPartitionId

`func (o *ReplicaStatusDataAllOf) GetPartitionId() int32`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *ReplicaStatusDataAllOf) GetPartitionIdOk() (*int32, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *ReplicaStatusDataAllOf) SetPartitionId(v int32)`

SetPartitionId sets PartitionId field to given value.


### GetLeader

`func (o *ReplicaStatusDataAllOf) GetLeader() bool`

GetLeader returns the Leader field if non-nil, zero value otherwise.

### GetLeaderOk

`func (o *ReplicaStatusDataAllOf) GetLeaderOk() (*bool, bool)`

GetLeaderOk returns a tuple with the Leader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeader

`func (o *ReplicaStatusDataAllOf) SetLeader(v bool)`

SetLeader sets Leader field to given value.


### GetObserver

`func (o *ReplicaStatusDataAllOf) GetObserver() bool`

GetObserver returns the Observer field if non-nil, zero value otherwise.

### GetObserverOk

`func (o *ReplicaStatusDataAllOf) GetObserverOk() (*bool, bool)`

GetObserverOk returns a tuple with the Observer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObserver

`func (o *ReplicaStatusDataAllOf) SetObserver(v bool)`

SetObserver sets Observer field to given value.


### GetIsrEligible

`func (o *ReplicaStatusDataAllOf) GetIsrEligible() bool`

GetIsrEligible returns the IsrEligible field if non-nil, zero value otherwise.

### GetIsrEligibleOk

`func (o *ReplicaStatusDataAllOf) GetIsrEligibleOk() (*bool, bool)`

GetIsrEligibleOk returns a tuple with the IsrEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsrEligible

`func (o *ReplicaStatusDataAllOf) SetIsrEligible(v bool)`

SetIsrEligible sets IsrEligible field to given value.


### GetInIsr

`func (o *ReplicaStatusDataAllOf) GetInIsr() bool`

GetInIsr returns the InIsr field if non-nil, zero value otherwise.

### GetInIsrOk

`func (o *ReplicaStatusDataAllOf) GetInIsrOk() (*bool, bool)`

GetInIsrOk returns a tuple with the InIsr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInIsr

`func (o *ReplicaStatusDataAllOf) SetInIsr(v bool)`

SetInIsr sets InIsr field to given value.


### GetCaughtUp

`func (o *ReplicaStatusDataAllOf) GetCaughtUp() bool`

GetCaughtUp returns the CaughtUp field if non-nil, zero value otherwise.

### GetCaughtUpOk

`func (o *ReplicaStatusDataAllOf) GetCaughtUpOk() (*bool, bool)`

GetCaughtUpOk returns a tuple with the CaughtUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaughtUp

`func (o *ReplicaStatusDataAllOf) SetCaughtUp(v bool)`

SetCaughtUp sets CaughtUp field to given value.


### GetLogStartOffset

`func (o *ReplicaStatusDataAllOf) GetLogStartOffset() int32`

GetLogStartOffset returns the LogStartOffset field if non-nil, zero value otherwise.

### GetLogStartOffsetOk

`func (o *ReplicaStatusDataAllOf) GetLogStartOffsetOk() (*int32, bool)`

GetLogStartOffsetOk returns a tuple with the LogStartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStartOffset

`func (o *ReplicaStatusDataAllOf) SetLogStartOffset(v int32)`

SetLogStartOffset sets LogStartOffset field to given value.


### GetLogEndOffset

`func (o *ReplicaStatusDataAllOf) GetLogEndOffset() int32`

GetLogEndOffset returns the LogEndOffset field if non-nil, zero value otherwise.

### GetLogEndOffsetOk

`func (o *ReplicaStatusDataAllOf) GetLogEndOffsetOk() (*int32, bool)`

GetLogEndOffsetOk returns a tuple with the LogEndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogEndOffset

`func (o *ReplicaStatusDataAllOf) SetLogEndOffset(v int32)`

SetLogEndOffset sets LogEndOffset field to given value.


### GetLastCaughtUpTimeMs

`func (o *ReplicaStatusDataAllOf) GetLastCaughtUpTimeMs() int32`

GetLastCaughtUpTimeMs returns the LastCaughtUpTimeMs field if non-nil, zero value otherwise.

### GetLastCaughtUpTimeMsOk

`func (o *ReplicaStatusDataAllOf) GetLastCaughtUpTimeMsOk() (*int32, bool)`

GetLastCaughtUpTimeMsOk returns a tuple with the LastCaughtUpTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCaughtUpTimeMs

`func (o *ReplicaStatusDataAllOf) SetLastCaughtUpTimeMs(v int32)`

SetLastCaughtUpTimeMs sets LastCaughtUpTimeMs field to given value.


### GetLastFetchTimeMs

`func (o *ReplicaStatusDataAllOf) GetLastFetchTimeMs() int32`

GetLastFetchTimeMs returns the LastFetchTimeMs field if non-nil, zero value otherwise.

### GetLastFetchTimeMsOk

`func (o *ReplicaStatusDataAllOf) GetLastFetchTimeMsOk() (*int32, bool)`

GetLastFetchTimeMsOk returns a tuple with the LastFetchTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFetchTimeMs

`func (o *ReplicaStatusDataAllOf) SetLastFetchTimeMs(v int32)`

SetLastFetchTimeMs sets LastFetchTimeMs field to given value.


### GetLinkName

`func (o *ReplicaStatusDataAllOf) GetLinkName() string`

GetLinkName returns the LinkName field if non-nil, zero value otherwise.

### GetLinkNameOk

`func (o *ReplicaStatusDataAllOf) GetLinkNameOk() (*string, bool)`

GetLinkNameOk returns a tuple with the LinkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkName

`func (o *ReplicaStatusDataAllOf) SetLinkName(v string)`

SetLinkName sets LinkName field to given value.

### HasLinkName

`func (o *ReplicaStatusDataAllOf) HasLinkName() bool`

HasLinkName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


