# \UserNotificationsNotificationsV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNotificationsV1UserNotification**](UserNotificationsNotificationsV1Api.md#GetNotificationsV1UserNotification) | **Get** /notifications/v1/user-notifications/{id} | Read a User Notification
[**GetNotificationsV1UserNotificationsSummary**](UserNotificationsNotificationsV1Api.md#GetNotificationsV1UserNotificationsSummary) | **Get** /notifications/v1/user-notifications:summary | Get notification summary
[**ListNotificationsV1UserNotifications**](UserNotificationsNotificationsV1Api.md#ListNotificationsV1UserNotifications) | **Get** /notifications/v1/user-notifications | List of User Notifications
[**MarkAllNotificationsV1UserNotifications**](UserNotificationsNotificationsV1Api.md#MarkAllNotificationsV1UserNotifications) | **Patch** /notifications/v1/user-notifications:mark-all | Mark multiple notifications read or unread
[**UpdateNotificationsV1UserNotification**](UserNotificationsNotificationsV1Api.md#UpdateNotificationsV1UserNotification) | **Patch** /notifications/v1/user-notifications/{id} | Update a User Notification



## GetNotificationsV1UserNotification

> NotificationsV1UserNotification GetNotificationsV1UserNotification(ctx, id).Execute()

Read a User Notification



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
    id := "id_example" // string | The unique identifier for the user notification.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserNotificationsNotificationsV1Api.GetNotificationsV1UserNotification(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserNotificationsNotificationsV1Api.GetNotificationsV1UserNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsV1UserNotification`: NotificationsV1UserNotification
    fmt.Fprintf(os.Stdout, "Response from `UserNotificationsNotificationsV1Api.GetNotificationsV1UserNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the user notification. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsV1UserNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationsV1UserNotification**](notifications.v1.UserNotification.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationsV1UserNotificationsSummary

> NotificationsV1Summary GetNotificationsV1UserNotificationsSummary(ctx).Execute()

Get notification summary



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
    resp, r, err := api_client.UserNotificationsNotificationsV1Api.GetNotificationsV1UserNotificationsSummary(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserNotificationsNotificationsV1Api.GetNotificationsV1UserNotificationsSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsV1UserNotificationsSummary`: NotificationsV1Summary
    fmt.Fprintf(os.Stdout, "Response from `UserNotificationsNotificationsV1Api.GetNotificationsV1UserNotificationsSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsV1UserNotificationsSummaryRequest struct via the builder pattern


### Return type

[**NotificationsV1Summary**](NotificationsV1Summary.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationsV1UserNotifications

> NotificationsV1UserNotificationList ListNotificationsV1UserNotifications(ctx).Read(read).Severity(severity).Include(include).ResourceType(resourceType).ResourceCrn(resourceCrn).Search(search).TimeRange(timeRange).PageSize(pageSize).PageToken(pageToken).Sort(sort).Execute()

List of User Notifications



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
    read := true // bool | Filter the results where read is true or false. (optional)
    severity := *openapiclient.NewMultipleSearchFilter() // MultipleSearchFilter | Filter notifications by severity. Pass the parameter multiple times to match any of the given values (`?severity=CRITICAL&severity=WARN`). A notification matches if its `severity` equals any of the listed values.  (optional)
    include := "integrations,recommended_actions" // string | Comma-separated list of optional fields to populate in the response items. Allowed values: `integrations`, `recommended_actions`. By default these fields are omitted from list responses to keep collection payloads slim; set this parameter to opt in. This is a partial-response selector, not a value filter.  (optional)
    resourceType := *openapiclient.NewMultipleSearchFilter() // MultipleSearchFilter | Filter notifications by the Confluent Cloud resource type they relate to. Pass the parameter multiple times to match any of the given values (`?resource.type=CLUSTER&resource.type=CONNECTOR`). A notification matches if its `resource.type` equals any of the listed values.  (optional)
    resourceCrn :=  // MultipleSearchFilter | Filter notifications by the CRN of the Confluent Cloud resource they relate to. Pass the parameter multiple times to match any of the given CRNs; a notification matches if its `resource.crn` equals any of the listed values.  (optional)
    search := "cluster failure" // string | Free-text partial-match search across the embedded notification type's `display_name` and `description`.  (optional)
    timeRange := "PAST_24H" // string | Filter notifications by a preset time window relative to now. Allowed values: `PAST_24H` (last 24 hours), `PAST_7D` (last 7 days), `PAST_30D` (last 30 days).  (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 100)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)
    sort := []string{"Sort_example"} // []string | The list of fields and directions that are used to sort the collection. (optional) (default to ["-received_at"])

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserNotificationsNotificationsV1Api.ListNotificationsV1UserNotifications(context.Background()).Read(read).Severity(severity).Include(include).ResourceType(resourceType).ResourceCrn(resourceCrn).Search(search).TimeRange(timeRange).PageSize(pageSize).PageToken(pageToken).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserNotificationsNotificationsV1Api.ListNotificationsV1UserNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNotificationsV1UserNotifications`: NotificationsV1UserNotificationList
    fmt.Fprintf(os.Stdout, "Response from `UserNotificationsNotificationsV1Api.ListNotificationsV1UserNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationsV1UserNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **read** | **bool** | Filter the results where read is true or false. | 
 **severity** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter notifications by severity. Pass the parameter multiple times to match any of the given values (&#x60;?severity&#x3D;CRITICAL&amp;severity&#x3D;WARN&#x60;). A notification matches if its &#x60;severity&#x60; equals any of the listed values.  | 
 **include** | **string** | Comma-separated list of optional fields to populate in the response items. Allowed values: &#x60;integrations&#x60;, &#x60;recommended_actions&#x60;. By default these fields are omitted from list responses to keep collection payloads slim; set this parameter to opt in. This is a partial-response selector, not a value filter.  | 
 **resourceType** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter notifications by the Confluent Cloud resource type they relate to. Pass the parameter multiple times to match any of the given values (&#x60;?resource.type&#x3D;CLUSTER&amp;resource.type&#x3D;CONNECTOR&#x60;). A notification matches if its &#x60;resource.type&#x60; equals any of the listed values.  | 
 **resourceCrn** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter notifications by the CRN of the Confluent Cloud resource they relate to. Pass the parameter multiple times to match any of the given CRNs; a notification matches if its &#x60;resource.crn&#x60; equals any of the listed values.  | 
 **search** | **string** | Free-text partial-match search across the embedded notification type&#39;s &#x60;display_name&#x60; and &#x60;description&#x60;.  | 
 **timeRange** | **string** | Filter notifications by a preset time window relative to now. Allowed values: &#x60;PAST_24H&#x60; (last 24 hours), &#x60;PAST_7D&#x60; (last 7 days), &#x60;PAST_30D&#x60; (last 30 days).  | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 100]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 
 **sort** | **[]string** | The list of fields and directions that are used to sort the collection. | [default to [&quot;-received_at&quot;]]

### Return type

[**NotificationsV1UserNotificationList**](notifications.v1.UserNotificationList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkAllNotificationsV1UserNotifications

> MarkAllNotificationsV1UserNotifications(ctx).Read(read).Severity(severity).ResourceType(resourceType).ResourceCrn(resourceCrn).Search(search).TimeRange(timeRange).NotificationsV1UpdateUserNotificationsReadRequest(notificationsV1UpdateUserNotificationsReadRequest).Execute()

Mark multiple notifications read or unread



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
    read := "false" // string | Scope the update to notifications with the given read state. Accepts `true` or `false`. Combine with a body of `{ \"read\": true }` to mark all currently-unread notifications as read (or vice versa).  (optional)
    severity := *openapiclient.NewMultipleSearchFilter() // MultipleSearchFilter | Filter notifications by severity. Pass the parameter multiple times to match any of the given values (`?severity=CRITICAL&severity=WARN`). A notification matches if its `severity` equals any of the listed values.  (optional)
    resourceType := *openapiclient.NewMultipleSearchFilter() // MultipleSearchFilter | Filter notifications by the Confluent Cloud resource type they relate to. Pass the parameter multiple times to match any of the given values (`?resource.type=CLUSTER&resource.type=CONNECTOR`). A notification matches if its `resource.type` equals any of the listed values.  (optional)
    resourceCrn :=  // MultipleSearchFilter | Filter notifications by the CRN of the Confluent Cloud resource they relate to. Pass the parameter multiple times to match any of the given CRNs; a notification matches if its `resource.crn` equals any of the listed values.  (optional)
    search := "cluster failure" // string | Free-text partial-match search across the embedded notification type's `display_name` and `description`.  (optional)
    timeRange := "PAST_24H" // string | Filter notifications by a preset time window relative to now. Allowed values: `PAST_24H` (last 24 hours), `PAST_7D` (last 7 days), `PAST_30D` (last 30 days).  (optional)
    notificationsV1UpdateUserNotificationsReadRequest := *openapiclient.NewNotificationsV1UpdateUserNotificationsReadRequest(true) // NotificationsV1UpdateUserNotificationsReadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserNotificationsNotificationsV1Api.MarkAllNotificationsV1UserNotifications(context.Background()).Read(read).Severity(severity).ResourceType(resourceType).ResourceCrn(resourceCrn).Search(search).TimeRange(timeRange).NotificationsV1UpdateUserNotificationsReadRequest(notificationsV1UpdateUserNotificationsReadRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserNotificationsNotificationsV1Api.MarkAllNotificationsV1UserNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkAllNotificationsV1UserNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **read** | **string** | Scope the update to notifications with the given read state. Accepts &#x60;true&#x60; or &#x60;false&#x60;. Combine with a body of &#x60;{ \&quot;read\&quot;: true }&#x60; to mark all currently-unread notifications as read (or vice versa).  | 
 **severity** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter notifications by severity. Pass the parameter multiple times to match any of the given values (&#x60;?severity&#x3D;CRITICAL&amp;severity&#x3D;WARN&#x60;). A notification matches if its &#x60;severity&#x60; equals any of the listed values.  | 
 **resourceType** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter notifications by the Confluent Cloud resource type they relate to. Pass the parameter multiple times to match any of the given values (&#x60;?resource.type&#x3D;CLUSTER&amp;resource.type&#x3D;CONNECTOR&#x60;). A notification matches if its &#x60;resource.type&#x60; equals any of the listed values.  | 
 **resourceCrn** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter notifications by the CRN of the Confluent Cloud resource they relate to. Pass the parameter multiple times to match any of the given CRNs; a notification matches if its &#x60;resource.crn&#x60; equals any of the listed values.  | 
 **search** | **string** | Free-text partial-match search across the embedded notification type&#39;s &#x60;display_name&#x60; and &#x60;description&#x60;.  | 
 **timeRange** | **string** | Filter notifications by a preset time window relative to now. Allowed values: &#x60;PAST_24H&#x60; (last 24 hours), &#x60;PAST_7D&#x60; (last 7 days), &#x60;PAST_30D&#x60; (last 30 days).  | 
 **notificationsV1UpdateUserNotificationsReadRequest** | [**NotificationsV1UpdateUserNotificationsReadRequest**](NotificationsV1UpdateUserNotificationsReadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationsV1UserNotification

> NotificationsV1UserNotification UpdateNotificationsV1UserNotification(ctx, id).NotificationsV1UserNotification(notificationsV1UserNotification).Execute()

Update a User Notification



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
    id := "id_example" // string | The unique identifier for the user notification.
    notificationsV1UserNotification := *openapiclient.NewNotificationsV1UserNotification() // NotificationsV1UserNotification |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserNotificationsNotificationsV1Api.UpdateNotificationsV1UserNotification(context.Background(), id).NotificationsV1UserNotification(notificationsV1UserNotification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserNotificationsNotificationsV1Api.UpdateNotificationsV1UserNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotificationsV1UserNotification`: NotificationsV1UserNotification
    fmt.Fprintf(os.Stdout, "Response from `UserNotificationsNotificationsV1Api.UpdateNotificationsV1UserNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the user notification. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationsV1UserNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationsV1UserNotification** | [**NotificationsV1UserNotification**](NotificationsV1UserNotification.md) |  | 

### Return type

[**NotificationsV1UserNotification**](notifications.v1.UserNotification.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

