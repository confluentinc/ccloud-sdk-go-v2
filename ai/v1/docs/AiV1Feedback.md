# AiV1Feedback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**AiSessionId** | Pointer to **string** |  | [optional] 
**Reaction** | Pointer to **string** | The feedback reaction types are:  &#39;THUMBS_UP&#39;: when a user upvotes AI Assistant&#39;s answer.  &#39;THUMBS_DOWN&#39;: when a user downvotes the AI Assistant&#39;s answer.  | [optional] [default to "UNSPECIFIED"]
**Copied** | Pointer to **string** | The feedback copy types are:  &#39;ANSWER&#39;: when a user copies the AI Assistant&#39;s answer.  &#39;CODE&#39;: when a user copies the code snippet from the AI Assistant&#39;s answer.  | [optional] [default to "UNSPECIFIED"]
**Comment** | Pointer to **NullableString** | Additional text feedback optionally provided by a user. | [optional] 

## Methods

### NewAiV1Feedback

`func NewAiV1Feedback() *AiV1Feedback`

NewAiV1Feedback instantiates a new AiV1Feedback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiV1FeedbackWithDefaults

`func NewAiV1FeedbackWithDefaults() *AiV1Feedback`

NewAiV1FeedbackWithDefaults instantiates a new AiV1Feedback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *AiV1Feedback) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AiV1Feedback) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AiV1Feedback) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AiV1Feedback) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *AiV1Feedback) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AiV1Feedback) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AiV1Feedback) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AiV1Feedback) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *AiV1Feedback) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiV1Feedback) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiV1Feedback) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiV1Feedback) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *AiV1Feedback) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AiV1Feedback) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AiV1Feedback) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AiV1Feedback) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAiSessionId

`func (o *AiV1Feedback) GetAiSessionId() string`

GetAiSessionId returns the AiSessionId field if non-nil, zero value otherwise.

### GetAiSessionIdOk

`func (o *AiV1Feedback) GetAiSessionIdOk() (*string, bool)`

GetAiSessionIdOk returns a tuple with the AiSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiSessionId

`func (o *AiV1Feedback) SetAiSessionId(v string)`

SetAiSessionId sets AiSessionId field to given value.

### HasAiSessionId

`func (o *AiV1Feedback) HasAiSessionId() bool`

HasAiSessionId returns a boolean if a field has been set.

### GetReaction

`func (o *AiV1Feedback) GetReaction() string`

GetReaction returns the Reaction field if non-nil, zero value otherwise.

### GetReactionOk

`func (o *AiV1Feedback) GetReactionOk() (*string, bool)`

GetReactionOk returns a tuple with the Reaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReaction

`func (o *AiV1Feedback) SetReaction(v string)`

SetReaction sets Reaction field to given value.

### HasReaction

`func (o *AiV1Feedback) HasReaction() bool`

HasReaction returns a boolean if a field has been set.

### GetCopied

`func (o *AiV1Feedback) GetCopied() string`

GetCopied returns the Copied field if non-nil, zero value otherwise.

### GetCopiedOk

`func (o *AiV1Feedback) GetCopiedOk() (*string, bool)`

GetCopiedOk returns a tuple with the Copied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopied

`func (o *AiV1Feedback) SetCopied(v string)`

SetCopied sets Copied field to given value.

### HasCopied

`func (o *AiV1Feedback) HasCopied() bool`

HasCopied returns a boolean if a field has been set.

### GetComment

`func (o *AiV1Feedback) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AiV1Feedback) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AiV1Feedback) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AiV1Feedback) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *AiV1Feedback) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *AiV1Feedback) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


