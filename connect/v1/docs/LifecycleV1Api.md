# \LifecycleV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PauseConnectv1Connector**](LifecycleV1Api.md#PauseConnectv1Connector) | **Put** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors/{connector_name}/pause | Pause a Connector
[**ResumeConnectv1Connector**](LifecycleV1Api.md#ResumeConnectv1Connector) | **Put** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors/{connector_name}/resume | Resume a Connector



## PauseConnectv1Connector

> PauseConnectv1Connector(ctx, connectorName, environmentId, kafkaClusterId).Execute()

Pause a Connector



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
    connectorName := "connectorName_example" // string | The unique name of the connector.
    environmentId := "environmentId_example" // string | The unique identifier of the environment this resource belongs to.
    kafkaClusterId := "kafkaClusterId_example" // string | The unique identifier for the Kafka cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LifecycleV1Api.PauseConnectv1Connector(context.Background(), connectorName, environmentId, kafkaClusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LifecycleV1Api.PauseConnectv1Connector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorName** | **string** | The unique name of the connector. | 
**environmentId** | **string** | The unique identifier of the environment this resource belongs to. | 
**kafkaClusterId** | **string** | The unique identifier for the Kafka cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseConnectv1ConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeConnectv1Connector

> ResumeConnectv1Connector(ctx, connectorName, environmentId, kafkaClusterId).Execute()

Resume a Connector



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
    connectorName := "connectorName_example" // string | The unique name of the connector.
    environmentId := "environmentId_example" // string | The unique identifier of the environment this resource belongs to.
    kafkaClusterId := "kafkaClusterId_example" // string | The unique identifier for the Kafka cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LifecycleV1Api.ResumeConnectv1Connector(context.Background(), connectorName, environmentId, kafkaClusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LifecycleV1Api.ResumeConnectv1Connector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorName** | **string** | The unique name of the connector. | 
**environmentId** | **string** | The unique identifier of the environment this resource belongs to. | 
**kafkaClusterId** | **string** | The unique identifier for the Kafka cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeConnectv1ConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

