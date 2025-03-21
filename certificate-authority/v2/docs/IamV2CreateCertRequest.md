# IamV2CreateCertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**DisplayName** | Pointer to **string** | The human-readable name of the certificate authority. | [optional] 
**Description** | Pointer to **string** | A description of the certificate authority. | [optional] 
**CertificateChain** | Pointer to **string** | The PEM encoded string containing the signing certificate chain used to validate client certs. | [optional] 
**CertificateChainFilename** | Pointer to **string** | The name of the certificate file. | [optional] 
**CrlUrl** | Pointer to **string** | The url from which to fetch the CRL for the certificate authority if crl_source is URL. | [optional] 
**CrlChain** | Pointer to **string** | The PEM encoded string containing the CRL for this certificate authority. Defaults to this over &#x60;crl_url&#x60; if available. | [optional] 

## Methods

### NewIamV2CreateCertRequest

`func NewIamV2CreateCertRequest() *IamV2CreateCertRequest`

NewIamV2CreateCertRequest instantiates a new IamV2CreateCertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2CreateCertRequestWithDefaults

`func NewIamV2CreateCertRequestWithDefaults() *IamV2CreateCertRequest`

NewIamV2CreateCertRequestWithDefaults instantiates a new IamV2CreateCertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IamV2CreateCertRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IamV2CreateCertRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IamV2CreateCertRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IamV2CreateCertRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IamV2CreateCertRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IamV2CreateCertRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IamV2CreateCertRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IamV2CreateCertRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *IamV2CreateCertRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamV2CreateCertRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamV2CreateCertRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamV2CreateCertRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *IamV2CreateCertRequest) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamV2CreateCertRequest) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamV2CreateCertRequest) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IamV2CreateCertRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamV2CreateCertRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamV2CreateCertRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamV2CreateCertRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamV2CreateCertRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *IamV2CreateCertRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamV2CreateCertRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamV2CreateCertRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamV2CreateCertRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCertificateChain

`func (o *IamV2CreateCertRequest) GetCertificateChain() string`

GetCertificateChain returns the CertificateChain field if non-nil, zero value otherwise.

### GetCertificateChainOk

`func (o *IamV2CreateCertRequest) GetCertificateChainOk() (*string, bool)`

GetCertificateChainOk returns a tuple with the CertificateChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateChain

`func (o *IamV2CreateCertRequest) SetCertificateChain(v string)`

SetCertificateChain sets CertificateChain field to given value.

### HasCertificateChain

`func (o *IamV2CreateCertRequest) HasCertificateChain() bool`

HasCertificateChain returns a boolean if a field has been set.

### GetCertificateChainFilename

`func (o *IamV2CreateCertRequest) GetCertificateChainFilename() string`

GetCertificateChainFilename returns the CertificateChainFilename field if non-nil, zero value otherwise.

### GetCertificateChainFilenameOk

`func (o *IamV2CreateCertRequest) GetCertificateChainFilenameOk() (*string, bool)`

GetCertificateChainFilenameOk returns a tuple with the CertificateChainFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateChainFilename

`func (o *IamV2CreateCertRequest) SetCertificateChainFilename(v string)`

SetCertificateChainFilename sets CertificateChainFilename field to given value.

### HasCertificateChainFilename

`func (o *IamV2CreateCertRequest) HasCertificateChainFilename() bool`

HasCertificateChainFilename returns a boolean if a field has been set.

### GetCrlUrl

`func (o *IamV2CreateCertRequest) GetCrlUrl() string`

GetCrlUrl returns the CrlUrl field if non-nil, zero value otherwise.

### GetCrlUrlOk

`func (o *IamV2CreateCertRequest) GetCrlUrlOk() (*string, bool)`

GetCrlUrlOk returns a tuple with the CrlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlUrl

`func (o *IamV2CreateCertRequest) SetCrlUrl(v string)`

SetCrlUrl sets CrlUrl field to given value.

### HasCrlUrl

`func (o *IamV2CreateCertRequest) HasCrlUrl() bool`

HasCrlUrl returns a boolean if a field has been set.

### GetCrlChain

`func (o *IamV2CreateCertRequest) GetCrlChain() string`

GetCrlChain returns the CrlChain field if non-nil, zero value otherwise.

### GetCrlChainOk

`func (o *IamV2CreateCertRequest) GetCrlChainOk() (*string, bool)`

GetCrlChainOk returns a tuple with the CrlChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlChain

`func (o *IamV2CreateCertRequest) SetCrlChain(v string)`

SetCrlChain sets CrlChain field to given value.

### HasCrlChain

`func (o *IamV2CreateCertRequest) HasCrlChain() bool`

HasCrlChain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


