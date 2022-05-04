# \UsagesCliV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCliV1Usage**](UsagesCliV1Api.md#CreateCliV1Usage) | **Post** /cli/v1/usages | Create a Usage



## CreateCliV1Usage

> CreateCliV1Usage(ctx).CliV1Usage(cliV1Usage).Execute()

Create a Usage



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
    cliV1Usage := *openapiclient.NewCliV1Usage() // CliV1Usage |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsagesCliV1Api.CreateCliV1Usage(context.Background()).CliV1Usage(cliV1Usage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsagesCliV1Api.CreateCliV1Usage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCliV1UsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cliV1Usage** | [**CliV1Usage**](CliV1Usage.md) |  | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

