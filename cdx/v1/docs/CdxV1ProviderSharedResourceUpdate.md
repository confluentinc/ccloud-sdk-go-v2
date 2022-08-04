# CdxV1ProviderSharedResourceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**DisplayName** | Pointer to **string** | Shared resource display name | [optional] 
**Description** | Pointer to **string** | Description of shared resource | [optional] 
**Labels** | Pointer to **[]string** | list of labels | [optional] 
**Examples** | Pointer to **[]string** | List of example data in JSON format | [optional] 
**Schemas** | Pointer to **[]string** | List of schemas in JSON format | [optional] 
**OrganizationName** | Pointer to **string** | Organization to which the shared resource belongs | [optional] 
**OrganizationDetails** | Pointer to **string** | Details of the organization to which the shared resource belongs | [optional] 
**OrganizationContact** | Pointer to **string** | Email of contact person from the organization | [optional] 
**EnvironmentName** | Pointer to **string** | The environment name of the shared resource | [optional] 
**ClusterName** | Pointer to **string** | The cluster display name of the shared resource | [optional] 
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

### GetDescription

`func (o *CdxV1ProviderSharedResourceUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CdxV1ProviderSharedResourceUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CdxV1ProviderSharedResourceUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CdxV1ProviderSharedResourceUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *CdxV1ProviderSharedResourceUpdate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CdxV1ProviderSharedResourceUpdate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CdxV1ProviderSharedResourceUpdate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CdxV1ProviderSharedResourceUpdate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetExamples

`func (o *CdxV1ProviderSharedResourceUpdate) GetExamples() []string`

GetExamples returns the Examples field if non-nil, zero value otherwise.

### GetExamplesOk

`func (o *CdxV1ProviderSharedResourceUpdate) GetExamplesOk() (*[]string, bool)`

GetExamplesOk returns a tuple with the Examples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamples

`func (o *CdxV1ProviderSharedResourceUpdate) SetExamples(v []string)`

SetExamples sets Examples field to given value.

### HasExamples

`func (o *CdxV1ProviderSharedResourceUpdate) HasExamples() bool`

HasExamples returns a boolean if a field has been set.

### GetSchemas

`func (o *CdxV1ProviderSharedResourceUpdate) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *CdxV1ProviderSharedResourceUpdate) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *CdxV1ProviderSharedResourceUpdate) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *CdxV1ProviderSharedResourceUpdate) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetOrganizationName

`func (o *CdxV1ProviderSharedResourceUpdate) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *CdxV1ProviderSharedResourceUpdate) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *CdxV1ProviderSharedResourceUpdate) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *CdxV1ProviderSharedResourceUpdate) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetOrganizationDetails

`func (o *CdxV1ProviderSharedResourceUpdate) GetOrganizationDetails() string`

GetOrganizationDetails returns the OrganizationDetails field if non-nil, zero value otherwise.

### GetOrganizationDetailsOk

`func (o *CdxV1ProviderSharedResourceUpdate) GetOrganizationDetailsOk() (*string, bool)`

GetOrganizationDetailsOk returns a tuple with the OrganizationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationDetails

`func (o *CdxV1ProviderSharedResourceUpdate) SetOrganizationDetails(v string)`

SetOrganizationDetails sets OrganizationDetails field to given value.

### HasOrganizationDetails

`func (o *CdxV1ProviderSharedResourceUpdate) HasOrganizationDetails() bool`

HasOrganizationDetails returns a boolean if a field has been set.

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

### GetEnvironmentName

`func (o *CdxV1ProviderSharedResourceUpdate) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *CdxV1ProviderSharedResourceUpdate) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *CdxV1ProviderSharedResourceUpdate) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.

### HasEnvironmentName

`func (o *CdxV1ProviderSharedResourceUpdate) HasEnvironmentName() bool`

HasEnvironmentName returns a boolean if a field has been set.

### GetClusterName

`func (o *CdxV1ProviderSharedResourceUpdate) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *CdxV1ProviderSharedResourceUpdate) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *CdxV1ProviderSharedResourceUpdate) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *CdxV1ProviderSharedResourceUpdate) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

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


