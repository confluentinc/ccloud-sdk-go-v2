# CcpmV1CustomConnectPluginList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]CcpmV1CustomConnectPlugin**](CcpmV1CustomConnectPlugin.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewCcpmV1CustomConnectPluginList

`func NewCcpmV1CustomConnectPluginList(apiVersion string, kind string, metadata ListMeta, data []CcpmV1CustomConnectPlugin, ) *CcpmV1CustomConnectPluginList`

NewCcpmV1CustomConnectPluginList instantiates a new CcpmV1CustomConnectPluginList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCcpmV1CustomConnectPluginListWithDefaults

`func NewCcpmV1CustomConnectPluginListWithDefaults() *CcpmV1CustomConnectPluginList`

NewCcpmV1CustomConnectPluginListWithDefaults instantiates a new CcpmV1CustomConnectPluginList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CcpmV1CustomConnectPluginList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CcpmV1CustomConnectPluginList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CcpmV1CustomConnectPluginList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *CcpmV1CustomConnectPluginList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CcpmV1CustomConnectPluginList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CcpmV1CustomConnectPluginList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *CcpmV1CustomConnectPluginList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CcpmV1CustomConnectPluginList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CcpmV1CustomConnectPluginList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *CcpmV1CustomConnectPluginList) GetData() []CcpmV1CustomConnectPlugin`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CcpmV1CustomConnectPluginList) GetDataOk() (*[]CcpmV1CustomConnectPlugin, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CcpmV1CustomConnectPluginList) SetData(v []CcpmV1CustomConnectPlugin)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


