# OrgV2ScimToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**ConnectionName** | Pointer to **string** | The SSO connection name associated with this token. | [optional] [readonly] 
**Token** | Pointer to **string** | The SCIM bearer token. Only provided in create responses, not in &#x60;list&#x60;. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | The date and time when the token was created. | [optional] [readonly] 
**ExpiresAt** | Pointer to **time.Time** | The date and time when the token expires. | [optional] [readonly] 

## Methods

### NewOrgV2ScimToken

`func NewOrgV2ScimToken() *OrgV2ScimToken`

NewOrgV2ScimToken instantiates a new OrgV2ScimToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgV2ScimTokenWithDefaults

`func NewOrgV2ScimTokenWithDefaults() *OrgV2ScimToken`

NewOrgV2ScimTokenWithDefaults instantiates a new OrgV2ScimToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *OrgV2ScimToken) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *OrgV2ScimToken) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *OrgV2ScimToken) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *OrgV2ScimToken) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *OrgV2ScimToken) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OrgV2ScimToken) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OrgV2ScimToken) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *OrgV2ScimToken) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *OrgV2ScimToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgV2ScimToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgV2ScimToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrgV2ScimToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *OrgV2ScimToken) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrgV2ScimToken) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrgV2ScimToken) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrgV2ScimToken) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetConnectionName

`func (o *OrgV2ScimToken) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *OrgV2ScimToken) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *OrgV2ScimToken) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.

### HasConnectionName

`func (o *OrgV2ScimToken) HasConnectionName() bool`

HasConnectionName returns a boolean if a field has been set.

### GetToken

`func (o *OrgV2ScimToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OrgV2ScimToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OrgV2ScimToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *OrgV2ScimToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrgV2ScimToken) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrgV2ScimToken) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrgV2ScimToken) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrgV2ScimToken) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *OrgV2ScimToken) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OrgV2ScimToken) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OrgV2ScimToken) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OrgV2ScimToken) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


