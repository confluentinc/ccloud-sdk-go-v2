# SdV1PipelineStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The current state of the pipeline.:   DRAFT:  the pipeline is a draft and not activated yet;   ACTIVATING:  the pipeline activation is in progress;   DEACTIVATING:  the pipeline deactivation is in progress;   ACTIVE:  the pipeline is actived and running;   FAILED:  the pipeline activation or deactivation failed;   DELETED:  the pipeline is deleted  | [optional] [readonly] 
**TopicCount** | Pointer to **int32** | The number of Kafka topics defined in the pipeline. | [optional] [readonly] 
**ConnectorCount** | Pointer to **int32** | The number of connectors defined in the pipeline. | [optional] [readonly] 
**QueryCount** | Pointer to **int32** | The number of KSQL queries defined in the pipeline. | [optional] [readonly] 

## Methods

### NewSdV1PipelineStatus

`func NewSdV1PipelineStatus() *SdV1PipelineStatus`

NewSdV1PipelineStatus instantiates a new SdV1PipelineStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdV1PipelineStatusWithDefaults

`func NewSdV1PipelineStatusWithDefaults() *SdV1PipelineStatus`

NewSdV1PipelineStatusWithDefaults instantiates a new SdV1PipelineStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *SdV1PipelineStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SdV1PipelineStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SdV1PipelineStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SdV1PipelineStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTopicCount

`func (o *SdV1PipelineStatus) GetTopicCount() int32`

GetTopicCount returns the TopicCount field if non-nil, zero value otherwise.

### GetTopicCountOk

`func (o *SdV1PipelineStatus) GetTopicCountOk() (*int32, bool)`

GetTopicCountOk returns a tuple with the TopicCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicCount

`func (o *SdV1PipelineStatus) SetTopicCount(v int32)`

SetTopicCount sets TopicCount field to given value.

### HasTopicCount

`func (o *SdV1PipelineStatus) HasTopicCount() bool`

HasTopicCount returns a boolean if a field has been set.

### GetConnectorCount

`func (o *SdV1PipelineStatus) GetConnectorCount() int32`

GetConnectorCount returns the ConnectorCount field if non-nil, zero value otherwise.

### GetConnectorCountOk

`func (o *SdV1PipelineStatus) GetConnectorCountOk() (*int32, bool)`

GetConnectorCountOk returns a tuple with the ConnectorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorCount

`func (o *SdV1PipelineStatus) SetConnectorCount(v int32)`

SetConnectorCount sets ConnectorCount field to given value.

### HasConnectorCount

`func (o *SdV1PipelineStatus) HasConnectorCount() bool`

HasConnectorCount returns a boolean if a field has been set.

### GetQueryCount

`func (o *SdV1PipelineStatus) GetQueryCount() int32`

GetQueryCount returns the QueryCount field if non-nil, zero value otherwise.

### GetQueryCountOk

`func (o *SdV1PipelineStatus) GetQueryCountOk() (*int32, bool)`

GetQueryCountOk returns a tuple with the QueryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryCount

`func (o *SdV1PipelineStatus) SetQueryCount(v int32)`

SetQueryCount sets QueryCount field to given value.

### HasQueryCount

`func (o *SdV1PipelineStatus) HasQueryCount() bool`

HasQueryCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


