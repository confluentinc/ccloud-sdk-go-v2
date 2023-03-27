# \ResourcePoolsFrpmV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFrpmV2ResourcePool**](ResourcePoolsFrpmV2Api.md#CreateFrpmV2ResourcePool) | **Post** /frpm/v2/resource-pools | Create a Resource Pool
[**DeleteFrpmV2ResourcePool**](ResourcePoolsFrpmV2Api.md#DeleteFrpmV2ResourcePool) | **Delete** /frpm/v2/resource-pools/{id} | Delete a Resource Pool
[**GetFrpmV2ResourcePool**](ResourcePoolsFrpmV2Api.md#GetFrpmV2ResourcePool) | **Get** /frpm/v2/resource-pools/{id} | Read a Resource Pool
[**ListFrpmV2ResourcePools**](ResourcePoolsFrpmV2Api.md#ListFrpmV2ResourcePools) | **Get** /frpm/v2/resource-pools | List of Resource Pools
[**UpdateFrpmV2ResourcePool**](ResourcePoolsFrpmV2Api.md#UpdateFrpmV2ResourcePool) | **Patch** /frpm/v2/resource-pools/{id} | Update a Resource Pool



## CreateFrpmV2ResourcePool

> FrpmV2ResourcePool CreateFrpmV2ResourcePool(ctx).FrpmV2ResourcePool(frpmV2ResourcePool).Execute()

Create a Resource Pool



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
    frpmV2ResourcePool := *openapiclient.NewFrpmV2ResourcePool() // FrpmV2ResourcePool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcePoolsFrpmV2Api.CreateFrpmV2ResourcePool(context.Background()).FrpmV2ResourcePool(frpmV2ResourcePool).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolsFrpmV2Api.CreateFrpmV2ResourcePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFrpmV2ResourcePool`: FrpmV2ResourcePool
    fmt.Fprintf(os.Stdout, "Response from `ResourcePoolsFrpmV2Api.CreateFrpmV2ResourcePool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFrpmV2ResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **frpmV2ResourcePool** | [**FrpmV2ResourcePool**](FrpmV2ResourcePool.md) |  | 

### Return type

[**FrpmV2ResourcePool**](frpm.v2.ResourcePool.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFrpmV2ResourcePool

> DeleteFrpmV2ResourcePool(ctx, id).Execute()

Delete a Resource Pool



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
    id := "id_example" // string | The unique identifier for the resource pool.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcePoolsFrpmV2Api.DeleteFrpmV2ResourcePool(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolsFrpmV2Api.DeleteFrpmV2ResourcePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the resource pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFrpmV2ResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetFrpmV2ResourcePool

> FrpmV2ResourcePool GetFrpmV2ResourcePool(ctx, id).Execute()

Read a Resource Pool



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
    id := "id_example" // string | The unique identifier for the resource pool.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcePoolsFrpmV2Api.GetFrpmV2ResourcePool(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolsFrpmV2Api.GetFrpmV2ResourcePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFrpmV2ResourcePool`: FrpmV2ResourcePool
    fmt.Fprintf(os.Stdout, "Response from `ResourcePoolsFrpmV2Api.GetFrpmV2ResourcePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the resource pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFrpmV2ResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FrpmV2ResourcePool**](frpm.v2.ResourcePool.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFrpmV2ResourcePools

> FrpmV2ResourcePoolList ListFrpmV2ResourcePools(ctx).SpecRegion(specRegion).Environment(environment).SpecNetwork(specNetwork).PageSize(pageSize).PageToken(pageToken).Execute()

List of Resource Pools



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
    specRegion := "us-west-1" // string | Filter the results by exact match for spec.region.
    environment := "env-00000" // string | Filter the results by exact match for environment. (optional)
    specNetwork := "n-00000" // string | Filter the results by exact match for spec.network. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcePoolsFrpmV2Api.ListFrpmV2ResourcePools(context.Background()).SpecRegion(specRegion).Environment(environment).SpecNetwork(specNetwork).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolsFrpmV2Api.ListFrpmV2ResourcePools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFrpmV2ResourcePools`: FrpmV2ResourcePoolList
    fmt.Fprintf(os.Stdout, "Response from `ResourcePoolsFrpmV2Api.ListFrpmV2ResourcePools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFrpmV2ResourcePoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **specRegion** | **string** | Filter the results by exact match for spec.region. | 
 **environment** | **string** | Filter the results by exact match for environment. | 
 **specNetwork** | **string** | Filter the results by exact match for spec.network. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**FrpmV2ResourcePoolList**](frpm.v2.ResourcePoolList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFrpmV2ResourcePool

> FrpmV2ResourcePool UpdateFrpmV2ResourcePool(ctx, id).FrpmV2ResourcePoolUpdate(frpmV2ResourcePoolUpdate).Execute()

Update a Resource Pool



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
    id := "id_example" // string | The unique identifier for the resource pool.
    frpmV2ResourcePoolUpdate := *openapiclient.NewFrpmV2ResourcePoolUpdate() // FrpmV2ResourcePoolUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcePoolsFrpmV2Api.UpdateFrpmV2ResourcePool(context.Background(), id).FrpmV2ResourcePoolUpdate(frpmV2ResourcePoolUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolsFrpmV2Api.UpdateFrpmV2ResourcePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFrpmV2ResourcePool`: FrpmV2ResourcePool
    fmt.Fprintf(os.Stdout, "Response from `ResourcePoolsFrpmV2Api.UpdateFrpmV2ResourcePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the resource pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFrpmV2ResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **frpmV2ResourcePoolUpdate** | [**FrpmV2ResourcePoolUpdate**](FrpmV2ResourcePoolUpdate.md) |  | 

### Return type

[**FrpmV2ResourcePool**](frpm.v2.ResourcePool.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

