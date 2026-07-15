# ConfigurationcontrolV1PolicyUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**ConfigurationcontrolV1PolicySpecUpdate**](ConfigurationcontrolV1PolicySpecUpdate.md) |  | [optional] 
**Status** | Pointer to [**ConfigurationcontrolV1PolicyStatus**](ConfigurationcontrolV1PolicyStatus.md) |  | [optional] 

## Methods

### NewConfigurationcontrolV1PolicyUpdate

`func NewConfigurationcontrolV1PolicyUpdate() *ConfigurationcontrolV1PolicyUpdate`

NewConfigurationcontrolV1PolicyUpdate instantiates a new ConfigurationcontrolV1PolicyUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationcontrolV1PolicyUpdateWithDefaults

`func NewConfigurationcontrolV1PolicyUpdateWithDefaults() *ConfigurationcontrolV1PolicyUpdate`

NewConfigurationcontrolV1PolicyUpdateWithDefaults instantiates a new ConfigurationcontrolV1PolicyUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ConfigurationcontrolV1PolicyUpdate) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ConfigurationcontrolV1PolicyUpdate) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ConfigurationcontrolV1PolicyUpdate) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ConfigurationcontrolV1PolicyUpdate) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ConfigurationcontrolV1PolicyUpdate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConfigurationcontrolV1PolicyUpdate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConfigurationcontrolV1PolicyUpdate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ConfigurationcontrolV1PolicyUpdate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *ConfigurationcontrolV1PolicyUpdate) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConfigurationcontrolV1PolicyUpdate) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConfigurationcontrolV1PolicyUpdate) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConfigurationcontrolV1PolicyUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *ConfigurationcontrolV1PolicyUpdate) GetSpec() ConfigurationcontrolV1PolicySpecUpdate`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ConfigurationcontrolV1PolicyUpdate) GetSpecOk() (*ConfigurationcontrolV1PolicySpecUpdate, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ConfigurationcontrolV1PolicyUpdate) SetSpec(v ConfigurationcontrolV1PolicySpecUpdate)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ConfigurationcontrolV1PolicyUpdate) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *ConfigurationcontrolV1PolicyUpdate) GetStatus() ConfigurationcontrolV1PolicyStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConfigurationcontrolV1PolicyUpdate) GetStatusOk() (*ConfigurationcontrolV1PolicyStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConfigurationcontrolV1PolicyUpdate) SetStatus(v ConfigurationcontrolV1PolicyStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConfigurationcontrolV1PolicyUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


