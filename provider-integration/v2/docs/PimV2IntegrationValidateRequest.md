# PimV2IntegrationValidateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | The ID of the provider integration to validate. | [optional] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Config** | Pointer to [**PimV2IntegrationValidateRequestConfigOneOf**](PimV2IntegrationValidateRequestConfigOneOf.md) | Cloud provider specific configuration for the provider integration. Required only for integrations in &#x60;DRAFT&#x60; status.  | [optional] 
**Environment** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewPimV2IntegrationValidateRequest

`func NewPimV2IntegrationValidateRequest() *PimV2IntegrationValidateRequest`

NewPimV2IntegrationValidateRequest instantiates a new PimV2IntegrationValidateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPimV2IntegrationValidateRequestWithDefaults

`func NewPimV2IntegrationValidateRequestWithDefaults() *PimV2IntegrationValidateRequest`

NewPimV2IntegrationValidateRequestWithDefaults instantiates a new PimV2IntegrationValidateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *PimV2IntegrationValidateRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *PimV2IntegrationValidateRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *PimV2IntegrationValidateRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *PimV2IntegrationValidateRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *PimV2IntegrationValidateRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PimV2IntegrationValidateRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PimV2IntegrationValidateRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PimV2IntegrationValidateRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *PimV2IntegrationValidateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PimV2IntegrationValidateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PimV2IntegrationValidateRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PimV2IntegrationValidateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *PimV2IntegrationValidateRequest) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PimV2IntegrationValidateRequest) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PimV2IntegrationValidateRequest) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PimV2IntegrationValidateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetConfig

`func (o *PimV2IntegrationValidateRequest) GetConfig() PimV2IntegrationValidateRequestConfigOneOf`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PimV2IntegrationValidateRequest) GetConfigOk() (*PimV2IntegrationValidateRequestConfigOneOf, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PimV2IntegrationValidateRequest) SetConfig(v PimV2IntegrationValidateRequestConfigOneOf)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *PimV2IntegrationValidateRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnvironment

`func (o *PimV2IntegrationValidateRequest) GetEnvironment() GlobalObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PimV2IntegrationValidateRequest) GetEnvironmentOk() (*GlobalObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PimV2IntegrationValidateRequest) SetEnvironment(v GlobalObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PimV2IntegrationValidateRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


