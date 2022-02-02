# BrokerReplicaExclusionDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**BrokerId** | **int32** |  | 
**Reason** | **string** |  | 
**Broker** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewBrokerReplicaExclusionDataAllOf

`func NewBrokerReplicaExclusionDataAllOf(clusterId string, brokerId int32, reason string, broker Relationship, ) *BrokerReplicaExclusionDataAllOf`

NewBrokerReplicaExclusionDataAllOf instantiates a new BrokerReplicaExclusionDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrokerReplicaExclusionDataAllOfWithDefaults

`func NewBrokerReplicaExclusionDataAllOfWithDefaults() *BrokerReplicaExclusionDataAllOf`

NewBrokerReplicaExclusionDataAllOfWithDefaults instantiates a new BrokerReplicaExclusionDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *BrokerReplicaExclusionDataAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *BrokerReplicaExclusionDataAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *BrokerReplicaExclusionDataAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetBrokerId

`func (o *BrokerReplicaExclusionDataAllOf) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *BrokerReplicaExclusionDataAllOf) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *BrokerReplicaExclusionDataAllOf) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.


### GetReason

`func (o *BrokerReplicaExclusionDataAllOf) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BrokerReplicaExclusionDataAllOf) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BrokerReplicaExclusionDataAllOf) SetReason(v string)`

SetReason sets Reason field to given value.


### GetBroker

`func (o *BrokerReplicaExclusionDataAllOf) GetBroker() Relationship`

GetBroker returns the Broker field if non-nil, zero value otherwise.

### GetBrokerOk

`func (o *BrokerReplicaExclusionDataAllOf) GetBrokerOk() (*Relationship, bool)`

GetBrokerOk returns a tuple with the Broker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroker

`func (o *BrokerReplicaExclusionDataAllOf) SetBroker(v Relationship)`

SetBroker sets Broker field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


