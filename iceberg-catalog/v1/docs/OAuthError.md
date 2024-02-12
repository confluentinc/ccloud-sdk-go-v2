# OAuthError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**ErrorDescription** | Pointer to **string** |  | [optional] 
**ErrorUri** | Pointer to **string** |  | [optional] 

## Methods

### NewOAuthError

`func NewOAuthError(error_ string, ) *OAuthError`

NewOAuthError instantiates a new OAuthError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthErrorWithDefaults

`func NewOAuthErrorWithDefaults() *OAuthError`

NewOAuthErrorWithDefaults instantiates a new OAuthError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *OAuthError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OAuthError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OAuthError) SetError(v string)`

SetError sets Error field to given value.


### GetErrorDescription

`func (o *OAuthError) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *OAuthError) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *OAuthError) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *OAuthError) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetErrorUri

`func (o *OAuthError) GetErrorUri() string`

GetErrorUri returns the ErrorUri field if non-nil, zero value otherwise.

### GetErrorUriOk

`func (o *OAuthError) GetErrorUriOk() (*string, bool)`

GetErrorUriOk returns a tuple with the ErrorUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorUri

`func (o *OAuthError) SetErrorUri(v string)`

SetErrorUri sets ErrorUri field to given value.

### HasErrorUri

`func (o *OAuthError) HasErrorUri() bool`

HasErrorUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


