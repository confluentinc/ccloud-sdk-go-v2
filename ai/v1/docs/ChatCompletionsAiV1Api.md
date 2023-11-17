# \ChatCompletionsAiV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryAiV1ChatCompletion**](ChatCompletionsAiV1Api.md#QueryAiV1ChatCompletion) | **Post** /ai/v1/chat-completions | Query a Chat Completion



## QueryAiV1ChatCompletion

> AiV1ChatCompletionsReply QueryAiV1ChatCompletion(ctx).AiV1ChatCompletionsRequest(aiV1ChatCompletionsRequest).Execute()

Query a Chat Completion



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
    resp, r, err := api_client.ChatCompletionsAiV1Api.QueryAiV1ChatCompletion(context.Background()).AiV1ChatCompletionsRequest(aiV1ChatCompletionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatCompletionsAiV1Api.QueryAiV1ChatCompletion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryAiV1ChatCompletion`: AiV1ChatCompletionsReply
    fmt.Fprintf(os.Stdout, "Response from `ChatCompletionsAiV1Api.QueryAiV1ChatCompletion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryAiV1ChatCompletionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiV1ChatCompletionsRequest** | [**AiV1ChatCompletionsRequest**](AiV1ChatCompletionsRequest.md) |  | 

### Return type

[**AiV1ChatCompletionsReply**](ai.v1.ChatCompletionsReply.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

