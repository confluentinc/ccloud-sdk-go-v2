# \ResourceSubscriptionsNotificationsV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationsV1ResourceSubscription**](ResourceSubscriptionsNotificationsV1Api.md#CreateNotificationsV1ResourceSubscription) | **Post** /notifications/v1/resource-subscriptions | Create a Resource Subscription
[**DeleteNotificationsV1ResourceSubscription**](ResourceSubscriptionsNotificationsV1Api.md#DeleteNotificationsV1ResourceSubscription) | **Delete** /notifications/v1/resource-subscriptions/{id} | Delete a Resource Subscription
[**GetNotificationsV1ResourceSubscription**](ResourceSubscriptionsNotificationsV1Api.md#GetNotificationsV1ResourceSubscription) | **Get** /notifications/v1/resource-subscriptions/{id} | Read a Resource Subscription
[**ListNotificationsV1ResourceSubscriptionsByFilter**](ResourceSubscriptionsNotificationsV1Api.md#ListNotificationsV1ResourceSubscriptionsByFilter) | **Get** /notifications/v1/resource-subscriptions:lookup | Lookup a list of resource subscription by filter
[**UpdateNotificationsV1ResourceSubscription**](ResourceSubscriptionsNotificationsV1Api.md#UpdateNotificationsV1ResourceSubscription) | **Patch** /notifications/v1/resource-subscriptions/{id} | Update a Resource Subscription



## CreateNotificationsV1ResourceSubscription

> NotificationsV1ResourceSubscription CreateNotificationsV1ResourceSubscription(ctx).NotificationsV1ResourceSubscription(notificationsV1ResourceSubscription).Execute()

Create a Resource Subscription



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
    notificationsV1ResourceSubscription := *openapiclient.NewNotificationsV1ResourceSubscription() // NotificationsV1ResourceSubscription |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceSubscriptionsNotificationsV1Api.CreateNotificationsV1ResourceSubscription(context.Background()).NotificationsV1ResourceSubscription(notificationsV1ResourceSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSubscriptionsNotificationsV1Api.CreateNotificationsV1ResourceSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsV1ResourceSubscription`: NotificationsV1ResourceSubscription
    fmt.Fprintf(os.Stdout, "Response from `ResourceSubscriptionsNotificationsV1Api.CreateNotificationsV1ResourceSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationsV1ResourceSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationsV1ResourceSubscription** | [**NotificationsV1ResourceSubscription**](NotificationsV1ResourceSubscription.md) |  | 

### Return type

[**NotificationsV1ResourceSubscription**](notifications.v1.ResourceSubscription.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token), [confluent_auth](../README.md#confluent_auth), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationsV1ResourceSubscription

> DeleteNotificationsV1ResourceSubscription(ctx, id).Execute()

Delete a Resource Subscription



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
    id := "id_example" // string | The unique identifier for the resource subscription.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceSubscriptionsNotificationsV1Api.DeleteNotificationsV1ResourceSubscription(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSubscriptionsNotificationsV1Api.DeleteNotificationsV1ResourceSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the resource subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationsV1ResourceSubscriptionRequest struct via the builder pattern


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


## GetNotificationsV1ResourceSubscription

> NotificationsV1ResourceSubscription GetNotificationsV1ResourceSubscription(ctx, id).Execute()

Read a Resource Subscription



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
    id := "id_example" // string | The unique identifier for the resource subscription.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceSubscriptionsNotificationsV1Api.GetNotificationsV1ResourceSubscription(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSubscriptionsNotificationsV1Api.GetNotificationsV1ResourceSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsV1ResourceSubscription`: NotificationsV1ResourceSubscription
    fmt.Fprintf(os.Stdout, "Response from `ResourceSubscriptionsNotificationsV1Api.GetNotificationsV1ResourceSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the resource subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsV1ResourceSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationsV1ResourceSubscription**](notifications.v1.ResourceSubscription.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token), [confluent_auth](../README.md#confluent_auth), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationsV1ResourceSubscriptionsByFilter

> NotificationsV1ResourceSubscriptionList ListNotificationsV1ResourceSubscriptionsByFilter(ctx).Resource(resource).ResourceType(resourceType).PageSize(pageSize).PageToken(pageToken).Execute()

Lookup a list of resource subscription by filter



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
    resource := "resource_example" // string | Confluent Cloud resource definition
    resourceType := "resourceType_example" // string | Confluent Cloud resource type
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 100)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceSubscriptionsNotificationsV1Api.ListNotificationsV1ResourceSubscriptionsByFilter(context.Background()).Resource(resource).ResourceType(resourceType).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSubscriptionsNotificationsV1Api.ListNotificationsV1ResourceSubscriptionsByFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNotificationsV1ResourceSubscriptionsByFilter`: NotificationsV1ResourceSubscriptionList
    fmt.Fprintf(os.Stdout, "Response from `ResourceSubscriptionsNotificationsV1Api.ListNotificationsV1ResourceSubscriptionsByFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationsV1ResourceSubscriptionsByFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resource** | **string** | Confluent Cloud resource definition | 
 **resourceType** | **string** | Confluent Cloud resource type | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 100]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**NotificationsV1ResourceSubscriptionList**](notifications.v1.ResourceSubscriptionList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token), [confluent_auth](../README.md#confluent_auth), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationsV1ResourceSubscription

> NotificationsV1ResourceSubscription UpdateNotificationsV1ResourceSubscription(ctx, id).NotificationsV1ResourceSubscriptionUpdate(notificationsV1ResourceSubscriptionUpdate).Execute()

Update a Resource Subscription



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
    id := "id_example" // string | The unique identifier for the resource subscription.
    notificationsV1ResourceSubscriptionUpdate := *openapiclient.NewNotificationsV1ResourceSubscriptionUpdate() // NotificationsV1ResourceSubscriptionUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceSubscriptionsNotificationsV1Api.UpdateNotificationsV1ResourceSubscription(context.Background(), id).NotificationsV1ResourceSubscriptionUpdate(notificationsV1ResourceSubscriptionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSubscriptionsNotificationsV1Api.UpdateNotificationsV1ResourceSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotificationsV1ResourceSubscription`: NotificationsV1ResourceSubscription
    fmt.Fprintf(os.Stdout, "Response from `ResourceSubscriptionsNotificationsV1Api.UpdateNotificationsV1ResourceSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the resource subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationsV1ResourceSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationsV1ResourceSubscriptionUpdate** | [**NotificationsV1ResourceSubscriptionUpdate**](NotificationsV1ResourceSubscriptionUpdate.md) |  | 

### Return type

[**NotificationsV1ResourceSubscription**](notifications.v1.ResourceSubscription.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token), [confluent_auth](../README.md#confluent_auth), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

