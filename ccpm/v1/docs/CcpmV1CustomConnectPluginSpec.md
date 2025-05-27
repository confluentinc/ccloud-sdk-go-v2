# CcpmV1CustomConnectPluginSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cloud** | Pointer to **string** | Cloud provider where the Custom Connect Plugin archive is uploaded. | [optional] 
**DisplayName** | Pointer to **string** | Display name of Custom Connect Plugin. | [optional] 
**Description** | Pointer to **string** | Description of Custom Connect Plugin. | [optional] 
**RuntimeLanguage** | Pointer to **string** | Runtime language of Custom Connect Plugin. | [optional] [readonly] 
**Environment** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewCcpmV1CustomConnectPluginSpec

`func NewCcpmV1CustomConnectPluginSpec() *CcpmV1CustomConnectPluginSpec`

NewCcpmV1CustomConnectPluginSpec instantiates a new CcpmV1CustomConnectPluginSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCcpmV1CustomConnectPluginSpecWithDefaults

`func NewCcpmV1CustomConnectPluginSpecWithDefaults() *CcpmV1CustomConnectPluginSpec`

NewCcpmV1CustomConnectPluginSpecWithDefaults instantiates a new CcpmV1CustomConnectPluginSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloud

`func (o *CcpmV1CustomConnectPluginSpec) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *CcpmV1CustomConnectPluginSpec) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *CcpmV1CustomConnectPluginSpec) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *CcpmV1CustomConnectPluginSpec) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetDisplayName

`func (o *CcpmV1CustomConnectPluginSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CcpmV1CustomConnectPluginSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CcpmV1CustomConnectPluginSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CcpmV1CustomConnectPluginSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *CcpmV1CustomConnectPluginSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CcpmV1CustomConnectPluginSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CcpmV1CustomConnectPluginSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CcpmV1CustomConnectPluginSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRuntimeLanguage

`func (o *CcpmV1CustomConnectPluginSpec) GetRuntimeLanguage() string`

GetRuntimeLanguage returns the RuntimeLanguage field if non-nil, zero value otherwise.

### GetRuntimeLanguageOk

`func (o *CcpmV1CustomConnectPluginSpec) GetRuntimeLanguageOk() (*string, bool)`

GetRuntimeLanguageOk returns a tuple with the RuntimeLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeLanguage

`func (o *CcpmV1CustomConnectPluginSpec) SetRuntimeLanguage(v string)`

SetRuntimeLanguage sets RuntimeLanguage field to given value.

### HasRuntimeLanguage

`func (o *CcpmV1CustomConnectPluginSpec) HasRuntimeLanguage() bool`

HasRuntimeLanguage returns a boolean if a field has been set.

### GetEnvironment

`func (o *CcpmV1CustomConnectPluginSpec) GetEnvironment() EnvScopedObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CcpmV1CustomConnectPluginSpec) GetEnvironmentOk() (*EnvScopedObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CcpmV1CustomConnectPluginSpec) SetEnvironment(v EnvScopedObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CcpmV1CustomConnectPluginSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


