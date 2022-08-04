# CdxV1ConsumerSharedResourceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]CdxV1ConsumerSharedResource**](CdxV1ConsumerSharedResource.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewCdxV1ConsumerSharedResourceList

`func NewCdxV1ConsumerSharedResourceList(apiVersion string, kind string, metadata ListMeta, data []CdxV1ConsumerSharedResource, ) *CdxV1ConsumerSharedResourceList`

NewCdxV1ConsumerSharedResourceList instantiates a new CdxV1ConsumerSharedResourceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdxV1ConsumerSharedResourceListWithDefaults

`func NewCdxV1ConsumerSharedResourceListWithDefaults() *CdxV1ConsumerSharedResourceList`

NewCdxV1ConsumerSharedResourceListWithDefaults instantiates a new CdxV1ConsumerSharedResourceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CdxV1ConsumerSharedResourceList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CdxV1ConsumerSharedResourceList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CdxV1ConsumerSharedResourceList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *CdxV1ConsumerSharedResourceList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CdxV1ConsumerSharedResourceList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CdxV1ConsumerSharedResourceList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *CdxV1ConsumerSharedResourceList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CdxV1ConsumerSharedResourceList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CdxV1ConsumerSharedResourceList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *CdxV1ConsumerSharedResourceList) GetData() []CdxV1ConsumerSharedResource`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CdxV1ConsumerSharedResourceList) GetDataOk() (*[]CdxV1ConsumerSharedResource, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CdxV1ConsumerSharedResourceList) SetData(v []CdxV1ConsumerSharedResource)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


