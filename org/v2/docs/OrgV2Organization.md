# OrgV2Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**DisplayName** | Pointer to **string** | A human-readable name for the Organization | [optional] 
**JitEnabled** | Pointer to **bool** | The flag to toggle Just-In-Time user provisioning for SSO-enabled organization. Available for early access only. | [optional] 
**ScimEnabled** | Pointer to **bool** | The flag to toggle SCIM user provisioning for an SSO-enabled organization. | [optional] 

## Methods

### NewOrgV2Organization

`func NewOrgV2Organization() *OrgV2Organization`

NewOrgV2Organization instantiates a new OrgV2Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgV2OrganizationWithDefaults

`func NewOrgV2OrganizationWithDefaults() *OrgV2Organization`

NewOrgV2OrganizationWithDefaults instantiates a new OrgV2Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *OrgV2Organization) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *OrgV2Organization) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *OrgV2Organization) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *OrgV2Organization) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *OrgV2Organization) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OrgV2Organization) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OrgV2Organization) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *OrgV2Organization) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *OrgV2Organization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgV2Organization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgV2Organization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrgV2Organization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *OrgV2Organization) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrgV2Organization) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrgV2Organization) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrgV2Organization) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *OrgV2Organization) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OrgV2Organization) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OrgV2Organization) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *OrgV2Organization) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetJitEnabled

`func (o *OrgV2Organization) GetJitEnabled() bool`

GetJitEnabled returns the JitEnabled field if non-nil, zero value otherwise.

### GetJitEnabledOk

`func (o *OrgV2Organization) GetJitEnabledOk() (*bool, bool)`

GetJitEnabledOk returns a tuple with the JitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitEnabled

`func (o *OrgV2Organization) SetJitEnabled(v bool)`

SetJitEnabled sets JitEnabled field to given value.

### HasJitEnabled

`func (o *OrgV2Organization) HasJitEnabled() bool`

HasJitEnabled returns a boolean if a field has been set.

### GetScimEnabled

`func (o *OrgV2Organization) GetScimEnabled() bool`

GetScimEnabled returns the ScimEnabled field if non-nil, zero value otherwise.

### GetScimEnabledOk

`func (o *OrgV2Organization) GetScimEnabledOk() (*bool, bool)`

GetScimEnabledOk returns a tuple with the ScimEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimEnabled

`func (o *OrgV2Organization) SetScimEnabled(v bool)`

SetScimEnabled sets ScimEnabled field to given value.

### HasScimEnabled

`func (o *OrgV2Organization) HasScimEnabled() bool`

HasScimEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


