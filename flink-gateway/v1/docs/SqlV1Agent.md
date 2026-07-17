# SqlV1Agent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ObjectMeta**](ObjectMeta.md) |  | 
**Name** | **string** | The user-provided name of the agent, unique within this environment. | 
**OrganizationId** | **string** | The unique identifier for the organization. | [readonly] 
**EnvironmentId** | **string** | The unique identifier for the environment. | [readonly] 
**Spec** | [**SqlV1AgentSpec**](SqlV1AgentSpec.md) |  | 
**Status** | Pointer to [**SqlV1AgentStatus**](SqlV1AgentStatus.md) |  | [optional] [readonly] 

## Methods

### NewSqlV1Agent

`func NewSqlV1Agent(apiVersion string, kind string, metadata ObjectMeta, name string, organizationId string, environmentId string, spec SqlV1AgentSpec, ) *SqlV1Agent`

NewSqlV1Agent instantiates a new SqlV1Agent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1AgentWithDefaults

`func NewSqlV1AgentWithDefaults() *SqlV1Agent`

NewSqlV1AgentWithDefaults instantiates a new SqlV1Agent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SqlV1Agent) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SqlV1Agent) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SqlV1Agent) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *SqlV1Agent) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SqlV1Agent) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SqlV1Agent) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *SqlV1Agent) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SqlV1Agent) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SqlV1Agent) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.


### GetName

`func (o *SqlV1Agent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlV1Agent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlV1Agent) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *SqlV1Agent) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SqlV1Agent) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SqlV1Agent) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetEnvironmentId

`func (o *SqlV1Agent) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *SqlV1Agent) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *SqlV1Agent) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetSpec

`func (o *SqlV1Agent) GetSpec() SqlV1AgentSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SqlV1Agent) GetSpecOk() (*SqlV1AgentSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SqlV1Agent) SetSpec(v SqlV1AgentSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *SqlV1Agent) GetStatus() SqlV1AgentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SqlV1Agent) GetStatusOk() (*SqlV1AgentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SqlV1Agent) SetStatus(v SqlV1AgentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SqlV1Agent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


