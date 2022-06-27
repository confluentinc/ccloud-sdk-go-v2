# CdxV1ConsumerShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**ConsumerUserName** | Pointer to **string** | Name of the consumer | [optional] 
**ConsumerOrganizationName** | Pointer to **string** | Consumer organization name | [optional] 
**ProviderUserName** | Pointer to **string** | Name of the provider | [optional] 
**Status** | Pointer to **string** | Status of share | [optional] 
**InviteExpiresAt** | Pointer to **time.Time** | The date and time at which the invitation will expire. Only for invited shares | [optional] 
**SharedResource** | Pointer to [**ObjectReference**](ObjectReference.md) | The shared_resource to which this belongs. | [optional] 

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

### GetStatus

`func (o *CdxV1ConsumerShare) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CdxV1ConsumerShare) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CdxV1ConsumerShare) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CdxV1ConsumerShare) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

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

### GetSharedResource

`func (o *CdxV1ConsumerShare) GetSharedResource() ObjectReference`

GetSharedResource returns the SharedResource field if non-nil, zero value otherwise.

### GetSharedResourceOk

`func (o *CdxV1ConsumerShare) GetSharedResourceOk() (*ObjectReference, bool)`

GetSharedResourceOk returns a tuple with the SharedResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedResource

`func (o *CdxV1ConsumerShare) SetSharedResource(v ObjectReference)`

SetSharedResource sets SharedResource field to given value.

### HasSharedResource

`func (o *CdxV1ConsumerShare) HasSharedResource() bool`

HasSharedResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


