# RemoveBrokerTaskData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
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

### NewRemoveBrokerTaskData

`func NewRemoveBrokerTaskData(kind string, metadata ResourceMetadata, clusterId string, brokerId int32, shutdownScheduled bool, brokerReplicaExclusionStatus string, partitionReassignmentStatus string, brokerShutdownStatus string, broker Relationship, ) *RemoveBrokerTaskData`

NewRemoveBrokerTaskData instantiates a new RemoveBrokerTaskData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveBrokerTaskDataWithDefaults

`func NewRemoveBrokerTaskDataWithDefaults() *RemoveBrokerTaskData`

NewRemoveBrokerTaskDataWithDefaults instantiates a new RemoveBrokerTaskData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *RemoveBrokerTaskData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RemoveBrokerTaskData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RemoveBrokerTaskData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *RemoveBrokerTaskData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RemoveBrokerTaskData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RemoveBrokerTaskData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *RemoveBrokerTaskData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *RemoveBrokerTaskData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *RemoveBrokerTaskData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetBrokerId

`func (o *RemoveBrokerTaskData) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *RemoveBrokerTaskData) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *RemoveBrokerTaskData) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.


### GetShutdownScheduled

`func (o *RemoveBrokerTaskData) GetShutdownScheduled() bool`

GetShutdownScheduled returns the ShutdownScheduled field if non-nil, zero value otherwise.

### GetShutdownScheduledOk

`func (o *RemoveBrokerTaskData) GetShutdownScheduledOk() (*bool, bool)`

GetShutdownScheduledOk returns a tuple with the ShutdownScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownScheduled

`func (o *RemoveBrokerTaskData) SetShutdownScheduled(v bool)`

SetShutdownScheduled sets ShutdownScheduled field to given value.


### GetBrokerReplicaExclusionStatus

`func (o *RemoveBrokerTaskData) GetBrokerReplicaExclusionStatus() string`

GetBrokerReplicaExclusionStatus returns the BrokerReplicaExclusionStatus field if non-nil, zero value otherwise.

### GetBrokerReplicaExclusionStatusOk

`func (o *RemoveBrokerTaskData) GetBrokerReplicaExclusionStatusOk() (*string, bool)`

GetBrokerReplicaExclusionStatusOk returns a tuple with the BrokerReplicaExclusionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerReplicaExclusionStatus

`func (o *RemoveBrokerTaskData) SetBrokerReplicaExclusionStatus(v string)`

SetBrokerReplicaExclusionStatus sets BrokerReplicaExclusionStatus field to given value.


### GetPartitionReassignmentStatus

`func (o *RemoveBrokerTaskData) GetPartitionReassignmentStatus() string`

GetPartitionReassignmentStatus returns the PartitionReassignmentStatus field if non-nil, zero value otherwise.

### GetPartitionReassignmentStatusOk

`func (o *RemoveBrokerTaskData) GetPartitionReassignmentStatusOk() (*string, bool)`

GetPartitionReassignmentStatusOk returns a tuple with the PartitionReassignmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionReassignmentStatus

`func (o *RemoveBrokerTaskData) SetPartitionReassignmentStatus(v string)`

SetPartitionReassignmentStatus sets PartitionReassignmentStatus field to given value.


### GetBrokerShutdownStatus

`func (o *RemoveBrokerTaskData) GetBrokerShutdownStatus() string`

GetBrokerShutdownStatus returns the BrokerShutdownStatus field if non-nil, zero value otherwise.

### GetBrokerShutdownStatusOk

`func (o *RemoveBrokerTaskData) GetBrokerShutdownStatusOk() (*string, bool)`

GetBrokerShutdownStatusOk returns a tuple with the BrokerShutdownStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerShutdownStatus

`func (o *RemoveBrokerTaskData) SetBrokerShutdownStatus(v string)`

SetBrokerShutdownStatus sets BrokerShutdownStatus field to given value.


### GetErrorCode

`func (o *RemoveBrokerTaskData) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *RemoveBrokerTaskData) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *RemoveBrokerTaskData) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *RemoveBrokerTaskData) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *RemoveBrokerTaskData) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *RemoveBrokerTaskData) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorMessage

`func (o *RemoveBrokerTaskData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *RemoveBrokerTaskData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *RemoveBrokerTaskData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *RemoveBrokerTaskData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *RemoveBrokerTaskData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *RemoveBrokerTaskData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetBroker

`func (o *RemoveBrokerTaskData) GetBroker() Relationship`

GetBroker returns the Broker field if non-nil, zero value otherwise.

### GetBrokerOk

`func (o *RemoveBrokerTaskData) GetBrokerOk() (*Relationship, bool)`

GetBrokerOk returns a tuple with the Broker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroker

`func (o *RemoveBrokerTaskData) SetBroker(v Relationship)`

SetBroker sets Broker field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


