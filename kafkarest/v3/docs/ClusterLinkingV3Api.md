# \ClusterLinkingV3Api

All URIs are relative to *https://pkc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKafkaLink**](ClusterLinkingV3Api.md#CreateKafkaLink) | **Post** /kafka/v3/clusters/{cluster_id}/links | Create a cluster link
[**CreateKafkaMirrorTopic**](ClusterLinkingV3Api.md#CreateKafkaMirrorTopic) | **Post** /kafka/v3/clusters/{cluster_id}/links/{link_name}/mirrors | Create a mirror topic
[**DeleteKafkaLink**](ClusterLinkingV3Api.md#DeleteKafkaLink) | **Delete** /kafka/v3/clusters/{cluster_id}/links/{link_name} | Delete the cluster link
[**DeleteKafkaLinkConfig**](ClusterLinkingV3Api.md#DeleteKafkaLinkConfig) | **Delete** /kafka/v3/clusters/{cluster_id}/links/{link_name}/configs/{config_name} | Reset the given config to default value
[**GetKafkaLink**](ClusterLinkingV3Api.md#GetKafkaLink) | **Get** /kafka/v3/clusters/{cluster_id}/links/{link_name} | Describe the cluster link
[**GetKafkaLinkConfigs**](ClusterLinkingV3Api.md#GetKafkaLinkConfigs) | **Get** /kafka/v3/clusters/{cluster_id}/links/{link_name}/configs/{config_name} | Describe the config under the cluster link
[**ListKafkaLinkConfigs**](ClusterLinkingV3Api.md#ListKafkaLinkConfigs) | **Get** /kafka/v3/clusters/{cluster_id}/links/{link_name}/configs | List all configs of the cluster link
[**ListKafkaLinks**](ClusterLinkingV3Api.md#ListKafkaLinks) | **Get** /kafka/v3/clusters/{cluster_id}/links | List all cluster links in the dest cluster
[**ListKafkaMirrorTopics**](ClusterLinkingV3Api.md#ListKafkaMirrorTopics) | **Get** /kafka/v3/clusters/{cluster_id}/links/-/mirrors | List mirror topics
[**ListKafkaMirrorTopicsUnderLink**](ClusterLinkingV3Api.md#ListKafkaMirrorTopicsUnderLink) | **Get** /kafka/v3/clusters/{cluster_id}/links/{link_name}/mirrors | List mirror topics
[**ReadKafkaMirrorTopic**](ClusterLinkingV3Api.md#ReadKafkaMirrorTopic) | **Get** /kafka/v3/clusters/{cluster_id}/links/{link_name}/mirrors/{mirror_topic_name} | Describe the mirror topic
[**UpdateKafkaLinkConfig**](ClusterLinkingV3Api.md#UpdateKafkaLinkConfig) | **Put** /kafka/v3/clusters/{cluster_id}/links/{link_name}/configs/{config_name} | Alter the config under the cluster link
[**UpdateKafkaLinkConfigBatch**](ClusterLinkingV3Api.md#UpdateKafkaLinkConfigBatch) | **Put** /kafka/v3/clusters/{cluster_id}/links/{link_name}/configs:alter | Batch Alter Cluster Link Configs
[**UpdateKafkaMirrorTopicsFailover**](ClusterLinkingV3Api.md#UpdateKafkaMirrorTopicsFailover) | **Post** /kafka/v3/clusters/{cluster_id}/links/{link_name}/mirrors:failover | Failover the mirror topics
[**UpdateKafkaMirrorTopicsPause**](ClusterLinkingV3Api.md#UpdateKafkaMirrorTopicsPause) | **Post** /kafka/v3/clusters/{cluster_id}/links/{link_name}/mirrors:pause | Pause the mirror topics
[**UpdateKafkaMirrorTopicsPromote**](ClusterLinkingV3Api.md#UpdateKafkaMirrorTopicsPromote) | **Post** /kafka/v3/clusters/{cluster_id}/links/{link_name}/mirrors:promote | Promote the mirror topics
[**UpdateKafkaMirrorTopicsResume**](ClusterLinkingV3Api.md#UpdateKafkaMirrorTopicsResume) | **Post** /kafka/v3/clusters/{cluster_id}/links/{link_name}/mirrors:resume | Resume the mirror topics
[**UpdateKafkaMirrorTopicsReverseAndPauseMirror**](ClusterLinkingV3Api.md#UpdateKafkaMirrorTopicsReverseAndPauseMirror) | **Post** /kafka/v3/clusters/{cluster_id}/links/{link_name}/mirrors:reverse-and-pause-mirror | Reverse the local mirror topic and Pause the remote mirror topic
[**UpdateKafkaMirrorTopicsReverseAndStartMirror**](ClusterLinkingV3Api.md#UpdateKafkaMirrorTopicsReverseAndStartMirror) | **Post** /kafka/v3/clusters/{cluster_id}/links/{link_name}/mirrors:reverse-and-start-mirror | Reverse the local mirror topic and start the remote mirror topic
[**UpdateKafkaMirrorTopicsTruncateAndRestoreMirror**](ClusterLinkingV3Api.md#UpdateKafkaMirrorTopicsTruncateAndRestoreMirror) | **Post** /kafka/v3/clusters/{cluster_id}/links/{link_name}/mirrors:truncate-and-restore | Truncates the local topic to the remote stopped mirror log end offsets and restores mirroring to the local topic to mirror from the remote topic



## CreateKafkaLink

> CreateKafkaLink(ctx, clusterId).LinkName(linkName).ValidateOnly(validateOnly).ValidateLink(validateLink).CreateLinkRequestData(createLinkRequestData).Execute()

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
    linkName := "link-sb1" // string | The link name
    validateOnly := false // bool | To validate the action can be performed successfully or not. Default: false (optional)
    validateLink := false // bool | To synchronously validate that the source cluster ID is expected and the dest cluster has the permission to read topics in the source cluster. Default: true (optional)
    createLinkRequestData := *openapiclient.NewCreateLinkRequestData() // CreateLinkRequestData | Create a cluster link (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.CreateKafkaLink(context.Background(), clusterId).LinkName(linkName).ValidateOnly(validateOnly).ValidateLink(validateLink).CreateLinkRequestData(createLinkRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.CreateKafkaLink``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateKafkaLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **linkName** | **string** | The link name | 
 **validateOnly** | **bool** | To validate the action can be performed successfully or not. Default: false | 
 **validateLink** | **bool** | To synchronously validate that the source cluster ID is expected and the dest cluster has the permission to read topics in the source cluster. Default: true | 
 **createLinkRequestData** | [**CreateLinkRequestData**](CreateLinkRequestData.md) | Create a cluster link | 

### Return type

 (empty response body)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateKafkaMirrorTopic

> CreateKafkaMirrorTopic(ctx, clusterId, linkName).CreateMirrorTopicRequestData(createMirrorTopicRequestData).Execute()

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
    createMirrorTopicRequestData := *openapiclient.NewCreateMirrorTopicRequestData("SourceTopicName_example") // CreateMirrorTopicRequestData | Name and configs of the topics mirroring from and mirroring to. Note that Confluent Cloud allows only specific replication factor values. Because of that the replication factor field should either be omitted or it should use one of the allowed values (see https://docs.confluent.io/cloud/current/client-apps/optimizing/durability.html). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.CreateKafkaMirrorTopic(context.Background(), clusterId, linkName).CreateMirrorTopicRequestData(createMirrorTopicRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.CreateKafkaMirrorTopic``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateKafkaMirrorTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createMirrorTopicRequestData** | [**CreateMirrorTopicRequestData**](CreateMirrorTopicRequestData.md) | Name and configs of the topics mirroring from and mirroring to. Note that Confluent Cloud allows only specific replication factor values. Because of that the replication factor field should either be omitted or it should use one of the allowed values (see https://docs.confluent.io/cloud/current/client-apps/optimizing/durability.html). | 

### Return type

 (empty response body)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKafkaLink

> DeleteKafkaLink(ctx, clusterId, linkName).Force(force).ValidateOnly(validateOnly).Execute()

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
    force := false // bool | Force the action. Default: false (optional)
    validateOnly := false // bool | To validate the action can be performed successfully or not. Default: false (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.DeleteKafkaLink(context.Background(), clusterId, linkName).Force(force).ValidateOnly(validateOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.DeleteKafkaLink``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteKafkaLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **bool** | Force the action. Default: false | 
 **validateOnly** | **bool** | To validate the action can be performed successfully or not. Default: false | 

### Return type

 (empty response body)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKafkaLinkConfig

> DeleteKafkaLinkConfig(ctx, clusterId, linkName, configName).Execute()

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
    resp, r, err := api_client.ClusterLinkingV3Api.DeleteKafkaLinkConfig(context.Background(), clusterId, linkName, configName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.DeleteKafkaLinkConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteKafkaLinkConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaLink

> ListLinksResponseData GetKafkaLink(ctx, clusterId, linkName).IncludeTasks(includeTasks).Execute()

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
    includeTasks := false // bool | Whether to include cluster linking tasks in the response. Default: false (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.GetKafkaLink(context.Background(), clusterId, linkName).IncludeTasks(includeTasks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.GetKafkaLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaLink`: ListLinksResponseData
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.GetKafkaLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**linkName** | **string** | The link name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeTasks** | **bool** | Whether to include cluster linking tasks in the response. Default: false | 

### Return type

[**ListLinksResponseData**](ListLinksResponseData.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaLinkConfigs

> ListLinkConfigsResponseData GetKafkaLinkConfigs(ctx, clusterId, linkName, configName).Execute()

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
    resp, r, err := api_client.ClusterLinkingV3Api.GetKafkaLinkConfigs(context.Background(), clusterId, linkName, configName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.GetKafkaLinkConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaLinkConfigs`: ListLinkConfigsResponseData
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.GetKafkaLinkConfigs`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetKafkaLinkConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ListLinkConfigsResponseData**](ListLinkConfigsResponseData.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaLinkConfigs

> ListLinkConfigsResponseDataList ListKafkaLinkConfigs(ctx, clusterId, linkName).Execute()

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
    resp, r, err := api_client.ClusterLinkingV3Api.ListKafkaLinkConfigs(context.Background(), clusterId, linkName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.ListKafkaLinkConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaLinkConfigs`: ListLinkConfigsResponseDataList
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.ListKafkaLinkConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**linkName** | **string** | The link name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaLinkConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListLinkConfigsResponseDataList**](ListLinkConfigsResponseDataList.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaLinks

> ListLinksResponseDataList ListKafkaLinks(ctx, clusterId).Execute()

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
    resp, r, err := api_client.ClusterLinkingV3Api.ListKafkaLinks(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.ListKafkaLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaLinks`: ListLinksResponseDataList
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.ListKafkaLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListLinksResponseDataList**](ListLinksResponseDataList.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaMirrorTopics

> ListMirrorTopicsResponseDataList ListKafkaMirrorTopics(ctx, clusterId).MirrorStatus(mirrorStatus).Execute()

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
    mirrorStatus := openapiclient.MirrorTopicStatus("ACTIVE") // MirrorTopicStatus | The status of the mirror topic. If not specified, all mirror topics will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.ListKafkaMirrorTopics(context.Background(), clusterId).MirrorStatus(mirrorStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.ListKafkaMirrorTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaMirrorTopics`: ListMirrorTopicsResponseDataList
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.ListKafkaMirrorTopics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaMirrorTopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mirrorStatus** | [**MirrorTopicStatus**](MirrorTopicStatus.md) | The status of the mirror topic. If not specified, all mirror topics will be returned. | 

### Return type

[**ListMirrorTopicsResponseDataList**](ListMirrorTopicsResponseDataList.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaMirrorTopicsUnderLink

> ListMirrorTopicsResponseDataList ListKafkaMirrorTopicsUnderLink(ctx, clusterId, linkName).MirrorStatus(mirrorStatus).Execute()

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
    mirrorStatus := openapiclient.MirrorTopicStatus("ACTIVE") // MirrorTopicStatus | The status of the mirror topic. If not specified, all mirror topics will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.ListKafkaMirrorTopicsUnderLink(context.Background(), clusterId, linkName).MirrorStatus(mirrorStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.ListKafkaMirrorTopicsUnderLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaMirrorTopicsUnderLink`: ListMirrorTopicsResponseDataList
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.ListKafkaMirrorTopicsUnderLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**linkName** | **string** | The link name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaMirrorTopicsUnderLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mirrorStatus** | [**MirrorTopicStatus**](MirrorTopicStatus.md) | The status of the mirror topic. If not specified, all mirror topics will be returned. | 

### Return type

[**ListMirrorTopicsResponseDataList**](ListMirrorTopicsResponseDataList.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadKafkaMirrorTopic

> ListMirrorTopicsResponseData ReadKafkaMirrorTopic(ctx, clusterId, linkName, mirrorTopicName).IncludeStateTransitionErrors(includeStateTransitionErrors).Execute()

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
    includeStateTransitionErrors := false // bool | Whether to include mirror state transition errors in the response. Default: false (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.ReadKafkaMirrorTopic(context.Background(), clusterId, linkName, mirrorTopicName).IncludeStateTransitionErrors(includeStateTransitionErrors).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.ReadKafkaMirrorTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadKafkaMirrorTopic`: ListMirrorTopicsResponseData
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.ReadKafkaMirrorTopic`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiReadKafkaMirrorTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **includeStateTransitionErrors** | **bool** | Whether to include mirror state transition errors in the response. Default: false | 

### Return type

[**ListMirrorTopicsResponseData**](ListMirrorTopicsResponseData.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaLinkConfig

> UpdateKafkaLinkConfig(ctx, clusterId, linkName, configName).UpdateLinkConfigRequestData(updateLinkConfigRequestData).Execute()

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
    resp, r, err := api_client.ClusterLinkingV3Api.UpdateKafkaLinkConfig(context.Background(), clusterId, linkName, configName).UpdateLinkConfigRequestData(updateLinkConfigRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.UpdateKafkaLinkConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateKafkaLinkConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateLinkConfigRequestData** | [**UpdateLinkConfigRequestData**](UpdateLinkConfigRequestData.md) | Link config value to update | 

### Return type

 (empty response body)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaLinkConfigBatch

> UpdateKafkaLinkConfigBatch(ctx, clusterId, linkName).ValidateOnly(validateOnly).AlterConfigBatchRequestData(alterConfigBatchRequestData).Execute()

Batch Alter Cluster Link Configs



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
    validateOnly := false // bool | To validate the action can be performed successfully or not. Default: false (optional)
    alterConfigBatchRequestData := *openapiclient.NewAlterConfigBatchRequestData([]openapiclient.AlterConfigBatchRequestDataData{*openapiclient.NewAlterConfigBatchRequestDataData("Name_example")}) // AlterConfigBatchRequestData |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.UpdateKafkaLinkConfigBatch(context.Background(), clusterId, linkName).ValidateOnly(validateOnly).AlterConfigBatchRequestData(alterConfigBatchRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.UpdateKafkaLinkConfigBatch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateKafkaLinkConfigBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **bool** | To validate the action can be performed successfully or not. Default: false | 
 **alterConfigBatchRequestData** | [**AlterConfigBatchRequestData**](AlterConfigBatchRequestData.md) |  | 

### Return type

 (empty response body)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaMirrorTopicsFailover

> AlterMirrorStatusResponseDataList UpdateKafkaMirrorTopicsFailover(ctx, clusterId, linkName).ValidateOnly(validateOnly).AlterMirrorsRequestData(alterMirrorsRequestData).Execute()

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
    validateOnly := false // bool | To validate the action can be performed successfully or not. Default: false (optional)
    alterMirrorsRequestData := *openapiclient.NewAlterMirrorsRequestData() // AlterMirrorsRequestData | Mirror topics to be altered. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.UpdateKafkaMirrorTopicsFailover(context.Background(), clusterId, linkName).ValidateOnly(validateOnly).AlterMirrorsRequestData(alterMirrorsRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.UpdateKafkaMirrorTopicsFailover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKafkaMirrorTopicsFailover`: AlterMirrorStatusResponseDataList
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.UpdateKafkaMirrorTopicsFailover`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**linkName** | **string** | The link name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKafkaMirrorTopicsFailoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **bool** | To validate the action can be performed successfully or not. Default: false | 
 **alterMirrorsRequestData** | [**AlterMirrorsRequestData**](AlterMirrorsRequestData.md) | Mirror topics to be altered. | 

### Return type

[**AlterMirrorStatusResponseDataList**](AlterMirrorStatusResponseDataList.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaMirrorTopicsPause

> AlterMirrorStatusResponseDataList UpdateKafkaMirrorTopicsPause(ctx, clusterId, linkName).ValidateOnly(validateOnly).AlterMirrorsRequestData(alterMirrorsRequestData).Execute()

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
    validateOnly := false // bool | To validate the action can be performed successfully or not. Default: false (optional)
    alterMirrorsRequestData := *openapiclient.NewAlterMirrorsRequestData() // AlterMirrorsRequestData | Mirror topics to be altered. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.UpdateKafkaMirrorTopicsPause(context.Background(), clusterId, linkName).ValidateOnly(validateOnly).AlterMirrorsRequestData(alterMirrorsRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.UpdateKafkaMirrorTopicsPause``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKafkaMirrorTopicsPause`: AlterMirrorStatusResponseDataList
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.UpdateKafkaMirrorTopicsPause`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**linkName** | **string** | The link name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKafkaMirrorTopicsPauseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **bool** | To validate the action can be performed successfully or not. Default: false | 
 **alterMirrorsRequestData** | [**AlterMirrorsRequestData**](AlterMirrorsRequestData.md) | Mirror topics to be altered. | 

### Return type

[**AlterMirrorStatusResponseDataList**](AlterMirrorStatusResponseDataList.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaMirrorTopicsPromote

> AlterMirrorStatusResponseDataList UpdateKafkaMirrorTopicsPromote(ctx, clusterId, linkName).ValidateOnly(validateOnly).AlterMirrorsRequestData(alterMirrorsRequestData).Execute()

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
    validateOnly := false // bool | To validate the action can be performed successfully or not. Default: false (optional)
    alterMirrorsRequestData := *openapiclient.NewAlterMirrorsRequestData() // AlterMirrorsRequestData | Mirror topics to be altered. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.UpdateKafkaMirrorTopicsPromote(context.Background(), clusterId, linkName).ValidateOnly(validateOnly).AlterMirrorsRequestData(alterMirrorsRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.UpdateKafkaMirrorTopicsPromote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKafkaMirrorTopicsPromote`: AlterMirrorStatusResponseDataList
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.UpdateKafkaMirrorTopicsPromote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**linkName** | **string** | The link name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKafkaMirrorTopicsPromoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **bool** | To validate the action can be performed successfully or not. Default: false | 
 **alterMirrorsRequestData** | [**AlterMirrorsRequestData**](AlterMirrorsRequestData.md) | Mirror topics to be altered. | 

### Return type

[**AlterMirrorStatusResponseDataList**](AlterMirrorStatusResponseDataList.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaMirrorTopicsResume

> AlterMirrorStatusResponseDataList UpdateKafkaMirrorTopicsResume(ctx, clusterId, linkName).ValidateOnly(validateOnly).AlterMirrorsRequestData(alterMirrorsRequestData).Execute()

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
    validateOnly := false // bool | To validate the action can be performed successfully or not. Default: false (optional)
    alterMirrorsRequestData := *openapiclient.NewAlterMirrorsRequestData() // AlterMirrorsRequestData | Mirror topics to be altered. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.UpdateKafkaMirrorTopicsResume(context.Background(), clusterId, linkName).ValidateOnly(validateOnly).AlterMirrorsRequestData(alterMirrorsRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.UpdateKafkaMirrorTopicsResume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKafkaMirrorTopicsResume`: AlterMirrorStatusResponseDataList
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.UpdateKafkaMirrorTopicsResume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**linkName** | **string** | The link name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKafkaMirrorTopicsResumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **bool** | To validate the action can be performed successfully or not. Default: false | 
 **alterMirrorsRequestData** | [**AlterMirrorsRequestData**](AlterMirrorsRequestData.md) | Mirror topics to be altered. | 

### Return type

[**AlterMirrorStatusResponseDataList**](AlterMirrorStatusResponseDataList.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaMirrorTopicsReverseAndPauseMirror

> AlterMirrorStatusResponseDataList UpdateKafkaMirrorTopicsReverseAndPauseMirror(ctx, clusterId, linkName).ValidateOnly(validateOnly).AlterMirrorsRequestData(alterMirrorsRequestData).Execute()

Reverse the local mirror topic and Pause the remote mirror topic



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
    validateOnly := false // bool | To validate the action can be performed successfully or not. Default: false (optional)
    alterMirrorsRequestData := *openapiclient.NewAlterMirrorsRequestData() // AlterMirrorsRequestData | Mirror topics to be altered. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.UpdateKafkaMirrorTopicsReverseAndPauseMirror(context.Background(), clusterId, linkName).ValidateOnly(validateOnly).AlterMirrorsRequestData(alterMirrorsRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.UpdateKafkaMirrorTopicsReverseAndPauseMirror``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKafkaMirrorTopicsReverseAndPauseMirror`: AlterMirrorStatusResponseDataList
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.UpdateKafkaMirrorTopicsReverseAndPauseMirror`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**linkName** | **string** | The link name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKafkaMirrorTopicsReverseAndPauseMirrorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **bool** | To validate the action can be performed successfully or not. Default: false | 
 **alterMirrorsRequestData** | [**AlterMirrorsRequestData**](AlterMirrorsRequestData.md) | Mirror topics to be altered. | 

### Return type

[**AlterMirrorStatusResponseDataList**](AlterMirrorStatusResponseDataList.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaMirrorTopicsReverseAndStartMirror

> AlterMirrorStatusResponseDataList UpdateKafkaMirrorTopicsReverseAndStartMirror(ctx, clusterId, linkName).ValidateOnly(validateOnly).AlterMirrorsRequestData(alterMirrorsRequestData).Execute()

Reverse the local mirror topic and start the remote mirror topic



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
    validateOnly := false // bool | To validate the action can be performed successfully or not. Default: false (optional)
    alterMirrorsRequestData := *openapiclient.NewAlterMirrorsRequestData() // AlterMirrorsRequestData | Mirror topics to be altered. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.UpdateKafkaMirrorTopicsReverseAndStartMirror(context.Background(), clusterId, linkName).ValidateOnly(validateOnly).AlterMirrorsRequestData(alterMirrorsRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.UpdateKafkaMirrorTopicsReverseAndStartMirror``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKafkaMirrorTopicsReverseAndStartMirror`: AlterMirrorStatusResponseDataList
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.UpdateKafkaMirrorTopicsReverseAndStartMirror`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**linkName** | **string** | The link name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKafkaMirrorTopicsReverseAndStartMirrorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **bool** | To validate the action can be performed successfully or not. Default: false | 
 **alterMirrorsRequestData** | [**AlterMirrorsRequestData**](AlterMirrorsRequestData.md) | Mirror topics to be altered. | 

### Return type

[**AlterMirrorStatusResponseDataList**](AlterMirrorStatusResponseDataList.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaMirrorTopicsTruncateAndRestoreMirror

> AlterMirrorStatusResponseDataList UpdateKafkaMirrorTopicsTruncateAndRestoreMirror(ctx, clusterId, linkName).IncludePartitionLevelTruncationData(includePartitionLevelTruncationData).ValidateOnly(validateOnly).AlterMirrorsRequestData(alterMirrorsRequestData).Execute()

Truncates the local topic to the remote stopped mirror log end offsets and restores mirroring to the local topic to mirror from the remote topic



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
    includePartitionLevelTruncationData := false // bool | Whether to include partition level truncation information when truncating and restoring a topic in the response. Default: false (optional)
    validateOnly := false // bool | To validate the action can be performed successfully or not. Default: false (optional)
    alterMirrorsRequestData := *openapiclient.NewAlterMirrorsRequestData() // AlterMirrorsRequestData | Mirror topics to be altered. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLinkingV3Api.UpdateKafkaMirrorTopicsTruncateAndRestoreMirror(context.Background(), clusterId, linkName).IncludePartitionLevelTruncationData(includePartitionLevelTruncationData).ValidateOnly(validateOnly).AlterMirrorsRequestData(alterMirrorsRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLinkingV3Api.UpdateKafkaMirrorTopicsTruncateAndRestoreMirror``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKafkaMirrorTopicsTruncateAndRestoreMirror`: AlterMirrorStatusResponseDataList
    fmt.Fprintf(os.Stdout, "Response from `ClusterLinkingV3Api.UpdateKafkaMirrorTopicsTruncateAndRestoreMirror`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**linkName** | **string** | The link name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKafkaMirrorTopicsTruncateAndRestoreMirrorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includePartitionLevelTruncationData** | **bool** | Whether to include partition level truncation information when truncating and restoring a topic in the response. Default: false | 
 **validateOnly** | **bool** | To validate the action can be performed successfully or not. Default: false | 
 **alterMirrorsRequestData** | [**AlterMirrorsRequestData**](AlterMirrorsRequestData.md) | Mirror topics to be altered. | 

### Return type

[**AlterMirrorStatusResponseDataList**](AlterMirrorStatusResponseDataList.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

