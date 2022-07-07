# IamV2Jwks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kty** | **string** | Specifies the cryptographic algorithm family used with the key | 
**Kid** | **string** | Specifies the key-id issued by the OpenIDProvider for the particular tenant | 
**Alg** | **string** | Specifies the algorithm to be used to generate the public key | 
**Use** | Pointer to **string** | Specifies the intended use of the public key | [optional] 
**KeyOps** | Pointer to **string** | Identifies the operation(s) for which the key is intended to be used | [optional] 

## Methods

### NewIamV2Jwks

`func NewIamV2Jwks(kty string, kid string, alg string, ) *IamV2Jwks`

NewIamV2Jwks instantiates a new IamV2Jwks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2JwksWithDefaults

`func NewIamV2JwksWithDefaults() *IamV2Jwks`

NewIamV2JwksWithDefaults instantiates a new IamV2Jwks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKty

`func (o *IamV2Jwks) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *IamV2Jwks) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *IamV2Jwks) SetKty(v string)`

SetKty sets Kty field to given value.


### GetKid

`func (o *IamV2Jwks) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *IamV2Jwks) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *IamV2Jwks) SetKid(v string)`

SetKid sets Kid field to given value.


### GetAlg

`func (o *IamV2Jwks) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *IamV2Jwks) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *IamV2Jwks) SetAlg(v string)`

SetAlg sets Alg field to given value.


### GetUse

`func (o *IamV2Jwks) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *IamV2Jwks) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *IamV2Jwks) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *IamV2Jwks) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetKeyOps

`func (o *IamV2Jwks) GetKeyOps() string`

GetKeyOps returns the KeyOps field if non-nil, zero value otherwise.

### GetKeyOpsOk

`func (o *IamV2Jwks) GetKeyOpsOk() (*string, bool)`

GetKeyOpsOk returns a tuple with the KeyOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyOps

`func (o *IamV2Jwks) SetKeyOps(v string)`

SetKeyOps sets KeyOps field to given value.

### HasKeyOps

`func (o *IamV2Jwks) HasKeyOps() bool`

HasKeyOps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


