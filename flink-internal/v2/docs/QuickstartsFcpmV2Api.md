# \QuickstartsFcpmV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFcpmV2Quickstart**](QuickstartsFcpmV2Api.md#CreateFcpmV2Quickstart) | **Post** /fcpm/v2/quickstart | quickstart compute pools



## CreateFcpmV2Quickstart

> FcpmV2Quickstart CreateFcpmV2Quickstart(ctx).FcpmV2Quickstart(fcpmV2Quickstart).Execute()

quickstart compute pools



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
    fcpmV2Quickstart := *openapiclient.NewFcpmV2Quickstart() // FcpmV2Quickstart |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuickstartsFcpmV2Api.CreateFcpmV2Quickstart(context.Background()).FcpmV2Quickstart(fcpmV2Quickstart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuickstartsFcpmV2Api.CreateFcpmV2Quickstart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFcpmV2Quickstart`: FcpmV2Quickstart
    fmt.Fprintf(os.Stdout, "Response from `QuickstartsFcpmV2Api.CreateFcpmV2Quickstart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFcpmV2QuickstartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fcpmV2Quickstart** | [**FcpmV2Quickstart**](FcpmV2Quickstart.md) |  | 

### Return type

[**FcpmV2Quickstart**](fcpm.v2.Quickstart.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

