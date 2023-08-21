# \StatementsSqlV1beta1Api

All URIs are relative to *https://flink.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSqlv1beta1Statement**](StatementsSqlV1beta1Api.md#CreateSqlv1beta1Statement) | **Post** /sql/v1beta1/organizations/{organization_id}/environments/{environment_id}/statements | Create a Statement
[**DeleteSqlv1beta1Statement**](StatementsSqlV1beta1Api.md#DeleteSqlv1beta1Statement) | **Delete** /sql/v1beta1/organizations/{organization_id}/environments/{environment_id}/statements/{statement_name} | Delete a Statement
[**GetSqlv1beta1Statement**](StatementsSqlV1beta1Api.md#GetSqlv1beta1Statement) | **Get** /sql/v1beta1/organizations/{organization_id}/environments/{environment_id}/statements/{statement_name} | Read a Statement
[**ListSqlv1beta1Statements**](StatementsSqlV1beta1Api.md#ListSqlv1beta1Statements) | **Get** /sql/v1beta1/organizations/{organization_id}/environments/{environment_id}/statements | List of Statements



## CreateSqlv1beta1Statement

> SqlV1beta1Statement CreateSqlv1beta1Statement(ctx, organizationId, environmentId).SqlV1beta1Statement(sqlV1beta1Statement).Execute()

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
    sqlV1beta1Statement := *openapiclient.NewSqlV1beta1Statement() // SqlV1beta1Statement |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatementsSqlV1beta1Api.CreateSqlv1beta1Statement(context.Background(), organizationId, environmentId).SqlV1beta1Statement(sqlV1beta1Statement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatementsSqlV1beta1Api.CreateSqlv1beta1Statement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSqlv1beta1Statement`: SqlV1beta1Statement
    fmt.Fprintf(os.Stdout, "Response from `StatementsSqlV1beta1Api.CreateSqlv1beta1Statement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSqlv1beta1StatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sqlV1beta1Statement** | [**SqlV1beta1Statement**](SqlV1beta1Statement.md) |  | 

### Return type

[**SqlV1beta1Statement**](sql.v1beta1.Statement.md)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSqlv1beta1Statement

> DeleteSqlv1beta1Statement(ctx, organizationId, environmentId, statementName).Execute()

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
    resp, r, err := api_client.StatementsSqlV1beta1Api.DeleteSqlv1beta1Statement(context.Background(), organizationId, environmentId, statementName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatementsSqlV1beta1Api.DeleteSqlv1beta1Statement``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSqlv1beta1StatementRequest struct via the builder pattern


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


## GetSqlv1beta1Statement

> SqlV1beta1Statement GetSqlv1beta1Statement(ctx, organizationId, environmentId, statementName).Execute()

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
    resp, r, err := api_client.StatementsSqlV1beta1Api.GetSqlv1beta1Statement(context.Background(), organizationId, environmentId, statementName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatementsSqlV1beta1Api.GetSqlv1beta1Statement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSqlv1beta1Statement`: SqlV1beta1Statement
    fmt.Fprintf(os.Stdout, "Response from `StatementsSqlV1beta1Api.GetSqlv1beta1Statement`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetSqlv1beta1StatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SqlV1beta1Statement**](sql.v1beta1.Statement.md)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSqlv1beta1Statements

> SqlV1beta1StatementList ListSqlv1beta1Statements(ctx, organizationId, environmentId).SpecComputePoolId(specComputePoolId).PageSize(pageSize).PageToken(pageToken).Execute()

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
    specComputePoolId := "lfcp-00000" // string | Filter the results by exact match for spec.compute_pool. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatementsSqlV1beta1Api.ListSqlv1beta1Statements(context.Background(), organizationId, environmentId).SpecComputePoolId(specComputePoolId).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatementsSqlV1beta1Api.ListSqlv1beta1Statements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSqlv1beta1Statements`: SqlV1beta1StatementList
    fmt.Fprintf(os.Stdout, "Response from `StatementsSqlV1beta1Api.ListSqlv1beta1Statements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSqlv1beta1StatementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **specComputePoolId** | **string** | Filter the results by exact match for spec.compute_pool. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**SqlV1beta1StatementList**](SqlV1beta1StatementList.md)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

