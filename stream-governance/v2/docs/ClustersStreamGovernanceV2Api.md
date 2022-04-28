# \ClustersStreamGovernanceV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStreamGovernanceV2Cluster**](ClustersStreamGovernanceV2Api.md#CreateStreamGovernanceV2Cluster) | **Post** /stream-governance/v2/clusters | Create a Cluster
[**DeleteStreamGovernanceV2Cluster**](ClustersStreamGovernanceV2Api.md#DeleteStreamGovernanceV2Cluster) | **Delete** /stream-governance/v2/clusters/{id} | Delete a Cluster
[**GetStreamGovernanceV2Cluster**](ClustersStreamGovernanceV2Api.md#GetStreamGovernanceV2Cluster) | **Get** /stream-governance/v2/clusters/{id} | Read a Cluster
[**ListStreamGovernanceV2Clusters**](ClustersStreamGovernanceV2Api.md#ListStreamGovernanceV2Clusters) | **Get** /stream-governance/v2/clusters | List of Clusters
[**UpdateStreamGovernanceV2Cluster**](ClustersStreamGovernanceV2Api.md#UpdateStreamGovernanceV2Cluster) | **Patch** /stream-governance/v2/clusters/{id} | Update a Cluster



## CreateStreamGovernanceV2Cluster

> StreamGovernanceV2Cluster CreateStreamGovernanceV2Cluster(ctx).StreamGovernanceV2Cluster(streamGovernanceV2Cluster).Execute()

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
    streamGovernanceV2Cluster := *openapiclient.NewStreamGovernanceV2Cluster() // StreamGovernanceV2Cluster |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersStreamGovernanceV2Api.CreateStreamGovernanceV2Cluster(context.Background()).StreamGovernanceV2Cluster(streamGovernanceV2Cluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersStreamGovernanceV2Api.CreateStreamGovernanceV2Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStreamGovernanceV2Cluster`: StreamGovernanceV2Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersStreamGovernanceV2Api.CreateStreamGovernanceV2Cluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStreamGovernanceV2ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **streamGovernanceV2Cluster** | [**StreamGovernanceV2Cluster**](StreamGovernanceV2Cluster.md) |  | 

### Return type

[**StreamGovernanceV2Cluster**](stream-governance.v2.Cluster.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStreamGovernanceV2Cluster

> DeleteStreamGovernanceV2Cluster(ctx, id).Execute()

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
    resp, r, err := api_client.ClustersStreamGovernanceV2Api.DeleteStreamGovernanceV2Cluster(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersStreamGovernanceV2Api.DeleteStreamGovernanceV2Cluster``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteStreamGovernanceV2ClusterRequest struct via the builder pattern


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


## GetStreamGovernanceV2Cluster

> StreamGovernanceV2Cluster GetStreamGovernanceV2Cluster(ctx, id).Execute()

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
    resp, r, err := api_client.ClustersStreamGovernanceV2Api.GetStreamGovernanceV2Cluster(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersStreamGovernanceV2Api.GetStreamGovernanceV2Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStreamGovernanceV2Cluster`: StreamGovernanceV2Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersStreamGovernanceV2Api.GetStreamGovernanceV2Cluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamGovernanceV2ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StreamGovernanceV2Cluster**](stream-governance.v2.Cluster.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStreamGovernanceV2Clusters

> StreamGovernanceV2ClusterList ListStreamGovernanceV2Clusters(ctx).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()

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
    resp, r, err := api_client.ClustersStreamGovernanceV2Api.ListStreamGovernanceV2Clusters(context.Background()).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersStreamGovernanceV2Api.ListStreamGovernanceV2Clusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStreamGovernanceV2Clusters`: StreamGovernanceV2ClusterList
    fmt.Fprintf(os.Stdout, "Response from `ClustersStreamGovernanceV2Api.ListStreamGovernanceV2Clusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListStreamGovernanceV2ClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**StreamGovernanceV2ClusterList**](stream-governance.v2.ClusterList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStreamGovernanceV2Cluster

> StreamGovernanceV2Cluster UpdateStreamGovernanceV2Cluster(ctx, id).StreamGovernanceV2ClusterUpdate(streamGovernanceV2ClusterUpdate).Execute()

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
    streamGovernanceV2ClusterUpdate := *openapiclient.NewStreamGovernanceV2ClusterUpdate() // StreamGovernanceV2ClusterUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersStreamGovernanceV2Api.UpdateStreamGovernanceV2Cluster(context.Background(), id).StreamGovernanceV2ClusterUpdate(streamGovernanceV2ClusterUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersStreamGovernanceV2Api.UpdateStreamGovernanceV2Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStreamGovernanceV2Cluster`: StreamGovernanceV2Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersStreamGovernanceV2Api.UpdateStreamGovernanceV2Cluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStreamGovernanceV2ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **streamGovernanceV2ClusterUpdate** | [**StreamGovernanceV2ClusterUpdate**](StreamGovernanceV2ClusterUpdate.md) |  | 

### Return type

[**StreamGovernanceV2Cluster**](stream-governance.v2.Cluster.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

