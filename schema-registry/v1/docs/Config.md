# Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | If alias is specified, then this subject is an alias for the subject named by the alias. That means that any reference to this subject will be replaced by the alias. | [optional] 
**Normalize** | Pointer to **bool** | If true, then schemas are automatically normalized when registered or when passed during lookups. This means that clients do not have to pass the \&quot;normalize\&quot; query parameter to have normalization occur. | [optional] 
**CompatibilityLevel** | Pointer to **string** | Compatibility Level | [optional] 
**CompatibilityGroup** | Pointer to **string** | Only schemas that belong to the same compatibility group will be checked for compatibility. | [optional] 
**DefaultMetadata** | Pointer to [**ConfigDefaultMetadata**](ConfigDefaultMetadata.md) |  | [optional] 
**OverrideMetadata** | Pointer to [**ConfigOverrideMetadata**](ConfigOverrideMetadata.md) |  | [optional] 
**DefaultRuleSet** | Pointer to [**ConfigDefaultRuleSet**](ConfigDefaultRuleSet.md) |  | [optional] 
**OverrideRuleSet** | Pointer to [**ConfigOverrideRuleSet**](ConfigOverrideRuleSet.md) |  | [optional] 

## Methods

### NewConfig

`func NewConfig() *Config`

NewConfig instantiates a new Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigWithDefaults

`func NewConfigWithDefaults() *Config`

NewConfigWithDefaults instantiates a new Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *Config) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *Config) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *Config) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *Config) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetNormalize

`func (o *Config) GetNormalize() bool`

GetNormalize returns the Normalize field if non-nil, zero value otherwise.

### GetNormalizeOk

`func (o *Config) GetNormalizeOk() (*bool, bool)`

GetNormalizeOk returns a tuple with the Normalize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalize

`func (o *Config) SetNormalize(v bool)`

SetNormalize sets Normalize field to given value.

### HasNormalize

`func (o *Config) HasNormalize() bool`

HasNormalize returns a boolean if a field has been set.

### GetCompatibilityLevel

`func (o *Config) GetCompatibilityLevel() string`

GetCompatibilityLevel returns the CompatibilityLevel field if non-nil, zero value otherwise.

### GetCompatibilityLevelOk

`func (o *Config) GetCompatibilityLevelOk() (*string, bool)`

GetCompatibilityLevelOk returns a tuple with the CompatibilityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityLevel

`func (o *Config) SetCompatibilityLevel(v string)`

SetCompatibilityLevel sets CompatibilityLevel field to given value.

### HasCompatibilityLevel

`func (o *Config) HasCompatibilityLevel() bool`

HasCompatibilityLevel returns a boolean if a field has been set.

### GetCompatibilityGroup

`func (o *Config) GetCompatibilityGroup() string`

GetCompatibilityGroup returns the CompatibilityGroup field if non-nil, zero value otherwise.

### GetCompatibilityGroupOk

`func (o *Config) GetCompatibilityGroupOk() (*string, bool)`

GetCompatibilityGroupOk returns a tuple with the CompatibilityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityGroup

`func (o *Config) SetCompatibilityGroup(v string)`

SetCompatibilityGroup sets CompatibilityGroup field to given value.

### HasCompatibilityGroup

`func (o *Config) HasCompatibilityGroup() bool`

HasCompatibilityGroup returns a boolean if a field has been set.

### GetDefaultMetadata

`func (o *Config) GetDefaultMetadata() ConfigDefaultMetadata`

GetDefaultMetadata returns the DefaultMetadata field if non-nil, zero value otherwise.

### GetDefaultMetadataOk

`func (o *Config) GetDefaultMetadataOk() (*ConfigDefaultMetadata, bool)`

GetDefaultMetadataOk returns a tuple with the DefaultMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMetadata

`func (o *Config) SetDefaultMetadata(v ConfigDefaultMetadata)`

SetDefaultMetadata sets DefaultMetadata field to given value.

### HasDefaultMetadata

`func (o *Config) HasDefaultMetadata() bool`

HasDefaultMetadata returns a boolean if a field has been set.

### GetOverrideMetadata

`func (o *Config) GetOverrideMetadata() ConfigOverrideMetadata`

GetOverrideMetadata returns the OverrideMetadata field if non-nil, zero value otherwise.

### GetOverrideMetadataOk

`func (o *Config) GetOverrideMetadataOk() (*ConfigOverrideMetadata, bool)`

GetOverrideMetadataOk returns a tuple with the OverrideMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideMetadata

`func (o *Config) SetOverrideMetadata(v ConfigOverrideMetadata)`

SetOverrideMetadata sets OverrideMetadata field to given value.

### HasOverrideMetadata

`func (o *Config) HasOverrideMetadata() bool`

HasOverrideMetadata returns a boolean if a field has been set.

### GetDefaultRuleSet

`func (o *Config) GetDefaultRuleSet() ConfigDefaultRuleSet`

GetDefaultRuleSet returns the DefaultRuleSet field if non-nil, zero value otherwise.

### GetDefaultRuleSetOk

`func (o *Config) GetDefaultRuleSetOk() (*ConfigDefaultRuleSet, bool)`

GetDefaultRuleSetOk returns a tuple with the DefaultRuleSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRuleSet

`func (o *Config) SetDefaultRuleSet(v ConfigDefaultRuleSet)`

SetDefaultRuleSet sets DefaultRuleSet field to given value.

### HasDefaultRuleSet

`func (o *Config) HasDefaultRuleSet() bool`

HasDefaultRuleSet returns a boolean if a field has been set.

### GetOverrideRuleSet

`func (o *Config) GetOverrideRuleSet() ConfigOverrideRuleSet`

GetOverrideRuleSet returns the OverrideRuleSet field if non-nil, zero value otherwise.

### GetOverrideRuleSetOk

`func (o *Config) GetOverrideRuleSetOk() (*ConfigOverrideRuleSet, bool)`

GetOverrideRuleSetOk returns a tuple with the OverrideRuleSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideRuleSet

`func (o *Config) SetOverrideRuleSet(v ConfigOverrideRuleSet)`

SetOverrideRuleSet sets OverrideRuleSet field to given value.

### HasOverrideRuleSet

`func (o *Config) HasOverrideRuleSet() bool`

HasOverrideRuleSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


