# IamV2RoleBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Principal** | Pointer to **string** | The principal User or Group to bind the role to | [optional] 
**RoleName** | Pointer to **string** | The name of the role to bind to the principal | [optional] 
**CrnPattern** | Pointer to **string** | A CRN that specifies the scope and resource patterns necessary for the role to bind | [optional] 

## Methods

### NewIamV2RoleBinding

`func NewIamV2RoleBinding() *IamV2RoleBinding`

NewIamV2RoleBinding instantiates a new IamV2RoleBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2RoleBindingWithDefaults

`func NewIamV2RoleBindingWithDefaults() *IamV2RoleBinding`

NewIamV2RoleBindingWithDefaults instantiates a new IamV2RoleBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IamV2RoleBinding) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IamV2RoleBinding) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IamV2RoleBinding) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IamV2RoleBinding) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IamV2RoleBinding) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IamV2RoleBinding) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IamV2RoleBinding) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IamV2RoleBinding) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *IamV2RoleBinding) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamV2RoleBinding) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamV2RoleBinding) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamV2RoleBinding) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *IamV2RoleBinding) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamV2RoleBinding) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamV2RoleBinding) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IamV2RoleBinding) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPrincipal

`func (o *IamV2RoleBinding) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *IamV2RoleBinding) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *IamV2RoleBinding) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *IamV2RoleBinding) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetRoleName

`func (o *IamV2RoleBinding) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *IamV2RoleBinding) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *IamV2RoleBinding) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *IamV2RoleBinding) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetCrnPattern

`func (o *IamV2RoleBinding) GetCrnPattern() string`

GetCrnPattern returns the CrnPattern field if non-nil, zero value otherwise.

### GetCrnPatternOk

`func (o *IamV2RoleBinding) GetCrnPatternOk() (*string, bool)`

GetCrnPatternOk returns a tuple with the CrnPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrnPattern

`func (o *IamV2RoleBinding) SetCrnPattern(v string)`

SetCrnPattern sets CrnPattern field to given value.

### HasCrnPattern

`func (o *IamV2RoleBinding) HasCrnPattern() bool`

HasCrnPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


