# \ScopesServiceQuotaV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceQuotaV2Scope**](ScopesServiceQuotaV2Api.md#GetServiceQuotaV2Scope) | **Get** /service-quota/v2/scopes/{id} | Read a Scope
[**ListServiceQuotaV2Scopes**](ScopesServiceQuotaV2Api.md#ListServiceQuotaV2Scopes) | **Get** /service-quota/v2/scopes | List of Scopes



## GetServiceQuotaV2Scope

> ServiceQuotaV2Scope GetServiceQuotaV2Scope(ctx, id).Execute()

Read a Scope



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
    id := "id_example" // string | The unique identifier for the scope.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScopesServiceQuotaV2Api.GetServiceQuotaV2Scope(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScopesServiceQuotaV2Api.GetServiceQuotaV2Scope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceQuotaV2Scope`: ServiceQuotaV2Scope
    fmt.Fprintf(os.Stdout, "Response from `ScopesServiceQuotaV2Api.GetServiceQuotaV2Scope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceQuotaV2ScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceQuotaV2Scope**](service-quota.v2.Scope.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceQuotaV2Scopes

> ServiceQuotaV2ScopeList ListServiceQuotaV2Scopes(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Scopes



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
    resp, r, err := api_client.ScopesServiceQuotaV2Api.ListServiceQuotaV2Scopes(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScopesServiceQuotaV2Api.ListServiceQuotaV2Scopes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceQuotaV2Scopes`: ServiceQuotaV2ScopeList
    fmt.Fprintf(os.Stdout, "Response from `ScopesServiceQuotaV2Api.ListServiceQuotaV2Scopes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServiceQuotaV2ScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**ServiceQuotaV2ScopeList**](service-quota.v2.ScopeList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

