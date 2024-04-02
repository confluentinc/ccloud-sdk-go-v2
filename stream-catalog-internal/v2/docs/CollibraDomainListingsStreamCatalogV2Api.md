# \CollibraDomainListingsStreamCatalogV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostStreamCatalogV2CollibraDomainListing**](CollibraDomainListingsStreamCatalogV2Api.md#PostStreamCatalogV2CollibraDomainListing) | **Post** /stream-catalog/v2/collibra-sink/domains | List collibra sink domains



## PostStreamCatalogV2CollibraDomainListing

> StreamCatalogV2CollibraDomains PostStreamCatalogV2CollibraDomainListing(ctx).StreamCatalogV2CollibraDomainsRequest(streamCatalogV2CollibraDomainsRequest).Execute()

List collibra sink domains



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
    streamCatalogV2CollibraDomainsRequest := *openapiclient.NewStreamCatalogV2CollibraDomainsRequest() // StreamCatalogV2CollibraDomainsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollibraDomainListingsStreamCatalogV2Api.PostStreamCatalogV2CollibraDomainListing(context.Background()).StreamCatalogV2CollibraDomainsRequest(streamCatalogV2CollibraDomainsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollibraDomainListingsStreamCatalogV2Api.PostStreamCatalogV2CollibraDomainListing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostStreamCatalogV2CollibraDomainListing`: StreamCatalogV2CollibraDomains
    fmt.Fprintf(os.Stdout, "Response from `CollibraDomainListingsStreamCatalogV2Api.PostStreamCatalogV2CollibraDomainListing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostStreamCatalogV2CollibraDomainListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **streamCatalogV2CollibraDomainsRequest** | [**StreamCatalogV2CollibraDomainsRequest**](StreamCatalogV2CollibraDomainsRequest.md) |  | 

### Return type

[**StreamCatalogV2CollibraDomains**](StreamCatalogV2CollibraDomains.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

