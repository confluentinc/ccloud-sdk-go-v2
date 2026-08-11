# NotificationsV1NotificationType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**DisplayName** | Pointer to **string** | Human readable display name of the notification type  | [optional] 
**Category** | Pointer to **string** | Represents the group with which the notification is associated. Notifications are grouped under certain categories for better organization. - BILLING_LICENSING: All billing, payments or licensing related notifications are grouped here. - SECURITY: All Confluent Cloud and Platform security related notifications are grouped here. - SERVICE: All Confluent services (eg. Kafka, Schema Registry, Connect etc.) related notifications are   grouped here. - ACCOUNT: All Confluent account related notifications are grouped here. For example: Billing, payment or license related notifications are grouped in BILLING_LICENSING category.  | [optional] 
**Description** | Pointer to **string** | Human readable description of the notification type  | [optional] 
**SubscriptionPriority** | Pointer to **string** | Indicates whether the notification is auto-subscribed and if the user can opt-out. - REQUIRED: the user is auto-subscribed to this notification and can&#39;t opt-out. - RECOMMENDED: the user is auto-subscribed to this notification and can opt-out. - OPTIONAL: the user is not auto-subscribed to this notification but can explicitly subscribe to it.  | [optional] 
**IsIncludedInPlan** | Pointer to **bool** | Whether this notification is available to subscribe or not as per the user&#39;s current billing plan.  | [optional] 
**Severity** | Pointer to **string** | Severity indicates the impact of this notification. - CRITICAL: a high impact notification which needs immediate attention. - WARN: a warning notification which can be addressed now or later. - INFO: an informational notification.  | [optional] 

## Methods

### NewNotificationsV1NotificationType

`func NewNotificationsV1NotificationType() *NotificationsV1NotificationType`

NewNotificationsV1NotificationType instantiates a new NotificationsV1NotificationType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsV1NotificationTypeWithDefaults

`func NewNotificationsV1NotificationTypeWithDefaults() *NotificationsV1NotificationType`

NewNotificationsV1NotificationTypeWithDefaults instantiates a new NotificationsV1NotificationType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NotificationsV1NotificationType) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NotificationsV1NotificationType) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NotificationsV1NotificationType) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NotificationsV1NotificationType) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *NotificationsV1NotificationType) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NotificationsV1NotificationType) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NotificationsV1NotificationType) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NotificationsV1NotificationType) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *NotificationsV1NotificationType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationsV1NotificationType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationsV1NotificationType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationsV1NotificationType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *NotificationsV1NotificationType) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NotificationsV1NotificationType) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NotificationsV1NotificationType) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NotificationsV1NotificationType) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *NotificationsV1NotificationType) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NotificationsV1NotificationType) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NotificationsV1NotificationType) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NotificationsV1NotificationType) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCategory

`func (o *NotificationsV1NotificationType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *NotificationsV1NotificationType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *NotificationsV1NotificationType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *NotificationsV1NotificationType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *NotificationsV1NotificationType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotificationsV1NotificationType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotificationsV1NotificationType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotificationsV1NotificationType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubscriptionPriority

`func (o *NotificationsV1NotificationType) GetSubscriptionPriority() string`

GetSubscriptionPriority returns the SubscriptionPriority field if non-nil, zero value otherwise.

### GetSubscriptionPriorityOk

`func (o *NotificationsV1NotificationType) GetSubscriptionPriorityOk() (*string, bool)`

GetSubscriptionPriorityOk returns a tuple with the SubscriptionPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPriority

`func (o *NotificationsV1NotificationType) SetSubscriptionPriority(v string)`

SetSubscriptionPriority sets SubscriptionPriority field to given value.

### HasSubscriptionPriority

`func (o *NotificationsV1NotificationType) HasSubscriptionPriority() bool`

HasSubscriptionPriority returns a boolean if a field has been set.

### GetIsIncludedInPlan

`func (o *NotificationsV1NotificationType) GetIsIncludedInPlan() bool`

GetIsIncludedInPlan returns the IsIncludedInPlan field if non-nil, zero value otherwise.

### GetIsIncludedInPlanOk

`func (o *NotificationsV1NotificationType) GetIsIncludedInPlanOk() (*bool, bool)`

GetIsIncludedInPlanOk returns a tuple with the IsIncludedInPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncludedInPlan

`func (o *NotificationsV1NotificationType) SetIsIncludedInPlan(v bool)`

SetIsIncludedInPlan sets IsIncludedInPlan field to given value.

### HasIsIncludedInPlan

`func (o *NotificationsV1NotificationType) HasIsIncludedInPlan() bool`

HasIsIncludedInPlan returns a boolean if a field has been set.

### GetSeverity

`func (o *NotificationsV1NotificationType) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *NotificationsV1NotificationType) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *NotificationsV1NotificationType) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *NotificationsV1NotificationType) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


