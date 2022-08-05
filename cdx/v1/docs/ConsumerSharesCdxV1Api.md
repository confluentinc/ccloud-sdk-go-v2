# \ConsumerSharesCdxV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCdxV1ConsumerShare**](ConsumerSharesCdxV1Api.md#DeleteCdxV1ConsumerShare) | **Delete** /cdx/v1/consumer-shares/{id} | Delete a Consumer Share
[**GetCdxV1ConsumerShare**](ConsumerSharesCdxV1Api.md#GetCdxV1ConsumerShare) | **Get** /cdx/v1/consumer-shares/{id} | Read a Consumer Share
[**ListCdxV1ConsumerShares**](ConsumerSharesCdxV1Api.md#ListCdxV1ConsumerShares) | **Get** /cdx/v1/consumer-shares | List of Consumer Shares



## DeleteCdxV1ConsumerShare

> DeleteCdxV1ConsumerShare(ctx, id).Execute()

Delete a Consumer Share



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
    id := "id_example" // string | The unique identifier for the consumer share.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConsumerSharesCdxV1Api.DeleteCdxV1ConsumerShare(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsumerSharesCdxV1Api.DeleteCdxV1ConsumerShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the consumer share. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCdxV1ConsumerShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCdxV1ConsumerShare

> CdxV1ConsumerShare GetCdxV1ConsumerShare(ctx, id).Execute()

Read a Consumer Share



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
    id := "id_example" // string | The unique identifier for the consumer share.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConsumerSharesCdxV1Api.GetCdxV1ConsumerShare(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsumerSharesCdxV1Api.GetCdxV1ConsumerShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCdxV1ConsumerShare`: CdxV1ConsumerShare
    fmt.Fprintf(os.Stdout, "Response from `ConsumerSharesCdxV1Api.GetCdxV1ConsumerShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the consumer share. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCdxV1ConsumerShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CdxV1ConsumerShare**](cdx.v1.ConsumerShare.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCdxV1ConsumerShares

> CdxV1ConsumerShareList ListCdxV1ConsumerShares(ctx).SharedResource(sharedResource).PageSize(pageSize).PageToken(pageToken).Execute()

List of Consumer Shares



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
    sharedResource := "sharedResource_example" // string | Filter the results by exact match for shared_resource. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConsumerSharesCdxV1Api.ListCdxV1ConsumerShares(context.Background()).SharedResource(sharedResource).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsumerSharesCdxV1Api.ListCdxV1ConsumerShares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCdxV1ConsumerShares`: CdxV1ConsumerShareList
    fmt.Fprintf(os.Stdout, "Response from `ConsumerSharesCdxV1Api.ListCdxV1ConsumerShares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCdxV1ConsumerSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sharedResource** | **string** | Filter the results by exact match for shared_resource. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**CdxV1ConsumerShareList**](cdx.v1.ConsumerShareList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

