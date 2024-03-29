# \IPAddressesNetworkingV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListNetworkingV1IpAddresses**](IPAddressesNetworkingV1Api.md#ListNetworkingV1IpAddresses) | **Get** /networking/v1/ip-addresses | List of IP Addresses



## ListNetworkingV1IpAddresses

> NetworkingV1IpAddressList ListNetworkingV1IpAddresses(ctx).Cloud(cloud).Region(region).Services(services).AddressType(addressType).PageSize(pageSize).PageToken(pageToken).Execute()

List of IP Addresses



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
    cloud := []string{"Inner_example"} // []string | Filter the results by exact match for cloud. Pass multiple times to see results matching any of the values. (optional)
    region := []string{"Inner_example"} // []string | Filter the results by exact match for region. Pass multiple times to see results matching any of the values. (optional)
    services := []string{"Inner_example"} // []string | Filter the results by exact match for services. Pass multiple times to see results matching any of the values. (optional)
    addressType := []string{"Inner_example"} // []string | Filter the results by exact match for address_type. Pass multiple times to see results matching any of the values. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPAddressesNetworkingV1Api.ListNetworkingV1IpAddresses(context.Background()).Cloud(cloud).Region(region).Services(services).AddressType(addressType).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAddressesNetworkingV1Api.ListNetworkingV1IpAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkingV1IpAddresses`: NetworkingV1IpAddressList
    fmt.Fprintf(os.Stdout, "Response from `IPAddressesNetworkingV1Api.ListNetworkingV1IpAddresses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkingV1IpAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud** | **[]string** | Filter the results by exact match for cloud. Pass multiple times to see results matching any of the values. | 
 **region** | **[]string** | Filter the results by exact match for region. Pass multiple times to see results matching any of the values. | 
 **services** | **[]string** | Filter the results by exact match for services. Pass multiple times to see results matching any of the values. | 
 **addressType** | **[]string** | Filter the results by exact match for address_type. Pass multiple times to see results matching any of the values. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**NetworkingV1IpAddressList**](networking.v1.IpAddressList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

