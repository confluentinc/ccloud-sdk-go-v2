# CdxV1ConsumerShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**ProviderOrganizationName** | Pointer to **string** | Provider organization name | [optional] [readonly] 
**ProviderUserName** | Pointer to **string** | Name or email of the provider user | [optional] [readonly] 
**InviteExpiresAt** | Pointer to **time.Time** | The date and time at which the invitation will expire. Only for invited shares | [optional] [readonly] 
**ConsumerOrganizationName** | Pointer to **string** | Consumer organization name. Deprecated | [optional] [readonly] 
**ConsumerUserName** | Pointer to **string** | Name of the consumer. Deprecated | [optional] [readonly] 
**ConsumerUser** | Pointer to [**ObjectReference**](ObjectReference.md) | The consumer user/invitee | [optional] [readonly] 
**Status** | Pointer to [**CdxV1ConsumerShareStatus**](CdxV1ConsumerShareStatus.md) |  | [optional] 

## Methods

### NewCdxV1ConsumerShare

`func NewCdxV1ConsumerShare() *CdxV1ConsumerShare`

NewCdxV1ConsumerShare instantiates a new CdxV1ConsumerShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdxV1ConsumerShareWithDefaults

`func NewCdxV1ConsumerShareWithDefaults() *CdxV1ConsumerShare`

NewCdxV1ConsumerShareWithDefaults instantiates a new CdxV1ConsumerShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CdxV1ConsumerShare) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CdxV1ConsumerShare) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CdxV1ConsumerShare) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CdxV1ConsumerShare) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CdxV1ConsumerShare) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CdxV1ConsumerShare) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CdxV1ConsumerShare) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CdxV1ConsumerShare) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CdxV1ConsumerShare) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdxV1ConsumerShare) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdxV1ConsumerShare) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CdxV1ConsumerShare) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CdxV1ConsumerShare) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CdxV1ConsumerShare) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CdxV1ConsumerShare) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CdxV1ConsumerShare) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProviderOrganizationName

`func (o *CdxV1ConsumerShare) GetProviderOrganizationName() string`

GetProviderOrganizationName returns the ProviderOrganizationName field if non-nil, zero value otherwise.

### GetProviderOrganizationNameOk

`func (o *CdxV1ConsumerShare) GetProviderOrganizationNameOk() (*string, bool)`

GetProviderOrganizationNameOk returns a tuple with the ProviderOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderOrganizationName

`func (o *CdxV1ConsumerShare) SetProviderOrganizationName(v string)`

SetProviderOrganizationName sets ProviderOrganizationName field to given value.

### HasProviderOrganizationName

`func (o *CdxV1ConsumerShare) HasProviderOrganizationName() bool`

HasProviderOrganizationName returns a boolean if a field has been set.

### GetProviderUserName

`func (o *CdxV1ConsumerShare) GetProviderUserName() string`

GetProviderUserName returns the ProviderUserName field if non-nil, zero value otherwise.

### GetProviderUserNameOk

`func (o *CdxV1ConsumerShare) GetProviderUserNameOk() (*string, bool)`

GetProviderUserNameOk returns a tuple with the ProviderUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUserName

`func (o *CdxV1ConsumerShare) SetProviderUserName(v string)`

SetProviderUserName sets ProviderUserName field to given value.

### HasProviderUserName

`func (o *CdxV1ConsumerShare) HasProviderUserName() bool`

HasProviderUserName returns a boolean if a field has been set.

### GetInviteExpiresAt

`func (o *CdxV1ConsumerShare) GetInviteExpiresAt() time.Time`

GetInviteExpiresAt returns the InviteExpiresAt field if non-nil, zero value otherwise.

### GetInviteExpiresAtOk

`func (o *CdxV1ConsumerShare) GetInviteExpiresAtOk() (*time.Time, bool)`

GetInviteExpiresAtOk returns a tuple with the InviteExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteExpiresAt

`func (o *CdxV1ConsumerShare) SetInviteExpiresAt(v time.Time)`

SetInviteExpiresAt sets InviteExpiresAt field to given value.

### HasInviteExpiresAt

`func (o *CdxV1ConsumerShare) HasInviteExpiresAt() bool`

HasInviteExpiresAt returns a boolean if a field has been set.

### GetConsumerOrganizationName

`func (o *CdxV1ConsumerShare) GetConsumerOrganizationName() string`

GetConsumerOrganizationName returns the ConsumerOrganizationName field if non-nil, zero value otherwise.

### GetConsumerOrganizationNameOk

`func (o *CdxV1ConsumerShare) GetConsumerOrganizationNameOk() (*string, bool)`

GetConsumerOrganizationNameOk returns a tuple with the ConsumerOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerOrganizationName

`func (o *CdxV1ConsumerShare) SetConsumerOrganizationName(v string)`

SetConsumerOrganizationName sets ConsumerOrganizationName field to given value.

### HasConsumerOrganizationName

`func (o *CdxV1ConsumerShare) HasConsumerOrganizationName() bool`

HasConsumerOrganizationName returns a boolean if a field has been set.

### GetConsumerUserName

`func (o *CdxV1ConsumerShare) GetConsumerUserName() string`

GetConsumerUserName returns the ConsumerUserName field if non-nil, zero value otherwise.

### GetConsumerUserNameOk

`func (o *CdxV1ConsumerShare) GetConsumerUserNameOk() (*string, bool)`

GetConsumerUserNameOk returns a tuple with the ConsumerUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerUserName

`func (o *CdxV1ConsumerShare) SetConsumerUserName(v string)`

SetConsumerUserName sets ConsumerUserName field to given value.

### HasConsumerUserName

`func (o *CdxV1ConsumerShare) HasConsumerUserName() bool`

HasConsumerUserName returns a boolean if a field has been set.

### GetConsumerUser

`func (o *CdxV1ConsumerShare) GetConsumerUser() ObjectReference`

GetConsumerUser returns the ConsumerUser field if non-nil, zero value otherwise.

### GetConsumerUserOk

`func (o *CdxV1ConsumerShare) GetConsumerUserOk() (*ObjectReference, bool)`

GetConsumerUserOk returns a tuple with the ConsumerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerUser

`func (o *CdxV1ConsumerShare) SetConsumerUser(v ObjectReference)`

SetConsumerUser sets ConsumerUser field to given value.

### HasConsumerUser

`func (o *CdxV1ConsumerShare) HasConsumerUser() bool`

HasConsumerUser returns a boolean if a field has been set.

### GetStatus

`func (o *CdxV1ConsumerShare) GetStatus() CdxV1ConsumerShareStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CdxV1ConsumerShare) GetStatusOk() (*CdxV1ConsumerShareStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CdxV1ConsumerShare) SetStatus(v CdxV1ConsumerShareStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CdxV1ConsumerShare) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


