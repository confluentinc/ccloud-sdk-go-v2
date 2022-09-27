# SdV1PipelineSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the pipeline. | [optional] 
**Description** | Pointer to **string** | The description of the pipeline. | [optional] 
**Activated** | Pointer to **bool** | The desired state of the pipeline. | [optional] 
**ActivationPrivilege** | Pointer to **string** | Whether the pipeline has privileges it needs for activation and deactivation. | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 
**KafkaCluster** | Pointer to [**ObjectReference**](ObjectReference.md) | The kafka_cluster to which this belongs. | [optional] 
**KsqlCluster** | Pointer to [**ObjectReference**](ObjectReference.md) | The ksql_cluster to which this belongs. | [optional] 
**StreamGovernanceCluster** | Pointer to [**ObjectReference**](ObjectReference.md) | The stream_governance_cluster to which this belongs. | [optional] 

## Methods

### NewSdV1PipelineSpec

`func NewSdV1PipelineSpec() *SdV1PipelineSpec`

NewSdV1PipelineSpec instantiates a new SdV1PipelineSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdV1PipelineSpecWithDefaults

`func NewSdV1PipelineSpecWithDefaults() *SdV1PipelineSpec`

NewSdV1PipelineSpecWithDefaults instantiates a new SdV1PipelineSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *SdV1PipelineSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SdV1PipelineSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SdV1PipelineSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SdV1PipelineSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *SdV1PipelineSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SdV1PipelineSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SdV1PipelineSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SdV1PipelineSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActivated

`func (o *SdV1PipelineSpec) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *SdV1PipelineSpec) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *SdV1PipelineSpec) SetActivated(v bool)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *SdV1PipelineSpec) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetActivationPrivilege

`func (o *SdV1PipelineSpec) GetActivationPrivilege() string`

GetActivationPrivilege returns the ActivationPrivilege field if non-nil, zero value otherwise.

### GetActivationPrivilegeOk

`func (o *SdV1PipelineSpec) GetActivationPrivilegeOk() (*string, bool)`

GetActivationPrivilegeOk returns a tuple with the ActivationPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationPrivilege

`func (o *SdV1PipelineSpec) SetActivationPrivilege(v string)`

SetActivationPrivilege sets ActivationPrivilege field to given value.

### HasActivationPrivilege

`func (o *SdV1PipelineSpec) HasActivationPrivilege() bool`

HasActivationPrivilege returns a boolean if a field has been set.

### GetEnvironment

`func (o *SdV1PipelineSpec) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SdV1PipelineSpec) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SdV1PipelineSpec) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SdV1PipelineSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetKafkaCluster

`func (o *SdV1PipelineSpec) GetKafkaCluster() ObjectReference`

GetKafkaCluster returns the KafkaCluster field if non-nil, zero value otherwise.

### GetKafkaClusterOk

`func (o *SdV1PipelineSpec) GetKafkaClusterOk() (*ObjectReference, bool)`

GetKafkaClusterOk returns a tuple with the KafkaCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaCluster

`func (o *SdV1PipelineSpec) SetKafkaCluster(v ObjectReference)`

SetKafkaCluster sets KafkaCluster field to given value.

### HasKafkaCluster

`func (o *SdV1PipelineSpec) HasKafkaCluster() bool`

HasKafkaCluster returns a boolean if a field has been set.

### GetKsqlCluster

`func (o *SdV1PipelineSpec) GetKsqlCluster() ObjectReference`

GetKsqlCluster returns the KsqlCluster field if non-nil, zero value otherwise.

### GetKsqlClusterOk

`func (o *SdV1PipelineSpec) GetKsqlClusterOk() (*ObjectReference, bool)`

GetKsqlClusterOk returns a tuple with the KsqlCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKsqlCluster

`func (o *SdV1PipelineSpec) SetKsqlCluster(v ObjectReference)`

SetKsqlCluster sets KsqlCluster field to given value.

### HasKsqlCluster

`func (o *SdV1PipelineSpec) HasKsqlCluster() bool`

HasKsqlCluster returns a boolean if a field has been set.

### GetStreamGovernanceCluster

`func (o *SdV1PipelineSpec) GetStreamGovernanceCluster() ObjectReference`

GetStreamGovernanceCluster returns the StreamGovernanceCluster field if non-nil, zero value otherwise.

### GetStreamGovernanceClusterOk

`func (o *SdV1PipelineSpec) GetStreamGovernanceClusterOk() (*ObjectReference, bool)`

GetStreamGovernanceClusterOk returns a tuple with the StreamGovernanceCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamGovernanceCluster

`func (o *SdV1PipelineSpec) SetStreamGovernanceCluster(v ObjectReference)`

SetStreamGovernanceCluster sets StreamGovernanceCluster field to given value.

### HasStreamGovernanceCluster

`func (o *SdV1PipelineSpec) HasStreamGovernanceCluster() bool`

HasStreamGovernanceCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


