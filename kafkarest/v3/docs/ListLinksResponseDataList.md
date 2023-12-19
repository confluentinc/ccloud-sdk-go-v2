# ListLinksResponseDataList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceCollectionMetadata**](ResourceCollectionMetadata.md) |  | 
**Data** | [**[]ListLinksResponseData**](ListLinksResponseData.md) |  | 

## Methods

### NewListLinksResponseDataList

`func NewListLinksResponseDataList(kind string, metadata ResourceCollectionMetadata, data []ListLinksResponseData, ) *ListLinksResponseDataList`

NewListLinksResponseDataList instantiates a new ListLinksResponseDataList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLinksResponseDataListWithDefaults

`func NewListLinksResponseDataListWithDefaults() *ListLinksResponseDataList`

NewListLinksResponseDataListWithDefaults instantiates a new ListLinksResponseDataList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ListLinksResponseDataList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListLinksResponseDataList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListLinksResponseDataList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ListLinksResponseDataList) GetMetadata() ResourceCollectionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListLinksResponseDataList) GetMetadataOk() (*ResourceCollectionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListLinksResponseDataList) SetMetadata(v ResourceCollectionMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ListLinksResponseDataList) GetData() []ListLinksResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListLinksResponseDataList) GetDataOk() (*[]ListLinksResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListLinksResponseDataList) SetData(v []ListLinksResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


