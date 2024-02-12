# \OAuth2APIAPI

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetToken**](OAuth2APIAPI.md#GetToken) | **Post** /v1/oauth/tokens | Get a token using an OAuth2 flow



## GetToken

> OAuthTokenResponse GetToken(ctx).GrantType(grantType).Scope(scope).ClientId(clientId).ClientSecret(clientSecret).RequestedTokenType(requestedTokenType).SubjectToken(subjectToken).SubjectTokenType(subjectTokenType).ActorToken(actorToken).ActorTokenType(actorTokenType).Execute()

Get a token using an OAuth2 flow



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/confluentinc/ccloud-sdk-go-v2"
)

func main() {
	grantType := "grantType_example" // string |  (optional)
	scope := "scope_example" // string |  (optional)
	clientId := "clientId_example" // string | Client ID  This can be sent in the request body, but OAuth2 recommends sending it in a Basic Authorization header. (optional)
	clientSecret := "clientSecret_example" // string | Client secret  This can be sent in the request body, but OAuth2 recommends sending it in a Basic Authorization header. (optional)
	requestedTokenType := openapiclient.TokenType("urn:ietf:params:oauth:token-type:access_token") // TokenType |  (optional)
	subjectToken := "subjectToken_example" // string | Subject token for token exchange request (optional)
	subjectTokenType := openapiclient.TokenType("urn:ietf:params:oauth:token-type:access_token") // TokenType |  (optional)
	actorToken := "actorToken_example" // string | Actor token for token exchange request (optional)
	actorTokenType := openapiclient.TokenType("urn:ietf:params:oauth:token-type:access_token") // TokenType |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth2APIAPI.GetToken(context.Background()).GrantType(grantType).Scope(scope).ClientId(clientId).ClientSecret(clientSecret).RequestedTokenType(requestedTokenType).SubjectToken(subjectToken).SubjectTokenType(subjectTokenType).ActorToken(actorToken).ActorTokenType(actorTokenType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth2APIAPI.GetToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToken`: OAuthTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuth2APIAPI.GetToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** |  | 
 **scope** | **string** |  | 
 **clientId** | **string** | Client ID  This can be sent in the request body, but OAuth2 recommends sending it in a Basic Authorization header. | 
 **clientSecret** | **string** | Client secret  This can be sent in the request body, but OAuth2 recommends sending it in a Basic Authorization header. | 
 **requestedTokenType** | [**TokenType**](TokenType.md) |  | 
 **subjectToken** | **string** | Subject token for token exchange request | 
 **subjectTokenType** | [**TokenType**](TokenType.md) |  | 
 **actorToken** | **string** | Actor token for token exchange request | 
 **actorTokenType** | [**TokenType**](TokenType.md) |  | 

### Return type

[**OAuthTokenResponse**](OAuthTokenResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

