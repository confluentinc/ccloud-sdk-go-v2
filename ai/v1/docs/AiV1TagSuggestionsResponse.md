# AiV1TagSuggestionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**TagSuggestions** | Pointer to [**[]AiV1TaggedSchemaField**](AiV1TaggedSchemaField.md) | List of fields with their corresponding tags | [optional] 

## Methods

### NewAiV1TagSuggestionsResponse

`func NewAiV1TagSuggestionsResponse() *AiV1TagSuggestionsResponse`

NewAiV1TagSuggestionsResponse instantiates a new AiV1TagSuggestionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiV1TagSuggestionsResponseWithDefaults

`func NewAiV1TagSuggestionsResponseWithDefaults() *AiV1TagSuggestionsResponse`

NewAiV1TagSuggestionsResponseWithDefaults instantiates a new AiV1TagSuggestionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *AiV1TagSuggestionsResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AiV1TagSuggestionsResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AiV1TagSuggestionsResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AiV1TagSuggestionsResponse) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *AiV1TagSuggestionsResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AiV1TagSuggestionsResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AiV1TagSuggestionsResponse) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AiV1TagSuggestionsResponse) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *AiV1TagSuggestionsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiV1TagSuggestionsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiV1TagSuggestionsResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiV1TagSuggestionsResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *AiV1TagSuggestionsResponse) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AiV1TagSuggestionsResponse) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AiV1TagSuggestionsResponse) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AiV1TagSuggestionsResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTagSuggestions

`func (o *AiV1TagSuggestionsResponse) GetTagSuggestions() []AiV1TaggedSchemaField`

GetTagSuggestions returns the TagSuggestions field if non-nil, zero value otherwise.

### GetTagSuggestionsOk

`func (o *AiV1TagSuggestionsResponse) GetTagSuggestionsOk() (*[]AiV1TaggedSchemaField, bool)`

GetTagSuggestionsOk returns a tuple with the TagSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagSuggestions

`func (o *AiV1TagSuggestionsResponse) SetTagSuggestions(v []AiV1TaggedSchemaField)`

SetTagSuggestions sets TagSuggestions field to given value.

### HasTagSuggestions

`func (o *AiV1TagSuggestionsResponse) HasTagSuggestions() bool`

HasTagSuggestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


