# StsV1TokenExchangeReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | An JWT access token, issued by Confluent, in response to the token exchange request. Client application could use the access token to access confluent public api  | 
**IssuedTokenType** | **string** | The token type. Always matches the value of requested_token_type from the request. | 
**TokenType** | **string** | Indicates the token type value. The only type that Confluent supports is Bearer | 
**ExpiresIn** | **int32** | The length of time, in seconds, that the access token is valid. | 

## Methods

### NewStsV1TokenExchangeReply

`func NewStsV1TokenExchangeReply(accessToken string, issuedTokenType string, tokenType string, expiresIn int32, ) *StsV1TokenExchangeReply`

NewStsV1TokenExchangeReply instantiates a new StsV1TokenExchangeReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStsV1TokenExchangeReplyWithDefaults

`func NewStsV1TokenExchangeReplyWithDefaults() *StsV1TokenExchangeReply`

NewStsV1TokenExchangeReplyWithDefaults instantiates a new StsV1TokenExchangeReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *StsV1TokenExchangeReply) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *StsV1TokenExchangeReply) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *StsV1TokenExchangeReply) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetIssuedTokenType

`func (o *StsV1TokenExchangeReply) GetIssuedTokenType() string`

GetIssuedTokenType returns the IssuedTokenType field if non-nil, zero value otherwise.

### GetIssuedTokenTypeOk

`func (o *StsV1TokenExchangeReply) GetIssuedTokenTypeOk() (*string, bool)`

GetIssuedTokenTypeOk returns a tuple with the IssuedTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedTokenType

`func (o *StsV1TokenExchangeReply) SetIssuedTokenType(v string)`

SetIssuedTokenType sets IssuedTokenType field to given value.


### GetTokenType

`func (o *StsV1TokenExchangeReply) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *StsV1TokenExchangeReply) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *StsV1TokenExchangeReply) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetExpiresIn

`func (o *StsV1TokenExchangeReply) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *StsV1TokenExchangeReply) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *StsV1TokenExchangeReply) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


