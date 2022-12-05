# \ModesV1Api

All URIs are relative to *https://psrc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSubjectMode**](ModesV1Api.md#DeleteSubjectMode) | **Delete** /mode/{subject} | Delete subject mode
[**GetMode**](ModesV1Api.md#GetMode) | **Get** /mode/{subject} | Get subject mode
[**GetTopLevelMode**](ModesV1Api.md#GetTopLevelMode) | **Get** /mode | Get global mode
[**UpdateMode**](ModesV1Api.md#UpdateMode) | **Put** /mode/{subject} | Update subject mode
[**UpdateTopLevelMode**](ModesV1Api.md#UpdateTopLevelMode) | **Put** /mode | Update global mode



## DeleteSubjectMode

> Mode DeleteSubjectMode(ctx, subject).Execute()

Delete subject mode



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
    subject := "subject_example" // string | Name of the subject

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ModesV1Api.DeleteSubjectMode(context.Background(), subject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModesV1Api.DeleteSubjectMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSubjectMode`: Mode
    fmt.Fprintf(os.Stdout, "Response from `ModesV1Api.DeleteSubjectMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Name of the subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubjectModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Mode**](Mode.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMode

> Mode GetMode(ctx, subject).DefaultToGlobal(defaultToGlobal).Execute()

Get subject mode



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
    subject := "subject_example" // string | Name of the subject
    defaultToGlobal := true // bool | Whether to return the global mode if subject mode not found (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ModesV1Api.GetMode(context.Background(), subject).DefaultToGlobal(defaultToGlobal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModesV1Api.GetMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMode`: Mode
    fmt.Fprintf(os.Stdout, "Response from `ModesV1Api.GetMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Name of the subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **defaultToGlobal** | **bool** | Whether to return the global mode if subject mode not found | 

### Return type

[**Mode**](Mode.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopLevelMode

> Mode GetTopLevelMode(ctx).Execute()

Get global mode



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
    resp, r, err := api_client.ModesV1Api.GetTopLevelMode(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModesV1Api.GetTopLevelMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopLevelMode`: Mode
    fmt.Fprintf(os.Stdout, "Response from `ModesV1Api.GetTopLevelMode`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopLevelModeRequest struct via the builder pattern


### Return type

[**Mode**](Mode.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMode

> ModeUpdateRequest UpdateMode(ctx, subject).ModeUpdateRequest(modeUpdateRequest).Force(force).Execute()

Update subject mode



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
    subject := "subject_example" // string | Name of the subject
    modeUpdateRequest := *openapiclient.NewModeUpdateRequest() // ModeUpdateRequest | Update Request
    force := true // bool | Whether to force update if setting mode to IMPORT and schemas currently exist (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ModesV1Api.UpdateMode(context.Background(), subject).ModeUpdateRequest(modeUpdateRequest).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModesV1Api.UpdateMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMode`: ModeUpdateRequest
    fmt.Fprintf(os.Stdout, "Response from `ModesV1Api.UpdateMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Name of the subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modeUpdateRequest** | [**ModeUpdateRequest**](ModeUpdateRequest.md) | Update Request | 
 **force** | **bool** | Whether to force update if setting mode to IMPORT and schemas currently exist | 

### Return type

[**ModeUpdateRequest**](ModeUpdateRequest.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTopLevelMode

> ModeUpdateRequest UpdateTopLevelMode(ctx).ModeUpdateRequest(modeUpdateRequest).Force(force).Execute()

Update global mode



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
    modeUpdateRequest := *openapiclient.NewModeUpdateRequest() // ModeUpdateRequest | Update Request
    force := true // bool | Whether to force update if setting mode to IMPORT and schemas currently exist (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ModesV1Api.UpdateTopLevelMode(context.Background()).ModeUpdateRequest(modeUpdateRequest).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModesV1Api.UpdateTopLevelMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTopLevelMode`: ModeUpdateRequest
    fmt.Fprintf(os.Stdout, "Response from `ModesV1Api.UpdateTopLevelMode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTopLevelModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modeUpdateRequest** | [**ModeUpdateRequest**](ModeUpdateRequest.md) | Update Request | 
 **force** | **bool** | Whether to force update if setting mode to IMPORT and schemas currently exist | 

### Return type

[**ModeUpdateRequest**](ModeUpdateRequest.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

