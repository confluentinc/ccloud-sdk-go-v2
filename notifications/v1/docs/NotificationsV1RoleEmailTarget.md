# NotificationsV1RoleEmailTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Email Integration type for Role | 
**RoleName** | **string** | name of the role | 

## Methods

### NewNotificationsV1RoleEmailTarget

`func NewNotificationsV1RoleEmailTarget(kind string, roleName string, ) *NotificationsV1RoleEmailTarget`

NewNotificationsV1RoleEmailTarget instantiates a new NotificationsV1RoleEmailTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsV1RoleEmailTargetWithDefaults

`func NewNotificationsV1RoleEmailTargetWithDefaults() *NotificationsV1RoleEmailTarget`

NewNotificationsV1RoleEmailTargetWithDefaults instantiates a new NotificationsV1RoleEmailTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NotificationsV1RoleEmailTarget) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NotificationsV1RoleEmailTarget) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NotificationsV1RoleEmailTarget) SetKind(v string)`

SetKind sets Kind field to given value.


### GetRoleName

`func (o *NotificationsV1RoleEmailTarget) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *NotificationsV1RoleEmailTarget) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *NotificationsV1RoleEmailTarget) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.



### AsNotificationsV1Target

`func (s *NotificationsV1RoleEmailTarget) AsNotificationsV1Target() NotificationsV1Target`

Convenience method to wrap this instance of NotificationsV1RoleEmailTarget in NotificationsV1Target

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


