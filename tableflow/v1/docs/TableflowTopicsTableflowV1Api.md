# \TableflowTopicsTableflowV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTableflowV1TableflowTopic**](TableflowTopicsTableflowV1Api.md#CreateTableflowV1TableflowTopic) | **Post** /tableflow/v1/tableflow-topics | Create a Tableflow Topic
[**DeleteTableflowV1TableflowTopic**](TableflowTopicsTableflowV1Api.md#DeleteTableflowV1TableflowTopic) | **Delete** /tableflow/v1/tableflow-topics/{display_name} | Delete a Tableflow Topic
[**GetTableflowV1TableflowTopic**](TableflowTopicsTableflowV1Api.md#GetTableflowV1TableflowTopic) | **Get** /tableflow/v1/tableflow-topics/{display_name} | Read a Tableflow Topic
[**ListTableflowV1TableflowTopics**](TableflowTopicsTableflowV1Api.md#ListTableflowV1TableflowTopics) | **Get** /tableflow/v1/tableflow-topics | List of Tableflow Topics
[**UpdateTableflowV1TableflowTopic**](TableflowTopicsTableflowV1Api.md#UpdateTableflowV1TableflowTopic) | **Patch** /tableflow/v1/tableflow-topics/{display_name} | Update a Tableflow Topic



## CreateTableflowV1TableflowTopic

> TableflowV1TableflowTopic CreateTableflowV1TableflowTopic(ctx).TableflowV1TableflowTopic(tableflowV1TableflowTopic).Execute()

Create a Tableflow Topic



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
    tableflowV1TableflowTopic := *openapiclient.NewTableflowV1TableflowTopic() // TableflowV1TableflowTopic |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TableflowTopicsTableflowV1Api.CreateTableflowV1TableflowTopic(context.Background()).TableflowV1TableflowTopic(tableflowV1TableflowTopic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableflowTopicsTableflowV1Api.CreateTableflowV1TableflowTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTableflowV1TableflowTopic`: TableflowV1TableflowTopic
    fmt.Fprintf(os.Stdout, "Response from `TableflowTopicsTableflowV1Api.CreateTableflowV1TableflowTopic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTableflowV1TableflowTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tableflowV1TableflowTopic** | [**TableflowV1TableflowTopic**](TableflowV1TableflowTopic.md) |  | 

### Return type

[**TableflowV1TableflowTopic**](tableflow.v1.TableflowTopic.md)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTableflowV1TableflowTopic

> DeleteTableflowV1TableflowTopic(ctx, displayName).Environment(environment).SpecKafkaCluster(specKafkaCluster).Execute()

Delete a Tableflow Topic



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
    specKafkaCluster := "lkc-00000" // string | Scope the operation to the given spec.kafka_cluster.
    displayName := "displayName_example" // string | The name of the Kafka topic for which Tableflow is enabled.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TableflowTopicsTableflowV1Api.DeleteTableflowV1TableflowTopic(context.Background(), displayName).Environment(environment).SpecKafkaCluster(specKafkaCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableflowTopicsTableflowV1Api.DeleteTableflowV1TableflowTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**displayName** | **string** | The name of the Kafka topic for which Tableflow is enabled. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTableflowV1TableflowTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 
 **specKafkaCluster** | **string** | Scope the operation to the given spec.kafka_cluster. | 


### Return type

 (empty response body)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTableflowV1TableflowTopic

> TableflowV1TableflowTopic GetTableflowV1TableflowTopic(ctx, displayName).Environment(environment).SpecKafkaCluster(specKafkaCluster).Execute()

Read a Tableflow Topic



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
    specKafkaCluster := "lkc-00000" // string | Scope the operation to the given spec.kafka_cluster.
    displayName := "displayName_example" // string | The name of the Kafka topic for which Tableflow is enabled.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TableflowTopicsTableflowV1Api.GetTableflowV1TableflowTopic(context.Background(), displayName).Environment(environment).SpecKafkaCluster(specKafkaCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableflowTopicsTableflowV1Api.GetTableflowV1TableflowTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTableflowV1TableflowTopic`: TableflowV1TableflowTopic
    fmt.Fprintf(os.Stdout, "Response from `TableflowTopicsTableflowV1Api.GetTableflowV1TableflowTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**displayName** | **string** | The name of the Kafka topic for which Tableflow is enabled. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTableflowV1TableflowTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 
 **specKafkaCluster** | **string** | Scope the operation to the given spec.kafka_cluster. | 


### Return type

[**TableflowV1TableflowTopic**](tableflow.v1.TableflowTopic.md)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTableflowV1TableflowTopics

> TableflowV1TableflowTopicList ListTableflowV1TableflowTopics(ctx).Environment(environment).SpecKafkaCluster(specKafkaCluster).PageSize(pageSize).PageToken(pageToken).Execute()

List of Tableflow Topics



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
    specKafkaCluster := "lkc-00000" // string | Filter the results by exact match for spec.kafka_cluster.
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TableflowTopicsTableflowV1Api.ListTableflowV1TableflowTopics(context.Background()).Environment(environment).SpecKafkaCluster(specKafkaCluster).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableflowTopicsTableflowV1Api.ListTableflowV1TableflowTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTableflowV1TableflowTopics`: TableflowV1TableflowTopicList
    fmt.Fprintf(os.Stdout, "Response from `TableflowTopicsTableflowV1Api.ListTableflowV1TableflowTopics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTableflowV1TableflowTopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **specKafkaCluster** | **string** | Filter the results by exact match for spec.kafka_cluster. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**TableflowV1TableflowTopicList**](tableflow.v1.TableflowTopicList.md)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTableflowV1TableflowTopic

> TableflowV1TableflowTopic UpdateTableflowV1TableflowTopic(ctx, displayName).TableflowV1TableflowTopicUpdate(tableflowV1TableflowTopicUpdate).Execute()

Update a Tableflow Topic



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
    displayName := "displayName_example" // string | The name of the Kafka topic for which Tableflow is enabled.
    tableflowV1TableflowTopicUpdate := *openapiclient.NewTableflowV1TableflowTopicUpdate() // TableflowV1TableflowTopicUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TableflowTopicsTableflowV1Api.UpdateTableflowV1TableflowTopic(context.Background(), displayName).TableflowV1TableflowTopicUpdate(tableflowV1TableflowTopicUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableflowTopicsTableflowV1Api.UpdateTableflowV1TableflowTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTableflowV1TableflowTopic`: TableflowV1TableflowTopic
    fmt.Fprintf(os.Stdout, "Response from `TableflowTopicsTableflowV1Api.UpdateTableflowV1TableflowTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**displayName** | **string** | The name of the Kafka topic for which Tableflow is enabled. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTableflowV1TableflowTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tableflowV1TableflowTopicUpdate** | [**TableflowV1TableflowTopicUpdate**](TableflowV1TableflowTopicUpdate.md) |  | 

### Return type

[**TableflowV1TableflowTopic**](tableflow.v1.TableflowTopic.md)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

