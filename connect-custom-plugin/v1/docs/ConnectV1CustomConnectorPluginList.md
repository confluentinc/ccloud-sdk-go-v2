# ConnectV1CustomConnectorPluginList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]ConnectV1CustomConnectorPlugin**](ConnectV1CustomConnectorPlugin.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewConnectV1CustomConnectorPluginList

`func NewConnectV1CustomConnectorPluginList(apiVersion string, kind string, metadata ListMeta, data []ConnectV1CustomConnectorPlugin, ) *ConnectV1CustomConnectorPluginList`

NewConnectV1CustomConnectorPluginList instantiates a new ConnectV1CustomConnectorPluginList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1CustomConnectorPluginListWithDefaults

`func NewConnectV1CustomConnectorPluginListWithDefaults() *ConnectV1CustomConnectorPluginList`

NewConnectV1CustomConnectorPluginListWithDefaults instantiates a new ConnectV1CustomConnectorPluginList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ConnectV1CustomConnectorPluginList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ConnectV1CustomConnectorPluginList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ConnectV1CustomConnectorPluginList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ConnectV1CustomConnectorPluginList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConnectV1CustomConnectorPluginList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConnectV1CustomConnectorPluginList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ConnectV1CustomConnectorPluginList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConnectV1CustomConnectorPluginList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConnectV1CustomConnectorPluginList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ConnectV1CustomConnectorPluginList) GetData() []ConnectV1CustomConnectorPlugin`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConnectV1CustomConnectorPluginList) GetDataOk() (*[]ConnectV1CustomConnectorPlugin, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConnectV1CustomConnectorPluginList) SetData(v []ConnectV1CustomConnectorPlugin)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


