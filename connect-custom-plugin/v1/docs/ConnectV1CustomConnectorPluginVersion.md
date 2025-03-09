# ConnectV1CustomConnectorPluginVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Version** | Pointer to **string** | Version of the Custom Connector Plugin. The version must comply with SemVer (e.g., &#x60;1.2.3&#x60;, &#x60;1.2.3-beta&#x60;, &#x60;1.2.3-rc.123&#x60;, &#x60;1.2.3-rc.123+build.456&#x60;).  | [optional] 
**SensitiveConfigProperties** | Pointer to **[]string** | A sensitive property is a connector configuration property that must be hidden after a user enters property value when setting up connector. | [optional] 
**UploadSource** | Pointer to [**ConnectV1CustomConnectorPluginVersionUploadSourceOneOf**](ConnectV1CustomConnectorPluginVersionUploadSourceOneOf.md) | Upload source of Custom Connector Plugin Version. Only required in &#x60;create&#x60; request, will be ignored in &#x60;read&#x60;, &#x60;update&#x60; or &#x60;list&#x60;. | [optional] 

## Methods

### NewConnectV1CustomConnectorPluginVersion

`func NewConnectV1CustomConnectorPluginVersion() *ConnectV1CustomConnectorPluginVersion`

NewConnectV1CustomConnectorPluginVersion instantiates a new ConnectV1CustomConnectorPluginVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1CustomConnectorPluginVersionWithDefaults

`func NewConnectV1CustomConnectorPluginVersionWithDefaults() *ConnectV1CustomConnectorPluginVersion`

NewConnectV1CustomConnectorPluginVersionWithDefaults instantiates a new ConnectV1CustomConnectorPluginVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ConnectV1CustomConnectorPluginVersion) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ConnectV1CustomConnectorPluginVersion) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ConnectV1CustomConnectorPluginVersion) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ConnectV1CustomConnectorPluginVersion) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ConnectV1CustomConnectorPluginVersion) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConnectV1CustomConnectorPluginVersion) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConnectV1CustomConnectorPluginVersion) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ConnectV1CustomConnectorPluginVersion) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *ConnectV1CustomConnectorPluginVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectV1CustomConnectorPluginVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectV1CustomConnectorPluginVersion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectV1CustomConnectorPluginVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *ConnectV1CustomConnectorPluginVersion) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConnectV1CustomConnectorPluginVersion) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConnectV1CustomConnectorPluginVersion) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConnectV1CustomConnectorPluginVersion) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetVersion

`func (o *ConnectV1CustomConnectorPluginVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConnectV1CustomConnectorPluginVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConnectV1CustomConnectorPluginVersion) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConnectV1CustomConnectorPluginVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSensitiveConfigProperties

`func (o *ConnectV1CustomConnectorPluginVersion) GetSensitiveConfigProperties() []string`

GetSensitiveConfigProperties returns the SensitiveConfigProperties field if non-nil, zero value otherwise.

### GetSensitiveConfigPropertiesOk

`func (o *ConnectV1CustomConnectorPluginVersion) GetSensitiveConfigPropertiesOk() (*[]string, bool)`

GetSensitiveConfigPropertiesOk returns a tuple with the SensitiveConfigProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveConfigProperties

`func (o *ConnectV1CustomConnectorPluginVersion) SetSensitiveConfigProperties(v []string)`

SetSensitiveConfigProperties sets SensitiveConfigProperties field to given value.

### HasSensitiveConfigProperties

`func (o *ConnectV1CustomConnectorPluginVersion) HasSensitiveConfigProperties() bool`

HasSensitiveConfigProperties returns a boolean if a field has been set.

### GetUploadSource

`func (o *ConnectV1CustomConnectorPluginVersion) GetUploadSource() ConnectV1CustomConnectorPluginVersionUploadSourceOneOf`

GetUploadSource returns the UploadSource field if non-nil, zero value otherwise.

### GetUploadSourceOk

`func (o *ConnectV1CustomConnectorPluginVersion) GetUploadSourceOk() (*ConnectV1CustomConnectorPluginVersionUploadSourceOneOf, bool)`

GetUploadSourceOk returns a tuple with the UploadSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadSource

`func (o *ConnectV1CustomConnectorPluginVersion) SetUploadSource(v ConnectV1CustomConnectorPluginVersionUploadSourceOneOf)`

SetUploadSource sets UploadSource field to given value.

### HasUploadSource

`func (o *ConnectV1CustomConnectorPluginVersion) HasUploadSource() bool`

HasUploadSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


