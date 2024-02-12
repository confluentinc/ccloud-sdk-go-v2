# SignS3Request200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** |  | 
**Headers** | [**S3Headers**](S3Headers.md) |  | 

## Methods

### NewSignS3Request200Response

`func NewSignS3Request200Response(uri string, headers S3Headers, ) *SignS3Request200Response`

NewSignS3Request200Response instantiates a new SignS3Request200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignS3Request200ResponseWithDefaults

`func NewSignS3Request200ResponseWithDefaults() *SignS3Request200Response`

NewSignS3Request200ResponseWithDefaults instantiates a new SignS3Request200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *SignS3Request200Response) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *SignS3Request200Response) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *SignS3Request200Response) SetUri(v string)`

SetUri sets Uri field to given value.


### GetHeaders

`func (o *SignS3Request200Response) GetHeaders() S3Headers`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *SignS3Request200Response) GetHeadersOk() (*S3Headers, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *SignS3Request200Response) SetHeaders(v S3Headers)`

SetHeaders sets Headers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


