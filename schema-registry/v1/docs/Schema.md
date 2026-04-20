# Schema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **string** | Name of the subject | [optional] 
**Version** | Pointer to **int32** | Version number | [optional] 
**Id** | Pointer to **int32** | Globally unique identifier of the schema | [optional] 
**Guid** | Pointer to **string** | Globally unique identifier of the schema. In Confluent Cloud, this field is returned only if Confluent-Accept-Unknown-Properties header is passed in | [optional] 
**SchemaType** | Pointer to **string** | Schema type. In Confluent Cloud, this field is returned only if Confluent-Accept-Unknown-Properties header is passed in | [optional] 
**References** | Pointer to [**[]SchemaReference**](SchemaReference.md) | References to other schemas | [optional] 
**Metadata** | Pointer to [**NullableMetadata**](Metadata.md) |  | [optional] 
**RuleSet** | Pointer to [**NullableRuleSet**](RuleSet.md) |  | [optional] 
**Schema** | Pointer to **string** | Schema definition string | [optional] 
**SchemaTags** | Pointer to [**[]SchemaTags**](SchemaTags.md) | Schema tags | [optional] 
**Ts** | Pointer to **int64** | Timestamp when the schema was created. In Confluent Cloud, this field is returned only if Confluent-Accept-Unknown-Properties header is passed in | [optional] 
**Deleted** | Pointer to **bool** | Whether the schema has been deleted. In Confluent Cloud, this field is returned only if Confluent-Accept-Unknown-Properties header is passed in | [optional] 

## Methods

### NewSchema

`func NewSchema() *Schema`

NewSchema instantiates a new Schema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaWithDefaults

`func NewSchemaWithDefaults() *Schema`

NewSchemaWithDefaults instantiates a new Schema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *Schema) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Schema) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Schema) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Schema) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetVersion

`func (o *Schema) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Schema) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Schema) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Schema) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetId

`func (o *Schema) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Schema) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Schema) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Schema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGuid

`func (o *Schema) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Schema) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Schema) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *Schema) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetSchemaType

`func (o *Schema) GetSchemaType() string`

GetSchemaType returns the SchemaType field if non-nil, zero value otherwise.

### GetSchemaTypeOk

`func (o *Schema) GetSchemaTypeOk() (*string, bool)`

GetSchemaTypeOk returns a tuple with the SchemaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaType

`func (o *Schema) SetSchemaType(v string)`

SetSchemaType sets SchemaType field to given value.

### HasSchemaType

`func (o *Schema) HasSchemaType() bool`

HasSchemaType returns a boolean if a field has been set.

### GetReferences

`func (o *Schema) GetReferences() []SchemaReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *Schema) GetReferencesOk() (*[]SchemaReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *Schema) SetReferences(v []SchemaReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *Schema) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetMetadata

`func (o *Schema) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Schema) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Schema) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Schema) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *Schema) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *Schema) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetRuleSet

`func (o *Schema) GetRuleSet() RuleSet`

GetRuleSet returns the RuleSet field if non-nil, zero value otherwise.

### GetRuleSetOk

`func (o *Schema) GetRuleSetOk() (*RuleSet, bool)`

GetRuleSetOk returns a tuple with the RuleSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleSet

`func (o *Schema) SetRuleSet(v RuleSet)`

SetRuleSet sets RuleSet field to given value.

### HasRuleSet

`func (o *Schema) HasRuleSet() bool`

HasRuleSet returns a boolean if a field has been set.

### SetRuleSetNil

`func (o *Schema) SetRuleSetNil(b bool)`

 SetRuleSetNil sets the value for RuleSet to be an explicit nil

### UnsetRuleSet
`func (o *Schema) UnsetRuleSet()`

UnsetRuleSet ensures that no value is present for RuleSet, not even an explicit nil
### GetSchema

`func (o *Schema) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *Schema) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *Schema) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *Schema) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetSchemaTags

`func (o *Schema) GetSchemaTags() []SchemaTags`

GetSchemaTags returns the SchemaTags field if non-nil, zero value otherwise.

### GetSchemaTagsOk

`func (o *Schema) GetSchemaTagsOk() (*[]SchemaTags, bool)`

GetSchemaTagsOk returns a tuple with the SchemaTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaTags

`func (o *Schema) SetSchemaTags(v []SchemaTags)`

SetSchemaTags sets SchemaTags field to given value.

### HasSchemaTags

`func (o *Schema) HasSchemaTags() bool`

HasSchemaTags returns a boolean if a field has been set.

### GetTs

`func (o *Schema) GetTs() int64`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *Schema) GetTsOk() (*int64, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *Schema) SetTs(v int64)`

SetTs sets Ts field to given value.

### HasTs

`func (o *Schema) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetDeleted

`func (o *Schema) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Schema) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Schema) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Schema) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


