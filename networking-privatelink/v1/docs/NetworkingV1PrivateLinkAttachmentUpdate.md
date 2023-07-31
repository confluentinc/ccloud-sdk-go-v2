# NetworkingV1PrivateLinkAttachmentUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**NetworkingV1PrivateLinkAttachmentSpecUpdate**](NetworkingV1PrivateLinkAttachmentSpecUpdate.md) |  | [optional] 
**Status** | Pointer to [**NetworkingV1PrivateLinkAttachmentStatus**](NetworkingV1PrivateLinkAttachmentStatus.md) |  | [optional] 

## Methods

### NewNetworkingV1PrivateLinkAttachmentUpdate

`func NewNetworkingV1PrivateLinkAttachmentUpdate() *NetworkingV1PrivateLinkAttachmentUpdate`

NewNetworkingV1PrivateLinkAttachmentUpdate instantiates a new NetworkingV1PrivateLinkAttachmentUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1PrivateLinkAttachmentUpdateWithDefaults

`func NewNetworkingV1PrivateLinkAttachmentUpdateWithDefaults() *NetworkingV1PrivateLinkAttachmentUpdate`

NewNetworkingV1PrivateLinkAttachmentUpdateWithDefaults instantiates a new NetworkingV1PrivateLinkAttachmentUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkingV1PrivateLinkAttachmentUpdate) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkingV1PrivateLinkAttachmentUpdate) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkingV1PrivateLinkAttachmentUpdate) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkingV1PrivateLinkAttachmentUpdate) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *NetworkingV1PrivateLinkAttachmentUpdate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1PrivateLinkAttachmentUpdate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1PrivateLinkAttachmentUpdate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkingV1PrivateLinkAttachmentUpdate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *NetworkingV1PrivateLinkAttachmentUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkingV1PrivateLinkAttachmentUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkingV1PrivateLinkAttachmentUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkingV1PrivateLinkAttachmentUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *NetworkingV1PrivateLinkAttachmentUpdate) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkingV1PrivateLinkAttachmentUpdate) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkingV1PrivateLinkAttachmentUpdate) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NetworkingV1PrivateLinkAttachmentUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *NetworkingV1PrivateLinkAttachmentUpdate) GetSpec() NetworkingV1PrivateLinkAttachmentSpecUpdate`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NetworkingV1PrivateLinkAttachmentUpdate) GetSpecOk() (*NetworkingV1PrivateLinkAttachmentSpecUpdate, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NetworkingV1PrivateLinkAttachmentUpdate) SetSpec(v NetworkingV1PrivateLinkAttachmentSpecUpdate)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *NetworkingV1PrivateLinkAttachmentUpdate) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkingV1PrivateLinkAttachmentUpdate) GetStatus() NetworkingV1PrivateLinkAttachmentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkingV1PrivateLinkAttachmentUpdate) GetStatusOk() (*NetworkingV1PrivateLinkAttachmentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkingV1PrivateLinkAttachmentUpdate) SetStatus(v NetworkingV1PrivateLinkAttachmentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkingV1PrivateLinkAttachmentUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


