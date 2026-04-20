# CreateDekRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **string** | Subject of the dek | [optional] 
**Version** | Pointer to **int32** | Version of the dek | [optional] 
**Algorithm** | Pointer to **string** | Algorithm of the dek | [optional] 
**EncryptedKeyMaterial** | Pointer to **string** | Encrypted key material of the dek | [optional] 
**Deleted** | Pointer to **bool** | Whether the dek is deleted | [optional] 

## Methods

### NewCreateDekRequest

`func NewCreateDekRequest() *CreateDekRequest`

NewCreateDekRequest instantiates a new CreateDekRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDekRequestWithDefaults

`func NewCreateDekRequestWithDefaults() *CreateDekRequest`

NewCreateDekRequestWithDefaults instantiates a new CreateDekRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *CreateDekRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CreateDekRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CreateDekRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CreateDekRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetVersion

`func (o *CreateDekRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateDekRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateDekRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CreateDekRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAlgorithm

`func (o *CreateDekRequest) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *CreateDekRequest) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *CreateDekRequest) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *CreateDekRequest) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetEncryptedKeyMaterial

`func (o *CreateDekRequest) GetEncryptedKeyMaterial() string`

GetEncryptedKeyMaterial returns the EncryptedKeyMaterial field if non-nil, zero value otherwise.

### GetEncryptedKeyMaterialOk

`func (o *CreateDekRequest) GetEncryptedKeyMaterialOk() (*string, bool)`

GetEncryptedKeyMaterialOk returns a tuple with the EncryptedKeyMaterial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedKeyMaterial

`func (o *CreateDekRequest) SetEncryptedKeyMaterial(v string)`

SetEncryptedKeyMaterial sets EncryptedKeyMaterial field to given value.

### HasEncryptedKeyMaterial

`func (o *CreateDekRequest) HasEncryptedKeyMaterial() bool`

HasEncryptedKeyMaterial returns a boolean if a field has been set.

### GetDeleted

`func (o *CreateDekRequest) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *CreateDekRequest) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *CreateDekRequest) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *CreateDekRequest) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


