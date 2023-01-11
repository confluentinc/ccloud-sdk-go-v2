# ByokV1AzureKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **string** | The Application ID created for this key-environment combination.  | [optional] [readonly] 
**KeyId** | **string** | The unique Key Object Identifier URL of an Azure Key Vault key.  | 
**KeyVaultId** | **string** | Key Vault ID containing the key  | 
**Kind** | **string** | BYOK kind type.  | 
**TenantId** | **string** | Tenant ID (uuid) hosting the Key Vault containing the key  | 

## Methods

### NewByokV1AzureKey

`func NewByokV1AzureKey(keyId string, keyVaultId string, kind string, tenantId string, ) *ByokV1AzureKey`

NewByokV1AzureKey instantiates a new ByokV1AzureKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewByokV1AzureKeyWithDefaults

`func NewByokV1AzureKeyWithDefaults() *ByokV1AzureKey`

NewByokV1AzureKeyWithDefaults instantiates a new ByokV1AzureKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *ByokV1AzureKey) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ByokV1AzureKey) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ByokV1AzureKey) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ByokV1AzureKey) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetKeyId

`func (o *ByokV1AzureKey) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *ByokV1AzureKey) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *ByokV1AzureKey) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetKeyVaultId

`func (o *ByokV1AzureKey) GetKeyVaultId() string`

GetKeyVaultId returns the KeyVaultId field if non-nil, zero value otherwise.

### GetKeyVaultIdOk

`func (o *ByokV1AzureKey) GetKeyVaultIdOk() (*string, bool)`

GetKeyVaultIdOk returns a tuple with the KeyVaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVaultId

`func (o *ByokV1AzureKey) SetKeyVaultId(v string)`

SetKeyVaultId sets KeyVaultId field to given value.


### GetKind

`func (o *ByokV1AzureKey) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ByokV1AzureKey) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ByokV1AzureKey) SetKind(v string)`

SetKind sets Kind field to given value.


### GetTenantId

`func (o *ByokV1AzureKey) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ByokV1AzureKey) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ByokV1AzureKey) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



### AsByokV1KeyKeyOneOf

`func (s *ByokV1AzureKey) AsByokV1KeyKeyOneOf() ByokV1KeyKeyOneOf`

Convenience method to wrap this instance of ByokV1AzureKey in ByokV1KeyKeyOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


