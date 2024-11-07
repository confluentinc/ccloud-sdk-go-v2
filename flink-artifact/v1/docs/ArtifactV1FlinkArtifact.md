# ArtifactV1FlinkArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Cloud** | Pointer to **string** | Cloud provider where the Flink Artifact archive is uploaded. | [optional] 
**Region** | Pointer to **string** | The Cloud provider region the Flink Artifact archive is uploaded. | [optional] 
**Environment** | Pointer to **string** | Environment the Flink Artifact belongs to. | [optional] 
**UniqueName** | Pointer to **string** | Unique name of the Flink Artifact per cloud, region, environment scope. | [optional] 
**Class** | Pointer to **string** | Java class or alias for the artifact as provided by developer. Deprecated | [optional] 
**ContentFormat** | Pointer to **string** | Archive format of the Flink Artifact. | [optional] 
**Description** | Pointer to **string** | Description of the Flink Artifact. | [optional] 
**DocumentationLink** | Pointer to **string** | Documentation link of the Flink Artifact. | [optional] 
**RuntimeLanguage** | Pointer to **string** | Runtime language of the Flink Artifact. | [optional] [default to "JAVA"]
**Versions** | Pointer to [**[]ArtifactV1FlinkArtifactVersion**](ArtifactV1FlinkArtifactVersion.md) | Versions associated with this Flink Artifact. | [optional] 

## Methods

### NewArtifactV1FlinkArtifact

`func NewArtifactV1FlinkArtifact() *ArtifactV1FlinkArtifact`

NewArtifactV1FlinkArtifact instantiates a new ArtifactV1FlinkArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactV1FlinkArtifactWithDefaults

`func NewArtifactV1FlinkArtifactWithDefaults() *ArtifactV1FlinkArtifact`

NewArtifactV1FlinkArtifactWithDefaults instantiates a new ArtifactV1FlinkArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ArtifactV1FlinkArtifact) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ArtifactV1FlinkArtifact) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ArtifactV1FlinkArtifact) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ArtifactV1FlinkArtifact) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ArtifactV1FlinkArtifact) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ArtifactV1FlinkArtifact) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ArtifactV1FlinkArtifact) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ArtifactV1FlinkArtifact) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *ArtifactV1FlinkArtifact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArtifactV1FlinkArtifact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArtifactV1FlinkArtifact) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArtifactV1FlinkArtifact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *ArtifactV1FlinkArtifact) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ArtifactV1FlinkArtifact) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ArtifactV1FlinkArtifact) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ArtifactV1FlinkArtifact) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCloud

`func (o *ArtifactV1FlinkArtifact) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ArtifactV1FlinkArtifact) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ArtifactV1FlinkArtifact) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *ArtifactV1FlinkArtifact) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetRegion

`func (o *ArtifactV1FlinkArtifact) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ArtifactV1FlinkArtifact) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ArtifactV1FlinkArtifact) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ArtifactV1FlinkArtifact) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetEnvironment

`func (o *ArtifactV1FlinkArtifact) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ArtifactV1FlinkArtifact) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ArtifactV1FlinkArtifact) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ArtifactV1FlinkArtifact) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetUniqueName

`func (o *ArtifactV1FlinkArtifact) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *ArtifactV1FlinkArtifact) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *ArtifactV1FlinkArtifact) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *ArtifactV1FlinkArtifact) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### GetClass

`func (o *ArtifactV1FlinkArtifact) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *ArtifactV1FlinkArtifact) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *ArtifactV1FlinkArtifact) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *ArtifactV1FlinkArtifact) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetContentFormat

`func (o *ArtifactV1FlinkArtifact) GetContentFormat() string`

GetContentFormat returns the ContentFormat field if non-nil, zero value otherwise.

### GetContentFormatOk

`func (o *ArtifactV1FlinkArtifact) GetContentFormatOk() (*string, bool)`

GetContentFormatOk returns a tuple with the ContentFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFormat

`func (o *ArtifactV1FlinkArtifact) SetContentFormat(v string)`

SetContentFormat sets ContentFormat field to given value.

### HasContentFormat

`func (o *ArtifactV1FlinkArtifact) HasContentFormat() bool`

HasContentFormat returns a boolean if a field has been set.

### GetDescription

`func (o *ArtifactV1FlinkArtifact) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArtifactV1FlinkArtifact) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArtifactV1FlinkArtifact) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArtifactV1FlinkArtifact) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentationLink

`func (o *ArtifactV1FlinkArtifact) GetDocumentationLink() string`

GetDocumentationLink returns the DocumentationLink field if non-nil, zero value otherwise.

### GetDocumentationLinkOk

`func (o *ArtifactV1FlinkArtifact) GetDocumentationLinkOk() (*string, bool)`

GetDocumentationLinkOk returns a tuple with the DocumentationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationLink

`func (o *ArtifactV1FlinkArtifact) SetDocumentationLink(v string)`

SetDocumentationLink sets DocumentationLink field to given value.

### HasDocumentationLink

`func (o *ArtifactV1FlinkArtifact) HasDocumentationLink() bool`

HasDocumentationLink returns a boolean if a field has been set.

### GetRuntimeLanguage

`func (o *ArtifactV1FlinkArtifact) GetRuntimeLanguage() string`

GetRuntimeLanguage returns the RuntimeLanguage field if non-nil, zero value otherwise.

### GetRuntimeLanguageOk

`func (o *ArtifactV1FlinkArtifact) GetRuntimeLanguageOk() (*string, bool)`

GetRuntimeLanguageOk returns a tuple with the RuntimeLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeLanguage

`func (o *ArtifactV1FlinkArtifact) SetRuntimeLanguage(v string)`

SetRuntimeLanguage sets RuntimeLanguage field to given value.

### HasRuntimeLanguage

`func (o *ArtifactV1FlinkArtifact) HasRuntimeLanguage() bool`

HasRuntimeLanguage returns a boolean if a field has been set.

### GetVersions

`func (o *ArtifactV1FlinkArtifact) GetVersions() []ArtifactV1FlinkArtifactVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ArtifactV1FlinkArtifact) GetVersionsOk() (*[]ArtifactV1FlinkArtifactVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ArtifactV1FlinkArtifact) SetVersions(v []ArtifactV1FlinkArtifactVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ArtifactV1FlinkArtifact) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


