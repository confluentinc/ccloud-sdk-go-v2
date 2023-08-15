# \WorkspacesWsV1beta1Api

All URIs are relative to *https://flink.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWsV1beta1Workspace**](WorkspacesWsV1beta1Api.md#CreateWsV1beta1Workspace) | **Post** /ws/v1beta1/organizations/{organization_id}/environments/{environment_id}/workspaces | Create a Workspace
[**DeleteWsV1beta1Workspace**](WorkspacesWsV1beta1Api.md#DeleteWsV1beta1Workspace) | **Delete** /ws/v1beta1/organizations/{organization_id}/environments/{environment_id}/workspaces/{name} | Delete a Workspace
[**GetWsV1beta1Workspace**](WorkspacesWsV1beta1Api.md#GetWsV1beta1Workspace) | **Get** /ws/v1beta1/organizations/{organization_id}/environments/{environment_id}/workspaces/{name} | Read a Workspace
[**ListWsV1beta1Workspaces**](WorkspacesWsV1beta1Api.md#ListWsV1beta1Workspaces) | **Get** /ws/v1beta1/organizations/{organization_id}/environments/{environment_id}/workspaces | List of Workspaces
[**PutWsV1beta1Workspace**](WorkspacesWsV1beta1Api.md#PutWsV1beta1Workspace) | **Put** /ws/v1beta1/organizations/{organization_id}/environments/{environment_id}/workspaces/{name} | Update a Workspace



## CreateWsV1beta1Workspace

> WsV1beta1Workspace CreateWsV1beta1Workspace(ctx, organizationId, environmentId).WsV1beta1Workspace(wsV1beta1Workspace).Execute()

Create a Workspace



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
    organizationId := TODO // string | The unique identifier for the organization.
    environmentId := "environmentId_example" // string | The unique identifier for the environment.
    wsV1beta1Workspace := *openapiclient.NewWsV1beta1Workspace() // WsV1beta1Workspace |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkspacesWsV1beta1Api.CreateWsV1beta1Workspace(context.Background(), organizationId, environmentId).WsV1beta1Workspace(wsV1beta1Workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesWsV1beta1Api.CreateWsV1beta1Workspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWsV1beta1Workspace`: WsV1beta1Workspace
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesWsV1beta1Api.CreateWsV1beta1Workspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWsV1beta1WorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wsV1beta1Workspace** | [**WsV1beta1Workspace**](WsV1beta1Workspace.md) |  | 

### Return type

[**WsV1beta1Workspace**](ws.v1beta1.Workspace.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWsV1beta1Workspace

> DeleteWsV1beta1Workspace(ctx, organizationId, environmentId, name).Execute()

Delete a Workspace



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
    organizationId := TODO // string | The unique identifier for the organization.
    environmentId := "environmentId_example" // string | The unique identifier for the environment.
    name := "name_example" // string | The workspace name that is unique across the environment and region.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkspacesWsV1beta1Api.DeleteWsV1beta1Workspace(context.Background(), organizationId, environmentId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesWsV1beta1Api.DeleteWsV1beta1Workspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 
**name** | **string** | The workspace name that is unique across the environment and region. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWsV1beta1WorkspaceRequest struct via the builder pattern


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


## GetWsV1beta1Workspace

> WsV1beta1Workspace GetWsV1beta1Workspace(ctx, organizationId, environmentId, name).Execute()

Read a Workspace



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
    organizationId := TODO // string | The unique identifier for the organization.
    environmentId := "environmentId_example" // string | The unique identifier for the environment.
    name := "name_example" // string | The workspace name that is unique across the environment and region.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkspacesWsV1beta1Api.GetWsV1beta1Workspace(context.Background(), organizationId, environmentId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesWsV1beta1Api.GetWsV1beta1Workspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWsV1beta1Workspace`: WsV1beta1Workspace
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesWsV1beta1Api.GetWsV1beta1Workspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 
**name** | **string** | The workspace name that is unique across the environment and region. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWsV1beta1WorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WsV1beta1Workspace**](ws.v1beta1.Workspace.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWsV1beta1Workspaces

> WsV1beta1WorkspaceList ListWsV1beta1Workspaces(ctx, organizationId, environmentId).SpecComputePool(specComputePool).PageSize(pageSize).PageToken(pageToken).Execute()

List of Workspaces



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
    organizationId := TODO // string | The unique identifier for the organization.
    environmentId := "environmentId_example" // string | The unique identifier for the environment.
    specComputePool := "lfcp-00000" // string | Filter the results by exact match for spec.compute_pool.id. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkspacesWsV1beta1Api.ListWsV1beta1Workspaces(context.Background(), organizationId, environmentId).SpecComputePool(specComputePool).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesWsV1beta1Api.ListWsV1beta1Workspaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWsV1beta1Workspaces`: WsV1beta1WorkspaceList
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesWsV1beta1Api.ListWsV1beta1Workspaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWsV1beta1WorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **specComputePool** | **string** | Filter the results by exact match for spec.compute_pool.id. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**WsV1beta1WorkspaceList**](WsV1beta1WorkspaceList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutWsV1beta1Workspace

> WsV1beta1Workspace PutWsV1beta1Workspace(ctx, organizationId, environmentId, name).WsV1beta1Workspace(wsV1beta1Workspace).Execute()

Update a Workspace



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
    organizationId := TODO // string | The unique identifier for the organization.
    environmentId := "environmentId_example" // string | The unique identifier for the environment.
    name := "name_example" // string | The workspace name that is unique across the environment and region.
    wsV1beta1Workspace := *openapiclient.NewWsV1beta1Workspace() // WsV1beta1Workspace |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkspacesWsV1beta1Api.PutWsV1beta1Workspace(context.Background(), organizationId, environmentId, name).WsV1beta1Workspace(wsV1beta1Workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesWsV1beta1Api.PutWsV1beta1Workspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutWsV1beta1Workspace`: WsV1beta1Workspace
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesWsV1beta1Api.PutWsV1beta1Workspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 
**name** | **string** | The workspace name that is unique across the environment and region. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutWsV1beta1WorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **wsV1beta1Workspace** | [**WsV1beta1Workspace**](WsV1beta1Workspace.md) |  | 

### Return type

[**WsV1beta1Workspace**](ws.v1beta1.Workspace.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

