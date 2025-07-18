# ConnectV1CustomConnectorRuntime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**CustomConnectPluginRuntimeName** | Pointer to **string** | Name of the runtime that is being used while provisioning a custom connector. This corresponds to the property custom.connect.plugin.runtime in the connector configuration.  | [optional] [readonly] 
**RuntimeAkVersion** | Pointer to **string** | The underlying version of Apache Kafka which bundles the connect runtime | [optional] [readonly] 
**SupportedJavaVersions** | Pointer to **[]string** | List of supported Java versions | [optional] [readonly] 
**ProductMaturity** | Pointer to **string** | The product maturity phase for the plugin runtime. EA (Early Access), GA (Generally Available), or Preview.  | [optional] [readonly] 
**EndOfLifeAt** | Pointer to **time.Time** | End of Life date for the runtime | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the runtime | [optional] [readonly] 

## Methods

### NewConnectV1CustomConnectorRuntime

`func NewConnectV1CustomConnectorRuntime() *ConnectV1CustomConnectorRuntime`

NewConnectV1CustomConnectorRuntime instantiates a new ConnectV1CustomConnectorRuntime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1CustomConnectorRuntimeWithDefaults

`func NewConnectV1CustomConnectorRuntimeWithDefaults() *ConnectV1CustomConnectorRuntime`

NewConnectV1CustomConnectorRuntimeWithDefaults instantiates a new ConnectV1CustomConnectorRuntime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ConnectV1CustomConnectorRuntime) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ConnectV1CustomConnectorRuntime) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ConnectV1CustomConnectorRuntime) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ConnectV1CustomConnectorRuntime) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ConnectV1CustomConnectorRuntime) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConnectV1CustomConnectorRuntime) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConnectV1CustomConnectorRuntime) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ConnectV1CustomConnectorRuntime) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *ConnectV1CustomConnectorRuntime) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectV1CustomConnectorRuntime) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectV1CustomConnectorRuntime) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectV1CustomConnectorRuntime) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCustomConnectPluginRuntimeName

`func (o *ConnectV1CustomConnectorRuntime) GetCustomConnectPluginRuntimeName() string`

GetCustomConnectPluginRuntimeName returns the CustomConnectPluginRuntimeName field if non-nil, zero value otherwise.

### GetCustomConnectPluginRuntimeNameOk

`func (o *ConnectV1CustomConnectorRuntime) GetCustomConnectPluginRuntimeNameOk() (*string, bool)`

GetCustomConnectPluginRuntimeNameOk returns a tuple with the CustomConnectPluginRuntimeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConnectPluginRuntimeName

`func (o *ConnectV1CustomConnectorRuntime) SetCustomConnectPluginRuntimeName(v string)`

SetCustomConnectPluginRuntimeName sets CustomConnectPluginRuntimeName field to given value.

### HasCustomConnectPluginRuntimeName

`func (o *ConnectV1CustomConnectorRuntime) HasCustomConnectPluginRuntimeName() bool`

HasCustomConnectPluginRuntimeName returns a boolean if a field has been set.

### GetRuntimeAkVersion

`func (o *ConnectV1CustomConnectorRuntime) GetRuntimeAkVersion() string`

GetRuntimeAkVersion returns the RuntimeAkVersion field if non-nil, zero value otherwise.

### GetRuntimeAkVersionOk

`func (o *ConnectV1CustomConnectorRuntime) GetRuntimeAkVersionOk() (*string, bool)`

GetRuntimeAkVersionOk returns a tuple with the RuntimeAkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeAkVersion

`func (o *ConnectV1CustomConnectorRuntime) SetRuntimeAkVersion(v string)`

SetRuntimeAkVersion sets RuntimeAkVersion field to given value.

### HasRuntimeAkVersion

`func (o *ConnectV1CustomConnectorRuntime) HasRuntimeAkVersion() bool`

HasRuntimeAkVersion returns a boolean if a field has been set.

### GetSupportedJavaVersions

`func (o *ConnectV1CustomConnectorRuntime) GetSupportedJavaVersions() []string`

GetSupportedJavaVersions returns the SupportedJavaVersions field if non-nil, zero value otherwise.

### GetSupportedJavaVersionsOk

`func (o *ConnectV1CustomConnectorRuntime) GetSupportedJavaVersionsOk() (*[]string, bool)`

GetSupportedJavaVersionsOk returns a tuple with the SupportedJavaVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedJavaVersions

`func (o *ConnectV1CustomConnectorRuntime) SetSupportedJavaVersions(v []string)`

SetSupportedJavaVersions sets SupportedJavaVersions field to given value.

### HasSupportedJavaVersions

`func (o *ConnectV1CustomConnectorRuntime) HasSupportedJavaVersions() bool`

HasSupportedJavaVersions returns a boolean if a field has been set.

### GetProductMaturity

`func (o *ConnectV1CustomConnectorRuntime) GetProductMaturity() string`

GetProductMaturity returns the ProductMaturity field if non-nil, zero value otherwise.

### GetProductMaturityOk

`func (o *ConnectV1CustomConnectorRuntime) GetProductMaturityOk() (*string, bool)`

GetProductMaturityOk returns a tuple with the ProductMaturity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductMaturity

`func (o *ConnectV1CustomConnectorRuntime) SetProductMaturity(v string)`

SetProductMaturity sets ProductMaturity field to given value.

### HasProductMaturity

`func (o *ConnectV1CustomConnectorRuntime) HasProductMaturity() bool`

HasProductMaturity returns a boolean if a field has been set.

### GetEndOfLifeAt

`func (o *ConnectV1CustomConnectorRuntime) GetEndOfLifeAt() time.Time`

GetEndOfLifeAt returns the EndOfLifeAt field if non-nil, zero value otherwise.

### GetEndOfLifeAtOk

`func (o *ConnectV1CustomConnectorRuntime) GetEndOfLifeAtOk() (*time.Time, bool)`

GetEndOfLifeAtOk returns a tuple with the EndOfLifeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfLifeAt

`func (o *ConnectV1CustomConnectorRuntime) SetEndOfLifeAt(v time.Time)`

SetEndOfLifeAt sets EndOfLifeAt field to given value.

### HasEndOfLifeAt

`func (o *ConnectV1CustomConnectorRuntime) HasEndOfLifeAt() bool`

HasEndOfLifeAt returns a boolean if a field has been set.

### GetDescription

`func (o *ConnectV1CustomConnectorRuntime) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectV1CustomConnectorRuntime) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectV1CustomConnectorRuntime) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectV1CustomConnectorRuntime) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


