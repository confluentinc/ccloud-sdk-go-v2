# IamV2CertificateIdentityPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**DisplayName** | Pointer to **string** | The name of the &#x60;IdentityPool&#x60;. | [optional] 
**Description** | Pointer to **string** | A description of how this &#x60;IdentityPool&#x60; is used | [optional] 
**ExternalIdentity** | Pointer to **string** | The certificate field that will be used to represent the pool&#39;s external identity for audit logging. | [optional] 
**Filter** | Pointer to **string** | A filter expression in [Supported Common Expression Language (CEL)](https://docs.confluent.io/cloud/current/access-management/authenticate/oauth/identity-pools.html#supported-common-expression-language-cel-filters) that specifies which identities can authenticate using your identity pool (see [Set identity pool filters](https://docs.confluent.io/cloud/current/access-management/authenticate/oauth/identity-pools.html#set-identity-pool-filters) for more details). | [optional] 
**Principal** | Pointer to **string** | Represents the federated identity associated with this pool. | [optional] [readonly] 
**State** | Pointer to **string** | The current state of the identity pool | [optional] [readonly] 

## Methods

### NewIamV2CertificateIdentityPool

`func NewIamV2CertificateIdentityPool() *IamV2CertificateIdentityPool`

NewIamV2CertificateIdentityPool instantiates a new IamV2CertificateIdentityPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2CertificateIdentityPoolWithDefaults

`func NewIamV2CertificateIdentityPoolWithDefaults() *IamV2CertificateIdentityPool`

NewIamV2CertificateIdentityPoolWithDefaults instantiates a new IamV2CertificateIdentityPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IamV2CertificateIdentityPool) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IamV2CertificateIdentityPool) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IamV2CertificateIdentityPool) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IamV2CertificateIdentityPool) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IamV2CertificateIdentityPool) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IamV2CertificateIdentityPool) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IamV2CertificateIdentityPool) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IamV2CertificateIdentityPool) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *IamV2CertificateIdentityPool) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamV2CertificateIdentityPool) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamV2CertificateIdentityPool) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamV2CertificateIdentityPool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *IamV2CertificateIdentityPool) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamV2CertificateIdentityPool) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamV2CertificateIdentityPool) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IamV2CertificateIdentityPool) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamV2CertificateIdentityPool) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamV2CertificateIdentityPool) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamV2CertificateIdentityPool) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamV2CertificateIdentityPool) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *IamV2CertificateIdentityPool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamV2CertificateIdentityPool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamV2CertificateIdentityPool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamV2CertificateIdentityPool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalIdentity

`func (o *IamV2CertificateIdentityPool) GetExternalIdentity() string`

GetExternalIdentity returns the ExternalIdentity field if non-nil, zero value otherwise.

### GetExternalIdentityOk

`func (o *IamV2CertificateIdentityPool) GetExternalIdentityOk() (*string, bool)`

GetExternalIdentityOk returns a tuple with the ExternalIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentity

`func (o *IamV2CertificateIdentityPool) SetExternalIdentity(v string)`

SetExternalIdentity sets ExternalIdentity field to given value.

### HasExternalIdentity

`func (o *IamV2CertificateIdentityPool) HasExternalIdentity() bool`

HasExternalIdentity returns a boolean if a field has been set.

### GetFilter

`func (o *IamV2CertificateIdentityPool) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *IamV2CertificateIdentityPool) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *IamV2CertificateIdentityPool) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *IamV2CertificateIdentityPool) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPrincipal

`func (o *IamV2CertificateIdentityPool) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *IamV2CertificateIdentityPool) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *IamV2CertificateIdentityPool) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *IamV2CertificateIdentityPool) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetState

`func (o *IamV2CertificateIdentityPool) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IamV2CertificateIdentityPool) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IamV2CertificateIdentityPool) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IamV2CertificateIdentityPool) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


