# \IntegrationsNotificationsV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationsV1Integration**](IntegrationsNotificationsV1Api.md#CreateNotificationsV1Integration) | **Post** /notifications/v1/integrations | Create an Integration
[**DeleteNotificationsV1Integration**](IntegrationsNotificationsV1Api.md#DeleteNotificationsV1Integration) | **Delete** /notifications/v1/integrations/{id} | Delete an Integration
[**GetNotificationsV1Integration**](IntegrationsNotificationsV1Api.md#GetNotificationsV1Integration) | **Get** /notifications/v1/integrations/{id} | Read an Integration
[**ListNotificationsV1Integrations**](IntegrationsNotificationsV1Api.md#ListNotificationsV1Integrations) | **Get** /notifications/v1/integrations | List of Integrations
[**TestNotificationsV1Integration**](IntegrationsNotificationsV1Api.md#TestNotificationsV1Integration) | **Post** /notifications/v1/integrations:test | Test a Webhook, Slack or Microsoft Teams integration
[**UpdateNotificationsV1Integration**](IntegrationsNotificationsV1Api.md#UpdateNotificationsV1Integration) | **Patch** /notifications/v1/integrations/{id} | Update an Integration



## CreateNotificationsV1Integration

> NotificationsV1Integration CreateNotificationsV1Integration(ctx).NotificationsV1Integration(notificationsV1Integration).Execute()

Create an Integration



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
    notificationsV1Integration := *openapiclient.NewNotificationsV1Integration() // NotificationsV1Integration |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsNotificationsV1Api.CreateNotificationsV1Integration(context.Background()).NotificationsV1Integration(notificationsV1Integration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsNotificationsV1Api.CreateNotificationsV1Integration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsV1Integration`: NotificationsV1Integration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsNotificationsV1Api.CreateNotificationsV1Integration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationsV1IntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationsV1Integration** | [**NotificationsV1Integration**](NotificationsV1Integration.md) |  | 

### Return type

[**NotificationsV1Integration**](notifications.v1.Integration.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token), [confluent_auth](../README.md#confluent_auth), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationsV1Integration

> DeleteNotificationsV1Integration(ctx, id).Execute()

Delete an Integration



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
    id := "id_example" // string | The unique identifier for the integration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsNotificationsV1Api.DeleteNotificationsV1Integration(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsNotificationsV1Api.DeleteNotificationsV1Integration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the integration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationsV1IntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token), [confluent_auth](../README.md#confluent_auth), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationsV1Integration

> NotificationsV1Integration GetNotificationsV1Integration(ctx, id).Execute()

Read an Integration



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
    id := "id_example" // string | The unique identifier for the integration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsNotificationsV1Api.GetNotificationsV1Integration(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsNotificationsV1Api.GetNotificationsV1Integration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsV1Integration`: NotificationsV1Integration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsNotificationsV1Api.GetNotificationsV1Integration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the integration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsV1IntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationsV1Integration**](notifications.v1.Integration.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token), [confluent_auth](../README.md#confluent_auth), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationsV1Integrations

> NotificationsV1IntegrationList ListNotificationsV1Integrations(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Integrations



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
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 100)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsNotificationsV1Api.ListNotificationsV1Integrations(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsNotificationsV1Api.ListNotificationsV1Integrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNotificationsV1Integrations`: NotificationsV1IntegrationList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsNotificationsV1Api.ListNotificationsV1Integrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationsV1IntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 100]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**NotificationsV1IntegrationList**](notifications.v1.IntegrationList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token), [confluent_auth](../README.md#confluent_auth), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestNotificationsV1Integration

> TestNotificationsV1Integration(ctx).NotificationsV1Integration(notificationsV1Integration).Execute()

Test a Webhook, Slack or Microsoft Teams integration



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
    notificationsV1Integration := *openapiclient.NewNotificationsV1Integration() // NotificationsV1Integration |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsNotificationsV1Api.TestNotificationsV1Integration(context.Background()).NotificationsV1Integration(notificationsV1Integration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsNotificationsV1Api.TestNotificationsV1Integration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestNotificationsV1IntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationsV1Integration** | [**NotificationsV1Integration**](NotificationsV1Integration.md) |  | 

### Return type

 (empty response body)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token), [confluent_auth](../README.md#confluent_auth), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationsV1Integration

> NotificationsV1Integration UpdateNotificationsV1Integration(ctx, id).NotificationsV1Integration(notificationsV1Integration).Execute()

Update an Integration



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
    id := "id_example" // string | The unique identifier for the integration.
    notificationsV1Integration := *openapiclient.NewNotificationsV1Integration() // NotificationsV1Integration |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsNotificationsV1Api.UpdateNotificationsV1Integration(context.Background(), id).NotificationsV1Integration(notificationsV1Integration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsNotificationsV1Api.UpdateNotificationsV1Integration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotificationsV1Integration`: NotificationsV1Integration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsNotificationsV1Api.UpdateNotificationsV1Integration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the integration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationsV1IntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationsV1Integration** | [**NotificationsV1Integration**](NotificationsV1Integration.md) |  | 

### Return type

[**NotificationsV1Integration**](notifications.v1.Integration.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token), [confluent_auth](../README.md#confluent_auth), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

