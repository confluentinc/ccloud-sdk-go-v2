# BrokerReplicaExclusionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**ClusterId** | **string** |  | 
**BrokerId** | **int32** |  | 
**Reason** | **string** |  | 
**Broker** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewBrokerReplicaExclusionData

`func NewBrokerReplicaExclusionData(kind string, metadata ResourceMetadata, clusterId string, brokerId int32, reason string, broker Relationship, ) *BrokerReplicaExclusionData`

NewBrokerReplicaExclusionData instantiates a new BrokerReplicaExclusionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrokerReplicaExclusionDataWithDefaults

`func NewBrokerReplicaExclusionDataWithDefaults() *BrokerReplicaExclusionData`

NewBrokerReplicaExclusionDataWithDefaults instantiates a new BrokerReplicaExclusionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *BrokerReplicaExclusionData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BrokerReplicaExclusionData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BrokerReplicaExclusionData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *BrokerReplicaExclusionData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BrokerReplicaExclusionData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BrokerReplicaExclusionData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *BrokerReplicaExclusionData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *BrokerReplicaExclusionData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *BrokerReplicaExclusionData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetBrokerId

`func (o *BrokerReplicaExclusionData) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *BrokerReplicaExclusionData) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *BrokerReplicaExclusionData) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.


### GetReason

`func (o *BrokerReplicaExclusionData) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BrokerReplicaExclusionData) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BrokerReplicaExclusionData) SetReason(v string)`

SetReason sets Reason field to given value.


### GetBroker

`func (o *BrokerReplicaExclusionData) GetBroker() Relationship`

GetBroker returns the Broker field if non-nil, zero value otherwise.

### GetBrokerOk

`func (o *BrokerReplicaExclusionData) GetBrokerOk() (*Relationship, bool)`

GetBrokerOk returns a tuple with the Broker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroker

`func (o *BrokerReplicaExclusionData) SetBroker(v Relationship)`

SetBroker sets Broker field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


