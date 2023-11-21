# \DocCompletionsAiV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryAiV1DocCompletion**](DocCompletionsAiV1Api.md#QueryAiV1DocCompletion) | **Post** /ai/v1/doc-completions | Query a Doc Completion



## QueryAiV1DocCompletion

> AiV1ChatCompletionsReply QueryAiV1DocCompletion(ctx).AiV1ChatCompletionsRequest(aiV1ChatCompletionsRequest).Execute()

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
    resp, r, err := api_client.DocCompletionsAiV1Api.QueryAiV1DocCompletion(context.Background()).AiV1ChatCompletionsRequest(aiV1ChatCompletionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocCompletionsAiV1Api.QueryAiV1DocCompletion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryAiV1DocCompletion`: AiV1ChatCompletionsReply
    fmt.Fprintf(os.Stdout, "Response from `DocCompletionsAiV1Api.QueryAiV1DocCompletion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryAiV1DocCompletionRequest struct via the builder pattern


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

