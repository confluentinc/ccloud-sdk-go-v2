# SqlV1MaterializedTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ObjectMeta**](ObjectMeta.md) |  | 
**Name** | **string** | The user-provided name of the resource, unique within the Kafka cluster. May contain ASCII alphanumerics, &#39;.&#39;, &#39;_&#39; and &#39;-&#39;; must not be &#39;.&#39; or &#39;..&#39;; max length 249. | 
**OrganizationId** | **string** | The unique identifier for the organization. | [readonly] 
**EnvironmentId** | **string** | The unique identifier for the environment. | [readonly] 
**Spec** | [**SqlV1MaterializedTableSpec**](SqlV1MaterializedTableSpec.md) |  | 
**Status** | Pointer to [**SqlV1MaterializedTableStatus**](SqlV1MaterializedTableStatus.md) |  | [optional] [readonly] 

## Methods

### NewSqlV1MaterializedTable

`func NewSqlV1MaterializedTable(apiVersion string, kind string, metadata ObjectMeta, name string, organizationId string, environmentId string, spec SqlV1MaterializedTableSpec, ) *SqlV1MaterializedTable`

NewSqlV1MaterializedTable instantiates a new SqlV1MaterializedTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1MaterializedTableWithDefaults

`func NewSqlV1MaterializedTableWithDefaults() *SqlV1MaterializedTable`

NewSqlV1MaterializedTableWithDefaults instantiates a new SqlV1MaterializedTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SqlV1MaterializedTable) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SqlV1MaterializedTable) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SqlV1MaterializedTable) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *SqlV1MaterializedTable) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SqlV1MaterializedTable) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SqlV1MaterializedTable) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *SqlV1MaterializedTable) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SqlV1MaterializedTable) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SqlV1MaterializedTable) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.


### GetName

`func (o *SqlV1MaterializedTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlV1MaterializedTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlV1MaterializedTable) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *SqlV1MaterializedTable) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SqlV1MaterializedTable) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SqlV1MaterializedTable) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetEnvironmentId

`func (o *SqlV1MaterializedTable) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *SqlV1MaterializedTable) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *SqlV1MaterializedTable) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetSpec

`func (o *SqlV1MaterializedTable) GetSpec() SqlV1MaterializedTableSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SqlV1MaterializedTable) GetSpecOk() (*SqlV1MaterializedTableSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SqlV1MaterializedTable) SetSpec(v SqlV1MaterializedTableSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *SqlV1MaterializedTable) GetStatus() SqlV1MaterializedTableStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SqlV1MaterializedTable) GetStatusOk() (*SqlV1MaterializedTableStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SqlV1MaterializedTable) SetStatus(v SqlV1MaterializedTableStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SqlV1MaterializedTable) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


