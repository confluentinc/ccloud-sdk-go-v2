# NetworkingV1TransitGatewayAttachmentList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]NetworkingV1TransitGatewayAttachment**](NetworkingV1TransitGatewayAttachment.md) |  | 

## Methods

### NewNetworkingV1TransitGatewayAttachmentList

`func NewNetworkingV1TransitGatewayAttachmentList(apiVersion string, kind string, metadata ListMeta, data []NetworkingV1TransitGatewayAttachment, ) *NetworkingV1TransitGatewayAttachmentList`

NewNetworkingV1TransitGatewayAttachmentList instantiates a new NetworkingV1TransitGatewayAttachmentList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1TransitGatewayAttachmentListWithDefaults

`func NewNetworkingV1TransitGatewayAttachmentListWithDefaults() *NetworkingV1TransitGatewayAttachmentList`

NewNetworkingV1TransitGatewayAttachmentListWithDefaults instantiates a new NetworkingV1TransitGatewayAttachmentList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkingV1TransitGatewayAttachmentList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkingV1TransitGatewayAttachmentList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkingV1TransitGatewayAttachmentList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *NetworkingV1TransitGatewayAttachmentList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1TransitGatewayAttachmentList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1TransitGatewayAttachmentList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *NetworkingV1TransitGatewayAttachmentList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkingV1TransitGatewayAttachmentList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkingV1TransitGatewayAttachmentList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *NetworkingV1TransitGatewayAttachmentList) GetData() []NetworkingV1TransitGatewayAttachment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NetworkingV1TransitGatewayAttachmentList) GetDataOk() (*[]NetworkingV1TransitGatewayAttachment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NetworkingV1TransitGatewayAttachmentList) SetData(v []NetworkingV1TransitGatewayAttachment)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


