# AlterMirrorStatusResponseDataList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceCollectionMetadata**](ResourceCollectionMetadata.md) |  | 
**Data** | [**[]AlterMirrorStatusResponseData**](AlterMirrorStatusResponseData.md) |  | 

## Methods

### NewAlterMirrorStatusResponseDataList

`func NewAlterMirrorStatusResponseDataList(kind string, metadata ResourceCollectionMetadata, data []AlterMirrorStatusResponseData, ) *AlterMirrorStatusResponseDataList`

NewAlterMirrorStatusResponseDataList instantiates a new AlterMirrorStatusResponseDataList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlterMirrorStatusResponseDataListWithDefaults

`func NewAlterMirrorStatusResponseDataListWithDefaults() *AlterMirrorStatusResponseDataList`

NewAlterMirrorStatusResponseDataListWithDefaults instantiates a new AlterMirrorStatusResponseDataList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *AlterMirrorStatusResponseDataList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AlterMirrorStatusResponseDataList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AlterMirrorStatusResponseDataList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *AlterMirrorStatusResponseDataList) GetMetadata() ResourceCollectionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AlterMirrorStatusResponseDataList) GetMetadataOk() (*ResourceCollectionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AlterMirrorStatusResponseDataList) SetMetadata(v ResourceCollectionMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *AlterMirrorStatusResponseDataList) GetData() []AlterMirrorStatusResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AlterMirrorStatusResponseDataList) GetDataOk() (*[]AlterMirrorStatusResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AlterMirrorStatusResponseDataList) SetData(v []AlterMirrorStatusResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


