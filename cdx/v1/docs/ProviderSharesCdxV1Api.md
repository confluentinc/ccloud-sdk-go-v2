# \ProviderSharesCdxV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCdxV1ProviderShare**](ProviderSharesCdxV1Api.md#CreateCdxV1ProviderShare) | **Post** /cdx/v1/provider-shares | Create a share
[**DeleteCdxV1ProviderShare**](ProviderSharesCdxV1Api.md#DeleteCdxV1ProviderShare) | **Delete** /cdx/v1/provider-shares/{id} | Delete a Provider Share
[**GetCdxV1ProviderShare**](ProviderSharesCdxV1Api.md#GetCdxV1ProviderShare) | **Get** /cdx/v1/provider-shares/{id} | Read a Provider Share
[**ListCdxV1ProviderShares**](ProviderSharesCdxV1Api.md#ListCdxV1ProviderShares) | **Get** /cdx/v1/provider-shares | List of Provider Shares
[**ResendCdxV1ProviderShare**](ProviderSharesCdxV1Api.md#ResendCdxV1ProviderShare) | **Post** /cdx/v1/provider-shares/{id}:resend | Resend



## CreateCdxV1ProviderShare

> CdxV1ProviderShare CreateCdxV1ProviderShare(ctx).CdxV1CreateShareRequest(cdxV1CreateShareRequest).Execute()

Create a share



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
    cdxV1CreateShareRequest := *openapiclient.NewCdxV1CreateShareRequest() // CdxV1CreateShareRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderSharesCdxV1Api.CreateCdxV1ProviderShare(context.Background()).CdxV1CreateShareRequest(cdxV1CreateShareRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderSharesCdxV1Api.CreateCdxV1ProviderShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCdxV1ProviderShare`: CdxV1ProviderShare
    fmt.Fprintf(os.Stdout, "Response from `ProviderSharesCdxV1Api.CreateCdxV1ProviderShare`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCdxV1ProviderShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cdxV1CreateShareRequest** | [**CdxV1CreateShareRequest**](CdxV1CreateShareRequest.md) |  | 

### Return type

[**CdxV1ProviderShare**](CdxV1ProviderShare.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCdxV1ProviderShare

> DeleteCdxV1ProviderShare(ctx, id).Execute()

Delete a Provider Share



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
    id := "id_example" // string | The unique identifier for the provider share.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderSharesCdxV1Api.DeleteCdxV1ProviderShare(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderSharesCdxV1Api.DeleteCdxV1ProviderShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the provider share. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCdxV1ProviderShareRequest struct via the builder pattern


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


## GetCdxV1ProviderShare

> CdxV1ProviderShare GetCdxV1ProviderShare(ctx, id).Execute()

Read a Provider Share



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
    id := "id_example" // string | The unique identifier for the provider share.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderSharesCdxV1Api.GetCdxV1ProviderShare(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderSharesCdxV1Api.GetCdxV1ProviderShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCdxV1ProviderShare`: CdxV1ProviderShare
    fmt.Fprintf(os.Stdout, "Response from `ProviderSharesCdxV1Api.GetCdxV1ProviderShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the provider share. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCdxV1ProviderShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CdxV1ProviderShare**](cdx.v1.ProviderShare.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCdxV1ProviderShares

> CdxV1ProviderShareList ListCdxV1ProviderShares(ctx).SharedResource(sharedResource).PageSize(pageSize).PageToken(pageToken).Execute()

List of Provider Shares



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
    resp, r, err := api_client.ProviderSharesCdxV1Api.ListCdxV1ProviderShares(context.Background()).SharedResource(sharedResource).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderSharesCdxV1Api.ListCdxV1ProviderShares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCdxV1ProviderShares`: CdxV1ProviderShareList
    fmt.Fprintf(os.Stdout, "Response from `ProviderSharesCdxV1Api.ListCdxV1ProviderShares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCdxV1ProviderSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sharedResource** | **string** | Filter the results by exact match for shared_resource. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**CdxV1ProviderShareList**](cdx.v1.ProviderShareList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendCdxV1ProviderShare

> ResendCdxV1ProviderShare(ctx, id).Execute()

Resend



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
    id := "id_example" // string | The unique identifier for the provider share.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderSharesCdxV1Api.ResendCdxV1ProviderShare(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderSharesCdxV1Api.ResendCdxV1ProviderShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the provider share. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendCdxV1ProviderShareRequest struct via the builder pattern


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

