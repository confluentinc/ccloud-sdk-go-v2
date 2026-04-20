# SchemaString

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **string** | Name of the subject. In Confluent Cloud, this field is returned only if Confluent-Accept-Unknown-Properties header is passed in | [optional] 
**Version** | Pointer to **int32** | Version number. In Confluent Cloud, this field is returned only if Confluent-Accept-Unknown-Properties header is passed in | [optional] 
**Guid** | Pointer to **string** | Globally unique identifier of the schema. In Confluent Cloud, this field is returned only if Confluent-Accept-Unknown-Properties header is passed in | [optional] 
**SchemaType** | Pointer to **string** | Schema type. In Confluent Cloud, this field is returned only if Confluent-Accept-Unknown-Properties header is passed in | [optional] 
**Schema** | Pointer to **string** | Schema string identified by the ID | [optional] 
**References** | Pointer to [**[]SchemaReference**](SchemaReference.md) | References to other schemas | [optional] 
**Metadata** | Pointer to [**NullableMetadata**](Metadata.md) |  | [optional] 
**RuleSet** | Pointer to [**NullableRuleSet**](RuleSet.md) |  | [optional] 
**SchemaTags** | Pointer to [**[]SchemaTags**](SchemaTags.md) | Schema tags | [optional] 
**MaxId** | Pointer to **int32** | Maximum ID | [optional] 
**Ts** | Pointer to **int64** | Timestamp when the schema was created. In Confluent Cloud, this field is returned only if Confluent-Accept-Unknown-Properties header is passed in | [optional] 
**Deleted** | Pointer to **bool** | Whether the schema has been deleted. In Confluent Cloud, this field is returned only if Confluent-Accept-Unknown-Properties header is passed in | [optional] 

## Methods

### NewSchemaString

`func NewSchemaString() *SchemaString`

NewSchemaString instantiates a new SchemaString object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaStringWithDefaults

`func NewSchemaStringWithDefaults() *SchemaString`

NewSchemaStringWithDefaults instantiates a new SchemaString object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *SchemaString) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SchemaString) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SchemaString) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SchemaString) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetVersion

`func (o *SchemaString) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SchemaString) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SchemaString) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SchemaString) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetGuid

`func (o *SchemaString) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *SchemaString) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *SchemaString) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *SchemaString) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetSchemaType

`func (o *SchemaString) GetSchemaType() string`

GetSchemaType returns the SchemaType field if non-nil, zero value otherwise.

### GetSchemaTypeOk

`func (o *SchemaString) GetSchemaTypeOk() (*string, bool)`

GetSchemaTypeOk returns a tuple with the SchemaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaType

`func (o *SchemaString) SetSchemaType(v string)`

SetSchemaType sets SchemaType field to given value.

### HasSchemaType

`func (o *SchemaString) HasSchemaType() bool`

HasSchemaType returns a boolean if a field has been set.

### GetSchema

`func (o *SchemaString) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *SchemaString) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *SchemaString) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *SchemaString) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetReferences

`func (o *SchemaString) GetReferences() []SchemaReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *SchemaString) GetReferencesOk() (*[]SchemaReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *SchemaString) SetReferences(v []SchemaReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *SchemaString) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetMetadata

`func (o *SchemaString) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SchemaString) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SchemaString) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SchemaString) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *SchemaString) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *SchemaString) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetRuleSet

`func (o *SchemaString) GetRuleSet() RuleSet`

GetRuleSet returns the RuleSet field if non-nil, zero value otherwise.

### GetRuleSetOk

`func (o *SchemaString) GetRuleSetOk() (*RuleSet, bool)`

GetRuleSetOk returns a tuple with the RuleSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleSet

`func (o *SchemaString) SetRuleSet(v RuleSet)`

SetRuleSet sets RuleSet field to given value.

### HasRuleSet

`func (o *SchemaString) HasRuleSet() bool`

HasRuleSet returns a boolean if a field has been set.

### SetRuleSetNil

`func (o *SchemaString) SetRuleSetNil(b bool)`

 SetRuleSetNil sets the value for RuleSet to be an explicit nil

### UnsetRuleSet
`func (o *SchemaString) UnsetRuleSet()`

UnsetRuleSet ensures that no value is present for RuleSet, not even an explicit nil
### GetSchemaTags

`func (o *SchemaString) GetSchemaTags() []SchemaTags`

GetSchemaTags returns the SchemaTags field if non-nil, zero value otherwise.

### GetSchemaTagsOk

`func (o *SchemaString) GetSchemaTagsOk() (*[]SchemaTags, bool)`

GetSchemaTagsOk returns a tuple with the SchemaTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaTags

`func (o *SchemaString) SetSchemaTags(v []SchemaTags)`

SetSchemaTags sets SchemaTags field to given value.

### HasSchemaTags

`func (o *SchemaString) HasSchemaTags() bool`

HasSchemaTags returns a boolean if a field has been set.

### GetMaxId

`func (o *SchemaString) GetMaxId() int32`

GetMaxId returns the MaxId field if non-nil, zero value otherwise.

### GetMaxIdOk

`func (o *SchemaString) GetMaxIdOk() (*int32, bool)`

GetMaxIdOk returns a tuple with the MaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxId

`func (o *SchemaString) SetMaxId(v int32)`

SetMaxId sets MaxId field to given value.

### HasMaxId

`func (o *SchemaString) HasMaxId() bool`

HasMaxId returns a boolean if a field has been set.

### GetTs

`func (o *SchemaString) GetTs() int64`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *SchemaString) GetTsOk() (*int64, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *SchemaString) SetTs(v int64)`

SetTs sets Ts field to given value.

### HasTs

`func (o *SchemaString) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetDeleted

`func (o *SchemaString) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *SchemaString) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *SchemaString) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *SchemaString) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


