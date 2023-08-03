# CcpV1UploadSourcePresignedUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **string** | Location of custom plugin source.  | 
**UploadId** | **string** | Upload id returned by &#x60;/presigned-upload-url&#x60; API. This field returns empty string in all responses. | 

## Methods

### NewCcpV1UploadSourcePresignedUrl

`func NewCcpV1UploadSourcePresignedUrl(location string, uploadId string, ) *CcpV1UploadSourcePresignedUrl`

NewCcpV1UploadSourcePresignedUrl instantiates a new CcpV1UploadSourcePresignedUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCcpV1UploadSourcePresignedUrlWithDefaults

`func NewCcpV1UploadSourcePresignedUrlWithDefaults() *CcpV1UploadSourcePresignedUrl`

NewCcpV1UploadSourcePresignedUrlWithDefaults instantiates a new CcpV1UploadSourcePresignedUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *CcpV1UploadSourcePresignedUrl) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CcpV1UploadSourcePresignedUrl) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CcpV1UploadSourcePresignedUrl) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetUploadId

`func (o *CcpV1UploadSourcePresignedUrl) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *CcpV1UploadSourcePresignedUrl) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *CcpV1UploadSourcePresignedUrl) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.



### AsCcpV1CustomPluginUploadSourceOneOf

`func (s *CcpV1UploadSourcePresignedUrl) AsCcpV1CustomPluginUploadSourceOneOf() CcpV1CustomPluginUploadSourceOneOf`

Convenience method to wrap this instance of CcpV1UploadSourcePresignedUrl in CcpV1CustomPluginUploadSourceOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


