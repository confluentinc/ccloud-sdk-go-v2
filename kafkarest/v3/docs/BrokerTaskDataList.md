# BrokerTaskDataList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceCollectionMetadata**](ResourceCollectionMetadata.md) |  | 
**Data** | [**[]BrokerTaskData**](BrokerTaskData.md) |  | 

## Methods

### NewBrokerTaskDataList

`func NewBrokerTaskDataList(kind string, metadata ResourceCollectionMetadata, data []BrokerTaskData, ) *BrokerTaskDataList`

NewBrokerTaskDataList instantiates a new BrokerTaskDataList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrokerTaskDataListWithDefaults

`func NewBrokerTaskDataListWithDefaults() *BrokerTaskDataList`

NewBrokerTaskDataListWithDefaults instantiates a new BrokerTaskDataList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *BrokerTaskDataList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BrokerTaskDataList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BrokerTaskDataList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *BrokerTaskDataList) GetMetadata() ResourceCollectionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BrokerTaskDataList) GetMetadataOk() (*ResourceCollectionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BrokerTaskDataList) SetMetadata(v ResourceCollectionMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *BrokerTaskDataList) GetData() []BrokerTaskData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BrokerTaskDataList) GetDataOk() (*[]BrokerTaskData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BrokerTaskDataList) SetData(v []BrokerTaskData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


