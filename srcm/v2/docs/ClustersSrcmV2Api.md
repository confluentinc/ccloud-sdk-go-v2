# \ClustersSrcmV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSrcmV2Cluster**](ClustersSrcmV2Api.md#CreateSrcmV2Cluster) | **Post** /srcm/v2/clusters | Create a Cluster
[**DeleteSrcmV2Cluster**](ClustersSrcmV2Api.md#DeleteSrcmV2Cluster) | **Delete** /srcm/v2/clusters/{id} | Delete a Cluster
[**GetSrcmV2Cluster**](ClustersSrcmV2Api.md#GetSrcmV2Cluster) | **Get** /srcm/v2/clusters/{id} | Read a Cluster
[**ListSrcmV2Clusters**](ClustersSrcmV2Api.md#ListSrcmV2Clusters) | **Get** /srcm/v2/clusters | List of Clusters
[**UpdateSrcmV2Cluster**](ClustersSrcmV2Api.md#UpdateSrcmV2Cluster) | **Patch** /srcm/v2/clusters/{id} | Update a Cluster



## CreateSrcmV2Cluster

> SrcmV2Cluster CreateSrcmV2Cluster(ctx).SrcmV2Cluster(srcmV2Cluster).Execute()

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
    srcmV2Cluster := *openapiclient.NewSrcmV2Cluster() // SrcmV2Cluster |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersSrcmV2Api.CreateSrcmV2Cluster(context.Background()).SrcmV2Cluster(srcmV2Cluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersSrcmV2Api.CreateSrcmV2Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSrcmV2Cluster`: SrcmV2Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersSrcmV2Api.CreateSrcmV2Cluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSrcmV2ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcmV2Cluster** | [**SrcmV2Cluster**](SrcmV2Cluster.md) |  | 

### Return type

[**SrcmV2Cluster**](srcm.v2.Cluster.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSrcmV2Cluster

> DeleteSrcmV2Cluster(ctx, id).Environment(environment).Execute()

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
    resp, r, err := api_client.ClustersSrcmV2Api.DeleteSrcmV2Cluster(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersSrcmV2Api.DeleteSrcmV2Cluster``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSrcmV2ClusterRequest struct via the builder pattern


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


## GetSrcmV2Cluster

> SrcmV2Cluster GetSrcmV2Cluster(ctx, id).Environment(environment).Execute()

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
    resp, r, err := api_client.ClustersSrcmV2Api.GetSrcmV2Cluster(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersSrcmV2Api.GetSrcmV2Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSrcmV2Cluster`: SrcmV2Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersSrcmV2Api.GetSrcmV2Cluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSrcmV2ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**SrcmV2Cluster**](srcm.v2.Cluster.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSrcmV2Clusters

> SrcmV2ClusterList ListSrcmV2Clusters(ctx).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()

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
    resp, r, err := api_client.ClustersSrcmV2Api.ListSrcmV2Clusters(context.Background()).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersSrcmV2Api.ListSrcmV2Clusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSrcmV2Clusters`: SrcmV2ClusterList
    fmt.Fprintf(os.Stdout, "Response from `ClustersSrcmV2Api.ListSrcmV2Clusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSrcmV2ClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**SrcmV2ClusterList**](srcm.v2.ClusterList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSrcmV2Cluster

> SrcmV2Cluster UpdateSrcmV2Cluster(ctx, id).SrcmV2ClusterUpdate(srcmV2ClusterUpdate).Execute()

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
    srcmV2ClusterUpdate := *openapiclient.NewSrcmV2ClusterUpdate() // SrcmV2ClusterUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersSrcmV2Api.UpdateSrcmV2Cluster(context.Background(), id).SrcmV2ClusterUpdate(srcmV2ClusterUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersSrcmV2Api.UpdateSrcmV2Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSrcmV2Cluster`: SrcmV2Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersSrcmV2Api.UpdateSrcmV2Cluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSrcmV2ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **srcmV2ClusterUpdate** | [**SrcmV2ClusterUpdate**](SrcmV2ClusterUpdate.md) |  | 

### Return type

[**SrcmV2Cluster**](srcm.v2.Cluster.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

