# BrokerReplicaExclusionRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrokerId** | **int32** |  | 
**Reason** | **string** |  | 

## Methods

### NewBrokerReplicaExclusionRequestData

`func NewBrokerReplicaExclusionRequestData(brokerId int32, reason string, ) *BrokerReplicaExclusionRequestData`

NewBrokerReplicaExclusionRequestData instantiates a new BrokerReplicaExclusionRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrokerReplicaExclusionRequestDataWithDefaults

`func NewBrokerReplicaExclusionRequestDataWithDefaults() *BrokerReplicaExclusionRequestData`

NewBrokerReplicaExclusionRequestDataWithDefaults instantiates a new BrokerReplicaExclusionRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrokerId

`func (o *BrokerReplicaExclusionRequestData) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *BrokerReplicaExclusionRequestData) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *BrokerReplicaExclusionRequestData) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.


### GetReason

`func (o *BrokerReplicaExclusionRequestData) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BrokerReplicaExclusionRequestData) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BrokerReplicaExclusionRequestData) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


