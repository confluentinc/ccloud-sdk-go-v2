# \WidgetsTemplateV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTemplateV1Widget**](WidgetsTemplateV1Api.md#CreateTemplateV1Widget) | **Post** /template/v1/widgets | Create a Widget
[**DeleteTemplateV1Widget**](WidgetsTemplateV1Api.md#DeleteTemplateV1Widget) | **Delete** /template/v1/widgets/{id} | Delete a Widget
[**GetTemplateV1Widget**](WidgetsTemplateV1Api.md#GetTemplateV1Widget) | **Get** /template/v1/widgets/{id} | Read a Widget
[**ListTemplateV1Widgets**](WidgetsTemplateV1Api.md#ListTemplateV1Widgets) | **Get** /template/v1/widgets | List of Widgets
[**UpdateTemplateV1Widget**](WidgetsTemplateV1Api.md#UpdateTemplateV1Widget) | **Patch** /template/v1/widgets/{id} | Update a Widget



## CreateTemplateV1Widget

> TemplateV1Widget CreateTemplateV1Widget(ctx).TemplateV1Widget(templateV1Widget).Execute()

Create a Widget



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
    templateV1Widget := *openapiclient.NewTemplateV1Widget() // TemplateV1Widget |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WidgetsTemplateV1Api.CreateTemplateV1Widget(context.Background()).TemplateV1Widget(templateV1Widget).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WidgetsTemplateV1Api.CreateTemplateV1Widget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTemplateV1Widget`: TemplateV1Widget
    fmt.Fprintf(os.Stdout, "Response from `WidgetsTemplateV1Api.CreateTemplateV1Widget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateV1WidgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateV1Widget** | [**TemplateV1Widget**](TemplateV1Widget.md) |  | 

### Return type

[**TemplateV1Widget**](template.v1.Widget.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTemplateV1Widget

> DeleteTemplateV1Widget(ctx, id).Execute()

Delete a Widget



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
    id := "id_example" // string | The unique identifier for the widget.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WidgetsTemplateV1Api.DeleteTemplateV1Widget(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WidgetsTemplateV1Api.DeleteTemplateV1Widget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the widget. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplateV1WidgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateV1Widget

> TemplateV1Widget GetTemplateV1Widget(ctx, id).Execute()

Read a Widget



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
    id := "id_example" // string | The unique identifier for the widget.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WidgetsTemplateV1Api.GetTemplateV1Widget(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WidgetsTemplateV1Api.GetTemplateV1Widget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplateV1Widget`: TemplateV1Widget
    fmt.Fprintf(os.Stdout, "Response from `WidgetsTemplateV1Api.GetTemplateV1Widget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the widget. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateV1WidgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TemplateV1Widget**](template.v1.Widget.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTemplateV1Widgets

> TemplateV1WidgetList ListTemplateV1Widgets(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Widgets



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
    resp, r, err := api_client.WidgetsTemplateV1Api.ListTemplateV1Widgets(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WidgetsTemplateV1Api.ListTemplateV1Widgets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTemplateV1Widgets`: TemplateV1WidgetList
    fmt.Fprintf(os.Stdout, "Response from `WidgetsTemplateV1Api.ListTemplateV1Widgets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTemplateV1WidgetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**TemplateV1WidgetList**](template.v1.WidgetList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTemplateV1Widget

> TemplateV1Widget UpdateTemplateV1Widget(ctx, id).TemplateV1WidgetUpdate(templateV1WidgetUpdate).Execute()

Update a Widget



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
    id := "id_example" // string | The unique identifier for the widget.
    templateV1WidgetUpdate := *openapiclient.NewTemplateV1WidgetUpdate() // TemplateV1WidgetUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WidgetsTemplateV1Api.UpdateTemplateV1Widget(context.Background(), id).TemplateV1WidgetUpdate(templateV1WidgetUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WidgetsTemplateV1Api.UpdateTemplateV1Widget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTemplateV1Widget`: TemplateV1Widget
    fmt.Fprintf(os.Stdout, "Response from `WidgetsTemplateV1Api.UpdateTemplateV1Widget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the widget. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTemplateV1WidgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateV1WidgetUpdate** | [**TemplateV1WidgetUpdate**](TemplateV1WidgetUpdate.md) |  | 

### Return type

[**TemplateV1Widget**](template.v1.Widget.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

