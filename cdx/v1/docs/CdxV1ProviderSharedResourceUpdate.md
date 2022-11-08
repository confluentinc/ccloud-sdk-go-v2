# CdxV1ProviderSharedResourceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Resources** | Pointer to **[]string** | List of resource crns that are shared together | [optional] 
**DisplayName** | Pointer to **string** | Shared resource display name | [optional] 
**OrganizationDescription** | Pointer to **string** | Shared resource&#39;s organization description | [optional] 
**OrganizationContact** | Pointer to **string** | Email of contact person from the organization | [optional] 
**LogoUrl** | Pointer to **string** | Resource logo url | [optional] [readonly] 

## Methods

### NewCdxV1ProviderSharedResourceUpdate

`func NewCdxV1ProviderSharedResourceUpdate() *CdxV1ProviderSharedResourceUpdate`

NewCdxV1ProviderSharedResourceUpdate instantiates a new CdxV1ProviderSharedResourceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdxV1ProviderSharedResourceUpdateWithDefaults

`func NewCdxV1ProviderSharedResourceUpdateWithDefaults() *CdxV1ProviderSharedResourceUpdate`

NewCdxV1ProviderSharedResourceUpdateWithDefaults instantiates a new CdxV1ProviderSharedResourceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CdxV1ProviderSharedResourceUpdate) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CdxV1ProviderSharedResourceUpdate) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CdxV1ProviderSharedResourceUpdate) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CdxV1ProviderSharedResourceUpdate) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CdxV1ProviderSharedResourceUpdate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CdxV1ProviderSharedResourceUpdate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CdxV1ProviderSharedResourceUpdate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CdxV1ProviderSharedResourceUpdate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CdxV1ProviderSharedResourceUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdxV1ProviderSharedResourceUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdxV1ProviderSharedResourceUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CdxV1ProviderSharedResourceUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CdxV1ProviderSharedResourceUpdate) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CdxV1ProviderSharedResourceUpdate) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CdxV1ProviderSharedResourceUpdate) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CdxV1ProviderSharedResourceUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResources

`func (o *CdxV1ProviderSharedResourceUpdate) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CdxV1ProviderSharedResourceUpdate) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CdxV1ProviderSharedResourceUpdate) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *CdxV1ProviderSharedResourceUpdate) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetDisplayName

`func (o *CdxV1ProviderSharedResourceUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CdxV1ProviderSharedResourceUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CdxV1ProviderSharedResourceUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CdxV1ProviderSharedResourceUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetOrganizationDescription

`func (o *CdxV1ProviderSharedResourceUpdate) GetOrganizationDescription() string`

GetOrganizationDescription returns the OrganizationDescription field if non-nil, zero value otherwise.

### GetOrganizationDescriptionOk

`func (o *CdxV1ProviderSharedResourceUpdate) GetOrganizationDescriptionOk() (*string, bool)`

GetOrganizationDescriptionOk returns a tuple with the OrganizationDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationDescription

`func (o *CdxV1ProviderSharedResourceUpdate) SetOrganizationDescription(v string)`

SetOrganizationDescription sets OrganizationDescription field to given value.

### HasOrganizationDescription

`func (o *CdxV1ProviderSharedResourceUpdate) HasOrganizationDescription() bool`

HasOrganizationDescription returns a boolean if a field has been set.

### GetOrganizationContact

`func (o *CdxV1ProviderSharedResourceUpdate) GetOrganizationContact() string`

GetOrganizationContact returns the OrganizationContact field if non-nil, zero value otherwise.

### GetOrganizationContactOk

`func (o *CdxV1ProviderSharedResourceUpdate) GetOrganizationContactOk() (*string, bool)`

GetOrganizationContactOk returns a tuple with the OrganizationContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationContact

`func (o *CdxV1ProviderSharedResourceUpdate) SetOrganizationContact(v string)`

SetOrganizationContact sets OrganizationContact field to given value.

### HasOrganizationContact

`func (o *CdxV1ProviderSharedResourceUpdate) HasOrganizationContact() bool`

HasOrganizationContact returns a boolean if a field has been set.

### GetLogoUrl

`func (o *CdxV1ProviderSharedResourceUpdate) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *CdxV1ProviderSharedResourceUpdate) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *CdxV1ProviderSharedResourceUpdate) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *CdxV1ProviderSharedResourceUpdate) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


