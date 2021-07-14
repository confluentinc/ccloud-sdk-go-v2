# \APIGroupsMetaV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetaV1APIGroup**](APIGroupsMetaV1Api.md#GetMetaV1APIGroup) | **Get** /{name} | Read an API Group
[**ListMetaV1APIGroups**](APIGroupsMetaV1Api.md#ListMetaV1APIGroups) | **Get** / | List of API Groups



## GetMetaV1APIGroup

> MetaV1APIGroup GetMetaV1APIGroup(ctx, name).Execute()

Read an API Group



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
    name := "name_example" // string | The API Group name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.APIGroupsMetaV1Api.GetMetaV1APIGroup(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIGroupsMetaV1Api.GetMetaV1APIGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetaV1APIGroup`: MetaV1APIGroup
    fmt.Fprintf(os.Stdout, "Response from `APIGroupsMetaV1Api.GetMetaV1APIGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The API Group name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetaV1APIGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetaV1APIGroup**](meta.v1.APIGroup.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetaV1APIGroups

> MetaV1APIGroupList ListMetaV1APIGroups(ctx).Execute()

List of API Groups



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.APIGroupsMetaV1Api.ListMetaV1APIGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIGroupsMetaV1Api.ListMetaV1APIGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMetaV1APIGroups`: MetaV1APIGroupList
    fmt.Fprintf(os.Stdout, "Response from `APIGroupsMetaV1Api.ListMetaV1APIGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMetaV1APIGroupsRequest struct via the builder pattern


### Return type

[**MetaV1APIGroupList**](MetaV1APIGroupList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

