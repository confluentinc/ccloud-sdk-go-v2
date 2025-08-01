# ConnectV1CustomConnectorRuntimeList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]ConnectV1CustomConnectorRuntime**](ConnectV1CustomConnectorRuntime.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewConnectV1CustomConnectorRuntimeList

`func NewConnectV1CustomConnectorRuntimeList(apiVersion string, kind string, metadata ListMeta, data []ConnectV1CustomConnectorRuntime, ) *ConnectV1CustomConnectorRuntimeList`

NewConnectV1CustomConnectorRuntimeList instantiates a new ConnectV1CustomConnectorRuntimeList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1CustomConnectorRuntimeListWithDefaults

`func NewConnectV1CustomConnectorRuntimeListWithDefaults() *ConnectV1CustomConnectorRuntimeList`

NewConnectV1CustomConnectorRuntimeListWithDefaults instantiates a new ConnectV1CustomConnectorRuntimeList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ConnectV1CustomConnectorRuntimeList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ConnectV1CustomConnectorRuntimeList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ConnectV1CustomConnectorRuntimeList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ConnectV1CustomConnectorRuntimeList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConnectV1CustomConnectorRuntimeList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConnectV1CustomConnectorRuntimeList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ConnectV1CustomConnectorRuntimeList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConnectV1CustomConnectorRuntimeList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConnectV1CustomConnectorRuntimeList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ConnectV1CustomConnectorRuntimeList) GetData() []ConnectV1CustomConnectorRuntime`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConnectV1CustomConnectorRuntimeList) GetDataOk() (*[]ConnectV1CustomConnectorRuntime, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConnectV1CustomConnectorRuntimeList) SetData(v []ConnectV1CustomConnectorRuntime)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


