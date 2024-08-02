# ArtifactV1FlinkArtifactList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]ArtifactV1FlinkArtifact**](ArtifactV1FlinkArtifact.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewArtifactV1FlinkArtifactList

`func NewArtifactV1FlinkArtifactList(apiVersion string, kind string, metadata ListMeta, data []ArtifactV1FlinkArtifact, ) *ArtifactV1FlinkArtifactList`

NewArtifactV1FlinkArtifactList instantiates a new ArtifactV1FlinkArtifactList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactV1FlinkArtifactListWithDefaults

`func NewArtifactV1FlinkArtifactListWithDefaults() *ArtifactV1FlinkArtifactList`

NewArtifactV1FlinkArtifactListWithDefaults instantiates a new ArtifactV1FlinkArtifactList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ArtifactV1FlinkArtifactList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ArtifactV1FlinkArtifactList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ArtifactV1FlinkArtifactList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ArtifactV1FlinkArtifactList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ArtifactV1FlinkArtifactList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ArtifactV1FlinkArtifactList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ArtifactV1FlinkArtifactList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ArtifactV1FlinkArtifactList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ArtifactV1FlinkArtifactList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ArtifactV1FlinkArtifactList) GetData() []ArtifactV1FlinkArtifact`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ArtifactV1FlinkArtifactList) GetDataOk() (*[]ArtifactV1FlinkArtifact, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ArtifactV1FlinkArtifactList) SetData(v []ArtifactV1FlinkArtifact)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


