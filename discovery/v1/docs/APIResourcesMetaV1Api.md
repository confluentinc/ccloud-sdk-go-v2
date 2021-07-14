# \APIResourcesMetaV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMetaV1APIResources**](APIResourcesMetaV1Api.md#ListMetaV1APIResources) | **Get** /{group}/{version} | List of API Resources



## ListMetaV1APIResources

> MetaV1APIResourceList ListMetaV1APIResources(ctx, group, version).Execute()

List of API Resources



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
    group := "group_example" // string | The Group
    version := "version_example" // string | The Version

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.APIResourcesMetaV1Api.ListMetaV1APIResources(context.Background(), group, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIResourcesMetaV1Api.ListMetaV1APIResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMetaV1APIResources`: MetaV1APIResourceList
    fmt.Fprintf(os.Stdout, "Response from `APIResourcesMetaV1Api.ListMetaV1APIResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**group** | **string** | The Group | 
**version** | **string** | The Version | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMetaV1APIResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MetaV1APIResourceList**](MetaV1APIResourceList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

