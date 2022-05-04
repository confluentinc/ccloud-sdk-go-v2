# IamV2IdentityProviderUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**DisplayName** | Pointer to **string** | The name of the Provider. | [optional] 
**Description** | Pointer to **string** | A description for your provider | [optional] 
**Keys** | Pointer to [**[]IamV2Jwks**](IamV2Jwks.md) | The JWKS provided by the Provider. We only express the &#x60;kid&#x60; and &#x60;algorithm&#x60; for each key set | [optional] [readonly] 
**Deactivated** | Pointer to **bool** | Indicates whether the current provider is disabled. All corresponding &#x60;IdentityPool&#x60; objects would also be disabled  | [optional] [default to false]

## Methods

### NewIamV2IdentityProviderUpdate

`func NewIamV2IdentityProviderUpdate() *IamV2IdentityProviderUpdate`

NewIamV2IdentityProviderUpdate instantiates a new IamV2IdentityProviderUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2IdentityProviderUpdateWithDefaults

`func NewIamV2IdentityProviderUpdateWithDefaults() *IamV2IdentityProviderUpdate`

NewIamV2IdentityProviderUpdateWithDefaults instantiates a new IamV2IdentityProviderUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IamV2IdentityProviderUpdate) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IamV2IdentityProviderUpdate) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IamV2IdentityProviderUpdate) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IamV2IdentityProviderUpdate) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IamV2IdentityProviderUpdate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IamV2IdentityProviderUpdate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IamV2IdentityProviderUpdate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IamV2IdentityProviderUpdate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *IamV2IdentityProviderUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamV2IdentityProviderUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamV2IdentityProviderUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamV2IdentityProviderUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *IamV2IdentityProviderUpdate) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamV2IdentityProviderUpdate) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamV2IdentityProviderUpdate) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IamV2IdentityProviderUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamV2IdentityProviderUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamV2IdentityProviderUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamV2IdentityProviderUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamV2IdentityProviderUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *IamV2IdentityProviderUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamV2IdentityProviderUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamV2IdentityProviderUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamV2IdentityProviderUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKeys

`func (o *IamV2IdentityProviderUpdate) GetKeys() []IamV2Jwks`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *IamV2IdentityProviderUpdate) GetKeysOk() (*[]IamV2Jwks, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *IamV2IdentityProviderUpdate) SetKeys(v []IamV2Jwks)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *IamV2IdentityProviderUpdate) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetDeactivated

`func (o *IamV2IdentityProviderUpdate) GetDeactivated() bool`

GetDeactivated returns the Deactivated field if non-nil, zero value otherwise.

### GetDeactivatedOk

`func (o *IamV2IdentityProviderUpdate) GetDeactivatedOk() (*bool, bool)`

GetDeactivatedOk returns a tuple with the Deactivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivated

`func (o *IamV2IdentityProviderUpdate) SetDeactivated(v bool)`

SetDeactivated sets Deactivated field to given value.

### HasDeactivated

`func (o *IamV2IdentityProviderUpdate) HasDeactivated() bool`

HasDeactivated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


