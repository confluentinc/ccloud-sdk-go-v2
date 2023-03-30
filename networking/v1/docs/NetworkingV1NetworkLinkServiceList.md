# NetworkingV1NetworkLinkServiceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]NetworkingV1NetworkLinkService**](NetworkingV1NetworkLinkService.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewNetworkingV1NetworkLinkServiceList

`func NewNetworkingV1NetworkLinkServiceList(apiVersion string, kind string, metadata ListMeta, data []NetworkingV1NetworkLinkService, ) *NetworkingV1NetworkLinkServiceList`

NewNetworkingV1NetworkLinkServiceList instantiates a new NetworkingV1NetworkLinkServiceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1NetworkLinkServiceListWithDefaults

`func NewNetworkingV1NetworkLinkServiceListWithDefaults() *NetworkingV1NetworkLinkServiceList`

NewNetworkingV1NetworkLinkServiceListWithDefaults instantiates a new NetworkingV1NetworkLinkServiceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkingV1NetworkLinkServiceList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkingV1NetworkLinkServiceList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkingV1NetworkLinkServiceList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *NetworkingV1NetworkLinkServiceList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1NetworkLinkServiceList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1NetworkLinkServiceList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *NetworkingV1NetworkLinkServiceList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkingV1NetworkLinkServiceList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkingV1NetworkLinkServiceList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *NetworkingV1NetworkLinkServiceList) GetData() []NetworkingV1NetworkLinkService`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NetworkingV1NetworkLinkServiceList) GetDataOk() (*[]NetworkingV1NetworkLinkService, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NetworkingV1NetworkLinkServiceList) SetData(v []NetworkingV1NetworkLinkService)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


