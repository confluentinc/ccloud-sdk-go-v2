# S3SignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** |  | 
**Uri** | **string** |  | 
**Method** | **string** |  | 
**Headers** | [**S3Headers**](S3Headers.md) |  | 
**Properties** | Pointer to **map[string]string** |  | [optional] 
**Body** | Pointer to **string** | Optional body of the S3 request to send to the signing API. This should only be populated for S3 requests which do not have the relevant data in the URI itself (e.g. DeleteObjects requests) | [optional] 

## Methods

### NewS3SignRequest

`func NewS3SignRequest(region string, uri string, method string, headers S3Headers, ) *S3SignRequest`

NewS3SignRequest instantiates a new S3SignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3SignRequestWithDefaults

`func NewS3SignRequestWithDefaults() *S3SignRequest`

NewS3SignRequestWithDefaults instantiates a new S3SignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *S3SignRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *S3SignRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *S3SignRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetUri

`func (o *S3SignRequest) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *S3SignRequest) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *S3SignRequest) SetUri(v string)`

SetUri sets Uri field to given value.


### GetMethod

`func (o *S3SignRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *S3SignRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *S3SignRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetHeaders

`func (o *S3SignRequest) GetHeaders() S3Headers`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *S3SignRequest) GetHeadersOk() (*S3Headers, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *S3SignRequest) SetHeaders(v S3Headers)`

SetHeaders sets Headers field to given value.


### GetProperties

`func (o *S3SignRequest) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *S3SignRequest) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *S3SignRequest) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *S3SignRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetBody

`func (o *S3SignRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *S3SignRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *S3SignRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *S3SignRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


