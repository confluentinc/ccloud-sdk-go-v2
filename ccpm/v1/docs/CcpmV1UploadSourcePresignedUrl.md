# CcpmV1UploadSourcePresignedUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **string** | Location of the Custom Connect Plugin source.  | 
**UploadId** | **string** | Upload ID returned by the &#x60;/presigned-upload-url&#x60; API. This field returns an empty string in all responses. | 

## Methods

### NewCcpmV1UploadSourcePresignedUrl

`func NewCcpmV1UploadSourcePresignedUrl(location string, uploadId string, ) *CcpmV1UploadSourcePresignedUrl`

NewCcpmV1UploadSourcePresignedUrl instantiates a new CcpmV1UploadSourcePresignedUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCcpmV1UploadSourcePresignedUrlWithDefaults

`func NewCcpmV1UploadSourcePresignedUrlWithDefaults() *CcpmV1UploadSourcePresignedUrl`

NewCcpmV1UploadSourcePresignedUrlWithDefaults instantiates a new CcpmV1UploadSourcePresignedUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *CcpmV1UploadSourcePresignedUrl) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CcpmV1UploadSourcePresignedUrl) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CcpmV1UploadSourcePresignedUrl) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetUploadId

`func (o *CcpmV1UploadSourcePresignedUrl) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *CcpmV1UploadSourcePresignedUrl) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *CcpmV1UploadSourcePresignedUrl) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.



### AsCcpmV1CustomConnectPluginVersionSpecUploadSourceOneOf

`func (s *CcpmV1UploadSourcePresignedUrl) AsCcpmV1CustomConnectPluginVersionSpecUploadSourceOneOf() CcpmV1CustomConnectPluginVersionSpecUploadSourceOneOf`

Convenience method to wrap this instance of CcpmV1UploadSourcePresignedUrl in CcpmV1CustomConnectPluginVersionSpecUploadSourceOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


