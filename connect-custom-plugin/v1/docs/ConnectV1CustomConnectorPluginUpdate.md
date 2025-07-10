# ConnectV1CustomConnectorPluginUpdate

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
**SensitiveConfigProperties** | Pointer to **[]string** | A sensitive property is a connector configuration property that must be hidden after a user enters property value when setting up connector.  | [optional] 
**UploadSource** | Pointer to [**ConnectV1CustomConnectorPluginUpdateUploadSourceOneOf**](ConnectV1CustomConnectorPluginUpdateUploadSourceOneOf.md) | Upload source of Custom Connector Plugin. Only required in &#x60;create&#x60; request, will be ignored in &#x60;read&#x60;, &#x60;update&#x60; or &#x60;list&#x60;. | [optional] 
**RuntimeLanguage** | Pointer to **string** | Runtime language of Custom Connector Plugin. | [optional] [default to "JAVA"]

## Methods

### NewConnectV1CustomConnectorPluginUpdate

`func NewConnectV1CustomConnectorPluginUpdate() *ConnectV1CustomConnectorPluginUpdate`

NewConnectV1CustomConnectorPluginUpdate instantiates a new ConnectV1CustomConnectorPluginUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1CustomConnectorPluginUpdateWithDefaults

`func NewConnectV1CustomConnectorPluginUpdateWithDefaults() *ConnectV1CustomConnectorPluginUpdate`

NewConnectV1CustomConnectorPluginUpdateWithDefaults instantiates a new ConnectV1CustomConnectorPluginUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ConnectV1CustomConnectorPluginUpdate) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ConnectV1CustomConnectorPluginUpdate) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ConnectV1CustomConnectorPluginUpdate) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ConnectV1CustomConnectorPluginUpdate) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ConnectV1CustomConnectorPluginUpdate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConnectV1CustomConnectorPluginUpdate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConnectV1CustomConnectorPluginUpdate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ConnectV1CustomConnectorPluginUpdate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *ConnectV1CustomConnectorPluginUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectV1CustomConnectorPluginUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectV1CustomConnectorPluginUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectV1CustomConnectorPluginUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *ConnectV1CustomConnectorPluginUpdate) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConnectV1CustomConnectorPluginUpdate) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConnectV1CustomConnectorPluginUpdate) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConnectV1CustomConnectorPluginUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *ConnectV1CustomConnectorPluginUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ConnectV1CustomConnectorPluginUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ConnectV1CustomConnectorPluginUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ConnectV1CustomConnectorPluginUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetContentFormat

`func (o *ConnectV1CustomConnectorPluginUpdate) GetContentFormat() string`

GetContentFormat returns the ContentFormat field if non-nil, zero value otherwise.

### GetContentFormatOk

`func (o *ConnectV1CustomConnectorPluginUpdate) GetContentFormatOk() (*string, bool)`

GetContentFormatOk returns a tuple with the ContentFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFormat

`func (o *ConnectV1CustomConnectorPluginUpdate) SetContentFormat(v string)`

SetContentFormat sets ContentFormat field to given value.

### HasContentFormat

`func (o *ConnectV1CustomConnectorPluginUpdate) HasContentFormat() bool`

HasContentFormat returns a boolean if a field has been set.

### GetDescription

`func (o *ConnectV1CustomConnectorPluginUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectV1CustomConnectorPluginUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectV1CustomConnectorPluginUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectV1CustomConnectorPluginUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentationLink

`func (o *ConnectV1CustomConnectorPluginUpdate) GetDocumentationLink() string`

GetDocumentationLink returns the DocumentationLink field if non-nil, zero value otherwise.

### GetDocumentationLinkOk

`func (o *ConnectV1CustomConnectorPluginUpdate) GetDocumentationLinkOk() (*string, bool)`

GetDocumentationLinkOk returns a tuple with the DocumentationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationLink

`func (o *ConnectV1CustomConnectorPluginUpdate) SetDocumentationLink(v string)`

SetDocumentationLink sets DocumentationLink field to given value.

### HasDocumentationLink

`func (o *ConnectV1CustomConnectorPluginUpdate) HasDocumentationLink() bool`

HasDocumentationLink returns a boolean if a field has been set.

### GetSensitiveConfigProperties

`func (o *ConnectV1CustomConnectorPluginUpdate) GetSensitiveConfigProperties() []string`

GetSensitiveConfigProperties returns the SensitiveConfigProperties field if non-nil, zero value otherwise.

### GetSensitiveConfigPropertiesOk

`func (o *ConnectV1CustomConnectorPluginUpdate) GetSensitiveConfigPropertiesOk() (*[]string, bool)`

GetSensitiveConfigPropertiesOk returns a tuple with the SensitiveConfigProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveConfigProperties

`func (o *ConnectV1CustomConnectorPluginUpdate) SetSensitiveConfigProperties(v []string)`

SetSensitiveConfigProperties sets SensitiveConfigProperties field to given value.

### HasSensitiveConfigProperties

`func (o *ConnectV1CustomConnectorPluginUpdate) HasSensitiveConfigProperties() bool`

HasSensitiveConfigProperties returns a boolean if a field has been set.

### GetUploadSource

`func (o *ConnectV1CustomConnectorPluginUpdate) GetUploadSource() ConnectV1CustomConnectorPluginUpdateUploadSourceOneOf`

GetUploadSource returns the UploadSource field if non-nil, zero value otherwise.

### GetUploadSourceOk

`func (o *ConnectV1CustomConnectorPluginUpdate) GetUploadSourceOk() (*ConnectV1CustomConnectorPluginUpdateUploadSourceOneOf, bool)`

GetUploadSourceOk returns a tuple with the UploadSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadSource

`func (o *ConnectV1CustomConnectorPluginUpdate) SetUploadSource(v ConnectV1CustomConnectorPluginUpdateUploadSourceOneOf)`

SetUploadSource sets UploadSource field to given value.

### HasUploadSource

`func (o *ConnectV1CustomConnectorPluginUpdate) HasUploadSource() bool`

HasUploadSource returns a boolean if a field has been set.

### GetRuntimeLanguage

`func (o *ConnectV1CustomConnectorPluginUpdate) GetRuntimeLanguage() string`

GetRuntimeLanguage returns the RuntimeLanguage field if non-nil, zero value otherwise.

### GetRuntimeLanguageOk

`func (o *ConnectV1CustomConnectorPluginUpdate) GetRuntimeLanguageOk() (*string, bool)`

GetRuntimeLanguageOk returns a tuple with the RuntimeLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeLanguage

`func (o *ConnectV1CustomConnectorPluginUpdate) SetRuntimeLanguage(v string)`

SetRuntimeLanguage sets RuntimeLanguage field to given value.

### HasRuntimeLanguage

`func (o *ConnectV1CustomConnectorPluginUpdate) HasRuntimeLanguage() bool`

HasRuntimeLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


