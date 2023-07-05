# BrokerTaskData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**ClusterId** | **string** |  | 
**BrokerId** | **int32** |  | 
**TaskType** | [**BrokerTaskType**](BrokerTaskType.md) |  | 
**TaskStatus** | **string** |  | 
**ShutdownScheduled** | Pointer to **NullableBool** |  | [optional] 
**SubTaskStatuses** | **map[string]string** |  | 
**CreatedAt** | **time.Time** | The date and time at which this task was created. | [readonly] 
**UpdatedAt** | **time.Time** | The date and time at which this task was last updated. | [readonly] 
**ErrorCode** | Pointer to **NullableInt32** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**Broker** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewBrokerTaskData

`func NewBrokerTaskData(kind string, metadata ResourceMetadata, clusterId string, brokerId int32, taskType BrokerTaskType, taskStatus string, subTaskStatuses map[string]string, createdAt time.Time, updatedAt time.Time, broker Relationship, ) *BrokerTaskData`

NewBrokerTaskData instantiates a new BrokerTaskData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrokerTaskDataWithDefaults

`func NewBrokerTaskDataWithDefaults() *BrokerTaskData`

NewBrokerTaskDataWithDefaults instantiates a new BrokerTaskData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *BrokerTaskData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BrokerTaskData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BrokerTaskData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *BrokerTaskData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BrokerTaskData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BrokerTaskData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *BrokerTaskData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *BrokerTaskData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *BrokerTaskData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetBrokerId

`func (o *BrokerTaskData) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *BrokerTaskData) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *BrokerTaskData) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.


### GetTaskType

`func (o *BrokerTaskData) GetTaskType() BrokerTaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *BrokerTaskData) GetTaskTypeOk() (*BrokerTaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *BrokerTaskData) SetTaskType(v BrokerTaskType)`

SetTaskType sets TaskType field to given value.


### GetTaskStatus

`func (o *BrokerTaskData) GetTaskStatus() string`

GetTaskStatus returns the TaskStatus field if non-nil, zero value otherwise.

### GetTaskStatusOk

`func (o *BrokerTaskData) GetTaskStatusOk() (*string, bool)`

GetTaskStatusOk returns a tuple with the TaskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStatus

`func (o *BrokerTaskData) SetTaskStatus(v string)`

SetTaskStatus sets TaskStatus field to given value.


### GetShutdownScheduled

`func (o *BrokerTaskData) GetShutdownScheduled() bool`

GetShutdownScheduled returns the ShutdownScheduled field if non-nil, zero value otherwise.

### GetShutdownScheduledOk

`func (o *BrokerTaskData) GetShutdownScheduledOk() (*bool, bool)`

GetShutdownScheduledOk returns a tuple with the ShutdownScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownScheduled

`func (o *BrokerTaskData) SetShutdownScheduled(v bool)`

SetShutdownScheduled sets ShutdownScheduled field to given value.

### HasShutdownScheduled

`func (o *BrokerTaskData) HasShutdownScheduled() bool`

HasShutdownScheduled returns a boolean if a field has been set.

### SetShutdownScheduledNil

`func (o *BrokerTaskData) SetShutdownScheduledNil(b bool)`

 SetShutdownScheduledNil sets the value for ShutdownScheduled to be an explicit nil

### UnsetShutdownScheduled
`func (o *BrokerTaskData) UnsetShutdownScheduled()`

UnsetShutdownScheduled ensures that no value is present for ShutdownScheduled, not even an explicit nil
### GetSubTaskStatuses

`func (o *BrokerTaskData) GetSubTaskStatuses() map[string]string`

GetSubTaskStatuses returns the SubTaskStatuses field if non-nil, zero value otherwise.

### GetSubTaskStatusesOk

`func (o *BrokerTaskData) GetSubTaskStatusesOk() (*map[string]string, bool)`

GetSubTaskStatusesOk returns a tuple with the SubTaskStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTaskStatuses

`func (o *BrokerTaskData) SetSubTaskStatuses(v map[string]string)`

SetSubTaskStatuses sets SubTaskStatuses field to given value.


### GetCreatedAt

`func (o *BrokerTaskData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BrokerTaskData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BrokerTaskData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BrokerTaskData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BrokerTaskData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BrokerTaskData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetErrorCode

`func (o *BrokerTaskData) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *BrokerTaskData) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *BrokerTaskData) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *BrokerTaskData) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *BrokerTaskData) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *BrokerTaskData) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorMessage

`func (o *BrokerTaskData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BrokerTaskData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BrokerTaskData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *BrokerTaskData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *BrokerTaskData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *BrokerTaskData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetBroker

`func (o *BrokerTaskData) GetBroker() Relationship`

GetBroker returns the Broker field if non-nil, zero value otherwise.

### GetBrokerOk

`func (o *BrokerTaskData) GetBrokerOk() (*Relationship, bool)`

GetBrokerOk returns a tuple with the Broker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroker

`func (o *BrokerTaskData) SetBroker(v Relationship)`

SetBroker sets Broker field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


