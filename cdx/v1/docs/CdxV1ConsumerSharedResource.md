# CdxV1ConsumerSharedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Cloud** | Pointer to **string** |  | [optional] [readonly] 
**NetworkConnectionTypes** | Pointer to [**CdxV1ConnectionTypes**](cdx.v1.ConnectionTypes.md) |  | [optional] [readonly] 
**DisplayName** | Pointer to **string** | Consumer resource display name | [optional] [readonly] 
**Description** | Pointer to **string** | Description of consumer resource | [optional] [readonly] 
**Labels** | Pointer to **[]string** | list of labels | [optional] [readonly] 
**Examples** | Pointer to **[]string** | List of example data in JSON format. This field is work in progress and subject to changes. | [optional] [readonly] 
**Schemas** | Pointer to **[]string** | List of schemas in JSON format. This field is work in progress and subject to changes. | [optional] [readonly] 
**OrganizationName** | Pointer to **string** | Shared resource&#39;s organization name | [optional] [readonly] 
**OrganizationDescription** | Pointer to **string** | Shared resource&#39;s organization description | [optional] [readonly] 
**OrganizationContact** | Pointer to **string** | Email of the shared resource&#39;s organization contact | [optional] [readonly] 
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

### GetCloud

`func (o *CdxV1ConsumerSharedResource) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *CdxV1ConsumerSharedResource) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *CdxV1ConsumerSharedResource) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *CdxV1ConsumerSharedResource) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetNetworkConnectionTypes

`func (o *CdxV1ConsumerSharedResource) GetNetworkConnectionTypes() CdxV1ConnectionTypes`

GetNetworkConnectionTypes returns the NetworkConnectionTypes field if non-nil, zero value otherwise.

### GetNetworkConnectionTypesOk

`func (o *CdxV1ConsumerSharedResource) GetNetworkConnectionTypesOk() (*CdxV1ConnectionTypes, bool)`

GetNetworkConnectionTypesOk returns a tuple with the NetworkConnectionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConnectionTypes

`func (o *CdxV1ConsumerSharedResource) SetNetworkConnectionTypes(v CdxV1ConnectionTypes)`

SetNetworkConnectionTypes sets NetworkConnectionTypes field to given value.

### HasNetworkConnectionTypes

`func (o *CdxV1ConsumerSharedResource) HasNetworkConnectionTypes() bool`

HasNetworkConnectionTypes returns a boolean if a field has been set.

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

### GetOrganizationDescription

`func (o *CdxV1ConsumerSharedResource) GetOrganizationDescription() string`

GetOrganizationDescription returns the OrganizationDescription field if non-nil, zero value otherwise.

### GetOrganizationDescriptionOk

`func (o *CdxV1ConsumerSharedResource) GetOrganizationDescriptionOk() (*string, bool)`

GetOrganizationDescriptionOk returns a tuple with the OrganizationDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationDescription

`func (o *CdxV1ConsumerSharedResource) SetOrganizationDescription(v string)`

SetOrganizationDescription sets OrganizationDescription field to given value.

### HasOrganizationDescription

`func (o *CdxV1ConsumerSharedResource) HasOrganizationDescription() bool`

HasOrganizationDescription returns a boolean if a field has been set.

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


