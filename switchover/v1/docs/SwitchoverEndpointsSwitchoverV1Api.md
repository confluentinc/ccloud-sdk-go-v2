# \SwitchoverEndpointsSwitchoverV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSwitchoverV1SwitchoverEndpoint**](SwitchoverEndpointsSwitchoverV1Api.md#CreateSwitchoverV1SwitchoverEndpoint) | **Post** /switchover/v1/switchover-endpoints | Create a Switchover Endpoint
[**DeleteSwitchoverV1SwitchoverEndpoint**](SwitchoverEndpointsSwitchoverV1Api.md#DeleteSwitchoverV1SwitchoverEndpoint) | **Delete** /switchover/v1/switchover-endpoints/{id} | Delete a Switchover Endpoint
[**GetSwitchoverV1SwitchoverEndpoint**](SwitchoverEndpointsSwitchoverV1Api.md#GetSwitchoverV1SwitchoverEndpoint) | **Get** /switchover/v1/switchover-endpoints/{id} | Read a Switchover Endpoint
[**ListSwitchoverV1SwitchoverEndpoints**](SwitchoverEndpointsSwitchoverV1Api.md#ListSwitchoverV1SwitchoverEndpoints) | **Get** /switchover/v1/switchover-endpoints | List of Switchover Endpoints
[**UpdateSwitchoverV1SwitchoverEndpoint**](SwitchoverEndpointsSwitchoverV1Api.md#UpdateSwitchoverV1SwitchoverEndpoint) | **Put** /switchover/v1/switchover-endpoints/{id} | Update a Switchover Endpoint



## CreateSwitchoverV1SwitchoverEndpoint

> SwitchoverV1SwitchoverEndpoint CreateSwitchoverV1SwitchoverEndpoint(ctx).SwitchoverV1SwitchoverEndpoint(switchoverV1SwitchoverEndpoint).Execute()

Create a Switchover Endpoint



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
    switchoverV1SwitchoverEndpoint := *openapiclient.NewSwitchoverV1SwitchoverEndpoint() // SwitchoverV1SwitchoverEndpoint |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SwitchoverEndpointsSwitchoverV1Api.CreateSwitchoverV1SwitchoverEndpoint(context.Background()).SwitchoverV1SwitchoverEndpoint(switchoverV1SwitchoverEndpoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SwitchoverEndpointsSwitchoverV1Api.CreateSwitchoverV1SwitchoverEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSwitchoverV1SwitchoverEndpoint`: SwitchoverV1SwitchoverEndpoint
    fmt.Fprintf(os.Stdout, "Response from `SwitchoverEndpointsSwitchoverV1Api.CreateSwitchoverV1SwitchoverEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSwitchoverV1SwitchoverEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **switchoverV1SwitchoverEndpoint** | [**SwitchoverV1SwitchoverEndpoint**](SwitchoverV1SwitchoverEndpoint.md) |  | 

### Return type

[**SwitchoverV1SwitchoverEndpoint**](switchover.v1.SwitchoverEndpoint.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSwitchoverV1SwitchoverEndpoint

> SwitchoverV1SwitchoverEndpoint DeleteSwitchoverV1SwitchoverEndpoint(ctx, id).Environment(environment).Execute()

Delete a Switchover Endpoint



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
    environment := "env-00000" // string | Scope the operation to the environment with this ID. The bare-ID form of `environment_crn`; it is the standard Confluent collection filter, passed as the `?environment=` query parameter. 
    id := "id_example" // string | The unique identifier for the switchover endpoint.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SwitchoverEndpointsSwitchoverV1Api.DeleteSwitchoverV1SwitchoverEndpoint(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SwitchoverEndpointsSwitchoverV1Api.DeleteSwitchoverV1SwitchoverEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSwitchoverV1SwitchoverEndpoint`: SwitchoverV1SwitchoverEndpoint
    fmt.Fprintf(os.Stdout, "Response from `SwitchoverEndpointsSwitchoverV1Api.DeleteSwitchoverV1SwitchoverEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the switchover endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSwitchoverV1SwitchoverEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the environment with this ID. The bare-ID form of &#x60;environment_crn&#x60;; it is the standard Confluent collection filter, passed as the &#x60;?environment&#x3D;&#x60; query parameter.  | 


### Return type

[**SwitchoverV1SwitchoverEndpoint**](SwitchoverV1SwitchoverEndpoint.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSwitchoverV1SwitchoverEndpoint

> SwitchoverV1SwitchoverEndpoint GetSwitchoverV1SwitchoverEndpoint(ctx, id).Environment(environment).Execute()

Read a Switchover Endpoint



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
    environment := "env-00000" // string | Scope the operation to the environment with this ID. The bare-ID form of `environment_crn`; it is the standard Confluent collection filter, passed as the `?environment=` query parameter. 
    id := "id_example" // string | The unique identifier for the switchover endpoint.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SwitchoverEndpointsSwitchoverV1Api.GetSwitchoverV1SwitchoverEndpoint(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SwitchoverEndpointsSwitchoverV1Api.GetSwitchoverV1SwitchoverEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSwitchoverV1SwitchoverEndpoint`: SwitchoverV1SwitchoverEndpoint
    fmt.Fprintf(os.Stdout, "Response from `SwitchoverEndpointsSwitchoverV1Api.GetSwitchoverV1SwitchoverEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the switchover endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSwitchoverV1SwitchoverEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the environment with this ID. The bare-ID form of &#x60;environment_crn&#x60;; it is the standard Confluent collection filter, passed as the &#x60;?environment&#x3D;&#x60; query parameter.  | 


### Return type

[**SwitchoverV1SwitchoverEndpoint**](switchover.v1.SwitchoverEndpoint.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSwitchoverV1SwitchoverEndpoints

> SwitchoverV1SwitchoverEndpointList ListSwitchoverV1SwitchoverEndpoints(ctx).Environment(environment).SwitchoverPair(switchoverPair).PageSize(pageSize).PageToken(pageToken).Execute()

List of Switchover Endpoints



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
    environment := "env-00000" // string | Scope the operation to the environment with this ID. The bare-ID form of `environment_crn`; it is the standard Confluent collection filter, passed as the `?environment=` query parameter. 
    switchoverPair := "sw-00000" // string | Filter the results by the associated `SwitchoverPair` ID. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SwitchoverEndpointsSwitchoverV1Api.ListSwitchoverV1SwitchoverEndpoints(context.Background()).Environment(environment).SwitchoverPair(switchoverPair).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SwitchoverEndpointsSwitchoverV1Api.ListSwitchoverV1SwitchoverEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSwitchoverV1SwitchoverEndpoints`: SwitchoverV1SwitchoverEndpointList
    fmt.Fprintf(os.Stdout, "Response from `SwitchoverEndpointsSwitchoverV1Api.ListSwitchoverV1SwitchoverEndpoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSwitchoverV1SwitchoverEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the environment with this ID. The bare-ID form of &#x60;environment_crn&#x60;; it is the standard Confluent collection filter, passed as the &#x60;?environment&#x3D;&#x60; query parameter.  | 
 **switchoverPair** | **string** | Filter the results by the associated &#x60;SwitchoverPair&#x60; ID. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**SwitchoverV1SwitchoverEndpointList**](switchover.v1.SwitchoverEndpointList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSwitchoverV1SwitchoverEndpoint

> SwitchoverV1SwitchoverEndpoint UpdateSwitchoverV1SwitchoverEndpoint(ctx, id).Environment(environment).SwitchoverV1SwitchoverEndpointUpdateRequest(switchoverV1SwitchoverEndpointUpdateRequest).Execute()

Update a Switchover Endpoint



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
    environment := "env-00000" // string | Scope the operation to the environment with this ID. The bare-ID form of `environment_crn`; it is the standard Confluent collection filter, passed as the `?environment=` query parameter. 
    id := "id_example" // string | The unique identifier for the switchover endpoint.
    switchoverV1SwitchoverEndpointUpdateRequest := *openapiclient.NewSwitchoverV1SwitchoverEndpointUpdateRequest(*openapiclient.NewSwitchoverV1SwitchoverEndpointUpdateRequestSpec()) // SwitchoverV1SwitchoverEndpointUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SwitchoverEndpointsSwitchoverV1Api.UpdateSwitchoverV1SwitchoverEndpoint(context.Background(), id).Environment(environment).SwitchoverV1SwitchoverEndpointUpdateRequest(switchoverV1SwitchoverEndpointUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SwitchoverEndpointsSwitchoverV1Api.UpdateSwitchoverV1SwitchoverEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSwitchoverV1SwitchoverEndpoint`: SwitchoverV1SwitchoverEndpoint
    fmt.Fprintf(os.Stdout, "Response from `SwitchoverEndpointsSwitchoverV1Api.UpdateSwitchoverV1SwitchoverEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the switchover endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSwitchoverV1SwitchoverEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the environment with this ID. The bare-ID form of &#x60;environment_crn&#x60;; it is the standard Confluent collection filter, passed as the &#x60;?environment&#x3D;&#x60; query parameter.  | 

 **switchoverV1SwitchoverEndpointUpdateRequest** | [**SwitchoverV1SwitchoverEndpointUpdateRequest**](SwitchoverV1SwitchoverEndpointUpdateRequest.md) |  | 

### Return type

[**SwitchoverV1SwitchoverEndpoint**](switchover.v1.SwitchoverEndpoint.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

