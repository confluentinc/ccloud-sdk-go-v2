# BillingV1alphaPromoCodeClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | Pointer to **float64** | Remaining amount of promo code in dollars. | [optional] 
**ClaimedAt** | Pointer to **time.Time** | Claim date of promo code. | [optional] 
**ExpiresAt** | Pointer to **time.Time** | Expiration date of promo code. | [optional] 

## Methods

### NewBillingV1alphaPromoCodeClaim

`func NewBillingV1alphaPromoCodeClaim() *BillingV1alphaPromoCodeClaim`

NewBillingV1alphaPromoCodeClaim instantiates a new BillingV1alphaPromoCodeClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingV1alphaPromoCodeClaimWithDefaults

`func NewBillingV1alphaPromoCodeClaimWithDefaults() *BillingV1alphaPromoCodeClaim`

NewBillingV1alphaPromoCodeClaimWithDefaults instantiates a new BillingV1alphaPromoCodeClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *BillingV1alphaPromoCodeClaim) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *BillingV1alphaPromoCodeClaim) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *BillingV1alphaPromoCodeClaim) SetBalance(v float64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *BillingV1alphaPromoCodeClaim) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetClaimedAt

`func (o *BillingV1alphaPromoCodeClaim) GetClaimedAt() time.Time`

GetClaimedAt returns the ClaimedAt field if non-nil, zero value otherwise.

### GetClaimedAtOk

`func (o *BillingV1alphaPromoCodeClaim) GetClaimedAtOk() (*time.Time, bool)`

GetClaimedAtOk returns a tuple with the ClaimedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedAt

`func (o *BillingV1alphaPromoCodeClaim) SetClaimedAt(v time.Time)`

SetClaimedAt sets ClaimedAt field to given value.

### HasClaimedAt

`func (o *BillingV1alphaPromoCodeClaim) HasClaimedAt() bool`

HasClaimedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *BillingV1alphaPromoCodeClaim) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *BillingV1alphaPromoCodeClaim) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *BillingV1alphaPromoCodeClaim) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *BillingV1alphaPromoCodeClaim) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


