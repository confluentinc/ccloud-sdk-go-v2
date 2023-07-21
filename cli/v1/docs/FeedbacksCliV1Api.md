# \FeedbacksCliV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCliV1Feedback**](FeedbacksCliV1Api.md#CreateCliV1Feedback) | **Post** /cli/v1/feedbacks | Create a Feedback



## CreateCliV1Feedback

> CreateCliV1Feedback(ctx).CliV1Feedback(cliV1Feedback).Execute()

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
    cliV1Feedback := *openapiclient.NewCliV1Feedback() // CliV1Feedback |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeedbacksCliV1Api.CreateCliV1Feedback(context.Background()).CliV1Feedback(cliV1Feedback).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeedbacksCliV1Api.CreateCliV1Feedback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCliV1FeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cliV1Feedback** | [**CliV1Feedback**](CliV1Feedback.md) |  | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

