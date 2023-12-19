# BrokerTaskDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewBrokerTaskDataAllOf

`func NewBrokerTaskDataAllOf(clusterId string, brokerId int32, taskType BrokerTaskType, taskStatus string, subTaskStatuses map[string]string, createdAt time.Time, updatedAt time.Time, broker Relationship, ) *BrokerTaskDataAllOf`

NewBrokerTaskDataAllOf instantiates a new BrokerTaskDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrokerTaskDataAllOfWithDefaults

`func NewBrokerTaskDataAllOfWithDefaults() *BrokerTaskDataAllOf`

NewBrokerTaskDataAllOfWithDefaults instantiates a new BrokerTaskDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *BrokerTaskDataAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *BrokerTaskDataAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *BrokerTaskDataAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetBrokerId

`func (o *BrokerTaskDataAllOf) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *BrokerTaskDataAllOf) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *BrokerTaskDataAllOf) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.


### GetTaskType

`func (o *BrokerTaskDataAllOf) GetTaskType() BrokerTaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *BrokerTaskDataAllOf) GetTaskTypeOk() (*BrokerTaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *BrokerTaskDataAllOf) SetTaskType(v BrokerTaskType)`

SetTaskType sets TaskType field to given value.


### GetTaskStatus

`func (o *BrokerTaskDataAllOf) GetTaskStatus() string`

GetTaskStatus returns the TaskStatus field if non-nil, zero value otherwise.

### GetTaskStatusOk

`func (o *BrokerTaskDataAllOf) GetTaskStatusOk() (*string, bool)`

GetTaskStatusOk returns a tuple with the TaskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStatus

`func (o *BrokerTaskDataAllOf) SetTaskStatus(v string)`

SetTaskStatus sets TaskStatus field to given value.


### GetShutdownScheduled

`func (o *BrokerTaskDataAllOf) GetShutdownScheduled() bool`

GetShutdownScheduled returns the ShutdownScheduled field if non-nil, zero value otherwise.

### GetShutdownScheduledOk

`func (o *BrokerTaskDataAllOf) GetShutdownScheduledOk() (*bool, bool)`

GetShutdownScheduledOk returns a tuple with the ShutdownScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownScheduled

`func (o *BrokerTaskDataAllOf) SetShutdownScheduled(v bool)`

SetShutdownScheduled sets ShutdownScheduled field to given value.

### HasShutdownScheduled

`func (o *BrokerTaskDataAllOf) HasShutdownScheduled() bool`

HasShutdownScheduled returns a boolean if a field has been set.

### SetShutdownScheduledNil

`func (o *BrokerTaskDataAllOf) SetShutdownScheduledNil(b bool)`

 SetShutdownScheduledNil sets the value for ShutdownScheduled to be an explicit nil

### UnsetShutdownScheduled
`func (o *BrokerTaskDataAllOf) UnsetShutdownScheduled()`

UnsetShutdownScheduled ensures that no value is present for ShutdownScheduled, not even an explicit nil
### GetSubTaskStatuses

`func (o *BrokerTaskDataAllOf) GetSubTaskStatuses() map[string]string`

GetSubTaskStatuses returns the SubTaskStatuses field if non-nil, zero value otherwise.

### GetSubTaskStatusesOk

`func (o *BrokerTaskDataAllOf) GetSubTaskStatusesOk() (*map[string]string, bool)`

GetSubTaskStatusesOk returns a tuple with the SubTaskStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTaskStatuses

`func (o *BrokerTaskDataAllOf) SetSubTaskStatuses(v map[string]string)`

SetSubTaskStatuses sets SubTaskStatuses field to given value.


### GetCreatedAt

`func (o *BrokerTaskDataAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BrokerTaskDataAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BrokerTaskDataAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BrokerTaskDataAllOf) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BrokerTaskDataAllOf) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BrokerTaskDataAllOf) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetErrorCode

`func (o *BrokerTaskDataAllOf) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *BrokerTaskDataAllOf) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *BrokerTaskDataAllOf) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *BrokerTaskDataAllOf) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *BrokerTaskDataAllOf) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *BrokerTaskDataAllOf) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorMessage

`func (o *BrokerTaskDataAllOf) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BrokerTaskDataAllOf) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BrokerTaskDataAllOf) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *BrokerTaskDataAllOf) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *BrokerTaskDataAllOf) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *BrokerTaskDataAllOf) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetBroker

`func (o *BrokerTaskDataAllOf) GetBroker() Relationship`

GetBroker returns the Broker field if non-nil, zero value otherwise.

### GetBrokerOk

`func (o *BrokerTaskDataAllOf) GetBrokerOk() (*Relationship, bool)`

GetBrokerOk returns a tuple with the Broker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroker

`func (o *BrokerTaskDataAllOf) SetBroker(v Relationship)`

SetBroker sets Broker field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


