# SqlV1Statement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Metadata** | Pointer to [**StatementObjectMeta**](StatementObjectMeta.md) |  | [optional] 
**Name** | Pointer to **string** | The user provided name of the resource, unique within this environment. | [optional] 
**OrganizationId** | Pointer to **string** | The unique identifier for the organization. | [optional] 
**EnvironmentId** | Pointer to **string** | The unique identifier for the environment. | [optional] 
**Spec** | Pointer to [**SqlV1StatementSpec**](SqlV1StatementSpec.md) |  | [optional] 
**Status** | Pointer to [**SqlV1StatementStatus**](SqlV1StatementStatus.md) |  | [optional] 
**Result** | Pointer to [**SqlV1StatementResult**](SqlV1StatementResult.md) |  | [optional] 

## Methods

### NewSqlV1Statement

`func NewSqlV1Statement() *SqlV1Statement`

NewSqlV1Statement instantiates a new SqlV1Statement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1StatementWithDefaults

`func NewSqlV1StatementWithDefaults() *SqlV1Statement`

NewSqlV1StatementWithDefaults instantiates a new SqlV1Statement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SqlV1Statement) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SqlV1Statement) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SqlV1Statement) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SqlV1Statement) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *SqlV1Statement) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SqlV1Statement) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SqlV1Statement) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SqlV1Statement) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *SqlV1Statement) GetMetadata() StatementObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SqlV1Statement) GetMetadataOk() (*StatementObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SqlV1Statement) SetMetadata(v StatementObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SqlV1Statement) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *SqlV1Statement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlV1Statement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlV1Statement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SqlV1Statement) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *SqlV1Statement) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SqlV1Statement) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SqlV1Statement) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *SqlV1Statement) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *SqlV1Statement) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *SqlV1Statement) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *SqlV1Statement) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *SqlV1Statement) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetSpec

`func (o *SqlV1Statement) GetSpec() SqlV1StatementSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SqlV1Statement) GetSpecOk() (*SqlV1StatementSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SqlV1Statement) SetSpec(v SqlV1StatementSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *SqlV1Statement) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *SqlV1Statement) GetStatus() SqlV1StatementStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SqlV1Statement) GetStatusOk() (*SqlV1StatementStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SqlV1Statement) SetStatus(v SqlV1StatementStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SqlV1Statement) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *SqlV1Statement) GetResult() SqlV1StatementResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SqlV1Statement) GetResultOk() (*SqlV1StatementResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SqlV1Statement) SetResult(v SqlV1StatementResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *SqlV1Statement) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


