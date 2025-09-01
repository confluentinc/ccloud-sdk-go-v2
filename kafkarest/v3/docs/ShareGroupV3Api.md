# \ShareGroupV3Api

All URIs are relative to *https://pkc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKafkaShareGroup**](ShareGroupV3Api.md#GetKafkaShareGroup) | **Get** /kafka/v3/clusters/{cluster_id}/share-groups/{group_id} | Get Share Group
[**GetKafkaShareGroupConsumer**](ShareGroupV3Api.md#GetKafkaShareGroupConsumer) | **Get** /kafka/v3/clusters/{cluster_id}/share-groups/{group_id}/consumers/{consumer_id} | Get Share Group Consumer
[**ListKafkaShareGroupConsumerAssignments**](ShareGroupV3Api.md#ListKafkaShareGroupConsumerAssignments) | **Get** /kafka/v3/clusters/{cluster_id}/share-groups/{group_id}/consumers/{consumer_id}/assignments | List Share Group Consumer Assignments
[**ListKafkaShareGroupConsumers**](ShareGroupV3Api.md#ListKafkaShareGroupConsumers) | **Get** /kafka/v3/clusters/{cluster_id}/share-groups/{group_id}/consumers | List Share Group Consumers
[**ListKafkaShareGroups**](ShareGroupV3Api.md#ListKafkaShareGroups) | **Get** /kafka/v3/clusters/{cluster_id}/share-groups | List Share Groups



## GetKafkaShareGroup

> ShareGroupData GetKafkaShareGroup(ctx, clusterId, groupId).Execute()

Get Share Group



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
    groupId := "group-1" // string | The group ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ShareGroupV3Api.GetKafkaShareGroup(context.Background(), clusterId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShareGroupV3Api.GetKafkaShareGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaShareGroup`: ShareGroupData
    fmt.Fprintf(os.Stdout, "Response from `ShareGroupV3Api.GetKafkaShareGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**groupId** | **string** | The group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaShareGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ShareGroupData**](ShareGroupData.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaShareGroupConsumer

> ShareGroupConsumerData GetKafkaShareGroupConsumer(ctx, clusterId, groupId, consumerId).Execute()

Get Share Group Consumer



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
    groupId := "group-1" // string | The group ID.
    consumerId := "consumer-1" // string | The consumer ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ShareGroupV3Api.GetKafkaShareGroupConsumer(context.Background(), clusterId, groupId, consumerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShareGroupV3Api.GetKafkaShareGroupConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaShareGroupConsumer`: ShareGroupConsumerData
    fmt.Fprintf(os.Stdout, "Response from `ShareGroupV3Api.GetKafkaShareGroupConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**groupId** | **string** | The group ID. | 
**consumerId** | **string** | The consumer ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaShareGroupConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ShareGroupConsumerData**](ShareGroupConsumerData.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaShareGroupConsumerAssignments

> ShareGroupConsumerAssignmentDataList ListKafkaShareGroupConsumerAssignments(ctx, clusterId, groupId, consumerId).Execute()

List Share Group Consumer Assignments



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
    groupId := "group-1" // string | The group ID.
    consumerId := "consumer-1" // string | The consumer ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ShareGroupV3Api.ListKafkaShareGroupConsumerAssignments(context.Background(), clusterId, groupId, consumerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShareGroupV3Api.ListKafkaShareGroupConsumerAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaShareGroupConsumerAssignments`: ShareGroupConsumerAssignmentDataList
    fmt.Fprintf(os.Stdout, "Response from `ShareGroupV3Api.ListKafkaShareGroupConsumerAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**groupId** | **string** | The group ID. | 
**consumerId** | **string** | The consumer ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaShareGroupConsumerAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ShareGroupConsumerAssignmentDataList**](ShareGroupConsumerAssignmentDataList.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaShareGroupConsumers

> ShareGroupConsumerDataList ListKafkaShareGroupConsumers(ctx, clusterId, groupId).Execute()

List Share Group Consumers



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
    groupId := "group-1" // string | The group ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ShareGroupV3Api.ListKafkaShareGroupConsumers(context.Background(), clusterId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShareGroupV3Api.ListKafkaShareGroupConsumers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaShareGroupConsumers`: ShareGroupConsumerDataList
    fmt.Fprintf(os.Stdout, "Response from `ShareGroupV3Api.ListKafkaShareGroupConsumers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**groupId** | **string** | The group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaShareGroupConsumersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ShareGroupConsumerDataList**](ShareGroupConsumerDataList.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaShareGroups

> ShareGroupDataList ListKafkaShareGroups(ctx, clusterId).Execute()

List Share Groups



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
    resp, r, err := api_client.ShareGroupV3Api.ListKafkaShareGroups(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShareGroupV3Api.ListKafkaShareGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaShareGroups`: ShareGroupDataList
    fmt.Fprintf(os.Stdout, "Response from `ShareGroupV3Api.ListKafkaShareGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaShareGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShareGroupDataList**](ShareGroupDataList.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

