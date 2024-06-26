# AiV1TagSuggestionsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**TagSuggestions** | Pointer to [**AiV1TagSuggestionsResponse**](ai.v1.TagSuggestionsResponse.md) |  | [optional] 

## Methods

### NewAiV1TagSuggestionsReply

`func NewAiV1TagSuggestionsReply() *AiV1TagSuggestionsReply`

NewAiV1TagSuggestionsReply instantiates a new AiV1TagSuggestionsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiV1TagSuggestionsReplyWithDefaults

`func NewAiV1TagSuggestionsReplyWithDefaults() *AiV1TagSuggestionsReply`

NewAiV1TagSuggestionsReplyWithDefaults instantiates a new AiV1TagSuggestionsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *AiV1TagSuggestionsReply) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AiV1TagSuggestionsReply) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AiV1TagSuggestionsReply) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AiV1TagSuggestionsReply) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *AiV1TagSuggestionsReply) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AiV1TagSuggestionsReply) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AiV1TagSuggestionsReply) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AiV1TagSuggestionsReply) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *AiV1TagSuggestionsReply) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiV1TagSuggestionsReply) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiV1TagSuggestionsReply) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiV1TagSuggestionsReply) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *AiV1TagSuggestionsReply) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AiV1TagSuggestionsReply) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AiV1TagSuggestionsReply) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AiV1TagSuggestionsReply) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTagSuggestions

`func (o *AiV1TagSuggestionsReply) GetTagSuggestions() AiV1TagSuggestionsResponse`

GetTagSuggestions returns the TagSuggestions field if non-nil, zero value otherwise.

### GetTagSuggestionsOk

`func (o *AiV1TagSuggestionsReply) GetTagSuggestionsOk() (*AiV1TagSuggestionsResponse, bool)`

GetTagSuggestionsOk returns a tuple with the TagSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagSuggestions

`func (o *AiV1TagSuggestionsReply) SetTagSuggestions(v AiV1TagSuggestionsResponse)`

SetTagSuggestions sets TagSuggestions field to given value.

### HasTagSuggestions

`func (o *AiV1TagSuggestionsReply) HasTagSuggestions() bool`

HasTagSuggestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


