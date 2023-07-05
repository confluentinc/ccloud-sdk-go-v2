# \ConfigsV3Api

All URIs are relative to *https://pkc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteKafkaClusterConfig**](ConfigsV3Api.md#DeleteKafkaClusterConfig) | **Delete** /kafka/v3/clusters/{cluster_id}/broker-configs/{name} | Reset Dynamic Broker Config
[**DeleteKafkaTopicConfig**](ConfigsV3Api.md#DeleteKafkaTopicConfig) | **Delete** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/configs/{name} | Reset Topic Config
[**GetKafkaClusterConfig**](ConfigsV3Api.md#GetKafkaClusterConfig) | **Get** /kafka/v3/clusters/{cluster_id}/broker-configs/{name} | Get Dynamic Broker Config
[**GetKafkaTopicConfig**](ConfigsV3Api.md#GetKafkaTopicConfig) | **Get** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/configs/{name} | Get Topic Config
[**ListKafkaAllTopicConfigs**](ConfigsV3Api.md#ListKafkaAllTopicConfigs) | **Get** /kafka/v3/clusters/{cluster_id}/topics/-/configs | List All Topic Configs
[**ListKafkaClusterConfigs**](ConfigsV3Api.md#ListKafkaClusterConfigs) | **Get** /kafka/v3/clusters/{cluster_id}/broker-configs | List Dynamic Broker Configs
[**ListKafkaDefaultTopicConfigs**](ConfigsV3Api.md#ListKafkaDefaultTopicConfigs) | **Get** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/default-configs | List New Topic Default Configs
[**ListKafkaTopicConfigs**](ConfigsV3Api.md#ListKafkaTopicConfigs) | **Get** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/configs | List Topic Configs
[**UpdateKafkaClusterConfig**](ConfigsV3Api.md#UpdateKafkaClusterConfig) | **Put** /kafka/v3/clusters/{cluster_id}/broker-configs/{name} | Update Dynamic Broker Config
[**UpdateKafkaClusterConfigs**](ConfigsV3Api.md#UpdateKafkaClusterConfigs) | **Post** /kafka/v3/clusters/{cluster_id}/broker-configs:alter | Batch Alter Dynamic Broker Configs
[**UpdateKafkaTopicConfig**](ConfigsV3Api.md#UpdateKafkaTopicConfig) | **Put** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/configs/{name} | Update Topic Config
[**UpdateKafkaTopicConfigBatch**](ConfigsV3Api.md#UpdateKafkaTopicConfigBatch) | **Post** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/configs:alter | Batch Alter Topic Configs



## DeleteKafkaClusterConfig

> DeleteKafkaClusterConfig(ctx, clusterId, name).Execute()

Reset Dynamic Broker Config



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
    name := "compression.type" // string | The configuration parameter name.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigsV3Api.DeleteKafkaClusterConfig(context.Background(), clusterId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsV3Api.DeleteKafkaClusterConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**name** | **string** | The configuration parameter name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKafkaClusterConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKafkaTopicConfig

> DeleteKafkaTopicConfig(ctx, clusterId, topicName, name).Execute()

Reset Topic Config



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
    name := "compression.type" // string | The configuration parameter name.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigsV3Api.DeleteKafkaTopicConfig(context.Background(), clusterId, topicName, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsV3Api.DeleteKafkaTopicConfig``: %v\n", err)
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
**name** | **string** | The configuration parameter name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKafkaTopicConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaClusterConfig

> ClusterConfigData GetKafkaClusterConfig(ctx, clusterId, name).Execute()

Get Dynamic Broker Config



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
    name := "compression.type" // string | The configuration parameter name.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigsV3Api.GetKafkaClusterConfig(context.Background(), clusterId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsV3Api.GetKafkaClusterConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaClusterConfig`: ClusterConfigData
    fmt.Fprintf(os.Stdout, "Response from `ConfigsV3Api.GetKafkaClusterConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**name** | **string** | The configuration parameter name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaClusterConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterConfigData**](ClusterConfigData.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaTopicConfig

> TopicConfigData GetKafkaTopicConfig(ctx, clusterId, topicName, name).Execute()

Get Topic Config



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
    name := "compression.type" // string | The configuration parameter name.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigsV3Api.GetKafkaTopicConfig(context.Background(), clusterId, topicName, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsV3Api.GetKafkaTopicConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaTopicConfig`: TopicConfigData
    fmt.Fprintf(os.Stdout, "Response from `ConfigsV3Api.GetKafkaTopicConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**topicName** | **string** | The topic name. | 
**name** | **string** | The configuration parameter name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaTopicConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TopicConfigData**](TopicConfigData.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaAllTopicConfigs

> TopicConfigDataList ListKafkaAllTopicConfigs(ctx, clusterId).Execute()

List All Topic Configs



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
    resp, r, err := api_client.ConfigsV3Api.ListKafkaAllTopicConfigs(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsV3Api.ListKafkaAllTopicConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaAllTopicConfigs`: TopicConfigDataList
    fmt.Fprintf(os.Stdout, "Response from `ConfigsV3Api.ListKafkaAllTopicConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaAllTopicConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TopicConfigDataList**](TopicConfigDataList.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaClusterConfigs

> ClusterConfigDataList ListKafkaClusterConfigs(ctx, clusterId).Execute()

List Dynamic Broker Configs



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
    resp, r, err := api_client.ConfigsV3Api.ListKafkaClusterConfigs(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsV3Api.ListKafkaClusterConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaClusterConfigs`: ClusterConfigDataList
    fmt.Fprintf(os.Stdout, "Response from `ConfigsV3Api.ListKafkaClusterConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaClusterConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterConfigDataList**](ClusterConfigDataList.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaDefaultTopicConfigs

> TopicConfigDataList ListKafkaDefaultTopicConfigs(ctx, clusterId, topicName).Execute()

List New Topic Default Configs



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
    resp, r, err := api_client.ConfigsV3Api.ListKafkaDefaultTopicConfigs(context.Background(), clusterId, topicName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsV3Api.ListKafkaDefaultTopicConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaDefaultTopicConfigs`: TopicConfigDataList
    fmt.Fprintf(os.Stdout, "Response from `ConfigsV3Api.ListKafkaDefaultTopicConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**topicName** | **string** | The topic name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaDefaultTopicConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TopicConfigDataList**](TopicConfigDataList.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaTopicConfigs

> TopicConfigDataList ListKafkaTopicConfigs(ctx, clusterId, topicName).Execute()

List Topic Configs



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
    resp, r, err := api_client.ConfigsV3Api.ListKafkaTopicConfigs(context.Background(), clusterId, topicName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsV3Api.ListKafkaTopicConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaTopicConfigs`: TopicConfigDataList
    fmt.Fprintf(os.Stdout, "Response from `ConfigsV3Api.ListKafkaTopicConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**topicName** | **string** | The topic name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaTopicConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TopicConfigDataList**](TopicConfigDataList.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaClusterConfig

> UpdateKafkaClusterConfig(ctx, clusterId, name).UpdateConfigRequestData(updateConfigRequestData).Execute()

Update Dynamic Broker Config



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
    name := "compression.type" // string | The configuration parameter name.
    updateConfigRequestData := *openapiclient.NewUpdateConfigRequestData() // UpdateConfigRequestData | The cluster configuration parameter update request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigsV3Api.UpdateKafkaClusterConfig(context.Background(), clusterId, name).UpdateConfigRequestData(updateConfigRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsV3Api.UpdateKafkaClusterConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**name** | **string** | The configuration parameter name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKafkaClusterConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateConfigRequestData** | [**UpdateConfigRequestData**](UpdateConfigRequestData.md) | The cluster configuration parameter update request. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaClusterConfigs

> UpdateKafkaClusterConfigs(ctx, clusterId).AlterConfigBatchRequestData(alterConfigBatchRequestData).Execute()

Batch Alter Dynamic Broker Configs



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
    alterConfigBatchRequestData := *openapiclient.NewAlterConfigBatchRequestData([]openapiclient.AlterConfigBatchRequestDataData{*openapiclient.NewAlterConfigBatchRequestDataData("Name_example")}) // AlterConfigBatchRequestData | The alter cluster configuration parameter batch request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigsV3Api.UpdateKafkaClusterConfigs(context.Background(), clusterId).AlterConfigBatchRequestData(alterConfigBatchRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsV3Api.UpdateKafkaClusterConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKafkaClusterConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alterConfigBatchRequestData** | [**AlterConfigBatchRequestData**](AlterConfigBatchRequestData.md) | The alter cluster configuration parameter batch request. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaTopicConfig

> UpdateKafkaTopicConfig(ctx, clusterId, topicName, name).UpdateConfigRequestData(updateConfigRequestData).Execute()

Update Topic Config



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
    name := "compression.type" // string | The configuration parameter name.
    updateConfigRequestData := *openapiclient.NewUpdateConfigRequestData() // UpdateConfigRequestData | The topic configuration parameter update request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigsV3Api.UpdateKafkaTopicConfig(context.Background(), clusterId, topicName, name).UpdateConfigRequestData(updateConfigRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsV3Api.UpdateKafkaTopicConfig``: %v\n", err)
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
**name** | **string** | The configuration parameter name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKafkaTopicConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateConfigRequestData** | [**UpdateConfigRequestData**](UpdateConfigRequestData.md) | The topic configuration parameter update request. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaTopicConfigBatch

> UpdateKafkaTopicConfigBatch(ctx, clusterId, topicName).AlterConfigBatchRequestData(alterConfigBatchRequestData).Execute()

Batch Alter Topic Configs



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
    alterConfigBatchRequestData := *openapiclient.NewAlterConfigBatchRequestData([]openapiclient.AlterConfigBatchRequestDataData{*openapiclient.NewAlterConfigBatchRequestDataData("Name_example")}) // AlterConfigBatchRequestData | The alter topic configuration parameter batch request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigsV3Api.UpdateKafkaTopicConfigBatch(context.Background(), clusterId, topicName).AlterConfigBatchRequestData(alterConfigBatchRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsV3Api.UpdateKafkaTopicConfigBatch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateKafkaTopicConfigBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **alterConfigBatchRequestData** | [**AlterConfigBatchRequestData**](AlterConfigBatchRequestData.md) | The alter topic configuration parameter batch request. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

