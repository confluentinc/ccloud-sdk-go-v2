# \ClustersFlinkcmV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFlinkcmV2Cluster**](ClustersFlinkcmV2Api.md#CreateFlinkcmV2Cluster) | **Post** /flinkcm/v2/clusters | Create a Cluster
[**DeleteFlinkcmV2Cluster**](ClustersFlinkcmV2Api.md#DeleteFlinkcmV2Cluster) | **Delete** /flinkcm/v2/clusters/{id} | Delete a Cluster
[**GetFlinkcmV2Cluster**](ClustersFlinkcmV2Api.md#GetFlinkcmV2Cluster) | **Get** /flinkcm/v2/clusters/{id} | Read a Cluster
[**ListFlinkcmV2Clusters**](ClustersFlinkcmV2Api.md#ListFlinkcmV2Clusters) | **Get** /flinkcm/v2/clusters | List of Clusters
[**UpdateFlinkcmV2Cluster**](ClustersFlinkcmV2Api.md#UpdateFlinkcmV2Cluster) | **Patch** /flinkcm/v2/clusters/{id} | Update a Cluster



## CreateFlinkcmV2Cluster

> FlinkcmV2Cluster CreateFlinkcmV2Cluster(ctx).FlinkcmV2Cluster(flinkcmV2Cluster).Execute()

Create a Cluster



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
    flinkcmV2Cluster := *openapiclient.NewFlinkcmV2Cluster() // FlinkcmV2Cluster |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersFlinkcmV2Api.CreateFlinkcmV2Cluster(context.Background()).FlinkcmV2Cluster(flinkcmV2Cluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersFlinkcmV2Api.CreateFlinkcmV2Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFlinkcmV2Cluster`: FlinkcmV2Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersFlinkcmV2Api.CreateFlinkcmV2Cluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFlinkcmV2ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flinkcmV2Cluster** | [**FlinkcmV2Cluster**](FlinkcmV2Cluster.md) |  | 

### Return type

[**FlinkcmV2Cluster**](flinkcm.v2.Cluster.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFlinkcmV2Cluster

> DeleteFlinkcmV2Cluster(ctx, id).Environment(environment).Execute()

Delete a Cluster



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
    id := "id_example" // string | The unique identifier for the cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersFlinkcmV2Api.DeleteFlinkcmV2Cluster(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersFlinkcmV2Api.DeleteFlinkcmV2Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFlinkcmV2ClusterRequest struct via the builder pattern


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


## GetFlinkcmV2Cluster

> FlinkcmV2Cluster GetFlinkcmV2Cluster(ctx, id).Environment(environment).Execute()

Read a Cluster



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
    id := "id_example" // string | The unique identifier for the cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersFlinkcmV2Api.GetFlinkcmV2Cluster(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersFlinkcmV2Api.GetFlinkcmV2Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlinkcmV2Cluster`: FlinkcmV2Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersFlinkcmV2Api.GetFlinkcmV2Cluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlinkcmV2ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**FlinkcmV2Cluster**](flinkcm.v2.Cluster.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFlinkcmV2Clusters

> FlinkcmV2ClusterList ListFlinkcmV2Clusters(ctx).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()

List of Clusters



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
    resp, r, err := api_client.ClustersFlinkcmV2Api.ListFlinkcmV2Clusters(context.Background()).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersFlinkcmV2Api.ListFlinkcmV2Clusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFlinkcmV2Clusters`: FlinkcmV2ClusterList
    fmt.Fprintf(os.Stdout, "Response from `ClustersFlinkcmV2Api.ListFlinkcmV2Clusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFlinkcmV2ClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**FlinkcmV2ClusterList**](flinkcm.v2.ClusterList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFlinkcmV2Cluster

> FlinkcmV2Cluster UpdateFlinkcmV2Cluster(ctx, id).FlinkcmV2ClusterUpdate(flinkcmV2ClusterUpdate).Execute()

Update a Cluster



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
    id := "id_example" // string | The unique identifier for the cluster.
    flinkcmV2ClusterUpdate := *openapiclient.NewFlinkcmV2ClusterUpdate() // FlinkcmV2ClusterUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersFlinkcmV2Api.UpdateFlinkcmV2Cluster(context.Background(), id).FlinkcmV2ClusterUpdate(flinkcmV2ClusterUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersFlinkcmV2Api.UpdateFlinkcmV2Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFlinkcmV2Cluster`: FlinkcmV2Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersFlinkcmV2Api.UpdateFlinkcmV2Cluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFlinkcmV2ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flinkcmV2ClusterUpdate** | [**FlinkcmV2ClusterUpdate**](FlinkcmV2ClusterUpdate.md) |  | 

### Return type

[**FlinkcmV2Cluster**](flinkcm.v2.Cluster.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

