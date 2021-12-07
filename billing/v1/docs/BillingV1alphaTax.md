# BillingV1alphaTax

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line1** | **string** | Address line 1 (e.g., street, PO Box, or company name). | 
**Line2** | Pointer to **string** | Address line 2 (e.g., apartment, suite, unit, or building). | [optional] 
**City** | **string** | City, district, suburb, town, or village.. | 
**State** | **string** | State, county, province, or region. | 
**Country** | **string** | Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)). | 
**PostalCode** | **string** | ZIP or postal code. | 
**TaxIds** | Pointer to [**[]BillingV1alphaTaxId**](BillingV1alphaTaxId.md) | The list of tax ID objects. | [optional] 

## Methods

### NewBillingV1alphaTax

`func NewBillingV1alphaTax(line1 string, city string, state string, country string, postalCode string, ) *BillingV1alphaTax`

NewBillingV1alphaTax instantiates a new BillingV1alphaTax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingV1alphaTaxWithDefaults

`func NewBillingV1alphaTaxWithDefaults() *BillingV1alphaTax`

NewBillingV1alphaTaxWithDefaults instantiates a new BillingV1alphaTax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine1

`func (o *BillingV1alphaTax) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *BillingV1alphaTax) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *BillingV1alphaTax) SetLine1(v string)`

SetLine1 sets Line1 field to given value.


### GetLine2

`func (o *BillingV1alphaTax) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *BillingV1alphaTax) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *BillingV1alphaTax) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *BillingV1alphaTax) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### GetCity

`func (o *BillingV1alphaTax) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *BillingV1alphaTax) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *BillingV1alphaTax) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *BillingV1alphaTax) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BillingV1alphaTax) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BillingV1alphaTax) SetState(v string)`

SetState sets State field to given value.


### GetCountry

`func (o *BillingV1alphaTax) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BillingV1alphaTax) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BillingV1alphaTax) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetPostalCode

`func (o *BillingV1alphaTax) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *BillingV1alphaTax) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *BillingV1alphaTax) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetTaxIds

`func (o *BillingV1alphaTax) GetTaxIds() []BillingV1alphaTaxId`

GetTaxIds returns the TaxIds field if non-nil, zero value otherwise.

### GetTaxIdsOk

`func (o *BillingV1alphaTax) GetTaxIdsOk() (*[]BillingV1alphaTaxId, bool)`

GetTaxIdsOk returns a tuple with the TaxIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIds

`func (o *BillingV1alphaTax) SetTaxIds(v []BillingV1alphaTaxId)`

SetTaxIds sets TaxIds field to given value.

### HasTaxIds

`func (o *BillingV1alphaTax) HasTaxIds() bool`

HasTaxIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


