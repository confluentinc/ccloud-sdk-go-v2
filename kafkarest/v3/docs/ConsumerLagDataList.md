# ConsumerLagDataList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceCollectionMetadata**](ResourceCollectionMetadata.md) |  | 
**Data** | [**[]ConsumerLagData**](ConsumerLagData.md) |  | 

## Methods

### NewConsumerLagDataList

`func NewConsumerLagDataList(kind string, metadata ResourceCollectionMetadata, data []ConsumerLagData, ) *ConsumerLagDataList`

NewConsumerLagDataList instantiates a new ConsumerLagDataList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerLagDataListWithDefaults

`func NewConsumerLagDataListWithDefaults() *ConsumerLagDataList`

NewConsumerLagDataListWithDefaults instantiates a new ConsumerLagDataList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ConsumerLagDataList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConsumerLagDataList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConsumerLagDataList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ConsumerLagDataList) GetMetadata() ResourceCollectionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConsumerLagDataList) GetMetadataOk() (*ResourceCollectionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConsumerLagDataList) SetMetadata(v ResourceCollectionMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ConsumerLagDataList) GetData() []ConsumerLagData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConsumerLagDataList) GetDataOk() (*[]ConsumerLagData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConsumerLagDataList) SetData(v []ConsumerLagData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


