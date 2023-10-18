# AiV1DocCompletionsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Answer** | Pointer to **string** |  | [optional] 
**AskedAt** | Pointer to **time.Time** | The date and time at which this question was asked. It is represented in RFC3339 format and is in UTC. | [optional] 

## Methods

### NewAiV1DocCompletionsReply

`func NewAiV1DocCompletionsReply() *AiV1DocCompletionsReply`

NewAiV1DocCompletionsReply instantiates a new AiV1DocCompletionsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiV1DocCompletionsReplyWithDefaults

`func NewAiV1DocCompletionsReplyWithDefaults() *AiV1DocCompletionsReply`

NewAiV1DocCompletionsReplyWithDefaults instantiates a new AiV1DocCompletionsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *AiV1DocCompletionsReply) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AiV1DocCompletionsReply) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AiV1DocCompletionsReply) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AiV1DocCompletionsReply) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *AiV1DocCompletionsReply) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AiV1DocCompletionsReply) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AiV1DocCompletionsReply) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AiV1DocCompletionsReply) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *AiV1DocCompletionsReply) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiV1DocCompletionsReply) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiV1DocCompletionsReply) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiV1DocCompletionsReply) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *AiV1DocCompletionsReply) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AiV1DocCompletionsReply) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AiV1DocCompletionsReply) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AiV1DocCompletionsReply) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAnswer

`func (o *AiV1DocCompletionsReply) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *AiV1DocCompletionsReply) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *AiV1DocCompletionsReply) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *AiV1DocCompletionsReply) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetAskedAt

`func (o *AiV1DocCompletionsReply) GetAskedAt() time.Time`

GetAskedAt returns the AskedAt field if non-nil, zero value otherwise.

### GetAskedAtOk

`func (o *AiV1DocCompletionsReply) GetAskedAtOk() (*time.Time, bool)`

GetAskedAtOk returns a tuple with the AskedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskedAt

`func (o *AiV1DocCompletionsReply) SetAskedAt(v time.Time)`

SetAskedAt sets AskedAt field to given value.

### HasAskedAt

`func (o *AiV1DocCompletionsReply) HasAskedAt() bool`

HasAskedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


