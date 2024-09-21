# IamV2CertificateAuthority

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**DisplayName** | Pointer to **string** | The human-readable name of the certificate authority. | [optional] 
**Description** | Pointer to **string** | A description of the certificate authority. | [optional] 
**Fingerprints** | Pointer to **[]string** | The fingerprints for each certificate in the certificate chain. These are SHA-1 encoded strings that act as unique identifiers for the certificates in the chain. | [optional] [readonly] 
**ExpirationDates** | Pointer to [**[]time.Time**](time.Time.md) | The expiration dates of certificates in the chain. | [optional] [readonly] 
**SerialNumbers** | Pointer to **[]string** | The serial numbers for each certificate in the certificate chain. | [optional] [readonly] 
**CertificateChainFilename** | Pointer to **string** | The file name of the uploaded pem file for this certificate authority. | [optional] [readonly] 
**CrlSource** | Pointer to **string** | The source specifies whether the Certificate Revocation List (CRL) is updated from either local file uploaded (LOCAL) or from url of CRL (URL). | [optional] [readonly] 
**CrlUrl** | Pointer to **string** | The url from which to fetch the CRL for the certificate authority if crl_source is URL. | [optional] [readonly] 
**CrlUpdatedAt** | Pointer to **time.Time** | The timestamp for when CRL was last updated. | [optional] [readonly] 
**State** | Pointer to **string** | The current state of the certificate authority. | [optional] [readonly] 

## Methods

### NewIamV2CertificateAuthority

`func NewIamV2CertificateAuthority() *IamV2CertificateAuthority`

NewIamV2CertificateAuthority instantiates a new IamV2CertificateAuthority object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2CertificateAuthorityWithDefaults

`func NewIamV2CertificateAuthorityWithDefaults() *IamV2CertificateAuthority`

NewIamV2CertificateAuthorityWithDefaults instantiates a new IamV2CertificateAuthority object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IamV2CertificateAuthority) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IamV2CertificateAuthority) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IamV2CertificateAuthority) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IamV2CertificateAuthority) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IamV2CertificateAuthority) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IamV2CertificateAuthority) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IamV2CertificateAuthority) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IamV2CertificateAuthority) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *IamV2CertificateAuthority) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamV2CertificateAuthority) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamV2CertificateAuthority) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamV2CertificateAuthority) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *IamV2CertificateAuthority) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamV2CertificateAuthority) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamV2CertificateAuthority) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IamV2CertificateAuthority) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamV2CertificateAuthority) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamV2CertificateAuthority) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamV2CertificateAuthority) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamV2CertificateAuthority) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *IamV2CertificateAuthority) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamV2CertificateAuthority) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamV2CertificateAuthority) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamV2CertificateAuthority) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFingerprints

`func (o *IamV2CertificateAuthority) GetFingerprints() []string`

GetFingerprints returns the Fingerprints field if non-nil, zero value otherwise.

### GetFingerprintsOk

`func (o *IamV2CertificateAuthority) GetFingerprintsOk() (*[]string, bool)`

GetFingerprintsOk returns a tuple with the Fingerprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprints

`func (o *IamV2CertificateAuthority) SetFingerprints(v []string)`

SetFingerprints sets Fingerprints field to given value.

### HasFingerprints

`func (o *IamV2CertificateAuthority) HasFingerprints() bool`

HasFingerprints returns a boolean if a field has been set.

### GetExpirationDates

`func (o *IamV2CertificateAuthority) GetExpirationDates() []time.Time`

GetExpirationDates returns the ExpirationDates field if non-nil, zero value otherwise.

### GetExpirationDatesOk

`func (o *IamV2CertificateAuthority) GetExpirationDatesOk() (*[]time.Time, bool)`

GetExpirationDatesOk returns a tuple with the ExpirationDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDates

`func (o *IamV2CertificateAuthority) SetExpirationDates(v []time.Time)`

SetExpirationDates sets ExpirationDates field to given value.

### HasExpirationDates

`func (o *IamV2CertificateAuthority) HasExpirationDates() bool`

HasExpirationDates returns a boolean if a field has been set.

### GetSerialNumbers

`func (o *IamV2CertificateAuthority) GetSerialNumbers() []string`

GetSerialNumbers returns the SerialNumbers field if non-nil, zero value otherwise.

### GetSerialNumbersOk

`func (o *IamV2CertificateAuthority) GetSerialNumbersOk() (*[]string, bool)`

GetSerialNumbersOk returns a tuple with the SerialNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumbers

`func (o *IamV2CertificateAuthority) SetSerialNumbers(v []string)`

SetSerialNumbers sets SerialNumbers field to given value.

### HasSerialNumbers

`func (o *IamV2CertificateAuthority) HasSerialNumbers() bool`

HasSerialNumbers returns a boolean if a field has been set.

### GetCertificateChainFilename

`func (o *IamV2CertificateAuthority) GetCertificateChainFilename() string`

GetCertificateChainFilename returns the CertificateChainFilename field if non-nil, zero value otherwise.

### GetCertificateChainFilenameOk

`func (o *IamV2CertificateAuthority) GetCertificateChainFilenameOk() (*string, bool)`

GetCertificateChainFilenameOk returns a tuple with the CertificateChainFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateChainFilename

`func (o *IamV2CertificateAuthority) SetCertificateChainFilename(v string)`

SetCertificateChainFilename sets CertificateChainFilename field to given value.

### HasCertificateChainFilename

`func (o *IamV2CertificateAuthority) HasCertificateChainFilename() bool`

HasCertificateChainFilename returns a boolean if a field has been set.

### GetCrlSource

`func (o *IamV2CertificateAuthority) GetCrlSource() string`

GetCrlSource returns the CrlSource field if non-nil, zero value otherwise.

### GetCrlSourceOk

`func (o *IamV2CertificateAuthority) GetCrlSourceOk() (*string, bool)`

GetCrlSourceOk returns a tuple with the CrlSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlSource

`func (o *IamV2CertificateAuthority) SetCrlSource(v string)`

SetCrlSource sets CrlSource field to given value.

### HasCrlSource

`func (o *IamV2CertificateAuthority) HasCrlSource() bool`

HasCrlSource returns a boolean if a field has been set.

### GetCrlUrl

`func (o *IamV2CertificateAuthority) GetCrlUrl() string`

GetCrlUrl returns the CrlUrl field if non-nil, zero value otherwise.

### GetCrlUrlOk

`func (o *IamV2CertificateAuthority) GetCrlUrlOk() (*string, bool)`

GetCrlUrlOk returns a tuple with the CrlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlUrl

`func (o *IamV2CertificateAuthority) SetCrlUrl(v string)`

SetCrlUrl sets CrlUrl field to given value.

### HasCrlUrl

`func (o *IamV2CertificateAuthority) HasCrlUrl() bool`

HasCrlUrl returns a boolean if a field has been set.

### GetCrlUpdatedAt

`func (o *IamV2CertificateAuthority) GetCrlUpdatedAt() time.Time`

GetCrlUpdatedAt returns the CrlUpdatedAt field if non-nil, zero value otherwise.

### GetCrlUpdatedAtOk

`func (o *IamV2CertificateAuthority) GetCrlUpdatedAtOk() (*time.Time, bool)`

GetCrlUpdatedAtOk returns a tuple with the CrlUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlUpdatedAt

`func (o *IamV2CertificateAuthority) SetCrlUpdatedAt(v time.Time)`

SetCrlUpdatedAt sets CrlUpdatedAt field to given value.

### HasCrlUpdatedAt

`func (o *IamV2CertificateAuthority) HasCrlUpdatedAt() bool`

HasCrlUpdatedAt returns a boolean if a field has been set.

### GetState

`func (o *IamV2CertificateAuthority) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IamV2CertificateAuthority) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IamV2CertificateAuthority) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IamV2CertificateAuthority) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


