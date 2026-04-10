# StreamsTaskData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**SubtopologyId** | **string** | The unique identifier of the Streams subtopology. | 
**PartitionIds** | **[]int32** | The list of partition IDs assigned to the Streams task. | 

## Methods

### NewStreamsTaskData

`func NewStreamsTaskData(kind string, metadata ResourceMetadata, subtopologyId string, partitionIds []int32, ) *StreamsTaskData`

NewStreamsTaskData instantiates a new StreamsTaskData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsTaskDataWithDefaults

`func NewStreamsTaskDataWithDefaults() *StreamsTaskData`

NewStreamsTaskDataWithDefaults instantiates a new StreamsTaskData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *StreamsTaskData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StreamsTaskData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StreamsTaskData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *StreamsTaskData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StreamsTaskData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StreamsTaskData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetSubtopologyId

`func (o *StreamsTaskData) GetSubtopologyId() string`

GetSubtopologyId returns the SubtopologyId field if non-nil, zero value otherwise.

### GetSubtopologyIdOk

`func (o *StreamsTaskData) GetSubtopologyIdOk() (*string, bool)`

GetSubtopologyIdOk returns a tuple with the SubtopologyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtopologyId

`func (o *StreamsTaskData) SetSubtopologyId(v string)`

SetSubtopologyId sets SubtopologyId field to given value.


### GetPartitionIds

`func (o *StreamsTaskData) GetPartitionIds() []int32`

GetPartitionIds returns the PartitionIds field if non-nil, zero value otherwise.

### GetPartitionIdsOk

`func (o *StreamsTaskData) GetPartitionIdsOk() (*[]int32, bool)`

GetPartitionIdsOk returns a tuple with the PartitionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionIds

`func (o *StreamsTaskData) SetPartitionIds(v []int32)`

SetPartitionIds sets PartitionIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


