# \StatusConnectV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListConnectv1ConnectorTasks**](StatusConnectV1Api.md#ListConnectv1ConnectorTasks) | **Get** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors/{connector_name}/tasks | List of Connector Tasks
[**ReadConnectv1ConnectorStatus**](StatusConnectV1Api.md#ReadConnectv1ConnectorStatus) | **Get** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors/{connector_name}/status | Read a Connector Status



## ListConnectv1ConnectorTasks

> ConnectV1Connectors ListConnectv1ConnectorTasks(ctx, connectorName, environmentId, kafkaClusterId).Execute()

List of Connector Tasks



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
    resp, r, err := api_client.StatusConnectV1Api.ListConnectv1ConnectorTasks(context.Background(), connectorName, environmentId, kafkaClusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatusConnectV1Api.ListConnectv1ConnectorTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectv1ConnectorTasks`: ConnectV1Connectors
    fmt.Fprintf(os.Stdout, "Response from `StatusConnectV1Api.ListConnectv1ConnectorTasks`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiListConnectv1ConnectorTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ConnectV1Connectors**](ConnectV1Connectors.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadConnectv1ConnectorStatus

> InlineResponse2001 ReadConnectv1ConnectorStatus(ctx, connectorName, environmentId, kafkaClusterId).Execute()

Read a Connector Status



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
    resp, r, err := api_client.StatusConnectV1Api.ReadConnectv1ConnectorStatus(context.Background(), connectorName, environmentId, kafkaClusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatusConnectV1Api.ReadConnectv1ConnectorStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadConnectv1ConnectorStatus`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `StatusConnectV1Api.ReadConnectv1ConnectorStatus`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiReadConnectv1ConnectorStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

