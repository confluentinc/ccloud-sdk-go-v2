# AiV1ChatCompletionsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**AiSessionId** | Pointer to **string** |  | [optional] 
**Answer** | Pointer to **string** |  | [optional] 
**Question** | Pointer to **string** |  | [optional] 

## Methods

### NewAiV1ChatCompletionsReply

`func NewAiV1ChatCompletionsReply() *AiV1ChatCompletionsReply`

NewAiV1ChatCompletionsReply instantiates a new AiV1ChatCompletionsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiV1ChatCompletionsReplyWithDefaults

`func NewAiV1ChatCompletionsReplyWithDefaults() *AiV1ChatCompletionsReply`

NewAiV1ChatCompletionsReplyWithDefaults instantiates a new AiV1ChatCompletionsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *AiV1ChatCompletionsReply) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AiV1ChatCompletionsReply) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AiV1ChatCompletionsReply) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AiV1ChatCompletionsReply) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *AiV1ChatCompletionsReply) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AiV1ChatCompletionsReply) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AiV1ChatCompletionsReply) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AiV1ChatCompletionsReply) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *AiV1ChatCompletionsReply) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiV1ChatCompletionsReply) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiV1ChatCompletionsReply) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiV1ChatCompletionsReply) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *AiV1ChatCompletionsReply) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AiV1ChatCompletionsReply) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AiV1ChatCompletionsReply) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AiV1ChatCompletionsReply) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAiSessionId

`func (o *AiV1ChatCompletionsReply) GetAiSessionId() string`

GetAiSessionId returns the AiSessionId field if non-nil, zero value otherwise.

### GetAiSessionIdOk

`func (o *AiV1ChatCompletionsReply) GetAiSessionIdOk() (*string, bool)`

GetAiSessionIdOk returns a tuple with the AiSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiSessionId

`func (o *AiV1ChatCompletionsReply) SetAiSessionId(v string)`

SetAiSessionId sets AiSessionId field to given value.

### HasAiSessionId

`func (o *AiV1ChatCompletionsReply) HasAiSessionId() bool`

HasAiSessionId returns a boolean if a field has been set.

### GetAnswer

`func (o *AiV1ChatCompletionsReply) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *AiV1ChatCompletionsReply) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *AiV1ChatCompletionsReply) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *AiV1ChatCompletionsReply) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetQuestion

`func (o *AiV1ChatCompletionsReply) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *AiV1ChatCompletionsReply) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *AiV1ChatCompletionsReply) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *AiV1ChatCompletionsReply) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


