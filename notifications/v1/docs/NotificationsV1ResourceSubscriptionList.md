# NotificationsV1ResourceSubscriptionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]NotificationsV1ResourceSubscription**](NotificationsV1ResourceSubscription.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewNotificationsV1ResourceSubscriptionList

`func NewNotificationsV1ResourceSubscriptionList(apiVersion string, kind string, metadata ListMeta, data []NotificationsV1ResourceSubscription, ) *NotificationsV1ResourceSubscriptionList`

NewNotificationsV1ResourceSubscriptionList instantiates a new NotificationsV1ResourceSubscriptionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsV1ResourceSubscriptionListWithDefaults

`func NewNotificationsV1ResourceSubscriptionListWithDefaults() *NotificationsV1ResourceSubscriptionList`

NewNotificationsV1ResourceSubscriptionListWithDefaults instantiates a new NotificationsV1ResourceSubscriptionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NotificationsV1ResourceSubscriptionList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NotificationsV1ResourceSubscriptionList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NotificationsV1ResourceSubscriptionList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *NotificationsV1ResourceSubscriptionList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NotificationsV1ResourceSubscriptionList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NotificationsV1ResourceSubscriptionList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *NotificationsV1ResourceSubscriptionList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NotificationsV1ResourceSubscriptionList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NotificationsV1ResourceSubscriptionList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *NotificationsV1ResourceSubscriptionList) GetData() []NotificationsV1ResourceSubscription`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NotificationsV1ResourceSubscriptionList) GetDataOk() (*[]NotificationsV1ResourceSubscription, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NotificationsV1ResourceSubscriptionList) SetData(v []NotificationsV1ResourceSubscription)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


