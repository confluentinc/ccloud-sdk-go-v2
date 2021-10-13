# \TemplatesNotificationsAdminV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationsAdminV1Template**](TemplatesNotificationsAdminV1Api.md#CreateNotificationsAdminV1Template) | **Post** /notifications-admin/v1/templates | Create a Template
[**DeleteNotificationsAdminV1Template**](TemplatesNotificationsAdminV1Api.md#DeleteNotificationsAdminV1Template) | **Delete** /notifications-admin/v1/templates/{id} | Delete a Template
[**GetNotificationsAdminV1Template**](TemplatesNotificationsAdminV1Api.md#GetNotificationsAdminV1Template) | **Get** /notifications-admin/v1/templates/{id} | Read a Template
[**ListNotificationsAdminV1Templates**](TemplatesNotificationsAdminV1Api.md#ListNotificationsAdminV1Templates) | **Get** /notifications-admin/v1/templates | List of Templates
[**UpdateNotificationsAdminV1Template**](TemplatesNotificationsAdminV1Api.md#UpdateNotificationsAdminV1Template) | **Patch** /notifications-admin/v1/templates/{id} | Update a Template



## CreateNotificationsAdminV1Template

> NotificationsAdminV1Template CreateNotificationsAdminV1Template(ctx).NotificationsAdminV1Template(notificationsAdminV1Template).Execute()

Create a Template



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
    notificationsAdminV1Template := *openapiclient.NewNotificationsAdminV1Template() // NotificationsAdminV1Template |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplatesNotificationsAdminV1Api.CreateNotificationsAdminV1Template(context.Background()).NotificationsAdminV1Template(notificationsAdminV1Template).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesNotificationsAdminV1Api.CreateNotificationsAdminV1Template``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsAdminV1Template`: NotificationsAdminV1Template
    fmt.Fprintf(os.Stdout, "Response from `TemplatesNotificationsAdminV1Api.CreateNotificationsAdminV1Template`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationsAdminV1TemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationsAdminV1Template** | [**NotificationsAdminV1Template**](NotificationsAdminV1Template.md) |  | 

### Return type

[**NotificationsAdminV1Template**](notifications-admin.v1.Template.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationsAdminV1Template

> DeleteNotificationsAdminV1Template(ctx, id).Execute()

Delete a Template



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
    id := "id_example" // string | The unique identifier for the template.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplatesNotificationsAdminV1Api.DeleteNotificationsAdminV1Template(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesNotificationsAdminV1Api.DeleteNotificationsAdminV1Template``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationsAdminV1TemplateRequest struct via the builder pattern


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


## GetNotificationsAdminV1Template

> NotificationsAdminV1Template GetNotificationsAdminV1Template(ctx, id).Execute()

Read a Template



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
    id := "id_example" // string | The unique identifier for the template.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplatesNotificationsAdminV1Api.GetNotificationsAdminV1Template(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesNotificationsAdminV1Api.GetNotificationsAdminV1Template``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsAdminV1Template`: NotificationsAdminV1Template
    fmt.Fprintf(os.Stdout, "Response from `TemplatesNotificationsAdminV1Api.GetNotificationsAdminV1Template`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsAdminV1TemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationsAdminV1Template**](notifications-admin.v1.Template.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationsAdminV1Templates

> NotificationsAdminV1TemplateList ListNotificationsAdminV1Templates(ctx).NotificationType(notificationType).IntegrationType(integrationType).PageSize(pageSize).PageToken(pageToken).Execute()

List of Templates



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
    notificationType := "alert.disk_full" // string | Filter the results by exact match for notification_type. (optional)
    integrationType := "SLACK" // string | Filter the results by exact match for integration_type. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplatesNotificationsAdminV1Api.ListNotificationsAdminV1Templates(context.Background()).NotificationType(notificationType).IntegrationType(integrationType).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesNotificationsAdminV1Api.ListNotificationsAdminV1Templates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNotificationsAdminV1Templates`: NotificationsAdminV1TemplateList
    fmt.Fprintf(os.Stdout, "Response from `TemplatesNotificationsAdminV1Api.ListNotificationsAdminV1Templates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationsAdminV1TemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationType** | **string** | Filter the results by exact match for notification_type. | 
 **integrationType** | **string** | Filter the results by exact match for integration_type. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**NotificationsAdminV1TemplateList**](NotificationsAdminV1TemplateList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationsAdminV1Template

> NotificationsAdminV1Template UpdateNotificationsAdminV1Template(ctx, id).NotificationsAdminV1Template(notificationsAdminV1Template).Execute()

Update a Template



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
    id := "id_example" // string | The unique identifier for the template.
    notificationsAdminV1Template := *openapiclient.NewNotificationsAdminV1Template() // NotificationsAdminV1Template |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplatesNotificationsAdminV1Api.UpdateNotificationsAdminV1Template(context.Background(), id).NotificationsAdminV1Template(notificationsAdminV1Template).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesNotificationsAdminV1Api.UpdateNotificationsAdminV1Template``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotificationsAdminV1Template`: NotificationsAdminV1Template
    fmt.Fprintf(os.Stdout, "Response from `TemplatesNotificationsAdminV1Api.UpdateNotificationsAdminV1Template`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationsAdminV1TemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationsAdminV1Template** | [**NotificationsAdminV1Template**](NotificationsAdminV1Template.md) |  | 

### Return type

[**NotificationsAdminV1Template**](notifications-admin.v1.Template.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

