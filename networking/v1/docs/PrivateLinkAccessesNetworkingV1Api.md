# \PrivateLinkAccessesNetworkingV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkingV1PrivateLinkAccess**](PrivateLinkAccessesNetworkingV1Api.md#CreateNetworkingV1PrivateLinkAccess) | **Post** /networking/v1/private-link-accesses | Create a Private Link Access
[**DeleteNetworkingV1PrivateLinkAccess**](PrivateLinkAccessesNetworkingV1Api.md#DeleteNetworkingV1PrivateLinkAccess) | **Delete** /networking/v1/private-link-accesses/{id} | Delete a Private Link Access
[**GetNetworkingV1PrivateLinkAccess**](PrivateLinkAccessesNetworkingV1Api.md#GetNetworkingV1PrivateLinkAccess) | **Get** /networking/v1/private-link-accesses/{id} | Read a Private Link Access
[**ListNetworkingV1PrivateLinkAccesses**](PrivateLinkAccessesNetworkingV1Api.md#ListNetworkingV1PrivateLinkAccesses) | **Get** /networking/v1/private-link-accesses | List of Private Link Accesses
[**UpdateNetworkingV1PrivateLinkAccess**](PrivateLinkAccessesNetworkingV1Api.md#UpdateNetworkingV1PrivateLinkAccess) | **Patch** /networking/v1/private-link-accesses/{id} | Update a Private Link Access



## CreateNetworkingV1PrivateLinkAccess

> NetworkingV1PrivateLinkAccess CreateNetworkingV1PrivateLinkAccess(ctx).NetworkingV1PrivateLinkAccess(networkingV1PrivateLinkAccess).Execute()

Create a Private Link Access



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
    networkingV1PrivateLinkAccess := *openapiclient.NewNetworkingV1PrivateLinkAccess() // NetworkingV1PrivateLinkAccess |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivateLinkAccessesNetworkingV1Api.CreateNetworkingV1PrivateLinkAccess(context.Background()).NetworkingV1PrivateLinkAccess(networkingV1PrivateLinkAccess).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateLinkAccessesNetworkingV1Api.CreateNetworkingV1PrivateLinkAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkingV1PrivateLinkAccess`: NetworkingV1PrivateLinkAccess
    fmt.Fprintf(os.Stdout, "Response from `PrivateLinkAccessesNetworkingV1Api.CreateNetworkingV1PrivateLinkAccess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkingV1PrivateLinkAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkingV1PrivateLinkAccess** | [**NetworkingV1PrivateLinkAccess**](NetworkingV1PrivateLinkAccess.md) |  | 

### Return type

[**NetworkingV1PrivateLinkAccess**](networking.v1.PrivateLinkAccess.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkingV1PrivateLinkAccess

> DeleteNetworkingV1PrivateLinkAccess(ctx, id).Environment(environment).Execute()

Delete a Private Link Access



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
    id := "id_example" // string | The unique identifier for the private link access.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivateLinkAccessesNetworkingV1Api.DeleteNetworkingV1PrivateLinkAccess(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateLinkAccessesNetworkingV1Api.DeleteNetworkingV1PrivateLinkAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the private link access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkingV1PrivateLinkAccessRequest struct via the builder pattern


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


## GetNetworkingV1PrivateLinkAccess

> NetworkingV1PrivateLinkAccess GetNetworkingV1PrivateLinkAccess(ctx, id).Environment(environment).Execute()

Read a Private Link Access



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
    id := "id_example" // string | The unique identifier for the private link access.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivateLinkAccessesNetworkingV1Api.GetNetworkingV1PrivateLinkAccess(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateLinkAccessesNetworkingV1Api.GetNetworkingV1PrivateLinkAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkingV1PrivateLinkAccess`: NetworkingV1PrivateLinkAccess
    fmt.Fprintf(os.Stdout, "Response from `PrivateLinkAccessesNetworkingV1Api.GetNetworkingV1PrivateLinkAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the private link access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkingV1PrivateLinkAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**NetworkingV1PrivateLinkAccess**](networking.v1.PrivateLinkAccess.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkingV1PrivateLinkAccesses

> NetworkingV1PrivateLinkAccessList ListNetworkingV1PrivateLinkAccesses(ctx).Environment(environment).SpecDisplayName(specDisplayName).StatusPhase(statusPhase).SpecNetwork(specNetwork).PageSize(pageSize).PageToken(pageToken).Execute()

List of Private Link Accesses



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
    specDisplayName := []string{"Inner_example"} // []string | Filter the results by exact match for spec.display_name. Pass multiple times to see results matching any of the values. (optional)
    statusPhase := []string{"Inner_example"} // []string | Filter the results by exact match for status.phase. Pass multiple times to see results matching any of the values. (optional)
    specNetwork := []string{"Inner_example"} // []string | Filter the results by exact match for spec.network. Pass multiple times to see results matching any of the values. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivateLinkAccessesNetworkingV1Api.ListNetworkingV1PrivateLinkAccesses(context.Background()).Environment(environment).SpecDisplayName(specDisplayName).StatusPhase(statusPhase).SpecNetwork(specNetwork).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateLinkAccessesNetworkingV1Api.ListNetworkingV1PrivateLinkAccesses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkingV1PrivateLinkAccesses`: NetworkingV1PrivateLinkAccessList
    fmt.Fprintf(os.Stdout, "Response from `PrivateLinkAccessesNetworkingV1Api.ListNetworkingV1PrivateLinkAccesses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkingV1PrivateLinkAccessesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **specDisplayName** | **[]string** | Filter the results by exact match for spec.display_name. Pass multiple times to see results matching any of the values. | 
 **statusPhase** | **[]string** | Filter the results by exact match for status.phase. Pass multiple times to see results matching any of the values. | 
 **specNetwork** | **[]string** | Filter the results by exact match for spec.network. Pass multiple times to see results matching any of the values. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**NetworkingV1PrivateLinkAccessList**](networking.v1.PrivateLinkAccessList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkingV1PrivateLinkAccess

> NetworkingV1PrivateLinkAccess UpdateNetworkingV1PrivateLinkAccess(ctx, id).NetworkingV1PrivateLinkAccessUpdate(networkingV1PrivateLinkAccessUpdate).Execute()

Update a Private Link Access



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
    id := "id_example" // string | The unique identifier for the private link access.
    networkingV1PrivateLinkAccessUpdate := *openapiclient.NewNetworkingV1PrivateLinkAccessUpdate() // NetworkingV1PrivateLinkAccessUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivateLinkAccessesNetworkingV1Api.UpdateNetworkingV1PrivateLinkAccess(context.Background(), id).NetworkingV1PrivateLinkAccessUpdate(networkingV1PrivateLinkAccessUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateLinkAccessesNetworkingV1Api.UpdateNetworkingV1PrivateLinkAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkingV1PrivateLinkAccess`: NetworkingV1PrivateLinkAccess
    fmt.Fprintf(os.Stdout, "Response from `PrivateLinkAccessesNetworkingV1Api.UpdateNetworkingV1PrivateLinkAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the private link access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkingV1PrivateLinkAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkingV1PrivateLinkAccessUpdate** | [**NetworkingV1PrivateLinkAccessUpdate**](NetworkingV1PrivateLinkAccessUpdate.md) |  | 

### Return type

[**NetworkingV1PrivateLinkAccess**](networking.v1.PrivateLinkAccess.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

