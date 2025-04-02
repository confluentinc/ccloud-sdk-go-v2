# CamV1ConnectArtifactList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]CamV1ConnectArtifact**](CamV1ConnectArtifact.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewCamV1ConnectArtifactList

`func NewCamV1ConnectArtifactList(apiVersion string, kind string, metadata ListMeta, data []CamV1ConnectArtifact, ) *CamV1ConnectArtifactList`

NewCamV1ConnectArtifactList instantiates a new CamV1ConnectArtifactList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCamV1ConnectArtifactListWithDefaults

`func NewCamV1ConnectArtifactListWithDefaults() *CamV1ConnectArtifactList`

NewCamV1ConnectArtifactListWithDefaults instantiates a new CamV1ConnectArtifactList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CamV1ConnectArtifactList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CamV1ConnectArtifactList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CamV1ConnectArtifactList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *CamV1ConnectArtifactList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CamV1ConnectArtifactList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CamV1ConnectArtifactList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *CamV1ConnectArtifactList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CamV1ConnectArtifactList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CamV1ConnectArtifactList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *CamV1ConnectArtifactList) GetData() []CamV1ConnectArtifact`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CamV1ConnectArtifactList) GetDataOk() (*[]CamV1ConnectArtifact, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CamV1ConnectArtifactList) SetData(v []CamV1ConnectArtifact)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


