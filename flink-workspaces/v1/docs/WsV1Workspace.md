# WsV1Workspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**OrganizationId** | Pointer to **string** | The organization in which the workspace exists. | [optional] [readonly] 
**EnvironmentId** | Pointer to **string** | The environment in which the workspace exists. | [optional] [readonly] 
**Name** | Pointer to **string** | The workspace name that is unique across the environment and region. | [optional] 
**Spec** | Pointer to [**WsV1WorkspaceSpec**](WsV1WorkspaceSpec.md) |  | [optional] 
**Status** | Pointer to [**WsV1WorkspaceStatus**](WsV1WorkspaceStatus.md) |  | [optional] 

## Methods

### NewWsV1Workspace

`func NewWsV1Workspace() *WsV1Workspace`

NewWsV1Workspace instantiates a new WsV1Workspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsV1WorkspaceWithDefaults

`func NewWsV1WorkspaceWithDefaults() *WsV1Workspace`

NewWsV1WorkspaceWithDefaults instantiates a new WsV1Workspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *WsV1Workspace) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *WsV1Workspace) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *WsV1Workspace) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *WsV1Workspace) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *WsV1Workspace) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WsV1Workspace) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WsV1Workspace) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *WsV1Workspace) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *WsV1Workspace) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WsV1Workspace) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WsV1Workspace) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WsV1Workspace) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOrganizationId

`func (o *WsV1Workspace) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *WsV1Workspace) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *WsV1Workspace) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *WsV1Workspace) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *WsV1Workspace) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *WsV1Workspace) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *WsV1Workspace) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *WsV1Workspace) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetName

`func (o *WsV1Workspace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WsV1Workspace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WsV1Workspace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WsV1Workspace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpec

`func (o *WsV1Workspace) GetSpec() WsV1WorkspaceSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *WsV1Workspace) GetSpecOk() (*WsV1WorkspaceSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *WsV1Workspace) SetSpec(v WsV1WorkspaceSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *WsV1Workspace) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *WsV1Workspace) GetStatus() WsV1WorkspaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WsV1Workspace) GetStatusOk() (*WsV1WorkspaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WsV1Workspace) SetStatus(v WsV1WorkspaceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WsV1Workspace) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


