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

> DeleteKafkaV3ClusterConfig(ctx, clusterId, name)

Reset Cluster Config

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Resets the configuration parameter specified by ``name`` to its default value.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**name** | **string**| The configuration parameter name. | 

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


## DeleteKafkaV3TopicConfig

> DeleteKafkaV3TopicConfig(ctx, clusterId, topicName, name)

Reset Topic Config

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Resets the config with given `name` to its default value.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 
**name** | **string**| The configuration parameter name. | 

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


## GetKafkaV3ClusterConfig

> ClusterConfigData GetKafkaV3ClusterConfig(ctx, clusterId, name)

Get Cluster Config

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns the configuration parameter specified by ``name``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**name** | **string**| The configuration parameter name. | 

### Return type

[**ClusterConfigData**](ClusterConfigData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaV3TopicConfig

> TopicConfigData GetKafkaV3TopicConfig(ctx, clusterId, topicName, name)

Get Topic Config

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Return the config with the given `name`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 
**name** | **string**| The configuration parameter name. | 

### Return type

[**TopicConfigData**](TopicConfigData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaV3AllTopicConfigs

> TopicConfigDataList ListKafkaV3AllTopicConfigs(ctx, clusterId)

Get All Topic Configs

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns all topic configurations for topics hosted by the specified cluster.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 

### Return type

[**TopicConfigDataList**](TopicConfigDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaV3ClusterConfigs

> ClusterConfigDataList ListKafkaV3ClusterConfigs(ctx, clusterId)

List Cluster Configs

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns a list of configuration parameters for the specified Kafka cluster.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 

### Return type

[**ClusterConfigDataList**](ClusterConfigDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaV3TopicConfigs

> TopicConfigDataList ListKafkaV3TopicConfigs(ctx, clusterId, topicName)

List Topic Configs

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Return the list of configs that belong to the specified topic.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 

### Return type

[**TopicConfigDataList**](TopicConfigDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaV3ClusterConfig

> UpdateKafkaV3ClusterConfig(ctx, clusterId, name, optional)

Update Cluster Config

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Updates the configuration parameter specified by ``name``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**name** | **string**| The configuration parameter name. | 
 **optional** | ***UpdateKafkaV3ClusterConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateKafkaV3ClusterConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateConfigRequestData** | [**optional.Interface of UpdateConfigRequestData**](UpdateConfigRequestData.md)| The cluster configuration parameter update request. | 

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


## UpdateKafkaV3ClusterConfigs

> UpdateKafkaV3ClusterConfigs(ctx, clusterId, optional)

Batch Alter Cluster Configs

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Updates or deletes a set of Kafka cluster configuration parameters.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***UpdateKafkaV3ClusterConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateKafkaV3ClusterConfigsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alterConfigBatchRequestData** | [**optional.Interface of AlterConfigBatchRequestData**](AlterConfigBatchRequestData.md)| The alter cluster configuration parameter batch request. | 

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


## UpdateKafkaV3TopicConfig

> UpdateKafkaV3TopicConfig(ctx, clusterId, topicName, name, optional)

Update Topic Config

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Updates the config with given `name`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 
**name** | **string**| The configuration parameter name. | 
 **optional** | ***UpdateKafkaV3TopicConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateKafkaV3TopicConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateConfigRequestData** | [**optional.Interface of UpdateConfigRequestData**](UpdateConfigRequestData.md)| The topic configuration parameter update request. | 

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


## UpdateKafkaV3TopicConfigBatch

> UpdateKafkaV3TopicConfigBatch(ctx, clusterId, topicName, optional)

Batch Alter Topic Configs

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Updates or deletes a set of topic configs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 
 **optional** | ***UpdateKafkaV3TopicConfigBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateKafkaV3TopicConfigBatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **alterConfigBatchRequestData** | [**optional.Interface of AlterConfigBatchRequestData**](AlterConfigBatchRequestData.md)| The alter topic configuration parameter batch request. | 

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

