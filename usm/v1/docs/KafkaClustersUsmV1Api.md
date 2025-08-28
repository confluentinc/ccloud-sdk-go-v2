# \KafkaClustersUsmV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUsmV1KafkaCluster**](KafkaClustersUsmV1Api.md#CreateUsmV1KafkaCluster) | **Post** /usm/v1/kafka-clusters | Create a Kafka Cluster
[**DeleteUsmV1KafkaCluster**](KafkaClustersUsmV1Api.md#DeleteUsmV1KafkaCluster) | **Delete** /usm/v1/kafka-clusters/{id} | Delete a Kafka Cluster
[**GetUsmV1KafkaCluster**](KafkaClustersUsmV1Api.md#GetUsmV1KafkaCluster) | **Get** /usm/v1/kafka-clusters/{id} | Read a Kafka Cluster
[**ListUsmV1KafkaClusters**](KafkaClustersUsmV1Api.md#ListUsmV1KafkaClusters) | **Get** /usm/v1/kafka-clusters | List of Kafka Clusters



## CreateUsmV1KafkaCluster

> UsmV1KafkaCluster CreateUsmV1KafkaCluster(ctx).UsmV1KafkaCluster(usmV1KafkaCluster).Execute()

Create a Kafka Cluster



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
    usmV1KafkaCluster := *openapiclient.NewUsmV1KafkaCluster() // UsmV1KafkaCluster |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KafkaClustersUsmV1Api.CreateUsmV1KafkaCluster(context.Background()).UsmV1KafkaCluster(usmV1KafkaCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaClustersUsmV1Api.CreateUsmV1KafkaCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUsmV1KafkaCluster`: UsmV1KafkaCluster
    fmt.Fprintf(os.Stdout, "Response from `KafkaClustersUsmV1Api.CreateUsmV1KafkaCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUsmV1KafkaClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usmV1KafkaCluster** | [**UsmV1KafkaCluster**](UsmV1KafkaCluster.md) |  | 

### Return type

[**UsmV1KafkaCluster**](usm.v1.KafkaCluster.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUsmV1KafkaCluster

> DeleteUsmV1KafkaCluster(ctx, id).Environment(environment).Execute()

Delete a Kafka Cluster



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
    id := "id_example" // string | The unique identifier for the kafka cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KafkaClustersUsmV1Api.DeleteUsmV1KafkaCluster(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaClustersUsmV1Api.DeleteUsmV1KafkaCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the kafka cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsmV1KafkaClusterRequest struct via the builder pattern


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


## GetUsmV1KafkaCluster

> UsmV1KafkaCluster GetUsmV1KafkaCluster(ctx, id).Environment(environment).Execute()

Read a Kafka Cluster



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
    id := "id_example" // string | The unique identifier for the kafka cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KafkaClustersUsmV1Api.GetUsmV1KafkaCluster(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaClustersUsmV1Api.GetUsmV1KafkaCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsmV1KafkaCluster`: UsmV1KafkaCluster
    fmt.Fprintf(os.Stdout, "Response from `KafkaClustersUsmV1Api.GetUsmV1KafkaCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the kafka cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsmV1KafkaClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**UsmV1KafkaCluster**](usm.v1.KafkaCluster.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsmV1KafkaClusters

> UsmV1KafkaClusterList ListUsmV1KafkaClusters(ctx).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()

List of Kafka Clusters



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
    resp, r, err := api_client.KafkaClustersUsmV1Api.ListUsmV1KafkaClusters(context.Background()).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaClustersUsmV1Api.ListUsmV1KafkaClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsmV1KafkaClusters`: UsmV1KafkaClusterList
    fmt.Fprintf(os.Stdout, "Response from `KafkaClustersUsmV1Api.ListUsmV1KafkaClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsmV1KafkaClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**UsmV1KafkaClusterList**](usm.v1.KafkaClusterList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

