# V2UserInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**SendInvitation** | Pointer to **bool** | Send invitation to email address | [optional] 
**Status** | Pointer to **string** | The invitation status | [optional] 
**User** | Pointer to [**IamV2User**](iam.v2.User.md) | The invitation user object | [optional] 
**RoleBindings** | Pointer to [**[]IamV2RoleBinding**](IamV2RoleBinding.md) | Role bindings associated with the user invitation. | [optional] 

## Methods

### NewV2UserInvitation

`func NewV2UserInvitation() *V2UserInvitation`

NewV2UserInvitation instantiates a new V2UserInvitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2UserInvitationWithDefaults

`func NewV2UserInvitationWithDefaults() *V2UserInvitation`

NewV2UserInvitationWithDefaults instantiates a new V2UserInvitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V2UserInvitation) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V2UserInvitation) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V2UserInvitation) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V2UserInvitation) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *V2UserInvitation) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V2UserInvitation) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V2UserInvitation) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V2UserInvitation) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *V2UserInvitation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2UserInvitation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2UserInvitation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V2UserInvitation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *V2UserInvitation) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V2UserInvitation) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V2UserInvitation) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V2UserInvitation) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSendInvitation

`func (o *V2UserInvitation) GetSendInvitation() bool`

GetSendInvitation returns the SendInvitation field if non-nil, zero value otherwise.

### GetSendInvitationOk

`func (o *V2UserInvitation) GetSendInvitationOk() (*bool, bool)`

GetSendInvitationOk returns a tuple with the SendInvitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendInvitation

`func (o *V2UserInvitation) SetSendInvitation(v bool)`

SetSendInvitation sets SendInvitation field to given value.

### HasSendInvitation

`func (o *V2UserInvitation) HasSendInvitation() bool`

HasSendInvitation returns a boolean if a field has been set.

### GetStatus

`func (o *V2UserInvitation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V2UserInvitation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V2UserInvitation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V2UserInvitation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUser

`func (o *V2UserInvitation) GetUser() IamV2User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V2UserInvitation) GetUserOk() (*IamV2User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V2UserInvitation) SetUser(v IamV2User)`

SetUser sets User field to given value.

### HasUser

`func (o *V2UserInvitation) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRoleBindings

`func (o *V2UserInvitation) GetRoleBindings() []IamV2RoleBinding`

GetRoleBindings returns the RoleBindings field if non-nil, zero value otherwise.

### GetRoleBindingsOk

`func (o *V2UserInvitation) GetRoleBindingsOk() (*[]IamV2RoleBinding, bool)`

GetRoleBindingsOk returns a tuple with the RoleBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleBindings

`func (o *V2UserInvitation) SetRoleBindings(v []IamV2RoleBinding)`

SetRoleBindings sets RoleBindings field to given value.

### HasRoleBindings

`func (o *V2UserInvitation) HasRoleBindings() bool`

HasRoleBindings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


