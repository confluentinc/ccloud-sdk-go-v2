# CdxV1ProviderSharedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Crn** | Pointer to **string** | A CRN that specifies the shared resource | [optional] 
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

### NewCdxV1ProviderSharedResource

`func NewCdxV1ProviderSharedResource() *CdxV1ProviderSharedResource`

NewCdxV1ProviderSharedResource instantiates a new CdxV1ProviderSharedResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdxV1ProviderSharedResourceWithDefaults

`func NewCdxV1ProviderSharedResourceWithDefaults() *CdxV1ProviderSharedResource`

NewCdxV1ProviderSharedResourceWithDefaults instantiates a new CdxV1ProviderSharedResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CdxV1ProviderSharedResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CdxV1ProviderSharedResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CdxV1ProviderSharedResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CdxV1ProviderSharedResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CdxV1ProviderSharedResource) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CdxV1ProviderSharedResource) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CdxV1ProviderSharedResource) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CdxV1ProviderSharedResource) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CdxV1ProviderSharedResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdxV1ProviderSharedResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdxV1ProviderSharedResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CdxV1ProviderSharedResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CdxV1ProviderSharedResource) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CdxV1ProviderSharedResource) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CdxV1ProviderSharedResource) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CdxV1ProviderSharedResource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCrn

`func (o *CdxV1ProviderSharedResource) GetCrn() string`

GetCrn returns the Crn field if non-nil, zero value otherwise.

### GetCrnOk

`func (o *CdxV1ProviderSharedResource) GetCrnOk() (*string, bool)`

GetCrnOk returns a tuple with the Crn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrn

`func (o *CdxV1ProviderSharedResource) SetCrn(v string)`

SetCrn sets Crn field to given value.

### HasCrn

`func (o *CdxV1ProviderSharedResource) HasCrn() bool`

HasCrn returns a boolean if a field has been set.

### GetDisplayName

`func (o *CdxV1ProviderSharedResource) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CdxV1ProviderSharedResource) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CdxV1ProviderSharedResource) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CdxV1ProviderSharedResource) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *CdxV1ProviderSharedResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CdxV1ProviderSharedResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CdxV1ProviderSharedResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CdxV1ProviderSharedResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *CdxV1ProviderSharedResource) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CdxV1ProviderSharedResource) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CdxV1ProviderSharedResource) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CdxV1ProviderSharedResource) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetExamples

`func (o *CdxV1ProviderSharedResource) GetExamples() []string`

GetExamples returns the Examples field if non-nil, zero value otherwise.

### GetExamplesOk

`func (o *CdxV1ProviderSharedResource) GetExamplesOk() (*[]string, bool)`

GetExamplesOk returns a tuple with the Examples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamples

`func (o *CdxV1ProviderSharedResource) SetExamples(v []string)`

SetExamples sets Examples field to given value.

### HasExamples

`func (o *CdxV1ProviderSharedResource) HasExamples() bool`

HasExamples returns a boolean if a field has been set.

### GetSchemas

`func (o *CdxV1ProviderSharedResource) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *CdxV1ProviderSharedResource) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *CdxV1ProviderSharedResource) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *CdxV1ProviderSharedResource) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetOrganizationName

`func (o *CdxV1ProviderSharedResource) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *CdxV1ProviderSharedResource) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *CdxV1ProviderSharedResource) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *CdxV1ProviderSharedResource) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetOrganizationDetails

`func (o *CdxV1ProviderSharedResource) GetOrganizationDetails() string`

GetOrganizationDetails returns the OrganizationDetails field if non-nil, zero value otherwise.

### GetOrganizationDetailsOk

`func (o *CdxV1ProviderSharedResource) GetOrganizationDetailsOk() (*string, bool)`

GetOrganizationDetailsOk returns a tuple with the OrganizationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationDetails

`func (o *CdxV1ProviderSharedResource) SetOrganizationDetails(v string)`

SetOrganizationDetails sets OrganizationDetails field to given value.

### HasOrganizationDetails

`func (o *CdxV1ProviderSharedResource) HasOrganizationDetails() bool`

HasOrganizationDetails returns a boolean if a field has been set.

### GetOrganizationContact

`func (o *CdxV1ProviderSharedResource) GetOrganizationContact() string`

GetOrganizationContact returns the OrganizationContact field if non-nil, zero value otherwise.

### GetOrganizationContactOk

`func (o *CdxV1ProviderSharedResource) GetOrganizationContactOk() (*string, bool)`

GetOrganizationContactOk returns a tuple with the OrganizationContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationContact

`func (o *CdxV1ProviderSharedResource) SetOrganizationContact(v string)`

SetOrganizationContact sets OrganizationContact field to given value.

### HasOrganizationContact

`func (o *CdxV1ProviderSharedResource) HasOrganizationContact() bool`

HasOrganizationContact returns a boolean if a field has been set.

### GetEnvironmentName

`func (o *CdxV1ProviderSharedResource) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *CdxV1ProviderSharedResource) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *CdxV1ProviderSharedResource) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.

### HasEnvironmentName

`func (o *CdxV1ProviderSharedResource) HasEnvironmentName() bool`

HasEnvironmentName returns a boolean if a field has been set.

### GetClusterName

`func (o *CdxV1ProviderSharedResource) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *CdxV1ProviderSharedResource) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *CdxV1ProviderSharedResource) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *CdxV1ProviderSharedResource) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetLogoUrl

`func (o *CdxV1ProviderSharedResource) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *CdxV1ProviderSharedResource) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *CdxV1ProviderSharedResource) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *CdxV1ProviderSharedResource) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


