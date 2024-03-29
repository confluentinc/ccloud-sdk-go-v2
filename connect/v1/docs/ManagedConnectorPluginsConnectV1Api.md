# \ManagedConnectorPluginsConnectV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListConnectv1ConnectorPlugins**](ManagedConnectorPluginsConnectV1Api.md#ListConnectv1ConnectorPlugins) | **Get** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connector-plugins | List of Managed Connector plugins
[**ValidateConnectv1ConnectorPlugin**](ManagedConnectorPluginsConnectV1Api.md#ValidateConnectv1ConnectorPlugin) | **Put** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connector-plugins/{plugin_name}/config/validate | Validate a Managed Connector Plugin



## ListConnectv1ConnectorPlugins

> []InlineResponse2002 ListConnectv1ConnectorPlugins(ctx, environmentId, kafkaClusterId).Execute()

List of Managed Connector plugins



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
    environmentId := "environmentId_example" // string | The unique identifier of the environment this resource belongs to.
    kafkaClusterId := "kafkaClusterId_example" // string | The unique identifier for the Kafka cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedConnectorPluginsConnectV1Api.ListConnectv1ConnectorPlugins(context.Background(), environmentId, kafkaClusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedConnectorPluginsConnectV1Api.ListConnectv1ConnectorPlugins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectv1ConnectorPlugins`: []InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `ManagedConnectorPluginsConnectV1Api.ListConnectv1ConnectorPlugins`: %v\n", resp)
}
```

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

[**[]InlineResponse2002**](InlineResponse2002.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateConnectv1ConnectorPlugin

> InlineResponse2003 ValidateConnectv1ConnectorPlugin(ctx, pluginName, environmentId, kafkaClusterId).RequestBody(requestBody).Execute()

Validate a Managed Connector Plugin



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
    pluginName := "pluginName_example" // string | The unique name of the connector plugin.
    environmentId := "environmentId_example" // string | The unique identifier of the environment this resource belongs to.
    kafkaClusterId := "kafkaClusterId_example" // string | The unique identifier for the Kafka cluster.
    requestBody := map[string]string{"key": "Inner_example"} // map[string]string | Configuration parameters for the connector. All values should be strings. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedConnectorPluginsConnectV1Api.ValidateConnectv1ConnectorPlugin(context.Background(), pluginName, environmentId, kafkaClusterId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedConnectorPluginsConnectV1Api.ValidateConnectv1ConnectorPlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateConnectv1ConnectorPlugin`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `ManagedConnectorPluginsConnectV1Api.ValidateConnectv1ConnectorPlugin`: %v\n", resp)
}
```

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



 **requestBody** | **map[string]string** | Configuration parameters for the connector. All values should be strings. | 

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

