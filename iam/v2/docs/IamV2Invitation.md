# IamV2Invitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Email** | Pointer to **string** | The user/invitee&#39;s email address | [optional] 
**AuthType** | Pointer to **string** | The user/invitee&#39;s authentication type. Note that only the [OrganizationAdmin role](https://docs.confluent.io/cloud/current/access-management/access-control/cloud-rbac.html#organizationadmin) can invite AUTH_TYPE_LOCAL users to SSO organizations. The user&#39;s auth_type is set as AUTH_TYPE_SSO by default if the organization has SSO enabled. Otherwise, the user&#39;s auth_type is AUTH_TYPE_LOCAL by default.  | [optional] 
**Status** | Pointer to **string** | The status of invitations | [optional] [readonly] 
**AcceptedAt** | Pointer to **NullableTime** | The timestamp that the invitation was accepted | [optional] [readonly] 
**ExpiresAt** | Pointer to **time.Time** | The timestamp that the invitation will expire | [optional] [readonly] 
**User** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The user/invitee | [optional] [readonly] 
**Creator** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The invitation creator | [optional] [readonly] 

## Methods

### NewIamV2Invitation

`func NewIamV2Invitation() *IamV2Invitation`

NewIamV2Invitation instantiates a new IamV2Invitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2InvitationWithDefaults

`func NewIamV2InvitationWithDefaults() *IamV2Invitation`

NewIamV2InvitationWithDefaults instantiates a new IamV2Invitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IamV2Invitation) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IamV2Invitation) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IamV2Invitation) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IamV2Invitation) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IamV2Invitation) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IamV2Invitation) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IamV2Invitation) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IamV2Invitation) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *IamV2Invitation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamV2Invitation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamV2Invitation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamV2Invitation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *IamV2Invitation) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamV2Invitation) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamV2Invitation) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IamV2Invitation) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetEmail

`func (o *IamV2Invitation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IamV2Invitation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IamV2Invitation) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IamV2Invitation) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAuthType

`func (o *IamV2Invitation) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *IamV2Invitation) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *IamV2Invitation) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *IamV2Invitation) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetStatus

`func (o *IamV2Invitation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IamV2Invitation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IamV2Invitation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IamV2Invitation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAcceptedAt

`func (o *IamV2Invitation) GetAcceptedAt() time.Time`

GetAcceptedAt returns the AcceptedAt field if non-nil, zero value otherwise.

### GetAcceptedAtOk

`func (o *IamV2Invitation) GetAcceptedAtOk() (*time.Time, bool)`

GetAcceptedAtOk returns a tuple with the AcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedAt

`func (o *IamV2Invitation) SetAcceptedAt(v time.Time)`

SetAcceptedAt sets AcceptedAt field to given value.

### HasAcceptedAt

`func (o *IamV2Invitation) HasAcceptedAt() bool`

HasAcceptedAt returns a boolean if a field has been set.

### SetAcceptedAtNil

`func (o *IamV2Invitation) SetAcceptedAtNil(b bool)`

 SetAcceptedAtNil sets the value for AcceptedAt to be an explicit nil

### UnsetAcceptedAt
`func (o *IamV2Invitation) UnsetAcceptedAt()`

UnsetAcceptedAt ensures that no value is present for AcceptedAt, not even an explicit nil
### GetExpiresAt

`func (o *IamV2Invitation) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *IamV2Invitation) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *IamV2Invitation) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *IamV2Invitation) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetUser

`func (o *IamV2Invitation) GetUser() GlobalObjectReference`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamV2Invitation) GetUserOk() (*GlobalObjectReference, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamV2Invitation) SetUser(v GlobalObjectReference)`

SetUser sets User field to given value.

### HasUser

`func (o *IamV2Invitation) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetCreator

`func (o *IamV2Invitation) GetCreator() GlobalObjectReference`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *IamV2Invitation) GetCreatorOk() (*GlobalObjectReference, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *IamV2Invitation) SetCreator(v GlobalObjectReference)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *IamV2Invitation) HasCreator() bool`

HasCreator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


