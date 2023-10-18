# AiV1DocCompletionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Question** | Pointer to **string** |  | [optional] 
**History** | Pointer to [**[]AiV1DocCompletionsHistory**](AiV1DocCompletionsHistory.md) | Question and answer pairs from previous interactions. | [optional] 

## Methods

### NewAiV1DocCompletionsRequest

`func NewAiV1DocCompletionsRequest() *AiV1DocCompletionsRequest`

NewAiV1DocCompletionsRequest instantiates a new AiV1DocCompletionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiV1DocCompletionsRequestWithDefaults

`func NewAiV1DocCompletionsRequestWithDefaults() *AiV1DocCompletionsRequest`

NewAiV1DocCompletionsRequestWithDefaults instantiates a new AiV1DocCompletionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *AiV1DocCompletionsRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AiV1DocCompletionsRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AiV1DocCompletionsRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AiV1DocCompletionsRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *AiV1DocCompletionsRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AiV1DocCompletionsRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AiV1DocCompletionsRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AiV1DocCompletionsRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *AiV1DocCompletionsRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiV1DocCompletionsRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiV1DocCompletionsRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiV1DocCompletionsRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *AiV1DocCompletionsRequest) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AiV1DocCompletionsRequest) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AiV1DocCompletionsRequest) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AiV1DocCompletionsRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetQuestion

`func (o *AiV1DocCompletionsRequest) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *AiV1DocCompletionsRequest) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *AiV1DocCompletionsRequest) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *AiV1DocCompletionsRequest) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetHistory

`func (o *AiV1DocCompletionsRequest) GetHistory() []AiV1DocCompletionsHistory`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *AiV1DocCompletionsRequest) GetHistoryOk() (*[]AiV1DocCompletionsHistory, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *AiV1DocCompletionsRequest) SetHistory(v []AiV1DocCompletionsHistory)`

SetHistory sets History field to given value.

### HasHistory

`func (o *AiV1DocCompletionsRequest) HasHistory() bool`

HasHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


