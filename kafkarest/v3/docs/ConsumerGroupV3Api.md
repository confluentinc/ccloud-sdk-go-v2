# \ConsumerGroupV3Api

All URIs are relative to *https://pkc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKafkaV3Consumer**](ConsumerGroupV3Api.md#GetKafkaV3Consumer) | **Get** /kafka/v3/clusters/{cluster_id}/consumer-groups/{consumer_group_id}/consumers/{consumer_id} | Get Consumer
[**GetKafkaV3ConsumerGroup**](ConsumerGroupV3Api.md#GetKafkaV3ConsumerGroup) | **Get** /kafka/v3/clusters/{cluster_id}/consumer-groups/{consumer_group_id} | Get Consumer Group
[**GetKafkaV3ConsumerGroupLagSummary**](ConsumerGroupV3Api.md#GetKafkaV3ConsumerGroupLagSummary) | **Get** /kafka/v3/clusters/{cluster_id}/consumer-groups/{consumer_group_id}/lag-summary | Get Consumer Group Lag Summary.
[**ListKafkaV3ConsumerGroups**](ConsumerGroupV3Api.md#ListKafkaV3ConsumerGroups) | **Get** /kafka/v3/clusters/{cluster_id}/consumer-groups | List Consumer Groups
[**ListKafkaV3ConsumerLags**](ConsumerGroupV3Api.md#ListKafkaV3ConsumerLags) | **Get** /kafka/v3/clusters/{cluster_id}/consumer-groups/{consumer_group_id}/lags | List Consumer Lags
[**ListKafkaV3Consumers**](ConsumerGroupV3Api.md#ListKafkaV3Consumers) | **Get** /kafka/v3/clusters/{cluster_id}/consumer-groups/{consumer_group_id}/consumers | List Consumers



## GetKafkaV3Consumer

> ConsumerData GetKafkaV3Consumer(ctx, clusterId, consumerGroupId, consumerId).Execute()

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
    resp, r, err := api_client.ConsumerGroupV3Api.GetKafkaV3Consumer(context.Background(), clusterId, consumerGroupId, consumerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsumerGroupV3Api.GetKafkaV3Consumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaV3Consumer`: ConsumerData
    fmt.Fprintf(os.Stdout, "Response from `ConsumerGroupV3Api.GetKafkaV3Consumer`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetKafkaV3ConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ConsumerData**](ConsumerData.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaV3ConsumerGroup

> ConsumerGroupData GetKafkaV3ConsumerGroup(ctx, clusterId, consumerGroupId).Execute()

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
    resp, r, err := api_client.ConsumerGroupV3Api.GetKafkaV3ConsumerGroup(context.Background(), clusterId, consumerGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsumerGroupV3Api.GetKafkaV3ConsumerGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaV3ConsumerGroup`: ConsumerGroupData
    fmt.Fprintf(os.Stdout, "Response from `ConsumerGroupV3Api.GetKafkaV3ConsumerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**consumerGroupId** | **string** | The consumer group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaV3ConsumerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConsumerGroupData**](ConsumerGroupData.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaV3ConsumerGroupLagSummary

> ConsumerGroupLagSummaryData GetKafkaV3ConsumerGroupLagSummary(ctx, clusterId, consumerGroupId).Execute()

Get Consumer Group Lag Summary.



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
    resp, r, err := api_client.ConsumerGroupV3Api.GetKafkaV3ConsumerGroupLagSummary(context.Background(), clusterId, consumerGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsumerGroupV3Api.GetKafkaV3ConsumerGroupLagSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaV3ConsumerGroupLagSummary`: ConsumerGroupLagSummaryData
    fmt.Fprintf(os.Stdout, "Response from `ConsumerGroupV3Api.GetKafkaV3ConsumerGroupLagSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**consumerGroupId** | **string** | The consumer group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaV3ConsumerGroupLagSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConsumerGroupLagSummaryData**](ConsumerGroupLagSummaryData.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaV3ConsumerGroups

> ConsumerGroupDataList ListKafkaV3ConsumerGroups(ctx, clusterId).Execute()

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
    resp, r, err := api_client.ConsumerGroupV3Api.ListKafkaV3ConsumerGroups(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsumerGroupV3Api.ListKafkaV3ConsumerGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaV3ConsumerGroups`: ConsumerGroupDataList
    fmt.Fprintf(os.Stdout, "Response from `ConsumerGroupV3Api.ListKafkaV3ConsumerGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaV3ConsumerGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsumerGroupDataList**](ConsumerGroupDataList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaV3ConsumerLags

> ConsumerLagDataList ListKafkaV3ConsumerLags(ctx, clusterId, consumerGroupId).Execute()

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
    resp, r, err := api_client.ConsumerGroupV3Api.ListKafkaV3ConsumerLags(context.Background(), clusterId, consumerGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsumerGroupV3Api.ListKafkaV3ConsumerLags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaV3ConsumerLags`: ConsumerLagDataList
    fmt.Fprintf(os.Stdout, "Response from `ConsumerGroupV3Api.ListKafkaV3ConsumerLags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**consumerGroupId** | **string** | The consumer group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaV3ConsumerLagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConsumerLagDataList**](ConsumerLagDataList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaV3Consumers

> ConsumerDataList ListKafkaV3Consumers(ctx, clusterId, consumerGroupId).Execute()

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
    resp, r, err := api_client.ConsumerGroupV3Api.ListKafkaV3Consumers(context.Background(), clusterId, consumerGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsumerGroupV3Api.ListKafkaV3Consumers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaV3Consumers`: ConsumerDataList
    fmt.Fprintf(os.Stdout, "Response from `ConsumerGroupV3Api.ListKafkaV3Consumers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**consumerGroupId** | **string** | The consumer group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaV3ConsumersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConsumerDataList**](ConsumerDataList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

