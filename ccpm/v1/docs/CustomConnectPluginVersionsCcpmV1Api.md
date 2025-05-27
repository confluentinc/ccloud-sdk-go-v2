# \CustomConnectPluginVersionsCcpmV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCcpmV1CustomConnectPluginVersion**](CustomConnectPluginVersionsCcpmV1Api.md#CreateCcpmV1CustomConnectPluginVersion) | **Post** /ccpm/v1/plugins/{plugin_id}/versions | Create a Custom Connect Plugin Version
[**DeleteCcpmV1CustomConnectPluginVersion**](CustomConnectPluginVersionsCcpmV1Api.md#DeleteCcpmV1CustomConnectPluginVersion) | **Delete** /ccpm/v1/plugins/{plugin_id}/versions/{id} | Delete a Custom Connect Plugin Version
[**GetCcpmV1CustomConnectPluginVersion**](CustomConnectPluginVersionsCcpmV1Api.md#GetCcpmV1CustomConnectPluginVersion) | **Get** /ccpm/v1/plugins/{plugin_id}/versions/{id} | Read a Custom Connect Plugin Version
[**ListCcpmV1CustomConnectPluginVersions**](CustomConnectPluginVersionsCcpmV1Api.md#ListCcpmV1CustomConnectPluginVersions) | **Get** /ccpm/v1/plugins/{plugin_id}/versions | List of Custom Connect Plugin Versions



## CreateCcpmV1CustomConnectPluginVersion

> CcpmV1CustomConnectPluginVersion CreateCcpmV1CustomConnectPluginVersion(ctx, pluginId).CcpmV1CustomConnectPluginVersion(ccpmV1CustomConnectPluginVersion).Execute()

Create a Custom Connect Plugin Version



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
    ccpmV1CustomConnectPluginVersion := *openapiclient.NewCcpmV1CustomConnectPluginVersion() // CcpmV1CustomConnectPluginVersion |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomConnectPluginVersionsCcpmV1Api.CreateCcpmV1CustomConnectPluginVersion(context.Background(), pluginId).CcpmV1CustomConnectPluginVersion(ccpmV1CustomConnectPluginVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomConnectPluginVersionsCcpmV1Api.CreateCcpmV1CustomConnectPluginVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCcpmV1CustomConnectPluginVersion`: CcpmV1CustomConnectPluginVersion
    fmt.Fprintf(os.Stdout, "Response from `CustomConnectPluginVersionsCcpmV1Api.CreateCcpmV1CustomConnectPluginVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | The Plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCcpmV1CustomConnectPluginVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ccpmV1CustomConnectPluginVersion** | [**CcpmV1CustomConnectPluginVersion**](CcpmV1CustomConnectPluginVersion.md) |  | 

### Return type

[**CcpmV1CustomConnectPluginVersion**](ccpm.v1.CustomConnectPluginVersion.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCcpmV1CustomConnectPluginVersion

> DeleteCcpmV1CustomConnectPluginVersion(ctx, pluginId, id).Environment(environment).Execute()

Delete a Custom Connect Plugin Version



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
    pluginId := "pluginId_example" // string | The Plugin
    id := "id_example" // string | The unique identifier for the custom connect plugin version.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomConnectPluginVersionsCcpmV1Api.DeleteCcpmV1CustomConnectPluginVersion(context.Background(), pluginId, id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomConnectPluginVersionsCcpmV1Api.DeleteCcpmV1CustomConnectPluginVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | The Plugin | 
**id** | **string** | The unique identifier for the custom connect plugin version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCcpmV1CustomConnectPluginVersionRequest struct via the builder pattern


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


## GetCcpmV1CustomConnectPluginVersion

> CcpmV1CustomConnectPluginVersion GetCcpmV1CustomConnectPluginVersion(ctx, pluginId, id).Environment(environment).Execute()

Read a Custom Connect Plugin Version



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
    pluginId := "pluginId_example" // string | The Plugin
    id := "id_example" // string | The unique identifier for the custom connect plugin version.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomConnectPluginVersionsCcpmV1Api.GetCcpmV1CustomConnectPluginVersion(context.Background(), pluginId, id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomConnectPluginVersionsCcpmV1Api.GetCcpmV1CustomConnectPluginVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCcpmV1CustomConnectPluginVersion`: CcpmV1CustomConnectPluginVersion
    fmt.Fprintf(os.Stdout, "Response from `CustomConnectPluginVersionsCcpmV1Api.GetCcpmV1CustomConnectPluginVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | The Plugin | 
**id** | **string** | The unique identifier for the custom connect plugin version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCcpmV1CustomConnectPluginVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 



### Return type

[**CcpmV1CustomConnectPluginVersion**](ccpm.v1.CustomConnectPluginVersion.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCcpmV1CustomConnectPluginVersions

> CcpmV1CustomConnectPluginVersionList ListCcpmV1CustomConnectPluginVersions(ctx, pluginId).Environment(environment).Execute()

List of Custom Connect Plugin Versions



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
    environment := "env-00000" // string | Filter the results by exact match for environment.
    pluginId := "pluginId_example" // string | The Plugin

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomConnectPluginVersionsCcpmV1Api.ListCcpmV1CustomConnectPluginVersions(context.Background(), pluginId).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomConnectPluginVersionsCcpmV1Api.ListCcpmV1CustomConnectPluginVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCcpmV1CustomConnectPluginVersions`: CcpmV1CustomConnectPluginVersionList
    fmt.Fprintf(os.Stdout, "Response from `CustomConnectPluginVersionsCcpmV1Api.ListCcpmV1CustomConnectPluginVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | The Plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCcpmV1CustomConnectPluginVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 


### Return type

[**CcpmV1CustomConnectPluginVersionList**](ccpm.v1.CustomConnectPluginVersionList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

