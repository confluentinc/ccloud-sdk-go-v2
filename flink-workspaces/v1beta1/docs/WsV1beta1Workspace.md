# WsV1beta1Workspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Org** | Pointer to **string** | The organization in which the workspace exists. | [optional] [readonly] 
**Environment** | Pointer to **string** | The environment in which the workspace exists. | [optional] [readonly] 
**Name** | Pointer to **string** | The workspace name that is unique across the environment and region. | [optional] 
**Spec** | Pointer to [**WsV1beta1WorkspaceSpec**](WsV1beta1WorkspaceSpec.md) |  | [optional] 

## Methods

### NewWsV1beta1Workspace

`func NewWsV1beta1Workspace() *WsV1beta1Workspace`

NewWsV1beta1Workspace instantiates a new WsV1beta1Workspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsV1beta1WorkspaceWithDefaults

`func NewWsV1beta1WorkspaceWithDefaults() *WsV1beta1Workspace`

NewWsV1beta1WorkspaceWithDefaults instantiates a new WsV1beta1Workspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *WsV1beta1Workspace) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *WsV1beta1Workspace) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *WsV1beta1Workspace) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *WsV1beta1Workspace) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *WsV1beta1Workspace) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WsV1beta1Workspace) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WsV1beta1Workspace) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *WsV1beta1Workspace) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *WsV1beta1Workspace) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WsV1beta1Workspace) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WsV1beta1Workspace) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WsV1beta1Workspace) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOrg

`func (o *WsV1beta1Workspace) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *WsV1beta1Workspace) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *WsV1beta1Workspace) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *WsV1beta1Workspace) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetEnvironment

`func (o *WsV1beta1Workspace) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *WsV1beta1Workspace) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *WsV1beta1Workspace) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *WsV1beta1Workspace) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetName

`func (o *WsV1beta1Workspace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WsV1beta1Workspace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WsV1beta1Workspace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WsV1beta1Workspace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpec

`func (o *WsV1beta1Workspace) GetSpec() WsV1beta1WorkspaceSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *WsV1beta1Workspace) GetSpecOk() (*WsV1beta1WorkspaceSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *WsV1beta1Workspace) SetSpec(v WsV1beta1WorkspaceSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *WsV1beta1Workspace) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


