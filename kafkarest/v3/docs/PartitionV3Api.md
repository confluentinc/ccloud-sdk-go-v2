# \PartitionV3Api

All URIs are relative to *https://pkc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKafkaV3ConsumerLag**](PartitionV3Api.md#GetKafkaV3ConsumerLag) | **Get** /kafka/v3/clusters/{cluster_id}/consumer-groups/{consumer_group_id}/lags/{topic_name}/partitions/{partition_id} | Get Consumer Lag
[**GetKafkaV3Partition**](PartitionV3Api.md#GetKafkaV3Partition) | **Get** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/partitions/{partition_id} | Get Partition
[**ListKafkaV3Partitions**](PartitionV3Api.md#ListKafkaV3Partitions) | **Get** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/partitions | List Partitions



## GetKafkaV3ConsumerLag

> ConsumerLagData GetKafkaV3ConsumerLag(ctx, clusterId, consumerGroupId, topicName, partitionId).Execute()

Get Consumer Lag



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
    clusterId := "cluster-1" // string | The Kafka cluster ID.
    consumerGroupId := "consumer-group-1" // string | The consumer group ID.
    topicName := "topic-1" // string | The topic name.
    partitionId := int32(0) // int32 | The partition ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartitionV3Api.GetKafkaV3ConsumerLag(context.Background(), clusterId, consumerGroupId, topicName, partitionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartitionV3Api.GetKafkaV3ConsumerLag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaV3ConsumerLag`: ConsumerLagData
    fmt.Fprintf(os.Stdout, "Response from `PartitionV3Api.GetKafkaV3ConsumerLag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**consumerGroupId** | **string** | The consumer group ID. | 
**topicName** | **string** | The topic name. | 
**partitionId** | **int32** | The partition ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaV3ConsumerLagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ConsumerLagData**](ConsumerLagData.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaV3Partition

> PartitionData GetKafkaV3Partition(ctx, clusterId, topicName, partitionId).Execute()

Get Partition



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
    clusterId := "cluster-1" // string | The Kafka cluster ID.
    topicName := "topic-1" // string | The topic name.
    partitionId := int32(0) // int32 | The partition ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartitionV3Api.GetKafkaV3Partition(context.Background(), clusterId, topicName, partitionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartitionV3Api.GetKafkaV3Partition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaV3Partition`: PartitionData
    fmt.Fprintf(os.Stdout, "Response from `PartitionV3Api.GetKafkaV3Partition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**topicName** | **string** | The topic name. | 
**partitionId** | **int32** | The partition ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaV3PartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PartitionData**](PartitionData.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaV3Partitions

> PartitionDataList ListKafkaV3Partitions(ctx, clusterId, topicName).Execute()

List Partitions



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
    clusterId := "cluster-1" // string | The Kafka cluster ID.
    topicName := "topic-1" // string | The topic name.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartitionV3Api.ListKafkaV3Partitions(context.Background(), clusterId, topicName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartitionV3Api.ListKafkaV3Partitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaV3Partitions`: PartitionDataList
    fmt.Fprintf(os.Stdout, "Response from `PartitionV3Api.ListKafkaV3Partitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**topicName** | **string** | The topic name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaV3PartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PartitionDataList**](PartitionDataList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

