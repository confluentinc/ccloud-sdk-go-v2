# \DnsForwardersNetworkingV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkingV1DnsForwarder**](DnsForwardersNetworkingV1Api.md#CreateNetworkingV1DnsForwarder) | **Post** /networking/v1/dns-forwarders | Create a Dns Forwarder
[**DeleteNetworkingV1DnsForwarder**](DnsForwardersNetworkingV1Api.md#DeleteNetworkingV1DnsForwarder) | **Delete** /networking/v1/dns-forwarders/{id} | Delete a Dns Forwarder
[**GetNetworkingV1DnsForwarder**](DnsForwardersNetworkingV1Api.md#GetNetworkingV1DnsForwarder) | **Get** /networking/v1/dns-forwarders/{id} | Read a Dns Forwarder
[**ListNetworkingV1DnsForwarders**](DnsForwardersNetworkingV1Api.md#ListNetworkingV1DnsForwarders) | **Get** /networking/v1/dns-forwarders | List of Dns Forwarders
[**UpdateNetworkingV1DnsForwarder**](DnsForwardersNetworkingV1Api.md#UpdateNetworkingV1DnsForwarder) | **Patch** /networking/v1/dns-forwarders/{id} | Update a Dns Forwarder



## CreateNetworkingV1DnsForwarder

> NetworkingV1DnsForwarder CreateNetworkingV1DnsForwarder(ctx).NetworkingV1DnsForwarder(networkingV1DnsForwarder).Execute()

Create a Dns Forwarder



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
    networkingV1DnsForwarder := *openapiclient.NewNetworkingV1DnsForwarder() // NetworkingV1DnsForwarder |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsForwardersNetworkingV1Api.CreateNetworkingV1DnsForwarder(context.Background()).NetworkingV1DnsForwarder(networkingV1DnsForwarder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsForwardersNetworkingV1Api.CreateNetworkingV1DnsForwarder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkingV1DnsForwarder`: NetworkingV1DnsForwarder
    fmt.Fprintf(os.Stdout, "Response from `DnsForwardersNetworkingV1Api.CreateNetworkingV1DnsForwarder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkingV1DnsForwarderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkingV1DnsForwarder** | [**NetworkingV1DnsForwarder**](NetworkingV1DnsForwarder.md) |  | 

### Return type

[**NetworkingV1DnsForwarder**](networking.v1.DnsForwarder.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkingV1DnsForwarder

> DeleteNetworkingV1DnsForwarder(ctx, id).Environment(environment).Execute()

Delete a Dns Forwarder



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
    id := "id_example" // string | The unique identifier for the dns forwarder.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsForwardersNetworkingV1Api.DeleteNetworkingV1DnsForwarder(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsForwardersNetworkingV1Api.DeleteNetworkingV1DnsForwarder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the dns forwarder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkingV1DnsForwarderRequest struct via the builder pattern


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


## GetNetworkingV1DnsForwarder

> NetworkingV1DnsForwarder GetNetworkingV1DnsForwarder(ctx, id).Environment(environment).Execute()

Read a Dns Forwarder



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
    id := "id_example" // string | The unique identifier for the dns forwarder.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsForwardersNetworkingV1Api.GetNetworkingV1DnsForwarder(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsForwardersNetworkingV1Api.GetNetworkingV1DnsForwarder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkingV1DnsForwarder`: NetworkingV1DnsForwarder
    fmt.Fprintf(os.Stdout, "Response from `DnsForwardersNetworkingV1Api.GetNetworkingV1DnsForwarder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the dns forwarder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkingV1DnsForwarderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**NetworkingV1DnsForwarder**](networking.v1.DnsForwarder.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkingV1DnsForwarders

> NetworkingV1DnsForwarderList ListNetworkingV1DnsForwarders(ctx).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()

List of Dns Forwarders



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
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsForwardersNetworkingV1Api.ListNetworkingV1DnsForwarders(context.Background()).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsForwardersNetworkingV1Api.ListNetworkingV1DnsForwarders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkingV1DnsForwarders`: NetworkingV1DnsForwarderList
    fmt.Fprintf(os.Stdout, "Response from `DnsForwardersNetworkingV1Api.ListNetworkingV1DnsForwarders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkingV1DnsForwardersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**NetworkingV1DnsForwarderList**](networking.v1.DnsForwarderList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkingV1DnsForwarder

> NetworkingV1DnsForwarder UpdateNetworkingV1DnsForwarder(ctx, id).NetworkingV1DnsForwarderUpdate(networkingV1DnsForwarderUpdate).Execute()

Update a Dns Forwarder



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
    id := "id_example" // string | The unique identifier for the dns forwarder.
    networkingV1DnsForwarderUpdate := *openapiclient.NewNetworkingV1DnsForwarderUpdate() // NetworkingV1DnsForwarderUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsForwardersNetworkingV1Api.UpdateNetworkingV1DnsForwarder(context.Background(), id).NetworkingV1DnsForwarderUpdate(networkingV1DnsForwarderUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsForwardersNetworkingV1Api.UpdateNetworkingV1DnsForwarder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkingV1DnsForwarder`: NetworkingV1DnsForwarder
    fmt.Fprintf(os.Stdout, "Response from `DnsForwardersNetworkingV1Api.UpdateNetworkingV1DnsForwarder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the dns forwarder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkingV1DnsForwarderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkingV1DnsForwarderUpdate** | [**NetworkingV1DnsForwarderUpdate**](NetworkingV1DnsForwarderUpdate.md) |  | 

### Return type

[**NetworkingV1DnsForwarder**](networking.v1.DnsForwarder.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

