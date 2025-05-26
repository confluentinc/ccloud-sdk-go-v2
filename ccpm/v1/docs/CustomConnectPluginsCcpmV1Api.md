# \CustomConnectPluginsCcpmV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCcpmV1CustomConnectPlugin**](CustomConnectPluginsCcpmV1Api.md#CreateCcpmV1CustomConnectPlugin) | **Post** /ccpm/v1/plugins | Create a Custom Connect Plugin
[**DeleteCcpmV1CustomConnectPlugin**](CustomConnectPluginsCcpmV1Api.md#DeleteCcpmV1CustomConnectPlugin) | **Delete** /ccpm/v1/plugins/{id} | Delete a Custom Connect Plugin
[**GetCcpmV1CustomConnectPlugin**](CustomConnectPluginsCcpmV1Api.md#GetCcpmV1CustomConnectPlugin) | **Get** /ccpm/v1/plugins/{id} | Read a Custom Connect Plugin
[**ListCcpmV1CustomConnectPlugins**](CustomConnectPluginsCcpmV1Api.md#ListCcpmV1CustomConnectPlugins) | **Get** /ccpm/v1/plugins | List of Custom Connect Plugins
[**UpdateCcpmV1CustomConnectPlugin**](CustomConnectPluginsCcpmV1Api.md#UpdateCcpmV1CustomConnectPlugin) | **Patch** /ccpm/v1/plugins/{id} | Update a Custom Connect Plugin



## CreateCcpmV1CustomConnectPlugin

> CcpmV1CustomConnectPlugin CreateCcpmV1CustomConnectPlugin(ctx).CcpmV1CustomConnectPlugin(ccpmV1CustomConnectPlugin).Execute()

Create a Custom Connect Plugin



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
    ccpmV1CustomConnectPlugin := *openapiclient.NewCcpmV1CustomConnectPlugin() // CcpmV1CustomConnectPlugin |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomConnectPluginsCcpmV1Api.CreateCcpmV1CustomConnectPlugin(context.Background()).CcpmV1CustomConnectPlugin(ccpmV1CustomConnectPlugin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomConnectPluginsCcpmV1Api.CreateCcpmV1CustomConnectPlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCcpmV1CustomConnectPlugin`: CcpmV1CustomConnectPlugin
    fmt.Fprintf(os.Stdout, "Response from `CustomConnectPluginsCcpmV1Api.CreateCcpmV1CustomConnectPlugin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCcpmV1CustomConnectPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccpmV1CustomConnectPlugin** | [**CcpmV1CustomConnectPlugin**](CcpmV1CustomConnectPlugin.md) |  | 

### Return type

[**CcpmV1CustomConnectPlugin**](ccpm.v1.CustomConnectPlugin.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCcpmV1CustomConnectPlugin

> DeleteCcpmV1CustomConnectPlugin(ctx, id).Environment(environment).Execute()

Delete a Custom Connect Plugin



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
    id := "id_example" // string | The unique identifier for the custom connect plugin.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomConnectPluginsCcpmV1Api.DeleteCcpmV1CustomConnectPlugin(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomConnectPluginsCcpmV1Api.DeleteCcpmV1CustomConnectPlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the custom connect plugin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCcpmV1CustomConnectPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

 (empty response body)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCcpmV1CustomConnectPlugin

> CcpmV1CustomConnectPlugin GetCcpmV1CustomConnectPlugin(ctx, id).Environment(environment).Execute()

Read a Custom Connect Plugin



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
    id := "id_example" // string | The unique identifier for the custom connect plugin.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomConnectPluginsCcpmV1Api.GetCcpmV1CustomConnectPlugin(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomConnectPluginsCcpmV1Api.GetCcpmV1CustomConnectPlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCcpmV1CustomConnectPlugin`: CcpmV1CustomConnectPlugin
    fmt.Fprintf(os.Stdout, "Response from `CustomConnectPluginsCcpmV1Api.GetCcpmV1CustomConnectPlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the custom connect plugin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCcpmV1CustomConnectPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**CcpmV1CustomConnectPlugin**](ccpm.v1.CustomConnectPlugin.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCcpmV1CustomConnectPlugins

> CcpmV1CustomConnectPluginList ListCcpmV1CustomConnectPlugins(ctx).Environment(environment).SpecCloud(specCloud).PageSize(pageSize).PageToken(pageToken).Execute()

List of Custom Connect Plugins



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
    specCloud := "AWS" // string | Filter the results by exact match for spec.cloud. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomConnectPluginsCcpmV1Api.ListCcpmV1CustomConnectPlugins(context.Background()).Environment(environment).SpecCloud(specCloud).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomConnectPluginsCcpmV1Api.ListCcpmV1CustomConnectPlugins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCcpmV1CustomConnectPlugins`: CcpmV1CustomConnectPluginList
    fmt.Fprintf(os.Stdout, "Response from `CustomConnectPluginsCcpmV1Api.ListCcpmV1CustomConnectPlugins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCcpmV1CustomConnectPluginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **specCloud** | **string** | Filter the results by exact match for spec.cloud. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**CcpmV1CustomConnectPluginList**](ccpm.v1.CustomConnectPluginList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCcpmV1CustomConnectPlugin

> CcpmV1CustomConnectPlugin UpdateCcpmV1CustomConnectPlugin(ctx, id).CcpmV1CustomConnectPluginUpdate(ccpmV1CustomConnectPluginUpdate).Execute()

Update a Custom Connect Plugin



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
    id := "id_example" // string | The unique identifier for the custom connect plugin.
    ccpmV1CustomConnectPluginUpdate := *openapiclient.NewCcpmV1CustomConnectPluginUpdate() // CcpmV1CustomConnectPluginUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomConnectPluginsCcpmV1Api.UpdateCcpmV1CustomConnectPlugin(context.Background(), id).CcpmV1CustomConnectPluginUpdate(ccpmV1CustomConnectPluginUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomConnectPluginsCcpmV1Api.UpdateCcpmV1CustomConnectPlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCcpmV1CustomConnectPlugin`: CcpmV1CustomConnectPlugin
    fmt.Fprintf(os.Stdout, "Response from `CustomConnectPluginsCcpmV1Api.UpdateCcpmV1CustomConnectPlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the custom connect plugin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCcpmV1CustomConnectPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ccpmV1CustomConnectPluginUpdate** | [**CcpmV1CustomConnectPluginUpdate**](CcpmV1CustomConnectPluginUpdate.md) |  | 

### Return type

[**CcpmV1CustomConnectPlugin**](ccpm.v1.CustomConnectPlugin.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

