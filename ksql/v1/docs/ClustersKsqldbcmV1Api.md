# \ClustersKsqldbcmV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKsqldbcmV1Cluster**](ClustersKsqldbcmV1Api.md#CreateKsqldbcmV1Cluster) | **Post** /ksqldbcm/v1/clusters | Create a Cluster
[**DeleteKsqldbcmV1Cluster**](ClustersKsqldbcmV1Api.md#DeleteKsqldbcmV1Cluster) | **Delete** /ksqldbcm/v1/clusters/{id} | Delete a Cluster
[**GetKsqldbcmV1Cluster**](ClustersKsqldbcmV1Api.md#GetKsqldbcmV1Cluster) | **Get** /ksqldbcm/v1/clusters/{id} | Read a Cluster
[**ListKsqldbcmV1Clusters**](ClustersKsqldbcmV1Api.md#ListKsqldbcmV1Clusters) | **Get** /ksqldbcm/v1/clusters | List of Clusters



## CreateKsqldbcmV1Cluster

> KsqldbcmV1Cluster CreateKsqldbcmV1Cluster(ctx).KsqldbcmV1Cluster(ksqldbcmV1Cluster).Execute()

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
    ksqldbcmV1Cluster := *openapiclient.NewKsqldbcmV1Cluster() // KsqldbcmV1Cluster |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersKsqldbcmV1Api.CreateKsqldbcmV1Cluster(context.Background()).KsqldbcmV1Cluster(ksqldbcmV1Cluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersKsqldbcmV1Api.CreateKsqldbcmV1Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKsqldbcmV1Cluster`: KsqldbcmV1Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersKsqldbcmV1Api.CreateKsqldbcmV1Cluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKsqldbcmV1ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ksqldbcmV1Cluster** | [**KsqldbcmV1Cluster**](KsqldbcmV1Cluster.md) |  | 

### Return type

[**KsqldbcmV1Cluster**](ksqldbcm.v1.Cluster.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKsqldbcmV1Cluster

> DeleteKsqldbcmV1Cluster(ctx, id).Environment(environment).Execute()

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
    resp, r, err := api_client.ClustersKsqldbcmV1Api.DeleteKsqldbcmV1Cluster(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersKsqldbcmV1Api.DeleteKsqldbcmV1Cluster``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteKsqldbcmV1ClusterRequest struct via the builder pattern


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


## GetKsqldbcmV1Cluster

> KsqldbcmV1Cluster GetKsqldbcmV1Cluster(ctx, id).Environment(environment).Execute()

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
    resp, r, err := api_client.ClustersKsqldbcmV1Api.GetKsqldbcmV1Cluster(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersKsqldbcmV1Api.GetKsqldbcmV1Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKsqldbcmV1Cluster`: KsqldbcmV1Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersKsqldbcmV1Api.GetKsqldbcmV1Cluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKsqldbcmV1ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**KsqldbcmV1Cluster**](ksqldbcm.v1.Cluster.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKsqldbcmV1Clusters

> KsqldbcmV1ClusterList ListKsqldbcmV1Clusters(ctx).Environment(environment).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersKsqldbcmV1Api.ListKsqldbcmV1Clusters(context.Background()).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersKsqldbcmV1Api.ListKsqldbcmV1Clusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKsqldbcmV1Clusters`: KsqldbcmV1ClusterList
    fmt.Fprintf(os.Stdout, "Response from `ClustersKsqldbcmV1Api.ListKsqldbcmV1Clusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListKsqldbcmV1ClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 

### Return type

[**KsqldbcmV1ClusterList**](ksqldbcm.v1.ClusterList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

