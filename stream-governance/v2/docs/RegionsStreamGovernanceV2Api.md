# \RegionsStreamGovernanceV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStreamGovernanceV2Region**](RegionsStreamGovernanceV2Api.md#GetStreamGovernanceV2Region) | **Get** /stream-governance/v2/regions/{id} | Read a Region
[**ListStreamGovernanceV2Regions**](RegionsStreamGovernanceV2Api.md#ListStreamGovernanceV2Regions) | **Get** /stream-governance/v2/regions | List of Regions



## GetStreamGovernanceV2Region

> StreamGovernanceV2Region GetStreamGovernanceV2Region(ctx, id).Execute()

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
    resp, r, err := api_client.RegionsStreamGovernanceV2Api.GetStreamGovernanceV2Region(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionsStreamGovernanceV2Api.GetStreamGovernanceV2Region``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStreamGovernanceV2Region`: StreamGovernanceV2Region
    fmt.Fprintf(os.Stdout, "Response from `RegionsStreamGovernanceV2Api.GetStreamGovernanceV2Region`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the region. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamGovernanceV2RegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StreamGovernanceV2Region**](stream-governance.v2.Region.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStreamGovernanceV2Regions

> StreamGovernanceV2RegionList ListStreamGovernanceV2Regions(ctx).SpecCloud(specCloud).SpecRegionName(specRegionName).SpecPackages(specPackages).PageSize(pageSize).PageToken(pageToken).Execute()

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
    resp, r, err := api_client.RegionsStreamGovernanceV2Api.ListStreamGovernanceV2Regions(context.Background()).SpecCloud(specCloud).SpecRegionName(specRegionName).SpecPackages(specPackages).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionsStreamGovernanceV2Api.ListStreamGovernanceV2Regions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStreamGovernanceV2Regions`: StreamGovernanceV2RegionList
    fmt.Fprintf(os.Stdout, "Response from `RegionsStreamGovernanceV2Api.ListStreamGovernanceV2Regions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListStreamGovernanceV2RegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **specCloud** | **string** | Filter the results by exact match for spec.cloud. | 
 **specRegionName** | **string** | Filter the results by exact match for spec.region_name. | 
 **specPackages** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter the results by exact match for spec.packages. Pass multiple times to see results matching any of the values. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**StreamGovernanceV2RegionList**](stream-governance.v2.RegionList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

