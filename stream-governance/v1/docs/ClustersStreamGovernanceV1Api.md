# \ClustersStreamGovernanceV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStreamGovernanceV1Cluster**](ClustersStreamGovernanceV1Api.md#CreateStreamGovernanceV1Cluster) | **Post** /stream-governance/v1/clusters | Create a Cluster
[**DeleteStreamGovernanceV1Cluster**](ClustersStreamGovernanceV1Api.md#DeleteStreamGovernanceV1Cluster) | **Delete** /stream-governance/v1/clusters/{id} | Delete a Cluster
[**GetStreamGovernanceV1Cluster**](ClustersStreamGovernanceV1Api.md#GetStreamGovernanceV1Cluster) | **Get** /stream-governance/v1/clusters/{id} | Read a Cluster
[**ListStreamGovernanceV1Clusters**](ClustersStreamGovernanceV1Api.md#ListStreamGovernanceV1Clusters) | **Get** /stream-governance/v1/clusters | List of Clusters
[**UpdateStreamGovernanceV1Cluster**](ClustersStreamGovernanceV1Api.md#UpdateStreamGovernanceV1Cluster) | **Patch** /stream-governance/v1/clusters/{id} | Update a Cluster



## CreateStreamGovernanceV1Cluster

> StreamGovernanceV1Cluster CreateStreamGovernanceV1Cluster(ctx).StreamGovernanceV1Cluster(streamGovernanceV1Cluster).Execute()

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
    streamGovernanceV1Cluster := *openapiclient.NewStreamGovernanceV1Cluster() // StreamGovernanceV1Cluster |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersStreamGovernanceV1Api.CreateStreamGovernanceV1Cluster(context.Background()).StreamGovernanceV1Cluster(streamGovernanceV1Cluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersStreamGovernanceV1Api.CreateStreamGovernanceV1Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStreamGovernanceV1Cluster`: StreamGovernanceV1Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersStreamGovernanceV1Api.CreateStreamGovernanceV1Cluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStreamGovernanceV1ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **streamGovernanceV1Cluster** | [**StreamGovernanceV1Cluster**](StreamGovernanceV1Cluster.md) |  | 

### Return type

[**StreamGovernanceV1Cluster**](stream-governance.v1.Cluster.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStreamGovernanceV1Cluster

> DeleteStreamGovernanceV1Cluster(ctx, id).Execute()

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
    id := "id_example" // string | The unique identifier for the cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersStreamGovernanceV1Api.DeleteStreamGovernanceV1Cluster(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersStreamGovernanceV1Api.DeleteStreamGovernanceV1Cluster``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteStreamGovernanceV1ClusterRequest struct via the builder pattern


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


## GetStreamGovernanceV1Cluster

> StreamGovernanceV1Cluster GetStreamGovernanceV1Cluster(ctx, id).Execute()

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
    id := "id_example" // string | The unique identifier for the cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersStreamGovernanceV1Api.GetStreamGovernanceV1Cluster(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersStreamGovernanceV1Api.GetStreamGovernanceV1Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStreamGovernanceV1Cluster`: StreamGovernanceV1Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersStreamGovernanceV1Api.GetStreamGovernanceV1Cluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamGovernanceV1ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StreamGovernanceV1Cluster**](stream-governance.v1.Cluster.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStreamGovernanceV1Clusters

> StreamGovernanceV1ClusterList ListStreamGovernanceV1Clusters(ctx).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()

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
    environment := "env-00000" // string | Filter the results by exact match for environment. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersStreamGovernanceV1Api.ListStreamGovernanceV1Clusters(context.Background()).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersStreamGovernanceV1Api.ListStreamGovernanceV1Clusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStreamGovernanceV1Clusters`: StreamGovernanceV1ClusterList
    fmt.Fprintf(os.Stdout, "Response from `ClustersStreamGovernanceV1Api.ListStreamGovernanceV1Clusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListStreamGovernanceV1ClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**StreamGovernanceV1ClusterList**](stream-governance.v1.ClusterList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStreamGovernanceV1Cluster

> StreamGovernanceV1Cluster UpdateStreamGovernanceV1Cluster(ctx, id).StreamGovernanceV1ClusterUpdate(streamGovernanceV1ClusterUpdate).Execute()

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
    streamGovernanceV1ClusterUpdate := *openapiclient.NewStreamGovernanceV1ClusterUpdate() // StreamGovernanceV1ClusterUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersStreamGovernanceV1Api.UpdateStreamGovernanceV1Cluster(context.Background(), id).StreamGovernanceV1ClusterUpdate(streamGovernanceV1ClusterUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersStreamGovernanceV1Api.UpdateStreamGovernanceV1Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStreamGovernanceV1Cluster`: StreamGovernanceV1Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersStreamGovernanceV1Api.UpdateStreamGovernanceV1Cluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStreamGovernanceV1ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **streamGovernanceV1ClusterUpdate** | [**StreamGovernanceV1ClusterUpdate**](StreamGovernanceV1ClusterUpdate.md) |  | 

### Return type

[**StreamGovernanceV1Cluster**](stream-governance.v1.Cluster.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

