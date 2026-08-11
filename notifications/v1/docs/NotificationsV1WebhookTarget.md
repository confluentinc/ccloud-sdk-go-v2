# NotificationsV1WebhookTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Integration Type | 
**Url** | **string** | URL endpoint for the webhook | 

## Methods

### NewNotificationsV1WebhookTarget

`func NewNotificationsV1WebhookTarget(kind string, url string, ) *NotificationsV1WebhookTarget`

NewNotificationsV1WebhookTarget instantiates a new NotificationsV1WebhookTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsV1WebhookTargetWithDefaults

`func NewNotificationsV1WebhookTargetWithDefaults() *NotificationsV1WebhookTarget`

NewNotificationsV1WebhookTargetWithDefaults instantiates a new NotificationsV1WebhookTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NotificationsV1WebhookTarget) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NotificationsV1WebhookTarget) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NotificationsV1WebhookTarget) SetKind(v string)`

SetKind sets Kind field to given value.


### GetUrl

`func (o *NotificationsV1WebhookTarget) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationsV1WebhookTarget) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationsV1WebhookTarget) SetUrl(v string)`

SetUrl sets Url field to given value.



### AsNotificationsV1Target

`func (s *NotificationsV1WebhookTarget) AsNotificationsV1Target() NotificationsV1Target`

Convenience method to wrap this instance of NotificationsV1WebhookTarget in NotificationsV1Target

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


