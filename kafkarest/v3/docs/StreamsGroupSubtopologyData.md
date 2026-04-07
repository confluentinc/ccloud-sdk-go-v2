# StreamsGroupSubtopologyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**ClusterId** | **string** | The unique identifier of the Kafka cluster. | 
**GroupId** | **string** | The unique identifier of the Streams group. | 
**SubtopologyId** | **string** | The unique identifier of the Streams subtopology. | 
**SourceTopics** | **[]string** | The list of source topics for the subtopology. | 

## Methods

### NewStreamsGroupSubtopologyData

`func NewStreamsGroupSubtopologyData(kind string, metadata ResourceMetadata, clusterId string, groupId string, subtopologyId string, sourceTopics []string, ) *StreamsGroupSubtopologyData`

NewStreamsGroupSubtopologyData instantiates a new StreamsGroupSubtopologyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsGroupSubtopologyDataWithDefaults

`func NewStreamsGroupSubtopologyDataWithDefaults() *StreamsGroupSubtopologyData`

NewStreamsGroupSubtopologyDataWithDefaults instantiates a new StreamsGroupSubtopologyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *StreamsGroupSubtopologyData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StreamsGroupSubtopologyData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StreamsGroupSubtopologyData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *StreamsGroupSubtopologyData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StreamsGroupSubtopologyData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StreamsGroupSubtopologyData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *StreamsGroupSubtopologyData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *StreamsGroupSubtopologyData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *StreamsGroupSubtopologyData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetGroupId

`func (o *StreamsGroupSubtopologyData) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *StreamsGroupSubtopologyData) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *StreamsGroupSubtopologyData) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetSubtopologyId

`func (o *StreamsGroupSubtopologyData) GetSubtopologyId() string`

GetSubtopologyId returns the SubtopologyId field if non-nil, zero value otherwise.

### GetSubtopologyIdOk

`func (o *StreamsGroupSubtopologyData) GetSubtopologyIdOk() (*string, bool)`

GetSubtopologyIdOk returns a tuple with the SubtopologyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtopologyId

`func (o *StreamsGroupSubtopologyData) SetSubtopologyId(v string)`

SetSubtopologyId sets SubtopologyId field to given value.


### GetSourceTopics

`func (o *StreamsGroupSubtopologyData) GetSourceTopics() []string`

GetSourceTopics returns the SourceTopics field if non-nil, zero value otherwise.

### GetSourceTopicsOk

`func (o *StreamsGroupSubtopologyData) GetSourceTopicsOk() (*[]string, bool)`

GetSourceTopicsOk returns a tuple with the SourceTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTopics

`func (o *StreamsGroupSubtopologyData) SetSourceTopics(v []string)`

SetSourceTopics sets SourceTopics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


