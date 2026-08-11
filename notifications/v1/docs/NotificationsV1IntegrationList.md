# NotificationsV1IntegrationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]NotificationsV1Integration**](NotificationsV1Integration.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewNotificationsV1IntegrationList

`func NewNotificationsV1IntegrationList(apiVersion string, kind string, metadata ListMeta, data []NotificationsV1Integration, ) *NotificationsV1IntegrationList`

NewNotificationsV1IntegrationList instantiates a new NotificationsV1IntegrationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsV1IntegrationListWithDefaults

`func NewNotificationsV1IntegrationListWithDefaults() *NotificationsV1IntegrationList`

NewNotificationsV1IntegrationListWithDefaults instantiates a new NotificationsV1IntegrationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NotificationsV1IntegrationList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NotificationsV1IntegrationList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NotificationsV1IntegrationList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *NotificationsV1IntegrationList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NotificationsV1IntegrationList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NotificationsV1IntegrationList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *NotificationsV1IntegrationList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NotificationsV1IntegrationList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NotificationsV1IntegrationList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *NotificationsV1IntegrationList) GetData() []NotificationsV1Integration`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NotificationsV1IntegrationList) GetDataOk() (*[]NotificationsV1Integration, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NotificationsV1IntegrationList) SetData(v []NotificationsV1Integration)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


