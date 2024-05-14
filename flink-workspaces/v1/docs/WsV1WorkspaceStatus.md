# WsV1WorkspaceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkKind** | Pointer to **string** | The networking type used by the Workspace:  PUBLIC: Workspace is using public networking;  PRIVATE: Workspace is using private networking;  | [optional] [readonly] 

## Methods

### NewWsV1WorkspaceStatus

`func NewWsV1WorkspaceStatus() *WsV1WorkspaceStatus`

NewWsV1WorkspaceStatus instantiates a new WsV1WorkspaceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsV1WorkspaceStatusWithDefaults

`func NewWsV1WorkspaceStatusWithDefaults() *WsV1WorkspaceStatus`

NewWsV1WorkspaceStatusWithDefaults instantiates a new WsV1WorkspaceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkKind

`func (o *WsV1WorkspaceStatus) GetNetworkKind() string`

GetNetworkKind returns the NetworkKind field if non-nil, zero value otherwise.

### GetNetworkKindOk

`func (o *WsV1WorkspaceStatus) GetNetworkKindOk() (*string, bool)`

GetNetworkKindOk returns a tuple with the NetworkKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkKind

`func (o *WsV1WorkspaceStatus) SetNetworkKind(v string)`

SetNetworkKind sets NetworkKind field to given value.

### HasNetworkKind

`func (o *WsV1WorkspaceStatus) HasNetworkKind() bool`

HasNetworkKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


