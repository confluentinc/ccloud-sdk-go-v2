# ConfigurationcontrolV1PolicySpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Lifecycle position of the policy.   * &#x60;DRAFT&#x60; — stored, never evaluated.   * &#x60;AUDIT&#x60; — evaluated and surfaced (logs/metrics); not enforced.   * &#x60;ACTIVE&#x60; — evaluated and enforced according to &#x60;action&#x60;.  | [optional] [default to "DRAFT"]
**Action** | **string** | Action taken when any of the policy&#39;s rules return false on an &#x60;ACTIVE&#x60; policy. Must be one the policy type allows.   * &#x60;DENY&#x60; — reject the operation.   * &#x60;OVERRIDE&#x60; — apply the failing rule&#39;s mutation to the resource.  | 
**Rules** | [**[]ConfigurationcontrolV1PolicyRule**](ConfigurationcontrolV1PolicyRule.md) | The list of CEL rules evaluated for the policy. All rules must evaluate true for the policy to pass; the platform AND-joins them at compile time.  | 

## Methods

### NewConfigurationcontrolV1PolicySpecUpdate

`func NewConfigurationcontrolV1PolicySpecUpdate(action string, rules []ConfigurationcontrolV1PolicyRule, ) *ConfigurationcontrolV1PolicySpecUpdate`

NewConfigurationcontrolV1PolicySpecUpdate instantiates a new ConfigurationcontrolV1PolicySpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationcontrolV1PolicySpecUpdateWithDefaults

`func NewConfigurationcontrolV1PolicySpecUpdateWithDefaults() *ConfigurationcontrolV1PolicySpecUpdate`

NewConfigurationcontrolV1PolicySpecUpdateWithDefaults instantiates a new ConfigurationcontrolV1PolicySpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *ConfigurationcontrolV1PolicySpecUpdate) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ConfigurationcontrolV1PolicySpecUpdate) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ConfigurationcontrolV1PolicySpecUpdate) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ConfigurationcontrolV1PolicySpecUpdate) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetAction

`func (o *ConfigurationcontrolV1PolicySpecUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ConfigurationcontrolV1PolicySpecUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ConfigurationcontrolV1PolicySpecUpdate) SetAction(v string)`

SetAction sets Action field to given value.


### GetRules

`func (o *ConfigurationcontrolV1PolicySpecUpdate) GetRules() []ConfigurationcontrolV1PolicyRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ConfigurationcontrolV1PolicySpecUpdate) GetRulesOk() (*[]ConfigurationcontrolV1PolicyRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ConfigurationcontrolV1PolicySpecUpdate) SetRules(v []ConfigurationcontrolV1PolicyRule)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


