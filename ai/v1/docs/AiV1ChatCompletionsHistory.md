# AiV1ChatCompletionsHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Question** | Pointer to **string** | Question for the AI assistant. | [optional] 
**Answer** | Pointer to **string** | Markdown-formatted answer from the AI assistant. | [optional] 
**AskedAt** | Pointer to **time.Time** | The date and time at which this question was asked. It is represented in RFC3339 format and is in UTC. | [optional] 

## Methods

### NewAiV1ChatCompletionsHistory

`func NewAiV1ChatCompletionsHistory() *AiV1ChatCompletionsHistory`

NewAiV1ChatCompletionsHistory instantiates a new AiV1ChatCompletionsHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiV1ChatCompletionsHistoryWithDefaults

`func NewAiV1ChatCompletionsHistoryWithDefaults() *AiV1ChatCompletionsHistory`

NewAiV1ChatCompletionsHistoryWithDefaults instantiates a new AiV1ChatCompletionsHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestion

`func (o *AiV1ChatCompletionsHistory) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *AiV1ChatCompletionsHistory) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *AiV1ChatCompletionsHistory) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *AiV1ChatCompletionsHistory) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetAnswer

`func (o *AiV1ChatCompletionsHistory) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *AiV1ChatCompletionsHistory) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *AiV1ChatCompletionsHistory) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *AiV1ChatCompletionsHistory) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetAskedAt

`func (o *AiV1ChatCompletionsHistory) GetAskedAt() time.Time`

GetAskedAt returns the AskedAt field if non-nil, zero value otherwise.

### GetAskedAtOk

`func (o *AiV1ChatCompletionsHistory) GetAskedAtOk() (*time.Time, bool)`

GetAskedAtOk returns a tuple with the AskedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskedAt

`func (o *AiV1ChatCompletionsHistory) SetAskedAt(v time.Time)`

SetAskedAt sets AskedAt field to given value.

### HasAskedAt

`func (o *AiV1ChatCompletionsHistory) HasAskedAt() bool`

HasAskedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


