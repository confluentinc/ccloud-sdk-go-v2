# \SsoConnectionsV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateV2SsoConnection**](SsoConnectionsV2Api.md#CreateV2SsoConnection) | **Post** /v2/sso-connections | Create a Sso Connection



## CreateV2SsoConnection

> V2SsoConnection CreateV2SsoConnection(ctx).V2SsoConnection(v2SsoConnection).Execute()

Create a Sso Connection



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
    v2SsoConnection := *openapiclient.NewV2SsoConnection() // V2SsoConnection |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SsoConnectionsV2Api.CreateV2SsoConnection(context.Background()).V2SsoConnection(v2SsoConnection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsoConnectionsV2Api.CreateV2SsoConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateV2SsoConnection`: V2SsoConnection
    fmt.Fprintf(os.Stdout, "Response from `SsoConnectionsV2Api.CreateV2SsoConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateV2SsoConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2SsoConnection** | [**V2SsoConnection**](V2SsoConnection.md) |  | 

### Return type

[**V2SsoConnection**](v2.SsoConnection.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

