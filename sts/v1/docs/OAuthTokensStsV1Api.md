# \OAuthTokensStsV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExchangeStsV1OauthToken**](OAuthTokensStsV1Api.md#ExchangeStsV1OauthToken) | **Post** /sts/v1/oauth2/token | Exchange an OAuth Token



## ExchangeStsV1OauthToken

> StsV1TokenExchangeReply ExchangeStsV1OauthToken(ctx).ApiVersion(apiVersion).Kind(kind).Id(id).Metadata(metadata).GrantType(grantType).SubjectToken(subjectToken).IdentityPoolId(identityPoolId).SubjectTokenType(subjectTokenType).RequestedTokenType(requestedTokenType).ExpiresIn(expiresIn).Execute()

Exchange an OAuth Token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    apiVersion := "apiVersion_example" // string | APIVersion defines the schema version of this representation of a resource. (optional)
    kind := "kind_example" // string | Kind defines the object this REST resource represents. (optional)
    id := "id_example" // string | ID is the \\\"natural identifier\\\" for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\\\"time\\\"); however, it may collide with IDs for other object `kinds` or objects of the same `kind` within a different scope/namespace (\\\"space\\\"). (optional)
    metadata := TODO // ObjectMeta |  (optional)
    grantType := "grantType_example" // string | The grant type. Must be urn:ietf:params:oauth:grant-type:token-exchange, which indicates a token exchange.  (optional)
    subjectToken := "subjectToken_example" // string | Confluent Cloud only accepts JSON Web Token (JWT) access tokens from customer identity provider (optional)
    identityPoolId := "identityPoolId_example" // string | Identity pool is a group of external identities that are assigned a certain level of access based on policy  (optional)
    subjectTokenType := "subjectTokenType_example" // string | An identifier for the type of requested security token. Supported values is urn:ietf:params:oauth:token-type:jwt.  (optional)
    requestedTokenType := "requestedTokenType_example" // string | An identifier for the type of requested security token. Supported values is urn:ietf:params:oauth:token-type:access_token.  (optional)
    expiresIn := int32(56) // int32 | The amount of time, in seconds, between the time when the access token was issued and the time when the access token will expire  (optional) (default to 900)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OAuthTokensStsV1Api.ExchangeStsV1OauthToken(context.Background()).ApiVersion(apiVersion).Kind(kind).Id(id).Metadata(metadata).GrantType(grantType).SubjectToken(subjectToken).IdentityPoolId(identityPoolId).SubjectTokenType(subjectTokenType).RequestedTokenType(requestedTokenType).ExpiresIn(expiresIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthTokensStsV1Api.ExchangeStsV1OauthToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExchangeStsV1OauthToken`: StsV1TokenExchangeReply
    fmt.Fprintf(os.Stdout, "Response from `OAuthTokensStsV1Api.ExchangeStsV1OauthToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExchangeStsV1OauthTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | 
 **kind** | **string** | Kind defines the object this REST resource represents. | 
 **id** | **string** | ID is the \\\&quot;natural identifier\\\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\\\&quot;time\\\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\\\&quot;space\\\&quot;). | 
 **metadata** | [**ObjectMeta**](ObjectMeta.md) |  | 
 **grantType** | **string** | The grant type. Must be urn:ietf:params:oauth:grant-type:token-exchange, which indicates a token exchange.  | 
 **subjectToken** | **string** | Confluent Cloud only accepts JSON Web Token (JWT) access tokens from customer identity provider | 
 **identityPoolId** | **string** | Identity pool is a group of external identities that are assigned a certain level of access based on policy  | 
 **subjectTokenType** | **string** | An identifier for the type of requested security token. Supported values is urn:ietf:params:oauth:token-type:jwt.  | 
 **requestedTokenType** | **string** | An identifier for the type of requested security token. Supported values is urn:ietf:params:oauth:token-type:access_token.  | 
 **expiresIn** | **int32** | The amount of time, in seconds, between the time when the access token was issued and the time when the access token will expire  | [default to 900]

### Return type

[**StsV1TokenExchangeReply**](StsV1TokenExchangeReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

