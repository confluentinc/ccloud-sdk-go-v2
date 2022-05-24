# NetworkingAdminV1PrivateLinkAccessList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]NetworkingAdminV1PrivateLinkAccess**](NetworkingAdminV1PrivateLinkAccess.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewNetworkingAdminV1PrivateLinkAccessList

`func NewNetworkingAdminV1PrivateLinkAccessList(apiVersion string, kind string, metadata ListMeta, data []NetworkingAdminV1PrivateLinkAccess, ) *NetworkingAdminV1PrivateLinkAccessList`

NewNetworkingAdminV1PrivateLinkAccessList instantiates a new NetworkingAdminV1PrivateLinkAccessList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingAdminV1PrivateLinkAccessListWithDefaults

`func NewNetworkingAdminV1PrivateLinkAccessListWithDefaults() *NetworkingAdminV1PrivateLinkAccessList`

NewNetworkingAdminV1PrivateLinkAccessListWithDefaults instantiates a new NetworkingAdminV1PrivateLinkAccessList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkingAdminV1PrivateLinkAccessList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkingAdminV1PrivateLinkAccessList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkingAdminV1PrivateLinkAccessList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *NetworkingAdminV1PrivateLinkAccessList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingAdminV1PrivateLinkAccessList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingAdminV1PrivateLinkAccessList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *NetworkingAdminV1PrivateLinkAccessList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkingAdminV1PrivateLinkAccessList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkingAdminV1PrivateLinkAccessList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *NetworkingAdminV1PrivateLinkAccessList) GetData() []NetworkingAdminV1PrivateLinkAccess`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NetworkingAdminV1PrivateLinkAccessList) GetDataOk() (*[]NetworkingAdminV1PrivateLinkAccess, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NetworkingAdminV1PrivateLinkAccessList) SetData(v []NetworkingAdminV1PrivateLinkAccess)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


