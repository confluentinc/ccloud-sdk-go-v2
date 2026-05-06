# \RegionsRtceV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListRtceV1Regions**](RegionsRtceV1Api.md#ListRtceV1Regions) | **Get** /rtce/v1/regions | List of Regions



## ListRtceV1Regions

> RtceV1RegionList ListRtceV1Regions(ctx).Cloud(cloud).Region(region).PageSize(pageSize).PageToken(pageToken).Execute()

List of Regions



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
    cloud := "AWS" // string | Filter the results by exact match for cloud. (optional)
    region := "us-east-2" // string | Filter the results by exact match for region. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegionsRtceV1Api.ListRtceV1Regions(context.Background()).Cloud(cloud).Region(region).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionsRtceV1Api.ListRtceV1Regions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRtceV1Regions`: RtceV1RegionList
    fmt.Fprintf(os.Stdout, "Response from `RegionsRtceV1Api.ListRtceV1Regions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRtceV1RegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud** | **string** | Filter the results by exact match for cloud. | 
 **region** | **string** | Filter the results by exact match for region. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**RtceV1RegionList**](rtce.v1.RegionList.md)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

