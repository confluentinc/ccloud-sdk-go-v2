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

### NewReplicaStatusData

`func NewReplicaStatusData(kind string, metadata ResourceMetadata, clusterId string, topicName string, brokerId int32, partitionId int32, leader bool, observer bool, isrEligible bool, inIsr bool, caughtUp bool, logStartOffset int32, logEndOffset int32, lastCaughtUpTimeMs int32, lastFetchTimeMs int32, ) *ReplicaStatusData`

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


### GetLeader

`func (o *ReplicaStatusData) GetLeader() bool`

GetLeader returns the Leader field if non-nil, zero value otherwise.

### GetLeaderOk

`func (o *ReplicaStatusData) GetLeaderOk() (*bool, bool)`

GetLeaderOk returns a tuple with the Leader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeader

`func (o *ReplicaStatusData) SetLeader(v bool)`

SetLeader sets Leader field to given value.


### GetObserver

`func (o *ReplicaStatusData) GetObserver() bool`

GetObserver returns the Observer field if non-nil, zero value otherwise.

### GetObserverOk

`func (o *ReplicaStatusData) GetObserverOk() (*bool, bool)`

GetObserverOk returns a tuple with the Observer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObserver

`func (o *ReplicaStatusData) SetObserver(v bool)`

SetObserver sets Observer field to given value.


### GetIsrEligible

`func (o *ReplicaStatusData) GetIsrEligible() bool`

GetIsrEligible returns the IsrEligible field if non-nil, zero value otherwise.

### GetIsrEligibleOk

`func (o *ReplicaStatusData) GetIsrEligibleOk() (*bool, bool)`

GetIsrEligibleOk returns a tuple with the IsrEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsrEligible

`func (o *ReplicaStatusData) SetIsrEligible(v bool)`

SetIsrEligible sets IsrEligible field to given value.


### GetInIsr

`func (o *ReplicaStatusData) GetInIsr() bool`

GetInIsr returns the InIsr field if non-nil, zero value otherwise.

### GetInIsrOk

`func (o *ReplicaStatusData) GetInIsrOk() (*bool, bool)`

GetInIsrOk returns a tuple with the InIsr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInIsr

`func (o *ReplicaStatusData) SetInIsr(v bool)`

SetInIsr sets InIsr field to given value.


### GetCaughtUp

`func (o *ReplicaStatusData) GetCaughtUp() bool`

GetCaughtUp returns the CaughtUp field if non-nil, zero value otherwise.

### GetCaughtUpOk

`func (o *ReplicaStatusData) GetCaughtUpOk() (*bool, bool)`

GetCaughtUpOk returns a tuple with the CaughtUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaughtUp

`func (o *ReplicaStatusData) SetCaughtUp(v bool)`

SetCaughtUp sets CaughtUp field to given value.


### GetLogStartOffset

`func (o *ReplicaStatusData) GetLogStartOffset() int32`

GetLogStartOffset returns the LogStartOffset field if non-nil, zero value otherwise.

### GetLogStartOffsetOk

`func (o *ReplicaStatusData) GetLogStartOffsetOk() (*int32, bool)`

GetLogStartOffsetOk returns a tuple with the LogStartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStartOffset

`func (o *ReplicaStatusData) SetLogStartOffset(v int32)`

SetLogStartOffset sets LogStartOffset field to given value.


### GetLogEndOffset

`func (o *ReplicaStatusData) GetLogEndOffset() int32`

GetLogEndOffset returns the LogEndOffset field if non-nil, zero value otherwise.

### GetLogEndOffsetOk

`func (o *ReplicaStatusData) GetLogEndOffsetOk() (*int32, bool)`

GetLogEndOffsetOk returns a tuple with the LogEndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogEndOffset

`func (o *ReplicaStatusData) SetLogEndOffset(v int32)`

SetLogEndOffset sets LogEndOffset field to given value.


### GetLastCaughtUpTimeMs

`func (o *ReplicaStatusData) GetLastCaughtUpTimeMs() int32`

GetLastCaughtUpTimeMs returns the LastCaughtUpTimeMs field if non-nil, zero value otherwise.

### GetLastCaughtUpTimeMsOk

`func (o *ReplicaStatusData) GetLastCaughtUpTimeMsOk() (*int32, bool)`

GetLastCaughtUpTimeMsOk returns a tuple with the LastCaughtUpTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCaughtUpTimeMs

`func (o *ReplicaStatusData) SetLastCaughtUpTimeMs(v int32)`

SetLastCaughtUpTimeMs sets LastCaughtUpTimeMs field to given value.


### GetLastFetchTimeMs

`func (o *ReplicaStatusData) GetLastFetchTimeMs() int32`

GetLastFetchTimeMs returns the LastFetchTimeMs field if non-nil, zero value otherwise.

### GetLastFetchTimeMsOk

`func (o *ReplicaStatusData) GetLastFetchTimeMsOk() (*int32, bool)`

GetLastFetchTimeMsOk returns a tuple with the LastFetchTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFetchTimeMs

`func (o *ReplicaStatusData) SetLastFetchTimeMs(v int32)`

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


