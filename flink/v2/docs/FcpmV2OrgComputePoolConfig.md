# FcpmV2OrgComputePoolConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**OrganizationId** | Pointer to **string** | The ID of the organization. | [optional] [readonly] 
**Spec** | Pointer to [**FcpmV2OrgComputePoolConfigSpec**](FcpmV2OrgComputePoolConfigSpec.md) |  | [optional] 

## Methods

### NewFcpmV2OrgComputePoolConfig

`func NewFcpmV2OrgComputePoolConfig() *FcpmV2OrgComputePoolConfig`

NewFcpmV2OrgComputePoolConfig instantiates a new FcpmV2OrgComputePoolConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpmV2OrgComputePoolConfigWithDefaults

`func NewFcpmV2OrgComputePoolConfigWithDefaults() *FcpmV2OrgComputePoolConfig`

NewFcpmV2OrgComputePoolConfigWithDefaults instantiates a new FcpmV2OrgComputePoolConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *FcpmV2OrgComputePoolConfig) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FcpmV2OrgComputePoolConfig) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FcpmV2OrgComputePoolConfig) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *FcpmV2OrgComputePoolConfig) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *FcpmV2OrgComputePoolConfig) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FcpmV2OrgComputePoolConfig) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FcpmV2OrgComputePoolConfig) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FcpmV2OrgComputePoolConfig) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetOrganizationId

`func (o *FcpmV2OrgComputePoolConfig) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *FcpmV2OrgComputePoolConfig) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *FcpmV2OrgComputePoolConfig) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *FcpmV2OrgComputePoolConfig) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetSpec

`func (o *FcpmV2OrgComputePoolConfig) GetSpec() FcpmV2OrgComputePoolConfigSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *FcpmV2OrgComputePoolConfig) GetSpecOk() (*FcpmV2OrgComputePoolConfigSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *FcpmV2OrgComputePoolConfig) SetSpec(v FcpmV2OrgComputePoolConfigSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *FcpmV2OrgComputePoolConfig) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


