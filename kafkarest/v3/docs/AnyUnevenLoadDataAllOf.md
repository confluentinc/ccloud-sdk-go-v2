# AnyUnevenLoadDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**Status** | **string** |  | 
**PreviousStatus** | **string** |  | 
**StatusUpdatedAt** | **time.Time** | The date and time at which this task was created. | [readonly] 
**PreviousStatusUpdatedAt** | **time.Time** | The date and time at which this task was created. | [readonly] 
**ErrorCode** | Pointer to **NullableInt32** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**BrokerTasks** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewAnyUnevenLoadDataAllOf

`func NewAnyUnevenLoadDataAllOf(clusterId string, status string, previousStatus string, statusUpdatedAt time.Time, previousStatusUpdatedAt time.Time, brokerTasks Relationship, ) *AnyUnevenLoadDataAllOf`

NewAnyUnevenLoadDataAllOf instantiates a new AnyUnevenLoadDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnyUnevenLoadDataAllOfWithDefaults

`func NewAnyUnevenLoadDataAllOfWithDefaults() *AnyUnevenLoadDataAllOf`

NewAnyUnevenLoadDataAllOfWithDefaults instantiates a new AnyUnevenLoadDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *AnyUnevenLoadDataAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *AnyUnevenLoadDataAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *AnyUnevenLoadDataAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetStatus

`func (o *AnyUnevenLoadDataAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AnyUnevenLoadDataAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AnyUnevenLoadDataAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPreviousStatus

`func (o *AnyUnevenLoadDataAllOf) GetPreviousStatus() string`

GetPreviousStatus returns the PreviousStatus field if non-nil, zero value otherwise.

### GetPreviousStatusOk

`func (o *AnyUnevenLoadDataAllOf) GetPreviousStatusOk() (*string, bool)`

GetPreviousStatusOk returns a tuple with the PreviousStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousStatus

`func (o *AnyUnevenLoadDataAllOf) SetPreviousStatus(v string)`

SetPreviousStatus sets PreviousStatus field to given value.


### GetStatusUpdatedAt

`func (o *AnyUnevenLoadDataAllOf) GetStatusUpdatedAt() time.Time`

GetStatusUpdatedAt returns the StatusUpdatedAt field if non-nil, zero value otherwise.

### GetStatusUpdatedAtOk

`func (o *AnyUnevenLoadDataAllOf) GetStatusUpdatedAtOk() (*time.Time, bool)`

GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUpdatedAt

`func (o *AnyUnevenLoadDataAllOf) SetStatusUpdatedAt(v time.Time)`

SetStatusUpdatedAt sets StatusUpdatedAt field to given value.


### GetPreviousStatusUpdatedAt

`func (o *AnyUnevenLoadDataAllOf) GetPreviousStatusUpdatedAt() time.Time`

GetPreviousStatusUpdatedAt returns the PreviousStatusUpdatedAt field if non-nil, zero value otherwise.

### GetPreviousStatusUpdatedAtOk

`func (o *AnyUnevenLoadDataAllOf) GetPreviousStatusUpdatedAtOk() (*time.Time, bool)`

GetPreviousStatusUpdatedAtOk returns a tuple with the PreviousStatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousStatusUpdatedAt

`func (o *AnyUnevenLoadDataAllOf) SetPreviousStatusUpdatedAt(v time.Time)`

SetPreviousStatusUpdatedAt sets PreviousStatusUpdatedAt field to given value.


### GetErrorCode

`func (o *AnyUnevenLoadDataAllOf) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *AnyUnevenLoadDataAllOf) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *AnyUnevenLoadDataAllOf) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *AnyUnevenLoadDataAllOf) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *AnyUnevenLoadDataAllOf) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *AnyUnevenLoadDataAllOf) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorMessage

`func (o *AnyUnevenLoadDataAllOf) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AnyUnevenLoadDataAllOf) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AnyUnevenLoadDataAllOf) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AnyUnevenLoadDataAllOf) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *AnyUnevenLoadDataAllOf) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *AnyUnevenLoadDataAllOf) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetBrokerTasks

`func (o *AnyUnevenLoadDataAllOf) GetBrokerTasks() Relationship`

GetBrokerTasks returns the BrokerTasks field if non-nil, zero value otherwise.

### GetBrokerTasksOk

`func (o *AnyUnevenLoadDataAllOf) GetBrokerTasksOk() (*Relationship, bool)`

GetBrokerTasksOk returns a tuple with the BrokerTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerTasks

`func (o *AnyUnevenLoadDataAllOf) SetBrokerTasks(v Relationship)`

SetBrokerTasks sets BrokerTasks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


