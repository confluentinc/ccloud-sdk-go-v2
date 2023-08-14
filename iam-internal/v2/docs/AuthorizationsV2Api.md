# \AuthorizationsV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizeV2Authorization**](AuthorizationsV2Api.md#AuthorizeV2Authorization) | **Post** /v2/authorize/{id}/authorize | Authorize an Authorization



## AuthorizeV2Authorization

> []string AuthorizeV2Authorization(ctx, id).IamV2AuthorizeRequest(iamV2AuthorizeRequest).Execute()

Authorize an Authorization



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
    id := "id_example" // string | The unique identifier for the authorization.
    iamV2AuthorizeRequest := *openapiclient.NewIamV2AuthorizeRequest() // IamV2AuthorizeRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthorizationsV2Api.AuthorizeV2Authorization(context.Background(), id).IamV2AuthorizeRequest(iamV2AuthorizeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationsV2Api.AuthorizeV2Authorization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthorizeV2Authorization`: []string
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationsV2Api.AuthorizeV2Authorization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the authorization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeV2AuthorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamV2AuthorizeRequest** | [**IamV2AuthorizeRequest**](IamV2AuthorizeRequest.md) |  | 

### Return type

**[]string**

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

