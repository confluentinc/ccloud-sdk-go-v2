# \FeedbacksAiV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAiV1ChatCompletionFeedback**](FeedbacksAiV1Api.md#CreateAiV1ChatCompletionFeedback) | **Post** /ai/v1/chat-completions/{chat_completion_id}/feedback | Create a Feedback
[**CreateAiV1DocCompletionFeedback**](FeedbacksAiV1Api.md#CreateAiV1DocCompletionFeedback) | **Post** /ai/v1/doc-completions/{doc_completion_id}/feedback | Create a Feedback



## CreateAiV1ChatCompletionFeedback

> AiV1Feedback CreateAiV1ChatCompletionFeedback(ctx, chatCompletionId).AiV1Feedback(aiV1Feedback).Execute()

Create a Feedback



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
    chatCompletionId := "chatCompletionId_example" // string | The Chat Completion
    aiV1Feedback := *openapiclient.NewAiV1Feedback() // AiV1Feedback |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeedbacksAiV1Api.CreateAiV1ChatCompletionFeedback(context.Background(), chatCompletionId).AiV1Feedback(aiV1Feedback).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeedbacksAiV1Api.CreateAiV1ChatCompletionFeedback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiV1ChatCompletionFeedback`: AiV1Feedback
    fmt.Fprintf(os.Stdout, "Response from `FeedbacksAiV1Api.CreateAiV1ChatCompletionFeedback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatCompletionId** | **string** | The Chat Completion | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiV1ChatCompletionFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aiV1Feedback** | [**AiV1Feedback**](AiV1Feedback.md) |  | 

### Return type

[**AiV1Feedback**](AiV1Feedback.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAiV1DocCompletionFeedback

> AiV1Feedback CreateAiV1DocCompletionFeedback(ctx, docCompletionId).AiV1Feedback(aiV1Feedback).Execute()

Create a Feedback



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
    docCompletionId := "docCompletionId_example" // string | The Doc Completion
    aiV1Feedback := *openapiclient.NewAiV1Feedback() // AiV1Feedback |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeedbacksAiV1Api.CreateAiV1DocCompletionFeedback(context.Background(), docCompletionId).AiV1Feedback(aiV1Feedback).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeedbacksAiV1Api.CreateAiV1DocCompletionFeedback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiV1DocCompletionFeedback`: AiV1Feedback
    fmt.Fprintf(os.Stdout, "Response from `FeedbacksAiV1Api.CreateAiV1DocCompletionFeedback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**docCompletionId** | **string** | The Doc Completion | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiV1DocCompletionFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aiV1Feedback** | [**AiV1Feedback**](AiV1Feedback.md) |  | 

### Return type

[**AiV1Feedback**](AiV1Feedback.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

