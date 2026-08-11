# NotificationsV1ResourceSubscriptionUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Resource** | Pointer to **string** | Denotes the Confluent Cloud resource definition. | [optional] 
**ResourceType** | Pointer to **string** | Denotes the Confluent Cloud resource type. | [optional] 
**CurrentState** | Pointer to **string** | Denotes the state of the resource preference. When the resource preference is ENABLED, the user will receive notifications for the Confluent Cloud resource. If the resource preference is DISABLED, the user will not receive any notification for the resource. Note that, you will still receive notifications for &#x60;REQUIRED&#x60; notification type even when it is DISABLED.  | [optional] 
**Integrations** | Pointer to [**[]GlobalObjectReference**](GlobalObjectReference.md) | Integrations to which notifications are to be sent. | [optional] 

## Methods

### NewNotificationsV1ResourceSubscriptionUpdate

`func NewNotificationsV1ResourceSubscriptionUpdate() *NotificationsV1ResourceSubscriptionUpdate`

NewNotificationsV1ResourceSubscriptionUpdate instantiates a new NotificationsV1ResourceSubscriptionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsV1ResourceSubscriptionUpdateWithDefaults

`func NewNotificationsV1ResourceSubscriptionUpdateWithDefaults() *NotificationsV1ResourceSubscriptionUpdate`

NewNotificationsV1ResourceSubscriptionUpdateWithDefaults instantiates a new NotificationsV1ResourceSubscriptionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NotificationsV1ResourceSubscriptionUpdate) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NotificationsV1ResourceSubscriptionUpdate) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NotificationsV1ResourceSubscriptionUpdate) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NotificationsV1ResourceSubscriptionUpdate) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *NotificationsV1ResourceSubscriptionUpdate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NotificationsV1ResourceSubscriptionUpdate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NotificationsV1ResourceSubscriptionUpdate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NotificationsV1ResourceSubscriptionUpdate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *NotificationsV1ResourceSubscriptionUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationsV1ResourceSubscriptionUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationsV1ResourceSubscriptionUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationsV1ResourceSubscriptionUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *NotificationsV1ResourceSubscriptionUpdate) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NotificationsV1ResourceSubscriptionUpdate) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NotificationsV1ResourceSubscriptionUpdate) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NotificationsV1ResourceSubscriptionUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResource

`func (o *NotificationsV1ResourceSubscriptionUpdate) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *NotificationsV1ResourceSubscriptionUpdate) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *NotificationsV1ResourceSubscriptionUpdate) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *NotificationsV1ResourceSubscriptionUpdate) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetResourceType

`func (o *NotificationsV1ResourceSubscriptionUpdate) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *NotificationsV1ResourceSubscriptionUpdate) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *NotificationsV1ResourceSubscriptionUpdate) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *NotificationsV1ResourceSubscriptionUpdate) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetCurrentState

`func (o *NotificationsV1ResourceSubscriptionUpdate) GetCurrentState() string`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *NotificationsV1ResourceSubscriptionUpdate) GetCurrentStateOk() (*string, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *NotificationsV1ResourceSubscriptionUpdate) SetCurrentState(v string)`

SetCurrentState sets CurrentState field to given value.

### HasCurrentState

`func (o *NotificationsV1ResourceSubscriptionUpdate) HasCurrentState() bool`

HasCurrentState returns a boolean if a field has been set.

### GetIntegrations

`func (o *NotificationsV1ResourceSubscriptionUpdate) GetIntegrations() []GlobalObjectReference`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *NotificationsV1ResourceSubscriptionUpdate) GetIntegrationsOk() (*[]GlobalObjectReference, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *NotificationsV1ResourceSubscriptionUpdate) SetIntegrations(v []GlobalObjectReference)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *NotificationsV1ResourceSubscriptionUpdate) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


