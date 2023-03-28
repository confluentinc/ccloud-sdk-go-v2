# \IdentityPoolEnvRegionBindingsFcpmV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFcpmV2IdentityPoolEnvRegionBinding**](IdentityPoolEnvRegionBindingsFcpmV2Api.md#CreateFcpmV2IdentityPoolEnvRegionBinding) | **Post** /fcpm/v2/identity-pool-env-region-bindings | Create an Identity Pool Env Region Binding
[**DeleteFcpmV2IdentityPoolEnvRegionBinding**](IdentityPoolEnvRegionBindingsFcpmV2Api.md#DeleteFcpmV2IdentityPoolEnvRegionBinding) | **Delete** /fcpm/v2/identity-pool-env-region-bindings/{id} | Delete an Identity Pool Env Region Binding
[**ListFcpmV2IdentityPoolEnvRegionBindings**](IdentityPoolEnvRegionBindingsFcpmV2Api.md#ListFcpmV2IdentityPoolEnvRegionBindings) | **Get** /fcpm/v2/identity-pool-env-region-bindings | List of Identity Pool Env Region Bindings



## CreateFcpmV2IdentityPoolEnvRegionBinding

> FcpmV2IdentityPoolEnvRegionBinding CreateFcpmV2IdentityPoolEnvRegionBinding(ctx).FcpmV2IdentityPoolEnvRegionBinding(fcpmV2IdentityPoolEnvRegionBinding).Execute()

Create an Identity Pool Env Region Binding



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
    fcpmV2IdentityPoolEnvRegionBinding := *openapiclient.NewFcpmV2IdentityPoolEnvRegionBinding() // FcpmV2IdentityPoolEnvRegionBinding |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityPoolEnvRegionBindingsFcpmV2Api.CreateFcpmV2IdentityPoolEnvRegionBinding(context.Background()).FcpmV2IdentityPoolEnvRegionBinding(fcpmV2IdentityPoolEnvRegionBinding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPoolEnvRegionBindingsFcpmV2Api.CreateFcpmV2IdentityPoolEnvRegionBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFcpmV2IdentityPoolEnvRegionBinding`: FcpmV2IdentityPoolEnvRegionBinding
    fmt.Fprintf(os.Stdout, "Response from `IdentityPoolEnvRegionBindingsFcpmV2Api.CreateFcpmV2IdentityPoolEnvRegionBinding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFcpmV2IdentityPoolEnvRegionBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fcpmV2IdentityPoolEnvRegionBinding** | [**FcpmV2IdentityPoolEnvRegionBinding**](FcpmV2IdentityPoolEnvRegionBinding.md) |  | 

### Return type

[**FcpmV2IdentityPoolEnvRegionBinding**](fcpm.v2.IdentityPoolEnvRegionBinding.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFcpmV2IdentityPoolEnvRegionBinding

> DeleteFcpmV2IdentityPoolEnvRegionBinding(ctx, id).Environment(environment).IdentityPool(identityPool).Execute()

Delete an Identity Pool Env Region Binding



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
    id := "id_example" // string | The unique identifier for the identity pool env region binding.
    identityPool := "iam-00000" // string | Scope the operation to the given identity_pool. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityPoolEnvRegionBindingsFcpmV2Api.DeleteFcpmV2IdentityPoolEnvRegionBinding(context.Background(), id).Environment(environment).IdentityPool(identityPool).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPoolEnvRegionBindingsFcpmV2Api.DeleteFcpmV2IdentityPoolEnvRegionBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the identity pool env region binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFcpmV2IdentityPoolEnvRegionBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 

 **identityPool** | **string** | Scope the operation to the given identity_pool. | 

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


## ListFcpmV2IdentityPoolEnvRegionBindings

> FcpmV2IdentityPoolEnvRegionBindingList ListFcpmV2IdentityPoolEnvRegionBindings(ctx).Environment(environment).Region(region).Cloud(cloud).IdentityPool(identityPool).PageSize(pageSize).PageToken(pageToken).Execute()

List of Identity Pool Env Region Bindings



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
    region := "us-west-1" // string | Filter the results by exact match for region. (optional)
    cloud := "AWS" // string | Filter the results by exact match for cloud. (optional)
    identityPool := "iam-00000" // string | Filter the results by exact match for identity_pool. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityPoolEnvRegionBindingsFcpmV2Api.ListFcpmV2IdentityPoolEnvRegionBindings(context.Background()).Environment(environment).Region(region).Cloud(cloud).IdentityPool(identityPool).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPoolEnvRegionBindingsFcpmV2Api.ListFcpmV2IdentityPoolEnvRegionBindings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFcpmV2IdentityPoolEnvRegionBindings`: FcpmV2IdentityPoolEnvRegionBindingList
    fmt.Fprintf(os.Stdout, "Response from `IdentityPoolEnvRegionBindingsFcpmV2Api.ListFcpmV2IdentityPoolEnvRegionBindings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFcpmV2IdentityPoolEnvRegionBindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **region** | **string** | Filter the results by exact match for region. | 
 **cloud** | **string** | Filter the results by exact match for cloud. | 
 **identityPool** | **string** | Filter the results by exact match for identity_pool. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**FcpmV2IdentityPoolEnvRegionBindingList**](fcpm.v2.IdentityPoolEnvRegionBindingList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

