# \S3SignerAPIAPI

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SignS3Request**](S3SignerAPIAPI.md#SignS3Request) | **Post** /v1/aws/s3/sign | Remotely signs S3 requests



## SignS3Request

> SignS3Request200Response SignS3Request(ctx).S3SignRequest(s3SignRequest).Execute()

Remotely signs S3 requests

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/confluentinc/ccloud-sdk-go-v2"
)

func main() {
	s3SignRequest := *openapiclient.NewS3SignRequest("Region_example", "Uri_example", "Method_example", *openapiclient.NewS3Headers()) // S3SignRequest | The request containing the headers to be signed

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3SignerAPIAPI.SignS3Request(context.Background()).S3SignRequest(s3SignRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3SignerAPIAPI.SignS3Request``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignS3Request`: SignS3Request200Response
	fmt.Fprintf(os.Stdout, "Response from `S3SignerAPIAPI.SignS3Request`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignS3RequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **s3SignRequest** | [**S3SignRequest**](S3SignRequest.md) | The request containing the headers to be signed | 

### Return type

[**SignS3Request200Response**](SignS3Request200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

