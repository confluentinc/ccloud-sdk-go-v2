# \ConfigsApi

All URIs are relative to *https://pkc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListKafkaV3DefaultTopicConfigs**](ConfigsApi.md#ListKafkaV3DefaultTopicConfigs) | **Get** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/default-configs | List Default Topic Configs



## ListKafkaV3DefaultTopicConfigs

> TopicConfigDataList ListKafkaV3DefaultTopicConfigs(ctx, clusterId, topicName)

List Default Topic Configs

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

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

