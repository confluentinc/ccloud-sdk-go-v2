# CcpmV1PresignedUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**ContentFormat** | Pointer to **string** | Content format of the Custom Connect Plugin archive. | [optional] 
**Cloud** | Pointer to **string** | Cloud provider where the Custom Connect Plugin archive is uploaded. | [optional] 
**UploadId** | Pointer to **string** | Unique identifier of this upload. | [optional] [readonly] 
**UploadUrl** | Pointer to **string** | Upload URL for the Custom Connect Plugin archive. | [optional] [readonly] 
**UploadFormData** | Pointer to **map[string]interface{}** | Upload form data of the Custom Connect Plugin. All values should be strings. | [optional] [readonly] 
**Environment** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewCcpmV1PresignedUrl

`func NewCcpmV1PresignedUrl() *CcpmV1PresignedUrl`

NewCcpmV1PresignedUrl instantiates a new CcpmV1PresignedUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCcpmV1PresignedUrlWithDefaults

`func NewCcpmV1PresignedUrlWithDefaults() *CcpmV1PresignedUrl`

NewCcpmV1PresignedUrlWithDefaults instantiates a new CcpmV1PresignedUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CcpmV1PresignedUrl) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CcpmV1PresignedUrl) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CcpmV1PresignedUrl) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CcpmV1PresignedUrl) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CcpmV1PresignedUrl) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CcpmV1PresignedUrl) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CcpmV1PresignedUrl) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CcpmV1PresignedUrl) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetContentFormat

`func (o *CcpmV1PresignedUrl) GetContentFormat() string`

GetContentFormat returns the ContentFormat field if non-nil, zero value otherwise.

### GetContentFormatOk

`func (o *CcpmV1PresignedUrl) GetContentFormatOk() (*string, bool)`

GetContentFormatOk returns a tuple with the ContentFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFormat

`func (o *CcpmV1PresignedUrl) SetContentFormat(v string)`

SetContentFormat sets ContentFormat field to given value.

### HasContentFormat

`func (o *CcpmV1PresignedUrl) HasContentFormat() bool`

HasContentFormat returns a boolean if a field has been set.

### GetCloud

`func (o *CcpmV1PresignedUrl) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *CcpmV1PresignedUrl) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *CcpmV1PresignedUrl) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *CcpmV1PresignedUrl) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetUploadId

`func (o *CcpmV1PresignedUrl) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *CcpmV1PresignedUrl) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *CcpmV1PresignedUrl) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.

### HasUploadId

`func (o *CcpmV1PresignedUrl) HasUploadId() bool`

HasUploadId returns a boolean if a field has been set.

### GetUploadUrl

`func (o *CcpmV1PresignedUrl) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *CcpmV1PresignedUrl) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *CcpmV1PresignedUrl) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.

### HasUploadUrl

`func (o *CcpmV1PresignedUrl) HasUploadUrl() bool`

HasUploadUrl returns a boolean if a field has been set.

### GetUploadFormData

`func (o *CcpmV1PresignedUrl) GetUploadFormData() map[string]interface{}`

GetUploadFormData returns the UploadFormData field if non-nil, zero value otherwise.

### GetUploadFormDataOk

`func (o *CcpmV1PresignedUrl) GetUploadFormDataOk() (*map[string]interface{}, bool)`

GetUploadFormDataOk returns a tuple with the UploadFormData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadFormData

`func (o *CcpmV1PresignedUrl) SetUploadFormData(v map[string]interface{})`

SetUploadFormData sets UploadFormData field to given value.

### HasUploadFormData

`func (o *CcpmV1PresignedUrl) HasUploadFormData() bool`

HasUploadFormData returns a boolean if a field has been set.

### GetEnvironment

`func (o *CcpmV1PresignedUrl) GetEnvironment() EnvScopedObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CcpmV1PresignedUrl) GetEnvironmentOk() (*EnvScopedObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CcpmV1PresignedUrl) SetEnvironment(v EnvScopedObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CcpmV1PresignedUrl) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


