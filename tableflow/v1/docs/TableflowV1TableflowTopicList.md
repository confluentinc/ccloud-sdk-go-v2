# TableflowV1TableflowTopicList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]TableflowV1TableflowTopic**](TableflowV1TableflowTopic.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewTableflowV1TableflowTopicList

`func NewTableflowV1TableflowTopicList(apiVersion string, kind string, metadata ListMeta, data []TableflowV1TableflowTopic, ) *TableflowV1TableflowTopicList`

NewTableflowV1TableflowTopicList instantiates a new TableflowV1TableflowTopicList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1TableflowTopicListWithDefaults

`func NewTableflowV1TableflowTopicListWithDefaults() *TableflowV1TableflowTopicList`

NewTableflowV1TableflowTopicListWithDefaults instantiates a new TableflowV1TableflowTopicList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *TableflowV1TableflowTopicList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *TableflowV1TableflowTopicList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *TableflowV1TableflowTopicList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *TableflowV1TableflowTopicList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TableflowV1TableflowTopicList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TableflowV1TableflowTopicList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *TableflowV1TableflowTopicList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TableflowV1TableflowTopicList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TableflowV1TableflowTopicList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *TableflowV1TableflowTopicList) GetData() []TableflowV1TableflowTopic`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TableflowV1TableflowTopicList) GetDataOk() (*[]TableflowV1TableflowTopic, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TableflowV1TableflowTopicList) SetData(v []TableflowV1TableflowTopic)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


