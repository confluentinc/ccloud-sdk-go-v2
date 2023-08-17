# CcpV1CustomPlugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**DisplayName** | Pointer to **string** | Display name of custom plugin. | [optional] 
**ContentFormat** | Pointer to **string** | Archive format of custom plugin. | [optional] [readonly] 
**Description** | Pointer to **string** | Description of custom plugin. | [optional] 
**DocumentationLink** | Pointer to **string** | Document link of custom plugin. | [optional] 
**Config** | Pointer to [**CcpV1CustomPluginConfigOneOf**](CcpV1CustomPluginConfigOneOf.md) | Plugin-specific config based on plugin type. | [optional] 
**UploadSource** | Pointer to [**CcpV1CustomPluginUploadSourceOneOf**](CcpV1CustomPluginUploadSourceOneOf.md) | Upload source of custom plugin. Only required in &#x60;create&#x60; request, not in &#x60;read&#x60;, &#x60;update&#x60; or &#x60;list&#x60;. | [optional] 

## Methods

### NewCcpV1CustomPlugin

`func NewCcpV1CustomPlugin() *CcpV1CustomPlugin`

NewCcpV1CustomPlugin instantiates a new CcpV1CustomPlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCcpV1CustomPluginWithDefaults

`func NewCcpV1CustomPluginWithDefaults() *CcpV1CustomPlugin`

NewCcpV1CustomPluginWithDefaults instantiates a new CcpV1CustomPlugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CcpV1CustomPlugin) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CcpV1CustomPlugin) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CcpV1CustomPlugin) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CcpV1CustomPlugin) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CcpV1CustomPlugin) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CcpV1CustomPlugin) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CcpV1CustomPlugin) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CcpV1CustomPlugin) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CcpV1CustomPlugin) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CcpV1CustomPlugin) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CcpV1CustomPlugin) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CcpV1CustomPlugin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CcpV1CustomPlugin) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CcpV1CustomPlugin) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CcpV1CustomPlugin) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CcpV1CustomPlugin) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *CcpV1CustomPlugin) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CcpV1CustomPlugin) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CcpV1CustomPlugin) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CcpV1CustomPlugin) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetContentFormat

`func (o *CcpV1CustomPlugin) GetContentFormat() string`

GetContentFormat returns the ContentFormat field if non-nil, zero value otherwise.

### GetContentFormatOk

`func (o *CcpV1CustomPlugin) GetContentFormatOk() (*string, bool)`

GetContentFormatOk returns a tuple with the ContentFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFormat

`func (o *CcpV1CustomPlugin) SetContentFormat(v string)`

SetContentFormat sets ContentFormat field to given value.

### HasContentFormat

`func (o *CcpV1CustomPlugin) HasContentFormat() bool`

HasContentFormat returns a boolean if a field has been set.

### GetDescription

`func (o *CcpV1CustomPlugin) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CcpV1CustomPlugin) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CcpV1CustomPlugin) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CcpV1CustomPlugin) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentationLink

`func (o *CcpV1CustomPlugin) GetDocumentationLink() string`

GetDocumentationLink returns the DocumentationLink field if non-nil, zero value otherwise.

### GetDocumentationLinkOk

`func (o *CcpV1CustomPlugin) GetDocumentationLinkOk() (*string, bool)`

GetDocumentationLinkOk returns a tuple with the DocumentationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationLink

`func (o *CcpV1CustomPlugin) SetDocumentationLink(v string)`

SetDocumentationLink sets DocumentationLink field to given value.

### HasDocumentationLink

`func (o *CcpV1CustomPlugin) HasDocumentationLink() bool`

HasDocumentationLink returns a boolean if a field has been set.

### GetConfig

`func (o *CcpV1CustomPlugin) GetConfig() CcpV1CustomPluginConfigOneOf`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CcpV1CustomPlugin) GetConfigOk() (*CcpV1CustomPluginConfigOneOf, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CcpV1CustomPlugin) SetConfig(v CcpV1CustomPluginConfigOneOf)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CcpV1CustomPlugin) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetUploadSource

`func (o *CcpV1CustomPlugin) GetUploadSource() CcpV1CustomPluginUploadSourceOneOf`

GetUploadSource returns the UploadSource field if non-nil, zero value otherwise.

### GetUploadSourceOk

`func (o *CcpV1CustomPlugin) GetUploadSourceOk() (*CcpV1CustomPluginUploadSourceOneOf, bool)`

GetUploadSourceOk returns a tuple with the UploadSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadSource

`func (o *CcpV1CustomPlugin) SetUploadSource(v CcpV1CustomPluginUploadSourceOneOf)`

SetUploadSource sets UploadSource field to given value.

### HasUploadSource

`func (o *CcpV1CustomPlugin) HasUploadSource() bool`

HasUploadSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


