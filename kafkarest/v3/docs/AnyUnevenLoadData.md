# AnyUnevenLoadData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**ClusterId** | **string** |  | 
**Status** | **string** |  | 
**PreviousStatus** | **string** |  | 
**StatusUpdatedAt** | **time.Time** | The date and time at which this task was created. | [readonly] 
**PreviousStatusUpdatedAt** | **time.Time** | The date and time at which this task was created. | [readonly] 
**ErrorCode** | Pointer to **NullableInt32** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**BrokerTasks** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewAnyUnevenLoadData

`func NewAnyUnevenLoadData(kind string, metadata ResourceMetadata, clusterId string, status string, previousStatus string, statusUpdatedAt time.Time, previousStatusUpdatedAt time.Time, brokerTasks Relationship, ) *AnyUnevenLoadData`

NewAnyUnevenLoadData instantiates a new AnyUnevenLoadData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnyUnevenLoadDataWithDefaults

`func NewAnyUnevenLoadDataWithDefaults() *AnyUnevenLoadData`

NewAnyUnevenLoadDataWithDefaults instantiates a new AnyUnevenLoadData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *AnyUnevenLoadData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AnyUnevenLoadData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AnyUnevenLoadData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *AnyUnevenLoadData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AnyUnevenLoadData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AnyUnevenLoadData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *AnyUnevenLoadData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *AnyUnevenLoadData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *AnyUnevenLoadData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetStatus

`func (o *AnyUnevenLoadData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AnyUnevenLoadData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AnyUnevenLoadData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPreviousStatus

`func (o *AnyUnevenLoadData) GetPreviousStatus() string`

GetPreviousStatus returns the PreviousStatus field if non-nil, zero value otherwise.

### GetPreviousStatusOk

`func (o *AnyUnevenLoadData) GetPreviousStatusOk() (*string, bool)`

GetPreviousStatusOk returns a tuple with the PreviousStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousStatus

`func (o *AnyUnevenLoadData) SetPreviousStatus(v string)`

SetPreviousStatus sets PreviousStatus field to given value.


### GetStatusUpdatedAt

`func (o *AnyUnevenLoadData) GetStatusUpdatedAt() time.Time`

GetStatusUpdatedAt returns the StatusUpdatedAt field if non-nil, zero value otherwise.

### GetStatusUpdatedAtOk

`func (o *AnyUnevenLoadData) GetStatusUpdatedAtOk() (*time.Time, bool)`

GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUpdatedAt

`func (o *AnyUnevenLoadData) SetStatusUpdatedAt(v time.Time)`

SetStatusUpdatedAt sets StatusUpdatedAt field to given value.


### GetPreviousStatusUpdatedAt

`func (o *AnyUnevenLoadData) GetPreviousStatusUpdatedAt() time.Time`

GetPreviousStatusUpdatedAt returns the PreviousStatusUpdatedAt field if non-nil, zero value otherwise.

### GetPreviousStatusUpdatedAtOk

`func (o *AnyUnevenLoadData) GetPreviousStatusUpdatedAtOk() (*time.Time, bool)`

GetPreviousStatusUpdatedAtOk returns a tuple with the PreviousStatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousStatusUpdatedAt

`func (o *AnyUnevenLoadData) SetPreviousStatusUpdatedAt(v time.Time)`

SetPreviousStatusUpdatedAt sets PreviousStatusUpdatedAt field to given value.


### GetErrorCode

`func (o *AnyUnevenLoadData) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *AnyUnevenLoadData) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *AnyUnevenLoadData) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *AnyUnevenLoadData) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *AnyUnevenLoadData) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *AnyUnevenLoadData) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorMessage

`func (o *AnyUnevenLoadData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AnyUnevenLoadData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AnyUnevenLoadData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AnyUnevenLoadData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *AnyUnevenLoadData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *AnyUnevenLoadData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetBrokerTasks

`func (o *AnyUnevenLoadData) GetBrokerTasks() Relationship`

GetBrokerTasks returns the BrokerTasks field if non-nil, zero value otherwise.

### GetBrokerTasksOk

`func (o *AnyUnevenLoadData) GetBrokerTasksOk() (*Relationship, bool)`

GetBrokerTasksOk returns a tuple with the BrokerTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerTasks

`func (o *AnyUnevenLoadData) SetBrokerTasks(v Relationship)`

SetBrokerTasks sets BrokerTasks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


