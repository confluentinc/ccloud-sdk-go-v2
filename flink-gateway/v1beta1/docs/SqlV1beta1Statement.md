# SqlV1beta1Statement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Name** | Pointer to **string** | The user provided name of the resource, unique within this environment. | [optional] 
**OrganizationId** | Pointer to **string** | The unique identifier for the organization. | [optional] 
**EnvironmentId** | Pointer to **string** | The unique identifier for the environment. | [optional] 
**Principal** | Pointer to **string** | The principal that the query runs as - can be a user account or service account | [optional] 
**Spec** | Pointer to [**SqlV1beta1StatementSpec**](SqlV1beta1StatementSpec.md) |  | [optional] 
**Status** | Pointer to [**SqlV1beta1StatementStatus**](SqlV1beta1StatementStatus.md) |  | [optional] 

## Methods

### NewSqlV1beta1Statement

`func NewSqlV1beta1Statement() *SqlV1beta1Statement`

NewSqlV1beta1Statement instantiates a new SqlV1beta1Statement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1beta1StatementWithDefaults

`func NewSqlV1beta1StatementWithDefaults() *SqlV1beta1Statement`

NewSqlV1beta1StatementWithDefaults instantiates a new SqlV1beta1Statement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SqlV1beta1Statement) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SqlV1beta1Statement) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SqlV1beta1Statement) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SqlV1beta1Statement) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *SqlV1beta1Statement) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SqlV1beta1Statement) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SqlV1beta1Statement) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SqlV1beta1Statement) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *SqlV1beta1Statement) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SqlV1beta1Statement) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SqlV1beta1Statement) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SqlV1beta1Statement) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *SqlV1beta1Statement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlV1beta1Statement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlV1beta1Statement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SqlV1beta1Statement) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *SqlV1beta1Statement) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SqlV1beta1Statement) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SqlV1beta1Statement) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *SqlV1beta1Statement) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *SqlV1beta1Statement) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *SqlV1beta1Statement) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *SqlV1beta1Statement) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *SqlV1beta1Statement) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetPrincipal

`func (o *SqlV1beta1Statement) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *SqlV1beta1Statement) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *SqlV1beta1Statement) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *SqlV1beta1Statement) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetSpec

`func (o *SqlV1beta1Statement) GetSpec() SqlV1beta1StatementSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SqlV1beta1Statement) GetSpecOk() (*SqlV1beta1StatementSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SqlV1beta1Statement) SetSpec(v SqlV1beta1StatementSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *SqlV1beta1Statement) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *SqlV1beta1Statement) GetStatus() SqlV1beta1StatementStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SqlV1beta1Statement) GetStatusOk() (*SqlV1beta1StatementStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SqlV1beta1Statement) SetStatus(v SqlV1beta1StatementStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SqlV1beta1Statement) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


