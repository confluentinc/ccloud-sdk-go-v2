# \StatementsSqlV1Api

All URIs are relative to *https://flink.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSqlv1Statement**](StatementsSqlV1Api.md#CreateSqlv1Statement) | **Post** /sql/v1/organizations/{organization_id}/environments/{environment_id}/statements | Create a Statement
[**DeleteSqlv1Statement**](StatementsSqlV1Api.md#DeleteSqlv1Statement) | **Delete** /sql/v1/organizations/{organization_id}/environments/{environment_id}/statements/{statement_name} | Delete a Statement
[**GetSqlv1Statement**](StatementsSqlV1Api.md#GetSqlv1Statement) | **Get** /sql/v1/organizations/{organization_id}/environments/{environment_id}/statements/{statement_name} | Read a Statement
[**ListSqlv1Statements**](StatementsSqlV1Api.md#ListSqlv1Statements) | **Get** /sql/v1/organizations/{organization_id}/environments/{environment_id}/statements | List of Statements
[**PatchSqlv1Statement**](StatementsSqlV1Api.md#PatchSqlv1Statement) | **Patch** /sql/v1/organizations/{organization_id}/environments/{environment_id}/statements/{statement_name} | Patch a Statement
[**UpdateSqlv1Statement**](StatementsSqlV1Api.md#UpdateSqlv1Statement) | **Put** /sql/v1/organizations/{organization_id}/environments/{environment_id}/statements/{statement_name} | Update a Statement



## CreateSqlv1Statement

> SqlV1Statement CreateSqlv1Statement(ctx, organizationId, environmentId).SqlV1Statement(sqlV1Statement).Execute()

Create a Statement



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
    sqlV1Statement := *openapiclient.NewSqlV1Statement() // SqlV1Statement |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatementsSqlV1Api.CreateSqlv1Statement(context.Background(), organizationId, environmentId).SqlV1Statement(sqlV1Statement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatementsSqlV1Api.CreateSqlv1Statement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSqlv1Statement`: SqlV1Statement
    fmt.Fprintf(os.Stdout, "Response from `StatementsSqlV1Api.CreateSqlv1Statement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSqlv1StatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sqlV1Statement** | [**SqlV1Statement**](SqlV1Statement.md) |  | 

### Return type

[**SqlV1Statement**](sql.v1.Statement.md)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSqlv1Statement

> DeleteSqlv1Statement(ctx, organizationId, environmentId, statementName).Execute()

Delete a Statement



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
    statementName := "statementName_example" // string | The unique identifier for the statement.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatementsSqlV1Api.DeleteSqlv1Statement(context.Background(), organizationId, environmentId, statementName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatementsSqlV1Api.DeleteSqlv1Statement``: %v\n", err)
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
**statementName** | **string** | The unique identifier for the statement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSqlv1StatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSqlv1Statement

> SqlV1Statement GetSqlv1Statement(ctx, organizationId, environmentId, statementName).Execute()

Read a Statement



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
    statementName := "statementName_example" // string | The unique identifier for the statement.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatementsSqlV1Api.GetSqlv1Statement(context.Background(), organizationId, environmentId, statementName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatementsSqlV1Api.GetSqlv1Statement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSqlv1Statement`: SqlV1Statement
    fmt.Fprintf(os.Stdout, "Response from `StatementsSqlV1Api.GetSqlv1Statement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 
**statementName** | **string** | The unique identifier for the statement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSqlv1StatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SqlV1Statement**](sql.v1.Statement.md)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSqlv1Statements

> SqlV1StatementList ListSqlv1Statements(ctx, organizationId, environmentId).SpecComputePoolId(specComputePoolId).PageSize(pageSize).PageToken(pageToken).LabelSelector(labelSelector).Execute()

List of Statements



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
    specComputePoolId := "lfcp-00000" // string | Filter the results by exact match for spec.compute_pool_id. When creating statements, if compute_pool_id is not specified, the statement will use the default compute pool. The default pool is automatically determined by the system. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)
    labelSelector := "labelSelector_example" // string | A comma-separated label selector to filter the statements. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatementsSqlV1Api.ListSqlv1Statements(context.Background(), organizationId, environmentId).SpecComputePoolId(specComputePoolId).PageSize(pageSize).PageToken(pageToken).LabelSelector(labelSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatementsSqlV1Api.ListSqlv1Statements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSqlv1Statements`: SqlV1StatementList
    fmt.Fprintf(os.Stdout, "Response from `StatementsSqlV1Api.ListSqlv1Statements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSqlv1StatementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **specComputePoolId** | **string** | Filter the results by exact match for spec.compute_pool_id. When creating statements, if compute_pool_id is not specified, the statement will use the default compute pool. The default pool is automatically determined by the system. | **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 
 **labelSelector** | **string** | A comma-separated label selector to filter the statements. | 

### Return type

[**SqlV1StatementList**](SqlV1StatementList.md)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSqlv1Statement

> SqlV1Statement PatchSqlv1Statement(ctx, organizationId, environmentId, statementName).PatchRequest(patchRequest).Execute()

Patch a Statement



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
    statementName := "statementName_example" // string | The unique identifier for the statement.
    patchRequest := *openapiclient.NewPatchRequest() // PatchRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatementsSqlV1Api.PatchSqlv1Statement(context.Background(), organizationId, environmentId, statementName).PatchRequest(patchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatementsSqlV1Api.PatchSqlv1Statement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchSqlv1Statement`: SqlV1Statement
    fmt.Fprintf(os.Stdout, "Response from `StatementsSqlV1Api.PatchSqlv1Statement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 
**statementName** | **string** | The unique identifier for the statement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSqlv1StatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchRequest** | [**PatchRequest**](PatchRequest.md) |  | 

### Return type

[**SqlV1Statement**](sql.v1.Statement.md)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSqlv1Statement

> UpdateSqlv1Statement(ctx, organizationId, environmentId, statementName).SqlV1Statement(sqlV1Statement).Execute()

Update a Statement



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
    statementName := "statementName_example" // string | The unique identifier for the statement.
    sqlV1Statement := *openapiclient.NewSqlV1Statement() // SqlV1Statement |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatementsSqlV1Api.UpdateSqlv1Statement(context.Background(), organizationId, environmentId, statementName).SqlV1Statement(sqlV1Statement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatementsSqlV1Api.UpdateSqlv1Statement``: %v\n", err)
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
**statementName** | **string** | The unique identifier for the statement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSqlv1StatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sqlV1Statement** | [**SqlV1Statement**](SqlV1Statement.md) |  | 

### Return type

 (empty response body)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

