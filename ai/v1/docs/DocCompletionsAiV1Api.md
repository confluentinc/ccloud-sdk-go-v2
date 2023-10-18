# \DocCompletionsAiV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryAiV1DocCompletion**](DocCompletionsAiV1Api.md#QueryAiV1DocCompletion) | **Post** /ai/v1/doc-completions | Query a Doc Completion



## QueryAiV1DocCompletion

> AiV1DocCompletionsReply QueryAiV1DocCompletion(ctx).AiV1DocCompletionsRequest(aiV1DocCompletionsRequest).Execute()

Query a Doc Completion



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
    aiV1DocCompletionsRequest := *openapiclient.NewAiV1DocCompletionsRequest() // AiV1DocCompletionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocCompletionsAiV1Api.QueryAiV1DocCompletion(context.Background()).AiV1DocCompletionsRequest(aiV1DocCompletionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocCompletionsAiV1Api.QueryAiV1DocCompletion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryAiV1DocCompletion`: AiV1DocCompletionsReply
    fmt.Fprintf(os.Stdout, "Response from `DocCompletionsAiV1Api.QueryAiV1DocCompletion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryAiV1DocCompletionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiV1DocCompletionsRequest** | [**AiV1DocCompletionsRequest**](AiV1DocCompletionsRequest.md) |  | 

### Return type

[**AiV1DocCompletionsReply**](AiV1DocCompletionsReply.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

