# \CustomPluginsCcpV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCcpV1CustomPlugin**](CustomPluginsCcpV1Api.md#CreateCcpV1CustomPlugin) | **Post** /ccp/v1/custom-plugins | Create a Custom Plugin
[**DeleteCcpV1CustomPlugin**](CustomPluginsCcpV1Api.md#DeleteCcpV1CustomPlugin) | **Delete** /ccp/v1/custom-plugins/{id} | Delete a Custom Plugin
[**GetCcpV1CustomPlugin**](CustomPluginsCcpV1Api.md#GetCcpV1CustomPlugin) | **Get** /ccp/v1/custom-plugins/{id} | Read a Custom Plugin
[**ListCcpV1CustomPlugins**](CustomPluginsCcpV1Api.md#ListCcpV1CustomPlugins) | **Get** /ccp/v1/custom-plugins | List of Custom Plugins
[**UpdateCcpV1CustomPlugin**](CustomPluginsCcpV1Api.md#UpdateCcpV1CustomPlugin) | **Patch** /ccp/v1/custom-plugins/{id} | Update a Custom Plugin



## CreateCcpV1CustomPlugin

> CcpV1CustomPlugin CreateCcpV1CustomPlugin(ctx).CcpV1CustomPlugin(ccpV1CustomPlugin).Execute()

Create a Custom Plugin



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
    ccpV1CustomPlugin := *openapiclient.NewCcpV1CustomPlugin() // CcpV1CustomPlugin |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomPluginsCcpV1Api.CreateCcpV1CustomPlugin(context.Background()).CcpV1CustomPlugin(ccpV1CustomPlugin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPluginsCcpV1Api.CreateCcpV1CustomPlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCcpV1CustomPlugin`: CcpV1CustomPlugin
    fmt.Fprintf(os.Stdout, "Response from `CustomPluginsCcpV1Api.CreateCcpV1CustomPlugin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCcpV1CustomPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccpV1CustomPlugin** | [**CcpV1CustomPlugin**](CcpV1CustomPlugin.md) |  | 

### Return type

[**CcpV1CustomPlugin**](ccp.v1.CustomPlugin.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCcpV1CustomPlugin

> DeleteCcpV1CustomPlugin(ctx, id).Execute()

Delete a Custom Plugin



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
    id := "id_example" // string | The unique identifier for the custom plugin.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomPluginsCcpV1Api.DeleteCcpV1CustomPlugin(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPluginsCcpV1Api.DeleteCcpV1CustomPlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the custom plugin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCcpV1CustomPluginRequest struct via the builder pattern


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


## GetCcpV1CustomPlugin

> CcpV1CustomPlugin GetCcpV1CustomPlugin(ctx, id).Execute()

Read a Custom Plugin



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
    id := "id_example" // string | The unique identifier for the custom plugin.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomPluginsCcpV1Api.GetCcpV1CustomPlugin(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPluginsCcpV1Api.GetCcpV1CustomPlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCcpV1CustomPlugin`: CcpV1CustomPlugin
    fmt.Fprintf(os.Stdout, "Response from `CustomPluginsCcpV1Api.GetCcpV1CustomPlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the custom plugin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCcpV1CustomPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CcpV1CustomPlugin**](ccp.v1.CustomPlugin.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCcpV1CustomPlugins

> CcpV1CustomPluginList ListCcpV1CustomPlugins(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Custom Plugins



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
    resp, r, err := api_client.CustomPluginsCcpV1Api.ListCcpV1CustomPlugins(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPluginsCcpV1Api.ListCcpV1CustomPlugins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCcpV1CustomPlugins`: CcpV1CustomPluginList
    fmt.Fprintf(os.Stdout, "Response from `CustomPluginsCcpV1Api.ListCcpV1CustomPlugins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCcpV1CustomPluginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**CcpV1CustomPluginList**](ccp.v1.CustomPluginList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCcpV1CustomPlugin

> CcpV1CustomPlugin UpdateCcpV1CustomPlugin(ctx, id).CcpV1CustomPlugin(ccpV1CustomPlugin).Execute()

Update a Custom Plugin



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
    id := "id_example" // string | The unique identifier for the custom plugin.
    ccpV1CustomPlugin := *openapiclient.NewCcpV1CustomPlugin() // CcpV1CustomPlugin |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomPluginsCcpV1Api.UpdateCcpV1CustomPlugin(context.Background(), id).CcpV1CustomPlugin(ccpV1CustomPlugin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPluginsCcpV1Api.UpdateCcpV1CustomPlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCcpV1CustomPlugin`: CcpV1CustomPlugin
    fmt.Fprintf(os.Stdout, "Response from `CustomPluginsCcpV1Api.UpdateCcpV1CustomPlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the custom plugin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCcpV1CustomPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ccpV1CustomPlugin** | [**CcpV1CustomPlugin**](CcpV1CustomPlugin.md) |  | 

### Return type

[**CcpV1CustomPlugin**](ccp.v1.CustomPlugin.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

