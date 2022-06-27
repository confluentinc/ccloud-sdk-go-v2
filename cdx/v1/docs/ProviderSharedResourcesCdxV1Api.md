# \ProviderSharedResourcesCdxV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteImageCdxV1ProviderSharedResource**](ProviderSharedResourcesCdxV1Api.md#DeleteImageCdxV1ProviderSharedResource) | **Delete** /cdx/v1/provider-shared-resources/{id}/images/{file_name} | Delete_Image a Provider Shared Resource
[**GetCdxV1ProviderSharedResource**](ProviderSharedResourcesCdxV1Api.md#GetCdxV1ProviderSharedResource) | **Get** /cdx/v1/provider-shared-resources/{id} | Read a Provider Shared Resource
[**ListCdxV1ProviderSharedResources**](ProviderSharedResourcesCdxV1Api.md#ListCdxV1ProviderSharedResources) | **Get** /cdx/v1/provider-shared-resources | List of Provider Shared Resources
[**UpdateCdxV1ProviderSharedResource**](ProviderSharedResourcesCdxV1Api.md#UpdateCdxV1ProviderSharedResource) | **Patch** /cdx/v1/provider-shared-resources/{id} | Update a Provider Shared Resource
[**UploadImageCdxV1ProviderSharedResource**](ProviderSharedResourcesCdxV1Api.md#UploadImageCdxV1ProviderSharedResource) | **Post** /cdx/v1/provider-shared-resources/{id}/images/{file_name} | Upload_Image a Provider Shared Resource
[**ViewImageCdxV1ProviderSharedResource**](ProviderSharedResourcesCdxV1Api.md#ViewImageCdxV1ProviderSharedResource) | **Get** /cdx/v1/provider-shared-resources/{id}/images/{file_name} | View_Image a Provider Shared Resource



## DeleteImageCdxV1ProviderSharedResource

> DeleteImageCdxV1ProviderSharedResource(ctx, id, fileName).Execute()

Delete_Image a Provider Shared Resource



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
    id := "id_example" // string | The unique identifier for the provider shared resource.
    fileName := "fileName_example" // string | The File Name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderSharedResourcesCdxV1Api.DeleteImageCdxV1ProviderSharedResource(context.Background(), id, fileName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderSharedResourcesCdxV1Api.DeleteImageCdxV1ProviderSharedResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the provider shared resource. | 
**fileName** | **string** | The File Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImageCdxV1ProviderSharedResourceRequest struct via the builder pattern


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


## GetCdxV1ProviderSharedResource

> CdxV1ProviderSharedResource GetCdxV1ProviderSharedResource(ctx, id).Execute()

Read a Provider Shared Resource



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
    id := "id_example" // string | The unique identifier for the provider shared resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderSharedResourcesCdxV1Api.GetCdxV1ProviderSharedResource(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderSharedResourcesCdxV1Api.GetCdxV1ProviderSharedResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCdxV1ProviderSharedResource`: CdxV1ProviderSharedResource
    fmt.Fprintf(os.Stdout, "Response from `ProviderSharedResourcesCdxV1Api.GetCdxV1ProviderSharedResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the provider shared resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCdxV1ProviderSharedResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CdxV1ProviderSharedResource**](cdx.v1.ProviderSharedResource.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCdxV1ProviderSharedResources

> CdxV1ProviderSharedResourceList ListCdxV1ProviderSharedResources(ctx).Crn(crn).PageSize(pageSize).PageToken(pageToken).Execute()

List of Provider Shared Resources



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
    crn := "crn://confluent.cloud/cloud-cluster=lkc-1111aaa/kafka=lkc-111aaa/topic=my.topic" // string | Filter the results by exact match for crn. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderSharedResourcesCdxV1Api.ListCdxV1ProviderSharedResources(context.Background()).Crn(crn).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderSharedResourcesCdxV1Api.ListCdxV1ProviderSharedResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCdxV1ProviderSharedResources`: CdxV1ProviderSharedResourceList
    fmt.Fprintf(os.Stdout, "Response from `ProviderSharedResourcesCdxV1Api.ListCdxV1ProviderSharedResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCdxV1ProviderSharedResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **crn** | **string** | Filter the results by exact match for crn. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**CdxV1ProviderSharedResourceList**](cdx.v1.ProviderSharedResourceList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCdxV1ProviderSharedResource

> CdxV1ProviderSharedResource UpdateCdxV1ProviderSharedResource(ctx, id).CdxV1ProviderSharedResourceUpdate(cdxV1ProviderSharedResourceUpdate).Execute()

Update a Provider Shared Resource



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
    id := "id_example" // string | The unique identifier for the provider shared resource.
    cdxV1ProviderSharedResourceUpdate := *openapiclient.NewCdxV1ProviderSharedResourceUpdate() // CdxV1ProviderSharedResourceUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderSharedResourcesCdxV1Api.UpdateCdxV1ProviderSharedResource(context.Background(), id).CdxV1ProviderSharedResourceUpdate(cdxV1ProviderSharedResourceUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderSharedResourcesCdxV1Api.UpdateCdxV1ProviderSharedResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCdxV1ProviderSharedResource`: CdxV1ProviderSharedResource
    fmt.Fprintf(os.Stdout, "Response from `ProviderSharedResourcesCdxV1Api.UpdateCdxV1ProviderSharedResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the provider shared resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCdxV1ProviderSharedResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cdxV1ProviderSharedResourceUpdate** | [**CdxV1ProviderSharedResourceUpdate**](CdxV1ProviderSharedResourceUpdate.md) |  | 

### Return type

[**CdxV1ProviderSharedResource**](cdx.v1.ProviderSharedResource.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadImageCdxV1ProviderSharedResource

> UploadImageCdxV1ProviderSharedResource(ctx, id, fileName).File(file).Execute()

Upload_Image a Provider Shared Resource



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
    id := "id_example" // string | The unique identifier for the provider shared resource.
    fileName := "fileName_example" // string | The File Name
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderSharedResourcesCdxV1Api.UploadImageCdxV1ProviderSharedResource(context.Background(), id, fileName).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderSharedResourcesCdxV1Api.UploadImageCdxV1ProviderSharedResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the provider shared resource. | 
**fileName** | **string** | The File Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadImageCdxV1ProviderSharedResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewImageCdxV1ProviderSharedResource

> *os.File ViewImageCdxV1ProviderSharedResource(ctx, id, fileName).Execute()

View_Image a Provider Shared Resource



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
    id := "id_example" // string | The unique identifier for the provider shared resource.
    fileName := "fileName_example" // string | The File Name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderSharedResourcesCdxV1Api.ViewImageCdxV1ProviderSharedResource(context.Background(), id, fileName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderSharedResourcesCdxV1Api.ViewImageCdxV1ProviderSharedResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ViewImageCdxV1ProviderSharedResource`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ProviderSharedResourcesCdxV1Api.ViewImageCdxV1ProviderSharedResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the provider shared resource. | 
**fileName** | **string** | The File Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewImageCdxV1ProviderSharedResourceRequest struct via the builder pattern


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

