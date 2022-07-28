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
**IdentityClaim** | Pointer to **string** | The JWT claim to extract the authenticating identity to Confluent resources | [optional] 
**Filter** | Pointer to **string** | A filter expression that must be evaluated to be true to use this identity pool | [optional] 
**Principal** | Pointer to **string** | Represents the federated identity associated with this pool. | [optional] [readonly] 
**State** | Pointer to **string** | The current state of the identity pool | [optional] [readonly] 

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

### GetIdentityClaim

`func (o *IamV2IdentityPool) GetIdentityClaim() string`

GetIdentityClaim returns the IdentityClaim field if non-nil, zero value otherwise.

### GetIdentityClaimOk

`func (o *IamV2IdentityPool) GetIdentityClaimOk() (*string, bool)`

GetIdentityClaimOk returns a tuple with the IdentityClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityClaim

`func (o *IamV2IdentityPool) SetIdentityClaim(v string)`

SetIdentityClaim sets IdentityClaim field to given value.

### HasIdentityClaim

`func (o *IamV2IdentityPool) HasIdentityClaim() bool`

HasIdentityClaim returns a boolean if a field has been set.

### GetFilter

`func (o *IamV2IdentityPool) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *IamV2IdentityPool) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *IamV2IdentityPool) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *IamV2IdentityPool) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPrincipal

`func (o *IamV2IdentityPool) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *IamV2IdentityPool) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *IamV2IdentityPool) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *IamV2IdentityPool) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetState

`func (o *IamV2IdentityPool) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IamV2IdentityPool) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IamV2IdentityPool) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IamV2IdentityPool) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


