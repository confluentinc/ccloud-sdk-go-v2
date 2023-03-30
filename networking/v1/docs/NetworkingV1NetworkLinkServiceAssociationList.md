# NetworkingV1NetworkLinkServiceAssociationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]NetworkingV1NetworkLinkServiceAssociation**](NetworkingV1NetworkLinkServiceAssociation.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewNetworkingV1NetworkLinkServiceAssociationList

`func NewNetworkingV1NetworkLinkServiceAssociationList(apiVersion string, kind string, metadata ListMeta, data []NetworkingV1NetworkLinkServiceAssociation, ) *NetworkingV1NetworkLinkServiceAssociationList`

NewNetworkingV1NetworkLinkServiceAssociationList instantiates a new NetworkingV1NetworkLinkServiceAssociationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1NetworkLinkServiceAssociationListWithDefaults

`func NewNetworkingV1NetworkLinkServiceAssociationListWithDefaults() *NetworkingV1NetworkLinkServiceAssociationList`

NewNetworkingV1NetworkLinkServiceAssociationListWithDefaults instantiates a new NetworkingV1NetworkLinkServiceAssociationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkingV1NetworkLinkServiceAssociationList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkingV1NetworkLinkServiceAssociationList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkingV1NetworkLinkServiceAssociationList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *NetworkingV1NetworkLinkServiceAssociationList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1NetworkLinkServiceAssociationList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1NetworkLinkServiceAssociationList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *NetworkingV1NetworkLinkServiceAssociationList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkingV1NetworkLinkServiceAssociationList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkingV1NetworkLinkServiceAssociationList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *NetworkingV1NetworkLinkServiceAssociationList) GetData() []NetworkingV1NetworkLinkServiceAssociation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NetworkingV1NetworkLinkServiceAssociationList) GetDataOk() (*[]NetworkingV1NetworkLinkServiceAssociation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NetworkingV1NetworkLinkServiceAssociationList) SetData(v []NetworkingV1NetworkLinkServiceAssociation)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


