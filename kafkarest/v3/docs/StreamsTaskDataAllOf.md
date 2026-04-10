# StreamsTaskDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubtopologyId** | **string** | The unique identifier of the Streams subtopology. | 
**PartitionIds** | **[]int32** | The list of partition IDs assigned to the Streams task. | 

## Methods

### NewStreamsTaskDataAllOf

`func NewStreamsTaskDataAllOf(subtopologyId string, partitionIds []int32, ) *StreamsTaskDataAllOf`

NewStreamsTaskDataAllOf instantiates a new StreamsTaskDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsTaskDataAllOfWithDefaults

`func NewStreamsTaskDataAllOfWithDefaults() *StreamsTaskDataAllOf`

NewStreamsTaskDataAllOfWithDefaults instantiates a new StreamsTaskDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubtopologyId

`func (o *StreamsTaskDataAllOf) GetSubtopologyId() string`

GetSubtopologyId returns the SubtopologyId field if non-nil, zero value otherwise.

### GetSubtopologyIdOk

`func (o *StreamsTaskDataAllOf) GetSubtopologyIdOk() (*string, bool)`

GetSubtopologyIdOk returns a tuple with the SubtopologyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtopologyId

`func (o *StreamsTaskDataAllOf) SetSubtopologyId(v string)`

SetSubtopologyId sets SubtopologyId field to given value.


### GetPartitionIds

`func (o *StreamsTaskDataAllOf) GetPartitionIds() []int32`

GetPartitionIds returns the PartitionIds field if non-nil, zero value otherwise.

### GetPartitionIdsOk

`func (o *StreamsTaskDataAllOf) GetPartitionIdsOk() (*[]int32, bool)`

GetPartitionIdsOk returns a tuple with the PartitionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionIds

`func (o *StreamsTaskDataAllOf) SetPartitionIds(v []int32)`

SetPartitionIds sets PartitionIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


