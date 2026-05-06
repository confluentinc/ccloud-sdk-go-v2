# \RtceTopicsRtceV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRtceV1RtceTopic**](RtceTopicsRtceV1Api.md#CreateRtceV1RtceTopic) | **Post** /rtce/v1/rtce-topics | Create a Rtce Topic
[**DeleteRtceV1RtceTopic**](RtceTopicsRtceV1Api.md#DeleteRtceV1RtceTopic) | **Delete** /rtce/v1/rtce-topics/{topic_name} | Delete a Rtce Topic
[**GetRtceV1RtceTopic**](RtceTopicsRtceV1Api.md#GetRtceV1RtceTopic) | **Get** /rtce/v1/rtce-topics/{topic_name} | Read a Rtce Topic
[**ListRtceV1RtceTopics**](RtceTopicsRtceV1Api.md#ListRtceV1RtceTopics) | **Get** /rtce/v1/rtce-topics | List of Rtce Topics
[**UpdateRtceV1RtceTopic**](RtceTopicsRtceV1Api.md#UpdateRtceV1RtceTopic) | **Patch** /rtce/v1/rtce-topics/{topic_name} | Update a Rtce Topic



## CreateRtceV1RtceTopic

> RtceV1RtceTopic CreateRtceV1RtceTopic(ctx).RtceV1RtceTopic(rtceV1RtceTopic).Execute()

Create a Rtce Topic



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
    rtceV1RtceTopic := *openapiclient.NewRtceV1RtceTopic() // RtceV1RtceTopic |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RtceTopicsRtceV1Api.CreateRtceV1RtceTopic(context.Background()).RtceV1RtceTopic(rtceV1RtceTopic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RtceTopicsRtceV1Api.CreateRtceV1RtceTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRtceV1RtceTopic`: RtceV1RtceTopic
    fmt.Fprintf(os.Stdout, "Response from `RtceTopicsRtceV1Api.CreateRtceV1RtceTopic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRtceV1RtceTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rtceV1RtceTopic** | [**RtceV1RtceTopic**](RtceV1RtceTopic.md) |  | 

### Return type

[**RtceV1RtceTopic**](rtce.v1.RtceTopic.md)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRtceV1RtceTopic

> DeleteRtceV1RtceTopic(ctx, topicName).Environment(environment).SpecKafkaCluster(specKafkaCluster).Execute()

Delete a Rtce Topic



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
    topicName := "topicName_example" // string | The Kafka topic name containing the data for the RTCE topic.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RtceTopicsRtceV1Api.DeleteRtceV1RtceTopic(context.Background(), topicName).Environment(environment).SpecKafkaCluster(specKafkaCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RtceTopicsRtceV1Api.DeleteRtceV1RtceTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicName** | **string** | The Kafka topic name containing the data for the RTCE topic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRtceV1RtceTopicRequest struct via the builder pattern


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


## GetRtceV1RtceTopic

> RtceV1RtceTopic GetRtceV1RtceTopic(ctx, topicName).Environment(environment).SpecKafkaCluster(specKafkaCluster).Execute()

Read a Rtce Topic



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
    topicName := "topicName_example" // string | The Kafka topic name containing the data for the RTCE topic.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RtceTopicsRtceV1Api.GetRtceV1RtceTopic(context.Background(), topicName).Environment(environment).SpecKafkaCluster(specKafkaCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RtceTopicsRtceV1Api.GetRtceV1RtceTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRtceV1RtceTopic`: RtceV1RtceTopic
    fmt.Fprintf(os.Stdout, "Response from `RtceTopicsRtceV1Api.GetRtceV1RtceTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicName** | **string** | The Kafka topic name containing the data for the RTCE topic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRtceV1RtceTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 
 **specKafkaCluster** | **string** | Scope the operation to the given spec.kafka_cluster. | 


### Return type

[**RtceV1RtceTopic**](rtce.v1.RtceTopic.md)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRtceV1RtceTopics

> RtceV1RtceTopicList ListRtceV1RtceTopics(ctx).Environment(environment).SpecKafkaCluster(specKafkaCluster).SpecCloud(specCloud).SpecRegion(specRegion).PageSize(pageSize).PageToken(pageToken).Execute()

List of Rtce Topics



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
    specCloud := "AWS" // string | Filter the results by exact match for spec.cloud. (optional)
    specRegion := "us-west-2" // string | Filter the results by exact match for spec.region. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RtceTopicsRtceV1Api.ListRtceV1RtceTopics(context.Background()).Environment(environment).SpecKafkaCluster(specKafkaCluster).SpecCloud(specCloud).SpecRegion(specRegion).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RtceTopicsRtceV1Api.ListRtceV1RtceTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRtceV1RtceTopics`: RtceV1RtceTopicList
    fmt.Fprintf(os.Stdout, "Response from `RtceTopicsRtceV1Api.ListRtceV1RtceTopics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRtceV1RtceTopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **specKafkaCluster** | **string** | Filter the results by exact match for spec.kafka_cluster. | 
 **specCloud** | **string** | Filter the results by exact match for spec.cloud. | 
 **specRegion** | **string** | Filter the results by exact match for spec.region. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**RtceV1RtceTopicList**](rtce.v1.RtceTopicList.md)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRtceV1RtceTopic

> RtceV1RtceTopic UpdateRtceV1RtceTopic(ctx, topicName).RtceV1RtceTopicUpdate(rtceV1RtceTopicUpdate).Execute()

Update a Rtce Topic



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
    topicName := "topicName_example" // string | The Kafka topic name containing the data for the RTCE topic.
    rtceV1RtceTopicUpdate := *openapiclient.NewRtceV1RtceTopicUpdate() // RtceV1RtceTopicUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RtceTopicsRtceV1Api.UpdateRtceV1RtceTopic(context.Background(), topicName).RtceV1RtceTopicUpdate(rtceV1RtceTopicUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RtceTopicsRtceV1Api.UpdateRtceV1RtceTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRtceV1RtceTopic`: RtceV1RtceTopic
    fmt.Fprintf(os.Stdout, "Response from `RtceTopicsRtceV1Api.UpdateRtceV1RtceTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicName** | **string** | The Kafka topic name containing the data for the RTCE topic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRtceV1RtceTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rtceV1RtceTopicUpdate** | [**RtceV1RtceTopicUpdate**](RtceV1RtceTopicUpdate.md) |  | 

### Return type

[**RtceV1RtceTopic**](rtce.v1.RtceTopic.md)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

