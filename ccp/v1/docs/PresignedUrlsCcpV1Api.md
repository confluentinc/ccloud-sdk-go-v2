# \PresignedUrlsCcpV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PresignedUploadUrlCcpV1PresignedUrl**](PresignedUrlsCcpV1Api.md#PresignedUploadUrlCcpV1PresignedUrl) | **Post** /ccp/v1/presigned-upload-url | Request a presigned upload url for new Custom Plugin.



## PresignedUploadUrlCcpV1PresignedUrl

> CcpV1PresignedUrl PresignedUploadUrlCcpV1PresignedUrl(ctx).CcpV1PresignedUrlRequest(ccpV1PresignedUrlRequest).Execute()

Request a presigned upload url for new Custom Plugin.



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
    ccpV1PresignedUrlRequest := *openapiclient.NewCcpV1PresignedUrlRequest() // CcpV1PresignedUrlRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PresignedUrlsCcpV1Api.PresignedUploadUrlCcpV1PresignedUrl(context.Background()).CcpV1PresignedUrlRequest(ccpV1PresignedUrlRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PresignedUrlsCcpV1Api.PresignedUploadUrlCcpV1PresignedUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PresignedUploadUrlCcpV1PresignedUrl`: CcpV1PresignedUrl
    fmt.Fprintf(os.Stdout, "Response from `PresignedUrlsCcpV1Api.PresignedUploadUrlCcpV1PresignedUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPresignedUploadUrlCcpV1PresignedUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccpV1PresignedUrlRequest** | [**CcpV1PresignedUrlRequest**](CcpV1PresignedUrlRequest.md) |  | 

### Return type

[**CcpV1PresignedUrl**](ccp.v1.PresignedUrl.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

