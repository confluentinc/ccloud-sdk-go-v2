# \KafkaClustersCmkV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCmkV2KafkaCluster**](KafkaClustersCmkV2Api.md#CreateCmkV2KafkaCluster) | **Post** /cmk/v2/clusters | Create a Kafka Cluster
[**DeleteCmkV2KafkaCluster**](KafkaClustersCmkV2Api.md#DeleteCmkV2KafkaCluster) | **Delete** /cmk/v2/clusters/{id} | Delete a Kafka Cluster
[**GetCmkV2KafkaCluster**](KafkaClustersCmkV2Api.md#GetCmkV2KafkaCluster) | **Get** /cmk/v2/clusters/{id} | Read a Kafka Cluster
[**ListCmkV2KafkaClusters**](KafkaClustersCmkV2Api.md#ListCmkV2KafkaClusters) | **Get** /cmk/v2/clusters | List of Kafka Clusters
[**UpdateCmkV2KafkaCluster**](KafkaClustersCmkV2Api.md#UpdateCmkV2KafkaCluster) | **Patch** /cmk/v2/clusters/{id} | Update a Kafka Cluster



## CreateCmkV2KafkaCluster

> CmkV2KafkaCluster CreateCmkV2KafkaCluster(ctx).CmkV2KafkaCluster(cmkV2KafkaCluster).Execute()

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
    cmkV2KafkaCluster := *openapiclient.NewCmkV2KafkaCluster() // CmkV2KafkaCluster |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KafkaClustersCmkV2Api.CreateCmkV2KafkaCluster(context.Background()).CmkV2KafkaCluster(cmkV2KafkaCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaClustersCmkV2Api.CreateCmkV2KafkaCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCmkV2KafkaCluster`: CmkV2KafkaCluster
    fmt.Fprintf(os.Stdout, "Response from `KafkaClustersCmkV2Api.CreateCmkV2KafkaCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCmkV2KafkaClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cmkV2KafkaCluster** | [**CmkV2KafkaCluster**](CmkV2KafkaCluster.md) |  | 

### Return type

[**CmkV2KafkaCluster**](cmk.v2.KafkaCluster.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCmkV2KafkaCluster

> DeleteCmkV2KafkaCluster(ctx, id).Environment(environment).Execute()

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
    resp, r, err := api_client.KafkaClustersCmkV2Api.DeleteCmkV2KafkaCluster(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaClustersCmkV2Api.DeleteCmkV2KafkaCluster``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCmkV2KafkaClusterRequest struct via the builder pattern


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


## GetCmkV2KafkaCluster

> CmkV2KafkaCluster GetCmkV2KafkaCluster(ctx, id).Environment(environment).Execute()

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
    resp, r, err := api_client.KafkaClustersCmkV2Api.GetCmkV2KafkaCluster(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaClustersCmkV2Api.GetCmkV2KafkaCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmkV2KafkaCluster`: CmkV2KafkaCluster
    fmt.Fprintf(os.Stdout, "Response from `KafkaClustersCmkV2Api.GetCmkV2KafkaCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the kafka cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmkV2KafkaClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**CmkV2KafkaCluster**](cmk.v2.KafkaCluster.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCmkV2KafkaClusters

> CmkV2KafkaClusterList ListCmkV2KafkaClusters(ctx).Environment(environment).Phase(phase).PageSize(pageSize).PageToken(pageToken).Execute()

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
    phase := "PROVISIONED" // string | Filter the results by exact match for phase. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KafkaClustersCmkV2Api.ListCmkV2KafkaClusters(context.Background()).Environment(environment).Phase(phase).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaClustersCmkV2Api.ListCmkV2KafkaClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCmkV2KafkaClusters`: CmkV2KafkaClusterList
    fmt.Fprintf(os.Stdout, "Response from `KafkaClustersCmkV2Api.ListCmkV2KafkaClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCmkV2KafkaClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **phase** | **string** | Filter the results by exact match for phase. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**CmkV2KafkaClusterList**](CmkV2KafkaClusterList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCmkV2KafkaCluster

> CmkV2KafkaCluster UpdateCmkV2KafkaCluster(ctx, id).CmkV2KafkaCluster(cmkV2KafkaCluster).Execute()

Update a Kafka Cluster



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
    id := "id_example" // string | The unique identifier for the kafka cluster.
    cmkV2KafkaCluster := *openapiclient.NewCmkV2KafkaCluster() // CmkV2KafkaCluster |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KafkaClustersCmkV2Api.UpdateCmkV2KafkaCluster(context.Background(), id).CmkV2KafkaCluster(cmkV2KafkaCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaClustersCmkV2Api.UpdateCmkV2KafkaCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCmkV2KafkaCluster`: CmkV2KafkaCluster
    fmt.Fprintf(os.Stdout, "Response from `KafkaClustersCmkV2Api.UpdateCmkV2KafkaCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the kafka cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCmkV2KafkaClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cmkV2KafkaCluster** | [**CmkV2KafkaCluster**](CmkV2KafkaCluster.md) |  | 

### Return type

[**CmkV2KafkaCluster**](cmk.v2.KafkaCluster.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

