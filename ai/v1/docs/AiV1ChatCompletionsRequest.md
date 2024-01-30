# AiV1ChatCompletionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**AiSessionId** | Pointer to **string** |  | [optional] 
**Question** | Pointer to **string** |  | [optional] 
**DriftEnabled** | Pointer to **bool** |  | [optional] 
**History** | Pointer to [**[]AiV1ChatCompletionsHistory**](AiV1ChatCompletionsHistory.md) | Completion objects from previous interactions. | [optional] 

## Methods

### NewAiV1ChatCompletionsRequest

`func NewAiV1ChatCompletionsRequest() *AiV1ChatCompletionsRequest`

NewAiV1ChatCompletionsRequest instantiates a new AiV1ChatCompletionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiV1ChatCompletionsRequestWithDefaults

`func NewAiV1ChatCompletionsRequestWithDefaults() *AiV1ChatCompletionsRequest`

NewAiV1ChatCompletionsRequestWithDefaults instantiates a new AiV1ChatCompletionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *AiV1ChatCompletionsRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AiV1ChatCompletionsRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AiV1ChatCompletionsRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AiV1ChatCompletionsRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *AiV1ChatCompletionsRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AiV1ChatCompletionsRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AiV1ChatCompletionsRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AiV1ChatCompletionsRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *AiV1ChatCompletionsRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiV1ChatCompletionsRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiV1ChatCompletionsRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiV1ChatCompletionsRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *AiV1ChatCompletionsRequest) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AiV1ChatCompletionsRequest) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AiV1ChatCompletionsRequest) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AiV1ChatCompletionsRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAiSessionId

`func (o *AiV1ChatCompletionsRequest) GetAiSessionId() string`

GetAiSessionId returns the AiSessionId field if non-nil, zero value otherwise.

### GetAiSessionIdOk

`func (o *AiV1ChatCompletionsRequest) GetAiSessionIdOk() (*string, bool)`

GetAiSessionIdOk returns a tuple with the AiSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiSessionId

`func (o *AiV1ChatCompletionsRequest) SetAiSessionId(v string)`

SetAiSessionId sets AiSessionId field to given value.

### HasAiSessionId

`func (o *AiV1ChatCompletionsRequest) HasAiSessionId() bool`

HasAiSessionId returns a boolean if a field has been set.

### GetQuestion

`func (o *AiV1ChatCompletionsRequest) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *AiV1ChatCompletionsRequest) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *AiV1ChatCompletionsRequest) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *AiV1ChatCompletionsRequest) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetDriftEnabled

`func (o *AiV1ChatCompletionsRequest) GetDriftEnabled() bool`

GetDriftEnabled returns the DriftEnabled field if non-nil, zero value otherwise.

### GetDriftEnabledOk

`func (o *AiV1ChatCompletionsRequest) GetDriftEnabledOk() (*bool, bool)`

GetDriftEnabledOk returns a tuple with the DriftEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriftEnabled

`func (o *AiV1ChatCompletionsRequest) SetDriftEnabled(v bool)`

SetDriftEnabled sets DriftEnabled field to given value.

### HasDriftEnabled

`func (o *AiV1ChatCompletionsRequest) HasDriftEnabled() bool`

HasDriftEnabled returns a boolean if a field has been set.

### GetHistory

`func (o *AiV1ChatCompletionsRequest) GetHistory() []AiV1ChatCompletionsHistory`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *AiV1ChatCompletionsRequest) GetHistoryOk() (*[]AiV1ChatCompletionsHistory, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *AiV1ChatCompletionsRequest) SetHistory(v []AiV1ChatCompletionsHistory)`

SetHistory sets History field to given value.

### HasHistory

`func (o *AiV1ChatCompletionsRequest) HasHistory() bool`

HasHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


