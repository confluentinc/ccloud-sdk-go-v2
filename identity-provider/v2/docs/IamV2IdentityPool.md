# IamV2IdentityPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**DisplayName** | Pointer to **string** | The name of the &#x60;IdentityPool&#x60;. | [optional] 
**Description** | Pointer to **string** | A description of how this &#x60;IdentityPool&#x60; is used | [optional] 
**PrincipalClaim** | Pointer to **string** | The JWT claim to extract the authenticating principal to Confluent resources | [optional] 
**Policy** | Pointer to **string** | A filter expression that must be evaluated to be true to use this identity pool | [optional] 
**Deactivated** | Pointer to **bool** | Indicates whether the current identity pool is disabled. | [optional] [default to false]

## Methods

### NewIamV2IdentityPool

`func NewIamV2IdentityPool() *IamV2IdentityPool`

NewIamV2IdentityPool instantiates a new IamV2IdentityPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2IdentityPoolWithDefaults

`func NewIamV2IdentityPoolWithDefaults() *IamV2IdentityPool`

NewIamV2IdentityPoolWithDefaults instantiates a new IamV2IdentityPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IamV2IdentityPool) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IamV2IdentityPool) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IamV2IdentityPool) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IamV2IdentityPool) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IamV2IdentityPool) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IamV2IdentityPool) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IamV2IdentityPool) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IamV2IdentityPool) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *IamV2IdentityPool) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamV2IdentityPool) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamV2IdentityPool) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamV2IdentityPool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *IamV2IdentityPool) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamV2IdentityPool) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamV2IdentityPool) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IamV2IdentityPool) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamV2IdentityPool) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamV2IdentityPool) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamV2IdentityPool) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamV2IdentityPool) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *IamV2IdentityPool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamV2IdentityPool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamV2IdentityPool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamV2IdentityPool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrincipalClaim

`func (o *IamV2IdentityPool) GetPrincipalClaim() string`

GetPrincipalClaim returns the PrincipalClaim field if non-nil, zero value otherwise.

### GetPrincipalClaimOk

`func (o *IamV2IdentityPool) GetPrincipalClaimOk() (*string, bool)`

GetPrincipalClaimOk returns a tuple with the PrincipalClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalClaim

`func (o *IamV2IdentityPool) SetPrincipalClaim(v string)`

SetPrincipalClaim sets PrincipalClaim field to given value.

### HasPrincipalClaim

`func (o *IamV2IdentityPool) HasPrincipalClaim() bool`

HasPrincipalClaim returns a boolean if a field has been set.

### GetPolicy

`func (o *IamV2IdentityPool) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *IamV2IdentityPool) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *IamV2IdentityPool) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *IamV2IdentityPool) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetDeactivated

`func (o *IamV2IdentityPool) GetDeactivated() bool`

GetDeactivated returns the Deactivated field if non-nil, zero value otherwise.

### GetDeactivatedOk

`func (o *IamV2IdentityPool) GetDeactivatedOk() (*bool, bool)`

GetDeactivatedOk returns a tuple with the Deactivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivated

`func (o *IamV2IdentityPool) SetDeactivated(v bool)`

SetDeactivated sets Deactivated field to given value.

### HasDeactivated

`func (o *IamV2IdentityPool) HasDeactivated() bool`

HasDeactivated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


