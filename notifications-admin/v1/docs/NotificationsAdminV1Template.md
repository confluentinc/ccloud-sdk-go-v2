# NotificationsAdminV1Template

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**TemplateSource** | Pointer to [**NotificationsAdminV1TemplateSource**](notifications-admin.v1.TemplateSource.md) | Contains message and title template for the notification with placeholders | [optional] 
**NotificationType** | Pointer to **string** | Type of the notification - Similar to the type in cloud event sent to the notification service while triggering notification.  | [optional] 
**IntegrationType** | Pointer to **string** | The type of the target system where alerts and notifications need to be delivered  | [optional] 

## Methods

### NewNotificationsAdminV1Template

`func NewNotificationsAdminV1Template() *NotificationsAdminV1Template`

NewNotificationsAdminV1Template instantiates a new NotificationsAdminV1Template object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsAdminV1TemplateWithDefaults

`func NewNotificationsAdminV1TemplateWithDefaults() *NotificationsAdminV1Template`

NewNotificationsAdminV1TemplateWithDefaults instantiates a new NotificationsAdminV1Template object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NotificationsAdminV1Template) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NotificationsAdminV1Template) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NotificationsAdminV1Template) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NotificationsAdminV1Template) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *NotificationsAdminV1Template) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NotificationsAdminV1Template) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NotificationsAdminV1Template) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NotificationsAdminV1Template) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *NotificationsAdminV1Template) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationsAdminV1Template) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationsAdminV1Template) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationsAdminV1Template) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *NotificationsAdminV1Template) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NotificationsAdminV1Template) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NotificationsAdminV1Template) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NotificationsAdminV1Template) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTemplateSource

`func (o *NotificationsAdminV1Template) GetTemplateSource() NotificationsAdminV1TemplateSource`

GetTemplateSource returns the TemplateSource field if non-nil, zero value otherwise.

### GetTemplateSourceOk

`func (o *NotificationsAdminV1Template) GetTemplateSourceOk() (*NotificationsAdminV1TemplateSource, bool)`

GetTemplateSourceOk returns a tuple with the TemplateSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSource

`func (o *NotificationsAdminV1Template) SetTemplateSource(v NotificationsAdminV1TemplateSource)`

SetTemplateSource sets TemplateSource field to given value.

### HasTemplateSource

`func (o *NotificationsAdminV1Template) HasTemplateSource() bool`

HasTemplateSource returns a boolean if a field has been set.

### GetNotificationType

`func (o *NotificationsAdminV1Template) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *NotificationsAdminV1Template) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *NotificationsAdminV1Template) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.

### HasNotificationType

`func (o *NotificationsAdminV1Template) HasNotificationType() bool`

HasNotificationType returns a boolean if a field has been set.

### GetIntegrationType

`func (o *NotificationsAdminV1Template) GetIntegrationType() string`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *NotificationsAdminV1Template) GetIntegrationTypeOk() (*string, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *NotificationsAdminV1Template) SetIntegrationType(v string)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *NotificationsAdminV1Template) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


