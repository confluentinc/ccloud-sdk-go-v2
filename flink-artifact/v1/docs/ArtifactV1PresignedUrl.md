# ArtifactV1PresignedUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**ContentFormat** | Pointer to **string** | Content format of the Flink Artifact archive. | [optional] [readonly] 
**Cloud** | Pointer to **string** | Cloud provider where the Flink Artifact archive is uploaded. | [optional] [readonly] 
**Region** | Pointer to **string** | The Cloud provider region the Flink Artifact archive is uploaded. | [optional] [readonly] 
**UploadId** | Pointer to **string** | Unique identifier of this upload. | [optional] [readonly] 
**UploadUrl** | Pointer to **string** | Upload URL for the Flink Artifact archive. | [optional] [readonly] 
**UploadFormData** | Pointer to **map[string]interface{}** | Upload form data of the Flink Artifact. All values should be strings. | [optional] [readonly] 

## Methods

### NewArtifactV1PresignedUrl

`func NewArtifactV1PresignedUrl() *ArtifactV1PresignedUrl`

NewArtifactV1PresignedUrl instantiates a new ArtifactV1PresignedUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactV1PresignedUrlWithDefaults

`func NewArtifactV1PresignedUrlWithDefaults() *ArtifactV1PresignedUrl`

NewArtifactV1PresignedUrlWithDefaults instantiates a new ArtifactV1PresignedUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ArtifactV1PresignedUrl) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ArtifactV1PresignedUrl) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ArtifactV1PresignedUrl) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ArtifactV1PresignedUrl) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ArtifactV1PresignedUrl) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ArtifactV1PresignedUrl) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ArtifactV1PresignedUrl) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ArtifactV1PresignedUrl) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetContentFormat

`func (o *ArtifactV1PresignedUrl) GetContentFormat() string`

GetContentFormat returns the ContentFormat field if non-nil, zero value otherwise.

### GetContentFormatOk

`func (o *ArtifactV1PresignedUrl) GetContentFormatOk() (*string, bool)`

GetContentFormatOk returns a tuple with the ContentFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFormat

`func (o *ArtifactV1PresignedUrl) SetContentFormat(v string)`

SetContentFormat sets ContentFormat field to given value.

### HasContentFormat

`func (o *ArtifactV1PresignedUrl) HasContentFormat() bool`

HasContentFormat returns a boolean if a field has been set.

### GetCloud

`func (o *ArtifactV1PresignedUrl) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ArtifactV1PresignedUrl) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ArtifactV1PresignedUrl) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *ArtifactV1PresignedUrl) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetRegion

`func (o *ArtifactV1PresignedUrl) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ArtifactV1PresignedUrl) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ArtifactV1PresignedUrl) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ArtifactV1PresignedUrl) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetUploadId

`func (o *ArtifactV1PresignedUrl) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *ArtifactV1PresignedUrl) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *ArtifactV1PresignedUrl) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.

### HasUploadId

`func (o *ArtifactV1PresignedUrl) HasUploadId() bool`

HasUploadId returns a boolean if a field has been set.

### GetUploadUrl

`func (o *ArtifactV1PresignedUrl) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *ArtifactV1PresignedUrl) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *ArtifactV1PresignedUrl) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.

### HasUploadUrl

`func (o *ArtifactV1PresignedUrl) HasUploadUrl() bool`

HasUploadUrl returns a boolean if a field has been set.

### GetUploadFormData

`func (o *ArtifactV1PresignedUrl) GetUploadFormData() map[string]interface{}`

GetUploadFormData returns the UploadFormData field if non-nil, zero value otherwise.

### GetUploadFormDataOk

`func (o *ArtifactV1PresignedUrl) GetUploadFormDataOk() (*map[string]interface{}, bool)`

GetUploadFormDataOk returns a tuple with the UploadFormData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadFormData

`func (o *ArtifactV1PresignedUrl) SetUploadFormData(v map[string]interface{})`

SetUploadFormData sets UploadFormData field to given value.

### HasUploadFormData

`func (o *ArtifactV1PresignedUrl) HasUploadFormData() bool`

HasUploadFormData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


