# \AccessesSrcmV3Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSrcmV3Access**](AccessesSrcmV3Api.md#GetSrcmV3Access) | **Get** /srcm/v3/clusters/{id}/access | Read an Access



## GetSrcmV3Access

> SrcmV3Access GetSrcmV3Access(ctx, id).Environment(environment).ClusterId(clusterId).Execute()

Read an Access



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
    id := "id_example" // string | The unique identifier for the access.
    environment := "env-00000" // string | Filter the results by exact match for environment.
    clusterId := "lkc-1234" // string | Logical cluster that will access schema registry.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessesSrcmV3Api.GetSrcmV3Access(context.Background(), id).Environment(environment).ClusterId(clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessesSrcmV3Api.GetSrcmV3Access``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSrcmV3Access`: SrcmV3Access
    fmt.Fprintf(os.Stdout, "Response from `AccessesSrcmV3Api.GetSrcmV3Access`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSrcmV3AccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environment** | **string** | Filter the results by exact match for environment. | 
 **clusterId** | **string** | Logical cluster that will access schema registry. | 

### Return type

[**SrcmV3Access**](srcm.v3.Access.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

