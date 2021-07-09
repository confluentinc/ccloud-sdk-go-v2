# \ClustersCmkV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCmkV2Cluster**](ClustersCmkV2Api.md#CreateCmkV2Cluster) | **Post** /cmk/v2/clusters | Create a Cluster
[**DeleteCmkV2Cluster**](ClustersCmkV2Api.md#DeleteCmkV2Cluster) | **Delete** /cmk/v2/clusters/{id} | Delete a Cluster
[**GetCmkV2Cluster**](ClustersCmkV2Api.md#GetCmkV2Cluster) | **Get** /cmk/v2/clusters/{id} | Read a Cluster
[**ListCmkV2Clusters**](ClustersCmkV2Api.md#ListCmkV2Clusters) | **Get** /cmk/v2/clusters | List of Clusters
[**UpdateCmkV2Cluster**](ClustersCmkV2Api.md#UpdateCmkV2Cluster) | **Patch** /cmk/v2/clusters/{id} | Update a Cluster



## CreateCmkV2Cluster

> CmkV2Cluster CreateCmkV2Cluster(ctx).CmkV2Cluster(cmkV2Cluster).Execute()

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
    cmkV2Cluster := *openapiclient.NewCmkV2Cluster() // CmkV2Cluster |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersCmkV2Api.CreateCmkV2Cluster(context.Background()).CmkV2Cluster(cmkV2Cluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersCmkV2Api.CreateCmkV2Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCmkV2Cluster`: CmkV2Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersCmkV2Api.CreateCmkV2Cluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCmkV2ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cmkV2Cluster** | [**CmkV2Cluster**](CmkV2Cluster.md) |  | 

### Return type

[**CmkV2Cluster**](cmk.v2.Cluster.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCmkV2Cluster

> DeleteCmkV2Cluster(ctx, id).Environment(environment).Execute()

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
    resp, r, err := api_client.ClustersCmkV2Api.DeleteCmkV2Cluster(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersCmkV2Api.DeleteCmkV2Cluster``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCmkV2ClusterRequest struct via the builder pattern


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


## GetCmkV2Cluster

> CmkV2Cluster GetCmkV2Cluster(ctx, id).Environment(environment).Execute()

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
    resp, r, err := api_client.ClustersCmkV2Api.GetCmkV2Cluster(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersCmkV2Api.GetCmkV2Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmkV2Cluster`: CmkV2Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersCmkV2Api.GetCmkV2Cluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmkV2ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**CmkV2Cluster**](cmk.v2.Cluster.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCmkV2Clusters

> CmkV2ClusterList ListCmkV2Clusters(ctx).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()

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
    resp, r, err := api_client.ClustersCmkV2Api.ListCmkV2Clusters(context.Background()).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersCmkV2Api.ListCmkV2Clusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCmkV2Clusters`: CmkV2ClusterList
    fmt.Fprintf(os.Stdout, "Response from `ClustersCmkV2Api.ListCmkV2Clusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCmkV2ClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**CmkV2ClusterList**](CmkV2ClusterList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCmkV2Cluster

> CmkV2Cluster UpdateCmkV2Cluster(ctx, id).CmkV2ClusterUpdate(cmkV2ClusterUpdate).Execute()

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
    cmkV2ClusterUpdate := *openapiclient.NewCmkV2ClusterUpdate() // CmkV2ClusterUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersCmkV2Api.UpdateCmkV2Cluster(context.Background(), id).CmkV2ClusterUpdate(cmkV2ClusterUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersCmkV2Api.UpdateCmkV2Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCmkV2Cluster`: CmkV2Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersCmkV2Api.UpdateCmkV2Cluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCmkV2ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cmkV2ClusterUpdate** | [**CmkV2ClusterUpdate**](CmkV2ClusterUpdate.md) |  | 

### Return type

[**CmkV2Cluster**](cmk.v2.Cluster.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

