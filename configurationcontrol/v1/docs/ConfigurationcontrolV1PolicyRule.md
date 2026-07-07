# ConfigurationcontrolV1PolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable label for the rule. Surfaced in errors, audit logs, and metrics so operators can identify which rule fired.  | 
**Match** | **string** | CEL expression returning a boolean that describes the required configuration. The rule passes when &#x60;match&#x60; returns true. When &#x60;match&#x60; returns false on a &#x60;DENY&#x60; policy the operation is rejected; on an &#x60;OVERRIDE&#x60; policy &#x60;mutate&#x60; is applied.  | 
**Mutate** | Pointer to **string** | CEL expression returning the mutation payload to apply when &#x60;match&#x60; returns false under an &#x60;OVERRIDE&#x60; policy. The output type must match a mutable variable binding declared by the policy type. Required on &#x60;OVERRIDE&#x60; policies; must be omitted (or the literal &#x60;false&#x60;) on &#x60;DENY&#x60; policies.  | [optional] 

## Methods

### NewConfigurationcontrolV1PolicyRule

`func NewConfigurationcontrolV1PolicyRule(name string, match string, ) *ConfigurationcontrolV1PolicyRule`

NewConfigurationcontrolV1PolicyRule instantiates a new ConfigurationcontrolV1PolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationcontrolV1PolicyRuleWithDefaults

`func NewConfigurationcontrolV1PolicyRuleWithDefaults() *ConfigurationcontrolV1PolicyRule`

NewConfigurationcontrolV1PolicyRuleWithDefaults instantiates a new ConfigurationcontrolV1PolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConfigurationcontrolV1PolicyRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigurationcontrolV1PolicyRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigurationcontrolV1PolicyRule) SetName(v string)`

SetName sets Name field to given value.


### GetMatch

`func (o *ConfigurationcontrolV1PolicyRule) GetMatch() string`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *ConfigurationcontrolV1PolicyRule) GetMatchOk() (*string, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *ConfigurationcontrolV1PolicyRule) SetMatch(v string)`

SetMatch sets Match field to given value.


### GetMutate

`func (o *ConfigurationcontrolV1PolicyRule) GetMutate() string`

GetMutate returns the Mutate field if non-nil, zero value otherwise.

### GetMutateOk

`func (o *ConfigurationcontrolV1PolicyRule) GetMutateOk() (*string, bool)`

GetMutateOk returns a tuple with the Mutate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutate

`func (o *ConfigurationcontrolV1PolicyRule) SetMutate(v string)`

SetMutate sets Mutate field to given value.

### HasMutate

`func (o *ConfigurationcontrolV1PolicyRule) HasMutate() bool`

HasMutate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


