# ConnectV1PresignedUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**ContentFormat** | Pointer to **string** | Content format of the Custom Connector Plugin archive. | [optional] [readonly] 
**Cloud** | Pointer to **string** | Cloud provider where the Custom Connector Plugin archive is uploaded. | [optional] [readonly] 
**RuntimeLanguage** | Pointer to **string** | Runtime language of Custom Connector Plugin. | [optional] [readonly] 
**UploadId** | Pointer to **string** | Unique identifier of this upload. | [optional] [readonly] 
**UploadUrl** | Pointer to **string** | Upload URL for the Custom Connector Plugin archive. | [optional] [readonly] 
**UploadFormData** | Pointer to **map[string]interface{}** | Upload form data of the Custom Connector Plugin. All values should be strings. | [optional] [readonly] 

## Methods

### NewConnectV1PresignedUrl

`func NewConnectV1PresignedUrl() *ConnectV1PresignedUrl`

NewConnectV1PresignedUrl instantiates a new ConnectV1PresignedUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1PresignedUrlWithDefaults

`func NewConnectV1PresignedUrlWithDefaults() *ConnectV1PresignedUrl`

NewConnectV1PresignedUrlWithDefaults instantiates a new ConnectV1PresignedUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ConnectV1PresignedUrl) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ConnectV1PresignedUrl) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ConnectV1PresignedUrl) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ConnectV1PresignedUrl) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ConnectV1PresignedUrl) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConnectV1PresignedUrl) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConnectV1PresignedUrl) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ConnectV1PresignedUrl) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetContentFormat

`func (o *ConnectV1PresignedUrl) GetContentFormat() string`

GetContentFormat returns the ContentFormat field if non-nil, zero value otherwise.

### GetContentFormatOk

`func (o *ConnectV1PresignedUrl) GetContentFormatOk() (*string, bool)`

GetContentFormatOk returns a tuple with the ContentFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFormat

`func (o *ConnectV1PresignedUrl) SetContentFormat(v string)`

SetContentFormat sets ContentFormat field to given value.

### HasContentFormat

`func (o *ConnectV1PresignedUrl) HasContentFormat() bool`

HasContentFormat returns a boolean if a field has been set.

### GetCloud

`func (o *ConnectV1PresignedUrl) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ConnectV1PresignedUrl) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ConnectV1PresignedUrl) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *ConnectV1PresignedUrl) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetRuntimeLanguage

`func (o *ConnectV1PresignedUrl) GetRuntimeLanguage() string`

GetRuntimeLanguage returns the RuntimeLanguage field if non-nil, zero value otherwise.

### GetRuntimeLanguageOk

`func (o *ConnectV1PresignedUrl) GetRuntimeLanguageOk() (*string, bool)`

GetRuntimeLanguageOk returns a tuple with the RuntimeLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeLanguage

`func (o *ConnectV1PresignedUrl) SetRuntimeLanguage(v string)`

SetRuntimeLanguage sets RuntimeLanguage field to given value.

### HasRuntimeLanguage

`func (o *ConnectV1PresignedUrl) HasRuntimeLanguage() bool`

HasRuntimeLanguage returns a boolean if a field has been set.

### GetUploadId

`func (o *ConnectV1PresignedUrl) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *ConnectV1PresignedUrl) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *ConnectV1PresignedUrl) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.

### HasUploadId

`func (o *ConnectV1PresignedUrl) HasUploadId() bool`

HasUploadId returns a boolean if a field has been set.

### GetUploadUrl

`func (o *ConnectV1PresignedUrl) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *ConnectV1PresignedUrl) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *ConnectV1PresignedUrl) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.

### HasUploadUrl

`func (o *ConnectV1PresignedUrl) HasUploadUrl() bool`

HasUploadUrl returns a boolean if a field has been set.

### GetUploadFormData

`func (o *ConnectV1PresignedUrl) GetUploadFormData() map[string]interface{}`

GetUploadFormData returns the UploadFormData field if non-nil, zero value otherwise.

### GetUploadFormDataOk

`func (o *ConnectV1PresignedUrl) GetUploadFormDataOk() (*map[string]interface{}, bool)`

GetUploadFormDataOk returns a tuple with the UploadFormData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadFormData

`func (o *ConnectV1PresignedUrl) SetUploadFormData(v map[string]interface{})`

SetUploadFormData sets UploadFormData field to given value.

### HasUploadFormData

`func (o *ConnectV1PresignedUrl) HasUploadFormData() bool`

HasUploadFormData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


