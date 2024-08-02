# ArtifactV1FlinkArtifactUpdate

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
**DisplayName** | Pointer to **string** | Display name of the Flink Artifact. | [optional] 
**ContentFormat** | Pointer to **string** | Archive format of the Flink Artifact. | [optional] 
**Description** | Pointer to **string** | Description of the Flink Artifact. | [optional] 
**DocumentationLink** | Pointer to **string** | Document link of the Flink Artifact. | [optional] 
**RuntimeLanguage** | Pointer to **string** | Runtime language of the Flink Artifact. | [optional] [default to "JAVA"]
**Versions** | Pointer to [**[]ArtifactV1FlinkArtifactVersion**](ArtifactV1FlinkArtifactVersion.md) | Versions associated with this Flink Artifact. | [optional] 

## Methods

### NewArtifactV1FlinkArtifactUpdate

`func NewArtifactV1FlinkArtifactUpdate() *ArtifactV1FlinkArtifactUpdate`

NewArtifactV1FlinkArtifactUpdate instantiates a new ArtifactV1FlinkArtifactUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactV1FlinkArtifactUpdateWithDefaults

`func NewArtifactV1FlinkArtifactUpdateWithDefaults() *ArtifactV1FlinkArtifactUpdate`

NewArtifactV1FlinkArtifactUpdateWithDefaults instantiates a new ArtifactV1FlinkArtifactUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ArtifactV1FlinkArtifactUpdate) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ArtifactV1FlinkArtifactUpdate) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ArtifactV1FlinkArtifactUpdate) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ArtifactV1FlinkArtifactUpdate) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ArtifactV1FlinkArtifactUpdate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ArtifactV1FlinkArtifactUpdate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ArtifactV1FlinkArtifactUpdate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ArtifactV1FlinkArtifactUpdate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *ArtifactV1FlinkArtifactUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArtifactV1FlinkArtifactUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArtifactV1FlinkArtifactUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArtifactV1FlinkArtifactUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *ArtifactV1FlinkArtifactUpdate) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ArtifactV1FlinkArtifactUpdate) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ArtifactV1FlinkArtifactUpdate) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ArtifactV1FlinkArtifactUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCloud

`func (o *ArtifactV1FlinkArtifactUpdate) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ArtifactV1FlinkArtifactUpdate) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ArtifactV1FlinkArtifactUpdate) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *ArtifactV1FlinkArtifactUpdate) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetRegion

`func (o *ArtifactV1FlinkArtifactUpdate) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ArtifactV1FlinkArtifactUpdate) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ArtifactV1FlinkArtifactUpdate) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ArtifactV1FlinkArtifactUpdate) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetEnvironment

`func (o *ArtifactV1FlinkArtifactUpdate) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ArtifactV1FlinkArtifactUpdate) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ArtifactV1FlinkArtifactUpdate) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ArtifactV1FlinkArtifactUpdate) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetDisplayName

`func (o *ArtifactV1FlinkArtifactUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ArtifactV1FlinkArtifactUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ArtifactV1FlinkArtifactUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ArtifactV1FlinkArtifactUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetContentFormat

`func (o *ArtifactV1FlinkArtifactUpdate) GetContentFormat() string`

GetContentFormat returns the ContentFormat field if non-nil, zero value otherwise.

### GetContentFormatOk

`func (o *ArtifactV1FlinkArtifactUpdate) GetContentFormatOk() (*string, bool)`

GetContentFormatOk returns a tuple with the ContentFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFormat

`func (o *ArtifactV1FlinkArtifactUpdate) SetContentFormat(v string)`

SetContentFormat sets ContentFormat field to given value.

### HasContentFormat

`func (o *ArtifactV1FlinkArtifactUpdate) HasContentFormat() bool`

HasContentFormat returns a boolean if a field has been set.

### GetDescription

`func (o *ArtifactV1FlinkArtifactUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArtifactV1FlinkArtifactUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArtifactV1FlinkArtifactUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArtifactV1FlinkArtifactUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentationLink

`func (o *ArtifactV1FlinkArtifactUpdate) GetDocumentationLink() string`

GetDocumentationLink returns the DocumentationLink field if non-nil, zero value otherwise.

### GetDocumentationLinkOk

`func (o *ArtifactV1FlinkArtifactUpdate) GetDocumentationLinkOk() (*string, bool)`

GetDocumentationLinkOk returns a tuple with the DocumentationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationLink

`func (o *ArtifactV1FlinkArtifactUpdate) SetDocumentationLink(v string)`

SetDocumentationLink sets DocumentationLink field to given value.

### HasDocumentationLink

`func (o *ArtifactV1FlinkArtifactUpdate) HasDocumentationLink() bool`

HasDocumentationLink returns a boolean if a field has been set.

### GetRuntimeLanguage

`func (o *ArtifactV1FlinkArtifactUpdate) GetRuntimeLanguage() string`

GetRuntimeLanguage returns the RuntimeLanguage field if non-nil, zero value otherwise.

### GetRuntimeLanguageOk

`func (o *ArtifactV1FlinkArtifactUpdate) GetRuntimeLanguageOk() (*string, bool)`

GetRuntimeLanguageOk returns a tuple with the RuntimeLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeLanguage

`func (o *ArtifactV1FlinkArtifactUpdate) SetRuntimeLanguage(v string)`

SetRuntimeLanguage sets RuntimeLanguage field to given value.

### HasRuntimeLanguage

`func (o *ArtifactV1FlinkArtifactUpdate) HasRuntimeLanguage() bool`

HasRuntimeLanguage returns a boolean if a field has been set.

### GetVersions

`func (o *ArtifactV1FlinkArtifactUpdate) GetVersions() []ArtifactV1FlinkArtifactVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ArtifactV1FlinkArtifactUpdate) GetVersionsOk() (*[]ArtifactV1FlinkArtifactVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ArtifactV1FlinkArtifactUpdate) SetVersions(v []ArtifactV1FlinkArtifactVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ArtifactV1FlinkArtifactUpdate) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


