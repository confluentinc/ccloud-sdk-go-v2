# QuotasV2AppliedQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Scope** | Pointer to **string** | The applied scope that this quota belongs to | [optional] 
**DisplayName** | Pointer to **string** | A human-readable name for the quota type name | [optional] 
**Code** | Pointer to **string** | The single quota type id that user want to query. If you only want to query a single specific quota, this is required.  | [optional] 
**AppliedLimit** | Pointer to **int32** | The latest applied service quota value, taking into account any limit adjustments.  | [optional] 
**User** | Pointer to [**ObjectReference**](ObjectReference.md) | The user to which this belongs. | [optional] 
**Organization** | Pointer to [**ObjectReference**](ObjectReference.md) | The organization associated with this object. | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 
**Cluster** | Pointer to [**ObjectReference**](ObjectReference.md) | The cluster to which this belongs. | [optional] 

## Methods

### NewQuotasV2AppliedQuota

`func NewQuotasV2AppliedQuota() *QuotasV2AppliedQuota`

NewQuotasV2AppliedQuota instantiates a new QuotasV2AppliedQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotasV2AppliedQuotaWithDefaults

`func NewQuotasV2AppliedQuotaWithDefaults() *QuotasV2AppliedQuota`

NewQuotasV2AppliedQuotaWithDefaults instantiates a new QuotasV2AppliedQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *QuotasV2AppliedQuota) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *QuotasV2AppliedQuota) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *QuotasV2AppliedQuota) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *QuotasV2AppliedQuota) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *QuotasV2AppliedQuota) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *QuotasV2AppliedQuota) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *QuotasV2AppliedQuota) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *QuotasV2AppliedQuota) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *QuotasV2AppliedQuota) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuotasV2AppliedQuota) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuotasV2AppliedQuota) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *QuotasV2AppliedQuota) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *QuotasV2AppliedQuota) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *QuotasV2AppliedQuota) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *QuotasV2AppliedQuota) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *QuotasV2AppliedQuota) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetScope

`func (o *QuotasV2AppliedQuota) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *QuotasV2AppliedQuota) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *QuotasV2AppliedQuota) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *QuotasV2AppliedQuota) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetDisplayName

`func (o *QuotasV2AppliedQuota) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *QuotasV2AppliedQuota) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *QuotasV2AppliedQuota) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *QuotasV2AppliedQuota) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCode

`func (o *QuotasV2AppliedQuota) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *QuotasV2AppliedQuota) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *QuotasV2AppliedQuota) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *QuotasV2AppliedQuota) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetAppliedLimit

`func (o *QuotasV2AppliedQuota) GetAppliedLimit() int32`

GetAppliedLimit returns the AppliedLimit field if non-nil, zero value otherwise.

### GetAppliedLimitOk

`func (o *QuotasV2AppliedQuota) GetAppliedLimitOk() (*int32, bool)`

GetAppliedLimitOk returns a tuple with the AppliedLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedLimit

`func (o *QuotasV2AppliedQuota) SetAppliedLimit(v int32)`

SetAppliedLimit sets AppliedLimit field to given value.

### HasAppliedLimit

`func (o *QuotasV2AppliedQuota) HasAppliedLimit() bool`

HasAppliedLimit returns a boolean if a field has been set.

### GetUser

`func (o *QuotasV2AppliedQuota) GetUser() ObjectReference`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *QuotasV2AppliedQuota) GetUserOk() (*ObjectReference, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *QuotasV2AppliedQuota) SetUser(v ObjectReference)`

SetUser sets User field to given value.

### HasUser

`func (o *QuotasV2AppliedQuota) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetOrganization

`func (o *QuotasV2AppliedQuota) GetOrganization() ObjectReference`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *QuotasV2AppliedQuota) GetOrganizationOk() (*ObjectReference, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *QuotasV2AppliedQuota) SetOrganization(v ObjectReference)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *QuotasV2AppliedQuota) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetEnvironment

`func (o *QuotasV2AppliedQuota) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *QuotasV2AppliedQuota) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *QuotasV2AppliedQuota) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *QuotasV2AppliedQuota) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetCluster

`func (o *QuotasV2AppliedQuota) GetCluster() ObjectReference`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *QuotasV2AppliedQuota) GetClusterOk() (*ObjectReference, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *QuotasV2AppliedQuota) SetCluster(v ObjectReference)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *QuotasV2AppliedQuota) HasCluster() bool`

HasCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


