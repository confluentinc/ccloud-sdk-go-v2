# \SearchV1Api

All URIs are relative to *https://psrc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchUsingAttribute**](SearchV1Api.md#SearchUsingAttribute) | **Get** /catalog/v1/search/attribute | Search by Attribute
[**SearchUsingBasic**](SearchV1Api.md#SearchUsingBasic) | **Get** /catalog/v1/search/basic | Search by Fulltext Query



## SearchUsingAttribute

> SearchResult SearchUsingAttribute(ctx).Type_(type_).Attr(attr).AttrName(attrName).AttrValuePrefix(attrValuePrefix).Tag(tag).SortBy(sortBy).SortOrder(sortOrder).Deleted(deleted).Limit(limit).Offset(offset).Execute()

Search by Attribute



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
    type_ := []string{"Inner_example"} // []string | Limit the result to only entities of specified types (optional)
    attr := []string{"Inner_example"} // []string | One of more additional attributes to return in the response (optional)
    attrName := []string{"Inner_example"} // []string | The attribute to search (optional)
    attrValuePrefix := []string{"Inner_example"} // []string | The prefix for the attribute value to search (optional)
    tag := "tag_example" // string | Limit the result to only entities tagged with the given tag (optional)
    sortBy := "sortBy_example" // string | An attribute to sort by (optional)
    sortOrder := "sortOrder_example" // string | Sort order, either ASCENDING (default) or DESCENDING (optional)
    deleted := true // bool | Whether to include deleted entities (optional)
    limit := int32(56) // int32 | Limit the result set to only include the specified number of entries (maximum 500) (optional)
    offset := int32(56) // int32 | Start offset of the result set (useful for pagination) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SearchV1Api.SearchUsingAttribute(context.Background()).Type_(type_).Attr(attr).AttrName(attrName).AttrValuePrefix(attrValuePrefix).Tag(tag).SortBy(sortBy).SortOrder(sortOrder).Deleted(deleted).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchV1Api.SearchUsingAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchUsingAttribute`: SearchResult
    fmt.Fprintf(os.Stdout, "Response from `SearchV1Api.SearchUsingAttribute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUsingAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **[]string** | Limit the result to only entities of specified types | 
 **attr** | **[]string** | One of more additional attributes to return in the response | 
 **attrName** | **[]string** | The attribute to search | 
 **attrValuePrefix** | **[]string** | The prefix for the attribute value to search | 
 **tag** | **string** | Limit the result to only entities tagged with the given tag | 
 **sortBy** | **string** | An attribute to sort by | 
 **sortOrder** | **string** | Sort order, either ASCENDING (default) or DESCENDING | 
 **deleted** | **bool** | Whether to include deleted entities | 
 **limit** | **int32** | Limit the result set to only include the specified number of entries (maximum 500) | 
 **offset** | **int32** | Start offset of the result set (useful for pagination) | 

### Return type

[**SearchResult**](SearchResult.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUsingBasic

> SearchResult SearchUsingBasic(ctx).Query(query).Type_(type_).Attr(attr).Tag(tag).SortBy(sortBy).SortOrder(sortOrder).Deleted(deleted).Limit(limit).Offset(offset).Execute()

Search by Fulltext Query



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
    query := "query_example" // string | The full-text query (optional)
    type_ := []string{"Inner_example"} // []string | Limit the result to only entities of specified types (optional)
    attr := []string{"Inner_example"} // []string | One of more additional attributes to return in the response (optional)
    tag := "tag_example" // string | Limit the result to only entities tagged with the given tag (optional)
    sortBy := "sortBy_example" // string | An attribute to sort by (optional)
    sortOrder := "sortOrder_example" // string | Sort order, either ASCENDING (default) or DESCENDING (optional)
    deleted := true // bool | Whether to include deleted entities (optional)
    limit := int32(56) // int32 | Limit the result set to only include the specified number of entries (maximum 500) (optional)
    offset := int32(56) // int32 | Start offset of the result set (useful for pagination) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SearchV1Api.SearchUsingBasic(context.Background()).Query(query).Type_(type_).Attr(attr).Tag(tag).SortBy(sortBy).SortOrder(sortOrder).Deleted(deleted).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchV1Api.SearchUsingBasic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchUsingBasic`: SearchResult
    fmt.Fprintf(os.Stdout, "Response from `SearchV1Api.SearchUsingBasic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUsingBasicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The full-text query | 
 **type_** | **[]string** | Limit the result to only entities of specified types | 
 **attr** | **[]string** | One of more additional attributes to return in the response | 
 **tag** | **string** | Limit the result to only entities tagged with the given tag | 
 **sortBy** | **string** | An attribute to sort by | 
 **sortOrder** | **string** | Sort order, either ASCENDING (default) or DESCENDING | 
 **deleted** | **bool** | Whether to include deleted entities | 
 **limit** | **int32** | Limit the result set to only include the specified number of entries (maximum 500) | 
 **offset** | **int32** | Start offset of the result set (useful for pagination) | 

### Return type

[**SearchResult**](SearchResult.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

