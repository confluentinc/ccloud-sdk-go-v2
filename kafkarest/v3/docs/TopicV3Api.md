# \TopicV3Api

All URIs are relative to *https://pkc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKafkaV3Topic**](TopicV3Api.md#CreateKafkaV3Topic) | **Post** /kafka/v3/clusters/{cluster_id}/topics | Create Topic
[**DeleteKafkaV3Topic**](TopicV3Api.md#DeleteKafkaV3Topic) | **Delete** /kafka/v3/clusters/{cluster_id}/topics/{topic_name} | Delete Topic
[**GetKafkaV3Topic**](TopicV3Api.md#GetKafkaV3Topic) | **Get** /kafka/v3/clusters/{cluster_id}/topics/{topic_name} | Get Topic
[**ListKafkaV3Topics**](TopicV3Api.md#ListKafkaV3Topics) | **Get** /kafka/v3/clusters/{cluster_id}/topics | List Topics



## CreateKafkaV3Topic

> TopicData CreateKafkaV3Topic(ctx, clusterId).CreateTopicRequestData(createTopicRequestData).Execute()

Create Topic



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
    createTopicRequestData := *openapiclient.NewCreateTopicRequestData("TopicName_example") // CreateTopicRequestData | The topic creation request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopicV3Api.CreateKafkaV3Topic(context.Background(), clusterId).CreateTopicRequestData(createTopicRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicV3Api.CreateKafkaV3Topic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKafkaV3Topic`: TopicData
    fmt.Fprintf(os.Stdout, "Response from `TopicV3Api.CreateKafkaV3Topic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKafkaV3TopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTopicRequestData** | [**CreateTopicRequestData**](CreateTopicRequestData.md) | The topic creation request. | 

### Return type

[**TopicData**](TopicData.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKafkaV3Topic

> DeleteKafkaV3Topic(ctx, clusterId, topicName).Execute()

Delete Topic



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
    resp, r, err := api_client.TopicV3Api.DeleteKafkaV3Topic(context.Background(), clusterId, topicName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicV3Api.DeleteKafkaV3Topic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**topicName** | **string** | The topic name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKafkaV3TopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaV3Topic

> TopicData GetKafkaV3Topic(ctx, clusterId, topicName).Execute()

Get Topic



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
    resp, r, err := api_client.TopicV3Api.GetKafkaV3Topic(context.Background(), clusterId, topicName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicV3Api.GetKafkaV3Topic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaV3Topic`: TopicData
    fmt.Fprintf(os.Stdout, "Response from `TopicV3Api.GetKafkaV3Topic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**topicName** | **string** | The topic name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaV3TopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TopicData**](TopicData.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaV3Topics

> TopicDataList ListKafkaV3Topics(ctx, clusterId).Execute()

List Topics



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
    resp, r, err := api_client.TopicV3Api.ListKafkaV3Topics(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicV3Api.ListKafkaV3Topics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaV3Topics`: TopicDataList
    fmt.Fprintf(os.Stdout, "Response from `TopicV3Api.ListKafkaV3Topics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaV3TopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TopicDataList**](TopicDataList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

