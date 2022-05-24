# \NetworksNetworkingAdminV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkingAdminV1Network**](NetworksNetworkingAdminV1Api.md#GetNetworkingAdminV1Network) | **Get** /networking-admin/v1/networks/{id} | Read a Network
[**ListNetworkingAdminV1Networks**](NetworksNetworkingAdminV1Api.md#ListNetworkingAdminV1Networks) | **Get** /networking-admin/v1/networks | List of Networks



## GetNetworkingAdminV1Network

> NetworkingAdminV1Network GetNetworkingAdminV1Network(ctx, id).Environment(environment).Execute()

Read a Network



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
    environment := "env-00000" // string | Scope the operation to the given environment.
    id := "id_example" // string | The unique identifier for the network.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksNetworkingAdminV1Api.GetNetworkingAdminV1Network(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksNetworkingAdminV1Api.GetNetworkingAdminV1Network``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkingAdminV1Network`: NetworkingAdminV1Network
    fmt.Fprintf(os.Stdout, "Response from `NetworksNetworkingAdminV1Api.GetNetworkingAdminV1Network`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkingAdminV1NetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**NetworkingAdminV1Network**](networking-admin.v1.Network.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkingAdminV1Networks

> NetworkingAdminV1NetworkList ListNetworkingAdminV1Networks(ctx).Environment(environment).SpecDisplayName(specDisplayName).SpecCloud(specCloud).SpecRegion(specRegion).SpecCidr(specCidr).StatusPhase(statusPhase).ConnectionType(connectionType).PageSize(pageSize).PageToken(pageToken).Execute()

List of Networks



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
    environment := "env-00000" // string | Filter the results by exact match for environment.
    specDisplayName := *openapiclient.NewMultipleSearchFilter() // MultipleSearchFilter | Filter the results by exact match for spec.display_name. Pass multiple times to see results matching any of the values. (optional)
    specCloud := *openapiclient.NewMultipleSearchFilter() // MultipleSearchFilter | Filter the results by exact match for spec.cloud. Pass multiple times to see results matching any of the values. (optional)
    specRegion :=  // MultipleSearchFilter | Filter the results by exact match for spec.region. Pass multiple times to see results matching any of the values. (optional)
    specCidr :=  // MultipleSearchFilter | Filter the results by exact match for spec.cidr. Pass multiple times to see results matching any of the values. (optional)
    statusPhase :=  // MultipleSearchFilter | Filter the results by exact match for status.phase. Pass multiple times to see results matching any of the values. (optional)
    connectionType :=  // MultipleSearchFilter | Filter the results by exact match for connection_type. Pass multiple times to see results matching any of the values. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksNetworkingAdminV1Api.ListNetworkingAdminV1Networks(context.Background()).Environment(environment).SpecDisplayName(specDisplayName).SpecCloud(specCloud).SpecRegion(specRegion).SpecCidr(specCidr).StatusPhase(statusPhase).ConnectionType(connectionType).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksNetworkingAdminV1Api.ListNetworkingAdminV1Networks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkingAdminV1Networks`: NetworkingAdminV1NetworkList
    fmt.Fprintf(os.Stdout, "Response from `NetworksNetworkingAdminV1Api.ListNetworkingAdminV1Networks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkingAdminV1NetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **specDisplayName** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter the results by exact match for spec.display_name. Pass multiple times to see results matching any of the values. | 
 **specCloud** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter the results by exact match for spec.cloud. Pass multiple times to see results matching any of the values. | 
 **specRegion** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter the results by exact match for spec.region. Pass multiple times to see results matching any of the values. | 
 **specCidr** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter the results by exact match for spec.cidr. Pass multiple times to see results matching any of the values. | 
 **statusPhase** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter the results by exact match for status.phase. Pass multiple times to see results matching any of the values. | 
 **connectionType** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter the results by exact match for connection_type. Pass multiple times to see results matching any of the values. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**NetworkingAdminV1NetworkList**](networking-admin.v1.NetworkList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

