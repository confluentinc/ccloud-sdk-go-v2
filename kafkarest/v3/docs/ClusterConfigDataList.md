# ClusterConfigDataList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceCollectionMetadata**](ResourceCollectionMetadata.md) |  | 
**Data** | [**[]ClusterConfigData**](ClusterConfigData.md) |  | 

## Methods

### NewClusterConfigDataList

`func NewClusterConfigDataList(kind string, metadata ResourceCollectionMetadata, data []ClusterConfigData, ) *ClusterConfigDataList`

NewClusterConfigDataList instantiates a new ClusterConfigDataList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConfigDataListWithDefaults

`func NewClusterConfigDataListWithDefaults() *ClusterConfigDataList`

NewClusterConfigDataListWithDefaults instantiates a new ClusterConfigDataList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ClusterConfigDataList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterConfigDataList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterConfigDataList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ClusterConfigDataList) GetMetadata() ResourceCollectionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ClusterConfigDataList) GetMetadataOk() (*ResourceCollectionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ClusterConfigDataList) SetMetadata(v ResourceCollectionMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ClusterConfigDataList) GetData() []ClusterConfigData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ClusterConfigDataList) GetDataOk() (*[]ClusterConfigData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ClusterConfigDataList) SetData(v []ClusterConfigData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


