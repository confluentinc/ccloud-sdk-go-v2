# BillingV1alphaPromoCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The promotional code. | [optional] 
**Amount** | Pointer to **float64** | Total value of promo code in dollars. | [optional] 
**Claim** | Pointer to [**BillingV1alphaPromoCodeClaim**](billing.v1alpha.PromoCodeClaim.md) | The promo code&#39;s claim details. | [optional] 

## Methods

### NewBillingV1alphaPromoCode

`func NewBillingV1alphaPromoCode() *BillingV1alphaPromoCode`

NewBillingV1alphaPromoCode instantiates a new BillingV1alphaPromoCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingV1alphaPromoCodeWithDefaults

`func NewBillingV1alphaPromoCodeWithDefaults() *BillingV1alphaPromoCode`

NewBillingV1alphaPromoCodeWithDefaults instantiates a new BillingV1alphaPromoCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *BillingV1alphaPromoCode) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BillingV1alphaPromoCode) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BillingV1alphaPromoCode) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *BillingV1alphaPromoCode) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetAmount

`func (o *BillingV1alphaPromoCode) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BillingV1alphaPromoCode) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BillingV1alphaPromoCode) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BillingV1alphaPromoCode) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetClaim

`func (o *BillingV1alphaPromoCode) GetClaim() BillingV1alphaPromoCodeClaim`

GetClaim returns the Claim field if non-nil, zero value otherwise.

### GetClaimOk

`func (o *BillingV1alphaPromoCode) GetClaimOk() (*BillingV1alphaPromoCodeClaim, bool)`

GetClaimOk returns a tuple with the Claim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaim

`func (o *BillingV1alphaPromoCode) SetClaim(v BillingV1alphaPromoCodeClaim)`

SetClaim sets Claim field to given value.

### HasClaim

`func (o *BillingV1alphaPromoCode) HasClaim() bool`

HasClaim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


