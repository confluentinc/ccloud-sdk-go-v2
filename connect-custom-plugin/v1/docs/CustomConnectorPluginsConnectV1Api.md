# \CustomConnectorPluginsConnectV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnectV1CustomConnectorPlugin**](CustomConnectorPluginsConnectV1Api.md#CreateConnectV1CustomConnectorPlugin) | **Post** /connect/v1/custom-connector-plugins | Create a Custom Connector Plugin
[**DeleteConnectV1CustomConnectorPlugin**](CustomConnectorPluginsConnectV1Api.md#DeleteConnectV1CustomConnectorPlugin) | **Delete** /connect/v1/custom-connector-plugins/{id} | Delete a Custom Connector Plugin
[**GetConnectV1CustomConnectorPlugin**](CustomConnectorPluginsConnectV1Api.md#GetConnectV1CustomConnectorPlugin) | **Get** /connect/v1/custom-connector-plugins/{id} | Read a Custom Connector Plugin
[**ListConnectV1CustomConnectorPlugins**](CustomConnectorPluginsConnectV1Api.md#ListConnectV1CustomConnectorPlugins) | **Get** /connect/v1/custom-connector-plugins | List of Custom Connector Plugins
[**UpdateConnectV1CustomConnectorPlugin**](CustomConnectorPluginsConnectV1Api.md#UpdateConnectV1CustomConnectorPlugin) | **Patch** /connect/v1/custom-connector-plugins/{id} | Update a Custom Connector Plugin



## CreateConnectV1CustomConnectorPlugin

> ConnectV1CustomConnectorPlugin CreateConnectV1CustomConnectorPlugin(ctx).ConnectV1CustomConnectorPlugin(connectV1CustomConnectorPlugin).Execute()

Create a Custom Connector Plugin



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
    connectV1CustomConnectorPlugin := *openapiclient.NewConnectV1CustomConnectorPlugin() // ConnectV1CustomConnectorPlugin |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomConnectorPluginsConnectV1Api.CreateConnectV1CustomConnectorPlugin(context.Background()).ConnectV1CustomConnectorPlugin(connectV1CustomConnectorPlugin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomConnectorPluginsConnectV1Api.CreateConnectV1CustomConnectorPlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConnectV1CustomConnectorPlugin`: ConnectV1CustomConnectorPlugin
    fmt.Fprintf(os.Stdout, "Response from `CustomConnectorPluginsConnectV1Api.CreateConnectV1CustomConnectorPlugin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectV1CustomConnectorPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectV1CustomConnectorPlugin** | [**ConnectV1CustomConnectorPlugin**](ConnectV1CustomConnectorPlugin.md) |  | 

### Return type

[**ConnectV1CustomConnectorPlugin**](connect.v1.CustomConnectorPlugin.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnectV1CustomConnectorPlugin

> DeleteConnectV1CustomConnectorPlugin(ctx, id).Execute()

Delete a Custom Connector Plugin



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
    id := "id_example" // string | The unique identifier for the custom connector plugin.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomConnectorPluginsConnectV1Api.DeleteConnectV1CustomConnectorPlugin(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomConnectorPluginsConnectV1Api.DeleteConnectV1CustomConnectorPlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the custom connector plugin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectV1CustomConnectorPluginRequest struct via the builder pattern


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


## GetConnectV1CustomConnectorPlugin

> ConnectV1CustomConnectorPlugin GetConnectV1CustomConnectorPlugin(ctx, id).Execute()

Read a Custom Connector Plugin



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
    id := "id_example" // string | The unique identifier for the custom connector plugin.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomConnectorPluginsConnectV1Api.GetConnectV1CustomConnectorPlugin(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomConnectorPluginsConnectV1Api.GetConnectV1CustomConnectorPlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectV1CustomConnectorPlugin`: ConnectV1CustomConnectorPlugin
    fmt.Fprintf(os.Stdout, "Response from `CustomConnectorPluginsConnectV1Api.GetConnectV1CustomConnectorPlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the custom connector plugin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectV1CustomConnectorPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectV1CustomConnectorPlugin**](connect.v1.CustomConnectorPlugin.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectV1CustomConnectorPlugins

> ConnectV1CustomConnectorPluginList ListConnectV1CustomConnectorPlugins(ctx).Cloud(cloud).PageSize(pageSize).PageToken(pageToken).Execute()

List of Custom Connector Plugins



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
    cloud := "AWS" // string | Filter the results by exact match for cloud. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomConnectorPluginsConnectV1Api.ListConnectV1CustomConnectorPlugins(context.Background()).Cloud(cloud).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomConnectorPluginsConnectV1Api.ListConnectV1CustomConnectorPlugins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectV1CustomConnectorPlugins`: ConnectV1CustomConnectorPluginList
    fmt.Fprintf(os.Stdout, "Response from `CustomConnectorPluginsConnectV1Api.ListConnectV1CustomConnectorPlugins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectV1CustomConnectorPluginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud** | **string** | Filter the results by exact match for cloud. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**ConnectV1CustomConnectorPluginList**](connect.v1.CustomConnectorPluginList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectV1CustomConnectorPlugin

> ConnectV1CustomConnectorPlugin UpdateConnectV1CustomConnectorPlugin(ctx, id).ConnectV1CustomConnectorPluginUpdate(connectV1CustomConnectorPluginUpdate).Execute()

Update a Custom Connector Plugin



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
    id := "id_example" // string | The unique identifier for the custom connector plugin.
    connectV1CustomConnectorPluginUpdate := *openapiclient.NewConnectV1CustomConnectorPluginUpdate() // ConnectV1CustomConnectorPluginUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomConnectorPluginsConnectV1Api.UpdateConnectV1CustomConnectorPlugin(context.Background(), id).ConnectV1CustomConnectorPluginUpdate(connectV1CustomConnectorPluginUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomConnectorPluginsConnectV1Api.UpdateConnectV1CustomConnectorPlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConnectV1CustomConnectorPlugin`: ConnectV1CustomConnectorPlugin
    fmt.Fprintf(os.Stdout, "Response from `CustomConnectorPluginsConnectV1Api.UpdateConnectV1CustomConnectorPlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the custom connector plugin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectV1CustomConnectorPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectV1CustomConnectorPluginUpdate** | [**ConnectV1CustomConnectorPluginUpdate**](ConnectV1CustomConnectorPluginUpdate.md) |  | 

### Return type

[**ConnectV1CustomConnectorPlugin**](connect.v1.CustomConnectorPlugin.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

