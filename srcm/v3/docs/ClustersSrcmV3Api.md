# \ClustersSrcmV3Api

All URIs below are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSrcmV3Cluster**](ClustersSrcmV3Api.md#GetSrcmV3Cluster) | **Get** /srcm/v3/clusters/{id} | Read a Cluster
[**ListSrcmV3Clusters**](ClustersSrcmV3Api.md#ListSrcmV3Clusters) | **Get** /srcm/v3/clusters | List of Clusters



## GetSrcmV3Cluster

> SrcmV3Cluster GetSrcmV3Cluster(ctx, id).Environment(environment).Execute()

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
    resp, r, err := api_client.ClustersSrcmV3Api.GetSrcmV3Cluster(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersSrcmV3Api.GetSrcmV3Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSrcmV3Cluster`: SrcmV3Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersSrcmV3Api.GetSrcmV3Cluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSrcmV3ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**SrcmV3Cluster**](srcm.v3.Cluster.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSrcmV3Clusters

> SrcmV3ClusterList ListSrcmV3Clusters(ctx).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()

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
    resp, r, err := api_client.ClustersSrcmV3Api.ListSrcmV3Clusters(context.Background()).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersSrcmV3Api.ListSrcmV3Clusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSrcmV3Clusters`: SrcmV3ClusterList
    fmt.Fprintf(os.Stdout, "Response from `ClustersSrcmV3Api.ListSrcmV3Clusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSrcmV3ClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**SrcmV3ClusterList**](srcm.v3.ClusterList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

