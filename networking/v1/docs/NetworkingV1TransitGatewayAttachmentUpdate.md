# NetworkingV1TransitGatewayAttachmentUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**NetworkingV1TransitGatewayAttachmentSpecUpdate**](NetworkingV1TransitGatewayAttachmentSpecUpdate.md) |  | [optional] 
**Status** | Pointer to [**NetworkingV1TransitGatewayAttachmentStatus**](NetworkingV1TransitGatewayAttachmentStatus.md) |  | [optional] 

## Methods

### NewNetworkingV1TransitGatewayAttachmentUpdate

`func NewNetworkingV1TransitGatewayAttachmentUpdate() *NetworkingV1TransitGatewayAttachmentUpdate`

NewNetworkingV1TransitGatewayAttachmentUpdate instantiates a new NetworkingV1TransitGatewayAttachmentUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1TransitGatewayAttachmentUpdateWithDefaults

`func NewNetworkingV1TransitGatewayAttachmentUpdateWithDefaults() *NetworkingV1TransitGatewayAttachmentUpdate`

NewNetworkingV1TransitGatewayAttachmentUpdateWithDefaults instantiates a new NetworkingV1TransitGatewayAttachmentUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkingV1TransitGatewayAttachmentUpdate) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkingV1TransitGatewayAttachmentUpdate) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkingV1TransitGatewayAttachmentUpdate) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkingV1TransitGatewayAttachmentUpdate) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *NetworkingV1TransitGatewayAttachmentUpdate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1TransitGatewayAttachmentUpdate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1TransitGatewayAttachmentUpdate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkingV1TransitGatewayAttachmentUpdate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *NetworkingV1TransitGatewayAttachmentUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkingV1TransitGatewayAttachmentUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkingV1TransitGatewayAttachmentUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkingV1TransitGatewayAttachmentUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *NetworkingV1TransitGatewayAttachmentUpdate) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkingV1TransitGatewayAttachmentUpdate) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkingV1TransitGatewayAttachmentUpdate) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NetworkingV1TransitGatewayAttachmentUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *NetworkingV1TransitGatewayAttachmentUpdate) GetSpec() NetworkingV1TransitGatewayAttachmentSpecUpdate`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NetworkingV1TransitGatewayAttachmentUpdate) GetSpecOk() (*NetworkingV1TransitGatewayAttachmentSpecUpdate, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NetworkingV1TransitGatewayAttachmentUpdate) SetSpec(v NetworkingV1TransitGatewayAttachmentSpecUpdate)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *NetworkingV1TransitGatewayAttachmentUpdate) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkingV1TransitGatewayAttachmentUpdate) GetStatus() NetworkingV1TransitGatewayAttachmentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkingV1TransitGatewayAttachmentUpdate) GetStatusOk() (*NetworkingV1TransitGatewayAttachmentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkingV1TransitGatewayAttachmentUpdate) SetStatus(v NetworkingV1TransitGatewayAttachmentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkingV1TransitGatewayAttachmentUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


