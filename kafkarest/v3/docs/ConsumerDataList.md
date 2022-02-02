# ConsumerDataList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceCollectionMetadata**](ResourceCollectionMetadata.md) |  | 
**Data** | [**[]ConsumerData**](ConsumerData.md) |  | 

## Methods

### NewConsumerDataList

`func NewConsumerDataList(kind string, metadata ResourceCollectionMetadata, data []ConsumerData, ) *ConsumerDataList`

NewConsumerDataList instantiates a new ConsumerDataList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerDataListWithDefaults

`func NewConsumerDataListWithDefaults() *ConsumerDataList`

NewConsumerDataListWithDefaults instantiates a new ConsumerDataList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ConsumerDataList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConsumerDataList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConsumerDataList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ConsumerDataList) GetMetadata() ResourceCollectionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConsumerDataList) GetMetadataOk() (*ResourceCollectionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConsumerDataList) SetMetadata(v ResourceCollectionMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ConsumerDataList) GetData() []ConsumerData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConsumerDataList) GetDataOk() (*[]ConsumerData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConsumerDataList) SetData(v []ConsumerData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


