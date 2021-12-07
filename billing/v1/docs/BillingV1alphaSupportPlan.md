# BillingV1alphaSupportPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionType** | Pointer to **string** | Support plan subscription type. | [optional] 
**EffectiveAt** | Pointer to **time.Time** | Effective date of support plan. | [optional] 

## Methods

### NewBillingV1alphaSupportPlan

`func NewBillingV1alphaSupportPlan() *BillingV1alphaSupportPlan`

NewBillingV1alphaSupportPlan instantiates a new BillingV1alphaSupportPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingV1alphaSupportPlanWithDefaults

`func NewBillingV1alphaSupportPlanWithDefaults() *BillingV1alphaSupportPlan`

NewBillingV1alphaSupportPlanWithDefaults instantiates a new BillingV1alphaSupportPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionType

`func (o *BillingV1alphaSupportPlan) GetSubscriptionType() string`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *BillingV1alphaSupportPlan) GetSubscriptionTypeOk() (*string, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *BillingV1alphaSupportPlan) SetSubscriptionType(v string)`

SetSubscriptionType sets SubscriptionType field to given value.

### HasSubscriptionType

`func (o *BillingV1alphaSupportPlan) HasSubscriptionType() bool`

HasSubscriptionType returns a boolean if a field has been set.

### GetEffectiveAt

`func (o *BillingV1alphaSupportPlan) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *BillingV1alphaSupportPlan) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *BillingV1alphaSupportPlan) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.

### HasEffectiveAt

`func (o *BillingV1alphaSupportPlan) HasEffectiveAt() bool`

HasEffectiveAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


