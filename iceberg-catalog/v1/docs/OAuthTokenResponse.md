# OAuthTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | The access token, for client credentials or token exchange | 
**TokenType** | **string** | Access token type for client credentials or token exchange  See https://datatracker.ietf.org/doc/html/rfc6749#section-7.1 | 
**ExpiresIn** | Pointer to **int32** | Lifetime of the access token in seconds for client credentials or token exchange | [optional] 
**IssuedTokenType** | Pointer to [**TokenType**](TokenType.md) |  | [optional] 
**RefreshToken** | Pointer to **string** | Refresh token for client credentials or token exchange | [optional] 
**Scope** | Pointer to **string** | Authorization scope for client credentials or token exchange | [optional] 

## Methods

### NewOAuthTokenResponse

`func NewOAuthTokenResponse(accessToken string, tokenType string, ) *OAuthTokenResponse`

NewOAuthTokenResponse instantiates a new OAuthTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthTokenResponseWithDefaults

`func NewOAuthTokenResponseWithDefaults() *OAuthTokenResponse`

NewOAuthTokenResponseWithDefaults instantiates a new OAuthTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *OAuthTokenResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *OAuthTokenResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *OAuthTokenResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetTokenType

`func (o *OAuthTokenResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *OAuthTokenResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *OAuthTokenResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetExpiresIn

`func (o *OAuthTokenResponse) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *OAuthTokenResponse) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *OAuthTokenResponse) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *OAuthTokenResponse) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetIssuedTokenType

`func (o *OAuthTokenResponse) GetIssuedTokenType() TokenType`

GetIssuedTokenType returns the IssuedTokenType field if non-nil, zero value otherwise.

### GetIssuedTokenTypeOk

`func (o *OAuthTokenResponse) GetIssuedTokenTypeOk() (*TokenType, bool)`

GetIssuedTokenTypeOk returns a tuple with the IssuedTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedTokenType

`func (o *OAuthTokenResponse) SetIssuedTokenType(v TokenType)`

SetIssuedTokenType sets IssuedTokenType field to given value.

### HasIssuedTokenType

`func (o *OAuthTokenResponse) HasIssuedTokenType() bool`

HasIssuedTokenType returns a boolean if a field has been set.

### GetRefreshToken

`func (o *OAuthTokenResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *OAuthTokenResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *OAuthTokenResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *OAuthTokenResponse) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetScope

`func (o *OAuthTokenResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OAuthTokenResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OAuthTokenResponse) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *OAuthTokenResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


