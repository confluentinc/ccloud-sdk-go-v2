# \ToolsSqlV1Api

All URIs are relative to *https://flink.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSqlv1Tool**](ToolsSqlV1Api.md#CreateSqlv1Tool) | **Post** /sql/v1/organizations/{organization_id}/environments/{environment_id}/databases/{database_name}/tools | Create a Tool
[**DeleteSqlv1Tool**](ToolsSqlV1Api.md#DeleteSqlv1Tool) | **Delete** /sql/v1/organizations/{organization_id}/environments/{environment_id}/databases/{database_name}/tools/{tool_name} | Delete a Tool
[**GetSqlv1Tool**](ToolsSqlV1Api.md#GetSqlv1Tool) | **Get** /sql/v1/organizations/{organization_id}/environments/{environment_id}/databases/{database_name}/tools/{tool_name} | Read a Tool
[**ListSqlv1Tools**](ToolsSqlV1Api.md#ListSqlv1Tools) | **Get** /sql/v1/organizations/{organization_id}/environments/{environment_id}/databases/{database_name}/tools | List of Tools



## CreateSqlv1Tool

> SqlV1Tool CreateSqlv1Tool(ctx, organizationId, environmentId, databaseName).SqlV1Tool(sqlV1Tool).Execute()

Create a Tool



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
    databaseName := "databaseName_example" // string | The name of the database.
    sqlV1Tool := *openapiclient.NewSqlV1Tool() // SqlV1Tool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ToolsSqlV1Api.CreateSqlv1Tool(context.Background(), organizationId, environmentId, databaseName).SqlV1Tool(sqlV1Tool).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ToolsSqlV1Api.CreateSqlv1Tool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSqlv1Tool`: SqlV1Tool
    fmt.Fprintf(os.Stdout, "Response from `ToolsSqlV1Api.CreateSqlv1Tool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 
**databaseName** | **string** | The name of the database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSqlv1ToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sqlV1Tool** | [**SqlV1Tool**](SqlV1Tool.md) |  | 

### Return type

[**SqlV1Tool**](SqlV1Tool.md)

### Authorization

[global-api-key](../README.md#global-api-key), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSqlv1Tool

> DeleteSqlv1Tool(ctx, organizationId, environmentId, databaseName, toolName).Execute()

Delete a Tool



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
    databaseName := "databaseName_example" // string | The name of the database.
    toolName := "toolName_example" // string | The user provided name of the Tool.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ToolsSqlV1Api.DeleteSqlv1Tool(context.Background(), organizationId, environmentId, databaseName, toolName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ToolsSqlV1Api.DeleteSqlv1Tool``: %v\n", err)
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
**databaseName** | **string** | The name of the database. | 
**toolName** | **string** | The user provided name of the Tool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSqlv1ToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[global-api-key](../README.md#global-api-key), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSqlv1Tool

> SqlV1Tool GetSqlv1Tool(ctx, organizationId, environmentId, databaseName, toolName).Execute()

Read a Tool



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
    databaseName := "databaseName_example" // string | The name of the database.
    toolName := "toolName_example" // string | The user provided name of the Tool.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ToolsSqlV1Api.GetSqlv1Tool(context.Background(), organizationId, environmentId, databaseName, toolName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ToolsSqlV1Api.GetSqlv1Tool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSqlv1Tool`: SqlV1Tool
    fmt.Fprintf(os.Stdout, "Response from `ToolsSqlV1Api.GetSqlv1Tool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 
**databaseName** | **string** | The name of the database. | 
**toolName** | **string** | The user provided name of the Tool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSqlv1ToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**SqlV1Tool**](SqlV1Tool.md)

### Authorization

[global-api-key](../README.md#global-api-key), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSqlv1Tools

> SqlV1ToolList ListSqlv1Tools(ctx, organizationId, environmentId, databaseName).PageSize(pageSize).PageToken(pageToken).Execute()

List of Tools



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
    databaseName := "databaseName_example" // string | The name of the database.
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ToolsSqlV1Api.ListSqlv1Tools(context.Background(), organizationId, environmentId, databaseName).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ToolsSqlV1Api.ListSqlv1Tools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSqlv1Tools`: SqlV1ToolList
    fmt.Fprintf(os.Stdout, "Response from `ToolsSqlV1Api.ListSqlv1Tools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 
**databaseName** | **string** | The name of the database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSqlv1ToolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**SqlV1ToolList**](SqlV1ToolList.md)

### Authorization

[global-api-key](../README.md#global-api-key), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

