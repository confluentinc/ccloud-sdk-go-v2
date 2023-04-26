# BillingV1alphaProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The organization&#39;s billing email address. | [optional] 
**Tax** | Pointer to [**BillingV1alphaTax**](billing.v1alpha.Tax.md) | The organization&#39;s tax information. | [optional] 

## Methods

### NewBillingV1alphaProfile

`func NewBillingV1alphaProfile() *BillingV1alphaProfile`

NewBillingV1alphaProfile instantiates a new BillingV1alphaProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingV1alphaProfileWithDefaults

`func NewBillingV1alphaProfileWithDefaults() *BillingV1alphaProfile`

NewBillingV1alphaProfileWithDefaults instantiates a new BillingV1alphaProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *BillingV1alphaProfile) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BillingV1alphaProfile) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BillingV1alphaProfile) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BillingV1alphaProfile) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetTax

`func (o *BillingV1alphaProfile) GetTax() BillingV1alphaTax`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *BillingV1alphaProfile) GetTaxOk() (*BillingV1alphaTax, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *BillingV1alphaProfile) SetTax(v BillingV1alphaTax)`

SetTax sets Tax field to given value.

### HasTax

`func (o *BillingV1alphaProfile) HasTax() bool`

HasTax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


