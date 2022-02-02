# \ClusterLinkingV3Api

All URIs are relative to *https://pkc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKafkaV3Link**](ClusterLinkingV3Api.md#CreateKafkaV3Link) | **Post** /kafka/v3/clusters/{cluster_id}/links | Create a cluster link
[**CreateKafkaV3MirrorTopic**](ClusterLinkingV3Api.md#CreateKafkaV3MirrorTopic) | **Post** /kafka/v3/clusters/{cluster_id}/links/{link_name}/mirrors | Create a mirror topic
[**DeleteKafkaV3Link**](ClusterLinkingV3Api.md#DeleteKafkaV3Link) | **Delete** /kafka/v3/clusters/{cluster_id}/links/{link_name} | Delete the cluster link
[**DeleteKafkaV3LinkConfig**](ClusterLinkingV3Api.md#DeleteKafkaV3LinkConfig) | **Delete** /kafka/v3/clusters/{cluster_id}/links/{link_name}/configs/{config_name} | Reset the given config to default value
[**GetKafkaV3Link**](ClusterLinkingV3Api.md#GetKafkaV3Link) | **Get** /kafka/v3/clusters/{cluster_id}/links/{link_name} | Describe the cluster link
[**GetKafkaV3LinkConfigs**](ClusterLinkingV3Api.md#GetKafkaV3LinkConfigs) | **Get** /kafka/v3/clusters/{cluster_id}/links/{link_name}/configs/{config_name} | Describe the config under the cluster link
[**ListKafkaV3LinkConfigs**](ClusterLinkingV3Api.md#ListKafkaV3LinkConfigs) | **Get** /kafka/v3/clusters/{cluster_id}/links/{link_name}/configs | List all configs of the cluster link
[**ListKafkaV3Links**](ClusterLinkingV3Api.md#ListKafkaV3Links) | **Get** /kafka/v3/clusters/{cluster_id}/links | List all cluster links in the dest cluster
[**ListKafkaV3MirrorTopics**](ClusterLinkingV3Api.md#ListKafkaV3MirrorTopics) | **Get** /kafka/v3/clusters/{cluster_id}/links/-/mirrors | List mirror topics
[**ListKafkaV3MirrorTopicsUnderLink**](ClusterLinkingV3Api.md#ListKafkaV3MirrorTopicsUnderLink) | **Get** /kafka/v3/clusters/{cluster_id}/links/{link_name}/mirrors | List mirror topics
[**ReadKafkaV3MirrorTopic**](ClusterLinkingV3Api.md#ReadKafkaV3MirrorTopic) | **Get** /kafka/v3/clusters/{cluster_id}/links/{link_name}/mirrors/{mirror_topic_name} | Describe the mirror topic
[**UpdateKafkaV3LinkConfig**](ClusterLinkingV3Api.md#UpdateKafkaV3LinkConfig) | **Put** /kafka/v3/clusters/{cluster_id}/links/{link_name}/configs/{config_name} | Alter the config under the cluster link
[**UpdateKafkaV3LinkConfigBatch**](ClusterLinkingV3Api.md#UpdateKafkaV3LinkConfigBatch) | **Post** /kafka/v3/clusters/{cluster_id}/links/{link_name}/configs:alter | Batch Alter Topic Configs
[**UpdateKafkaV3MirrorTopicsFailover**](ClusterLinkingV3Api.md#UpdateKafkaV3MirrorTopicsFailover) | **Post** /kafka/v3/clusters/{cluster_id}/links/{link_name}/mirrors:failover | Failover the mirror topics
[**UpdateKafkaV3MirrorTopicsPause**](ClusterLinkingV3Api.md#UpdateKafkaV3MirrorTopicsPause) | **Post** /kafka/v3/clusters/{cluster_id}/links/{link_name}/mirrors:pause | Pause the mirror topics
[**UpdateKafkaV3MirrorTopicsPromote**](ClusterLinkingV3Api.md#UpdateKafkaV3MirrorTopicsPromote) | **Post** /kafka/v3/clusters/{cluster_id}/links/{link_name}/mirrors:promote | Promote the mirror topics
[**UpdateKafkaV3MirrorTopicsResume**](ClusterLinkingV3Api.md#UpdateKafkaV3MirrorTopicsResume) | **Post** /kafka/v3/clusters/{cluster_id}/links/{link_name}/mirrors:resume | Resume the mirror topics



## CreateKafkaV3Link

> CreateKafkaV3Link(ctx, clusterId).CreateLinkRequestData(createLinkRequestData).Execute()

Create a cluster link



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
    createLinkRequestData := *openapiclient.NewCreateLinkRequestData("SourceClusterId_example") // CreateLinkRequestData | Create a cluster link (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.CreateKafkaV3Link(context.Background(), clusterId).CreateLinkRequestData(createLinkRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.CreateKafkaV3Link``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateKafkaV3LinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createLinkRequestData** | [**CreateLinkRequestData**](CreateLinkRequestData.md) | Create a cluster link | 

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


## CreateKafkaV3MirrorTopic

> CreateKafkaV3MirrorTopic(ctx, clusterId, linkName).CreateMirrorTopicRequestData(createMirrorTopicRequestData).Execute()

Create a mirror topic



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
    linkName := "link-sb1" // string | The link name
    createMirrorTopicRequestData := *openapiclient.NewCreateMirrorTopicRequestData("SourceTopicName_example") // CreateMirrorTopicRequestData | Name and configs of the topics mirroring from and mirroring to (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.CreateKafkaV3MirrorTopic(context.Background(), clusterId, linkName).CreateMirrorTopicRequestData(createMirrorTopicRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.CreateKafkaV3MirrorTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**linkName** | **string** | The link name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKafkaV3MirrorTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createMirrorTopicRequestData** | [**CreateMirrorTopicRequestData**](CreateMirrorTopicRequestData.md) | Name and configs of the topics mirroring from and mirroring to | 

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


## DeleteKafkaV3Link

> DeleteKafkaV3Link(ctx, clusterId, linkName).Execute()

Delete the cluster link



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
    linkName := "link-sb1" // string | The link name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.DeleteKafkaV3Link(context.Background(), clusterId, linkName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.DeleteKafkaV3Link``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**linkName** | **string** | The link name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKafkaV3LinkRequest struct via the builder pattern


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


## DeleteKafkaV3LinkConfig

> DeleteKafkaV3LinkConfig(ctx, clusterId, linkName, configName).Execute()

Reset the given config to default value



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
    linkName := "link-sb1" // string | The link name
    configName := "consumer.offset.sync.enable" // string | The link config name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.DeleteKafkaV3LinkConfig(context.Background(), clusterId, linkName, configName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.DeleteKafkaV3LinkConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**linkName** | **string** | The link name | 
**configName** | **string** | The link config name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKafkaV3LinkConfigRequest struct via the builder pattern


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


## GetKafkaV3Link

> ListLinksResponseData GetKafkaV3Link(ctx, clusterId, linkName).Execute()

Describe the cluster link



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
    linkName := "link-sb1" // string | The link name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.GetKafkaV3Link(context.Background(), clusterId, linkName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.GetKafkaV3Link``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaV3Link`: ListLinksResponseData
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.GetKafkaV3Link`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**linkName** | **string** | The link name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaV3LinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListLinksResponseData**](ListLinksResponseData.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaV3LinkConfigs

> ListLinkConfigsResponseData GetKafkaV3LinkConfigs(ctx, clusterId, linkName, configName).Execute()

Describe the config under the cluster link



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
    linkName := "link-sb1" // string | The link name
    configName := "consumer.offset.sync.enable" // string | The link config name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.GetKafkaV3LinkConfigs(context.Background(), clusterId, linkName, configName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.GetKafkaV3LinkConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaV3LinkConfigs`: ListLinkConfigsResponseData
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.GetKafkaV3LinkConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**linkName** | **string** | The link name | 
**configName** | **string** | The link config name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaV3LinkConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ListLinkConfigsResponseData**](ListLinkConfigsResponseData.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaV3LinkConfigs

> ListLinkConfigsResponseDataList ListKafkaV3LinkConfigs(ctx, clusterId, linkName).Execute()

List all configs of the cluster link



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
    linkName := "link-sb1" // string | The link name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.ListKafkaV3LinkConfigs(context.Background(), clusterId, linkName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.ListKafkaV3LinkConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaV3LinkConfigs`: ListLinkConfigsResponseDataList
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.ListKafkaV3LinkConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**linkName** | **string** | The link name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaV3LinkConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListLinkConfigsResponseDataList**](ListLinkConfigsResponseDataList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaV3Links

> ListLinksResponseDataList ListKafkaV3Links(ctx, clusterId).Execute()

List all cluster links in the dest cluster



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
    resp, r, err := api_client.ClusterLinkingV3Api.ListKafkaV3Links(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.ListKafkaV3Links``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaV3Links`: ListLinksResponseDataList
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.ListKafkaV3Links`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaV3LinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListLinksResponseDataList**](ListLinksResponseDataList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaV3MirrorTopics

> ListMirrorTopicsResponseDataList ListKafkaV3MirrorTopics(ctx, clusterId).MirrorStatus(mirrorStatus).Execute()

List mirror topics



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
    mirrorStatus := openapiclient.MirrorTopicStatus("active") // MirrorTopicStatus | The status of the mirror topic. If not specified, all mirror topics will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.ListKafkaV3MirrorTopics(context.Background(), clusterId).MirrorStatus(mirrorStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.ListKafkaV3MirrorTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaV3MirrorTopics`: ListMirrorTopicsResponseDataList
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.ListKafkaV3MirrorTopics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaV3MirrorTopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mirrorStatus** | [**MirrorTopicStatus**](MirrorTopicStatus.md) | The status of the mirror topic. If not specified, all mirror topics will be returned. | 

### Return type

[**ListMirrorTopicsResponseDataList**](ListMirrorTopicsResponseDataList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaV3MirrorTopicsUnderLink

> ListMirrorTopicsResponseDataList ListKafkaV3MirrorTopicsUnderLink(ctx, clusterId, linkName).MirrorStatus(mirrorStatus).Execute()

List mirror topics



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
    linkName := "link-sb1" // string | The link name
    mirrorStatus := openapiclient.MirrorTopicStatus("active") // MirrorTopicStatus | The status of the mirror topic. If not specified, all mirror topics will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.ListKafkaV3MirrorTopicsUnderLink(context.Background(), clusterId, linkName).MirrorStatus(mirrorStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.ListKafkaV3MirrorTopicsUnderLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaV3MirrorTopicsUnderLink`: ListMirrorTopicsResponseDataList
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.ListKafkaV3MirrorTopicsUnderLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**linkName** | **string** | The link name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaV3MirrorTopicsUnderLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mirrorStatus** | [**MirrorTopicStatus**](MirrorTopicStatus.md) | The status of the mirror topic. If not specified, all mirror topics will be returned. | 

### Return type

[**ListMirrorTopicsResponseDataList**](ListMirrorTopicsResponseDataList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadKafkaV3MirrorTopic

> ListMirrorTopicsResponseData ReadKafkaV3MirrorTopic(ctx, clusterId, linkName, mirrorTopicName).Execute()

Describe the mirror topic



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
    linkName := "link-sb1" // string | The link name
    mirrorTopicName := "topic-1" // string | Cluster Linking mirror topic name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.ReadKafkaV3MirrorTopic(context.Background(), clusterId, linkName, mirrorTopicName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.ReadKafkaV3MirrorTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadKafkaV3MirrorTopic`: ListMirrorTopicsResponseData
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.ReadKafkaV3MirrorTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**linkName** | **string** | The link name | 
**mirrorTopicName** | **string** | Cluster Linking mirror topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadKafkaV3MirrorTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ListMirrorTopicsResponseData**](ListMirrorTopicsResponseData.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaV3LinkConfig

> UpdateKafkaV3LinkConfig(ctx, clusterId, linkName, configName).UpdateLinkConfigRequestData(updateLinkConfigRequestData).Execute()

Alter the config under the cluster link



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
    linkName := "link-sb1" // string | The link name
    configName := "consumer.offset.sync.enable" // string | The link config name
    updateLinkConfigRequestData := *openapiclient.NewUpdateLinkConfigRequestData("Value_example") // UpdateLinkConfigRequestData | Link config value to update (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.UpdateKafkaV3LinkConfig(context.Background(), clusterId, linkName, configName).UpdateLinkConfigRequestData(updateLinkConfigRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.UpdateKafkaV3LinkConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**linkName** | **string** | The link name | 
**configName** | **string** | The link config name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKafkaV3LinkConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateLinkConfigRequestData** | [**UpdateLinkConfigRequestData**](UpdateLinkConfigRequestData.md) | Link config value to update | 

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


## UpdateKafkaV3LinkConfigBatch

> UpdateKafkaV3LinkConfigBatch(ctx, clusterId, linkName).AlterConfigBatchRequestData(alterConfigBatchRequestData).Execute()

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
    linkName := "link-sb1" // string | The link name
    alterConfigBatchRequestData := *openapiclient.NewAlterConfigBatchRequestData([]openapiclient.AlterConfigBatchRequestDataData{*openapiclient.NewAlterConfigBatchRequestDataData("Name_example")}) // AlterConfigBatchRequestData |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.UpdateKafkaV3LinkConfigBatch(context.Background(), clusterId, linkName).AlterConfigBatchRequestData(alterConfigBatchRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.UpdateKafkaV3LinkConfigBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**linkName** | **string** | The link name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKafkaV3LinkConfigBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **alterConfigBatchRequestData** | [**AlterConfigBatchRequestData**](AlterConfigBatchRequestData.md) |  | 

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


## UpdateKafkaV3MirrorTopicsFailover

> AlterMirrorStatusResponseDataList UpdateKafkaV3MirrorTopicsFailover(ctx, clusterId, linkName).ValidateOnly(validateOnly).AlterMirrorsRequestData(alterMirrorsRequestData).Execute()

Failover the mirror topics



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
    linkName := "link-sb1" // string | The link name
    validateOnly := false // bool | To validate if the link can be created or not, but not to create it. Default: false (optional)
    alterMirrorsRequestData := *openapiclient.NewAlterMirrorsRequestData([]string{"MirrorTopicNames_example"}) // AlterMirrorsRequestData | Name of the topics to apply the changes (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.UpdateKafkaV3MirrorTopicsFailover(context.Background(), clusterId, linkName).ValidateOnly(validateOnly).AlterMirrorsRequestData(alterMirrorsRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.UpdateKafkaV3MirrorTopicsFailover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKafkaV3MirrorTopicsFailover`: AlterMirrorStatusResponseDataList
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.UpdateKafkaV3MirrorTopicsFailover`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**linkName** | **string** | The link name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKafkaV3MirrorTopicsFailoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **bool** | To validate if the link can be created or not, but not to create it. Default: false | 
 **alterMirrorsRequestData** | [**AlterMirrorsRequestData**](AlterMirrorsRequestData.md) | Name of the topics to apply the changes | 

### Return type

[**AlterMirrorStatusResponseDataList**](AlterMirrorStatusResponseDataList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaV3MirrorTopicsPause

> AlterMirrorStatusResponseDataList UpdateKafkaV3MirrorTopicsPause(ctx, clusterId, linkName).ValidateOnly(validateOnly).AlterMirrorsRequestData(alterMirrorsRequestData).Execute()

Pause the mirror topics



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
    linkName := "link-sb1" // string | The link name
    validateOnly := false // bool | To validate if the link can be created or not, but not to create it. Default: false (optional)
    alterMirrorsRequestData := *openapiclient.NewAlterMirrorsRequestData([]string{"MirrorTopicNames_example"}) // AlterMirrorsRequestData | Name of the topics to apply the changes (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.UpdateKafkaV3MirrorTopicsPause(context.Background(), clusterId, linkName).ValidateOnly(validateOnly).AlterMirrorsRequestData(alterMirrorsRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.UpdateKafkaV3MirrorTopicsPause``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKafkaV3MirrorTopicsPause`: AlterMirrorStatusResponseDataList
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.UpdateKafkaV3MirrorTopicsPause`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**linkName** | **string** | The link name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKafkaV3MirrorTopicsPauseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **bool** | To validate if the link can be created or not, but not to create it. Default: false | 
 **alterMirrorsRequestData** | [**AlterMirrorsRequestData**](AlterMirrorsRequestData.md) | Name of the topics to apply the changes | 

### Return type

[**AlterMirrorStatusResponseDataList**](AlterMirrorStatusResponseDataList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaV3MirrorTopicsPromote

> AlterMirrorStatusResponseDataList UpdateKafkaV3MirrorTopicsPromote(ctx, clusterId, linkName).ValidateOnly(validateOnly).AlterMirrorsRequestData(alterMirrorsRequestData).Execute()

Promote the mirror topics



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
    linkName := "link-sb1" // string | The link name
    validateOnly := false // bool | To validate if the link can be created or not, but not to create it. Default: false (optional)
    alterMirrorsRequestData := *openapiclient.NewAlterMirrorsRequestData([]string{"MirrorTopicNames_example"}) // AlterMirrorsRequestData | Name of the topics to apply the changes (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.UpdateKafkaV3MirrorTopicsPromote(context.Background(), clusterId, linkName).ValidateOnly(validateOnly).AlterMirrorsRequestData(alterMirrorsRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.UpdateKafkaV3MirrorTopicsPromote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKafkaV3MirrorTopicsPromote`: AlterMirrorStatusResponseDataList
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.UpdateKafkaV3MirrorTopicsPromote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**linkName** | **string** | The link name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKafkaV3MirrorTopicsPromoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **bool** | To validate if the link can be created or not, but not to create it. Default: false | 
 **alterMirrorsRequestData** | [**AlterMirrorsRequestData**](AlterMirrorsRequestData.md) | Name of the topics to apply the changes | 

### Return type

[**AlterMirrorStatusResponseDataList**](AlterMirrorStatusResponseDataList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaV3MirrorTopicsResume

> AlterMirrorStatusResponseDataList UpdateKafkaV3MirrorTopicsResume(ctx, clusterId, linkName).ValidateOnly(validateOnly).AlterMirrorsRequestData(alterMirrorsRequestData).Execute()

Resume the mirror topics



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
    linkName := "link-sb1" // string | The link name
    validateOnly := false // bool | To validate if the link can be created or not, but not to create it. Default: false (optional)
    alterMirrorsRequestData := *openapiclient.NewAlterMirrorsRequestData([]string{"MirrorTopicNames_example"}) // AlterMirrorsRequestData | Name of the topics to apply the changes (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.UpdateKafkaV3MirrorTopicsResume(context.Background(), clusterId, linkName).ValidateOnly(validateOnly).AlterMirrorsRequestData(alterMirrorsRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.UpdateKafkaV3MirrorTopicsResume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKafkaV3MirrorTopicsResume`: AlterMirrorStatusResponseDataList
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.UpdateKafkaV3MirrorTopicsResume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**linkName** | **string** | The link name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKafkaV3MirrorTopicsResumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **bool** | To validate if the link can be created or not, but not to create it. Default: false | 
 **alterMirrorsRequestData** | [**AlterMirrorsRequestData**](AlterMirrorsRequestData.md) | Name of the topics to apply the changes | 

### Return type

[**AlterMirrorStatusResponseDataList**](AlterMirrorStatusResponseDataList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

