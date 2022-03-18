# \UsagesCliUsageV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCliUsageV1Usage**](UsagesCliUsageV1Api.md#CreateCliUsageV1Usage) | **Post** /cli-usage/v1/usages | Create a Usage
[**GetCliUsageV1Usage**](UsagesCliUsageV1Api.md#GetCliUsageV1Usage) | **Get** /cli-usage/v1/usages/{id} | Read a Usage



## CreateCliUsageV1Usage

> CliUsageV1Usage CreateCliUsageV1Usage(ctx).CliUsageV1Usage(cliUsageV1Usage).Execute()

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
    cliUsageV1Usage := *openapiclient.NewCliUsageV1Usage() // CliUsageV1Usage |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsagesCliUsageV1Api.CreateCliUsageV1Usage(context.Background()).CliUsageV1Usage(cliUsageV1Usage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsagesCliUsageV1Api.CreateCliUsageV1Usage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCliUsageV1Usage`: CliUsageV1Usage
    fmt.Fprintf(os.Stdout, "Response from `UsagesCliUsageV1Api.CreateCliUsageV1Usage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCliUsageV1UsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cliUsageV1Usage** | [**CliUsageV1Usage**](CliUsageV1Usage.md) |  | 

### Return type

[**CliUsageV1Usage**](cli-usage.v1.Usage.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCliUsageV1Usage

> CliUsageV1Usage GetCliUsageV1Usage(ctx, id).Execute()

Read a Usage



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
    id := "id_example" // string | The unique identifier for the usage.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsagesCliUsageV1Api.GetCliUsageV1Usage(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsagesCliUsageV1Api.GetCliUsageV1Usage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCliUsageV1Usage`: CliUsageV1Usage
    fmt.Fprintf(os.Stdout, "Response from `UsagesCliUsageV1Api.GetCliUsageV1Usage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the usage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCliUsageV1UsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CliUsageV1Usage**](cli-usage.v1.Usage.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

