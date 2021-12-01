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

> CreateKafkaV3Link(ctx, clusterId, optional)

Create a cluster link

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***CreateKafkaV3LinkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateKafkaV3LinkOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createLinkRequestData** | [**optional.Interface of CreateLinkRequestData**](CreateLinkRequestData.md)| Create a cluster link | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateKafkaV3MirrorTopic

> CreateKafkaV3MirrorTopic(ctx, clusterId, linkName, optional)

Create a mirror topic

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Create a topic in the destination cluster mirroring a topic in the source cluster

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***CreateKafkaV3MirrorTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateKafkaV3MirrorTopicOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createMirrorTopicRequestData** | [**optional.Interface of CreateMirrorTopicRequestData**](CreateMirrorTopicRequestData.md)| Name and configs of the topics mirroring from and mirroring to | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKafkaV3Link

> DeleteKafkaV3Link(ctx, clusterId, linkName)

Delete the cluster link

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKafkaV3LinkConfig

> DeleteKafkaV3LinkConfig(ctx, clusterId, linkName, configName)

Reset the given config to default value

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
**configName** | **string**| The link config name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaV3Link

> ListLinksResponseData GetKafkaV3Link(ctx, clusterId, linkName)

Describe the cluster link

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 

### Return type

[**ListLinksResponseData**](ListLinksResponseData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaV3LinkConfigs

> ListLinkConfigsResponseData GetKafkaV3LinkConfigs(ctx, clusterId, linkName, configName)

Describe the config under the cluster link

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
**configName** | **string**| The link config name | 

### Return type

[**ListLinkConfigsResponseData**](ListLinkConfigsResponseData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaV3LinkConfigs

> ListLinkConfigsResponseDataList ListKafkaV3LinkConfigs(ctx, clusterId, linkName)

List all configs of the cluster link

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 

### Return type

[**ListLinkConfigsResponseDataList**](ListLinkConfigsResponseDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaV3Links

> ListLinksResponseDataList ListKafkaV3Links(ctx, clusterId)

List all cluster links in the dest cluster

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 

### Return type

[**ListLinksResponseDataList**](ListLinksResponseDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaV3MirrorTopics

> ListMirrorTopicsResponseDataList ListKafkaV3MirrorTopics(ctx, clusterId, optional)

List mirror topics

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  List all mirror topics in the cluster

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***ListKafkaV3MirrorTopicsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListKafkaV3MirrorTopicsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mirrorStatus** | [**optional.Interface of MirrorTopicStatus**](.md)| The status of the mirror topic. If not specified, all mirror topics will be returned. | 

### Return type

[**ListMirrorTopicsResponseDataList**](ListMirrorTopicsResponseDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaV3MirrorTopicsUnderLink

> ListMirrorTopicsResponseDataList ListKafkaV3MirrorTopicsUnderLink(ctx, clusterId, linkName, optional)

List mirror topics

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  List all mirror topics under the link

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***ListKafkaV3MirrorTopicsUnderLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListKafkaV3MirrorTopicsUnderLinkOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mirrorStatus** | [**optional.Interface of MirrorTopicStatus**](.md)| The status of the mirror topic. If not specified, all mirror topics will be returned. | 

### Return type

[**ListMirrorTopicsResponseDataList**](ListMirrorTopicsResponseDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadKafkaV3MirrorTopic

> ListMirrorTopicsResponseData ReadKafkaV3MirrorTopic(ctx, clusterId, linkName, mirrorTopicName)

Describe the mirror topic

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
**mirrorTopicName** | **string**| Cluster Linking mirror topic name | 

### Return type

[**ListMirrorTopicsResponseData**](ListMirrorTopicsResponseData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaV3LinkConfig

> UpdateKafkaV3LinkConfig(ctx, clusterId, linkName, configName, optional)

Alter the config under the cluster link

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
**configName** | **string**| The link config name | 
 **optional** | ***UpdateKafkaV3LinkConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateKafkaV3LinkConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateLinkConfigRequestData** | [**optional.Interface of UpdateLinkConfigRequestData**](UpdateLinkConfigRequestData.md)| Link config value to update | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaV3LinkConfigBatch

> UpdateKafkaV3LinkConfigBatch(ctx, clusterId, linkName, optional)

Batch Alter Topic Configs

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Batch Alter Link Configs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***UpdateKafkaV3LinkConfigBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateKafkaV3LinkConfigBatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **alterConfigBatchRequestData** | [**optional.Interface of AlterConfigBatchRequestData**](AlterConfigBatchRequestData.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaV3MirrorTopicsFailover

> AlterMirrorStatusResponseDataList UpdateKafkaV3MirrorTopicsFailover(ctx, clusterId, linkName, optional)

Failover the mirror topics

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***UpdateKafkaV3MirrorTopicsFailoverOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateKafkaV3MirrorTopicsFailoverOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **optional.Bool**| To validate if the link can be created or not, but not to create it. Default: false | 
 **alterMirrorsRequestData** | [**optional.Interface of AlterMirrorsRequestData**](AlterMirrorsRequestData.md)| Name of the topics to apply the changes | 

### Return type

[**AlterMirrorStatusResponseDataList**](AlterMirrorStatusResponseDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaV3MirrorTopicsPause

> AlterMirrorStatusResponseDataList UpdateKafkaV3MirrorTopicsPause(ctx, clusterId, linkName, optional)

Pause the mirror topics

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***UpdateKafkaV3MirrorTopicsPauseOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateKafkaV3MirrorTopicsPauseOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **optional.Bool**| To validate if the link can be created or not, but not to create it. Default: false | 
 **alterMirrorsRequestData** | [**optional.Interface of AlterMirrorsRequestData**](AlterMirrorsRequestData.md)| Name of the topics to apply the changes | 

### Return type

[**AlterMirrorStatusResponseDataList**](AlterMirrorStatusResponseDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaV3MirrorTopicsPromote

> AlterMirrorStatusResponseDataList UpdateKafkaV3MirrorTopicsPromote(ctx, clusterId, linkName, optional)

Promote the mirror topics

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***UpdateKafkaV3MirrorTopicsPromoteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateKafkaV3MirrorTopicsPromoteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **optional.Bool**| To validate if the link can be created or not, but not to create it. Default: false | 
 **alterMirrorsRequestData** | [**optional.Interface of AlterMirrorsRequestData**](AlterMirrorsRequestData.md)| Name of the topics to apply the changes | 

### Return type

[**AlterMirrorStatusResponseDataList**](AlterMirrorStatusResponseDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaV3MirrorTopicsResume

> AlterMirrorStatusResponseDataList UpdateKafkaV3MirrorTopicsResume(ctx, clusterId, linkName, optional)

Resume the mirror topics

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***UpdateKafkaV3MirrorTopicsResumeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateKafkaV3MirrorTopicsResumeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **optional.Bool**| To validate if the link can be created or not, but not to create it. Default: false | 
 **alterMirrorsRequestData** | [**optional.Interface of AlterMirrorsRequestData**](AlterMirrorsRequestData.md)| Name of the topics to apply the changes | 

### Return type

[**AlterMirrorStatusResponseDataList**](AlterMirrorStatusResponseDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

