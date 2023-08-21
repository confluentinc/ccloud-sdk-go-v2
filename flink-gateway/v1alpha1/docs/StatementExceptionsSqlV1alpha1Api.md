# \StatementExceptionsSqlV1alpha1Api

All URIs are relative to *https://flink.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSqlV1alpha1StatementExceptions**](StatementExceptionsSqlV1alpha1Api.md#GetSqlV1alpha1StatementExceptions) | **Get** /sql/v1alpha1/environments/{environment_id}/statements/{statement_name}/exceptions | List of Statement Exceptions



## GetSqlV1alpha1StatementExceptions

> SqlV1alpha1StatementExceptionList GetSqlV1alpha1StatementExceptions(ctx, environmentId, statementName).OrgId(orgId).Execute()

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
    orgId := TODO // string | The unique identifier for the organization.
    environmentId := "environmentId_example" // string | The unique identifier for the environment.
    statementName := "statementName_example" // string | The unique identifier for the statement.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatementExceptionsSqlV1alpha1Api.GetSqlV1alpha1StatementExceptions(context.Background(), environmentId, statementName).OrgId(orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatementExceptionsSqlV1alpha1Api.GetSqlV1alpha1StatementExceptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSqlV1alpha1StatementExceptions`: SqlV1alpha1StatementExceptionList
    fmt.Fprintf(os.Stdout, "Response from `StatementExceptionsSqlV1alpha1Api.GetSqlV1alpha1StatementExceptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The unique identifier for the environment. | 
**statementName** | **string** | The unique identifier for the statement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSqlV1alpha1StatementExceptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | [**string**](string.md) | The unique identifier for the organization. | 



### Return type

[**SqlV1alpha1StatementExceptionList**](SqlV1alpha1StatementExceptionList.md)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

