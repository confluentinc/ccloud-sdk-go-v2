# AiV1TaggedSchemaField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | Pointer to **string** | Qualified name of the schema field | [optional] 
**Tags** | Pointer to [**[]AiV1TagDefinition**](AiV1TagDefinition.md) | List of suggested tag definitions for the schema field | [optional] 

## Methods

### NewAiV1TaggedSchemaField

`func NewAiV1TaggedSchemaField() *AiV1TaggedSchemaField`

NewAiV1TaggedSchemaField instantiates a new AiV1TaggedSchemaField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiV1TaggedSchemaFieldWithDefaults

`func NewAiV1TaggedSchemaFieldWithDefaults() *AiV1TaggedSchemaField`

NewAiV1TaggedSchemaFieldWithDefaults instantiates a new AiV1TaggedSchemaField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *AiV1TaggedSchemaField) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *AiV1TaggedSchemaField) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *AiV1TaggedSchemaField) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *AiV1TaggedSchemaField) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetTags

`func (o *AiV1TaggedSchemaField) GetTags() []AiV1TagDefinition`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AiV1TaggedSchemaField) GetTagsOk() (*[]AiV1TagDefinition, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AiV1TaggedSchemaField) SetTags(v []AiV1TagDefinition)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AiV1TaggedSchemaField) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


