# \PluginsV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListConnectv1ConnectorPlugins**](PluginsV1Api.md#ListConnectv1ConnectorPlugins) | **Get** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connector-plugins | List of Connector Plugins
[**ValidateConnectv1ConnectorPlugin**](PluginsV1Api.md#ValidateConnectv1ConnectorPlugin) | **Put** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connector-plugins/{plugin_name}/config/validate | Validate a Connector Plugin



## ListConnectv1ConnectorPlugins

> []InlineResponse2002 ListConnectv1ConnectorPlugins(ctx, environmentId, kafkaClusterId).Execute()

List of Connector Plugins



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The unique identifier of the environment this resource belongs to. | 
**kafkaClusterId** | **string** | The unique identifier for the Kafka cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectv1ConnectorPluginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]InlineResponse2002**](inline_response_200_2.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateConnectv1ConnectorPlugin

> InlineResponse2003 ValidateConnectv1ConnectorPlugin(ctx, pluginName, environmentId, kafkaClusterId).RequestBody(requestBody).Execute()

Validate a Connector Plugin



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginName** | **string** | The unique name of the connector plugin. | 
**environmentId** | **string** | The unique identifier of the environment this resource belongs to. | 
**kafkaClusterId** | **string** | The unique identifier for the Kafka cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateConnectv1ConnectorPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **requestBody** | [**map[string]string**](string.md) | Configuration parameters for the connector. All values should be strings. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

