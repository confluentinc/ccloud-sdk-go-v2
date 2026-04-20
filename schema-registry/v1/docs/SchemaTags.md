# SchemaTags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaEntity** | Pointer to [**SchemaEntity**](SchemaEntity.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSchemaTags

`func NewSchemaTags() *SchemaTags`

NewSchemaTags instantiates a new SchemaTags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaTagsWithDefaults

`func NewSchemaTagsWithDefaults() *SchemaTags`

NewSchemaTagsWithDefaults instantiates a new SchemaTags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaEntity

`func (o *SchemaTags) GetSchemaEntity() SchemaEntity`

GetSchemaEntity returns the SchemaEntity field if non-nil, zero value otherwise.

### GetSchemaEntityOk

`func (o *SchemaTags) GetSchemaEntityOk() (*SchemaEntity, bool)`

GetSchemaEntityOk returns a tuple with the SchemaEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaEntity

`func (o *SchemaTags) SetSchemaEntity(v SchemaEntity)`

SetSchemaEntity sets SchemaEntity field to given value.

### HasSchemaEntity

`func (o *SchemaTags) HasSchemaEntity() bool`

HasSchemaEntity returns a boolean if a field has been set.

### GetTags

`func (o *SchemaTags) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SchemaTags) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SchemaTags) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SchemaTags) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


