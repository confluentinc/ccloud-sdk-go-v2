# \CustomCodeLoggingsCclV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCclV1CustomCodeLogging**](CustomCodeLoggingsCclV1Api.md#CreateCclV1CustomCodeLogging) | **Post** /ccl/v1/custom-code-loggings | Create a Custom Code Logging
[**DeleteCclV1CustomCodeLogging**](CustomCodeLoggingsCclV1Api.md#DeleteCclV1CustomCodeLogging) | **Delete** /ccl/v1/custom-code-loggings/{id} | Delete a Custom Code Logging
[**GetCclV1CustomCodeLogging**](CustomCodeLoggingsCclV1Api.md#GetCclV1CustomCodeLogging) | **Get** /ccl/v1/custom-code-loggings/{id} | Read a Custom Code Logging
[**ListCclV1CustomCodeLoggings**](CustomCodeLoggingsCclV1Api.md#ListCclV1CustomCodeLoggings) | **Get** /ccl/v1/custom-code-loggings | List of Custom Code Loggings
[**UpdateCclV1CustomCodeLogging**](CustomCodeLoggingsCclV1Api.md#UpdateCclV1CustomCodeLogging) | **Patch** /ccl/v1/custom-code-loggings/{id} | Update a Custom Code Logging



## CreateCclV1CustomCodeLogging

> CclV1CustomCodeLogging CreateCclV1CustomCodeLogging(ctx).CclV1CustomCodeLogging(cclV1CustomCodeLogging).Execute()

Create a Custom Code Logging



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
    cclV1CustomCodeLogging := *openapiclient.NewCclV1CustomCodeLogging() // CclV1CustomCodeLogging |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomCodeLoggingsCclV1Api.CreateCclV1CustomCodeLogging(context.Background()).CclV1CustomCodeLogging(cclV1CustomCodeLogging).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomCodeLoggingsCclV1Api.CreateCclV1CustomCodeLogging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCclV1CustomCodeLogging`: CclV1CustomCodeLogging
    fmt.Fprintf(os.Stdout, "Response from `CustomCodeLoggingsCclV1Api.CreateCclV1CustomCodeLogging`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCclV1CustomCodeLoggingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cclV1CustomCodeLogging** | [**CclV1CustomCodeLogging**](CclV1CustomCodeLogging.md) |  | 

### Return type

[**CclV1CustomCodeLogging**](ccl.v1.CustomCodeLogging.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCclV1CustomCodeLogging

> DeleteCclV1CustomCodeLogging(ctx, id).Execute()

Delete a Custom Code Logging



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
    id := "id_example" // string | The unique identifier for the custom code logging.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomCodeLoggingsCclV1Api.DeleteCclV1CustomCodeLogging(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomCodeLoggingsCclV1Api.DeleteCclV1CustomCodeLogging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the custom code logging. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCclV1CustomCodeLoggingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCclV1CustomCodeLogging

> CclV1CustomCodeLogging GetCclV1CustomCodeLogging(ctx, id).Execute()

Read a Custom Code Logging



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
    id := "id_example" // string | The unique identifier for the custom code logging.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomCodeLoggingsCclV1Api.GetCclV1CustomCodeLogging(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomCodeLoggingsCclV1Api.GetCclV1CustomCodeLogging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCclV1CustomCodeLogging`: CclV1CustomCodeLogging
    fmt.Fprintf(os.Stdout, "Response from `CustomCodeLoggingsCclV1Api.GetCclV1CustomCodeLogging`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the custom code logging. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCclV1CustomCodeLoggingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CclV1CustomCodeLogging**](ccl.v1.CustomCodeLogging.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCclV1CustomCodeLoggings

> CclV1CustomCodeLoggingList ListCclV1CustomCodeLoggings(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Custom Code Loggings



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
    resp, r, err := api_client.CustomCodeLoggingsCclV1Api.ListCclV1CustomCodeLoggings(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomCodeLoggingsCclV1Api.ListCclV1CustomCodeLoggings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCclV1CustomCodeLoggings`: CclV1CustomCodeLoggingList
    fmt.Fprintf(os.Stdout, "Response from `CustomCodeLoggingsCclV1Api.ListCclV1CustomCodeLoggings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCclV1CustomCodeLoggingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**CclV1CustomCodeLoggingList**](ccl.v1.CustomCodeLoggingList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCclV1CustomCodeLogging

> CclV1CustomCodeLogging UpdateCclV1CustomCodeLogging(ctx, id).CclV1CustomCodeLoggingUpdate(cclV1CustomCodeLoggingUpdate).Execute()

Update a Custom Code Logging



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
    id := "id_example" // string | The unique identifier for the custom code logging.
    cclV1CustomCodeLoggingUpdate := *openapiclient.NewCclV1CustomCodeLoggingUpdate() // CclV1CustomCodeLoggingUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomCodeLoggingsCclV1Api.UpdateCclV1CustomCodeLogging(context.Background(), id).CclV1CustomCodeLoggingUpdate(cclV1CustomCodeLoggingUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomCodeLoggingsCclV1Api.UpdateCclV1CustomCodeLogging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCclV1CustomCodeLogging`: CclV1CustomCodeLogging
    fmt.Fprintf(os.Stdout, "Response from `CustomCodeLoggingsCclV1Api.UpdateCclV1CustomCodeLogging`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the custom code logging. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCclV1CustomCodeLoggingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cclV1CustomCodeLoggingUpdate** | [**CclV1CustomCodeLoggingUpdate**](CclV1CustomCodeLoggingUpdate.md) |  | 

### Return type

[**CclV1CustomCodeLogging**](ccl.v1.CustomCodeLogging.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

