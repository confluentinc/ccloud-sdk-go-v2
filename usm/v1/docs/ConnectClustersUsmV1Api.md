# \ConnectClustersUsmV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUsmV1ConnectCluster**](ConnectClustersUsmV1Api.md#CreateUsmV1ConnectCluster) | **Post** /usm/v1/connect-clusters | Create a Connect Cluster
[**DeleteUsmV1ConnectCluster**](ConnectClustersUsmV1Api.md#DeleteUsmV1ConnectCluster) | **Delete** /usm/v1/connect-clusters/{id} | Delete a Connect Cluster
[**GetUsmV1ConnectCluster**](ConnectClustersUsmV1Api.md#GetUsmV1ConnectCluster) | **Get** /usm/v1/connect-clusters/{id} | Read a Connect Cluster
[**ListUsmV1ConnectClusters**](ConnectClustersUsmV1Api.md#ListUsmV1ConnectClusters) | **Get** /usm/v1/connect-clusters | List of Connect Clusters



## CreateUsmV1ConnectCluster

> UsmV1ConnectCluster CreateUsmV1ConnectCluster(ctx).UsmV1ConnectCluster(usmV1ConnectCluster).Execute()

Create a Connect Cluster



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
    usmV1ConnectCluster := *openapiclient.NewUsmV1ConnectCluster() // UsmV1ConnectCluster |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectClustersUsmV1Api.CreateUsmV1ConnectCluster(context.Background()).UsmV1ConnectCluster(usmV1ConnectCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectClustersUsmV1Api.CreateUsmV1ConnectCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUsmV1ConnectCluster`: UsmV1ConnectCluster
    fmt.Fprintf(os.Stdout, "Response from `ConnectClustersUsmV1Api.CreateUsmV1ConnectCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUsmV1ConnectClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usmV1ConnectCluster** | [**UsmV1ConnectCluster**](UsmV1ConnectCluster.md) |  | 

### Return type

[**UsmV1ConnectCluster**](usm.v1.ConnectCluster.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUsmV1ConnectCluster

> DeleteUsmV1ConnectCluster(ctx, id).Environment(environment).Execute()

Delete a Connect Cluster



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
    id := "id_example" // string | The unique identifier for the connect cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectClustersUsmV1Api.DeleteUsmV1ConnectCluster(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectClustersUsmV1Api.DeleteUsmV1ConnectCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the connect cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsmV1ConnectClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

 (empty response body)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsmV1ConnectCluster

> UsmV1ConnectCluster GetUsmV1ConnectCluster(ctx, id).Environment(environment).Execute()

Read a Connect Cluster



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
    id := "id_example" // string | The unique identifier for the connect cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectClustersUsmV1Api.GetUsmV1ConnectCluster(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectClustersUsmV1Api.GetUsmV1ConnectCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsmV1ConnectCluster`: UsmV1ConnectCluster
    fmt.Fprintf(os.Stdout, "Response from `ConnectClustersUsmV1Api.GetUsmV1ConnectCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the connect cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsmV1ConnectClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**UsmV1ConnectCluster**](usm.v1.ConnectCluster.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsmV1ConnectClusters

> UsmV1ConnectClusterList ListUsmV1ConnectClusters(ctx).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()

List of Connect Clusters



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
    resp, r, err := api_client.ConnectClustersUsmV1Api.ListUsmV1ConnectClusters(context.Background()).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectClustersUsmV1Api.ListUsmV1ConnectClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsmV1ConnectClusters`: UsmV1ConnectClusterList
    fmt.Fprintf(os.Stdout, "Response from `ConnectClustersUsmV1Api.ListUsmV1ConnectClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsmV1ConnectClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**UsmV1ConnectClusterList**](usm.v1.ConnectClusterList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

