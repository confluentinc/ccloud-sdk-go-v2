# \NotificationTopicConfigsStreamCatalogV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableNotificationsStreamCatalogV2NotificationTopicConfig**](NotificationTopicConfigsStreamCatalogV2Api.md#DisableNotificationsStreamCatalogV2NotificationTopicConfig) | **Delete** /stream-catalog/v2/clusters/{catalog_cluster_id}/notification-topic-configs | Delete the catalog notifications topic configuration
[**EnableNotificationsStreamCatalogV2NotificationTopicConfig**](NotificationTopicConfigsStreamCatalogV2Api.md#EnableNotificationsStreamCatalogV2NotificationTopicConfig) | **Post** /stream-catalog/v2/clusters/{catalog_cluster_id}/notification-topic-configs | Enable stream catalog notifications
[**GetNotificationsStreamCatalogV2NotificationTopicConfig**](NotificationTopicConfigsStreamCatalogV2Api.md#GetNotificationsStreamCatalogV2NotificationTopicConfig) | **Get** /stream-catalog/v2/clusters/{catalog_cluster_id}/notification-topic-configs | Get catalog notifications settings



## DisableNotificationsStreamCatalogV2NotificationTopicConfig

> DisableNotificationsStreamCatalogV2NotificationTopicConfig(ctx, catalogClusterId).Environment(environment).Execute()

Delete the catalog notifications topic configuration



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
    environment := "env-00000" // string | Scope the operation to the given environment.
    catalogClusterId := "catalogClusterId_example" // string | The Catalog Cluster

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationTopicConfigsStreamCatalogV2Api.DisableNotificationsStreamCatalogV2NotificationTopicConfig(context.Background(), catalogClusterId).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationTopicConfigsStreamCatalogV2Api.DisableNotificationsStreamCatalogV2NotificationTopicConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogClusterId** | **string** | The Catalog Cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableNotificationsStreamCatalogV2NotificationTopicConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


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


## EnableNotificationsStreamCatalogV2NotificationTopicConfig

> StreamCatalogV2NotificationTopicConfig EnableNotificationsStreamCatalogV2NotificationTopicConfig(ctx, catalogClusterId).Environment(environment).StreamCatalogV2NotificationTopicConfig(streamCatalogV2NotificationTopicConfig).Execute()

Enable stream catalog notifications



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
    environment := "env-00000" // string | Scope the operation to the given environment.
    catalogClusterId := "catalogClusterId_example" // string | The Catalog Cluster
    streamCatalogV2NotificationTopicConfig := *openapiclient.NewStreamCatalogV2NotificationTopicConfig() // StreamCatalogV2NotificationTopicConfig |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationTopicConfigsStreamCatalogV2Api.EnableNotificationsStreamCatalogV2NotificationTopicConfig(context.Background(), catalogClusterId).Environment(environment).StreamCatalogV2NotificationTopicConfig(streamCatalogV2NotificationTopicConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationTopicConfigsStreamCatalogV2Api.EnableNotificationsStreamCatalogV2NotificationTopicConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableNotificationsStreamCatalogV2NotificationTopicConfig`: StreamCatalogV2NotificationTopicConfig
    fmt.Fprintf(os.Stdout, "Response from `NotificationTopicConfigsStreamCatalogV2Api.EnableNotificationsStreamCatalogV2NotificationTopicConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogClusterId** | **string** | The Catalog Cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableNotificationsStreamCatalogV2NotificationTopicConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 

 **streamCatalogV2NotificationTopicConfig** | [**StreamCatalogV2NotificationTopicConfig**](StreamCatalogV2NotificationTopicConfig.md) |  | 

### Return type

[**StreamCatalogV2NotificationTopicConfig**](StreamCatalogV2NotificationTopicConfig.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationsStreamCatalogV2NotificationTopicConfig

> StreamCatalogV2NotificationTopicConfig GetNotificationsStreamCatalogV2NotificationTopicConfig(ctx, catalogClusterId).Environment(environment).Execute()

Get catalog notifications settings



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
    environment := "env-00000" // string | Scope the operation to the given environment.
    catalogClusterId := "catalogClusterId_example" // string | The Catalog Cluster

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationTopicConfigsStreamCatalogV2Api.GetNotificationsStreamCatalogV2NotificationTopicConfig(context.Background(), catalogClusterId).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationTopicConfigsStreamCatalogV2Api.GetNotificationsStreamCatalogV2NotificationTopicConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsStreamCatalogV2NotificationTopicConfig`: StreamCatalogV2NotificationTopicConfig
    fmt.Fprintf(os.Stdout, "Response from `NotificationTopicConfigsStreamCatalogV2Api.GetNotificationsStreamCatalogV2NotificationTopicConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogClusterId** | **string** | The Catalog Cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsStreamCatalogV2NotificationTopicConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**StreamCatalogV2NotificationTopicConfig**](StreamCatalogV2NotificationTopicConfig.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

