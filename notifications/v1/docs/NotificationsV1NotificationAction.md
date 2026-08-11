# NotificationsV1NotificationAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | Stable identifier for the action, suitable for analytics. Stable across notification deliveries that recommend the same action.  | 
**Url** | **string** | Confluent Cloud URL this action navigates to. | 
**Role** | **string** | Visual prominence of the action. &#x60;PRIMARY&#x60; is the recommended default action; &#x60;SECONDARY&#x60; is shown alongside as a less prominent option.  | 

## Methods

### NewNotificationsV1NotificationAction

`func NewNotificationsV1NotificationAction(identifier string, url string, role string, ) *NotificationsV1NotificationAction`

NewNotificationsV1NotificationAction instantiates a new NotificationsV1NotificationAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsV1NotificationActionWithDefaults

`func NewNotificationsV1NotificationActionWithDefaults() *NotificationsV1NotificationAction`

NewNotificationsV1NotificationActionWithDefaults instantiates a new NotificationsV1NotificationAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *NotificationsV1NotificationAction) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *NotificationsV1NotificationAction) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *NotificationsV1NotificationAction) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetUrl

`func (o *NotificationsV1NotificationAction) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationsV1NotificationAction) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationsV1NotificationAction) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRole

`func (o *NotificationsV1NotificationAction) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *NotificationsV1NotificationAction) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *NotificationsV1NotificationAction) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


