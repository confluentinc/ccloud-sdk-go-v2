# IamV2UpdateCertRequest

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
**CertificateChainFilename** | Pointer to **string** | The name of the certificate file. Must be set if certificate is updated. | [optional] 
**CrlUrl** | Pointer to **string** | The url from which to fetch the CRL for the certificate authority if crl_source is URL. | [optional] 
**CrlChain** | Pointer to **string** | The PEM encoded string containing the CRL for this certificate authority. Defaults to this over &#x60;crl_url&#x60; if available. | [optional] 

## Methods

### NewIamV2UpdateCertRequest

`func NewIamV2UpdateCertRequest() *IamV2UpdateCertRequest`

NewIamV2UpdateCertRequest instantiates a new IamV2UpdateCertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2UpdateCertRequestWithDefaults

`func NewIamV2UpdateCertRequestWithDefaults() *IamV2UpdateCertRequest`

NewIamV2UpdateCertRequestWithDefaults instantiates a new IamV2UpdateCertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IamV2UpdateCertRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IamV2UpdateCertRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IamV2UpdateCertRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IamV2UpdateCertRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IamV2UpdateCertRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IamV2UpdateCertRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IamV2UpdateCertRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IamV2UpdateCertRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *IamV2UpdateCertRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamV2UpdateCertRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamV2UpdateCertRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamV2UpdateCertRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *IamV2UpdateCertRequest) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamV2UpdateCertRequest) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamV2UpdateCertRequest) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IamV2UpdateCertRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamV2UpdateCertRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamV2UpdateCertRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamV2UpdateCertRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamV2UpdateCertRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *IamV2UpdateCertRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamV2UpdateCertRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamV2UpdateCertRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamV2UpdateCertRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCertificateChain

`func (o *IamV2UpdateCertRequest) GetCertificateChain() string`

GetCertificateChain returns the CertificateChain field if non-nil, zero value otherwise.

### GetCertificateChainOk

`func (o *IamV2UpdateCertRequest) GetCertificateChainOk() (*string, bool)`

GetCertificateChainOk returns a tuple with the CertificateChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateChain

`func (o *IamV2UpdateCertRequest) SetCertificateChain(v string)`

SetCertificateChain sets CertificateChain field to given value.

### HasCertificateChain

`func (o *IamV2UpdateCertRequest) HasCertificateChain() bool`

HasCertificateChain returns a boolean if a field has been set.

### GetCertificateChainFilename

`func (o *IamV2UpdateCertRequest) GetCertificateChainFilename() string`

GetCertificateChainFilename returns the CertificateChainFilename field if non-nil, zero value otherwise.

### GetCertificateChainFilenameOk

`func (o *IamV2UpdateCertRequest) GetCertificateChainFilenameOk() (*string, bool)`

GetCertificateChainFilenameOk returns a tuple with the CertificateChainFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateChainFilename

`func (o *IamV2UpdateCertRequest) SetCertificateChainFilename(v string)`

SetCertificateChainFilename sets CertificateChainFilename field to given value.

### HasCertificateChainFilename

`func (o *IamV2UpdateCertRequest) HasCertificateChainFilename() bool`

HasCertificateChainFilename returns a boolean if a field has been set.

### GetCrlUrl

`func (o *IamV2UpdateCertRequest) GetCrlUrl() string`

GetCrlUrl returns the CrlUrl field if non-nil, zero value otherwise.

### GetCrlUrlOk

`func (o *IamV2UpdateCertRequest) GetCrlUrlOk() (*string, bool)`

GetCrlUrlOk returns a tuple with the CrlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlUrl

`func (o *IamV2UpdateCertRequest) SetCrlUrl(v string)`

SetCrlUrl sets CrlUrl field to given value.

### HasCrlUrl

`func (o *IamV2UpdateCertRequest) HasCrlUrl() bool`

HasCrlUrl returns a boolean if a field has been set.

### GetCrlChain

`func (o *IamV2UpdateCertRequest) GetCrlChain() string`

GetCrlChain returns the CrlChain field if non-nil, zero value otherwise.

### GetCrlChainOk

`func (o *IamV2UpdateCertRequest) GetCrlChainOk() (*string, bool)`

GetCrlChainOk returns a tuple with the CrlChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlChain

`func (o *IamV2UpdateCertRequest) SetCrlChain(v string)`

SetCrlChain sets CrlChain field to given value.

### HasCrlChain

`func (o *IamV2UpdateCertRequest) HasCrlChain() bool`

HasCrlChain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


