# TopicConfigDataList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceCollectionMetadata**](ResourceCollectionMetadata.md) |  | 
**Data** | [**[]TopicConfigData**](TopicConfigData.md) |  | 

## Methods

### NewTopicConfigDataList

`func NewTopicConfigDataList(kind string, metadata ResourceCollectionMetadata, data []TopicConfigData, ) *TopicConfigDataList`

NewTopicConfigDataList instantiates a new TopicConfigDataList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicConfigDataListWithDefaults

`func NewTopicConfigDataListWithDefaults() *TopicConfigDataList`

NewTopicConfigDataListWithDefaults instantiates a new TopicConfigDataList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *TopicConfigDataList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TopicConfigDataList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TopicConfigDataList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *TopicConfigDataList) GetMetadata() ResourceCollectionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TopicConfigDataList) GetMetadataOk() (*ResourceCollectionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TopicConfigDataList) SetMetadata(v ResourceCollectionMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *TopicConfigDataList) GetData() []TopicConfigData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TopicConfigDataList) GetDataOk() (*[]TopicConfigData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TopicConfigDataList) SetData(v []TopicConfigData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


