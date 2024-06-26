# \TagSuggestionsAiV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryAiV1TagSuggestion**](TagSuggestionsAiV1Api.md#QueryAiV1TagSuggestion) | **Post** /ai/v1/tag-suggestions | Query a Tag Suggestion



## QueryAiV1TagSuggestion

> AiV1TagSuggestionsResponse QueryAiV1TagSuggestion(ctx).AiV1TagSuggestionsRequest(aiV1TagSuggestionsRequest).Execute()

Query a Tag Suggestion



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
    aiV1TagSuggestionsRequest := *openapiclient.NewAiV1TagSuggestionsRequest() // AiV1TagSuggestionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagSuggestionsAiV1Api.QueryAiV1TagSuggestion(context.Background()).AiV1TagSuggestionsRequest(aiV1TagSuggestionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagSuggestionsAiV1Api.QueryAiV1TagSuggestion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryAiV1TagSuggestion`: AiV1TagSuggestionsResponse
    fmt.Fprintf(os.Stdout, "Response from `TagSuggestionsAiV1Api.QueryAiV1TagSuggestion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryAiV1TagSuggestionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiV1TagSuggestionsRequest** | [**AiV1TagSuggestionsRequest**](AiV1TagSuggestionsRequest.md) |  | 

### Return type

[**AiV1TagSuggestionsResponse**](ai.v1.TagSuggestionsResponse.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

