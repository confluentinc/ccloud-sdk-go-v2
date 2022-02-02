# RemoveBrokerTaskDataList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceCollectionMetadata**](ResourceCollectionMetadata.md) |  | 
**Data** | [**[]RemoveBrokerTaskData**](RemoveBrokerTaskData.md) |  | 

## Methods

### NewRemoveBrokerTaskDataList

`func NewRemoveBrokerTaskDataList(kind string, metadata ResourceCollectionMetadata, data []RemoveBrokerTaskData, ) *RemoveBrokerTaskDataList`

NewRemoveBrokerTaskDataList instantiates a new RemoveBrokerTaskDataList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveBrokerTaskDataListWithDefaults

`func NewRemoveBrokerTaskDataListWithDefaults() *RemoveBrokerTaskDataList`

NewRemoveBrokerTaskDataListWithDefaults instantiates a new RemoveBrokerTaskDataList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *RemoveBrokerTaskDataList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RemoveBrokerTaskDataList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RemoveBrokerTaskDataList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *RemoveBrokerTaskDataList) GetMetadata() ResourceCollectionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RemoveBrokerTaskDataList) GetMetadataOk() (*ResourceCollectionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RemoveBrokerTaskDataList) SetMetadata(v ResourceCollectionMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *RemoveBrokerTaskDataList) GetData() []RemoveBrokerTaskData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RemoveBrokerTaskDataList) GetDataOk() (*[]RemoveBrokerTaskData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RemoveBrokerTaskDataList) SetData(v []RemoveBrokerTaskData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


