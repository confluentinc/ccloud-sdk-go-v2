# WsV1WorkspaceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | 
**Kind** | **string** | Kind defines the object this REST resource represents. | 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]WsV1Workspace**](WsV1Workspace.md) | The array of workspace resource. | 

## Methods

### NewWsV1WorkspaceList

`func NewWsV1WorkspaceList(apiVersion string, kind string, metadata ListMeta, data []WsV1Workspace, ) *WsV1WorkspaceList`

NewWsV1WorkspaceList instantiates a new WsV1WorkspaceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsV1WorkspaceListWithDefaults

`func NewWsV1WorkspaceListWithDefaults() *WsV1WorkspaceList`

NewWsV1WorkspaceListWithDefaults instantiates a new WsV1WorkspaceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *WsV1WorkspaceList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *WsV1WorkspaceList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *WsV1WorkspaceList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *WsV1WorkspaceList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WsV1WorkspaceList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WsV1WorkspaceList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *WsV1WorkspaceList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WsV1WorkspaceList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WsV1WorkspaceList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *WsV1WorkspaceList) GetData() []WsV1Workspace`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WsV1WorkspaceList) GetDataOk() (*[]WsV1Workspace, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WsV1WorkspaceList) SetData(v []WsV1Workspace)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


