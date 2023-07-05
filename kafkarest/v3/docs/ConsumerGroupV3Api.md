# \ConsumerGroupV3Api

All URIs are relative to *https://pkc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKafkaConsumer**](ConsumerGroupV3Api.md#GetKafkaConsumer) | **Get** /kafka/v3/clusters/{cluster_id}/consumer-groups/{consumer_group_id}/consumers/{consumer_id} | Get Consumer
[**GetKafkaConsumerGroup**](ConsumerGroupV3Api.md#GetKafkaConsumerGroup) | **Get** /kafka/v3/clusters/{cluster_id}/consumer-groups/{consumer_group_id} | Get Consumer Group
[**GetKafkaConsumerGroupLagSummary**](ConsumerGroupV3Api.md#GetKafkaConsumerGroupLagSummary) | **Get** /kafka/v3/clusters/{cluster_id}/consumer-groups/{consumer_group_id}/lag-summary | Get Consumer Group Lag Summary
[**GetKafkaConsumerLag**](ConsumerGroupV3Api.md#GetKafkaConsumerLag) | **Get** /kafka/v3/clusters/{cluster_id}/consumer-groups/{consumer_group_id}/lags/{topic_name}/partitions/{partition_id} | Get Consumer Lag
[**ListKafkaConsumerGroups**](ConsumerGroupV3Api.md#ListKafkaConsumerGroups) | **Get** /kafka/v3/clusters/{cluster_id}/consumer-groups | List Consumer Groups
[**ListKafkaConsumerLags**](ConsumerGroupV3Api.md#ListKafkaConsumerLags) | **Get** /kafka/v3/clusters/{cluster_id}/consumer-groups/{consumer_group_id}/lags | List Consumer Lags
[**ListKafkaConsumers**](ConsumerGroupV3Api.md#ListKafkaConsumers) | **Get** /kafka/v3/clusters/{cluster_id}/consumer-groups/{consumer_group_id}/consumers | List Consumers



## GetKafkaConsumer

> ConsumerData GetKafkaConsumer(ctx, clusterId, consumerGroupId, consumerId).Execute()

Get Consumer



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
    consumerId := "consumer-1" // string | The consumer ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConsumerGroupV3Api.GetKafkaConsumer(context.Background(), clusterId, consumerGroupId, consumerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsumerGroupV3Api.GetKafkaConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaConsumer`: ConsumerData
    fmt.Fprintf(os.Stdout, "Response from `ConsumerGroupV3Api.GetKafkaConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**consumerGroupId** | **string** | The consumer group ID. | 
**consumerId** | **string** | The consumer ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ConsumerData**](ConsumerData.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaConsumerGroup

> ConsumerGroupData GetKafkaConsumerGroup(ctx, clusterId, consumerGroupId).Execute()

Get Consumer Group



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConsumerGroupV3Api.GetKafkaConsumerGroup(context.Background(), clusterId, consumerGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsumerGroupV3Api.GetKafkaConsumerGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaConsumerGroup`: ConsumerGroupData
    fmt.Fprintf(os.Stdout, "Response from `ConsumerGroupV3Api.GetKafkaConsumerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**consumerGroupId** | **string** | The consumer group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaConsumerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConsumerGroupData**](ConsumerGroupData.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaConsumerGroupLagSummary

> ConsumerGroupLagSummaryData GetKafkaConsumerGroupLagSummary(ctx, clusterId, consumerGroupId).Execute()

Get Consumer Group Lag Summary



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConsumerGroupV3Api.GetKafkaConsumerGroupLagSummary(context.Background(), clusterId, consumerGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsumerGroupV3Api.GetKafkaConsumerGroupLagSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaConsumerGroupLagSummary`: ConsumerGroupLagSummaryData
    fmt.Fprintf(os.Stdout, "Response from `ConsumerGroupV3Api.GetKafkaConsumerGroupLagSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**consumerGroupId** | **string** | The consumer group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaConsumerGroupLagSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConsumerGroupLagSummaryData**](ConsumerGroupLagSummaryData.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaConsumerLag

> ConsumerLagData GetKafkaConsumerLag(ctx, clusterId, consumerGroupId, topicName, partitionId).Execute()

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
    resp, r, err := api_client.ConsumerGroupV3Api.GetKafkaConsumerLag(context.Background(), clusterId, consumerGroupId, topicName, partitionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsumerGroupV3Api.GetKafkaConsumerLag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaConsumerLag`: ConsumerLagData
    fmt.Fprintf(os.Stdout, "Response from `ConsumerGroupV3Api.GetKafkaConsumerLag`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetKafkaConsumerLagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ConsumerLagData**](ConsumerLagData.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaConsumerGroups

> ConsumerGroupDataList ListKafkaConsumerGroups(ctx, clusterId).Execute()

List Consumer Groups



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConsumerGroupV3Api.ListKafkaConsumerGroups(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsumerGroupV3Api.ListKafkaConsumerGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaConsumerGroups`: ConsumerGroupDataList
    fmt.Fprintf(os.Stdout, "Response from `ConsumerGroupV3Api.ListKafkaConsumerGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaConsumerGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsumerGroupDataList**](ConsumerGroupDataList.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaConsumerLags

> ConsumerLagDataList ListKafkaConsumerLags(ctx, clusterId, consumerGroupId).Execute()

List Consumer Lags



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConsumerGroupV3Api.ListKafkaConsumerLags(context.Background(), clusterId, consumerGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsumerGroupV3Api.ListKafkaConsumerLags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaConsumerLags`: ConsumerLagDataList
    fmt.Fprintf(os.Stdout, "Response from `ConsumerGroupV3Api.ListKafkaConsumerLags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**consumerGroupId** | **string** | The consumer group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaConsumerLagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConsumerLagDataList**](ConsumerLagDataList.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaConsumers

> ConsumerDataList ListKafkaConsumers(ctx, clusterId, consumerGroupId).Execute()

List Consumers



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConsumerGroupV3Api.ListKafkaConsumers(context.Background(), clusterId, consumerGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsumerGroupV3Api.ListKafkaConsumers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaConsumers`: ConsumerDataList
    fmt.Fprintf(os.Stdout, "Response from `ConsumerGroupV3Api.ListKafkaConsumers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**consumerGroupId** | **string** | The consumer group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaConsumersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConsumerDataList**](ConsumerDataList.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

