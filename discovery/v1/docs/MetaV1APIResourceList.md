# MetaV1APIResourceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]MetaV1APIResource**](MetaV1APIResource.md) |  | 

## Methods

### NewMetaV1APIResourceList

`func NewMetaV1APIResourceList(apiVersion string, kind string, metadata ListMeta, data []MetaV1APIResource, ) *MetaV1APIResourceList`

NewMetaV1APIResourceList instantiates a new MetaV1APIResourceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaV1APIResourceListWithDefaults

`func NewMetaV1APIResourceListWithDefaults() *MetaV1APIResourceList`

NewMetaV1APIResourceListWithDefaults instantiates a new MetaV1APIResourceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *MetaV1APIResourceList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MetaV1APIResourceList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MetaV1APIResourceList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *MetaV1APIResourceList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MetaV1APIResourceList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MetaV1APIResourceList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *MetaV1APIResourceList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MetaV1APIResourceList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MetaV1APIResourceList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *MetaV1APIResourceList) GetData() []MetaV1APIResource`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MetaV1APIResourceList) GetDataOk() (*[]MetaV1APIResource, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MetaV1APIResourceList) SetData(v []MetaV1APIResource)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


