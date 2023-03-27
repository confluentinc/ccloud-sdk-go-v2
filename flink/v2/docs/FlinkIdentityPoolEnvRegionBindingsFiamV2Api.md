# \FlinkIdentityPoolEnvRegionBindingsFiamV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFiamV2FlinkIdentityPoolEnvRegionBinding**](FlinkIdentityPoolEnvRegionBindingsFiamV2Api.md#CreateFiamV2FlinkIdentityPoolEnvRegionBinding) | **Post** /fiam/v2/iam-bindings | Create a Flink Identity Pool Env Region Binding
[**DeleteFiamV2FlinkIdentityPoolEnvRegionBinding**](FlinkIdentityPoolEnvRegionBindingsFiamV2Api.md#DeleteFiamV2FlinkIdentityPoolEnvRegionBinding) | **Delete** /fiam/v2/iam-bindings/{id} | Delete a Flink Identity Pool Env Region Binding
[**ListFiamV2FlinkIdentityPoolEnvRegionBindings**](FlinkIdentityPoolEnvRegionBindingsFiamV2Api.md#ListFiamV2FlinkIdentityPoolEnvRegionBindings) | **Get** /fiam/v2/iam-bindings | List of Flink Identity Pool Env Region Bindings



## CreateFiamV2FlinkIdentityPoolEnvRegionBinding

> FiamV2FlinkIdentityPoolEnvRegionBinding CreateFiamV2FlinkIdentityPoolEnvRegionBinding(ctx).IdentityPool(identityPool).Environment(environment).FiamV2FlinkIdentityPoolEnvRegionBinding(fiamV2FlinkIdentityPoolEnvRegionBinding).Execute()

Create a Flink Identity Pool Env Region Binding



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
    identityPool := "identityPool_example" // string | Scope the operation to the given identity_pool.
    environment := "env-00000" // string | Scope the operation to the given environment.
    fiamV2FlinkIdentityPoolEnvRegionBinding := *openapiclient.NewFiamV2FlinkIdentityPoolEnvRegionBinding() // FiamV2FlinkIdentityPoolEnvRegionBinding |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlinkIdentityPoolEnvRegionBindingsFiamV2Api.CreateFiamV2FlinkIdentityPoolEnvRegionBinding(context.Background()).IdentityPool(identityPool).Environment(environment).FiamV2FlinkIdentityPoolEnvRegionBinding(fiamV2FlinkIdentityPoolEnvRegionBinding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlinkIdentityPoolEnvRegionBindingsFiamV2Api.CreateFiamV2FlinkIdentityPoolEnvRegionBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFiamV2FlinkIdentityPoolEnvRegionBinding`: FiamV2FlinkIdentityPoolEnvRegionBinding
    fmt.Fprintf(os.Stdout, "Response from `FlinkIdentityPoolEnvRegionBindingsFiamV2Api.CreateFiamV2FlinkIdentityPoolEnvRegionBinding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFiamV2FlinkIdentityPoolEnvRegionBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityPool** | **string** | Scope the operation to the given identity_pool. | 
 **environment** | **string** | Scope the operation to the given environment. | 
 **fiamV2FlinkIdentityPoolEnvRegionBinding** | [**FiamV2FlinkIdentityPoolEnvRegionBinding**](FiamV2FlinkIdentityPoolEnvRegionBinding.md) |  | 

### Return type

[**FiamV2FlinkIdentityPoolEnvRegionBinding**](fiam.v2.FlinkIdentityPoolEnvRegionBinding.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFiamV2FlinkIdentityPoolEnvRegionBinding

> DeleteFiamV2FlinkIdentityPoolEnvRegionBinding(ctx, id).IdentityPool(identityPool).Environment(environment).Execute()

Delete a Flink Identity Pool Env Region Binding



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
    identityPool := "identityPool_example" // string | Scope the operation to the given identity_pool.
    environment := "env-00000" // string | Scope the operation to the given environment.
    id := "id_example" // string | The unique identifier for the flink identity pool env region binding.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlinkIdentityPoolEnvRegionBindingsFiamV2Api.DeleteFiamV2FlinkIdentityPoolEnvRegionBinding(context.Background(), id).IdentityPool(identityPool).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlinkIdentityPoolEnvRegionBindingsFiamV2Api.DeleteFiamV2FlinkIdentityPoolEnvRegionBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the flink identity pool env region binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFiamV2FlinkIdentityPoolEnvRegionBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityPool** | **string** | Scope the operation to the given identity_pool. | 
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


## ListFiamV2FlinkIdentityPoolEnvRegionBindings

> FiamV2FlinkIdentityPoolEnvRegionBindingList ListFiamV2FlinkIdentityPoolEnvRegionBindings(ctx).Region(region).IdentityPool(identityPool).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()

List of Flink Identity Pool Env Region Bindings



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
    region := "us-west-1" // string | Filter the results by exact match for region.
    identityPool := "identityPool_example" // string | Filter the results by exact match for identity_pool.
    environment := "env-00000" // string | Filter the results by exact match for environment.
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlinkIdentityPoolEnvRegionBindingsFiamV2Api.ListFiamV2FlinkIdentityPoolEnvRegionBindings(context.Background()).Region(region).IdentityPool(identityPool).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlinkIdentityPoolEnvRegionBindingsFiamV2Api.ListFiamV2FlinkIdentityPoolEnvRegionBindings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFiamV2FlinkIdentityPoolEnvRegionBindings`: FiamV2FlinkIdentityPoolEnvRegionBindingList
    fmt.Fprintf(os.Stdout, "Response from `FlinkIdentityPoolEnvRegionBindingsFiamV2Api.ListFiamV2FlinkIdentityPoolEnvRegionBindings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFiamV2FlinkIdentityPoolEnvRegionBindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **string** | Filter the results by exact match for region. | 
 **identityPool** | **string** | Filter the results by exact match for identity_pool. | 
 **environment** | **string** | Filter the results by exact match for environment. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**FiamV2FlinkIdentityPoolEnvRegionBindingList**](fiam.v2.FlinkIdentityPoolEnvRegionBindingList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

