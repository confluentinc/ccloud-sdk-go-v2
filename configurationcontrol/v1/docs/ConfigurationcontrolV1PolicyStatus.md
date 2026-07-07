# ConfigurationcontrolV1PolicyStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | Current lifecycle phase of the policy.   * &#x60;READY&#x60; — the policy has been accepted and is being evaluated     according to &#x60;mode&#x60;.   * &#x60;FAILED&#x60; — the policy is in an unrecoverable error state and is     not being evaluated. See &#x60;error_message&#x60; for details.  | [readonly] 
**ErrorMessage** | Pointer to **string** | Displayable error message if the policy is in a &#x60;FAILED&#x60; state.  | [optional] [readonly] 

## Methods

### NewConfigurationcontrolV1PolicyStatus

`func NewConfigurationcontrolV1PolicyStatus(phase string, ) *ConfigurationcontrolV1PolicyStatus`

NewConfigurationcontrolV1PolicyStatus instantiates a new ConfigurationcontrolV1PolicyStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationcontrolV1PolicyStatusWithDefaults

`func NewConfigurationcontrolV1PolicyStatusWithDefaults() *ConfigurationcontrolV1PolicyStatus`

NewConfigurationcontrolV1PolicyStatusWithDefaults instantiates a new ConfigurationcontrolV1PolicyStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *ConfigurationcontrolV1PolicyStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *ConfigurationcontrolV1PolicyStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *ConfigurationcontrolV1PolicyStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetErrorMessage

`func (o *ConfigurationcontrolV1PolicyStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ConfigurationcontrolV1PolicyStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ConfigurationcontrolV1PolicyStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ConfigurationcontrolV1PolicyStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


