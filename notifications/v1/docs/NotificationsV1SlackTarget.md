# NotificationsV1SlackTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Integration Type | 
**WebhookUrl** | **string** | Slack Webhook URL for the particular Slack channel | 

## Methods

### NewNotificationsV1SlackTarget

`func NewNotificationsV1SlackTarget(kind string, webhookUrl string, ) *NotificationsV1SlackTarget`

NewNotificationsV1SlackTarget instantiates a new NotificationsV1SlackTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsV1SlackTargetWithDefaults

`func NewNotificationsV1SlackTargetWithDefaults() *NotificationsV1SlackTarget`

NewNotificationsV1SlackTargetWithDefaults instantiates a new NotificationsV1SlackTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NotificationsV1SlackTarget) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NotificationsV1SlackTarget) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NotificationsV1SlackTarget) SetKind(v string)`

SetKind sets Kind field to given value.


### GetWebhookUrl

`func (o *NotificationsV1SlackTarget) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *NotificationsV1SlackTarget) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *NotificationsV1SlackTarget) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.



### AsNotificationsV1Target

`func (s *NotificationsV1SlackTarget) AsNotificationsV1Target() NotificationsV1Target`

Convenience method to wrap this instance of NotificationsV1SlackTarget in NotificationsV1Target

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


