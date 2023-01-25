# RemoveBrokerTaskDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**BrokerId** | **int32** |  | 
**ShutdownScheduled** | **bool** |  | 
**BrokerReplicaExclusionStatus** | **string** |  | 
**PartitionReassignmentStatus** | **string** |  | 
**BrokerShutdownStatus** | **string** |  | 
**ErrorCode** | Pointer to **NullableInt32** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**Broker** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewRemoveBrokerTaskDataAllOf

`func NewRemoveBrokerTaskDataAllOf(clusterId string, brokerId int32, shutdownScheduled bool, brokerReplicaExclusionStatus string, partitionReassignmentStatus string, brokerShutdownStatus string, broker Relationship, ) *RemoveBrokerTaskDataAllOf`

NewRemoveBrokerTaskDataAllOf instantiates a new RemoveBrokerTaskDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveBrokerTaskDataAllOfWithDefaults

`func NewRemoveBrokerTaskDataAllOfWithDefaults() *RemoveBrokerTaskDataAllOf`

NewRemoveBrokerTaskDataAllOfWithDefaults instantiates a new RemoveBrokerTaskDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *RemoveBrokerTaskDataAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *RemoveBrokerTaskDataAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *RemoveBrokerTaskDataAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetBrokerId

`func (o *RemoveBrokerTaskDataAllOf) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *RemoveBrokerTaskDataAllOf) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *RemoveBrokerTaskDataAllOf) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.


### GetShutdownScheduled

`func (o *RemoveBrokerTaskDataAllOf) GetShutdownScheduled() bool`

GetShutdownScheduled returns the ShutdownScheduled field if non-nil, zero value otherwise.

### GetShutdownScheduledOk

`func (o *RemoveBrokerTaskDataAllOf) GetShutdownScheduledOk() (*bool, bool)`

GetShutdownScheduledOk returns a tuple with the ShutdownScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownScheduled

`func (o *RemoveBrokerTaskDataAllOf) SetShutdownScheduled(v bool)`

SetShutdownScheduled sets ShutdownScheduled field to given value.


### GetBrokerReplicaExclusionStatus

`func (o *RemoveBrokerTaskDataAllOf) GetBrokerReplicaExclusionStatus() string`

GetBrokerReplicaExclusionStatus returns the BrokerReplicaExclusionStatus field if non-nil, zero value otherwise.

### GetBrokerReplicaExclusionStatusOk

`func (o *RemoveBrokerTaskDataAllOf) GetBrokerReplicaExclusionStatusOk() (*string, bool)`

GetBrokerReplicaExclusionStatusOk returns a tuple with the BrokerReplicaExclusionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerReplicaExclusionStatus

`func (o *RemoveBrokerTaskDataAllOf) SetBrokerReplicaExclusionStatus(v string)`

SetBrokerReplicaExclusionStatus sets BrokerReplicaExclusionStatus field to given value.


### GetPartitionReassignmentStatus

`func (o *RemoveBrokerTaskDataAllOf) GetPartitionReassignmentStatus() string`

GetPartitionReassignmentStatus returns the PartitionReassignmentStatus field if non-nil, zero value otherwise.

### GetPartitionReassignmentStatusOk

`func (o *RemoveBrokerTaskDataAllOf) GetPartitionReassignmentStatusOk() (*string, bool)`

GetPartitionReassignmentStatusOk returns a tuple with the PartitionReassignmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionReassignmentStatus

`func (o *RemoveBrokerTaskDataAllOf) SetPartitionReassignmentStatus(v string)`

SetPartitionReassignmentStatus sets PartitionReassignmentStatus field to given value.


### GetBrokerShutdownStatus

`func (o *RemoveBrokerTaskDataAllOf) GetBrokerShutdownStatus() string`

GetBrokerShutdownStatus returns the BrokerShutdownStatus field if non-nil, zero value otherwise.

### GetBrokerShutdownStatusOk

`func (o *RemoveBrokerTaskDataAllOf) GetBrokerShutdownStatusOk() (*string, bool)`

GetBrokerShutdownStatusOk returns a tuple with the BrokerShutdownStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerShutdownStatus

`func (o *RemoveBrokerTaskDataAllOf) SetBrokerShutdownStatus(v string)`

SetBrokerShutdownStatus sets BrokerShutdownStatus field to given value.


### GetErrorCode

`func (o *RemoveBrokerTaskDataAllOf) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *RemoveBrokerTaskDataAllOf) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *RemoveBrokerTaskDataAllOf) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *RemoveBrokerTaskDataAllOf) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *RemoveBrokerTaskDataAllOf) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *RemoveBrokerTaskDataAllOf) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorMessage

`func (o *RemoveBrokerTaskDataAllOf) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *RemoveBrokerTaskDataAllOf) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *RemoveBrokerTaskDataAllOf) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *RemoveBrokerTaskDataAllOf) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *RemoveBrokerTaskDataAllOf) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *RemoveBrokerTaskDataAllOf) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetBroker

`func (o *RemoveBrokerTaskDataAllOf) GetBroker() Relationship`

GetBroker returns the Broker field if non-nil, zero value otherwise.

### GetBrokerOk

`func (o *RemoveBrokerTaskDataAllOf) GetBrokerOk() (*Relationship, bool)`

GetBrokerOk returns a tuple with the Broker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroker

`func (o *RemoveBrokerTaskDataAllOf) SetBroker(v Relationship)`

SetBroker sets Broker field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


