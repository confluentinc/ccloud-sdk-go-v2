# \StatementExceptionsSqlV1Api

All URIs are relative to *https://flink.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSqlv1StatementExceptions**](StatementExceptionsSqlV1Api.md#GetSqlv1StatementExceptions) | **Get** /sql/v1/organizations/{organization_id}/environments/{environment_id}/statements/{statement_name}/exceptions | List of Statement Exceptions



## GetSqlv1StatementExceptions

> SqlV1StatementExceptionList GetSqlv1StatementExceptions(ctx, organizationId, environmentId, statementName).Execute()

List of Statement Exceptions



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
    resp, r, err := api_client.StatementExceptionsSqlV1Api.GetSqlv1StatementExceptions(context.Background(), organizationId, environmentId, statementName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatementExceptionsSqlV1Api.GetSqlv1StatementExceptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSqlv1StatementExceptions`: SqlV1StatementExceptionList
    fmt.Fprintf(os.Stdout, "Response from `StatementExceptionsSqlV1Api.GetSqlv1StatementExceptions`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetSqlv1StatementExceptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SqlV1StatementExceptionList**](SqlV1StatementExceptionList.md)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

