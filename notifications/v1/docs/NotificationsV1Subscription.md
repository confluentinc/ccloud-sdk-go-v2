# NotificationsV1Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**CurrentState** | Pointer to **string** | Denotes the state of the subscription. When the subscription is ENABLED, the user will receive notification on the configured Integrations. If the subscription is DISABLED, the user will not recieve any notification for the configured notification type. Note that, you cannot disable a subscription for &#x60;REQUIRED&#x60; notification type.  | [optional] 
**NotificationType** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The type of notification to subscribe to. | [optional] 
**Integrations** | Pointer to [**[]GlobalObjectReference**](GlobalObjectReference.md) | Integrations to which notifications are to be sent. | [optional] 

## Methods

### NewNotificationsV1Subscription

`func NewNotificationsV1Subscription() *NotificationsV1Subscription`

NewNotificationsV1Subscription instantiates a new NotificationsV1Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsV1SubscriptionWithDefaults

`func NewNotificationsV1SubscriptionWithDefaults() *NotificationsV1Subscription`

NewNotificationsV1SubscriptionWithDefaults instantiates a new NotificationsV1Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NotificationsV1Subscription) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NotificationsV1Subscription) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NotificationsV1Subscription) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NotificationsV1Subscription) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *NotificationsV1Subscription) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NotificationsV1Subscription) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NotificationsV1Subscription) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NotificationsV1Subscription) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *NotificationsV1Subscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationsV1Subscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationsV1Subscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationsV1Subscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *NotificationsV1Subscription) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NotificationsV1Subscription) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NotificationsV1Subscription) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NotificationsV1Subscription) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCurrentState

`func (o *NotificationsV1Subscription) GetCurrentState() string`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *NotificationsV1Subscription) GetCurrentStateOk() (*string, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *NotificationsV1Subscription) SetCurrentState(v string)`

SetCurrentState sets CurrentState field to given value.

### HasCurrentState

`func (o *NotificationsV1Subscription) HasCurrentState() bool`

HasCurrentState returns a boolean if a field has been set.

### GetNotificationType

`func (o *NotificationsV1Subscription) GetNotificationType() GlobalObjectReference`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *NotificationsV1Subscription) GetNotificationTypeOk() (*GlobalObjectReference, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *NotificationsV1Subscription) SetNotificationType(v GlobalObjectReference)`

SetNotificationType sets NotificationType field to given value.

### HasNotificationType

`func (o *NotificationsV1Subscription) HasNotificationType() bool`

HasNotificationType returns a boolean if a field has been set.

### GetIntegrations

`func (o *NotificationsV1Subscription) GetIntegrations() []GlobalObjectReference`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *NotificationsV1Subscription) GetIntegrationsOk() (*[]GlobalObjectReference, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *NotificationsV1Subscription) SetIntegrations(v []GlobalObjectReference)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *NotificationsV1Subscription) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


