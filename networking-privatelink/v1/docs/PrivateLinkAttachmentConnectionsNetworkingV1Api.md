# \PrivateLinkAttachmentConnectionsNetworkingV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkingV1PrivateLinkAttachmentConnection**](PrivateLinkAttachmentConnectionsNetworkingV1Api.md#CreateNetworkingV1PrivateLinkAttachmentConnection) | **Post** /networking/v1/private-link-attachment-connections | Create a Private Link Attachment Connection
[**DeleteNetworkingV1PrivateLinkAttachmentConnection**](PrivateLinkAttachmentConnectionsNetworkingV1Api.md#DeleteNetworkingV1PrivateLinkAttachmentConnection) | **Delete** /networking/v1/private-link-attachment-connections/{id} | Delete a Private Link Attachment Connection
[**GetNetworkingV1PrivateLinkAttachmentConnection**](PrivateLinkAttachmentConnectionsNetworkingV1Api.md#GetNetworkingV1PrivateLinkAttachmentConnection) | **Get** /networking/v1/private-link-attachment-connections/{id} | Read a Private Link Attachment Connection
[**ListNetworkingV1PrivateLinkAttachmentConnections**](PrivateLinkAttachmentConnectionsNetworkingV1Api.md#ListNetworkingV1PrivateLinkAttachmentConnections) | **Get** /networking/v1/private-link-attachment-connections | List of Private Link Attachment Connections
[**UpdateNetworkingV1PrivateLinkAttachmentConnection**](PrivateLinkAttachmentConnectionsNetworkingV1Api.md#UpdateNetworkingV1PrivateLinkAttachmentConnection) | **Patch** /networking/v1/private-link-attachment-connections/{id} | Update a Private Link Attachment Connection



## CreateNetworkingV1PrivateLinkAttachmentConnection

> NetworkingV1PrivateLinkAttachmentConnection CreateNetworkingV1PrivateLinkAttachmentConnection(ctx).NetworkingV1PrivateLinkAttachmentConnection(networkingV1PrivateLinkAttachmentConnection).Execute()

Create a Private Link Attachment Connection



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
    networkingV1PrivateLinkAttachmentConnection := *openapiclient.NewNetworkingV1PrivateLinkAttachmentConnection() // NetworkingV1PrivateLinkAttachmentConnection |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivateLinkAttachmentConnectionsNetworkingV1Api.CreateNetworkingV1PrivateLinkAttachmentConnection(context.Background()).NetworkingV1PrivateLinkAttachmentConnection(networkingV1PrivateLinkAttachmentConnection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateLinkAttachmentConnectionsNetworkingV1Api.CreateNetworkingV1PrivateLinkAttachmentConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkingV1PrivateLinkAttachmentConnection`: NetworkingV1PrivateLinkAttachmentConnection
    fmt.Fprintf(os.Stdout, "Response from `PrivateLinkAttachmentConnectionsNetworkingV1Api.CreateNetworkingV1PrivateLinkAttachmentConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkingV1PrivateLinkAttachmentConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkingV1PrivateLinkAttachmentConnection** | [**NetworkingV1PrivateLinkAttachmentConnection**](NetworkingV1PrivateLinkAttachmentConnection.md) |  | 

### Return type

[**NetworkingV1PrivateLinkAttachmentConnection**](networking.v1.PrivateLinkAttachmentConnection.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkingV1PrivateLinkAttachmentConnection

> DeleteNetworkingV1PrivateLinkAttachmentConnection(ctx, id).Environment(environment).Execute()

Delete a Private Link Attachment Connection



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
    id := "id_example" // string | The unique identifier for the private link attachment connection.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivateLinkAttachmentConnectionsNetworkingV1Api.DeleteNetworkingV1PrivateLinkAttachmentConnection(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateLinkAttachmentConnectionsNetworkingV1Api.DeleteNetworkingV1PrivateLinkAttachmentConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the private link attachment connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkingV1PrivateLinkAttachmentConnectionRequest struct via the builder pattern


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


## GetNetworkingV1PrivateLinkAttachmentConnection

> NetworkingV1PrivateLinkAttachmentConnection GetNetworkingV1PrivateLinkAttachmentConnection(ctx, id).Environment(environment).Execute()

Read a Private Link Attachment Connection



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
    id := "id_example" // string | The unique identifier for the private link attachment connection.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivateLinkAttachmentConnectionsNetworkingV1Api.GetNetworkingV1PrivateLinkAttachmentConnection(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateLinkAttachmentConnectionsNetworkingV1Api.GetNetworkingV1PrivateLinkAttachmentConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkingV1PrivateLinkAttachmentConnection`: NetworkingV1PrivateLinkAttachmentConnection
    fmt.Fprintf(os.Stdout, "Response from `PrivateLinkAttachmentConnectionsNetworkingV1Api.GetNetworkingV1PrivateLinkAttachmentConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the private link attachment connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkingV1PrivateLinkAttachmentConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**NetworkingV1PrivateLinkAttachmentConnection**](networking.v1.PrivateLinkAttachmentConnection.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkingV1PrivateLinkAttachmentConnections

> NetworkingV1PrivateLinkAttachmentConnectionList ListNetworkingV1PrivateLinkAttachmentConnections(ctx).Environment(environment).SpecPrivateLinkAttachment(specPrivateLinkAttachment).PageSize(pageSize).PageToken(pageToken).Execute()

List of Private Link Attachment Connections



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
    specPrivateLinkAttachment := "platt-00000" // string | Filter the results by exact match for spec.private_link_attachment.
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivateLinkAttachmentConnectionsNetworkingV1Api.ListNetworkingV1PrivateLinkAttachmentConnections(context.Background()).Environment(environment).SpecPrivateLinkAttachment(specPrivateLinkAttachment).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateLinkAttachmentConnectionsNetworkingV1Api.ListNetworkingV1PrivateLinkAttachmentConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkingV1PrivateLinkAttachmentConnections`: NetworkingV1PrivateLinkAttachmentConnectionList
    fmt.Fprintf(os.Stdout, "Response from `PrivateLinkAttachmentConnectionsNetworkingV1Api.ListNetworkingV1PrivateLinkAttachmentConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkingV1PrivateLinkAttachmentConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **specPrivateLinkAttachment** | **string** | Filter the results by exact match for spec.private_link_attachment. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**NetworkingV1PrivateLinkAttachmentConnectionList**](networking.v1.PrivateLinkAttachmentConnectionList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkingV1PrivateLinkAttachmentConnection

> NetworkingV1PrivateLinkAttachmentConnection UpdateNetworkingV1PrivateLinkAttachmentConnection(ctx, id).NetworkingV1PrivateLinkAttachmentConnectionUpdate(networkingV1PrivateLinkAttachmentConnectionUpdate).Execute()

Update a Private Link Attachment Connection



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
    id := "id_example" // string | The unique identifier for the private link attachment connection.
    networkingV1PrivateLinkAttachmentConnectionUpdate := *openapiclient.NewNetworkingV1PrivateLinkAttachmentConnectionUpdate() // NetworkingV1PrivateLinkAttachmentConnectionUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivateLinkAttachmentConnectionsNetworkingV1Api.UpdateNetworkingV1PrivateLinkAttachmentConnection(context.Background(), id).NetworkingV1PrivateLinkAttachmentConnectionUpdate(networkingV1PrivateLinkAttachmentConnectionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateLinkAttachmentConnectionsNetworkingV1Api.UpdateNetworkingV1PrivateLinkAttachmentConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkingV1PrivateLinkAttachmentConnection`: NetworkingV1PrivateLinkAttachmentConnection
    fmt.Fprintf(os.Stdout, "Response from `PrivateLinkAttachmentConnectionsNetworkingV1Api.UpdateNetworkingV1PrivateLinkAttachmentConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the private link attachment connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkingV1PrivateLinkAttachmentConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkingV1PrivateLinkAttachmentConnectionUpdate** | [**NetworkingV1PrivateLinkAttachmentConnectionUpdate**](NetworkingV1PrivateLinkAttachmentConnectionUpdate.md) |  | 

### Return type

[**NetworkingV1PrivateLinkAttachmentConnection**](networking.v1.PrivateLinkAttachmentConnection.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

