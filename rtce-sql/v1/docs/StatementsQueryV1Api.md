# \StatementsQueryV1Api

All URIs are relative to *https://sql.&lt;region&gt;.&lt;provider&gt;.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecuteQueryV1Statement**](StatementsQueryV1Api.md#ExecuteQueryV1Statement) | **Post** /query/v1 | Execute SQL Statement
[**GetQueryV1JobStatus**](StatementsQueryV1Api.md#GetQueryV1JobStatus) | **Get** /query/v1/jobs/{environment_id}/{kafka_cluster_id}/{topic_name}/{statement_id}/status | Query Async Job Status



## ExecuteQueryV1Statement

> QueryV1QueryResponseInline ExecuteQueryV1Statement(ctx).QueryV1QueryRequest(queryV1QueryRequest).Execute()

Execute SQL Statement



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
    queryV1QueryRequest := *openapiclient.NewQueryV1QueryRequest("CatalogName_example", "DatabaseName_example", "Query_example") // QueryV1QueryRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatementsQueryV1Api.ExecuteQueryV1Statement(context.Background()).QueryV1QueryRequest(queryV1QueryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatementsQueryV1Api.ExecuteQueryV1Statement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteQueryV1Statement`: QueryV1QueryResponseInline
    fmt.Fprintf(os.Stdout, "Response from `StatementsQueryV1Api.ExecuteQueryV1Statement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecuteQueryV1StatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryV1QueryRequest** | [**QueryV1QueryRequest**](QueryV1QueryRequest.md) |  | 

### Return type

[**QueryV1QueryResponseInline**](QueryV1QueryResponseInline.md)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueryV1JobStatus

> QueryV1JobStatusResponse GetQueryV1JobStatus(ctx, environmentId, kafkaClusterId, topicName, statementId).Execute()

Query Async Job Status



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
    environmentId := "environmentId_example" // string | Confluent Cloud environment ID (e.g., env-xyz123) the statement was run against.
    kafkaClusterId := "kafkaClusterId_example" // string | Kafka cluster ID (e.g., lkc-xyz123) the statement was run against.
    topicName := "topicName_example" // string | The name of the Kafka topic or Lightning table being queried.
    statementId := "statementId_example" // string | The target unique identifier of the background operational query process.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatementsQueryV1Api.GetQueryV1JobStatus(context.Background(), environmentId, kafkaClusterId, topicName, statementId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatementsQueryV1Api.GetQueryV1JobStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQueryV1JobStatus`: QueryV1JobStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `StatementsQueryV1Api.GetQueryV1JobStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Confluent Cloud environment ID (e.g., env-xyz123) the statement was run against. | 
**kafkaClusterId** | **string** | Kafka cluster ID (e.g., lkc-xyz123) the statement was run against. | 
**topicName** | **string** | The name of the Kafka topic or Lightning table being queried. | 
**statementId** | **string** | The target unique identifier of the background operational query process. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueryV1JobStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**QueryV1JobStatusResponse**](QueryV1JobStatusResponse.md)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

