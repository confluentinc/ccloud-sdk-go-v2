# NotificationsV1UpdateUserNotificationsReadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Read** | **bool** | The target read state to apply to all notifications matching the filter query parameters. &#x60;true&#x60; marks them as read; &#x60;false&#x60; marks them as unread.  | 

## Methods

### NewNotificationsV1UpdateUserNotificationsReadRequest

`func NewNotificationsV1UpdateUserNotificationsReadRequest(read bool, ) *NotificationsV1UpdateUserNotificationsReadRequest`

NewNotificationsV1UpdateUserNotificationsReadRequest instantiates a new NotificationsV1UpdateUserNotificationsReadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsV1UpdateUserNotificationsReadRequestWithDefaults

`func NewNotificationsV1UpdateUserNotificationsReadRequestWithDefaults() *NotificationsV1UpdateUserNotificationsReadRequest`

NewNotificationsV1UpdateUserNotificationsReadRequestWithDefaults instantiates a new NotificationsV1UpdateUserNotificationsReadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRead

`func (o *NotificationsV1UpdateUserNotificationsReadRequest) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *NotificationsV1UpdateUserNotificationsReadRequest) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *NotificationsV1UpdateUserNotificationsReadRequest) SetRead(v bool)`

SetRead sets Read field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


