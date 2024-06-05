# \AvailabilitiesAiV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAiV1Availability**](AvailabilitiesAiV1Api.md#GetAiV1Availability) | **Get** /ai/v1/availability | Read the organization&#39;s ai-assistant setting.



## GetAiV1Availability

> AiV1Availability GetAiV1Availability(ctx).Execute()

Read the organization's ai-assistant setting.



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AvailabilitiesAiV1Api.GetAiV1Availability(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvailabilitiesAiV1Api.GetAiV1Availability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiV1Availability`: AiV1Availability
    fmt.Fprintf(os.Stdout, "Response from `AvailabilitiesAiV1Api.GetAiV1Availability`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiV1AvailabilityRequest struct via the builder pattern


### Return type

[**AiV1Availability**](ai.v1.Availability.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

