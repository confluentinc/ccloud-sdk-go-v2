# \PrivateLinkAccessesNetworkingAdminV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkingAdminV1PrivateLinkAccess**](PrivateLinkAccessesNetworkingAdminV1Api.md#GetNetworkingAdminV1PrivateLinkAccess) | **Get** /networking-admin/v1/private-link-accesses/{id} | Read a Private Link Access
[**ListNetworkingAdminV1PrivateLinkAccesses**](PrivateLinkAccessesNetworkingAdminV1Api.md#ListNetworkingAdminV1PrivateLinkAccesses) | **Get** /networking-admin/v1/private-link-accesses | List of Private Link Accesses



## GetNetworkingAdminV1PrivateLinkAccess

> NetworkingAdminV1PrivateLinkAccess GetNetworkingAdminV1PrivateLinkAccess(ctx, id).Environment(environment).Execute()

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
    resp, r, err := api_client.PrivateLinkAccessesNetworkingAdminV1Api.GetNetworkingAdminV1PrivateLinkAccess(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateLinkAccessesNetworkingAdminV1Api.GetNetworkingAdminV1PrivateLinkAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkingAdminV1PrivateLinkAccess`: NetworkingAdminV1PrivateLinkAccess
    fmt.Fprintf(os.Stdout, "Response from `PrivateLinkAccessesNetworkingAdminV1Api.GetNetworkingAdminV1PrivateLinkAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the private link access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkingAdminV1PrivateLinkAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**NetworkingAdminV1PrivateLinkAccess**](networking-admin.v1.PrivateLinkAccess.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkingAdminV1PrivateLinkAccesses

> NetworkingAdminV1PrivateLinkAccessList ListNetworkingAdminV1PrivateLinkAccesses(ctx).Environment(environment).SpecDisplayName(specDisplayName).StatusPhase(statusPhase).SpecNetwork(specNetwork).PageSize(pageSize).PageToken(pageToken).Execute()

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
    specDisplayName := *openapiclient.NewMultipleSearchFilter() // MultipleSearchFilter | Filter the results by exact match for spec.display_name. Pass multiple times to see results matching any of the values. (optional)
    statusPhase := *openapiclient.NewMultipleSearchFilter() // MultipleSearchFilter | Filter the results by exact match for status.phase. Pass multiple times to see results matching any of the values. (optional)
    specNetwork :=  // MultipleSearchFilter | Filter the results by exact match for spec.network. Pass multiple times to see results matching any of the values. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivateLinkAccessesNetworkingAdminV1Api.ListNetworkingAdminV1PrivateLinkAccesses(context.Background()).Environment(environment).SpecDisplayName(specDisplayName).StatusPhase(statusPhase).SpecNetwork(specNetwork).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateLinkAccessesNetworkingAdminV1Api.ListNetworkingAdminV1PrivateLinkAccesses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkingAdminV1PrivateLinkAccesses`: NetworkingAdminV1PrivateLinkAccessList
    fmt.Fprintf(os.Stdout, "Response from `PrivateLinkAccessesNetworkingAdminV1Api.ListNetworkingAdminV1PrivateLinkAccesses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkingAdminV1PrivateLinkAccessesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **specDisplayName** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter the results by exact match for spec.display_name. Pass multiple times to see results matching any of the values. | 
 **statusPhase** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter the results by exact match for status.phase. Pass multiple times to see results matching any of the values. | 
 **specNetwork** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter the results by exact match for spec.network. Pass multiple times to see results matching any of the values. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**NetworkingAdminV1PrivateLinkAccessList**](networking-admin.v1.PrivateLinkAccessList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

