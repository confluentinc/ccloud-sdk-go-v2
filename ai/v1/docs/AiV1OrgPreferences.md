# AiV1OrgPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**AiAssistantEnabled** | Pointer to **bool** | Enable ai-assist for the organization | [optional] 

## Methods

### NewAiV1OrgPreferences

`func NewAiV1OrgPreferences() *AiV1OrgPreferences`

NewAiV1OrgPreferences instantiates a new AiV1OrgPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiV1OrgPreferencesWithDefaults

`func NewAiV1OrgPreferencesWithDefaults() *AiV1OrgPreferences`

NewAiV1OrgPreferencesWithDefaults instantiates a new AiV1OrgPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *AiV1OrgPreferences) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AiV1OrgPreferences) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AiV1OrgPreferences) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AiV1OrgPreferences) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *AiV1OrgPreferences) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AiV1OrgPreferences) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AiV1OrgPreferences) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AiV1OrgPreferences) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetAiAssistantEnabled

`func (o *AiV1OrgPreferences) GetAiAssistantEnabled() bool`

GetAiAssistantEnabled returns the AiAssistantEnabled field if non-nil, zero value otherwise.

### GetAiAssistantEnabledOk

`func (o *AiV1OrgPreferences) GetAiAssistantEnabledOk() (*bool, bool)`

GetAiAssistantEnabledOk returns a tuple with the AiAssistantEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiAssistantEnabled

`func (o *AiV1OrgPreferences) SetAiAssistantEnabled(v bool)`

SetAiAssistantEnabled sets AiAssistantEnabled field to given value.

### HasAiAssistantEnabled

`func (o *AiV1OrgPreferences) HasAiAssistantEnabled() bool`

HasAiAssistantEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


