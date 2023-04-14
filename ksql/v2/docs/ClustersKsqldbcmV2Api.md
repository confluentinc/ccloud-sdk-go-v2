# \ClustersKsqldbcmV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKsqldbcmV2Cluster**](ClustersKsqldbcmV2Api.md#CreateKsqldbcmV2Cluster) | **Post** /ksqldbcm/v2/clusters | Create a Cluster
[**DeleteKsqldbcmV2Cluster**](ClustersKsqldbcmV2Api.md#DeleteKsqldbcmV2Cluster) | **Delete** /ksqldbcm/v2/clusters/{id} | Delete a Cluster
[**GetKsqldbcmV2Cluster**](ClustersKsqldbcmV2Api.md#GetKsqldbcmV2Cluster) | **Get** /ksqldbcm/v2/clusters/{id} | Read a Cluster
[**ListKsqldbcmV2Clusters**](ClustersKsqldbcmV2Api.md#ListKsqldbcmV2Clusters) | **Get** /ksqldbcm/v2/clusters | List of Clusters



## CreateKsqldbcmV2Cluster

> KsqldbcmV2Cluster CreateKsqldbcmV2Cluster(ctx).KsqldbcmV2Cluster(ksqldbcmV2Cluster).Execute()

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
    ksqldbcmV2Cluster := *openapiclient.NewKsqldbcmV2Cluster() // KsqldbcmV2Cluster |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersKsqldbcmV2Api.CreateKsqldbcmV2Cluster(context.Background()).KsqldbcmV2Cluster(ksqldbcmV2Cluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersKsqldbcmV2Api.CreateKsqldbcmV2Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKsqldbcmV2Cluster`: KsqldbcmV2Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersKsqldbcmV2Api.CreateKsqldbcmV2Cluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKsqldbcmV2ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ksqldbcmV2Cluster** | [**KsqldbcmV2Cluster**](KsqldbcmV2Cluster.md) |  | 

### Return type

[**KsqldbcmV2Cluster**](ksqldbcm.v2.Cluster.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKsqldbcmV2Cluster

> DeleteKsqldbcmV2Cluster(ctx, id).Environment(environment).Execute()

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
    resp, r, err := api_client.ClustersKsqldbcmV2Api.DeleteKsqldbcmV2Cluster(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersKsqldbcmV2Api.DeleteKsqldbcmV2Cluster``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteKsqldbcmV2ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKsqldbcmV2Cluster

> KsqldbcmV2Cluster GetKsqldbcmV2Cluster(ctx, id).Environment(environment).Execute()

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
    resp, r, err := api_client.ClustersKsqldbcmV2Api.GetKsqldbcmV2Cluster(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersKsqldbcmV2Api.GetKsqldbcmV2Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKsqldbcmV2Cluster`: KsqldbcmV2Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersKsqldbcmV2Api.GetKsqldbcmV2Cluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKsqldbcmV2ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**KsqldbcmV2Cluster**](ksqldbcm.v2.Cluster.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKsqldbcmV2Clusters

> KsqldbcmV2ClusterList ListKsqldbcmV2Clusters(ctx).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()

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
    resp, r, err := api_client.ClustersKsqldbcmV2Api.ListKsqldbcmV2Clusters(context.Background()).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersKsqldbcmV2Api.ListKsqldbcmV2Clusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKsqldbcmV2Clusters`: KsqldbcmV2ClusterList
    fmt.Fprintf(os.Stdout, "Response from `ClustersKsqldbcmV2Api.ListKsqldbcmV2Clusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListKsqldbcmV2ClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**KsqldbcmV2ClusterList**](ksqldbcm.v2.ClusterList.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

