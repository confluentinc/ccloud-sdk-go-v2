# NotificationsV1UserNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Read** | Pointer to **bool** | Whether the notification has been read by the user. | [optional] 
**ReadAt** | Pointer to **NullableTime** | The time the notification was marked as read, or &#x60;null&#x60; if it is unread. | [optional] [readonly] 
**Severity** | Pointer to **string** | The severity level of the notification. - CRITICAL: a high impact notification which needs immediate attention. - WARN: a warning notification which can be addressed now or later. - INFO: an informational notification.  | [optional] [readonly] 
**ReceivedAt** | Pointer to **time.Time** | The time the underlying event was generated. | [optional] [readonly] 
**Resource** | Pointer to [**NotificationsV1ResourceSnapshot**](notifications.v1.ResourceSnapshot.md) | The Confluent Cloud resource this notification relates to, embedded as a point-in-time snapshot at delivery time. Values remain accurate even if the underlying resource is later renamed or deleted.  | [optional] [readonly] 
**Actions** | Pointer to [**[]NotificationsV1NotificationAction**](NotificationsV1NotificationAction.md) | Ordered list of user-facing actions associated with this notification. The first entry is the primary action (&#x60;role: PRIMARY&#x60;) and is always present; subsequent entries are secondary. Cardinality is open-ended — additional actions may be added over time without a breaking schema change.  | [optional] [readonly] 
**NotificationType** | Pointer to [**NotificationsV1NotificationType**](notifications.v1.NotificationType.md) | The notification type that triggered this notification, embedded as a point-in-time snapshot at delivery time so values remain accurate even if the underlying &#x60;NotificationType&#x60; is later modified.  | [optional] [readonly] 
**Integrations** | Pointer to [**[]NotificationsV1Integration**](NotificationsV1Integration.md) | The integrations this notification was delivered to. Each entry is a point-in-time snapshot of the integration at delivery time, so values remain accurate even if the underlying &#x60;Integration&#x60; is later modified or deleted. Populated on single-resource reads (&#x60;GET /user-notifications/{id}&#x60;); omitted from list responses.  | [optional] [readonly] 
**RecommendedActions** | Pointer to [**NotificationsV1RecommendedActions**](notifications.v1.RecommendedActions.md) | Versioned payload describing the recommended actions a user can take in response to this notification. The shape is stable per &#x60;version&#x60; and consumers should branch on &#x60;version&#x60; when deserializing. Populated on single-resource reads (&#x60;GET /user-notifications/{id}&#x60;); omitted from list responses.  | [optional] [readonly] 

## Methods

### NewNotificationsV1UserNotification

`func NewNotificationsV1UserNotification() *NotificationsV1UserNotification`

NewNotificationsV1UserNotification instantiates a new NotificationsV1UserNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsV1UserNotificationWithDefaults

`func NewNotificationsV1UserNotificationWithDefaults() *NotificationsV1UserNotification`

NewNotificationsV1UserNotificationWithDefaults instantiates a new NotificationsV1UserNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NotificationsV1UserNotification) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NotificationsV1UserNotification) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NotificationsV1UserNotification) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NotificationsV1UserNotification) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *NotificationsV1UserNotification) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NotificationsV1UserNotification) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NotificationsV1UserNotification) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NotificationsV1UserNotification) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *NotificationsV1UserNotification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationsV1UserNotification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationsV1UserNotification) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationsV1UserNotification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *NotificationsV1UserNotification) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NotificationsV1UserNotification) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NotificationsV1UserNotification) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NotificationsV1UserNotification) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRead

`func (o *NotificationsV1UserNotification) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *NotificationsV1UserNotification) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *NotificationsV1UserNotification) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *NotificationsV1UserNotification) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetReadAt

`func (o *NotificationsV1UserNotification) GetReadAt() time.Time`

GetReadAt returns the ReadAt field if non-nil, zero value otherwise.

### GetReadAtOk

`func (o *NotificationsV1UserNotification) GetReadAtOk() (*time.Time, bool)`

GetReadAtOk returns a tuple with the ReadAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAt

`func (o *NotificationsV1UserNotification) SetReadAt(v time.Time)`

SetReadAt sets ReadAt field to given value.

### HasReadAt

`func (o *NotificationsV1UserNotification) HasReadAt() bool`

HasReadAt returns a boolean if a field has been set.

### SetReadAtNil

`func (o *NotificationsV1UserNotification) SetReadAtNil(b bool)`

 SetReadAtNil sets the value for ReadAt to be an explicit nil

### UnsetReadAt
`func (o *NotificationsV1UserNotification) UnsetReadAt()`

UnsetReadAt ensures that no value is present for ReadAt, not even an explicit nil
### GetSeverity

`func (o *NotificationsV1UserNotification) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *NotificationsV1UserNotification) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *NotificationsV1UserNotification) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *NotificationsV1UserNotification) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetReceivedAt

`func (o *NotificationsV1UserNotification) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *NotificationsV1UserNotification) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *NotificationsV1UserNotification) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.

### HasReceivedAt

`func (o *NotificationsV1UserNotification) HasReceivedAt() bool`

HasReceivedAt returns a boolean if a field has been set.

### GetResource

`func (o *NotificationsV1UserNotification) GetResource() NotificationsV1ResourceSnapshot`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *NotificationsV1UserNotification) GetResourceOk() (*NotificationsV1ResourceSnapshot, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *NotificationsV1UserNotification) SetResource(v NotificationsV1ResourceSnapshot)`

SetResource sets Resource field to given value.

### HasResource

`func (o *NotificationsV1UserNotification) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetActions

`func (o *NotificationsV1UserNotification) GetActions() []NotificationsV1NotificationAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *NotificationsV1UserNotification) GetActionsOk() (*[]NotificationsV1NotificationAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *NotificationsV1UserNotification) SetActions(v []NotificationsV1NotificationAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *NotificationsV1UserNotification) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetNotificationType

`func (o *NotificationsV1UserNotification) GetNotificationType() NotificationsV1NotificationType`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *NotificationsV1UserNotification) GetNotificationTypeOk() (*NotificationsV1NotificationType, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *NotificationsV1UserNotification) SetNotificationType(v NotificationsV1NotificationType)`

SetNotificationType sets NotificationType field to given value.

### HasNotificationType

`func (o *NotificationsV1UserNotification) HasNotificationType() bool`

HasNotificationType returns a boolean if a field has been set.

### GetIntegrations

`func (o *NotificationsV1UserNotification) GetIntegrations() []NotificationsV1Integration`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *NotificationsV1UserNotification) GetIntegrationsOk() (*[]NotificationsV1Integration, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *NotificationsV1UserNotification) SetIntegrations(v []NotificationsV1Integration)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *NotificationsV1UserNotification) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### GetRecommendedActions

`func (o *NotificationsV1UserNotification) GetRecommendedActions() NotificationsV1RecommendedActions`

GetRecommendedActions returns the RecommendedActions field if non-nil, zero value otherwise.

### GetRecommendedActionsOk

`func (o *NotificationsV1UserNotification) GetRecommendedActionsOk() (*NotificationsV1RecommendedActions, bool)`

GetRecommendedActionsOk returns a tuple with the RecommendedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedActions

`func (o *NotificationsV1UserNotification) SetRecommendedActions(v NotificationsV1RecommendedActions)`

SetRecommendedActions sets RecommendedActions field to given value.

### HasRecommendedActions

`func (o *NotificationsV1UserNotification) HasRecommendedActions() bool`

HasRecommendedActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


