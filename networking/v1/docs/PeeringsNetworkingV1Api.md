# \PeeringsNetworkingV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkingV1Peering**](PeeringsNetworkingV1Api.md#CreateNetworkingV1Peering) | **Post** /networking/v1/peerings | Create a Peering
[**DeleteNetworkingV1Peering**](PeeringsNetworkingV1Api.md#DeleteNetworkingV1Peering) | **Delete** /networking/v1/peerings/{id} | Delete a Peering
[**GetNetworkingV1Peering**](PeeringsNetworkingV1Api.md#GetNetworkingV1Peering) | **Get** /networking/v1/peerings/{id} | Read a Peering
[**ListNetworkingV1Peerings**](PeeringsNetworkingV1Api.md#ListNetworkingV1Peerings) | **Get** /networking/v1/peerings | List of Peerings
[**UpdateNetworkingV1Peering**](PeeringsNetworkingV1Api.md#UpdateNetworkingV1Peering) | **Patch** /networking/v1/peerings/{id} | Update a Peering



## CreateNetworkingV1Peering

> NetworkingV1Peering CreateNetworkingV1Peering(ctx).NetworkingV1Peering(networkingV1Peering).Execute()

Create a Peering



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
    networkingV1Peering := *openapiclient.NewNetworkingV1Peering() // NetworkingV1Peering |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeeringsNetworkingV1Api.CreateNetworkingV1Peering(context.Background()).NetworkingV1Peering(networkingV1Peering).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeeringsNetworkingV1Api.CreateNetworkingV1Peering``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkingV1Peering`: NetworkingV1Peering
    fmt.Fprintf(os.Stdout, "Response from `PeeringsNetworkingV1Api.CreateNetworkingV1Peering`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkingV1PeeringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkingV1Peering** | [**NetworkingV1Peering**](NetworkingV1Peering.md) |  | 

### Return type

[**NetworkingV1Peering**](networking.v1.Peering.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkingV1Peering

> DeleteNetworkingV1Peering(ctx, id).Environment(environment).Execute()

Delete a Peering



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
    id := "id_example" // string | The unique identifier for the peering.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeeringsNetworkingV1Api.DeleteNetworkingV1Peering(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeeringsNetworkingV1Api.DeleteNetworkingV1Peering``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the peering. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkingV1PeeringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkingV1Peering

> NetworkingV1Peering GetNetworkingV1Peering(ctx, id).Environment(environment).Execute()

Read a Peering



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
    id := "id_example" // string | The unique identifier for the peering.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeeringsNetworkingV1Api.GetNetworkingV1Peering(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeeringsNetworkingV1Api.GetNetworkingV1Peering``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkingV1Peering`: NetworkingV1Peering
    fmt.Fprintf(os.Stdout, "Response from `PeeringsNetworkingV1Api.GetNetworkingV1Peering`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the peering. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkingV1PeeringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**NetworkingV1Peering**](networking.v1.Peering.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkingV1Peerings

> NetworkingV1PeeringList ListNetworkingV1Peerings(ctx).Environment(environment).SpecDisplayName(specDisplayName).StatusPhase(statusPhase).SpecNetwork(specNetwork).PageSize(pageSize).PageToken(pageToken).Execute()

List of Peerings



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
    statusPhase := *openapiclient.NewMultipleSearchFilter() // MultipleSearchFilter | Filter the results by exact match for status.phase. Pass multiple times to see results matching any of the values. (optional)
    specNetwork :=  // MultipleSearchFilter | Filter the results by exact match for spec.network. Pass multiple times to see results matching any of the values. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeeringsNetworkingV1Api.ListNetworkingV1Peerings(context.Background()).Environment(environment).SpecDisplayName(specDisplayName).StatusPhase(statusPhase).SpecNetwork(specNetwork).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeeringsNetworkingV1Api.ListNetworkingV1Peerings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkingV1Peerings`: NetworkingV1PeeringList
    fmt.Fprintf(os.Stdout, "Response from `PeeringsNetworkingV1Api.ListNetworkingV1Peerings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkingV1PeeringsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **specDisplayName** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter the results by exact match for spec.display_name. Pass multiple times to see results matching any of the values. | 
 **statusPhase** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter the results by exact match for status.phase. Pass multiple times to see results matching any of the values. | 
 **specNetwork** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter the results by exact match for spec.network. Pass multiple times to see results matching any of the values. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**NetworkingV1PeeringList**](networking.v1.PeeringList.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkingV1Peering

> NetworkingV1Peering UpdateNetworkingV1Peering(ctx, id).NetworkingV1PeeringUpdate(networkingV1PeeringUpdate).Execute()

Update a Peering



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
    id := "id_example" // string | The unique identifier for the peering.
    networkingV1PeeringUpdate := *openapiclient.NewNetworkingV1PeeringUpdate() // NetworkingV1PeeringUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeeringsNetworkingV1Api.UpdateNetworkingV1Peering(context.Background(), id).NetworkingV1PeeringUpdate(networkingV1PeeringUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeeringsNetworkingV1Api.UpdateNetworkingV1Peering``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkingV1Peering`: NetworkingV1Peering
    fmt.Fprintf(os.Stdout, "Response from `PeeringsNetworkingV1Api.UpdateNetworkingV1Peering`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the peering. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkingV1PeeringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkingV1PeeringUpdate** | [**NetworkingV1PeeringUpdate**](NetworkingV1PeeringUpdate.md) |  | 

### Return type

[**NetworkingV1Peering**](networking.v1.Peering.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

