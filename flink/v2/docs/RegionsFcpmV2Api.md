# \RegionsFcpmV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListFcpmV2Regions**](RegionsFcpmV2Api.md#ListFcpmV2Regions) | **Get** /fcpm/v2/regions | List of Regions



## ListFcpmV2Regions

> FcpmV2RegionList ListFcpmV2Regions(ctx).Cloud(cloud).RegionName(regionName).PageSize(pageSize).PageToken(pageToken).Execute()

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
    regionName := "us-east-2" // string | Filter the results by exact match for region_name. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegionsFcpmV2Api.ListFcpmV2Regions(context.Background()).Cloud(cloud).RegionName(regionName).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionsFcpmV2Api.ListFcpmV2Regions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFcpmV2Regions`: FcpmV2RegionList
    fmt.Fprintf(os.Stdout, "Response from `RegionsFcpmV2Api.ListFcpmV2Regions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFcpmV2RegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud** | **string** | Filter the results by exact match for cloud. | 
 **regionName** | **string** | Filter the results by exact match for region_name. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**FcpmV2RegionList**](fcpm.v2.RegionList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

