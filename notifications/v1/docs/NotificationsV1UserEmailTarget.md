# NotificationsV1UserEmailTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Email Integration type for User | 
**User** | [**GlobalObjectReference**](GlobalObjectReference.md) | Reference to the user | 

## Methods

### NewNotificationsV1UserEmailTarget

`func NewNotificationsV1UserEmailTarget(kind string, user GlobalObjectReference, ) *NotificationsV1UserEmailTarget`

NewNotificationsV1UserEmailTarget instantiates a new NotificationsV1UserEmailTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsV1UserEmailTargetWithDefaults

`func NewNotificationsV1UserEmailTargetWithDefaults() *NotificationsV1UserEmailTarget`

NewNotificationsV1UserEmailTargetWithDefaults instantiates a new NotificationsV1UserEmailTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NotificationsV1UserEmailTarget) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NotificationsV1UserEmailTarget) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NotificationsV1UserEmailTarget) SetKind(v string)`

SetKind sets Kind field to given value.


### GetUser

`func (o *NotificationsV1UserEmailTarget) GetUser() GlobalObjectReference`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *NotificationsV1UserEmailTarget) GetUserOk() (*GlobalObjectReference, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *NotificationsV1UserEmailTarget) SetUser(v GlobalObjectReference)`

SetUser sets User field to given value.



### AsNotificationsV1Target

`func (s *NotificationsV1UserEmailTarget) AsNotificationsV1Target() NotificationsV1Target`

Convenience method to wrap this instance of NotificationsV1UserEmailTarget in NotificationsV1Target

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


