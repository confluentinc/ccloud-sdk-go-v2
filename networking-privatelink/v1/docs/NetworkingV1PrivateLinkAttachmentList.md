# NetworkingV1PrivateLinkAttachmentList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]NetworkingV1PrivateLinkAttachment**](NetworkingV1PrivateLinkAttachment.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewNetworkingV1PrivateLinkAttachmentList

`func NewNetworkingV1PrivateLinkAttachmentList(apiVersion string, kind string, metadata ListMeta, data []NetworkingV1PrivateLinkAttachment, ) *NetworkingV1PrivateLinkAttachmentList`

NewNetworkingV1PrivateLinkAttachmentList instantiates a new NetworkingV1PrivateLinkAttachmentList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1PrivateLinkAttachmentListWithDefaults

`func NewNetworkingV1PrivateLinkAttachmentListWithDefaults() *NetworkingV1PrivateLinkAttachmentList`

NewNetworkingV1PrivateLinkAttachmentListWithDefaults instantiates a new NetworkingV1PrivateLinkAttachmentList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkingV1PrivateLinkAttachmentList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkingV1PrivateLinkAttachmentList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkingV1PrivateLinkAttachmentList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *NetworkingV1PrivateLinkAttachmentList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1PrivateLinkAttachmentList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1PrivateLinkAttachmentList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *NetworkingV1PrivateLinkAttachmentList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkingV1PrivateLinkAttachmentList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkingV1PrivateLinkAttachmentList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *NetworkingV1PrivateLinkAttachmentList) GetData() []NetworkingV1PrivateLinkAttachment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NetworkingV1PrivateLinkAttachmentList) GetDataOk() (*[]NetworkingV1PrivateLinkAttachment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NetworkingV1PrivateLinkAttachmentList) SetData(v []NetworkingV1PrivateLinkAttachment)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


