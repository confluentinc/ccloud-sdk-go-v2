# SrcmV3ClusterList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]SrcmV3Cluster**](SrcmV3Cluster.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewSrcmV3ClusterList

`func NewSrcmV3ClusterList(apiVersion string, kind string, metadata ListMeta, data []SrcmV3Cluster, ) *SrcmV3ClusterList`

NewSrcmV3ClusterList instantiates a new SrcmV3ClusterList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSrcmV3ClusterListWithDefaults

`func NewSrcmV3ClusterListWithDefaults() *SrcmV3ClusterList`

NewSrcmV3ClusterListWithDefaults instantiates a new SrcmV3ClusterList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SrcmV3ClusterList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SrcmV3ClusterList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SrcmV3ClusterList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *SrcmV3ClusterList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SrcmV3ClusterList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SrcmV3ClusterList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *SrcmV3ClusterList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SrcmV3ClusterList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SrcmV3ClusterList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *SrcmV3ClusterList) GetData() []SrcmV3Cluster`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SrcmV3ClusterList) GetDataOk() (*[]SrcmV3Cluster, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SrcmV3ClusterList) SetData(v []SrcmV3Cluster)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


