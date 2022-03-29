# ReplicaStatusData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
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

### NewReplicaStatusData

`func NewReplicaStatusData(kind string, metadata ResourceMetadata, clusterId string, topicName string, brokerId int32, partitionId int32, isLeader bool, isObserver bool, isIsrEligible bool, isInIsr bool, isCaughtUp bool, logStartOffset int64, logEndOffset int64, lastCaughtUpTimeMs int64, lastFetchTimeMs int64, ) *ReplicaStatusData`

NewReplicaStatusData instantiates a new ReplicaStatusData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicaStatusDataWithDefaults

`func NewReplicaStatusDataWithDefaults() *ReplicaStatusData`

NewReplicaStatusDataWithDefaults instantiates a new ReplicaStatusData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ReplicaStatusData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ReplicaStatusData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ReplicaStatusData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ReplicaStatusData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ReplicaStatusData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ReplicaStatusData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *ReplicaStatusData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ReplicaStatusData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ReplicaStatusData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetTopicName

`func (o *ReplicaStatusData) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *ReplicaStatusData) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *ReplicaStatusData) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.


### GetBrokerId

`func (o *ReplicaStatusData) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *ReplicaStatusData) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *ReplicaStatusData) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.


### GetPartitionId

`func (o *ReplicaStatusData) GetPartitionId() int32`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *ReplicaStatusData) GetPartitionIdOk() (*int32, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *ReplicaStatusData) SetPartitionId(v int32)`

SetPartitionId sets PartitionId field to given value.


### GetIsLeader

`func (o *ReplicaStatusData) GetIsLeader() bool`

GetIsLeader returns the IsLeader field if non-nil, zero value otherwise.

### GetIsLeaderOk

`func (o *ReplicaStatusData) GetIsLeaderOk() (*bool, bool)`

GetIsLeaderOk returns a tuple with the IsLeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLeader

`func (o *ReplicaStatusData) SetIsLeader(v bool)`

SetIsLeader sets IsLeader field to given value.


### GetIsObserver

`func (o *ReplicaStatusData) GetIsObserver() bool`

GetIsObserver returns the IsObserver field if non-nil, zero value otherwise.

### GetIsObserverOk

`func (o *ReplicaStatusData) GetIsObserverOk() (*bool, bool)`

GetIsObserverOk returns a tuple with the IsObserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsObserver

`func (o *ReplicaStatusData) SetIsObserver(v bool)`

SetIsObserver sets IsObserver field to given value.


### GetIsIsrEligible

`func (o *ReplicaStatusData) GetIsIsrEligible() bool`

GetIsIsrEligible returns the IsIsrEligible field if non-nil, zero value otherwise.

### GetIsIsrEligibleOk

`func (o *ReplicaStatusData) GetIsIsrEligibleOk() (*bool, bool)`

GetIsIsrEligibleOk returns a tuple with the IsIsrEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIsrEligible

`func (o *ReplicaStatusData) SetIsIsrEligible(v bool)`

SetIsIsrEligible sets IsIsrEligible field to given value.


### GetIsInIsr

`func (o *ReplicaStatusData) GetIsInIsr() bool`

GetIsInIsr returns the IsInIsr field if non-nil, zero value otherwise.

### GetIsInIsrOk

`func (o *ReplicaStatusData) GetIsInIsrOk() (*bool, bool)`

GetIsInIsrOk returns a tuple with the IsInIsr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInIsr

`func (o *ReplicaStatusData) SetIsInIsr(v bool)`

SetIsInIsr sets IsInIsr field to given value.


### GetIsCaughtUp

`func (o *ReplicaStatusData) GetIsCaughtUp() bool`

GetIsCaughtUp returns the IsCaughtUp field if non-nil, zero value otherwise.

### GetIsCaughtUpOk

`func (o *ReplicaStatusData) GetIsCaughtUpOk() (*bool, bool)`

GetIsCaughtUpOk returns a tuple with the IsCaughtUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCaughtUp

`func (o *ReplicaStatusData) SetIsCaughtUp(v bool)`

SetIsCaughtUp sets IsCaughtUp field to given value.


### GetLogStartOffset

`func (o *ReplicaStatusData) GetLogStartOffset() int64`

GetLogStartOffset returns the LogStartOffset field if non-nil, zero value otherwise.

### GetLogStartOffsetOk

`func (o *ReplicaStatusData) GetLogStartOffsetOk() (*int64, bool)`

GetLogStartOffsetOk returns a tuple with the LogStartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStartOffset

`func (o *ReplicaStatusData) SetLogStartOffset(v int64)`

SetLogStartOffset sets LogStartOffset field to given value.


### GetLogEndOffset

`func (o *ReplicaStatusData) GetLogEndOffset() int64`

GetLogEndOffset returns the LogEndOffset field if non-nil, zero value otherwise.

### GetLogEndOffsetOk

`func (o *ReplicaStatusData) GetLogEndOffsetOk() (*int64, bool)`

GetLogEndOffsetOk returns a tuple with the LogEndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogEndOffset

`func (o *ReplicaStatusData) SetLogEndOffset(v int64)`

SetLogEndOffset sets LogEndOffset field to given value.


### GetLastCaughtUpTimeMs

`func (o *ReplicaStatusData) GetLastCaughtUpTimeMs() int64`

GetLastCaughtUpTimeMs returns the LastCaughtUpTimeMs field if non-nil, zero value otherwise.

### GetLastCaughtUpTimeMsOk

`func (o *ReplicaStatusData) GetLastCaughtUpTimeMsOk() (*int64, bool)`

GetLastCaughtUpTimeMsOk returns a tuple with the LastCaughtUpTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCaughtUpTimeMs

`func (o *ReplicaStatusData) SetLastCaughtUpTimeMs(v int64)`

SetLastCaughtUpTimeMs sets LastCaughtUpTimeMs field to given value.


### GetLastFetchTimeMs

`func (o *ReplicaStatusData) GetLastFetchTimeMs() int64`

GetLastFetchTimeMs returns the LastFetchTimeMs field if non-nil, zero value otherwise.

### GetLastFetchTimeMsOk

`func (o *ReplicaStatusData) GetLastFetchTimeMsOk() (*int64, bool)`

GetLastFetchTimeMsOk returns a tuple with the LastFetchTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFetchTimeMs

`func (o *ReplicaStatusData) SetLastFetchTimeMs(v int64)`

SetLastFetchTimeMs sets LastFetchTimeMs field to given value.


### GetLinkName

`func (o *ReplicaStatusData) GetLinkName() string`

GetLinkName returns the LinkName field if non-nil, zero value otherwise.

### GetLinkNameOk

`func (o *ReplicaStatusData) GetLinkNameOk() (*string, bool)`

GetLinkNameOk returns a tuple with the LinkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkName

`func (o *ReplicaStatusData) SetLinkName(v string)`

SetLinkName sets LinkName field to given value.

### HasLinkName

`func (o *ReplicaStatusData) HasLinkName() bool`

HasLinkName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


