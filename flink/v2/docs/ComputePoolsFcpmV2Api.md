# \ComputePoolsFcpmV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFcpmV2ComputePool**](ComputePoolsFcpmV2Api.md#CreateFcpmV2ComputePool) | **Post** /fcpm/v2/compute-pools | Create a Compute Pool
[**DeleteFcpmV2ComputePool**](ComputePoolsFcpmV2Api.md#DeleteFcpmV2ComputePool) | **Delete** /fcpm/v2/compute-pools/{id} | Delete a Compute Pool
[**GetFcpmV2ComputePool**](ComputePoolsFcpmV2Api.md#GetFcpmV2ComputePool) | **Get** /fcpm/v2/compute-pools/{id} | Read a Compute Pool
[**ListFcpmV2ComputePools**](ComputePoolsFcpmV2Api.md#ListFcpmV2ComputePools) | **Get** /fcpm/v2/compute-pools | List of Compute Pools
[**UpdateFcpmV2ComputePool**](ComputePoolsFcpmV2Api.md#UpdateFcpmV2ComputePool) | **Patch** /fcpm/v2/compute-pools/{id} | Update a Compute Pool



## CreateFcpmV2ComputePool

> FcpmV2ComputePool CreateFcpmV2ComputePool(ctx).FcpmV2ComputePool(fcpmV2ComputePool).Execute()

Create a Compute Pool



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
    fcpmV2ComputePool := *openapiclient.NewFcpmV2ComputePool() // FcpmV2ComputePool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComputePoolsFcpmV2Api.CreateFcpmV2ComputePool(context.Background()).FcpmV2ComputePool(fcpmV2ComputePool).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputePoolsFcpmV2Api.CreateFcpmV2ComputePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFcpmV2ComputePool`: FcpmV2ComputePool
    fmt.Fprintf(os.Stdout, "Response from `ComputePoolsFcpmV2Api.CreateFcpmV2ComputePool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFcpmV2ComputePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fcpmV2ComputePool** | [**FcpmV2ComputePool**](FcpmV2ComputePool.md) |  | 

### Return type

[**FcpmV2ComputePool**](fcpm.v2.ComputePool.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFcpmV2ComputePool

> DeleteFcpmV2ComputePool(ctx, id).Environment(environment).Execute()

Delete a Compute Pool



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
    id := "id_example" // string | The unique identifier for the compute pool.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComputePoolsFcpmV2Api.DeleteFcpmV2ComputePool(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputePoolsFcpmV2Api.DeleteFcpmV2ComputePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the compute pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFcpmV2ComputePoolRequest struct via the builder pattern


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


## GetFcpmV2ComputePool

> FcpmV2ComputePool GetFcpmV2ComputePool(ctx, id).Environment(environment).Execute()

Read a Compute Pool



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
    id := "id_example" // string | The unique identifier for the compute pool.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComputePoolsFcpmV2Api.GetFcpmV2ComputePool(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputePoolsFcpmV2Api.GetFcpmV2ComputePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFcpmV2ComputePool`: FcpmV2ComputePool
    fmt.Fprintf(os.Stdout, "Response from `ComputePoolsFcpmV2Api.GetFcpmV2ComputePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the compute pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFcpmV2ComputePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**FcpmV2ComputePool**](fcpm.v2.ComputePool.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFcpmV2ComputePools

> FcpmV2ComputePoolList ListFcpmV2ComputePools(ctx).Environment(environment).SpecRegion(specRegion).SpecNetwork(specNetwork).PageSize(pageSize).PageToken(pageToken).Execute()

List of Compute Pools



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
    specRegion := "us-west-1" // string | Filter the results by exact match for spec.region. (optional)
    specNetwork := "n-00000" // string | Filter the results by exact match for spec.network. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComputePoolsFcpmV2Api.ListFcpmV2ComputePools(context.Background()).Environment(environment).SpecRegion(specRegion).SpecNetwork(specNetwork).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputePoolsFcpmV2Api.ListFcpmV2ComputePools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFcpmV2ComputePools`: FcpmV2ComputePoolList
    fmt.Fprintf(os.Stdout, "Response from `ComputePoolsFcpmV2Api.ListFcpmV2ComputePools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFcpmV2ComputePoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **specRegion** | **string** | Filter the results by exact match for spec.region. | 
 **specNetwork** | **string** | Filter the results by exact match for spec.network. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**FcpmV2ComputePoolList**](fcpm.v2.ComputePoolList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFcpmV2ComputePool

> FcpmV2ComputePool UpdateFcpmV2ComputePool(ctx, id).FcpmV2ComputePoolUpdate(fcpmV2ComputePoolUpdate).Execute()

Update a Compute Pool



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
    id := "id_example" // string | The unique identifier for the compute pool.
    fcpmV2ComputePoolUpdate := *openapiclient.NewFcpmV2ComputePoolUpdate() // FcpmV2ComputePoolUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComputePoolsFcpmV2Api.UpdateFcpmV2ComputePool(context.Background(), id).FcpmV2ComputePoolUpdate(fcpmV2ComputePoolUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputePoolsFcpmV2Api.UpdateFcpmV2ComputePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFcpmV2ComputePool`: FcpmV2ComputePool
    fmt.Fprintf(os.Stdout, "Response from `ComputePoolsFcpmV2Api.UpdateFcpmV2ComputePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the compute pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFcpmV2ComputePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fcpmV2ComputePoolUpdate** | [**FcpmV2ComputePoolUpdate**](FcpmV2ComputePoolUpdate.md) |  | 

### Return type

[**FcpmV2ComputePool**](fcpm.v2.ComputePool.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

