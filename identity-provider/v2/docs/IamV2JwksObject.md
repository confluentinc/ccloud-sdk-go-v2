# IamV2JwksObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kty** | **string** | Specifies the cryptographic algorithm family used with the key | 
**Kid** | **string** | Specifies the key-id issued by the OpenIDProvider for the particular tenant | 
**Alg** | **string** | Specifies the algorithm to be used to generate the public key | 
**Use** | Pointer to **string** | Specifies the intended usage of the key | [optional] 
**N** | Pointer to **string** | Specifies the modulus of the RSA public key. Represented as a Base64urlUInt-encoded value | [optional] 
**E** | Pointer to **string** | Specifies the exponent of the RSA public key. | [optional] 

## Methods

### NewIamV2JwksObject

`func NewIamV2JwksObject(kty string, kid string, alg string, ) *IamV2JwksObject`

NewIamV2JwksObject instantiates a new IamV2JwksObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2JwksObjectWithDefaults

`func NewIamV2JwksObjectWithDefaults() *IamV2JwksObject`

NewIamV2JwksObjectWithDefaults instantiates a new IamV2JwksObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKty

`func (o *IamV2JwksObject) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *IamV2JwksObject) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *IamV2JwksObject) SetKty(v string)`

SetKty sets Kty field to given value.


### GetKid

`func (o *IamV2JwksObject) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *IamV2JwksObject) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *IamV2JwksObject) SetKid(v string)`

SetKid sets Kid field to given value.


### GetAlg

`func (o *IamV2JwksObject) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *IamV2JwksObject) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *IamV2JwksObject) SetAlg(v string)`

SetAlg sets Alg field to given value.


### GetUse

`func (o *IamV2JwksObject) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *IamV2JwksObject) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *IamV2JwksObject) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *IamV2JwksObject) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetN

`func (o *IamV2JwksObject) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *IamV2JwksObject) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *IamV2JwksObject) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *IamV2JwksObject) HasN() bool`

HasN returns a boolean if a field has been set.

### GetE

`func (o *IamV2JwksObject) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *IamV2JwksObject) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *IamV2JwksObject) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *IamV2JwksObject) HasE() bool`

HasE returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


