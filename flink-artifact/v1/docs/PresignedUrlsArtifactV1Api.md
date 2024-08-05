# \PresignedUrlsArtifactV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PresignedUploadUrlArtifactV1PresignedUrl**](PresignedUrlsArtifactV1Api.md#PresignedUploadUrlArtifactV1PresignedUrl) | **Post** /artifact/v1/presigned-upload-url | Request a presigned upload URL for a new Flink Artifact.



## PresignedUploadUrlArtifactV1PresignedUrl

> ArtifactV1PresignedUrl PresignedUploadUrlArtifactV1PresignedUrl(ctx).ArtifactV1PresignedUrlRequest(artifactV1PresignedUrlRequest).Execute()

Request a presigned upload URL for a new Flink Artifact.



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
    artifactV1PresignedUrlRequest := *openapiclient.NewArtifactV1PresignedUrlRequest() // ArtifactV1PresignedUrlRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PresignedUrlsArtifactV1Api.PresignedUploadUrlArtifactV1PresignedUrl(context.Background()).ArtifactV1PresignedUrlRequest(artifactV1PresignedUrlRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PresignedUrlsArtifactV1Api.PresignedUploadUrlArtifactV1PresignedUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PresignedUploadUrlArtifactV1PresignedUrl`: ArtifactV1PresignedUrl
    fmt.Fprintf(os.Stdout, "Response from `PresignedUrlsArtifactV1Api.PresignedUploadUrlArtifactV1PresignedUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPresignedUploadUrlArtifactV1PresignedUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artifactV1PresignedUrlRequest** | [**ArtifactV1PresignedUrlRequest**](ArtifactV1PresignedUrlRequest.md) |  | 

### Return type

[**ArtifactV1PresignedUrl**](artifact.v1.PresignedUrl.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

