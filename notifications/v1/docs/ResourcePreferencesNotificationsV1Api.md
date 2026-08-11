# \ResourcePreferencesNotificationsV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationsV1ResourcePreference**](ResourcePreferencesNotificationsV1Api.md#CreateNotificationsV1ResourcePreference) | **Post** /notifications/v1/resource-preferences | Create a Resource Preference
[**DeleteNotificationsV1ResourcePreference**](ResourcePreferencesNotificationsV1Api.md#DeleteNotificationsV1ResourcePreference) | **Delete** /notifications/v1/resource-preferences/{id} | Delete a Resource Preference
[**GetNotificationsV1ResourcePreference**](ResourcePreferencesNotificationsV1Api.md#GetNotificationsV1ResourcePreference) | **Get** /notifications/v1/resource-preferences/{id} | Read a Resource Preference
[**GetNotificationsV1ResourcePreferenceByFilter**](ResourcePreferencesNotificationsV1Api.md#GetNotificationsV1ResourcePreferenceByFilter) | **Get** /notifications/v1/resource-preferences:lookup | Lookup a resource preference by filter (returns one)
[**UpdateNotificationsV1ResourcePreference**](ResourcePreferencesNotificationsV1Api.md#UpdateNotificationsV1ResourcePreference) | **Patch** /notifications/v1/resource-preferences/{id} | Update a Resource Preference



## CreateNotificationsV1ResourcePreference

> NotificationsV1ResourcePreference CreateNotificationsV1ResourcePreference(ctx).NotificationsV1ResourcePreference(notificationsV1ResourcePreference).Execute()

Create a Resource Preference



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
    notificationsV1ResourcePreference := *openapiclient.NewNotificationsV1ResourcePreference() // NotificationsV1ResourcePreference |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcePreferencesNotificationsV1Api.CreateNotificationsV1ResourcePreference(context.Background()).NotificationsV1ResourcePreference(notificationsV1ResourcePreference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePreferencesNotificationsV1Api.CreateNotificationsV1ResourcePreference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsV1ResourcePreference`: NotificationsV1ResourcePreference
    fmt.Fprintf(os.Stdout, "Response from `ResourcePreferencesNotificationsV1Api.CreateNotificationsV1ResourcePreference`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationsV1ResourcePreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationsV1ResourcePreference** | [**NotificationsV1ResourcePreference**](NotificationsV1ResourcePreference.md) |  | 

### Return type

[**NotificationsV1ResourcePreference**](notifications.v1.ResourcePreference.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token), [confluent_auth](../README.md#confluent_auth), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationsV1ResourcePreference

> DeleteNotificationsV1ResourcePreference(ctx, id).Execute()

Delete a Resource Preference



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
    id := "id_example" // string | The unique identifier for the resource preference.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcePreferencesNotificationsV1Api.DeleteNotificationsV1ResourcePreference(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePreferencesNotificationsV1Api.DeleteNotificationsV1ResourcePreference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the resource preference. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationsV1ResourcePreferenceRequest struct via the builder pattern


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


## GetNotificationsV1ResourcePreference

> NotificationsV1ResourcePreference GetNotificationsV1ResourcePreference(ctx, id).Execute()

Read a Resource Preference



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
    id := "id_example" // string | The unique identifier for the resource preference.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcePreferencesNotificationsV1Api.GetNotificationsV1ResourcePreference(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePreferencesNotificationsV1Api.GetNotificationsV1ResourcePreference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsV1ResourcePreference`: NotificationsV1ResourcePreference
    fmt.Fprintf(os.Stdout, "Response from `ResourcePreferencesNotificationsV1Api.GetNotificationsV1ResourcePreference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the resource preference. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsV1ResourcePreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationsV1ResourcePreference**](notifications.v1.ResourcePreference.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token), [confluent_auth](../README.md#confluent_auth), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationsV1ResourcePreferenceByFilter

> NotificationsV1ResourcePreference GetNotificationsV1ResourcePreferenceByFilter(ctx).Resource(resource).ResourceType(resourceType).PageSize(pageSize).PageToken(pageToken).Execute()

Lookup a resource preference by filter (returns one)



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
    resp, r, err := api_client.ResourcePreferencesNotificationsV1Api.GetNotificationsV1ResourcePreferenceByFilter(context.Background()).Resource(resource).ResourceType(resourceType).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePreferencesNotificationsV1Api.GetNotificationsV1ResourcePreferenceByFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsV1ResourcePreferenceByFilter`: NotificationsV1ResourcePreference
    fmt.Fprintf(os.Stdout, "Response from `ResourcePreferencesNotificationsV1Api.GetNotificationsV1ResourcePreferenceByFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsV1ResourcePreferenceByFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resource** | **string** | Confluent Cloud resource definition | 
 **resourceType** | **string** | Confluent Cloud resource type | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 100]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**NotificationsV1ResourcePreference**](notifications.v1.ResourcePreference.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token), [confluent_auth](../README.md#confluent_auth), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationsV1ResourcePreference

> NotificationsV1ResourcePreference UpdateNotificationsV1ResourcePreference(ctx, id).NotificationsV1ResourcePreference(notificationsV1ResourcePreference).Execute()

Update a Resource Preference



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
    id := "id_example" // string | The unique identifier for the resource preference.
    notificationsV1ResourcePreference := *openapiclient.NewNotificationsV1ResourcePreference() // NotificationsV1ResourcePreference |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcePreferencesNotificationsV1Api.UpdateNotificationsV1ResourcePreference(context.Background(), id).NotificationsV1ResourcePreference(notificationsV1ResourcePreference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePreferencesNotificationsV1Api.UpdateNotificationsV1ResourcePreference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotificationsV1ResourcePreference`: NotificationsV1ResourcePreference
    fmt.Fprintf(os.Stdout, "Response from `ResourcePreferencesNotificationsV1Api.UpdateNotificationsV1ResourcePreference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the resource preference. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationsV1ResourcePreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationsV1ResourcePreference** | [**NotificationsV1ResourcePreference**](NotificationsV1ResourcePreference.md) |  | 

### Return type

[**NotificationsV1ResourcePreference**](notifications.v1.ResourcePreference.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token), [confluent_auth](../README.md#confluent_auth), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

