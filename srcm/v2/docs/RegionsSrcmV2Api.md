# \RegionsSrcmV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSrcmV2Region**](RegionsSrcmV2Api.md#GetSrcmV2Region) | **Get** /srcm/v2/regions/{id} | Read a Region
[**ListSrcmV2Regions**](RegionsSrcmV2Api.md#ListSrcmV2Regions) | **Get** /srcm/v2/regions | List of Regions



## GetSrcmV2Region

> SrcmV2Region GetSrcmV2Region(ctx, id).Execute()

Read a Region



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
    id := "id_example" // string | The unique identifier for the region.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegionsSrcmV2Api.GetSrcmV2Region(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionsSrcmV2Api.GetSrcmV2Region``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSrcmV2Region`: SrcmV2Region
    fmt.Fprintf(os.Stdout, "Response from `RegionsSrcmV2Api.GetSrcmV2Region`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the region. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSrcmV2RegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SrcmV2Region**](srcm.v2.Region.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSrcmV2Regions

> SrcmV2RegionList ListSrcmV2Regions(ctx).SpecCloud(specCloud).SpecRegionName(specRegionName).SpecPackages(specPackages).PageSize(pageSize).PageToken(pageToken).Execute()

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
    specCloud := "AWS" // string | Filter the results by exact match for spec.cloud. (optional)
    specRegionName := "us-east-2" // string | Filter the results by exact match for spec.region_name. (optional)
    specPackages := *openapiclient.NewMultipleSearchFilter() // MultipleSearchFilter | Filter the results by exact match for spec.packages. Pass multiple times to see results matching any of the values. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegionsSrcmV2Api.ListSrcmV2Regions(context.Background()).SpecCloud(specCloud).SpecRegionName(specRegionName).SpecPackages(specPackages).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionsSrcmV2Api.ListSrcmV2Regions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSrcmV2Regions`: SrcmV2RegionList
    fmt.Fprintf(os.Stdout, "Response from `RegionsSrcmV2Api.ListSrcmV2Regions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSrcmV2RegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **specCloud** | **string** | Filter the results by exact match for spec.cloud. | 
 **specRegionName** | **string** | Filter the results by exact match for spec.region_name. | 
 **specPackages** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter the results by exact match for spec.packages. Pass multiple times to see results matching any of the values. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**SrcmV2RegionList**](srcm.v2.RegionList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

