# SdV1PipelineSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the pipeline. | [optional] 
**Description** | Pointer to **string** | The description of the pipeline. | [optional] 
**Activated** | Pointer to **bool** | The desired state of the pipeline. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


