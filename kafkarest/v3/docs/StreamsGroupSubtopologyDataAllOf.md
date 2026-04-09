# StreamsGroupSubtopologyDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** | The unique identifier of the Kafka cluster. | 
**GroupId** | **string** | The unique identifier of the Streams group. | 
**SubtopologyId** | **string** | The unique identifier of the Streams subtopology. | 
**SourceTopics** | **[]string** | The list of source topics for the subtopology. | 

## Methods

### NewStreamsGroupSubtopologyDataAllOf

`func NewStreamsGroupSubtopologyDataAllOf(clusterId string, groupId string, subtopologyId string, sourceTopics []string, ) *StreamsGroupSubtopologyDataAllOf`

NewStreamsGroupSubtopologyDataAllOf instantiates a new StreamsGroupSubtopologyDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsGroupSubtopologyDataAllOfWithDefaults

`func NewStreamsGroupSubtopologyDataAllOfWithDefaults() *StreamsGroupSubtopologyDataAllOf`

NewStreamsGroupSubtopologyDataAllOfWithDefaults instantiates a new StreamsGroupSubtopologyDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *StreamsGroupSubtopologyDataAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *StreamsGroupSubtopologyDataAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *StreamsGroupSubtopologyDataAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetGroupId

`func (o *StreamsGroupSubtopologyDataAllOf) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *StreamsGroupSubtopologyDataAllOf) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *StreamsGroupSubtopologyDataAllOf) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetSubtopologyId

`func (o *StreamsGroupSubtopologyDataAllOf) GetSubtopologyId() string`

GetSubtopologyId returns the SubtopologyId field if non-nil, zero value otherwise.

### GetSubtopologyIdOk

`func (o *StreamsGroupSubtopologyDataAllOf) GetSubtopologyIdOk() (*string, bool)`

GetSubtopologyIdOk returns a tuple with the SubtopologyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtopologyId

`func (o *StreamsGroupSubtopologyDataAllOf) SetSubtopologyId(v string)`

SetSubtopologyId sets SubtopologyId field to given value.


### GetSourceTopics

`func (o *StreamsGroupSubtopologyDataAllOf) GetSourceTopics() []string`

GetSourceTopics returns the SourceTopics field if non-nil, zero value otherwise.

### GetSourceTopicsOk

`func (o *StreamsGroupSubtopologyDataAllOf) GetSourceTopicsOk() (*[]string, bool)`

GetSourceTopicsOk returns a tuple with the SourceTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTopics

`func (o *StreamsGroupSubtopologyDataAllOf) SetSourceTopics(v []string)`

SetSourceTopics sets SourceTopics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


