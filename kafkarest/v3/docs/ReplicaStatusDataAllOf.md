# ReplicaStatusDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**TopicName** | **string** |  | 
**BrokerId** | **int32** |  | 
**PartitionId** | **int32** |  | 
**IsLeader** | **bool** |  | 
**IsObserver** | **bool** |  | 
**IsIsrEligible** | **bool** |  | 
**IsInIsr** | **bool** |  | 
**IsCaughtUp** | **bool** |  | 
**LogStartOffset** | **int64** |  | 
**LogEndOffset** | **int64** |  | 
**LastCaughtUpTimeMs** | **int64** |  | 
**LastFetchTimeMs** | **int64** |  | 
**LinkName** | Pointer to **string** |  | [optional] 

## Methods

### NewReplicaStatusDataAllOf

`func NewReplicaStatusDataAllOf(clusterId string, topicName string, brokerId int32, partitionId int32, isLeader bool, isObserver bool, isIsrEligible bool, isInIsr bool, isCaughtUp bool, logStartOffset int64, logEndOffset int64, lastCaughtUpTimeMs int64, lastFetchTimeMs int64, ) *ReplicaStatusDataAllOf`

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


### GetIsLeader

`func (o *ReplicaStatusDataAllOf) GetIsLeader() bool`

GetIsLeader returns the IsLeader field if non-nil, zero value otherwise.

### GetIsLeaderOk

`func (o *ReplicaStatusDataAllOf) GetIsLeaderOk() (*bool, bool)`

GetIsLeaderOk returns a tuple with the IsLeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLeader

`func (o *ReplicaStatusDataAllOf) SetIsLeader(v bool)`

SetIsLeader sets IsLeader field to given value.


### GetIsObserver

`func (o *ReplicaStatusDataAllOf) GetIsObserver() bool`

GetIsObserver returns the IsObserver field if non-nil, zero value otherwise.

### GetIsObserverOk

`func (o *ReplicaStatusDataAllOf) GetIsObserverOk() (*bool, bool)`

GetIsObserverOk returns a tuple with the IsObserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsObserver

`func (o *ReplicaStatusDataAllOf) SetIsObserver(v bool)`

SetIsObserver sets IsObserver field to given value.


### GetIsIsrEligible

`func (o *ReplicaStatusDataAllOf) GetIsIsrEligible() bool`

GetIsIsrEligible returns the IsIsrEligible field if non-nil, zero value otherwise.

### GetIsIsrEligibleOk

`func (o *ReplicaStatusDataAllOf) GetIsIsrEligibleOk() (*bool, bool)`

GetIsIsrEligibleOk returns a tuple with the IsIsrEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIsrEligible

`func (o *ReplicaStatusDataAllOf) SetIsIsrEligible(v bool)`

SetIsIsrEligible sets IsIsrEligible field to given value.


### GetIsInIsr

`func (o *ReplicaStatusDataAllOf) GetIsInIsr() bool`

GetIsInIsr returns the IsInIsr field if non-nil, zero value otherwise.

### GetIsInIsrOk

`func (o *ReplicaStatusDataAllOf) GetIsInIsrOk() (*bool, bool)`

GetIsInIsrOk returns a tuple with the IsInIsr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInIsr

`func (o *ReplicaStatusDataAllOf) SetIsInIsr(v bool)`

SetIsInIsr sets IsInIsr field to given value.


### GetIsCaughtUp

`func (o *ReplicaStatusDataAllOf) GetIsCaughtUp() bool`

GetIsCaughtUp returns the IsCaughtUp field if non-nil, zero value otherwise.

### GetIsCaughtUpOk

`func (o *ReplicaStatusDataAllOf) GetIsCaughtUpOk() (*bool, bool)`

GetIsCaughtUpOk returns a tuple with the IsCaughtUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCaughtUp

`func (o *ReplicaStatusDataAllOf) SetIsCaughtUp(v bool)`

SetIsCaughtUp sets IsCaughtUp field to given value.


### GetLogStartOffset

`func (o *ReplicaStatusDataAllOf) GetLogStartOffset() int64`

GetLogStartOffset returns the LogStartOffset field if non-nil, zero value otherwise.

### GetLogStartOffsetOk

`func (o *ReplicaStatusDataAllOf) GetLogStartOffsetOk() (*int64, bool)`

GetLogStartOffsetOk returns a tuple with the LogStartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStartOffset

`func (o *ReplicaStatusDataAllOf) SetLogStartOffset(v int64)`

SetLogStartOffset sets LogStartOffset field to given value.


### GetLogEndOffset

`func (o *ReplicaStatusDataAllOf) GetLogEndOffset() int64`

GetLogEndOffset returns the LogEndOffset field if non-nil, zero value otherwise.

### GetLogEndOffsetOk

`func (o *ReplicaStatusDataAllOf) GetLogEndOffsetOk() (*int64, bool)`

GetLogEndOffsetOk returns a tuple with the LogEndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogEndOffset

`func (o *ReplicaStatusDataAllOf) SetLogEndOffset(v int64)`

SetLogEndOffset sets LogEndOffset field to given value.


### GetLastCaughtUpTimeMs

`func (o *ReplicaStatusDataAllOf) GetLastCaughtUpTimeMs() int64`

GetLastCaughtUpTimeMs returns the LastCaughtUpTimeMs field if non-nil, zero value otherwise.

### GetLastCaughtUpTimeMsOk

`func (o *ReplicaStatusDataAllOf) GetLastCaughtUpTimeMsOk() (*int64, bool)`

GetLastCaughtUpTimeMsOk returns a tuple with the LastCaughtUpTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCaughtUpTimeMs

`func (o *ReplicaStatusDataAllOf) SetLastCaughtUpTimeMs(v int64)`

SetLastCaughtUpTimeMs sets LastCaughtUpTimeMs field to given value.


### GetLastFetchTimeMs

`func (o *ReplicaStatusDataAllOf) GetLastFetchTimeMs() int64`

GetLastFetchTimeMs returns the LastFetchTimeMs field if non-nil, zero value otherwise.

### GetLastFetchTimeMsOk

`func (o *ReplicaStatusDataAllOf) GetLastFetchTimeMsOk() (*int64, bool)`

GetLastFetchTimeMsOk returns a tuple with the LastFetchTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFetchTimeMs

`func (o *ReplicaStatusDataAllOf) SetLastFetchTimeMs(v int64)`

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


