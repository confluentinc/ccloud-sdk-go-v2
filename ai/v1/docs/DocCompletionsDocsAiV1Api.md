# \DocCompletionsDocsAiV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryDocsAiV1DocCompletion**](DocCompletionsDocsAiV1Api.md#QueryDocsAiV1DocCompletion) | **Post** /docs-ai/v1/doc-completions | Query a Doc Completion



## QueryDocsAiV1DocCompletion

> AiV1ChatCompletionsReply QueryDocsAiV1DocCompletion(ctx).AiV1ChatCompletionsRequest(aiV1ChatCompletionsRequest).Execute()

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
    aiV1ChatCompletionsRequest := *openapiclient.NewAiV1ChatCompletionsRequest() // AiV1ChatCompletionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocCompletionsDocsAiV1Api.QueryDocsAiV1DocCompletion(context.Background()).AiV1ChatCompletionsRequest(aiV1ChatCompletionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocCompletionsDocsAiV1Api.QueryDocsAiV1DocCompletion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryDocsAiV1DocCompletion`: AiV1ChatCompletionsReply
    fmt.Fprintf(os.Stdout, "Response from `DocCompletionsDocsAiV1Api.QueryDocsAiV1DocCompletion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryDocsAiV1DocCompletionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiV1ChatCompletionsRequest** | [**AiV1ChatCompletionsRequest**](AiV1ChatCompletionsRequest.md) |  | 

### Return type

[**AiV1ChatCompletionsReply**](ai.v1.ChatCompletionsReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

