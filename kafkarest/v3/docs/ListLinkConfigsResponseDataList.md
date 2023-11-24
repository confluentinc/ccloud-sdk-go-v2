# ListLinkConfigsResponseDataList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceCollectionMetadata**](ResourceCollectionMetadata.md) |  | 
**Data** | [**[]ListLinkConfigsResponseData**](ListLinkConfigsResponseData.md) |  | 

## Methods

### NewListLinkConfigsResponseDataList

`func NewListLinkConfigsResponseDataList(kind string, metadata ResourceCollectionMetadata, data []ListLinkConfigsResponseData, ) *ListLinkConfigsResponseDataList`

NewListLinkConfigsResponseDataList instantiates a new ListLinkConfigsResponseDataList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLinkConfigsResponseDataListWithDefaults

`func NewListLinkConfigsResponseDataListWithDefaults() *ListLinkConfigsResponseDataList`

NewListLinkConfigsResponseDataListWithDefaults instantiates a new ListLinkConfigsResponseDataList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ListLinkConfigsResponseDataList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListLinkConfigsResponseDataList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListLinkConfigsResponseDataList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ListLinkConfigsResponseDataList) GetMetadata() ResourceCollectionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListLinkConfigsResponseDataList) GetMetadataOk() (*ResourceCollectionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListLinkConfigsResponseDataList) SetMetadata(v ResourceCollectionMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ListLinkConfigsResponseDataList) GetData() []ListLinkConfigsResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListLinkConfigsResponseDataList) GetDataOk() (*[]ListLinkConfigsResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListLinkConfigsResponseDataList) SetData(v []ListLinkConfigsResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


