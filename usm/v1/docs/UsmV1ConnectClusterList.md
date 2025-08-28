# UsmV1ConnectClusterList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]UsmV1ConnectCluster**](UsmV1ConnectCluster.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewUsmV1ConnectClusterList

`func NewUsmV1ConnectClusterList(apiVersion string, kind string, metadata ListMeta, data []UsmV1ConnectCluster, ) *UsmV1ConnectClusterList`

NewUsmV1ConnectClusterList instantiates a new UsmV1ConnectClusterList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsmV1ConnectClusterListWithDefaults

`func NewUsmV1ConnectClusterListWithDefaults() *UsmV1ConnectClusterList`

NewUsmV1ConnectClusterListWithDefaults instantiates a new UsmV1ConnectClusterList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *UsmV1ConnectClusterList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *UsmV1ConnectClusterList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *UsmV1ConnectClusterList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *UsmV1ConnectClusterList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *UsmV1ConnectClusterList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *UsmV1ConnectClusterList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *UsmV1ConnectClusterList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UsmV1ConnectClusterList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UsmV1ConnectClusterList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *UsmV1ConnectClusterList) GetData() []UsmV1ConnectCluster`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UsmV1ConnectClusterList) GetDataOk() (*[]UsmV1ConnectCluster, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UsmV1ConnectClusterList) SetData(v []UsmV1ConnectCluster)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


