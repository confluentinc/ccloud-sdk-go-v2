# \NotificationTypesNotificationsV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNotificationsV1NotificationType**](NotificationTypesNotificationsV1Api.md#GetNotificationsV1NotificationType) | **Get** /notifications/v1/notification-types/{id} | Read a Notification Type
[**ListNotificationsV1NotificationTypes**](NotificationTypesNotificationsV1Api.md#ListNotificationsV1NotificationTypes) | **Get** /notifications/v1/notification-types | Retrieve a list of all notification types for the resource type.



## GetNotificationsV1NotificationType

> NotificationsV1NotificationType GetNotificationsV1NotificationType(ctx, id).Execute()

Read a Notification Type



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
    id := "id_example" // string | The unique identifier for the notification type.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationTypesNotificationsV1Api.GetNotificationsV1NotificationType(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationTypesNotificationsV1Api.GetNotificationsV1NotificationType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsV1NotificationType`: NotificationsV1NotificationType
    fmt.Fprintf(os.Stdout, "Response from `NotificationTypesNotificationsV1Api.GetNotificationsV1NotificationType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the notification type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsV1NotificationTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationsV1NotificationType**](notifications.v1.NotificationType.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token), [confluent_auth](../README.md#confluent_auth), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationsV1NotificationTypes

> NotificationsV1NotificationTypeList ListNotificationsV1NotificationTypes(ctx).ResourceType(resourceType).PageSize(pageSize).PageToken(pageToken).Execute()

Retrieve a list of all notification types for the resource type.



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
    resourceType := "resourceType_example" // string | Confluent Cloud resource type (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 100)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationTypesNotificationsV1Api.ListNotificationsV1NotificationTypes(context.Background()).ResourceType(resourceType).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationTypesNotificationsV1Api.ListNotificationsV1NotificationTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNotificationsV1NotificationTypes`: NotificationsV1NotificationTypeList
    fmt.Fprintf(os.Stdout, "Response from `NotificationTypesNotificationsV1Api.ListNotificationsV1NotificationTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationsV1NotificationTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceType** | **string** | Confluent Cloud resource type | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 100]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**NotificationsV1NotificationTypeList**](notifications.v1.NotificationTypeList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token), [confluent_auth](../README.md#confluent_auth), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

