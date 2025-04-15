# CamV1ConnectArtifactSpec

## Properties

Name | Type | Description                                                               | Notes
------------ | ------------- |---------------------------------------------------------------------------| -------------
**Cloud** | **string** | Cloud provider where the Connect Artifact archive is uploaded.            | 
**Environment** | **string** | Environment the Connect Artifact belongs to.                              | 
**DisplayName** | **string** | Unique name of the Connect Artifact Archive per cloud, environment scope. | 
**Description** | Pointer to **string** | Description of the Connect Artifact.                                      | [optional] 
**ContentFormat** | Pointer to **string** | Archive format of the Connect Artifact.                                   | [optional] [readonly] 
**UploadSource** | Pointer to [**CamV1ConnectArtifactSpecUploadSourceOneOf**](CamV1ConnectArtifactSpecUploadSourceOneOf.md) | Upload source of the Connect Artifact.                                    | [optional] 
**Plugins** | Pointer to [**[]CamV1Plugins**](CamV1Plugins.md) | List of classes present in the Connect Artifact uploaded                  | [optional] [readonly] 
**Usages** | Pointer to **[]string** | List of resource crns where this Connect artifact is being used.          | [optional] [readonly] 

## Methods

### NewCamV1ConnectArtifactSpec

`func NewCamV1ConnectArtifactSpec(cloud string, environment string, displayName string, ) *CamV1ConnectArtifactSpec`

NewCamV1ConnectArtifactSpec instantiates a new CamV1ConnectArtifactSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCamV1ConnectArtifactSpecWithDefaults

`func NewCamV1ConnectArtifactSpecWithDefaults() *CamV1ConnectArtifactSpec`

NewCamV1ConnectArtifactSpecWithDefaults instantiates a new CamV1ConnectArtifactSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloud

`func (o *CamV1ConnectArtifactSpec) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *CamV1ConnectArtifactSpec) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *CamV1ConnectArtifactSpec) SetCloud(v string)`

SetCloud sets Cloud field to given value.


### GetEnvironment

`func (o *CamV1ConnectArtifactSpec) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CamV1ConnectArtifactSpec) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CamV1ConnectArtifactSpec) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetDisplayName

`func (o *CamV1ConnectArtifactSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CamV1ConnectArtifactSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CamV1ConnectArtifactSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *CamV1ConnectArtifactSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CamV1ConnectArtifactSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CamV1ConnectArtifactSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CamV1ConnectArtifactSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContentFormat

`func (o *CamV1ConnectArtifactSpec) GetContentFormat() string`

GetContentFormat returns the ContentFormat field if non-nil, zero value otherwise.

### GetContentFormatOk

`func (o *CamV1ConnectArtifactSpec) GetContentFormatOk() (*string, bool)`

GetContentFormatOk returns a tuple with the ContentFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFormat

`func (o *CamV1ConnectArtifactSpec) SetContentFormat(v string)`

SetContentFormat sets ContentFormat field to given value.

### HasContentFormat

`func (o *CamV1ConnectArtifactSpec) HasContentFormat() bool`

HasContentFormat returns a boolean if a field has been set.

### GetUploadSource

`func (o *CamV1ConnectArtifactSpec) GetUploadSource() CamV1ConnectArtifactSpecUploadSourceOneOf`

GetUploadSource returns the UploadSource field if non-nil, zero value otherwise.

### GetUploadSourceOk

`func (o *CamV1ConnectArtifactSpec) GetUploadSourceOk() (*CamV1ConnectArtifactSpecUploadSourceOneOf, bool)`

GetUploadSourceOk returns a tuple with the UploadSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadSource

`func (o *CamV1ConnectArtifactSpec) SetUploadSource(v CamV1ConnectArtifactSpecUploadSourceOneOf)`

SetUploadSource sets UploadSource field to given value.

### HasUploadSource

`func (o *CamV1ConnectArtifactSpec) HasUploadSource() bool`

HasUploadSource returns a boolean if a field has been set.

### GetPlugins

`func (o *CamV1ConnectArtifactSpec) GetPlugins() []CamV1Plugins`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *CamV1ConnectArtifactSpec) GetPluginsOk() (*[]CamV1Plugins, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *CamV1ConnectArtifactSpec) SetPlugins(v []CamV1Plugins)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *CamV1ConnectArtifactSpec) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.

### GetUsages

`func (o *CamV1ConnectArtifactSpec) GetUsages() []string`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *CamV1ConnectArtifactSpec) GetUsagesOk() (*[]string, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *CamV1ConnectArtifactSpec) SetUsages(v []string)`

SetUsages sets Usages field to given value.

### HasUsages

`func (o *CamV1ConnectArtifactSpec) HasUsages() bool`

HasUsages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


