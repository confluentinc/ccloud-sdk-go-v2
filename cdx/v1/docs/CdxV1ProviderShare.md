# CdxV1ProviderShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**ConsumerUserName** | Pointer to **string** | Name of the consumer | [optional] [readonly] 
**ConsumerOrganizationName** | Pointer to **string** | Consumer organization name | [optional] [readonly] 
**ProviderUserName** | Pointer to **string** | Name or email of the provider user. Deprecated | [optional] [readonly] 
**DeliveryMethod** | Pointer to **string** | Method by which the invite will be delivered | [optional] 
**ConsumerRestriction** | Pointer to [**CdxV1ProviderShareConsumerRestrictionOneOf**](CdxV1ProviderShareConsumerRestrictionOneOf.md) | Restrictions on the consumer that can redeem this token | [optional] 
**InvitedAt** | Pointer to **time.Time** | The date and time at which consumer was invited | [optional] [readonly] 
**InviteExpiresAt** | Pointer to **time.Time** | The date and time at which the invitation will expire. Only for invited shares | [optional] [readonly] 
**RedeemedAt** | Pointer to **time.Time** | The date and time at which the invite was redeemed | [optional] [readonly] 
**ProviderUser** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The provider user/inviter | [optional] [readonly] 
**ServiceAccount** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The service account associated with this object. | [optional] 
**CloudCluster** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The cloud cluster to which this belongs. | [optional] 
**Status** | Pointer to [**CdxV1ProviderShareStatus**](CdxV1ProviderShareStatus.md) |  | [optional] 

## Methods

### NewCdxV1ProviderShare

`func NewCdxV1ProviderShare() *CdxV1ProviderShare`

NewCdxV1ProviderShare instantiates a new CdxV1ProviderShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdxV1ProviderShareWithDefaults

`func NewCdxV1ProviderShareWithDefaults() *CdxV1ProviderShare`

NewCdxV1ProviderShareWithDefaults instantiates a new CdxV1ProviderShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CdxV1ProviderShare) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CdxV1ProviderShare) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CdxV1ProviderShare) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CdxV1ProviderShare) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CdxV1ProviderShare) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CdxV1ProviderShare) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CdxV1ProviderShare) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CdxV1ProviderShare) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CdxV1ProviderShare) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdxV1ProviderShare) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdxV1ProviderShare) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CdxV1ProviderShare) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CdxV1ProviderShare) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CdxV1ProviderShare) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CdxV1ProviderShare) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CdxV1ProviderShare) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetConsumerUserName

`func (o *CdxV1ProviderShare) GetConsumerUserName() string`

GetConsumerUserName returns the ConsumerUserName field if non-nil, zero value otherwise.

### GetConsumerUserNameOk

`func (o *CdxV1ProviderShare) GetConsumerUserNameOk() (*string, bool)`

GetConsumerUserNameOk returns a tuple with the ConsumerUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerUserName

`func (o *CdxV1ProviderShare) SetConsumerUserName(v string)`

SetConsumerUserName sets ConsumerUserName field to given value.

### HasConsumerUserName

`func (o *CdxV1ProviderShare) HasConsumerUserName() bool`

HasConsumerUserName returns a boolean if a field has been set.

### GetConsumerOrganizationName

`func (o *CdxV1ProviderShare) GetConsumerOrganizationName() string`

GetConsumerOrganizationName returns the ConsumerOrganizationName field if non-nil, zero value otherwise.

### GetConsumerOrganizationNameOk

`func (o *CdxV1ProviderShare) GetConsumerOrganizationNameOk() (*string, bool)`

GetConsumerOrganizationNameOk returns a tuple with the ConsumerOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerOrganizationName

`func (o *CdxV1ProviderShare) SetConsumerOrganizationName(v string)`

SetConsumerOrganizationName sets ConsumerOrganizationName field to given value.

### HasConsumerOrganizationName

`func (o *CdxV1ProviderShare) HasConsumerOrganizationName() bool`

HasConsumerOrganizationName returns a boolean if a field has been set.

### GetProviderUserName

`func (o *CdxV1ProviderShare) GetProviderUserName() string`

GetProviderUserName returns the ProviderUserName field if non-nil, zero value otherwise.

### GetProviderUserNameOk

`func (o *CdxV1ProviderShare) GetProviderUserNameOk() (*string, bool)`

GetProviderUserNameOk returns a tuple with the ProviderUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUserName

`func (o *CdxV1ProviderShare) SetProviderUserName(v string)`

SetProviderUserName sets ProviderUserName field to given value.

### HasProviderUserName

`func (o *CdxV1ProviderShare) HasProviderUserName() bool`

HasProviderUserName returns a boolean if a field has been set.

### GetDeliveryMethod

`func (o *CdxV1ProviderShare) GetDeliveryMethod() string`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *CdxV1ProviderShare) GetDeliveryMethodOk() (*string, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *CdxV1ProviderShare) SetDeliveryMethod(v string)`

SetDeliveryMethod sets DeliveryMethod field to given value.

### HasDeliveryMethod

`func (o *CdxV1ProviderShare) HasDeliveryMethod() bool`

HasDeliveryMethod returns a boolean if a field has been set.

### GetConsumerRestriction

`func (o *CdxV1ProviderShare) GetConsumerRestriction() CdxV1ProviderShareConsumerRestrictionOneOf`

GetConsumerRestriction returns the ConsumerRestriction field if non-nil, zero value otherwise.

### GetConsumerRestrictionOk

`func (o *CdxV1ProviderShare) GetConsumerRestrictionOk() (*CdxV1ProviderShareConsumerRestrictionOneOf, bool)`

GetConsumerRestrictionOk returns a tuple with the ConsumerRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerRestriction

`func (o *CdxV1ProviderShare) SetConsumerRestriction(v CdxV1ProviderShareConsumerRestrictionOneOf)`

SetConsumerRestriction sets ConsumerRestriction field to given value.

### HasConsumerRestriction

`func (o *CdxV1ProviderShare) HasConsumerRestriction() bool`

HasConsumerRestriction returns a boolean if a field has been set.

### GetInvitedAt

`func (o *CdxV1ProviderShare) GetInvitedAt() time.Time`

GetInvitedAt returns the InvitedAt field if non-nil, zero value otherwise.

### GetInvitedAtOk

`func (o *CdxV1ProviderShare) GetInvitedAtOk() (*time.Time, bool)`

GetInvitedAtOk returns a tuple with the InvitedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedAt

`func (o *CdxV1ProviderShare) SetInvitedAt(v time.Time)`

SetInvitedAt sets InvitedAt field to given value.

### HasInvitedAt

`func (o *CdxV1ProviderShare) HasInvitedAt() bool`

HasInvitedAt returns a boolean if a field has been set.

### GetInviteExpiresAt

`func (o *CdxV1ProviderShare) GetInviteExpiresAt() time.Time`

GetInviteExpiresAt returns the InviteExpiresAt field if non-nil, zero value otherwise.

### GetInviteExpiresAtOk

`func (o *CdxV1ProviderShare) GetInviteExpiresAtOk() (*time.Time, bool)`

GetInviteExpiresAtOk returns a tuple with the InviteExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteExpiresAt

`func (o *CdxV1ProviderShare) SetInviteExpiresAt(v time.Time)`

SetInviteExpiresAt sets InviteExpiresAt field to given value.

### HasInviteExpiresAt

`func (o *CdxV1ProviderShare) HasInviteExpiresAt() bool`

HasInviteExpiresAt returns a boolean if a field has been set.

### GetRedeemedAt

`func (o *CdxV1ProviderShare) GetRedeemedAt() time.Time`

GetRedeemedAt returns the RedeemedAt field if non-nil, zero value otherwise.

### GetRedeemedAtOk

`func (o *CdxV1ProviderShare) GetRedeemedAtOk() (*time.Time, bool)`

GetRedeemedAtOk returns a tuple with the RedeemedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemedAt

`func (o *CdxV1ProviderShare) SetRedeemedAt(v time.Time)`

SetRedeemedAt sets RedeemedAt field to given value.

### HasRedeemedAt

`func (o *CdxV1ProviderShare) HasRedeemedAt() bool`

HasRedeemedAt returns a boolean if a field has been set.

### GetProviderUser

`func (o *CdxV1ProviderShare) GetProviderUser() GlobalObjectReference`

GetProviderUser returns the ProviderUser field if non-nil, zero value otherwise.

### GetProviderUserOk

`func (o *CdxV1ProviderShare) GetProviderUserOk() (*GlobalObjectReference, bool)`

GetProviderUserOk returns a tuple with the ProviderUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUser

`func (o *CdxV1ProviderShare) SetProviderUser(v GlobalObjectReference)`

SetProviderUser sets ProviderUser field to given value.

### HasProviderUser

`func (o *CdxV1ProviderShare) HasProviderUser() bool`

HasProviderUser returns a boolean if a field has been set.

### GetServiceAccount

`func (o *CdxV1ProviderShare) GetServiceAccount() GlobalObjectReference`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *CdxV1ProviderShare) GetServiceAccountOk() (*GlobalObjectReference, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *CdxV1ProviderShare) SetServiceAccount(v GlobalObjectReference)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *CdxV1ProviderShare) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetCloudCluster

`func (o *CdxV1ProviderShare) GetCloudCluster() EnvScopedObjectReference`

GetCloudCluster returns the CloudCluster field if non-nil, zero value otherwise.

### GetCloudClusterOk

`func (o *CdxV1ProviderShare) GetCloudClusterOk() (*EnvScopedObjectReference, bool)`

GetCloudClusterOk returns a tuple with the CloudCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCluster

`func (o *CdxV1ProviderShare) SetCloudCluster(v EnvScopedObjectReference)`

SetCloudCluster sets CloudCluster field to given value.

### HasCloudCluster

`func (o *CdxV1ProviderShare) HasCloudCluster() bool`

HasCloudCluster returns a boolean if a field has been set.

### GetStatus

`func (o *CdxV1ProviderShare) GetStatus() CdxV1ProviderShareStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CdxV1ProviderShare) GetStatusOk() (*CdxV1ProviderShareStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CdxV1ProviderShare) SetStatus(v CdxV1ProviderShareStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CdxV1ProviderShare) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


