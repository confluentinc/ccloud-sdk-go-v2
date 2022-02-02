# \ConfigsV3Api

All URIs are relative to *https://pkc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteKafkaV3ClusterConfig**](ConfigsV3Api.md#DeleteKafkaV3ClusterConfig) | **Delete** /kafka/v3/clusters/{cluster_id}/broker-configs/{name} | Reset Cluster Config
[**DeleteKafkaV3TopicConfig**](ConfigsV3Api.md#DeleteKafkaV3TopicConfig) | **Delete** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/configs/{name} | Reset Topic Config
[**GetKafkaV3ClusterConfig**](ConfigsV3Api.md#GetKafkaV3ClusterConfig) | **Get** /kafka/v3/clusters/{cluster_id}/broker-configs/{name} | Get Cluster Config
[**GetKafkaV3TopicConfig**](ConfigsV3Api.md#GetKafkaV3TopicConfig) | **Get** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/configs/{name} | Get Topic Config
[**ListKafkaV3AllTopicConfigs**](ConfigsV3Api.md#ListKafkaV3AllTopicConfigs) | **Get** /kafka/v3/clusters/{cluster_id}/topics/-/configs | Get All Topic Configs
[**ListKafkaV3ClusterConfigs**](ConfigsV3Api.md#ListKafkaV3ClusterConfigs) | **Get** /kafka/v3/clusters/{cluster_id}/broker-configs | List Cluster Configs
[**ListKafkaV3TopicConfigs**](ConfigsV3Api.md#ListKafkaV3TopicConfigs) | **Get** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/configs | List Topic Configs
[**UpdateKafkaV3ClusterConfig**](ConfigsV3Api.md#UpdateKafkaV3ClusterConfig) | **Put** /kafka/v3/clusters/{cluster_id}/broker-configs/{name} | Update Cluster Config
[**UpdateKafkaV3ClusterConfigs**](ConfigsV3Api.md#UpdateKafkaV3ClusterConfigs) | **Post** /kafka/v3/clusters/{cluster_id}/broker-configs:alter | Batch Alter Cluster Configs
[**UpdateKafkaV3TopicConfig**](ConfigsV3Api.md#UpdateKafkaV3TopicConfig) | **Put** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/configs/{name} | Update Topic Config
[**UpdateKafkaV3TopicConfigBatch**](ConfigsV3Api.md#UpdateKafkaV3TopicConfigBatch) | **Post** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/configs:alter | Batch Alter Topic Configs



## DeleteKafkaV3ClusterConfig

> DeleteKafkaV3ClusterConfig(ctx, clusterId, name).Execute()

Reset Cluster Config



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
    resp, r, err := api_client.ConfigsV3Api.DeleteKafkaV3ClusterConfig(context.Background(), clusterId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsV3Api.DeleteKafkaV3ClusterConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteKafkaV3ClusterConfigRequest struct via the builder pattern


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


## DeleteKafkaV3TopicConfig

> DeleteKafkaV3TopicConfig(ctx, clusterId, topicName, name).Execute()

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
    resp, r, err := api_client.ConfigsV3Api.DeleteKafkaV3TopicConfig(context.Background(), clusterId, topicName, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsV3Api.DeleteKafkaV3TopicConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteKafkaV3TopicConfigRequest struct via the builder pattern


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


## GetKafkaV3ClusterConfig

> ClusterConfigData GetKafkaV3ClusterConfig(ctx, clusterId, name).Execute()

Get Cluster Config



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
    resp, r, err := api_client.ConfigsV3Api.GetKafkaV3ClusterConfig(context.Background(), clusterId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsV3Api.GetKafkaV3ClusterConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaV3ClusterConfig`: ClusterConfigData
    fmt.Fprintf(os.Stdout, "Response from `ConfigsV3Api.GetKafkaV3ClusterConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**name** | **string** | The configuration parameter name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaV3ClusterConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterConfigData**](ClusterConfigData.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaV3TopicConfig

> TopicConfigData GetKafkaV3TopicConfig(ctx, clusterId, topicName, name).Execute()

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
    resp, r, err := api_client.ConfigsV3Api.GetKafkaV3TopicConfig(context.Background(), clusterId, topicName, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsV3Api.GetKafkaV3TopicConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaV3TopicConfig`: TopicConfigData
    fmt.Fprintf(os.Stdout, "Response from `ConfigsV3Api.GetKafkaV3TopicConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetKafkaV3TopicConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TopicConfigData**](TopicConfigData.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaV3AllTopicConfigs

> TopicConfigDataList ListKafkaV3AllTopicConfigs(ctx, clusterId).Execute()

Get All Topic Configs



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
    resp, r, err := api_client.ConfigsV3Api.ListKafkaV3AllTopicConfigs(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsV3Api.ListKafkaV3AllTopicConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaV3AllTopicConfigs`: TopicConfigDataList
    fmt.Fprintf(os.Stdout, "Response from `ConfigsV3Api.ListKafkaV3AllTopicConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaV3AllTopicConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TopicConfigDataList**](TopicConfigDataList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaV3ClusterConfigs

> ClusterConfigDataList ListKafkaV3ClusterConfigs(ctx, clusterId).Execute()

List Cluster Configs



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
    resp, r, err := api_client.ConfigsV3Api.ListKafkaV3ClusterConfigs(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsV3Api.ListKafkaV3ClusterConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaV3ClusterConfigs`: ClusterConfigDataList
    fmt.Fprintf(os.Stdout, "Response from `ConfigsV3Api.ListKafkaV3ClusterConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaV3ClusterConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterConfigDataList**](ClusterConfigDataList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaV3TopicConfigs

> TopicConfigDataList ListKafkaV3TopicConfigs(ctx, clusterId, topicName).Execute()

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
    resp, r, err := api_client.ConfigsV3Api.ListKafkaV3TopicConfigs(context.Background(), clusterId, topicName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsV3Api.ListKafkaV3TopicConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaV3TopicConfigs`: TopicConfigDataList
    fmt.Fprintf(os.Stdout, "Response from `ConfigsV3Api.ListKafkaV3TopicConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**topicName** | **string** | The topic name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaV3TopicConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TopicConfigDataList**](TopicConfigDataList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaV3ClusterConfig

> UpdateKafkaV3ClusterConfig(ctx, clusterId, name).UpdateConfigRequestData(updateConfigRequestData).Execute()

Update Cluster Config



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
    resp, r, err := api_client.ConfigsV3Api.UpdateKafkaV3ClusterConfig(context.Background(), clusterId, name).UpdateConfigRequestData(updateConfigRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsV3Api.UpdateKafkaV3ClusterConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateKafkaV3ClusterConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateConfigRequestData** | [**UpdateConfigRequestData**](UpdateConfigRequestData.md) | The cluster configuration parameter update request. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaV3ClusterConfigs

> UpdateKafkaV3ClusterConfigs(ctx, clusterId).AlterConfigBatchRequestData(alterConfigBatchRequestData).Execute()

Batch Alter Cluster Configs



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
    resp, r, err := api_client.ConfigsV3Api.UpdateKafkaV3ClusterConfigs(context.Background(), clusterId).AlterConfigBatchRequestData(alterConfigBatchRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsV3Api.UpdateKafkaV3ClusterConfigs``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateKafkaV3ClusterConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alterConfigBatchRequestData** | [**AlterConfigBatchRequestData**](AlterConfigBatchRequestData.md) | The alter cluster configuration parameter batch request. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaV3TopicConfig

> UpdateKafkaV3TopicConfig(ctx, clusterId, topicName, name).UpdateConfigRequestData(updateConfigRequestData).Execute()

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
    resp, r, err := api_client.ConfigsV3Api.UpdateKafkaV3TopicConfig(context.Background(), clusterId, topicName, name).UpdateConfigRequestData(updateConfigRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsV3Api.UpdateKafkaV3TopicConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateKafkaV3TopicConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateConfigRequestData** | [**UpdateConfigRequestData**](UpdateConfigRequestData.md) | The topic configuration parameter update request. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaV3TopicConfigBatch

> UpdateKafkaV3TopicConfigBatch(ctx, clusterId, topicName).AlterConfigBatchRequestData(alterConfigBatchRequestData).Execute()

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
    resp, r, err := api_client.ConfigsV3Api.UpdateKafkaV3TopicConfigBatch(context.Background(), clusterId, topicName).AlterConfigBatchRequestData(alterConfigBatchRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsV3Api.UpdateKafkaV3TopicConfigBatch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateKafkaV3TopicConfigBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **alterConfigBatchRequestData** | [**AlterConfigBatchRequestData**](AlterConfigBatchRequestData.md) | The alter topic configuration parameter batch request. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

