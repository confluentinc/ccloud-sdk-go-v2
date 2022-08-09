# \ConsumerSharedResourcesCdxV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCdxV1ConsumerSharedResource**](ConsumerSharedResourcesCdxV1Api.md#GetCdxV1ConsumerSharedResource) | **Get** /cdx/v1/consumer-shared-resources/{id} | Read a Consumer Shared Resource
[**ImageCdxV1ConsumerSharedResource**](ConsumerSharedResourcesCdxV1Api.md#ImageCdxV1ConsumerSharedResource) | **Get** /cdx/v1/consumer-shared-resources/{id}/images/{file_name} | Image a Consumer Shared Resource
[**ListCdxV1ConsumerSharedResources**](ConsumerSharedResourcesCdxV1Api.md#ListCdxV1ConsumerSharedResources) | **Get** /cdx/v1/consumer-shared-resources | List of Consumer Shared Resources
[**NetworkCdxV1ConsumerSharedResource**](ConsumerSharedResourcesCdxV1Api.md#NetworkCdxV1ConsumerSharedResource) | **Get** /cdx/v1/consumer-shared-resources/{id}:network | Network a Consumer Shared Resource



## GetCdxV1ConsumerSharedResource

> CdxV1ConsumerSharedResource GetCdxV1ConsumerSharedResource(ctx, id).Execute()

Read a Consumer Shared Resource



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
    id := "id_example" // string | The unique identifier for the consumer shared resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConsumerSharedResourcesCdxV1Api.GetCdxV1ConsumerSharedResource(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsumerSharedResourcesCdxV1Api.GetCdxV1ConsumerSharedResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCdxV1ConsumerSharedResource`: CdxV1ConsumerSharedResource
    fmt.Fprintf(os.Stdout, "Response from `ConsumerSharedResourcesCdxV1Api.GetCdxV1ConsumerSharedResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the consumer shared resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCdxV1ConsumerSharedResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CdxV1ConsumerSharedResource**](cdx.v1.ConsumerSharedResource.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageCdxV1ConsumerSharedResource

> *os.File ImageCdxV1ConsumerSharedResource(ctx, id, fileName).Execute()

Image a Consumer Shared Resource



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
    id := "id_example" // string | The unique identifier for the consumer shared resource.
    fileName := "fileName_example" // string | The File Name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConsumerSharedResourcesCdxV1Api.ImageCdxV1ConsumerSharedResource(context.Background(), id, fileName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsumerSharedResourcesCdxV1Api.ImageCdxV1ConsumerSharedResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImageCdxV1ConsumerSharedResource`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ConsumerSharedResourcesCdxV1Api.ImageCdxV1ConsumerSharedResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the consumer shared resource. | 
**fileName** | **string** | The File Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageCdxV1ConsumerSharedResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCdxV1ConsumerSharedResources

> CdxV1ConsumerSharedResourceList ListCdxV1ConsumerSharedResources(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Consumer Shared Resources



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
    resp, r, err := api_client.ConsumerSharedResourcesCdxV1Api.ListCdxV1ConsumerSharedResources(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsumerSharedResourcesCdxV1Api.ListCdxV1ConsumerSharedResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCdxV1ConsumerSharedResources`: CdxV1ConsumerSharedResourceList
    fmt.Fprintf(os.Stdout, "Response from `ConsumerSharedResourcesCdxV1Api.ListCdxV1ConsumerSharedResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCdxV1ConsumerSharedResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**CdxV1ConsumerSharedResourceList**](cdx.v1.ConsumerSharedResourceList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkCdxV1ConsumerSharedResource

> CdxV1Network NetworkCdxV1ConsumerSharedResource(ctx, id).Execute()

Network a Consumer Shared Resource



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
    id := "id_example" // string | The unique identifier for the consumer shared resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConsumerSharedResourcesCdxV1Api.NetworkCdxV1ConsumerSharedResource(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsumerSharedResourcesCdxV1Api.NetworkCdxV1ConsumerSharedResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkCdxV1ConsumerSharedResource`: CdxV1Network
    fmt.Fprintf(os.Stdout, "Response from `ConsumerSharedResourcesCdxV1Api.NetworkCdxV1ConsumerSharedResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the consumer shared resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkCdxV1ConsumerSharedResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CdxV1Network**](CdxV1Network.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

