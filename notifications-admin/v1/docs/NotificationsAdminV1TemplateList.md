# NotificationsAdminV1TemplateList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]NotificationsAdminV1Template**](NotificationsAdminV1Template.md) |  | 

## Methods

### NewNotificationsAdminV1TemplateList

`func NewNotificationsAdminV1TemplateList(apiVersion string, kind string, metadata ListMeta, data []NotificationsAdminV1Template, ) *NotificationsAdminV1TemplateList`

NewNotificationsAdminV1TemplateList instantiates a new NotificationsAdminV1TemplateList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsAdminV1TemplateListWithDefaults

`func NewNotificationsAdminV1TemplateListWithDefaults() *NotificationsAdminV1TemplateList`

NewNotificationsAdminV1TemplateListWithDefaults instantiates a new NotificationsAdminV1TemplateList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NotificationsAdminV1TemplateList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NotificationsAdminV1TemplateList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NotificationsAdminV1TemplateList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *NotificationsAdminV1TemplateList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NotificationsAdminV1TemplateList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NotificationsAdminV1TemplateList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *NotificationsAdminV1TemplateList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NotificationsAdminV1TemplateList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NotificationsAdminV1TemplateList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *NotificationsAdminV1TemplateList) GetData() []NotificationsAdminV1Template`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NotificationsAdminV1TemplateList) GetDataOk() (*[]NotificationsAdminV1Template, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NotificationsAdminV1TemplateList) SetData(v []NotificationsAdminV1Template)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


