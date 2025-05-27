# \PresignedUrlsCcpmV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCcpmV1PresignedUrl**](PresignedUrlsCcpmV1Api.md#CreateCcpmV1PresignedUrl) | **Post** /ccpm/v1/presigned-upload-url | Request a presigned upload URL for a new Custom Connect Plugin.



## CreateCcpmV1PresignedUrl

> CcpmV1PresignedUrl CreateCcpmV1PresignedUrl(ctx).CcpmV1PresignedUrl(ccpmV1PresignedUrl).Execute()

Request a presigned upload URL for a new Custom Connect Plugin.



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
    ccpmV1PresignedUrl := *openapiclient.NewCcpmV1PresignedUrl() // CcpmV1PresignedUrl |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PresignedUrlsCcpmV1Api.CreateCcpmV1PresignedUrl(context.Background()).CcpmV1PresignedUrl(ccpmV1PresignedUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PresignedUrlsCcpmV1Api.CreateCcpmV1PresignedUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCcpmV1PresignedUrl`: CcpmV1PresignedUrl
    fmt.Fprintf(os.Stdout, "Response from `PresignedUrlsCcpmV1Api.CreateCcpmV1PresignedUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCcpmV1PresignedUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccpmV1PresignedUrl** | [**CcpmV1PresignedUrl**](CcpmV1PresignedUrl.md) |  | 

### Return type

[**CcpmV1PresignedUrl**](ccpm.v1.PresignedUrl.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

