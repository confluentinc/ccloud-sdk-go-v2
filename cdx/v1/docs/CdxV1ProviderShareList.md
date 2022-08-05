# CdxV1ProviderShareList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]CdxV1ProviderShare**](CdxV1ProviderShare.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewCdxV1ProviderShareList

`func NewCdxV1ProviderShareList(apiVersion string, kind string, metadata ListMeta, data []CdxV1ProviderShare, ) *CdxV1ProviderShareList`

NewCdxV1ProviderShareList instantiates a new CdxV1ProviderShareList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdxV1ProviderShareListWithDefaults

`func NewCdxV1ProviderShareListWithDefaults() *CdxV1ProviderShareList`

NewCdxV1ProviderShareListWithDefaults instantiates a new CdxV1ProviderShareList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CdxV1ProviderShareList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CdxV1ProviderShareList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CdxV1ProviderShareList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *CdxV1ProviderShareList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CdxV1ProviderShareList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CdxV1ProviderShareList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *CdxV1ProviderShareList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CdxV1ProviderShareList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CdxV1ProviderShareList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *CdxV1ProviderShareList) GetData() []CdxV1ProviderShare`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CdxV1ProviderShareList) GetDataOk() (*[]CdxV1ProviderShare, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CdxV1ProviderShareList) SetData(v []CdxV1ProviderShare)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


