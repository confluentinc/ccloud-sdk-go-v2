# ConnectV1CustomConnectorPluginVersionUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**SensitiveConfigProperties** | Pointer to **[]string** | A sensitive property is a connector configuration property that must be hidden after a user enters property value when setting up connector. | [optional] 

## Methods

### NewConnectV1CustomConnectorPluginVersionUpdate

`func NewConnectV1CustomConnectorPluginVersionUpdate() *ConnectV1CustomConnectorPluginVersionUpdate`

NewConnectV1CustomConnectorPluginVersionUpdate instantiates a new ConnectV1CustomConnectorPluginVersionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1CustomConnectorPluginVersionUpdateWithDefaults

`func NewConnectV1CustomConnectorPluginVersionUpdateWithDefaults() *ConnectV1CustomConnectorPluginVersionUpdate`

NewConnectV1CustomConnectorPluginVersionUpdateWithDefaults instantiates a new ConnectV1CustomConnectorPluginVersionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ConnectV1CustomConnectorPluginVersionUpdate) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ConnectV1CustomConnectorPluginVersionUpdate) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ConnectV1CustomConnectorPluginVersionUpdate) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ConnectV1CustomConnectorPluginVersionUpdate) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ConnectV1CustomConnectorPluginVersionUpdate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConnectV1CustomConnectorPluginVersionUpdate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConnectV1CustomConnectorPluginVersionUpdate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ConnectV1CustomConnectorPluginVersionUpdate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *ConnectV1CustomConnectorPluginVersionUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectV1CustomConnectorPluginVersionUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectV1CustomConnectorPluginVersionUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectV1CustomConnectorPluginVersionUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *ConnectV1CustomConnectorPluginVersionUpdate) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConnectV1CustomConnectorPluginVersionUpdate) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConnectV1CustomConnectorPluginVersionUpdate) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConnectV1CustomConnectorPluginVersionUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSensitiveConfigProperties

`func (o *ConnectV1CustomConnectorPluginVersionUpdate) GetSensitiveConfigProperties() []string`

GetSensitiveConfigProperties returns the SensitiveConfigProperties field if non-nil, zero value otherwise.

### GetSensitiveConfigPropertiesOk

`func (o *ConnectV1CustomConnectorPluginVersionUpdate) GetSensitiveConfigPropertiesOk() (*[]string, bool)`

GetSensitiveConfigPropertiesOk returns a tuple with the SensitiveConfigProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveConfigProperties

`func (o *ConnectV1CustomConnectorPluginVersionUpdate) SetSensitiveConfigProperties(v []string)`

SetSensitiveConfigProperties sets SensitiveConfigProperties field to given value.

### HasSensitiveConfigProperties

`func (o *ConnectV1CustomConnectorPluginVersionUpdate) HasSensitiveConfigProperties() bool`

HasSensitiveConfigProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


