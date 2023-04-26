# BillingV1alphaProductTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**BillingV1alphaProduct**](billing.v1alpha.Product.md) | Billable product details. | [optional] 
**Tier** | Pointer to [**BillingV1alphaTier**](billing.v1alpha.Tier.md) | Billable tier details. | [optional] 
**EffectiveAt** | Pointer to **time.Time** | effective date for billing the selected product tier. | [optional] 

## Methods

### NewBillingV1alphaProductTier

`func NewBillingV1alphaProductTier() *BillingV1alphaProductTier`

NewBillingV1alphaProductTier instantiates a new BillingV1alphaProductTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingV1alphaProductTierWithDefaults

`func NewBillingV1alphaProductTierWithDefaults() *BillingV1alphaProductTier`

NewBillingV1alphaProductTierWithDefaults instantiates a new BillingV1alphaProductTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *BillingV1alphaProductTier) GetProduct() BillingV1alphaProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *BillingV1alphaProductTier) GetProductOk() (*BillingV1alphaProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *BillingV1alphaProductTier) SetProduct(v BillingV1alphaProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *BillingV1alphaProductTier) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetTier

`func (o *BillingV1alphaProductTier) GetTier() BillingV1alphaTier`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *BillingV1alphaProductTier) GetTierOk() (*BillingV1alphaTier, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *BillingV1alphaProductTier) SetTier(v BillingV1alphaTier)`

SetTier sets Tier field to given value.

### HasTier

`func (o *BillingV1alphaProductTier) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetEffectiveAt

`func (o *BillingV1alphaProductTier) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *BillingV1alphaProductTier) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *BillingV1alphaProductTier) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.

### HasEffectiveAt

`func (o *BillingV1alphaProductTier) HasEffectiveAt() bool`

HasEffectiveAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


