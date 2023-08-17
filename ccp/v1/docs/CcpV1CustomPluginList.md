# CcpV1CustomPluginList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]CcpV1CustomPlugin**](CcpV1CustomPlugin.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewCcpV1CustomPluginList

`func NewCcpV1CustomPluginList(apiVersion string, kind string, metadata ListMeta, data []CcpV1CustomPlugin, ) *CcpV1CustomPluginList`

NewCcpV1CustomPluginList instantiates a new CcpV1CustomPluginList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCcpV1CustomPluginListWithDefaults

`func NewCcpV1CustomPluginListWithDefaults() *CcpV1CustomPluginList`

NewCcpV1CustomPluginListWithDefaults instantiates a new CcpV1CustomPluginList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CcpV1CustomPluginList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CcpV1CustomPluginList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CcpV1CustomPluginList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *CcpV1CustomPluginList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CcpV1CustomPluginList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CcpV1CustomPluginList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *CcpV1CustomPluginList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CcpV1CustomPluginList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CcpV1CustomPluginList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *CcpV1CustomPluginList) GetData() []CcpV1CustomPlugin`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CcpV1CustomPluginList) GetDataOk() (*[]CcpV1CustomPlugin, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CcpV1CustomPluginList) SetData(v []CcpV1CustomPlugin)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


