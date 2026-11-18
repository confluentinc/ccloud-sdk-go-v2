# \LogsLoggingV1Api

All URIs are relative to *https://api.logging.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DescribeLoggingV1Source**](LogsLoggingV1Api.md#DescribeLoggingV1Source) | **Get** /logs/v1/sources/{name} | Describe one log source
[**ListLoggingV1Sources**](LogsLoggingV1Api.md#ListLoggingV1Sources) | **Get** /logs/v1/sources | List available log sources
[**SearchLoggingV1Logs**](LogsLoggingV1Api.md#SearchLoggingV1Logs) | **Post** /logs/v1/search | Search logs



## DescribeLoggingV1Source

> LoggingV1Source DescribeLoggingV1Source(ctx, name).Execute()

Describe one log source



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
    name := "name_example" // string | The source's catalog key, for example `connector` or `statement`.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogsLoggingV1Api.DescribeLoggingV1Source(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsLoggingV1Api.DescribeLoggingV1Source``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeLoggingV1Source`: LoggingV1Source
    fmt.Fprintf(os.Stdout, "Response from `LogsLoggingV1Api.DescribeLoggingV1Source`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The source&#39;s catalog key, for example &#x60;connector&#x60; or &#x60;statement&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeLoggingV1SourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoggingV1Source**](LoggingV1Source.md)

### Authorization

[confluent_auth](../README.md#confluent_auth), [data-plane-bearer-token](../README.md#data-plane-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLoggingV1Sources

> LoggingV1SourceList ListLoggingV1Sources(ctx).Execute()

List available log sources



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogsLoggingV1Api.ListLoggingV1Sources(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsLoggingV1Api.ListLoggingV1Sources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLoggingV1Sources`: LoggingV1SourceList
    fmt.Fprintf(os.Stdout, "Response from `LogsLoggingV1Api.ListLoggingV1Sources`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLoggingV1SourcesRequest struct via the builder pattern


### Return type

[**LoggingV1SourceList**](LoggingV1SourceList.md)

### Authorization

[confluent_auth](../README.md#confluent_auth), [data-plane-bearer-token](../README.md#data-plane-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchLoggingV1Logs

> LoggingV1SearchResponse SearchLoggingV1Logs(ctx).LoggingV1SearchRequest(loggingV1SearchRequest).PageSize(pageSize).PageToken(pageToken).Execute()

Search logs



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
    loggingV1SearchRequest := *openapiclient.NewLoggingV1SearchRequest("crn://confluent.cloud/organization=9bb441c4-.../environment=env-123/cloud-cluster=lkc-abc123/connector=my-connector", "2026-08-20T00:00:00Z", "2026-08-21T00:00:00Z") // LoggingV1SearchRequest | 
    pageSize := int32(56) // int32 | Number of documents per page. Values outside 1..1000 are rejected with 422. (optional) (default to 100)
    pageToken := "pageToken_example" // string | Opaque continuation token taken from `metadata.next` of the previous response. The page size encoded in the token takes precedence over `page_size`. A token that cannot be decoded is ignored and the first page is returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogsLoggingV1Api.SearchLoggingV1Logs(context.Background()).LoggingV1SearchRequest(loggingV1SearchRequest).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsLoggingV1Api.SearchLoggingV1Logs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchLoggingV1Logs`: LoggingV1SearchResponse
    fmt.Fprintf(os.Stdout, "Response from `LogsLoggingV1Api.SearchLoggingV1Logs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchLoggingV1LogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loggingV1SearchRequest** | [**LoggingV1SearchRequest**](LoggingV1SearchRequest.md) |  | 
 **pageSize** | **int32** | Number of documents per page. Values outside 1..1000 are rejected with 422. | [default to 100]
 **pageToken** | **string** | Opaque continuation token taken from &#x60;metadata.next&#x60; of the previous response. The page size encoded in the token takes precedence over &#x60;page_size&#x60;. A token that cannot be decoded is ignored and the first page is returned. | 

### Return type

[**LoggingV1SearchResponse**](LoggingV1SearchResponse.md)

### Authorization

[confluent_auth](../README.md#confluent_auth), [data-plane-bearer-token](../README.md#data-plane-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

