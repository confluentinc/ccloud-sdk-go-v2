# BudgetV1InferenceBudgetStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeriodStart** | Pointer to **time.Time** | The start of the current budget period, in RFC3339 UTC. | [optional] 
**PeriodEnd** | Pointer to **time.Time** | The end of the current budget period, in RFC3339 UTC. | [optional] 
**ConsumedUsd** | Pointer to **string** | The inference spend consumed so far in the current budget period, in US dollars as a decimal string with two fractional digits (e.g. \&quot;42.50\&quot;).  | [optional] 
**AsOf** | Pointer to **time.Time** | The date and time at which the consumption figures were computed, in RFC3339 UTC. | [optional] 
**PercentConsumed** | Pointer to **float64** | The percentage of the budget consumed so far in the current budget period. | [optional] 
**LastNotifiedPercent** | Pointer to **int32** | The highest notification threshold that has fired so far in the current budget period. | [optional] 
**LastReconciled** | Pointer to **time.Time** | The watermark up to which consumption has been reconciled against source usage records. Reconciliation runs hourly.  | [optional] 

## Methods

### NewBudgetV1InferenceBudgetStatus

`func NewBudgetV1InferenceBudgetStatus() *BudgetV1InferenceBudgetStatus`

NewBudgetV1InferenceBudgetStatus instantiates a new BudgetV1InferenceBudgetStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetV1InferenceBudgetStatusWithDefaults

`func NewBudgetV1InferenceBudgetStatusWithDefaults() *BudgetV1InferenceBudgetStatus`

NewBudgetV1InferenceBudgetStatusWithDefaults instantiates a new BudgetV1InferenceBudgetStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriodStart

`func (o *BudgetV1InferenceBudgetStatus) GetPeriodStart() time.Time`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *BudgetV1InferenceBudgetStatus) GetPeriodStartOk() (*time.Time, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *BudgetV1InferenceBudgetStatus) SetPeriodStart(v time.Time)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *BudgetV1InferenceBudgetStatus) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *BudgetV1InferenceBudgetStatus) GetPeriodEnd() time.Time`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *BudgetV1InferenceBudgetStatus) GetPeriodEndOk() (*time.Time, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *BudgetV1InferenceBudgetStatus) SetPeriodEnd(v time.Time)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *BudgetV1InferenceBudgetStatus) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetConsumedUsd

`func (o *BudgetV1InferenceBudgetStatus) GetConsumedUsd() string`

GetConsumedUsd returns the ConsumedUsd field if non-nil, zero value otherwise.

### GetConsumedUsdOk

`func (o *BudgetV1InferenceBudgetStatus) GetConsumedUsdOk() (*string, bool)`

GetConsumedUsdOk returns a tuple with the ConsumedUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedUsd

`func (o *BudgetV1InferenceBudgetStatus) SetConsumedUsd(v string)`

SetConsumedUsd sets ConsumedUsd field to given value.

### HasConsumedUsd

`func (o *BudgetV1InferenceBudgetStatus) HasConsumedUsd() bool`

HasConsumedUsd returns a boolean if a field has been set.

### GetAsOf

`func (o *BudgetV1InferenceBudgetStatus) GetAsOf() time.Time`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *BudgetV1InferenceBudgetStatus) GetAsOfOk() (*time.Time, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *BudgetV1InferenceBudgetStatus) SetAsOf(v time.Time)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *BudgetV1InferenceBudgetStatus) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetPercentConsumed

`func (o *BudgetV1InferenceBudgetStatus) GetPercentConsumed() float64`

GetPercentConsumed returns the PercentConsumed field if non-nil, zero value otherwise.

### GetPercentConsumedOk

`func (o *BudgetV1InferenceBudgetStatus) GetPercentConsumedOk() (*float64, bool)`

GetPercentConsumedOk returns a tuple with the PercentConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentConsumed

`func (o *BudgetV1InferenceBudgetStatus) SetPercentConsumed(v float64)`

SetPercentConsumed sets PercentConsumed field to given value.

### HasPercentConsumed

`func (o *BudgetV1InferenceBudgetStatus) HasPercentConsumed() bool`

HasPercentConsumed returns a boolean if a field has been set.

### GetLastNotifiedPercent

`func (o *BudgetV1InferenceBudgetStatus) GetLastNotifiedPercent() int32`

GetLastNotifiedPercent returns the LastNotifiedPercent field if non-nil, zero value otherwise.

### GetLastNotifiedPercentOk

`func (o *BudgetV1InferenceBudgetStatus) GetLastNotifiedPercentOk() (*int32, bool)`

GetLastNotifiedPercentOk returns a tuple with the LastNotifiedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNotifiedPercent

`func (o *BudgetV1InferenceBudgetStatus) SetLastNotifiedPercent(v int32)`

SetLastNotifiedPercent sets LastNotifiedPercent field to given value.

### HasLastNotifiedPercent

`func (o *BudgetV1InferenceBudgetStatus) HasLastNotifiedPercent() bool`

HasLastNotifiedPercent returns a boolean if a field has been set.

### GetLastReconciled

`func (o *BudgetV1InferenceBudgetStatus) GetLastReconciled() time.Time`

GetLastReconciled returns the LastReconciled field if non-nil, zero value otherwise.

### GetLastReconciledOk

`func (o *BudgetV1InferenceBudgetStatus) GetLastReconciledOk() (*time.Time, bool)`

GetLastReconciledOk returns a tuple with the LastReconciled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReconciled

`func (o *BudgetV1InferenceBudgetStatus) SetLastReconciled(v time.Time)`

SetLastReconciled sets LastReconciled field to given value.

### HasLastReconciled

`func (o *BudgetV1InferenceBudgetStatus) HasLastReconciled() bool`

HasLastReconciled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


