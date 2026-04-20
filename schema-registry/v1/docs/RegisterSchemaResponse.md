# RegisterSchemaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Globally unique identifier of the schema | [optional] 
**Version** | Pointer to **int32** | Version number. In Confluent Cloud, this field is returned only if Confluent-Accept-Unknown-Properties header is passed in | [optional] 
**Guid** | Pointer to **string** | Globally unique identifier of the schema. In Confluent Cloud, this field is returned only if Confluent-Accept-Unknown-Properties header is passed in | [optional] 
**SchemaType** | Pointer to **string** | Schema type. In Confluent Cloud, this field is returned only if Confluent-Accept-Unknown-Properties header is passed in | [optional] 
**References** | Pointer to [**[]SchemaReference**](SchemaReference.md) | References to other schemas. In Confluent Cloud, this field is returned only if Confluent-Accept-Unknown-Properties header is passed in | [optional] 
**Metadata** | Pointer to [**NullableMetadata**](Metadata.md) |  | [optional] 
**RuleSet** | Pointer to [**NullableRuleSet**](RuleSet.md) |  | [optional] 
**Schema** | Pointer to **string** | Schema definition string. In Confluent Cloud, this field is returned only if Confluent-Accept-Unknown-Properties header is passed in | [optional] 
**Ts** | Pointer to **int64** | Timestamp when the schema was created. In Confluent Cloud, this field is returned only if Confluent-Accept-Unknown-Properties header is passed in | [optional] 
**Deleted** | Pointer to **bool** | Whether the schema has been deleted. In Confluent Cloud, this field is returned only if Confluent-Accept-Unknown-Properties header is passed in | [optional] 

## Methods

### NewRegisterSchemaResponse

`func NewRegisterSchemaResponse() *RegisterSchemaResponse`

NewRegisterSchemaResponse instantiates a new RegisterSchemaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterSchemaResponseWithDefaults

`func NewRegisterSchemaResponseWithDefaults() *RegisterSchemaResponse`

NewRegisterSchemaResponseWithDefaults instantiates a new RegisterSchemaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegisterSchemaResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegisterSchemaResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegisterSchemaResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RegisterSchemaResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *RegisterSchemaResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RegisterSchemaResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RegisterSchemaResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RegisterSchemaResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetGuid

`func (o *RegisterSchemaResponse) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *RegisterSchemaResponse) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *RegisterSchemaResponse) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *RegisterSchemaResponse) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetSchemaType

`func (o *RegisterSchemaResponse) GetSchemaType() string`

GetSchemaType returns the SchemaType field if non-nil, zero value otherwise.

### GetSchemaTypeOk

`func (o *RegisterSchemaResponse) GetSchemaTypeOk() (*string, bool)`

GetSchemaTypeOk returns a tuple with the SchemaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaType

`func (o *RegisterSchemaResponse) SetSchemaType(v string)`

SetSchemaType sets SchemaType field to given value.

### HasSchemaType

`func (o *RegisterSchemaResponse) HasSchemaType() bool`

HasSchemaType returns a boolean if a field has been set.

### GetReferences

`func (o *RegisterSchemaResponse) GetReferences() []SchemaReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *RegisterSchemaResponse) GetReferencesOk() (*[]SchemaReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *RegisterSchemaResponse) SetReferences(v []SchemaReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *RegisterSchemaResponse) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetMetadata

`func (o *RegisterSchemaResponse) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RegisterSchemaResponse) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RegisterSchemaResponse) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RegisterSchemaResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *RegisterSchemaResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *RegisterSchemaResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetRuleSet

`func (o *RegisterSchemaResponse) GetRuleSet() RuleSet`

GetRuleSet returns the RuleSet field if non-nil, zero value otherwise.

### GetRuleSetOk

`func (o *RegisterSchemaResponse) GetRuleSetOk() (*RuleSet, bool)`

GetRuleSetOk returns a tuple with the RuleSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleSet

`func (o *RegisterSchemaResponse) SetRuleSet(v RuleSet)`

SetRuleSet sets RuleSet field to given value.

### HasRuleSet

`func (o *RegisterSchemaResponse) HasRuleSet() bool`

HasRuleSet returns a boolean if a field has been set.

### SetRuleSetNil

`func (o *RegisterSchemaResponse) SetRuleSetNil(b bool)`

 SetRuleSetNil sets the value for RuleSet to be an explicit nil

### UnsetRuleSet
`func (o *RegisterSchemaResponse) UnsetRuleSet()`

UnsetRuleSet ensures that no value is present for RuleSet, not even an explicit nil
### GetSchema

`func (o *RegisterSchemaResponse) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *RegisterSchemaResponse) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *RegisterSchemaResponse) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *RegisterSchemaResponse) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetTs

`func (o *RegisterSchemaResponse) GetTs() int64`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *RegisterSchemaResponse) GetTsOk() (*int64, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *RegisterSchemaResponse) SetTs(v int64)`

SetTs sets Ts field to given value.

### HasTs

`func (o *RegisterSchemaResponse) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetDeleted

`func (o *RegisterSchemaResponse) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *RegisterSchemaResponse) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *RegisterSchemaResponse) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *RegisterSchemaResponse) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


