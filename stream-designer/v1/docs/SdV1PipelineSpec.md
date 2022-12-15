# SdV1PipelineSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the pipeline. | [optional] 
**Description** | Pointer to **string** | The description of the pipeline. | [optional] 
**RetainedTopicNames** | Pointer to **[]string** | A list of Kafka topic names from the activated pipeline to be retained when this pipeline is deactivated.  | [optional] 
**Activated** | Pointer to **bool** | The desired state of the pipeline. | [optional] [default to false]
**ActivationPrivilege** | Pointer to **bool** | Whether the pipeline has privileges to be activated. | [optional] [default to false]
**SourceCode** | Pointer to [**SdV1SourceCodeObject**](sd.v1.SourceCodeObject.md) | A map of source code format and statements that define this pipeline. | [optional] 
**Secrets** | Pointer to **map[string]string** | A map of secrets used in the pipeline source code. | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 
**KafkaCluster** | Pointer to [**ObjectReference**](ObjectReference.md) | The kafka_cluster to which this belongs. | [optional] 
**KsqlCluster** | Pointer to [**ObjectReference**](ObjectReference.md) | The ksql_cluster associated with this object. | [optional] 
**StreamGovernanceCluster** | Pointer to [**ObjectReference**](ObjectReference.md) | The stream_governance_cluster associated with this object. | [optional] 

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

### GetRetainedTopicNames

`func (o *SdV1PipelineSpec) GetRetainedTopicNames() []string`

GetRetainedTopicNames returns the RetainedTopicNames field if non-nil, zero value otherwise.

### GetRetainedTopicNamesOk

`func (o *SdV1PipelineSpec) GetRetainedTopicNamesOk() (*[]string, bool)`

GetRetainedTopicNamesOk returns a tuple with the RetainedTopicNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainedTopicNames

`func (o *SdV1PipelineSpec) SetRetainedTopicNames(v []string)`

SetRetainedTopicNames sets RetainedTopicNames field to given value.

### HasRetainedTopicNames

`func (o *SdV1PipelineSpec) HasRetainedTopicNames() bool`

HasRetainedTopicNames returns a boolean if a field has been set.

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

`func (o *SdV1PipelineSpec) GetActivationPrivilege() bool`

GetActivationPrivilege returns the ActivationPrivilege field if non-nil, zero value otherwise.

### GetActivationPrivilegeOk

`func (o *SdV1PipelineSpec) GetActivationPrivilegeOk() (*bool, bool)`

GetActivationPrivilegeOk returns a tuple with the ActivationPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationPrivilege

`func (o *SdV1PipelineSpec) SetActivationPrivilege(v bool)`

SetActivationPrivilege sets ActivationPrivilege field to given value.

### HasActivationPrivilege

`func (o *SdV1PipelineSpec) HasActivationPrivilege() bool`

HasActivationPrivilege returns a boolean if a field has been set.

### GetSourceCode

`func (o *SdV1PipelineSpec) GetSourceCode() SdV1SourceCodeObject`

GetSourceCode returns the SourceCode field if non-nil, zero value otherwise.

### GetSourceCodeOk

`func (o *SdV1PipelineSpec) GetSourceCodeOk() (*SdV1SourceCodeObject, bool)`

GetSourceCodeOk returns a tuple with the SourceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCode

`func (o *SdV1PipelineSpec) SetSourceCode(v SdV1SourceCodeObject)`

SetSourceCode sets SourceCode field to given value.

### HasSourceCode

`func (o *SdV1PipelineSpec) HasSourceCode() bool`

HasSourceCode returns a boolean if a field has been set.

### GetSecrets

`func (o *SdV1PipelineSpec) GetSecrets() map[string]string`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *SdV1PipelineSpec) GetSecretsOk() (*map[string]string, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *SdV1PipelineSpec) SetSecrets(v map[string]string)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *SdV1PipelineSpec) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

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


