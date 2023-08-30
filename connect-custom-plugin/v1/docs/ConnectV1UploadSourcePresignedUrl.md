# ConnectV1UploadSourcePresignedUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **string** | Location of Custom Connector Plugin source.  | 
**UploadId** | **string** | Upload id returned by &#x60;/presigned-upload-url&#x60; API. This field returns empty string in all responses. | 

## Methods

### NewConnectV1UploadSourcePresignedUrl

`func NewConnectV1UploadSourcePresignedUrl(location string, uploadId string, ) *ConnectV1UploadSourcePresignedUrl`

NewConnectV1UploadSourcePresignedUrl instantiates a new ConnectV1UploadSourcePresignedUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1UploadSourcePresignedUrlWithDefaults

`func NewConnectV1UploadSourcePresignedUrlWithDefaults() *ConnectV1UploadSourcePresignedUrl`

NewConnectV1UploadSourcePresignedUrlWithDefaults instantiates a new ConnectV1UploadSourcePresignedUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *ConnectV1UploadSourcePresignedUrl) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ConnectV1UploadSourcePresignedUrl) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ConnectV1UploadSourcePresignedUrl) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetUploadId

`func (o *ConnectV1UploadSourcePresignedUrl) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *ConnectV1UploadSourcePresignedUrl) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *ConnectV1UploadSourcePresignedUrl) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.



### AsConnectV1CustomConnectorPluginUpdateUploadSourceOneOf

`func (s *ConnectV1UploadSourcePresignedUrl) AsConnectV1CustomConnectorPluginUpdateUploadSourceOneOf() ConnectV1CustomConnectorPluginUpdateUploadSourceOneOf`

Convenience method to wrap this instance of ConnectV1UploadSourcePresignedUrl in ConnectV1CustomConnectorPluginUpdateUploadSourceOneOf

### AsConnectV1CustomConnectorPluginUploadSourceOneOf

`func (s *ConnectV1UploadSourcePresignedUrl) AsConnectV1CustomConnectorPluginUploadSourceOneOf() ConnectV1CustomConnectorPluginUploadSourceOneOf`

Convenience method to wrap this instance of ConnectV1UploadSourcePresignedUrl in ConnectV1CustomConnectorPluginUploadSourceOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


