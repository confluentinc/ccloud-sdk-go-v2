# NetworkingV1NetworkList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]NetworkingV1Network**](NetworkingV1Network.md) |  | 

## Methods

### NewNetworkingV1NetworkList

`func NewNetworkingV1NetworkList(apiVersion string, kind string, metadata ListMeta, data []NetworkingV1Network, ) *NetworkingV1NetworkList`

NewNetworkingV1NetworkList instantiates a new NetworkingV1NetworkList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1NetworkListWithDefaults

`func NewNetworkingV1NetworkListWithDefaults() *NetworkingV1NetworkList`

NewNetworkingV1NetworkListWithDefaults instantiates a new NetworkingV1NetworkList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkingV1NetworkList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkingV1NetworkList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkingV1NetworkList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *NetworkingV1NetworkList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1NetworkList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1NetworkList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *NetworkingV1NetworkList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkingV1NetworkList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkingV1NetworkList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *NetworkingV1NetworkList) GetData() []NetworkingV1Network`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NetworkingV1NetworkList) GetDataOk() (*[]NetworkingV1Network, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NetworkingV1NetworkList) SetData(v []NetworkingV1Network)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


