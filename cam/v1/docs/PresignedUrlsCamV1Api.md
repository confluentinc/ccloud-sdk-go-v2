# \PresignedUrlsCamV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PresignedUploadUrlCamV1PresignedUrl**](PresignedUrlsCamV1Api.md#PresignedUploadUrlCamV1PresignedUrl) | **Post** /cam/v1/presigned-upload-url | Request a presigned upload URL for a new Connect Artifact.



## PresignedUploadUrlCamV1PresignedUrl

> CamV1PresignedUrl PresignedUploadUrlCamV1PresignedUrl(ctx).CamV1PresignedUrlRequest(camV1PresignedUrlRequest).Execute()

Request a presigned upload URL for a new Connect Artifact.



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
    camV1PresignedUrlRequest := *openapiclient.NewCamV1PresignedUrlRequest() // CamV1PresignedUrlRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PresignedUrlsCamV1Api.PresignedUploadUrlCamV1PresignedUrl(context.Background()).CamV1PresignedUrlRequest(camV1PresignedUrlRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PresignedUrlsCamV1Api.PresignedUploadUrlCamV1PresignedUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PresignedUploadUrlCamV1PresignedUrl`: CamV1PresignedUrl
    fmt.Fprintf(os.Stdout, "Response from `PresignedUrlsCamV1Api.PresignedUploadUrlCamV1PresignedUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPresignedUploadUrlCamV1PresignedUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **camV1PresignedUrlRequest** | [**CamV1PresignedUrlRequest**](CamV1PresignedUrlRequest.md) |  | 

### Return type

[**CamV1PresignedUrl**](cam.v1.PresignedUrl.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

