# CdxV1ProviderSharedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Crn** | Pointer to **string** | Deprecated please use resources attribute. | [optional] 
**Resources** | Pointer to **[]string** | List of resource crns that are shared together | [optional] 
**DisplayName** | Pointer to **string** | Shared resource display name | [optional] 
**Description** | Pointer to **string** | Description of shared resource | [optional] [readonly] 
**Labels** | Pointer to **[]string** | list of labels | [optional] [readonly] 
**Schemas** | Pointer to **[]string** | List of schemas in JSON format. This field is work in progress and subject to changes. | [optional] [readonly] 
**OrganizationDescription** | Pointer to **string** | Shared resource&#39;s organization description | [optional] 
**OrganizationContact** | Pointer to **string** | Email of contact person from the organization | [optional] 
**LogoUrl** | Pointer to **string** | Resource logo url | [optional] [readonly] 
**OrganizationName** | Pointer to **interface{}** | Organization to which the shared resource belongs. Deprecated | [optional] [readonly] 
**EnvironmentName** | Pointer to **string** | The environment name of the shared resource. Deprecated | [optional] [readonly] 
**ClusterName** | Pointer to **string** | The cluster display name of the shared resource. Deprecated | [optional] [readonly] 
**CloudCluster** | Pointer to [**ObjectReference**](ObjectReference.md) | The cloud cluster to which this belongs. | [optional] 

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

### GetResources

`func (o *CdxV1ProviderSharedResource) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CdxV1ProviderSharedResource) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CdxV1ProviderSharedResource) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *CdxV1ProviderSharedResource) HasResources() bool`

HasResources returns a boolean if a field has been set.

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

### GetOrganizationDescription

`func (o *CdxV1ProviderSharedResource) GetOrganizationDescription() string`

GetOrganizationDescription returns the OrganizationDescription field if non-nil, zero value otherwise.

### GetOrganizationDescriptionOk

`func (o *CdxV1ProviderSharedResource) GetOrganizationDescriptionOk() (*string, bool)`

GetOrganizationDescriptionOk returns a tuple with the OrganizationDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationDescription

`func (o *CdxV1ProviderSharedResource) SetOrganizationDescription(v string)`

SetOrganizationDescription sets OrganizationDescription field to given value.

### HasOrganizationDescription

`func (o *CdxV1ProviderSharedResource) HasOrganizationDescription() bool`

HasOrganizationDescription returns a boolean if a field has been set.

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

### GetOrganizationName

`func (o *CdxV1ProviderSharedResource) GetOrganizationName() interface{}`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *CdxV1ProviderSharedResource) GetOrganizationNameOk() (*interface{}, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *CdxV1ProviderSharedResource) SetOrganizationName(v interface{})`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *CdxV1ProviderSharedResource) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *CdxV1ProviderSharedResource) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *CdxV1ProviderSharedResource) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
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

### GetCloudCluster

`func (o *CdxV1ProviderSharedResource) GetCloudCluster() ObjectReference`

GetCloudCluster returns the CloudCluster field if non-nil, zero value otherwise.

### GetCloudClusterOk

`func (o *CdxV1ProviderSharedResource) GetCloudClusterOk() (*ObjectReference, bool)`

GetCloudClusterOk returns a tuple with the CloudCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCluster

`func (o *CdxV1ProviderSharedResource) SetCloudCluster(v ObjectReference)`

SetCloudCluster sets CloudCluster field to given value.

### HasCloudCluster

`func (o *CdxV1ProviderSharedResource) HasCloudCluster() bool`

HasCloudCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


