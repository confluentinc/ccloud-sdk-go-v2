# NetworkingAdminV1NetworkList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]NetworkingAdminV1Network**](NetworkingAdminV1Network.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewNetworkingAdminV1NetworkList

`func NewNetworkingAdminV1NetworkList(apiVersion string, kind string, metadata ListMeta, data []NetworkingAdminV1Network, ) *NetworkingAdminV1NetworkList`

NewNetworkingAdminV1NetworkList instantiates a new NetworkingAdminV1NetworkList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingAdminV1NetworkListWithDefaults

`func NewNetworkingAdminV1NetworkListWithDefaults() *NetworkingAdminV1NetworkList`

NewNetworkingAdminV1NetworkListWithDefaults instantiates a new NetworkingAdminV1NetworkList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkingAdminV1NetworkList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkingAdminV1NetworkList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkingAdminV1NetworkList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *NetworkingAdminV1NetworkList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingAdminV1NetworkList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingAdminV1NetworkList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *NetworkingAdminV1NetworkList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkingAdminV1NetworkList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkingAdminV1NetworkList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *NetworkingAdminV1NetworkList) GetData() []NetworkingAdminV1Network`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NetworkingAdminV1NetworkList) GetDataOk() (*[]NetworkingAdminV1Network, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NetworkingAdminV1NetworkList) SetData(v []NetworkingAdminV1Network)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


