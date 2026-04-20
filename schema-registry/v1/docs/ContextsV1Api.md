# \ContextsV1Api

All URIs are relative to *https://psrc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListContexts**](ContextsV1Api.md#ListContexts) | **Get** /contexts | List contexts



## ListContexts

> []string ListContexts(ctx).Offset(offset).Limit(limit).Execute()

List contexts



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
    offset := int32(56) // int32 | Pagination offset for results (optional) (default to 0)
    limit := int32(56) // int32 | Pagination size for results. Ignored if negative (optional) (default to -1)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContextsV1Api.ListContexts(context.Background()).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsV1Api.ListContexts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContexts`: []string
    fmt.Fprintf(os.Stdout, "Response from `ContextsV1Api.ListContexts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListContextsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Pagination offset for results | [default to 0]
 **limit** | **int32** | Pagination size for results. Ignored if negative | [default to -1]

### Return type

**[]string**

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

