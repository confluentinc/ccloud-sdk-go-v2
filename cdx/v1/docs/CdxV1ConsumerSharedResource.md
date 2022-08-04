# CdxV1ConsumerSharedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**DisplayName** | Pointer to **string** | Consumer resource display name | [optional] 
**Description** | Pointer to **string** | Description of consumer resource | [optional] 
**Labels** | Pointer to **[]string** | list of labels | [optional] 
**Examples** | Pointer to **[]string** | List of example data in JSON format | [optional] 
**Schemas** | Pointer to **[]string** | List of schemas in JSON format | [optional] 
**OrganizationName** | Pointer to **string** | Organization to which the shared resource belongs | [optional] 
**OrganizationDetails** | Pointer to **string** | Details of the organization to which the shared resource belongs | [optional] 
**OrganizationContact** | Pointer to **string** | Email of contact person from the organization | [optional] 
**LogoUrl** | Pointer to **string** | Resource logo url | [optional] [readonly] 

## Methods

### NewCdxV1ConsumerSharedResource

`func NewCdxV1ConsumerSharedResource() *CdxV1ConsumerSharedResource`

NewCdxV1ConsumerSharedResource instantiates a new CdxV1ConsumerSharedResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdxV1ConsumerSharedResourceWithDefaults

`func NewCdxV1ConsumerSharedResourceWithDefaults() *CdxV1ConsumerSharedResource`

NewCdxV1ConsumerSharedResourceWithDefaults instantiates a new CdxV1ConsumerSharedResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CdxV1ConsumerSharedResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CdxV1ConsumerSharedResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CdxV1ConsumerSharedResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CdxV1ConsumerSharedResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CdxV1ConsumerSharedResource) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CdxV1ConsumerSharedResource) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CdxV1ConsumerSharedResource) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CdxV1ConsumerSharedResource) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CdxV1ConsumerSharedResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdxV1ConsumerSharedResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdxV1ConsumerSharedResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CdxV1ConsumerSharedResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CdxV1ConsumerSharedResource) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CdxV1ConsumerSharedResource) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CdxV1ConsumerSharedResource) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CdxV1ConsumerSharedResource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *CdxV1ConsumerSharedResource) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CdxV1ConsumerSharedResource) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CdxV1ConsumerSharedResource) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CdxV1ConsumerSharedResource) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *CdxV1ConsumerSharedResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CdxV1ConsumerSharedResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CdxV1ConsumerSharedResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CdxV1ConsumerSharedResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *CdxV1ConsumerSharedResource) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CdxV1ConsumerSharedResource) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CdxV1ConsumerSharedResource) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CdxV1ConsumerSharedResource) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetExamples

`func (o *CdxV1ConsumerSharedResource) GetExamples() []string`

GetExamples returns the Examples field if non-nil, zero value otherwise.

### GetExamplesOk

`func (o *CdxV1ConsumerSharedResource) GetExamplesOk() (*[]string, bool)`

GetExamplesOk returns a tuple with the Examples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamples

`func (o *CdxV1ConsumerSharedResource) SetExamples(v []string)`

SetExamples sets Examples field to given value.

### HasExamples

`func (o *CdxV1ConsumerSharedResource) HasExamples() bool`

HasExamples returns a boolean if a field has been set.

### GetSchemas

`func (o *CdxV1ConsumerSharedResource) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *CdxV1ConsumerSharedResource) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *CdxV1ConsumerSharedResource) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *CdxV1ConsumerSharedResource) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetOrganizationName

`func (o *CdxV1ConsumerSharedResource) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *CdxV1ConsumerSharedResource) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *CdxV1ConsumerSharedResource) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *CdxV1ConsumerSharedResource) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetOrganizationDetails

`func (o *CdxV1ConsumerSharedResource) GetOrganizationDetails() string`

GetOrganizationDetails returns the OrganizationDetails field if non-nil, zero value otherwise.

### GetOrganizationDetailsOk

`func (o *CdxV1ConsumerSharedResource) GetOrganizationDetailsOk() (*string, bool)`

GetOrganizationDetailsOk returns a tuple with the OrganizationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationDetails

`func (o *CdxV1ConsumerSharedResource) SetOrganizationDetails(v string)`

SetOrganizationDetails sets OrganizationDetails field to given value.

### HasOrganizationDetails

`func (o *CdxV1ConsumerSharedResource) HasOrganizationDetails() bool`

HasOrganizationDetails returns a boolean if a field has been set.

### GetOrganizationContact

`func (o *CdxV1ConsumerSharedResource) GetOrganizationContact() string`

GetOrganizationContact returns the OrganizationContact field if non-nil, zero value otherwise.

### GetOrganizationContactOk

`func (o *CdxV1ConsumerSharedResource) GetOrganizationContactOk() (*string, bool)`

GetOrganizationContactOk returns a tuple with the OrganizationContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationContact

`func (o *CdxV1ConsumerSharedResource) SetOrganizationContact(v string)`

SetOrganizationContact sets OrganizationContact field to given value.

### HasOrganizationContact

`func (o *CdxV1ConsumerSharedResource) HasOrganizationContact() bool`

HasOrganizationContact returns a boolean if a field has been set.

### GetLogoUrl

`func (o *CdxV1ConsumerSharedResource) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *CdxV1ConsumerSharedResource) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *CdxV1ConsumerSharedResource) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *CdxV1ConsumerSharedResource) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


