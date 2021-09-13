# \PartitionV3Api

All URIs are relative to *https://pkc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKafkaV3ConsumerLag**](PartitionV3Api.md#GetKafkaV3ConsumerLag) | **Get** /kafka/v3/clusters/{cluster_id}/consumer-groups/{consumer_group_id}/lags/{topic_name}/partitions/{partition_id} | Get Consumer Lag
[**GetKafkaV3Partition**](PartitionV3Api.md#GetKafkaV3Partition) | **Get** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/partitions/{partition_id} | Get Partition
[**ListKafkaV3Partitions**](PartitionV3Api.md#ListKafkaV3Partitions) | **Get** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/partitions | List Partitions



## GetKafkaV3ConsumerLag

> ConsumerLagData GetKafkaV3ConsumerLag(ctx, clusterId, consumerGroupId, topicName, partitionId)

Get Consumer Lag

[![Early Access](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy) [![Request Access To Cluster Admin for Kafka (v3)](https://img.shields.io/badge/-Request%20Access%20To%20Cluster%20Admin%20For%20Kafka%20v3-%23bc8540)](mailto:ccloud-rest-api+consumer-lag-earlyaccess@confluent.io?subject=Request%20to%20join%20v3%20API%20Early%20Access&body=I%E2%80%99d%20like%20to%20join%20the%20Confluent%20Cluster%20Admin%20For%20Kafka%20v3%20Early%20Access%20to%20provide%20early%20feedback%20on%20consumer%20lag%20apis%21%20My%20Cloud%20Organization%20ID%20is%20%3Cretrieve%20from%20https%3A//confluent.cloud/settings/billing/payment%3E.)  Returns the consumer lag on a partition with the given `partition_id`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**consumerGroupId** | **string**| The consumer group ID. | 
**topicName** | **string**| The topic name. | 
**partitionId** | **int32**| The partition ID. | 

### Return type

[**ConsumerLagData**](ConsumerLagData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaV3Partition

> PartitionData GetKafkaV3Partition(ctx, clusterId, topicName, partitionId)

Get Partition

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns the partition with the given `partition_id`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 
**partitionId** | **int32**| The partition ID. | 

### Return type

[**PartitionData**](PartitionData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaV3Partitions

> PartitionDataList ListKafkaV3Partitions(ctx, clusterId, topicName)

List Partitions

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns the list of partitions that belong to the specified topic.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 

### Return type

[**PartitionDataList**](PartitionDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

