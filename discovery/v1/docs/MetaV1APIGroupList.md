# MetaV1APIGroupList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]MetaV1APIGroup**](MetaV1APIGroup.md) |  | 

## Methods

### NewMetaV1APIGroupList

`func NewMetaV1APIGroupList(apiVersion string, kind string, metadata ListMeta, data []MetaV1APIGroup, ) *MetaV1APIGroupList`

NewMetaV1APIGroupList instantiates a new MetaV1APIGroupList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaV1APIGroupListWithDefaults

`func NewMetaV1APIGroupListWithDefaults() *MetaV1APIGroupList`

NewMetaV1APIGroupListWithDefaults instantiates a new MetaV1APIGroupList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *MetaV1APIGroupList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MetaV1APIGroupList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MetaV1APIGroupList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *MetaV1APIGroupList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MetaV1APIGroupList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MetaV1APIGroupList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *MetaV1APIGroupList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MetaV1APIGroupList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MetaV1APIGroupList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *MetaV1APIGroupList) GetData() []MetaV1APIGroup`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MetaV1APIGroupList) GetDataOk() (*[]MetaV1APIGroup, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MetaV1APIGroupList) SetData(v []MetaV1APIGroup)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

