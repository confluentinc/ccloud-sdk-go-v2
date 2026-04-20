# ConfigUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | If alias is specified, then this subject is an alias for the subject named by the alias. That means that any reference to this subject will be replaced by the alias. | [optional] 
**AliasForDeks** | Pointer to **string** | If aliasForDeks is specified, then this subject is an alias for the  subject named by the alias when accessing DEKs. That means that any  reference to this subject will be replaced by the alias when accessing DEKs. | [optional] 
**Normalize** | Pointer to **bool** | If true, then schemas are automatically normalized when registered or when passed during lookups. This means that clients do not have to pass the \&quot;normalize\&quot; query parameter to have normalization occur. | [optional] 
**ValidateFields** | Pointer to **bool** | If true, then additional validation will be performed for fields. | [optional] 
**ValidateNewSchemas** | Pointer to **bool** | If true, then additional validation will be performed for new schemas. | [optional] 
**ValidateRules** | Pointer to **bool** | If true, then additional validation will be performed for rules. | [optional] 
**Compatibility** | Pointer to **string** | Compatibility Level | [optional] 
**CompatibilityPolicy** | Pointer to **string** | Compatibility policy | [optional] 
**CompatibilityGroup** | Pointer to **string** | Only schemas that belong to the same compatibility group will be checked for compatibility. | [optional] 
**DefaultMetadata** | Pointer to [**ConfigDefaultMetadata**](ConfigDefaultMetadata.md) |  | [optional] 
**OverrideMetadata** | Pointer to [**ConfigOverrideMetadata**](ConfigOverrideMetadata.md) |  | [optional] 
**DefaultRuleSet** | Pointer to [**ConfigDefaultRuleSet**](ConfigDefaultRuleSet.md) |  | [optional] 
**OverrideRuleSet** | Pointer to [**ConfigOverrideRuleSet**](ConfigOverrideRuleSet.md) |  | [optional] 

## Methods

### NewConfigUpdateRequest

`func NewConfigUpdateRequest() *ConfigUpdateRequest`

NewConfigUpdateRequest instantiates a new ConfigUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigUpdateRequestWithDefaults

`func NewConfigUpdateRequestWithDefaults() *ConfigUpdateRequest`

NewConfigUpdateRequestWithDefaults instantiates a new ConfigUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *ConfigUpdateRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ConfigUpdateRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ConfigUpdateRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ConfigUpdateRequest) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetAliasForDeks

`func (o *ConfigUpdateRequest) GetAliasForDeks() string`

GetAliasForDeks returns the AliasForDeks field if non-nil, zero value otherwise.

### GetAliasForDeksOk

`func (o *ConfigUpdateRequest) GetAliasForDeksOk() (*string, bool)`

GetAliasForDeksOk returns a tuple with the AliasForDeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasForDeks

`func (o *ConfigUpdateRequest) SetAliasForDeks(v string)`

SetAliasForDeks sets AliasForDeks field to given value.

### HasAliasForDeks

`func (o *ConfigUpdateRequest) HasAliasForDeks() bool`

HasAliasForDeks returns a boolean if a field has been set.

### GetNormalize

`func (o *ConfigUpdateRequest) GetNormalize() bool`

GetNormalize returns the Normalize field if non-nil, zero value otherwise.

### GetNormalizeOk

`func (o *ConfigUpdateRequest) GetNormalizeOk() (*bool, bool)`

GetNormalizeOk returns a tuple with the Normalize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalize

`func (o *ConfigUpdateRequest) SetNormalize(v bool)`

SetNormalize sets Normalize field to given value.

### HasNormalize

`func (o *ConfigUpdateRequest) HasNormalize() bool`

HasNormalize returns a boolean if a field has been set.

### GetValidateFields

`func (o *ConfigUpdateRequest) GetValidateFields() bool`

GetValidateFields returns the ValidateFields field if non-nil, zero value otherwise.

### GetValidateFieldsOk

`func (o *ConfigUpdateRequest) GetValidateFieldsOk() (*bool, bool)`

GetValidateFieldsOk returns a tuple with the ValidateFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateFields

`func (o *ConfigUpdateRequest) SetValidateFields(v bool)`

SetValidateFields sets ValidateFields field to given value.

### HasValidateFields

`func (o *ConfigUpdateRequest) HasValidateFields() bool`

HasValidateFields returns a boolean if a field has been set.

### GetValidateNewSchemas

`func (o *ConfigUpdateRequest) GetValidateNewSchemas() bool`

GetValidateNewSchemas returns the ValidateNewSchemas field if non-nil, zero value otherwise.

### GetValidateNewSchemasOk

`func (o *ConfigUpdateRequest) GetValidateNewSchemasOk() (*bool, bool)`

GetValidateNewSchemasOk returns a tuple with the ValidateNewSchemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateNewSchemas

`func (o *ConfigUpdateRequest) SetValidateNewSchemas(v bool)`

SetValidateNewSchemas sets ValidateNewSchemas field to given value.

### HasValidateNewSchemas

`func (o *ConfigUpdateRequest) HasValidateNewSchemas() bool`

HasValidateNewSchemas returns a boolean if a field has been set.

### GetValidateRules

`func (o *ConfigUpdateRequest) GetValidateRules() bool`

GetValidateRules returns the ValidateRules field if non-nil, zero value otherwise.

### GetValidateRulesOk

`func (o *ConfigUpdateRequest) GetValidateRulesOk() (*bool, bool)`

GetValidateRulesOk returns a tuple with the ValidateRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateRules

`func (o *ConfigUpdateRequest) SetValidateRules(v bool)`

SetValidateRules sets ValidateRules field to given value.

### HasValidateRules

`func (o *ConfigUpdateRequest) HasValidateRules() bool`

HasValidateRules returns a boolean if a field has been set.

### GetCompatibility

`func (o *ConfigUpdateRequest) GetCompatibility() string`

GetCompatibility returns the Compatibility field if non-nil, zero value otherwise.

### GetCompatibilityOk

`func (o *ConfigUpdateRequest) GetCompatibilityOk() (*string, bool)`

GetCompatibilityOk returns a tuple with the Compatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibility

`func (o *ConfigUpdateRequest) SetCompatibility(v string)`

SetCompatibility sets Compatibility field to given value.

### HasCompatibility

`func (o *ConfigUpdateRequest) HasCompatibility() bool`

HasCompatibility returns a boolean if a field has been set.

### GetCompatibilityPolicy

`func (o *ConfigUpdateRequest) GetCompatibilityPolicy() string`

GetCompatibilityPolicy returns the CompatibilityPolicy field if non-nil, zero value otherwise.

### GetCompatibilityPolicyOk

`func (o *ConfigUpdateRequest) GetCompatibilityPolicyOk() (*string, bool)`

GetCompatibilityPolicyOk returns a tuple with the CompatibilityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityPolicy

`func (o *ConfigUpdateRequest) SetCompatibilityPolicy(v string)`

SetCompatibilityPolicy sets CompatibilityPolicy field to given value.

### HasCompatibilityPolicy

`func (o *ConfigUpdateRequest) HasCompatibilityPolicy() bool`

HasCompatibilityPolicy returns a boolean if a field has been set.

### GetCompatibilityGroup

`func (o *ConfigUpdateRequest) GetCompatibilityGroup() string`

GetCompatibilityGroup returns the CompatibilityGroup field if non-nil, zero value otherwise.

### GetCompatibilityGroupOk

`func (o *ConfigUpdateRequest) GetCompatibilityGroupOk() (*string, bool)`

GetCompatibilityGroupOk returns a tuple with the CompatibilityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityGroup

`func (o *ConfigUpdateRequest) SetCompatibilityGroup(v string)`

SetCompatibilityGroup sets CompatibilityGroup field to given value.

### HasCompatibilityGroup

`func (o *ConfigUpdateRequest) HasCompatibilityGroup() bool`

HasCompatibilityGroup returns a boolean if a field has been set.

### GetDefaultMetadata

`func (o *ConfigUpdateRequest) GetDefaultMetadata() ConfigDefaultMetadata`

GetDefaultMetadata returns the DefaultMetadata field if non-nil, zero value otherwise.

### GetDefaultMetadataOk

`func (o *ConfigUpdateRequest) GetDefaultMetadataOk() (*ConfigDefaultMetadata, bool)`

GetDefaultMetadataOk returns a tuple with the DefaultMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMetadata

`func (o *ConfigUpdateRequest) SetDefaultMetadata(v ConfigDefaultMetadata)`

SetDefaultMetadata sets DefaultMetadata field to given value.

### HasDefaultMetadata

`func (o *ConfigUpdateRequest) HasDefaultMetadata() bool`

HasDefaultMetadata returns a boolean if a field has been set.

### GetOverrideMetadata

`func (o *ConfigUpdateRequest) GetOverrideMetadata() ConfigOverrideMetadata`

GetOverrideMetadata returns the OverrideMetadata field if non-nil, zero value otherwise.

### GetOverrideMetadataOk

`func (o *ConfigUpdateRequest) GetOverrideMetadataOk() (*ConfigOverrideMetadata, bool)`

GetOverrideMetadataOk returns a tuple with the OverrideMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideMetadata

`func (o *ConfigUpdateRequest) SetOverrideMetadata(v ConfigOverrideMetadata)`

SetOverrideMetadata sets OverrideMetadata field to given value.

### HasOverrideMetadata

`func (o *ConfigUpdateRequest) HasOverrideMetadata() bool`

HasOverrideMetadata returns a boolean if a field has been set.

### GetDefaultRuleSet

`func (o *ConfigUpdateRequest) GetDefaultRuleSet() ConfigDefaultRuleSet`

GetDefaultRuleSet returns the DefaultRuleSet field if non-nil, zero value otherwise.

### GetDefaultRuleSetOk

`func (o *ConfigUpdateRequest) GetDefaultRuleSetOk() (*ConfigDefaultRuleSet, bool)`

GetDefaultRuleSetOk returns a tuple with the DefaultRuleSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRuleSet

`func (o *ConfigUpdateRequest) SetDefaultRuleSet(v ConfigDefaultRuleSet)`

SetDefaultRuleSet sets DefaultRuleSet field to given value.

### HasDefaultRuleSet

`func (o *ConfigUpdateRequest) HasDefaultRuleSet() bool`

HasDefaultRuleSet returns a boolean if a field has been set.

### GetOverrideRuleSet

`func (o *ConfigUpdateRequest) GetOverrideRuleSet() ConfigOverrideRuleSet`

GetOverrideRuleSet returns the OverrideRuleSet field if non-nil, zero value otherwise.

### GetOverrideRuleSetOk

`func (o *ConfigUpdateRequest) GetOverrideRuleSetOk() (*ConfigOverrideRuleSet, bool)`

GetOverrideRuleSetOk returns a tuple with the OverrideRuleSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideRuleSet

`func (o *ConfigUpdateRequest) SetOverrideRuleSet(v ConfigOverrideRuleSet)`

SetOverrideRuleSet sets OverrideRuleSet field to given value.

### HasOverrideRuleSet

`func (o *ConfigUpdateRequest) HasOverrideRuleSet() bool`

HasOverrideRuleSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


