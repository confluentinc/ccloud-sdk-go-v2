# SdV1PipelineSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the pipeline. | [optional] 
**Description** | Pointer to **string** | The description of the pipeline. | [optional] 
**RetainedTopicNames** | Pointer to **[]string** | A list of Kafka topic names from the activated pipeline to be retained when this pipeline is deactivated.  | [optional] 
**Activated** | Pointer to **bool** | The desired state of the pipeline. | [optional] [default to false]
**ActivationPrivilege** | Pointer to **bool** | Whether the pipeline has privileges to be activated. | [optional] [default to false]
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 
**KafkaCluster** | Pointer to [**ObjectReference**](ObjectReference.md) | The kafka_cluster to which this belongs. | [optional] 

## Methods

### NewSdV1PipelineSpecUpdate

`func NewSdV1PipelineSpecUpdate() *SdV1PipelineSpecUpdate`

NewSdV1PipelineSpecUpdate instantiates a new SdV1PipelineSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdV1PipelineSpecUpdateWithDefaults

`func NewSdV1PipelineSpecUpdateWithDefaults() *SdV1PipelineSpecUpdate`

NewSdV1PipelineSpecUpdateWithDefaults instantiates a new SdV1PipelineSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *SdV1PipelineSpecUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SdV1PipelineSpecUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SdV1PipelineSpecUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SdV1PipelineSpecUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *SdV1PipelineSpecUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SdV1PipelineSpecUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SdV1PipelineSpecUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SdV1PipelineSpecUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRetainedTopicNames

`func (o *SdV1PipelineSpecUpdate) GetRetainedTopicNames() []string`

GetRetainedTopicNames returns the RetainedTopicNames field if non-nil, zero value otherwise.

### GetRetainedTopicNamesOk

`func (o *SdV1PipelineSpecUpdate) GetRetainedTopicNamesOk() (*[]string, bool)`

GetRetainedTopicNamesOk returns a tuple with the RetainedTopicNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainedTopicNames

`func (o *SdV1PipelineSpecUpdate) SetRetainedTopicNames(v []string)`

SetRetainedTopicNames sets RetainedTopicNames field to given value.

### HasRetainedTopicNames

`func (o *SdV1PipelineSpecUpdate) HasRetainedTopicNames() bool`

HasRetainedTopicNames returns a boolean if a field has been set.

### GetActivated

`func (o *SdV1PipelineSpecUpdate) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *SdV1PipelineSpecUpdate) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *SdV1PipelineSpecUpdate) SetActivated(v bool)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *SdV1PipelineSpecUpdate) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetActivationPrivilege

`func (o *SdV1PipelineSpecUpdate) GetActivationPrivilege() bool`

GetActivationPrivilege returns the ActivationPrivilege field if non-nil, zero value otherwise.

### GetActivationPrivilegeOk

`func (o *SdV1PipelineSpecUpdate) GetActivationPrivilegeOk() (*bool, bool)`

GetActivationPrivilegeOk returns a tuple with the ActivationPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationPrivilege

`func (o *SdV1PipelineSpecUpdate) SetActivationPrivilege(v bool)`

SetActivationPrivilege sets ActivationPrivilege field to given value.

### HasActivationPrivilege

`func (o *SdV1PipelineSpecUpdate) HasActivationPrivilege() bool`

HasActivationPrivilege returns a boolean if a field has been set.

### GetEnvironment

`func (o *SdV1PipelineSpecUpdate) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SdV1PipelineSpecUpdate) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SdV1PipelineSpecUpdate) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SdV1PipelineSpecUpdate) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetKafkaCluster

`func (o *SdV1PipelineSpecUpdate) GetKafkaCluster() ObjectReference`

GetKafkaCluster returns the KafkaCluster field if non-nil, zero value otherwise.

### GetKafkaClusterOk

`func (o *SdV1PipelineSpecUpdate) GetKafkaClusterOk() (*ObjectReference, bool)`

GetKafkaClusterOk returns a tuple with the KafkaCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaCluster

`func (o *SdV1PipelineSpecUpdate) SetKafkaCluster(v ObjectReference)`

SetKafkaCluster sets KafkaCluster field to given value.

### HasKafkaCluster

`func (o *SdV1PipelineSpecUpdate) HasKafkaCluster() bool`

HasKafkaCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


