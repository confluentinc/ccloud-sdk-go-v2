# FcpmV2RegionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]FcpmV2Region**](FcpmV2Region.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewFcpmV2RegionList

`func NewFcpmV2RegionList(apiVersion string, kind string, metadata ListMeta, data []FcpmV2Region, ) *FcpmV2RegionList`

NewFcpmV2RegionList instantiates a new FcpmV2RegionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpmV2RegionListWithDefaults

`func NewFcpmV2RegionListWithDefaults() *FcpmV2RegionList`

NewFcpmV2RegionListWithDefaults instantiates a new FcpmV2RegionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *FcpmV2RegionList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FcpmV2RegionList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FcpmV2RegionList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *FcpmV2RegionList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FcpmV2RegionList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FcpmV2RegionList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *FcpmV2RegionList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FcpmV2RegionList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FcpmV2RegionList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *FcpmV2RegionList) GetData() []FcpmV2Region`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FcpmV2RegionList) GetDataOk() (*[]FcpmV2Region, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FcpmV2RegionList) SetData(v []FcpmV2Region)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


