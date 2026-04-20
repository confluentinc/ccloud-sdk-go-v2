# RegisterSchemaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int32** | Version number | [optional] 
**Id** | Pointer to **int32** | Globally unique identifier of the schema | [optional] 
**SchemaType** | Pointer to **string** | Schema type | [optional] 
**References** | Pointer to [**[]SchemaReference**](SchemaReference.md) | References to other schemas | [optional] 
**Schema** | Pointer to **string** | Schema definition string | [optional] 
**Metadata** | Pointer to [**NullableMetadata**](Metadata.md) |  | [optional] 
**RuleSet** | Pointer to [**NullableRuleSet**](RuleSet.md) |  | [optional] 
**SchemaTagsToAdd** | Pointer to [**[]SchemaTags**](SchemaTags.md) |  | [optional] 
**SchemaTagsToRemove** | Pointer to [**[]SchemaTags**](SchemaTags.md) |  | [optional] 
**PropagateSchemaTags** | Pointer to **bool** |  | [optional] 

## Methods

### NewRegisterSchemaRequest

`func NewRegisterSchemaRequest() *RegisterSchemaRequest`

NewRegisterSchemaRequest instantiates a new RegisterSchemaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterSchemaRequestWithDefaults

`func NewRegisterSchemaRequestWithDefaults() *RegisterSchemaRequest`

NewRegisterSchemaRequestWithDefaults instantiates a new RegisterSchemaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *RegisterSchemaRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RegisterSchemaRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RegisterSchemaRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RegisterSchemaRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetId

`func (o *RegisterSchemaRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegisterSchemaRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegisterSchemaRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RegisterSchemaRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchemaType

`func (o *RegisterSchemaRequest) GetSchemaType() string`

GetSchemaType returns the SchemaType field if non-nil, zero value otherwise.

### GetSchemaTypeOk

`func (o *RegisterSchemaRequest) GetSchemaTypeOk() (*string, bool)`

GetSchemaTypeOk returns a tuple with the SchemaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaType

`func (o *RegisterSchemaRequest) SetSchemaType(v string)`

SetSchemaType sets SchemaType field to given value.

### HasSchemaType

`func (o *RegisterSchemaRequest) HasSchemaType() bool`

HasSchemaType returns a boolean if a field has been set.

### GetReferences

`func (o *RegisterSchemaRequest) GetReferences() []SchemaReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *RegisterSchemaRequest) GetReferencesOk() (*[]SchemaReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *RegisterSchemaRequest) SetReferences(v []SchemaReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *RegisterSchemaRequest) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSchema

`func (o *RegisterSchemaRequest) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *RegisterSchemaRequest) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *RegisterSchemaRequest) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *RegisterSchemaRequest) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetMetadata

`func (o *RegisterSchemaRequest) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RegisterSchemaRequest) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RegisterSchemaRequest) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RegisterSchemaRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *RegisterSchemaRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *RegisterSchemaRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetRuleSet

`func (o *RegisterSchemaRequest) GetRuleSet() RuleSet`

GetRuleSet returns the RuleSet field if non-nil, zero value otherwise.

### GetRuleSetOk

`func (o *RegisterSchemaRequest) GetRuleSetOk() (*RuleSet, bool)`

GetRuleSetOk returns a tuple with the RuleSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleSet

`func (o *RegisterSchemaRequest) SetRuleSet(v RuleSet)`

SetRuleSet sets RuleSet field to given value.

### HasRuleSet

`func (o *RegisterSchemaRequest) HasRuleSet() bool`

HasRuleSet returns a boolean if a field has been set.

### SetRuleSetNil

`func (o *RegisterSchemaRequest) SetRuleSetNil(b bool)`

 SetRuleSetNil sets the value for RuleSet to be an explicit nil

### UnsetRuleSet
`func (o *RegisterSchemaRequest) UnsetRuleSet()`

UnsetRuleSet ensures that no value is present for RuleSet, not even an explicit nil
### GetSchemaTagsToAdd

`func (o *RegisterSchemaRequest) GetSchemaTagsToAdd() []SchemaTags`

GetSchemaTagsToAdd returns the SchemaTagsToAdd field if non-nil, zero value otherwise.

### GetSchemaTagsToAddOk

`func (o *RegisterSchemaRequest) GetSchemaTagsToAddOk() (*[]SchemaTags, bool)`

GetSchemaTagsToAddOk returns a tuple with the SchemaTagsToAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaTagsToAdd

`func (o *RegisterSchemaRequest) SetSchemaTagsToAdd(v []SchemaTags)`

SetSchemaTagsToAdd sets SchemaTagsToAdd field to given value.

### HasSchemaTagsToAdd

`func (o *RegisterSchemaRequest) HasSchemaTagsToAdd() bool`

HasSchemaTagsToAdd returns a boolean if a field has been set.

### GetSchemaTagsToRemove

`func (o *RegisterSchemaRequest) GetSchemaTagsToRemove() []SchemaTags`

GetSchemaTagsToRemove returns the SchemaTagsToRemove field if non-nil, zero value otherwise.

### GetSchemaTagsToRemoveOk

`func (o *RegisterSchemaRequest) GetSchemaTagsToRemoveOk() (*[]SchemaTags, bool)`

GetSchemaTagsToRemoveOk returns a tuple with the SchemaTagsToRemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaTagsToRemove

`func (o *RegisterSchemaRequest) SetSchemaTagsToRemove(v []SchemaTags)`

SetSchemaTagsToRemove sets SchemaTagsToRemove field to given value.

### HasSchemaTagsToRemove

`func (o *RegisterSchemaRequest) HasSchemaTagsToRemove() bool`

HasSchemaTagsToRemove returns a boolean if a field has been set.

### GetPropagateSchemaTags

`func (o *RegisterSchemaRequest) GetPropagateSchemaTags() bool`

GetPropagateSchemaTags returns the PropagateSchemaTags field if non-nil, zero value otherwise.

### GetPropagateSchemaTagsOk

`func (o *RegisterSchemaRequest) GetPropagateSchemaTagsOk() (*bool, bool)`

GetPropagateSchemaTagsOk returns a tuple with the PropagateSchemaTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagateSchemaTags

`func (o *RegisterSchemaRequest) SetPropagateSchemaTags(v bool)`

SetPropagateSchemaTags sets PropagateSchemaTags field to given value.

### HasPropagateSchemaTags

`func (o *RegisterSchemaRequest) HasPropagateSchemaTags() bool`

HasPropagateSchemaTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


