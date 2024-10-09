# \GatewaysNetworkingV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkingV1Gateway**](GatewaysNetworkingV1Api.md#CreateNetworkingV1Gateway) | **Post** /networking/v1/gateways | Create a Gateway
[**DeleteNetworkingV1Gateway**](GatewaysNetworkingV1Api.md#DeleteNetworkingV1Gateway) | **Delete** /networking/v1/gateways/{id} | Delete a Gateway
[**GetNetworkingV1Gateway**](GatewaysNetworkingV1Api.md#GetNetworkingV1Gateway) | **Get** /networking/v1/gateways/{id} | Read a Gateway
[**ListNetworkingV1Gateways**](GatewaysNetworkingV1Api.md#ListNetworkingV1Gateways) | **Get** /networking/v1/gateways | List of Gateways
[**UpdateNetworkingV1Gateway**](GatewaysNetworkingV1Api.md#UpdateNetworkingV1Gateway) | **Patch** /networking/v1/gateways/{id} | Update a Gateway



## CreateNetworkingV1Gateway

> NetworkingV1Gateway CreateNetworkingV1Gateway(ctx).NetworkingV1Gateway(networkingV1Gateway).Execute()

Create a Gateway



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
    networkingV1Gateway := *openapiclient.NewNetworkingV1Gateway() // NetworkingV1Gateway |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GatewaysNetworkingV1Api.CreateNetworkingV1Gateway(context.Background()).NetworkingV1Gateway(networkingV1Gateway).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewaysNetworkingV1Api.CreateNetworkingV1Gateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkingV1Gateway`: NetworkingV1Gateway
    fmt.Fprintf(os.Stdout, "Response from `GatewaysNetworkingV1Api.CreateNetworkingV1Gateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkingV1GatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkingV1Gateway** | [**NetworkingV1Gateway**](NetworkingV1Gateway.md) |  | 

### Return type

[**NetworkingV1Gateway**](networking.v1.Gateway.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkingV1Gateway

> DeleteNetworkingV1Gateway(ctx, id).Environment(environment).Execute()

Delete a Gateway



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
    id := "id_example" // string | The unique identifier for the gateway.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GatewaysNetworkingV1Api.DeleteNetworkingV1Gateway(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewaysNetworkingV1Api.DeleteNetworkingV1Gateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the gateway. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkingV1GatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

 (empty response body)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkingV1Gateway

> NetworkingV1Gateway GetNetworkingV1Gateway(ctx, id).Environment(environment).Execute()

Read a Gateway



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
    id := "id_example" // string | The unique identifier for the gateway.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GatewaysNetworkingV1Api.GetNetworkingV1Gateway(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewaysNetworkingV1Api.GetNetworkingV1Gateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkingV1Gateway`: NetworkingV1Gateway
    fmt.Fprintf(os.Stdout, "Response from `GatewaysNetworkingV1Api.GetNetworkingV1Gateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the gateway. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkingV1GatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**NetworkingV1Gateway**](networking.v1.Gateway.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkingV1Gateways

> NetworkingV1GatewayList ListNetworkingV1Gateways(ctx).Environment(environment).GatewayType(gatewayType).Id(id).PageSize(pageSize).PageToken(pageToken).Execute()

List of Gateways



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
    gatewayType := []string{"Inner_example"} // []string | Filter the results by exact match for gateway_type. Pass multiple times to see results matching any of the values. (optional)
    id := []string{"Inner_example"} // []string | Filter the results by exact match for id. Pass multiple times to see results matching any of the values. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 100)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GatewaysNetworkingV1Api.ListNetworkingV1Gateways(context.Background()).Environment(environment).GatewayType(gatewayType).Id(id).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewaysNetworkingV1Api.ListNetworkingV1Gateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkingV1Gateways`: NetworkingV1GatewayList
    fmt.Fprintf(os.Stdout, "Response from `GatewaysNetworkingV1Api.ListNetworkingV1Gateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkingV1GatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **gatewayType** | **[]string** | Filter the results by exact match for gateway_type. Pass multiple times to see results matching any of the values. | 
 **id** | **[]string** | Filter the results by exact match for id. Pass multiple times to see results matching any of the values. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 100]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**NetworkingV1GatewayList**](networking.v1.GatewayList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkingV1Gateway

> NetworkingV1Gateway UpdateNetworkingV1Gateway(ctx, id).NetworkingV1GatewayUpdate(networkingV1GatewayUpdate).Execute()

Update a Gateway



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
    id := "id_example" // string | The unique identifier for the gateway.
    networkingV1GatewayUpdate := *openapiclient.NewNetworkingV1GatewayUpdate() // NetworkingV1GatewayUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GatewaysNetworkingV1Api.UpdateNetworkingV1Gateway(context.Background(), id).NetworkingV1GatewayUpdate(networkingV1GatewayUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewaysNetworkingV1Api.UpdateNetworkingV1Gateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkingV1Gateway`: NetworkingV1Gateway
    fmt.Fprintf(os.Stdout, "Response from `GatewaysNetworkingV1Api.UpdateNetworkingV1Gateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the gateway. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkingV1GatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkingV1GatewayUpdate** | [**NetworkingV1GatewayUpdate**](NetworkingV1GatewayUpdate.md) |  | 

### Return type

[**NetworkingV1Gateway**](networking.v1.Gateway.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

