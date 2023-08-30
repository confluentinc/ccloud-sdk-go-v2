# \PresignedUrlsConnectV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PresignedUploadUrlConnectV1PresignedUrl**](PresignedUrlsConnectV1Api.md#PresignedUploadUrlConnectV1PresignedUrl) | **Post** /connect/v1/presigned-upload-url | Request a presigned upload url for new Custom Connector Plugin.



## PresignedUploadUrlConnectV1PresignedUrl

> ConnectV1PresignedUrl PresignedUploadUrlConnectV1PresignedUrl(ctx).ConnectV1PresignedUrlRequest(connectV1PresignedUrlRequest).Execute()

Request a presigned upload url for new Custom Connector Plugin.



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
    connectV1PresignedUrlRequest := *openapiclient.NewConnectV1PresignedUrlRequest() // ConnectV1PresignedUrlRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PresignedUrlsConnectV1Api.PresignedUploadUrlConnectV1PresignedUrl(context.Background()).ConnectV1PresignedUrlRequest(connectV1PresignedUrlRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PresignedUrlsConnectV1Api.PresignedUploadUrlConnectV1PresignedUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PresignedUploadUrlConnectV1PresignedUrl`: ConnectV1PresignedUrl
    fmt.Fprintf(os.Stdout, "Response from `PresignedUrlsConnectV1Api.PresignedUploadUrlConnectV1PresignedUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPresignedUploadUrlConnectV1PresignedUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectV1PresignedUrlRequest** | [**ConnectV1PresignedUrlRequest**](ConnectV1PresignedUrlRequest.md) |  | 

### Return type

[**ConnectV1PresignedUrl**](connect.v1.PresignedUrl.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

