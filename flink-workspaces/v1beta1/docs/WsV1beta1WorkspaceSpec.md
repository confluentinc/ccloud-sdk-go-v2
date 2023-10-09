# WsV1beta1WorkspaceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | A human-readable display name for the workspace. | [optional] 
**ComputePool** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) |  | [optional] 
**ServiceAccount** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) |  | [optional] 
**Blocks** | Pointer to [**[]WsV1beta1Block**](WsV1beta1Block.md) | The ordered blocks for the new workspace | [optional] 
**Properties** | Pointer to **map[string]string** | A map (key-value pairs) of workspace-level properties. All key-value pairs are optional, and clients choose whether and how to use these key-value pairs. | [optional] 

## Methods

### NewWsV1beta1WorkspaceSpec

`func NewWsV1beta1WorkspaceSpec() *WsV1beta1WorkspaceSpec`

NewWsV1beta1WorkspaceSpec instantiates a new WsV1beta1WorkspaceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsV1beta1WorkspaceSpecWithDefaults

`func NewWsV1beta1WorkspaceSpecWithDefaults() *WsV1beta1WorkspaceSpec`

NewWsV1beta1WorkspaceSpecWithDefaults instantiates a new WsV1beta1WorkspaceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *WsV1beta1WorkspaceSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WsV1beta1WorkspaceSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WsV1beta1WorkspaceSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *WsV1beta1WorkspaceSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetComputePool

`func (o *WsV1beta1WorkspaceSpec) GetComputePool() EnvScopedObjectReference`

GetComputePool returns the ComputePool field if non-nil, zero value otherwise.

### GetComputePoolOk

`func (o *WsV1beta1WorkspaceSpec) GetComputePoolOk() (*EnvScopedObjectReference, bool)`

GetComputePoolOk returns a tuple with the ComputePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePool

`func (o *WsV1beta1WorkspaceSpec) SetComputePool(v EnvScopedObjectReference)`

SetComputePool sets ComputePool field to given value.

### HasComputePool

`func (o *WsV1beta1WorkspaceSpec) HasComputePool() bool`

HasComputePool returns a boolean if a field has been set.

### GetServiceAccount

`func (o *WsV1beta1WorkspaceSpec) GetServiceAccount() GlobalObjectReference`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *WsV1beta1WorkspaceSpec) GetServiceAccountOk() (*GlobalObjectReference, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *WsV1beta1WorkspaceSpec) SetServiceAccount(v GlobalObjectReference)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *WsV1beta1WorkspaceSpec) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetBlocks

`func (o *WsV1beta1WorkspaceSpec) GetBlocks() []WsV1beta1Block`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *WsV1beta1WorkspaceSpec) GetBlocksOk() (*[]WsV1beta1Block, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *WsV1beta1WorkspaceSpec) SetBlocks(v []WsV1beta1Block)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *WsV1beta1WorkspaceSpec) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetProperties

`func (o *WsV1beta1WorkspaceSpec) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WsV1beta1WorkspaceSpec) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WsV1beta1WorkspaceSpec) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WsV1beta1WorkspaceSpec) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


