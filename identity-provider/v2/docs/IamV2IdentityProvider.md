# IamV2IdentityProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**DisplayName** | Pointer to **string** | The human-readable name of the OAuth identity provider. | [optional] 
**Description** | Pointer to **string** | A description of the identity provider. | [optional] 
**IdentityClaim** | Pointer to **string** | The JSON Web Token (JWT) claim to extract the authenticating identity to Confluent resources from [Registered Claim Names](https://datatracker.ietf.org/doc/html/rfc7519#section-4.1). This appears in audit log records. Note: if the client specifies mapping to one identity pool ID, the identity claim configured with that pool will be used instead. | [optional] 
**State** | Pointer to **string** | The current state of the identity provider. | [optional] [readonly] 
**Issuer** | Pointer to **string** | A publicly accessible URL uniquely identifying the OAuth identity provider authorized to issue access tokens. | [optional] 
**JwksUri** | Pointer to **string** | A publicly accessible JSON Web Key Set (JWKS) URI for the OAuth identity provider. JWKS provides a set of crypotgraphic keys used to verify the authenticity and integrity of JSON Web Tokens (JWTs) issued by the OAuth identity provider. | [optional] 
**Keys** | Pointer to [**[]IamV2JwksObject**](IamV2JwksObject.md) | The JWKS issued by the OAuth identity provider. Only &#x60;kid&#x60; (key ID) and &#x60;alg&#x60; (algorithm) properties for each key set are included. | [optional] [readonly] 

## Methods

### NewIamV2IdentityProvider

`func NewIamV2IdentityProvider() *IamV2IdentityProvider`

NewIamV2IdentityProvider instantiates a new IamV2IdentityProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2IdentityProviderWithDefaults

`func NewIamV2IdentityProviderWithDefaults() *IamV2IdentityProvider`

NewIamV2IdentityProviderWithDefaults instantiates a new IamV2IdentityProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IamV2IdentityProvider) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IamV2IdentityProvider) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IamV2IdentityProvider) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IamV2IdentityProvider) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IamV2IdentityProvider) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IamV2IdentityProvider) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IamV2IdentityProvider) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IamV2IdentityProvider) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *IamV2IdentityProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamV2IdentityProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamV2IdentityProvider) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamV2IdentityProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *IamV2IdentityProvider) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamV2IdentityProvider) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamV2IdentityProvider) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IamV2IdentityProvider) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamV2IdentityProvider) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamV2IdentityProvider) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamV2IdentityProvider) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamV2IdentityProvider) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *IamV2IdentityProvider) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamV2IdentityProvider) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamV2IdentityProvider) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamV2IdentityProvider) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIdentityClaim

`func (o *IamV2IdentityProvider) GetIdentityClaim() string`

GetIdentityClaim returns the IdentityClaim field if non-nil, zero value otherwise.

### GetIdentityClaimOk

`func (o *IamV2IdentityProvider) GetIdentityClaimOk() (*string, bool)`

GetIdentityClaimOk returns a tuple with the IdentityClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityClaim

`func (o *IamV2IdentityProvider) SetIdentityClaim(v string)`

SetIdentityClaim sets IdentityClaim field to given value.

### HasIdentityClaim

`func (o *IamV2IdentityProvider) HasIdentityClaim() bool`

HasIdentityClaim returns a boolean if a field has been set.

### GetState

`func (o *IamV2IdentityProvider) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IamV2IdentityProvider) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IamV2IdentityProvider) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IamV2IdentityProvider) HasState() bool`

HasState returns a boolean if a field has been set.

### GetIssuer

`func (o *IamV2IdentityProvider) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *IamV2IdentityProvider) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *IamV2IdentityProvider) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *IamV2IdentityProvider) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetJwksUri

`func (o *IamV2IdentityProvider) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *IamV2IdentityProvider) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *IamV2IdentityProvider) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *IamV2IdentityProvider) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.

### GetKeys

`func (o *IamV2IdentityProvider) GetKeys() []IamV2JwksObject`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *IamV2IdentityProvider) GetKeysOk() (*[]IamV2JwksObject, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *IamV2IdentityProvider) SetKeys(v []IamV2JwksObject)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *IamV2IdentityProvider) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


