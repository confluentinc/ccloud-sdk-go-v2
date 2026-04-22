# SqlV1MaterializedTableVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ObjectMeta**](ObjectMeta.md) |  | 
**Name** | **string** | The resource version name, unique within the Kafka cluster. Name conforms to DNS Subdomain (RFC 1123).  | [readonly] 
**OrganizationId** | **string** | The unique identifier for the organization. | [readonly] 
**EnvironmentId** | **string** | The unique identifier for the environment. | [readonly] 
**Spec** | [**SqlV1MaterializedTableVersionSpec**](SqlV1MaterializedTableVersionSpec.md) |  | 

## Methods

### NewSqlV1MaterializedTableVersion

`func NewSqlV1MaterializedTableVersion(apiVersion string, kind string, metadata ObjectMeta, name string, organizationId string, environmentId string, spec SqlV1MaterializedTableVersionSpec, ) *SqlV1MaterializedTableVersion`

NewSqlV1MaterializedTableVersion instantiates a new SqlV1MaterializedTableVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1MaterializedTableVersionWithDefaults

`func NewSqlV1MaterializedTableVersionWithDefaults() *SqlV1MaterializedTableVersion`

NewSqlV1MaterializedTableVersionWithDefaults instantiates a new SqlV1MaterializedTableVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SqlV1MaterializedTableVersion) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SqlV1MaterializedTableVersion) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SqlV1MaterializedTableVersion) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *SqlV1MaterializedTableVersion) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SqlV1MaterializedTableVersion) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SqlV1MaterializedTableVersion) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *SqlV1MaterializedTableVersion) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SqlV1MaterializedTableVersion) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SqlV1MaterializedTableVersion) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.


### GetName

`func (o *SqlV1MaterializedTableVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlV1MaterializedTableVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlV1MaterializedTableVersion) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *SqlV1MaterializedTableVersion) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SqlV1MaterializedTableVersion) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SqlV1MaterializedTableVersion) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetEnvironmentId

`func (o *SqlV1MaterializedTableVersion) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *SqlV1MaterializedTableVersion) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *SqlV1MaterializedTableVersion) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetSpec

`func (o *SqlV1MaterializedTableVersion) GetSpec() SqlV1MaterializedTableVersionSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SqlV1MaterializedTableVersion) GetSpecOk() (*SqlV1MaterializedTableVersionSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SqlV1MaterializedTableVersion) SetSpec(v SqlV1MaterializedTableVersionSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


