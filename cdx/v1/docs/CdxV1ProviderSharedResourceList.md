# CdxV1ProviderSharedResourceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]CdxV1ProviderSharedResource**](CdxV1ProviderSharedResource.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewCdxV1ProviderSharedResourceList

`func NewCdxV1ProviderSharedResourceList(apiVersion string, kind string, metadata ListMeta, data []CdxV1ProviderSharedResource, ) *CdxV1ProviderSharedResourceList`

NewCdxV1ProviderSharedResourceList instantiates a new CdxV1ProviderSharedResourceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdxV1ProviderSharedResourceListWithDefaults

`func NewCdxV1ProviderSharedResourceListWithDefaults() *CdxV1ProviderSharedResourceList`

NewCdxV1ProviderSharedResourceListWithDefaults instantiates a new CdxV1ProviderSharedResourceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CdxV1ProviderSharedResourceList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CdxV1ProviderSharedResourceList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CdxV1ProviderSharedResourceList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *CdxV1ProviderSharedResourceList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CdxV1ProviderSharedResourceList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CdxV1ProviderSharedResourceList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *CdxV1ProviderSharedResourceList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CdxV1ProviderSharedResourceList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CdxV1ProviderSharedResourceList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *CdxV1ProviderSharedResourceList) GetData() []CdxV1ProviderSharedResource`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CdxV1ProviderSharedResourceList) GetDataOk() (*[]CdxV1ProviderSharedResource, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CdxV1ProviderSharedResourceList) SetData(v []CdxV1ProviderSharedResource)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


