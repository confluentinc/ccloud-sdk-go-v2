# \CustomConnectorPluginVersionsConnectV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnectV1CustomConnectorPluginVersion**](CustomConnectorPluginVersionsConnectV1Api.md#CreateConnectV1CustomConnectorPluginVersion) | **Post** /connect/v1/custom-connector-plugins/{plugin_id}/versions | Create a Custom Connector Plugin Version
[**DeleteConnectV1CustomConnectorPluginVersion**](CustomConnectorPluginVersionsConnectV1Api.md#DeleteConnectV1CustomConnectorPluginVersion) | **Delete** /connect/v1/custom-connector-plugins/{plugin_id}/versions/{id} | Delete a Custom Connector Plugin Version
[**GetConnectV1CustomConnectorPluginVersion**](CustomConnectorPluginVersionsConnectV1Api.md#GetConnectV1CustomConnectorPluginVersion) | **Get** /connect/v1/custom-connector-plugins/{plugin_id}/versions/{id} | Read a Custom Connector Plugin Version
[**ListConnectV1CustomConnectorPluginVersions**](CustomConnectorPluginVersionsConnectV1Api.md#ListConnectV1CustomConnectorPluginVersions) | **Get** /connect/v1/custom-connector-plugins/{plugin_id}/versions | List of Custom Connector Plugin Versions
[**UpdateConnectV1CustomConnectorPluginVersion**](CustomConnectorPluginVersionsConnectV1Api.md#UpdateConnectV1CustomConnectorPluginVersion) | **Patch** /connect/v1/custom-connector-plugins/{plugin_id}/versions/{id} | Update a Custom Connector Plugin Version



## CreateConnectV1CustomConnectorPluginVersion

> ConnectV1CustomConnectorPluginVersion CreateConnectV1CustomConnectorPluginVersion(ctx, pluginId).ConnectV1CustomConnectorPluginVersion(connectV1CustomConnectorPluginVersion).Execute()

Create a Custom Connector Plugin Version



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
    pluginId := "pluginId_example" // string | The Plugin
    connectV1CustomConnectorPluginVersion := *openapiclient.NewConnectV1CustomConnectorPluginVersion() // ConnectV1CustomConnectorPluginVersion |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomConnectorPluginVersionsConnectV1Api.CreateConnectV1CustomConnectorPluginVersion(context.Background(), pluginId).ConnectV1CustomConnectorPluginVersion(connectV1CustomConnectorPluginVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomConnectorPluginVersionsConnectV1Api.CreateConnectV1CustomConnectorPluginVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConnectV1CustomConnectorPluginVersion`: ConnectV1CustomConnectorPluginVersion
    fmt.Fprintf(os.Stdout, "Response from `CustomConnectorPluginVersionsConnectV1Api.CreateConnectV1CustomConnectorPluginVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | The Plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectV1CustomConnectorPluginVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectV1CustomConnectorPluginVersion** | [**ConnectV1CustomConnectorPluginVersion**](ConnectV1CustomConnectorPluginVersion.md) |  | 

### Return type

[**ConnectV1CustomConnectorPluginVersion**](connect.v1.CustomConnectorPluginVersion.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnectV1CustomConnectorPluginVersion

> DeleteConnectV1CustomConnectorPluginVersion(ctx, pluginId, id).Execute()

Delete a Custom Connector Plugin Version



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
    pluginId := "pluginId_example" // string | The Plugin
    id := "id_example" // string | The unique identifier for the custom connector plugin version.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomConnectorPluginVersionsConnectV1Api.DeleteConnectV1CustomConnectorPluginVersion(context.Background(), pluginId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomConnectorPluginVersionsConnectV1Api.DeleteConnectV1CustomConnectorPluginVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | The Plugin | 
**id** | **string** | The unique identifier for the custom connector plugin version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectV1CustomConnectorPluginVersionRequest struct via the builder pattern


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


## GetConnectV1CustomConnectorPluginVersion

> ConnectV1CustomConnectorPluginVersion GetConnectV1CustomConnectorPluginVersion(ctx, pluginId, id).Execute()

Read a Custom Connector Plugin Version



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
    pluginId := "pluginId_example" // string | The Plugin
    id := "id_example" // string | The unique identifier for the custom connector plugin version.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomConnectorPluginVersionsConnectV1Api.GetConnectV1CustomConnectorPluginVersion(context.Background(), pluginId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomConnectorPluginVersionsConnectV1Api.GetConnectV1CustomConnectorPluginVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectV1CustomConnectorPluginVersion`: ConnectV1CustomConnectorPluginVersion
    fmt.Fprintf(os.Stdout, "Response from `CustomConnectorPluginVersionsConnectV1Api.GetConnectV1CustomConnectorPluginVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | The Plugin | 
**id** | **string** | The unique identifier for the custom connector plugin version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectV1CustomConnectorPluginVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConnectV1CustomConnectorPluginVersion**](connect.v1.CustomConnectorPluginVersion.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectV1CustomConnectorPluginVersions

> ConnectV1CustomConnectorPluginVersionList ListConnectV1CustomConnectorPluginVersions(ctx, pluginId).PageSize(pageSize).PageToken(pageToken).Execute()

List of Custom Connector Plugin Versions



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
    pluginId := "pluginId_example" // string | The Plugin
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomConnectorPluginVersionsConnectV1Api.ListConnectV1CustomConnectorPluginVersions(context.Background(), pluginId).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomConnectorPluginVersionsConnectV1Api.ListConnectV1CustomConnectorPluginVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectV1CustomConnectorPluginVersions`: ConnectV1CustomConnectorPluginVersionList
    fmt.Fprintf(os.Stdout, "Response from `CustomConnectorPluginVersionsConnectV1Api.ListConnectV1CustomConnectorPluginVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | The Plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectV1CustomConnectorPluginVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**ConnectV1CustomConnectorPluginVersionList**](connect.v1.CustomConnectorPluginVersionList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectV1CustomConnectorPluginVersion

> ConnectV1CustomConnectorPluginVersion UpdateConnectV1CustomConnectorPluginVersion(ctx, pluginId, id).ConnectV1CustomConnectorPluginVersion(connectV1CustomConnectorPluginVersion).Execute()

Update a Custom Connector Plugin Version



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
    pluginId := "pluginId_example" // string | The Plugin
    id := "id_example" // string | The unique identifier for the custom connector plugin version.
    connectV1CustomConnectorPluginVersion := *openapiclient.NewConnectV1CustomConnectorPluginVersion() // ConnectV1CustomConnectorPluginVersion |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomConnectorPluginVersionsConnectV1Api.UpdateConnectV1CustomConnectorPluginVersion(context.Background(), pluginId, id).ConnectV1CustomConnectorPluginVersion(connectV1CustomConnectorPluginVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomConnectorPluginVersionsConnectV1Api.UpdateConnectV1CustomConnectorPluginVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConnectV1CustomConnectorPluginVersion`: ConnectV1CustomConnectorPluginVersion
    fmt.Fprintf(os.Stdout, "Response from `CustomConnectorPluginVersionsConnectV1Api.UpdateConnectV1CustomConnectorPluginVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | The Plugin | 
**id** | **string** | The unique identifier for the custom connector plugin version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectV1CustomConnectorPluginVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **connectV1CustomConnectorPluginVersion** | [**ConnectV1CustomConnectorPluginVersion**](ConnectV1CustomConnectorPluginVersion.md) |  | 

### Return type

[**ConnectV1CustomConnectorPluginVersion**](connect.v1.CustomConnectorPluginVersion.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

