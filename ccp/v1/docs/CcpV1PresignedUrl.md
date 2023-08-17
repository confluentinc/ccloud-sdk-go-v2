# CcpV1PresignedUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**PluginType** | Pointer to **string** | Custom plugin type. | [optional] [readonly] 
**ContentFormat** | Pointer to **string** | Content format of custom plugin archive. | [optional] [readonly] 
**UploadId** | Pointer to **string** | Unique identifier of this upload. | [optional] [readonly] 
**UploadUrl** | Pointer to **string** | Upload url for custom plugin archive. | [optional] [readonly] 
**UploadFormData** | Pointer to **map[string]interface{}** | Upload form data of custom plugin. All values should be strings. | [optional] [readonly] 

## Methods

### NewCcpV1PresignedUrl

`func NewCcpV1PresignedUrl() *CcpV1PresignedUrl`

NewCcpV1PresignedUrl instantiates a new CcpV1PresignedUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCcpV1PresignedUrlWithDefaults

`func NewCcpV1PresignedUrlWithDefaults() *CcpV1PresignedUrl`

NewCcpV1PresignedUrlWithDefaults instantiates a new CcpV1PresignedUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CcpV1PresignedUrl) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CcpV1PresignedUrl) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CcpV1PresignedUrl) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CcpV1PresignedUrl) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CcpV1PresignedUrl) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CcpV1PresignedUrl) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CcpV1PresignedUrl) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CcpV1PresignedUrl) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPluginType

`func (o *CcpV1PresignedUrl) GetPluginType() string`

GetPluginType returns the PluginType field if non-nil, zero value otherwise.

### GetPluginTypeOk

`func (o *CcpV1PresignedUrl) GetPluginTypeOk() (*string, bool)`

GetPluginTypeOk returns a tuple with the PluginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginType

`func (o *CcpV1PresignedUrl) SetPluginType(v string)`

SetPluginType sets PluginType field to given value.

### HasPluginType

`func (o *CcpV1PresignedUrl) HasPluginType() bool`

HasPluginType returns a boolean if a field has been set.

### GetContentFormat

`func (o *CcpV1PresignedUrl) GetContentFormat() string`

GetContentFormat returns the ContentFormat field if non-nil, zero value otherwise.

### GetContentFormatOk

`func (o *CcpV1PresignedUrl) GetContentFormatOk() (*string, bool)`

GetContentFormatOk returns a tuple with the ContentFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFormat

`func (o *CcpV1PresignedUrl) SetContentFormat(v string)`

SetContentFormat sets ContentFormat field to given value.

### HasContentFormat

`func (o *CcpV1PresignedUrl) HasContentFormat() bool`

HasContentFormat returns a boolean if a field has been set.

### GetUploadId

`func (o *CcpV1PresignedUrl) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *CcpV1PresignedUrl) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *CcpV1PresignedUrl) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.

### HasUploadId

`func (o *CcpV1PresignedUrl) HasUploadId() bool`

HasUploadId returns a boolean if a field has been set.

### GetUploadUrl

`func (o *CcpV1PresignedUrl) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *CcpV1PresignedUrl) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *CcpV1PresignedUrl) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.

### HasUploadUrl

`func (o *CcpV1PresignedUrl) HasUploadUrl() bool`

HasUploadUrl returns a boolean if a field has been set.

### GetUploadFormData

`func (o *CcpV1PresignedUrl) GetUploadFormData() map[string]interface{}`

GetUploadFormData returns the UploadFormData field if non-nil, zero value otherwise.

### GetUploadFormDataOk

`func (o *CcpV1PresignedUrl) GetUploadFormDataOk() (*map[string]interface{}, bool)`

GetUploadFormDataOk returns a tuple with the UploadFormData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadFormData

`func (o *CcpV1PresignedUrl) SetUploadFormData(v map[string]interface{})`

SetUploadFormData sets UploadFormData field to given value.

### HasUploadFormData

`func (o *CcpV1PresignedUrl) HasUploadFormData() bool`

HasUploadFormData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


