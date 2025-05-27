# CcpmV1CustomConnectPluginSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Display name of Custom Connect Plugin. | [optional] 
**Description** | Pointer to **string** | Description of Custom Connect Plugin. | [optional] 
**RuntimeLanguage** | Pointer to **string** | Runtime language of Custom Connect Plugin. | [optional] [readonly] 
**Environment** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewCcpmV1CustomConnectPluginSpecUpdate

`func NewCcpmV1CustomConnectPluginSpecUpdate() *CcpmV1CustomConnectPluginSpecUpdate`

NewCcpmV1CustomConnectPluginSpecUpdate instantiates a new CcpmV1CustomConnectPluginSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCcpmV1CustomConnectPluginSpecUpdateWithDefaults

`func NewCcpmV1CustomConnectPluginSpecUpdateWithDefaults() *CcpmV1CustomConnectPluginSpecUpdate`

NewCcpmV1CustomConnectPluginSpecUpdateWithDefaults instantiates a new CcpmV1CustomConnectPluginSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *CcpmV1CustomConnectPluginSpecUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CcpmV1CustomConnectPluginSpecUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CcpmV1CustomConnectPluginSpecUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CcpmV1CustomConnectPluginSpecUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *CcpmV1CustomConnectPluginSpecUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CcpmV1CustomConnectPluginSpecUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CcpmV1CustomConnectPluginSpecUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CcpmV1CustomConnectPluginSpecUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRuntimeLanguage

`func (o *CcpmV1CustomConnectPluginSpecUpdate) GetRuntimeLanguage() string`

GetRuntimeLanguage returns the RuntimeLanguage field if non-nil, zero value otherwise.

### GetRuntimeLanguageOk

`func (o *CcpmV1CustomConnectPluginSpecUpdate) GetRuntimeLanguageOk() (*string, bool)`

GetRuntimeLanguageOk returns a tuple with the RuntimeLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeLanguage

`func (o *CcpmV1CustomConnectPluginSpecUpdate) SetRuntimeLanguage(v string)`

SetRuntimeLanguage sets RuntimeLanguage field to given value.

### HasRuntimeLanguage

`func (o *CcpmV1CustomConnectPluginSpecUpdate) HasRuntimeLanguage() bool`

HasRuntimeLanguage returns a boolean if a field has been set.

### GetEnvironment

`func (o *CcpmV1CustomConnectPluginSpecUpdate) GetEnvironment() EnvScopedObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CcpmV1CustomConnectPluginSpecUpdate) GetEnvironmentOk() (*EnvScopedObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CcpmV1CustomConnectPluginSpecUpdate) SetEnvironment(v EnvScopedObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CcpmV1CustomConnectPluginSpecUpdate) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


