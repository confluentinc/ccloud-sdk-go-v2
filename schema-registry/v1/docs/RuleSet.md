# RuleSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MigrationRules** | Pointer to [**[]Rule**](Rule.md) |  | [optional] 
**DomainRules** | Pointer to [**[]Rule**](Rule.md) |  | [optional] 

## Methods

### NewRuleSet

`func NewRuleSet() *RuleSet`

NewRuleSet instantiates a new RuleSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleSetWithDefaults

`func NewRuleSetWithDefaults() *RuleSet`

NewRuleSetWithDefaults instantiates a new RuleSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMigrationRules

`func (o *RuleSet) GetMigrationRules() []Rule`

GetMigrationRules returns the MigrationRules field if non-nil, zero value otherwise.

### GetMigrationRulesOk

`func (o *RuleSet) GetMigrationRulesOk() (*[]Rule, bool)`

GetMigrationRulesOk returns a tuple with the MigrationRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationRules

`func (o *RuleSet) SetMigrationRules(v []Rule)`

SetMigrationRules sets MigrationRules field to given value.

### HasMigrationRules

`func (o *RuleSet) HasMigrationRules() bool`

HasMigrationRules returns a boolean if a field has been set.

### GetDomainRules

`func (o *RuleSet) GetDomainRules() []Rule`

GetDomainRules returns the DomainRules field if non-nil, zero value otherwise.

### GetDomainRulesOk

`func (o *RuleSet) GetDomainRulesOk() (*[]Rule, bool)`

GetDomainRulesOk returns a tuple with the DomainRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainRules

`func (o *RuleSet) SetDomainRules(v []Rule)`

SetDomainRules sets DomainRules field to given value.

### HasDomainRules

`func (o *RuleSet) HasDomainRules() bool`

HasDomainRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


