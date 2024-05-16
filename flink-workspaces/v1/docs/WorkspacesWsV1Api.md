# \WorkspacesWsV1Api

All URIs are relative to *https://flink.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWsV1Workspace**](WorkspacesWsV1Api.md#CreateWsV1Workspace) | **Post** /ws/v1/organizations/{organization_id}/environments/{environment_id}/workspaces | Create a Workspace
[**DeleteWsV1Workspace**](WorkspacesWsV1Api.md#DeleteWsV1Workspace) | **Delete** /ws/v1/organizations/{organization_id}/environments/{environment_id}/workspaces/{name} | Delete a Workspace
[**GetWsV1Workspace**](WorkspacesWsV1Api.md#GetWsV1Workspace) | **Get** /ws/v1/organizations/{organization_id}/environments/{environment_id}/workspaces/{name} | Read a Workspace
[**ListWsV1Workspaces**](WorkspacesWsV1Api.md#ListWsV1Workspaces) | **Get** /ws/v1/organizations/{organization_id}/environments/{environment_id}/workspaces | List of Workspaces
[**PatchWsV1Workspace**](WorkspacesWsV1Api.md#PatchWsV1Workspace) | **Patch** /ws/v1/organizations/{organization_id}/environments/{environment_id}/workspaces/{name} | Update a Workspace with patch request
[**PutWsV1Workspace**](WorkspacesWsV1Api.md#PutWsV1Workspace) | **Put** /ws/v1/organizations/{organization_id}/environments/{environment_id}/workspaces/{name} | Update a Workspace



## CreateWsV1Workspace

> WsV1Workspace CreateWsV1Workspace(ctx, organizationId, environmentId).WsV1Workspace(wsV1Workspace).Execute()

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
    wsV1Workspace := *openapiclient.NewWsV1Workspace() // WsV1Workspace |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkspacesWsV1Api.CreateWsV1Workspace(context.Background(), organizationId, environmentId).WsV1Workspace(wsV1Workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesWsV1Api.CreateWsV1Workspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWsV1Workspace`: WsV1Workspace
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesWsV1Api.CreateWsV1Workspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWsV1WorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wsV1Workspace** | [**WsV1Workspace**](WsV1Workspace.md) |  | 

### Return type

[**WsV1Workspace**](ws.v1.Workspace.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWsV1Workspace

> DeleteWsV1Workspace(ctx, organizationId, environmentId, name).Execute()

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
    resp, r, err := api_client.WorkspacesWsV1Api.DeleteWsV1Workspace(context.Background(), organizationId, environmentId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesWsV1Api.DeleteWsV1Workspace``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteWsV1WorkspaceRequest struct via the builder pattern


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


## GetWsV1Workspace

> WsV1Workspace GetWsV1Workspace(ctx, organizationId, environmentId, name).Execute()

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
    resp, r, err := api_client.WorkspacesWsV1Api.GetWsV1Workspace(context.Background(), organizationId, environmentId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesWsV1Api.GetWsV1Workspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWsV1Workspace`: WsV1Workspace
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesWsV1Api.GetWsV1Workspace`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetWsV1WorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WsV1Workspace**](ws.v1.Workspace.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWsV1Workspaces

> WsV1WorkspaceList ListWsV1Workspaces(ctx, organizationId, environmentId).SpecComputePool(specComputePool).All(all).PageSize(pageSize).PageToken(pageToken).Execute()

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
    all := true // bool | If true **and** the requester has either OrgAdmin or EnvAdmin roles, then return all workspaces in the environment. Otherwise return only those created by the caller (default). (optional) (default to false)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkspacesWsV1Api.ListWsV1Workspaces(context.Background(), organizationId, environmentId).SpecComputePool(specComputePool).All(all).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesWsV1Api.ListWsV1Workspaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWsV1Workspaces`: WsV1WorkspaceList
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesWsV1Api.ListWsV1Workspaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWsV1WorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **specComputePool** | **string** | Filter the results by exact match for spec.compute_pool.id. | 
 **all** | **bool** | If true **and** the requester has either OrgAdmin or EnvAdmin roles, then return all workspaces in the environment. Otherwise return only those created by the caller (default). | [default to false]
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**WsV1WorkspaceList**](WsV1WorkspaceList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchWsV1Workspace

> WsV1Workspace PatchWsV1Workspace(ctx, organizationId, environmentId, name).PatchRequest(patchRequest).Execute()

Update a Workspace with patch request



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
    patchRequest := *openapiclient.NewPatchRequest() // PatchRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkspacesWsV1Api.PatchWsV1Workspace(context.Background(), organizationId, environmentId, name).PatchRequest(patchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesWsV1Api.PatchWsV1Workspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchWsV1Workspace`: WsV1Workspace
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesWsV1Api.PatchWsV1Workspace`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPatchWsV1WorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchRequest** | [**PatchRequest**](PatchRequest.md) |  | 

### Return type

[**WsV1Workspace**](ws.v1.Workspace.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutWsV1Workspace

> WsV1Workspace PutWsV1Workspace(ctx, organizationId, environmentId, name).WsV1Workspace(wsV1Workspace).Execute()

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
    wsV1Workspace := *openapiclient.NewWsV1Workspace() // WsV1Workspace |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkspacesWsV1Api.PutWsV1Workspace(context.Background(), organizationId, environmentId, name).WsV1Workspace(wsV1Workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesWsV1Api.PutWsV1Workspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutWsV1Workspace`: WsV1Workspace
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesWsV1Api.PutWsV1Workspace`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPutWsV1WorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **wsV1Workspace** | [**WsV1Workspace**](WsV1Workspace.md) |  | 

### Return type

[**WsV1Workspace**](ws.v1.Workspace.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

