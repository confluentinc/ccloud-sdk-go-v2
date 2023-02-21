# ByokV1KeyList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]ByokV1Key**](ByokV1Key.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewByokV1KeyList

`func NewByokV1KeyList(apiVersion string, kind string, metadata ListMeta, data []ByokV1Key, ) *ByokV1KeyList`

NewByokV1KeyList instantiates a new ByokV1KeyList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewByokV1KeyListWithDefaults

`func NewByokV1KeyListWithDefaults() *ByokV1KeyList`

NewByokV1KeyListWithDefaults instantiates a new ByokV1KeyList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ByokV1KeyList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ByokV1KeyList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ByokV1KeyList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ByokV1KeyList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ByokV1KeyList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ByokV1KeyList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ByokV1KeyList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ByokV1KeyList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ByokV1KeyList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ByokV1KeyList) GetData() []ByokV1Key`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ByokV1KeyList) GetDataOk() (*[]ByokV1Key, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ByokV1KeyList) SetData(v []ByokV1Key)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


