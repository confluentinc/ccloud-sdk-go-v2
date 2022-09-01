# \NetworksNetworkingV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkingV1Network**](NetworksNetworkingV1Api.md#CreateNetworkingV1Network) | **Post** /networking/v1/networks | Create a Network
[**DeleteNetworkingV1Network**](NetworksNetworkingV1Api.md#DeleteNetworkingV1Network) | **Delete** /networking/v1/networks/{id} | Delete a Network
[**GetNetworkingV1Network**](NetworksNetworkingV1Api.md#GetNetworkingV1Network) | **Get** /networking/v1/networks/{id} | Read a Network
[**ListNetworkingV1Networks**](NetworksNetworkingV1Api.md#ListNetworkingV1Networks) | **Get** /networking/v1/networks | List of Networks
[**UpdateNetworkingV1Network**](NetworksNetworkingV1Api.md#UpdateNetworkingV1Network) | **Patch** /networking/v1/networks/{id} | Update a Network



## CreateNetworkingV1Network

> NetworkingV1Network CreateNetworkingV1Network(ctx).NetworkingV1Network(networkingV1Network).Execute()

Create a Network



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
    networkingV1Network := *openapiclient.NewNetworkingV1Network() // NetworkingV1Network |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksNetworkingV1Api.CreateNetworkingV1Network(context.Background()).NetworkingV1Network(networkingV1Network).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksNetworkingV1Api.CreateNetworkingV1Network``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkingV1Network`: NetworkingV1Network
    fmt.Fprintf(os.Stdout, "Response from `NetworksNetworkingV1Api.CreateNetworkingV1Network`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkingV1NetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkingV1Network** | [**NetworkingV1Network**](NetworkingV1Network.md) |  | 

### Return type

[**NetworkingV1Network**](networking.v1.Network.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkingV1Network

> DeleteNetworkingV1Network(ctx, id).Environment(environment).Execute()

Delete a Network



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
    resp, r, err := api_client.NetworksNetworkingV1Api.DeleteNetworkingV1Network(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksNetworkingV1Api.DeleteNetworkingV1Network``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkingV1NetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkingV1Network

> NetworkingV1Network GetNetworkingV1Network(ctx, id).Environment(environment).Execute()

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
    resp, r, err := api_client.NetworksNetworkingV1Api.GetNetworkingV1Network(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksNetworkingV1Api.GetNetworkingV1Network``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkingV1Network`: NetworkingV1Network
    fmt.Fprintf(os.Stdout, "Response from `NetworksNetworkingV1Api.GetNetworkingV1Network`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkingV1NetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**NetworkingV1Network**](networking.v1.Network.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkingV1Networks

> NetworkingV1NetworkList ListNetworkingV1Networks(ctx).Environment(environment).SpecDisplayName(specDisplayName).SpecCloud(specCloud).SpecRegion(specRegion).SpecCidr(specCidr).StatusPhase(statusPhase).ConnectionType(connectionType).PageSize(pageSize).PageToken(pageToken).Execute()

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
    resp, r, err := api_client.NetworksNetworkingV1Api.ListNetworkingV1Networks(context.Background()).Environment(environment).SpecDisplayName(specDisplayName).SpecCloud(specCloud).SpecRegion(specRegion).SpecCidr(specCidr).StatusPhase(statusPhase).ConnectionType(connectionType).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksNetworkingV1Api.ListNetworkingV1Networks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkingV1Networks`: NetworkingV1NetworkList
    fmt.Fprintf(os.Stdout, "Response from `NetworksNetworkingV1Api.ListNetworkingV1Networks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkingV1NetworksRequest struct via the builder pattern


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

[**NetworkingV1NetworkList**](networking.v1.NetworkList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkingV1Network

> NetworkingV1Network UpdateNetworkingV1Network(ctx, id).NetworkingV1NetworkUpdate(networkingV1NetworkUpdate).Execute()

Update a Network



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
    id := "id_example" // string | The unique identifier for the network.
    networkingV1NetworkUpdate := *openapiclient.NewNetworkingV1NetworkUpdate() // NetworkingV1NetworkUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksNetworkingV1Api.UpdateNetworkingV1Network(context.Background(), id).NetworkingV1NetworkUpdate(networkingV1NetworkUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksNetworkingV1Api.UpdateNetworkingV1Network``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkingV1Network`: NetworkingV1Network
    fmt.Fprintf(os.Stdout, "Response from `NetworksNetworkingV1Api.UpdateNetworkingV1Network`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkingV1NetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkingV1NetworkUpdate** | [**NetworkingV1NetworkUpdate**](NetworkingV1NetworkUpdate.md) |  | 

### Return type

[**NetworkingV1Network**](networking.v1.Network.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

