# CcpmV1CustomConnectPluginVersionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | Version of the Custom Connect Plugin. The version must comply with SemVer (e.g., &#x60;1.2.3&#x60;, &#x60;1.2.3-beta&#x60;, &#x60;1.2.3-rc.123&#x60;, &#x60;1.2.3-rc.123+build.456&#x60;).  | [optional] 
**SensitiveConfigProperties** | Pointer to **[]string** | A sensitive property is a connector configuration property that must be hidden after a user enters property value when setting up connector.  | [optional] 
**DocumentationLink** | Pointer to **string** | Document link of Custom Connect Plugin. | [optional] 
**ContentFormat** | Pointer to **string** | Archive format of Custom Connect Plugin. | [optional] [readonly] 
**ConnectorClasses** | Pointer to [**[]CcpmV1ConnectorClass**](CcpmV1ConnectorClass.md) | List of connector classes. The connector class must be a valid Java class name or alias for the connector. You can get the connector class from the connector documentation provided by the developer.  | [optional] 
**UploadSource** | Pointer to [**CcpmV1CustomConnectPluginVersionSpecUploadSourceOneOf**](CcpmV1CustomConnectPluginVersionSpecUploadSourceOneOf.md) | Upload source of Custom Connect Plugin Version. Only required in &#x60;create&#x60; request, will be ignored in &#x60;read&#x60;, &#x60;update&#x60; or &#x60;list&#x60;.  | [optional] 
**Environment** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewCcpmV1CustomConnectPluginVersionSpec

`func NewCcpmV1CustomConnectPluginVersionSpec() *CcpmV1CustomConnectPluginVersionSpec`

NewCcpmV1CustomConnectPluginVersionSpec instantiates a new CcpmV1CustomConnectPluginVersionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCcpmV1CustomConnectPluginVersionSpecWithDefaults

`func NewCcpmV1CustomConnectPluginVersionSpecWithDefaults() *CcpmV1CustomConnectPluginVersionSpec`

NewCcpmV1CustomConnectPluginVersionSpecWithDefaults instantiates a new CcpmV1CustomConnectPluginVersionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *CcpmV1CustomConnectPluginVersionSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CcpmV1CustomConnectPluginVersionSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CcpmV1CustomConnectPluginVersionSpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CcpmV1CustomConnectPluginVersionSpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSensitiveConfigProperties

`func (o *CcpmV1CustomConnectPluginVersionSpec) GetSensitiveConfigProperties() []string`

GetSensitiveConfigProperties returns the SensitiveConfigProperties field if non-nil, zero value otherwise.

### GetSensitiveConfigPropertiesOk

`func (o *CcpmV1CustomConnectPluginVersionSpec) GetSensitiveConfigPropertiesOk() (*[]string, bool)`

GetSensitiveConfigPropertiesOk returns a tuple with the SensitiveConfigProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveConfigProperties

`func (o *CcpmV1CustomConnectPluginVersionSpec) SetSensitiveConfigProperties(v []string)`

SetSensitiveConfigProperties sets SensitiveConfigProperties field to given value.

### HasSensitiveConfigProperties

`func (o *CcpmV1CustomConnectPluginVersionSpec) HasSensitiveConfigProperties() bool`

HasSensitiveConfigProperties returns a boolean if a field has been set.

### GetDocumentationLink

`func (o *CcpmV1CustomConnectPluginVersionSpec) GetDocumentationLink() string`

GetDocumentationLink returns the DocumentationLink field if non-nil, zero value otherwise.

### GetDocumentationLinkOk

`func (o *CcpmV1CustomConnectPluginVersionSpec) GetDocumentationLinkOk() (*string, bool)`

GetDocumentationLinkOk returns a tuple with the DocumentationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationLink

`func (o *CcpmV1CustomConnectPluginVersionSpec) SetDocumentationLink(v string)`

SetDocumentationLink sets DocumentationLink field to given value.

### HasDocumentationLink

`func (o *CcpmV1CustomConnectPluginVersionSpec) HasDocumentationLink() bool`

HasDocumentationLink returns a boolean if a field has been set.

### GetContentFormat

`func (o *CcpmV1CustomConnectPluginVersionSpec) GetContentFormat() string`

GetContentFormat returns the ContentFormat field if non-nil, zero value otherwise.

### GetContentFormatOk

`func (o *CcpmV1CustomConnectPluginVersionSpec) GetContentFormatOk() (*string, bool)`

GetContentFormatOk returns a tuple with the ContentFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFormat

`func (o *CcpmV1CustomConnectPluginVersionSpec) SetContentFormat(v string)`

SetContentFormat sets ContentFormat field to given value.

### HasContentFormat

`func (o *CcpmV1CustomConnectPluginVersionSpec) HasContentFormat() bool`

HasContentFormat returns a boolean if a field has been set.

### GetConnectorClasses

`func (o *CcpmV1CustomConnectPluginVersionSpec) GetConnectorClasses() []CcpmV1ConnectorClass`

GetConnectorClasses returns the ConnectorClasses field if non-nil, zero value otherwise.

### GetConnectorClassesOk

`func (o *CcpmV1CustomConnectPluginVersionSpec) GetConnectorClassesOk() (*[]CcpmV1ConnectorClass, bool)`

GetConnectorClassesOk returns a tuple with the ConnectorClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorClasses

`func (o *CcpmV1CustomConnectPluginVersionSpec) SetConnectorClasses(v []CcpmV1ConnectorClass)`

SetConnectorClasses sets ConnectorClasses field to given value.

### HasConnectorClasses

`func (o *CcpmV1CustomConnectPluginVersionSpec) HasConnectorClasses() bool`

HasConnectorClasses returns a boolean if a field has been set.

### GetUploadSource

`func (o *CcpmV1CustomConnectPluginVersionSpec) GetUploadSource() CcpmV1CustomConnectPluginVersionSpecUploadSourceOneOf`

GetUploadSource returns the UploadSource field if non-nil, zero value otherwise.

### GetUploadSourceOk

`func (o *CcpmV1CustomConnectPluginVersionSpec) GetUploadSourceOk() (*CcpmV1CustomConnectPluginVersionSpecUploadSourceOneOf, bool)`

GetUploadSourceOk returns a tuple with the UploadSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadSource

`func (o *CcpmV1CustomConnectPluginVersionSpec) SetUploadSource(v CcpmV1CustomConnectPluginVersionSpecUploadSourceOneOf)`

SetUploadSource sets UploadSource field to given value.

### HasUploadSource

`func (o *CcpmV1CustomConnectPluginVersionSpec) HasUploadSource() bool`

HasUploadSource returns a boolean if a field has been set.

### GetEnvironment

`func (o *CcpmV1CustomConnectPluginVersionSpec) GetEnvironment() EnvScopedObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CcpmV1CustomConnectPluginVersionSpec) GetEnvironmentOk() (*EnvScopedObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CcpmV1CustomConnectPluginVersionSpec) SetEnvironment(v EnvScopedObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CcpmV1CustomConnectPluginVersionSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


