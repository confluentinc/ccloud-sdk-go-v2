# \ScimTokensOrgV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgV2ScimToken**](ScimTokensOrgV2Api.md#CreateOrgV2ScimToken) | **Post** /org/v2/scim-tokens | Create a Scim Token
[**DeleteOrgV2ScimToken**](ScimTokensOrgV2Api.md#DeleteOrgV2ScimToken) | **Delete** /org/v2/scim-tokens/{id} | Delete a Scim Token
[**ListOrgV2ScimTokens**](ScimTokensOrgV2Api.md#ListOrgV2ScimTokens) | **Get** /org/v2/scim-tokens | List of Scim Tokens



## CreateOrgV2ScimToken

> OrgV2ScimToken CreateOrgV2ScimToken(ctx).InlineObject(inlineObject).Execute()

Create a Scim Token



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
    inlineObject := *openapiclient.NewInlineObject() // InlineObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScimTokensOrgV2Api.CreateOrgV2ScimToken(context.Background()).InlineObject(inlineObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimTokensOrgV2Api.CreateOrgV2ScimToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrgV2ScimToken`: OrgV2ScimToken
    fmt.Fprintf(os.Stdout, "Response from `ScimTokensOrgV2Api.CreateOrgV2ScimToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgV2ScimTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject** | [**InlineObject**](InlineObject.md) |  | 

### Return type

[**OrgV2ScimToken**](OrgV2ScimToken.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgV2ScimToken

> DeleteOrgV2ScimToken(ctx, id).Execute()

Delete a Scim Token



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
    id := "id_example" // string | The unique identifier for the scim token.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScimTokensOrgV2Api.DeleteOrgV2ScimToken(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimTokensOrgV2Api.DeleteOrgV2ScimToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the scim token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgV2ScimTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgV2ScimTokens

> OrgV2ScimTokenList ListOrgV2ScimTokens(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Scim Tokens



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
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScimTokensOrgV2Api.ListOrgV2ScimTokens(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimTokensOrgV2Api.ListOrgV2ScimTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgV2ScimTokens`: OrgV2ScimTokenList
    fmt.Fprintf(os.Stdout, "Response from `ScimTokensOrgV2Api.ListOrgV2ScimTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrgV2ScimTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**OrgV2ScimTokenList**](org.v2.ScimTokenList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

