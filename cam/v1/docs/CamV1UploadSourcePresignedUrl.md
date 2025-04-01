# CamV1UploadSourcePresignedUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **string** | Location of the Connect Artifact source.  | 
**UploadId** | **string** | Upload ID returned by the &#x60;/presigned-upload-url&#x60; API. This field returns an empty string in all responses. | 

## Methods

### NewCamV1UploadSourcePresignedUrl

`func NewCamV1UploadSourcePresignedUrl(location string, uploadId string, ) *CamV1UploadSourcePresignedUrl`

NewCamV1UploadSourcePresignedUrl instantiates a new CamV1UploadSourcePresignedUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCamV1UploadSourcePresignedUrlWithDefaults

`func NewCamV1UploadSourcePresignedUrlWithDefaults() *CamV1UploadSourcePresignedUrl`

NewCamV1UploadSourcePresignedUrlWithDefaults instantiates a new CamV1UploadSourcePresignedUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *CamV1UploadSourcePresignedUrl) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CamV1UploadSourcePresignedUrl) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CamV1UploadSourcePresignedUrl) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetUploadId

`func (o *CamV1UploadSourcePresignedUrl) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *CamV1UploadSourcePresignedUrl) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *CamV1UploadSourcePresignedUrl) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.



### AsCamV1ConnectArtifactSpecUploadSourceOneOf

`func (s *CamV1UploadSourcePresignedUrl) AsCamV1ConnectArtifactSpecUploadSourceOneOf() CamV1ConnectArtifactSpecUploadSourceOneOf`

Convenience method to wrap this instance of CamV1UploadSourcePresignedUrl in CamV1ConnectArtifactSpecUploadSourceOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


