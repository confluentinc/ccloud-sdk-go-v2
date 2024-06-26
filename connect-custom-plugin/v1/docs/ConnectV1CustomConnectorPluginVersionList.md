# ConnectV1CustomConnectorPluginVersionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]ConnectV1CustomConnectorPluginVersion**](ConnectV1CustomConnectorPluginVersion.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewConnectV1CustomConnectorPluginVersionList

`func NewConnectV1CustomConnectorPluginVersionList(apiVersion string, kind string, metadata ListMeta, data []ConnectV1CustomConnectorPluginVersion, ) *ConnectV1CustomConnectorPluginVersionList`

NewConnectV1CustomConnectorPluginVersionList instantiates a new ConnectV1CustomConnectorPluginVersionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1CustomConnectorPluginVersionListWithDefaults

`func NewConnectV1CustomConnectorPluginVersionListWithDefaults() *ConnectV1CustomConnectorPluginVersionList`

NewConnectV1CustomConnectorPluginVersionListWithDefaults instantiates a new ConnectV1CustomConnectorPluginVersionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ConnectV1CustomConnectorPluginVersionList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ConnectV1CustomConnectorPluginVersionList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ConnectV1CustomConnectorPluginVersionList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ConnectV1CustomConnectorPluginVersionList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConnectV1CustomConnectorPluginVersionList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConnectV1CustomConnectorPluginVersionList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ConnectV1CustomConnectorPluginVersionList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConnectV1CustomConnectorPluginVersionList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConnectV1CustomConnectorPluginVersionList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ConnectV1CustomConnectorPluginVersionList) GetData() []ConnectV1CustomConnectorPluginVersion`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConnectV1CustomConnectorPluginVersionList) GetDataOk() (*[]ConnectV1CustomConnectorPluginVersion, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConnectV1CustomConnectorPluginVersionList) SetData(v []ConnectV1CustomConnectorPluginVersion)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


