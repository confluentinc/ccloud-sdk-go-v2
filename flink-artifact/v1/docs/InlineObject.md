# InlineObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cloud** | **string** | Cloud provider where the Flink Artifact archive is uploaded. | 
**Region** | **string** | The Cloud provider region the Flink Artifact archive is uploaded. | 
**Environment** | **string** | Environment the Flink Artifact belongs to. | 
**DisplayName** | **string** | Display name of the Flink Artifact. | 
**Class** | **string** | Java class or alias for the artifact as provided by developer. | 
**ContentFormat** | Pointer to **string** | Archive format of the Flink Artifact. | [optional] 
**Description** | Pointer to **string** | Description of the Flink Artifact. | [optional] 
**DocumentationLink** | Pointer to **string** | Document link of the Flink Artifact. | [optional] 
**RuntimeLanguage** | Pointer to **string** | Runtime language of the Flink Artifact. | [optional] [default to "JAVA"]
**UploadSource** | [**InlineObjectUploadSourceOneOf**](InlineObjectUploadSourceOneOf.md) | Upload source of the Flink Artifact source. | 

## Methods

### NewInlineObject

`func NewInlineObject(cloud string, region string, environment string, displayName string, class string, uploadSource InlineObjectUploadSourceOneOf, ) *InlineObject`

NewInlineObject instantiates a new InlineObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObjectWithDefaults

`func NewInlineObjectWithDefaults() *InlineObject`

NewInlineObjectWithDefaults instantiates a new InlineObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloud

`func (o *InlineObject) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *InlineObject) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *InlineObject) SetCloud(v string)`

SetCloud sets Cloud field to given value.


### GetRegion

`func (o *InlineObject) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InlineObject) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InlineObject) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetEnvironment

`func (o *InlineObject) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *InlineObject) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *InlineObject) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetDisplayName

`func (o *InlineObject) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InlineObject) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InlineObject) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetClass

`func (o *InlineObject) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *InlineObject) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *InlineObject) SetClass(v string)`

SetClass sets Class field to given value.


### GetContentFormat

`func (o *InlineObject) GetContentFormat() string`

GetContentFormat returns the ContentFormat field if non-nil, zero value otherwise.

### GetContentFormatOk

`func (o *InlineObject) GetContentFormatOk() (*string, bool)`

GetContentFormatOk returns a tuple with the ContentFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFormat

`func (o *InlineObject) SetContentFormat(v string)`

SetContentFormat sets ContentFormat field to given value.

### HasContentFormat

`func (o *InlineObject) HasContentFormat() bool`

HasContentFormat returns a boolean if a field has been set.

### GetDescription

`func (o *InlineObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentationLink

`func (o *InlineObject) GetDocumentationLink() string`

GetDocumentationLink returns the DocumentationLink field if non-nil, zero value otherwise.

### GetDocumentationLinkOk

`func (o *InlineObject) GetDocumentationLinkOk() (*string, bool)`

GetDocumentationLinkOk returns a tuple with the DocumentationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationLink

`func (o *InlineObject) SetDocumentationLink(v string)`

SetDocumentationLink sets DocumentationLink field to given value.

### HasDocumentationLink

`func (o *InlineObject) HasDocumentationLink() bool`

HasDocumentationLink returns a boolean if a field has been set.

### GetRuntimeLanguage

`func (o *InlineObject) GetRuntimeLanguage() string`

GetRuntimeLanguage returns the RuntimeLanguage field if non-nil, zero value otherwise.

### GetRuntimeLanguageOk

`func (o *InlineObject) GetRuntimeLanguageOk() (*string, bool)`

GetRuntimeLanguageOk returns a tuple with the RuntimeLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeLanguage

`func (o *InlineObject) SetRuntimeLanguage(v string)`

SetRuntimeLanguage sets RuntimeLanguage field to given value.

### HasRuntimeLanguage

`func (o *InlineObject) HasRuntimeLanguage() bool`

HasRuntimeLanguage returns a boolean if a field has been set.

### GetUploadSource

`func (o *InlineObject) GetUploadSource() InlineObjectUploadSourceOneOf`

GetUploadSource returns the UploadSource field if non-nil, zero value otherwise.

### GetUploadSourceOk

`func (o *InlineObject) GetUploadSourceOk() (*InlineObjectUploadSourceOneOf, bool)`

GetUploadSourceOk returns a tuple with the UploadSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadSource

`func (o *InlineObject) SetUploadSource(v InlineObjectUploadSourceOneOf)`

SetUploadSource sets UploadSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


