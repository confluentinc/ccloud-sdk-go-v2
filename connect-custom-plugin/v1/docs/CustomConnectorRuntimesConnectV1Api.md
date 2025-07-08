# \CustomConnectorRuntimesConnectV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListConnectV1CustomConnectorRuntimes**](CustomConnectorRuntimesConnectV1Api.md#ListConnectV1CustomConnectorRuntimes) | **Get** /connect/v1/custom-connector-runtimes | List of Custom Connector Runtimes



## ListConnectV1CustomConnectorRuntimes

> ConnectV1CustomConnectorRuntimeList ListConnectV1CustomConnectorRuntimes(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Custom Connector Runtimes



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
    resp, r, err := api_client.CustomConnectorRuntimesConnectV1Api.ListConnectV1CustomConnectorRuntimes(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomConnectorRuntimesConnectV1Api.ListConnectV1CustomConnectorRuntimes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectV1CustomConnectorRuntimes`: ConnectV1CustomConnectorRuntimeList
    fmt.Fprintf(os.Stdout, "Response from `CustomConnectorRuntimesConnectV1Api.ListConnectV1CustomConnectorRuntimes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectV1CustomConnectorRuntimesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**ConnectV1CustomConnectorRuntimeList**](connect.v1.CustomConnectorRuntimeList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

