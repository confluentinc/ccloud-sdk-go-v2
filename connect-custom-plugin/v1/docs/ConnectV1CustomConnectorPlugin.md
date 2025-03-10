# ConnectV1CustomConnectorPlugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**DisplayName** | Pointer to **string** | Display name of Custom Connector Plugin. | [optional] 
**ContentFormat** | Pointer to **string** | Archive format of Custom Connector Plugin. | [optional] [readonly] 
**Description** | Pointer to **string** | Description of Custom Connector Plugin. | [optional] 
**DocumentationLink** | Pointer to **string** | Document link of Custom Connector Plugin. | [optional] 
**ConnectorClass** | Pointer to **string** | Java class or alias for connector. You can get connector class from connector documentation provided by developer. | [optional] 
**ConnectorType** | Pointer to **string** | Custom Connector type.  | [optional] 
**Cloud** | Pointer to **string** | Cloud provider where the Custom Connector Plugin archive is uploaded. | [optional] [default to "AWS"]
**SensitiveConfigProperties** | Pointer to **[]string** | A sensitive property is a connector configuration property that must be hidden after a user enters property value when setting up connector. - If the plugin has **only one version**, these properties apply to that version. - If the plugin has **multiple versions**, each version maintains its own set of sensitive properties and does not inherit or use the &#x60;sensitive_config_properties&#x60; of the plugin.  | [optional] 
**UploadSource** | Pointer to [**ConnectV1CustomConnectorPluginUploadSourceOneOf**](ConnectV1CustomConnectorPluginUploadSourceOneOf.md) | Upload source of Custom Connector Plugin. Only required in &#x60;create&#x60; request, will be ignored in &#x60;read&#x60;, &#x60;update&#x60; or &#x60;list&#x60;. | [optional] 
**RuntimeLanguage** | Pointer to **string** | Runtime language of Custom Connector Plugin. | [optional] [default to "JAVA"]
**Version** | Pointer to **string** | Initial Version of the Custom Connector Plugin. The version must comply with SemVer (e.g., &#x60;1.2.3&#x60;, &#x60;1.2.3-beta&#x60;, &#x60;1.2.3-rc.123&#x60;, &#x60;1.2.3-rc.123+build.456&#x60;).  | [optional] [default to "0.0.0"]

## Methods

### NewConnectV1CustomConnectorPlugin

`func NewConnectV1CustomConnectorPlugin() *ConnectV1CustomConnectorPlugin`

NewConnectV1CustomConnectorPlugin instantiates a new ConnectV1CustomConnectorPlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1CustomConnectorPluginWithDefaults

`func NewConnectV1CustomConnectorPluginWithDefaults() *ConnectV1CustomConnectorPlugin`

NewConnectV1CustomConnectorPluginWithDefaults instantiates a new ConnectV1CustomConnectorPlugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ConnectV1CustomConnectorPlugin) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ConnectV1CustomConnectorPlugin) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ConnectV1CustomConnectorPlugin) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ConnectV1CustomConnectorPlugin) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ConnectV1CustomConnectorPlugin) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConnectV1CustomConnectorPlugin) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConnectV1CustomConnectorPlugin) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ConnectV1CustomConnectorPlugin) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *ConnectV1CustomConnectorPlugin) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectV1CustomConnectorPlugin) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectV1CustomConnectorPlugin) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectV1CustomConnectorPlugin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *ConnectV1CustomConnectorPlugin) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConnectV1CustomConnectorPlugin) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConnectV1CustomConnectorPlugin) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConnectV1CustomConnectorPlugin) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *ConnectV1CustomConnectorPlugin) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ConnectV1CustomConnectorPlugin) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ConnectV1CustomConnectorPlugin) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ConnectV1CustomConnectorPlugin) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetContentFormat

`func (o *ConnectV1CustomConnectorPlugin) GetContentFormat() string`

GetContentFormat returns the ContentFormat field if non-nil, zero value otherwise.

### GetContentFormatOk

`func (o *ConnectV1CustomConnectorPlugin) GetContentFormatOk() (*string, bool)`

GetContentFormatOk returns a tuple with the ContentFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFormat

`func (o *ConnectV1CustomConnectorPlugin) SetContentFormat(v string)`

SetContentFormat sets ContentFormat field to given value.

### HasContentFormat

`func (o *ConnectV1CustomConnectorPlugin) HasContentFormat() bool`

HasContentFormat returns a boolean if a field has been set.

### GetDescription

`func (o *ConnectV1CustomConnectorPlugin) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectV1CustomConnectorPlugin) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectV1CustomConnectorPlugin) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectV1CustomConnectorPlugin) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentationLink

`func (o *ConnectV1CustomConnectorPlugin) GetDocumentationLink() string`

GetDocumentationLink returns the DocumentationLink field if non-nil, zero value otherwise.

### GetDocumentationLinkOk

`func (o *ConnectV1CustomConnectorPlugin) GetDocumentationLinkOk() (*string, bool)`

GetDocumentationLinkOk returns a tuple with the DocumentationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationLink

`func (o *ConnectV1CustomConnectorPlugin) SetDocumentationLink(v string)`

SetDocumentationLink sets DocumentationLink field to given value.

### HasDocumentationLink

`func (o *ConnectV1CustomConnectorPlugin) HasDocumentationLink() bool`

HasDocumentationLink returns a boolean if a field has been set.

### GetConnectorClass

`func (o *ConnectV1CustomConnectorPlugin) GetConnectorClass() string`

GetConnectorClass returns the ConnectorClass field if non-nil, zero value otherwise.

### GetConnectorClassOk

`func (o *ConnectV1CustomConnectorPlugin) GetConnectorClassOk() (*string, bool)`

GetConnectorClassOk returns a tuple with the ConnectorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorClass

`func (o *ConnectV1CustomConnectorPlugin) SetConnectorClass(v string)`

SetConnectorClass sets ConnectorClass field to given value.

### HasConnectorClass

`func (o *ConnectV1CustomConnectorPlugin) HasConnectorClass() bool`

HasConnectorClass returns a boolean if a field has been set.

### GetConnectorType

`func (o *ConnectV1CustomConnectorPlugin) GetConnectorType() string`

GetConnectorType returns the ConnectorType field if non-nil, zero value otherwise.

### GetConnectorTypeOk

`func (o *ConnectV1CustomConnectorPlugin) GetConnectorTypeOk() (*string, bool)`

GetConnectorTypeOk returns a tuple with the ConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorType

`func (o *ConnectV1CustomConnectorPlugin) SetConnectorType(v string)`

SetConnectorType sets ConnectorType field to given value.

### HasConnectorType

`func (o *ConnectV1CustomConnectorPlugin) HasConnectorType() bool`

HasConnectorType returns a boolean if a field has been set.

### GetCloud

`func (o *ConnectV1CustomConnectorPlugin) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ConnectV1CustomConnectorPlugin) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ConnectV1CustomConnectorPlugin) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *ConnectV1CustomConnectorPlugin) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetSensitiveConfigProperties

`func (o *ConnectV1CustomConnectorPlugin) GetSensitiveConfigProperties() []string`

GetSensitiveConfigProperties returns the SensitiveConfigProperties field if non-nil, zero value otherwise.

### GetSensitiveConfigPropertiesOk

`func (o *ConnectV1CustomConnectorPlugin) GetSensitiveConfigPropertiesOk() (*[]string, bool)`

GetSensitiveConfigPropertiesOk returns a tuple with the SensitiveConfigProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveConfigProperties

`func (o *ConnectV1CustomConnectorPlugin) SetSensitiveConfigProperties(v []string)`

SetSensitiveConfigProperties sets SensitiveConfigProperties field to given value.

### HasSensitiveConfigProperties

`func (o *ConnectV1CustomConnectorPlugin) HasSensitiveConfigProperties() bool`

HasSensitiveConfigProperties returns a boolean if a field has been set.

### GetUploadSource

`func (o *ConnectV1CustomConnectorPlugin) GetUploadSource() ConnectV1CustomConnectorPluginUploadSourceOneOf`

GetUploadSource returns the UploadSource field if non-nil, zero value otherwise.

### GetUploadSourceOk

`func (o *ConnectV1CustomConnectorPlugin) GetUploadSourceOk() (*ConnectV1CustomConnectorPluginUploadSourceOneOf, bool)`

GetUploadSourceOk returns a tuple with the UploadSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadSource

`func (o *ConnectV1CustomConnectorPlugin) SetUploadSource(v ConnectV1CustomConnectorPluginUploadSourceOneOf)`

SetUploadSource sets UploadSource field to given value.

### HasUploadSource

`func (o *ConnectV1CustomConnectorPlugin) HasUploadSource() bool`

HasUploadSource returns a boolean if a field has been set.

### GetRuntimeLanguage

`func (o *ConnectV1CustomConnectorPlugin) GetRuntimeLanguage() string`

GetRuntimeLanguage returns the RuntimeLanguage field if non-nil, zero value otherwise.

### GetRuntimeLanguageOk

`func (o *ConnectV1CustomConnectorPlugin) GetRuntimeLanguageOk() (*string, bool)`

GetRuntimeLanguageOk returns a tuple with the RuntimeLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeLanguage

`func (o *ConnectV1CustomConnectorPlugin) SetRuntimeLanguage(v string)`

SetRuntimeLanguage sets RuntimeLanguage field to given value.

### HasRuntimeLanguage

`func (o *ConnectV1CustomConnectorPlugin) HasRuntimeLanguage() bool`

HasRuntimeLanguage returns a boolean if a field has been set.

### GetVersion

`func (o *ConnectV1CustomConnectorPlugin) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConnectV1CustomConnectorPlugin) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConnectV1CustomConnectorPlugin) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConnectV1CustomConnectorPlugin) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


