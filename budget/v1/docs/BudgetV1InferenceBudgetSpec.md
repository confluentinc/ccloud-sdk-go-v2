# BudgetV1InferenceBudgetSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | Pointer to **string** | The period over which inference consumption accumulates and resets. The set of values may grow over time.  | [optional] 
**NotificationThresholds** | Pointer to **[]int32** | The percentages of the budget at which notifications fire, each a whole percentage from 1 to 100. Multiple thresholds may be configured, and each fires once as consumption crosses it. Defaults to 50, 90, and 100 percent.  | [optional] [default to [50,90,100]]
**Scope** | Pointer to **string** | The scope of resources whose inference consumption counts against this budget. The set of values may grow over time.  | [optional] 
**MaxInferenceUsd** | Pointer to **string** | The maximum inference spend permitted within the period, in US dollars as a decimal string with two fractional digits (e.g. \&quot;100.00\&quot;). A string rather than a float so the currency amount is exact.  | [optional] 
**EnforcementMode** | Pointer to **string** | The action taken when consumption reaches the cap. The set of values may grow over time.  | [optional] 

## Methods

### NewBudgetV1InferenceBudgetSpec

`func NewBudgetV1InferenceBudgetSpec() *BudgetV1InferenceBudgetSpec`

NewBudgetV1InferenceBudgetSpec instantiates a new BudgetV1InferenceBudgetSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetV1InferenceBudgetSpecWithDefaults

`func NewBudgetV1InferenceBudgetSpecWithDefaults() *BudgetV1InferenceBudgetSpec`

NewBudgetV1InferenceBudgetSpecWithDefaults instantiates a new BudgetV1InferenceBudgetSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *BudgetV1InferenceBudgetSpec) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *BudgetV1InferenceBudgetSpec) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *BudgetV1InferenceBudgetSpec) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *BudgetV1InferenceBudgetSpec) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetNotificationThresholds

`func (o *BudgetV1InferenceBudgetSpec) GetNotificationThresholds() []int32`

GetNotificationThresholds returns the NotificationThresholds field if non-nil, zero value otherwise.

### GetNotificationThresholdsOk

`func (o *BudgetV1InferenceBudgetSpec) GetNotificationThresholdsOk() (*[]int32, bool)`

GetNotificationThresholdsOk returns a tuple with the NotificationThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationThresholds

`func (o *BudgetV1InferenceBudgetSpec) SetNotificationThresholds(v []int32)`

SetNotificationThresholds sets NotificationThresholds field to given value.

### HasNotificationThresholds

`func (o *BudgetV1InferenceBudgetSpec) HasNotificationThresholds() bool`

HasNotificationThresholds returns a boolean if a field has been set.

### GetScope

`func (o *BudgetV1InferenceBudgetSpec) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *BudgetV1InferenceBudgetSpec) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *BudgetV1InferenceBudgetSpec) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *BudgetV1InferenceBudgetSpec) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetMaxInferenceUsd

`func (o *BudgetV1InferenceBudgetSpec) GetMaxInferenceUsd() string`

GetMaxInferenceUsd returns the MaxInferenceUsd field if non-nil, zero value otherwise.

### GetMaxInferenceUsdOk

`func (o *BudgetV1InferenceBudgetSpec) GetMaxInferenceUsdOk() (*string, bool)`

GetMaxInferenceUsdOk returns a tuple with the MaxInferenceUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInferenceUsd

`func (o *BudgetV1InferenceBudgetSpec) SetMaxInferenceUsd(v string)`

SetMaxInferenceUsd sets MaxInferenceUsd field to given value.

### HasMaxInferenceUsd

`func (o *BudgetV1InferenceBudgetSpec) HasMaxInferenceUsd() bool`

HasMaxInferenceUsd returns a boolean if a field has been set.

### GetEnforcementMode

`func (o *BudgetV1InferenceBudgetSpec) GetEnforcementMode() string`

GetEnforcementMode returns the EnforcementMode field if non-nil, zero value otherwise.

### GetEnforcementModeOk

`func (o *BudgetV1InferenceBudgetSpec) GetEnforcementModeOk() (*string, bool)`

GetEnforcementModeOk returns a tuple with the EnforcementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementMode

`func (o *BudgetV1InferenceBudgetSpec) SetEnforcementMode(v string)`

SetEnforcementMode sets EnforcementMode field to given value.

### HasEnforcementMode

`func (o *BudgetV1InferenceBudgetSpec) HasEnforcementMode() bool`

HasEnforcementMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


