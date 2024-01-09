# CreateKekRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**KmsType** | Pointer to **string** |  | [optional] 
**KmsKeyId** | Pointer to **string** |  | [optional] 
**KmsProps** | Pointer to **map[string]string** |  | [optional] 
**Doc** | Pointer to **string** |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateKekRequest

`func NewCreateKekRequest() *CreateKekRequest`

NewCreateKekRequest instantiates a new CreateKekRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateKekRequestWithDefaults

`func NewCreateKekRequestWithDefaults() *CreateKekRequest`

NewCreateKekRequestWithDefaults instantiates a new CreateKekRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateKekRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateKekRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateKekRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateKekRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKmsType

`func (o *CreateKekRequest) GetKmsType() string`

GetKmsType returns the KmsType field if non-nil, zero value otherwise.

### GetKmsTypeOk

`func (o *CreateKekRequest) GetKmsTypeOk() (*string, bool)`

GetKmsTypeOk returns a tuple with the KmsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsType

`func (o *CreateKekRequest) SetKmsType(v string)`

SetKmsType sets KmsType field to given value.

### HasKmsType

`func (o *CreateKekRequest) HasKmsType() bool`

HasKmsType returns a boolean if a field has been set.

### GetKmsKeyId

`func (o *CreateKekRequest) GetKmsKeyId() string`

GetKmsKeyId returns the KmsKeyId field if non-nil, zero value otherwise.

### GetKmsKeyIdOk

`func (o *CreateKekRequest) GetKmsKeyIdOk() (*string, bool)`

GetKmsKeyIdOk returns a tuple with the KmsKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsKeyId

`func (o *CreateKekRequest) SetKmsKeyId(v string)`

SetKmsKeyId sets KmsKeyId field to given value.

### HasKmsKeyId

`func (o *CreateKekRequest) HasKmsKeyId() bool`

HasKmsKeyId returns a boolean if a field has been set.

### GetKmsProps

`func (o *CreateKekRequest) GetKmsProps() map[string]string`

GetKmsProps returns the KmsProps field if non-nil, zero value otherwise.

### GetKmsPropsOk

`func (o *CreateKekRequest) GetKmsPropsOk() (*map[string]string, bool)`

GetKmsPropsOk returns a tuple with the KmsProps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsProps

`func (o *CreateKekRequest) SetKmsProps(v map[string]string)`

SetKmsProps sets KmsProps field to given value.

### HasKmsProps

`func (o *CreateKekRequest) HasKmsProps() bool`

HasKmsProps returns a boolean if a field has been set.

### GetDoc

`func (o *CreateKekRequest) GetDoc() string`

GetDoc returns the Doc field if non-nil, zero value otherwise.

### GetDocOk

`func (o *CreateKekRequest) GetDocOk() (*string, bool)`

GetDocOk returns a tuple with the Doc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoc

`func (o *CreateKekRequest) SetDoc(v string)`

SetDoc sets Doc field to given value.

### HasDoc

`func (o *CreateKekRequest) HasDoc() bool`

HasDoc returns a boolean if a field has been set.

### GetShared

`func (o *CreateKekRequest) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *CreateKekRequest) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *CreateKekRequest) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *CreateKekRequest) HasShared() bool`

HasShared returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


