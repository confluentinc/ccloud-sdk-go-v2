# NetworkingAdminV1TransitGatewayAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**NetworkingAdminV1TransitGatewayAttachmentSpec**](NetworkingAdminV1TransitGatewayAttachmentSpec.md) |  | [optional] 
**Status** | Pointer to [**NetworkingAdminV1TransitGatewayAttachmentStatus**](NetworkingAdminV1TransitGatewayAttachmentStatus.md) |  | [optional] 

## Methods

### NewNetworkingAdminV1TransitGatewayAttachment

`func NewNetworkingAdminV1TransitGatewayAttachment() *NetworkingAdminV1TransitGatewayAttachment`

NewNetworkingAdminV1TransitGatewayAttachment instantiates a new NetworkingAdminV1TransitGatewayAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingAdminV1TransitGatewayAttachmentWithDefaults

`func NewNetworkingAdminV1TransitGatewayAttachmentWithDefaults() *NetworkingAdminV1TransitGatewayAttachment`

NewNetworkingAdminV1TransitGatewayAttachmentWithDefaults instantiates a new NetworkingAdminV1TransitGatewayAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkingAdminV1TransitGatewayAttachment) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkingAdminV1TransitGatewayAttachment) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkingAdminV1TransitGatewayAttachment) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkingAdminV1TransitGatewayAttachment) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *NetworkingAdminV1TransitGatewayAttachment) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingAdminV1TransitGatewayAttachment) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingAdminV1TransitGatewayAttachment) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkingAdminV1TransitGatewayAttachment) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *NetworkingAdminV1TransitGatewayAttachment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkingAdminV1TransitGatewayAttachment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkingAdminV1TransitGatewayAttachment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkingAdminV1TransitGatewayAttachment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *NetworkingAdminV1TransitGatewayAttachment) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkingAdminV1TransitGatewayAttachment) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkingAdminV1TransitGatewayAttachment) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NetworkingAdminV1TransitGatewayAttachment) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *NetworkingAdminV1TransitGatewayAttachment) GetSpec() NetworkingAdminV1TransitGatewayAttachmentSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NetworkingAdminV1TransitGatewayAttachment) GetSpecOk() (*NetworkingAdminV1TransitGatewayAttachmentSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NetworkingAdminV1TransitGatewayAttachment) SetSpec(v NetworkingAdminV1TransitGatewayAttachmentSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *NetworkingAdminV1TransitGatewayAttachment) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkingAdminV1TransitGatewayAttachment) GetStatus() NetworkingAdminV1TransitGatewayAttachmentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkingAdminV1TransitGatewayAttachment) GetStatusOk() (*NetworkingAdminV1TransitGatewayAttachmentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkingAdminV1TransitGatewayAttachment) SetStatus(v NetworkingAdminV1TransitGatewayAttachmentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkingAdminV1TransitGatewayAttachment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


