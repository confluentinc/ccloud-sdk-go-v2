# \OffsetsConnectV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlterConnectv1ConnectorOffsetsRequest**](OffsetsConnectV1Api.md#AlterConnectv1ConnectorOffsetsRequest) | **Post** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors/{connector_name}/offsets/request | Request a Connector Offsets
[**GetConnectv1ConnectorOffsets**](OffsetsConnectV1Api.md#GetConnectv1ConnectorOffsets) | **Get** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors/{connector_name}/offsets | Get a Connector Offsets
[**GetConnectv1ConnectorOffsetsRequestStatus**](OffsetsConnectV1Api.md#GetConnectv1ConnectorOffsetsRequestStatus) | **Get** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors/{connector_name}/offsets/request/status | Get a Connector Offsets Request Status



## AlterConnectv1ConnectorOffsetsRequest

> ConnectV1AlterOffsetRequestInfo AlterConnectv1ConnectorOffsetsRequest(ctx, connectorName, environmentId, kafkaClusterId).ConnectV1AlterOffsetRequest(connectV1AlterOffsetRequest).Execute()

Request a Connector Offsets



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
    connectV1AlterOffsetRequest := *openapiclient.NewConnectV1AlterOffsetRequest(openapiclient.connect.v1.AlterOffsetRequestType("PATCH")) // ConnectV1AlterOffsetRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OffsetsConnectV1Api.AlterConnectv1ConnectorOffsetsRequest(context.Background(), connectorName, environmentId, kafkaClusterId).ConnectV1AlterOffsetRequest(connectV1AlterOffsetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OffsetsConnectV1Api.AlterConnectv1ConnectorOffsetsRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlterConnectv1ConnectorOffsetsRequest`: ConnectV1AlterOffsetRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `OffsetsConnectV1Api.AlterConnectv1ConnectorOffsetsRequest`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiAlterConnectv1ConnectorOffsetsRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **connectV1AlterOffsetRequest** | [**ConnectV1AlterOffsetRequest**](ConnectV1AlterOffsetRequest.md) |  | 

### Return type

[**ConnectV1AlterOffsetRequestInfo**](ConnectV1AlterOffsetRequestInfo.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectv1ConnectorOffsets

> ConnectV1ConnectorOffsets GetConnectv1ConnectorOffsets(ctx, connectorName, environmentId, kafkaClusterId).Execute()

Get a Connector Offsets



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
    resp, r, err := api_client.OffsetsConnectV1Api.GetConnectv1ConnectorOffsets(context.Background(), connectorName, environmentId, kafkaClusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OffsetsConnectV1Api.GetConnectv1ConnectorOffsets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectv1ConnectorOffsets`: ConnectV1ConnectorOffsets
    fmt.Fprintf(os.Stdout, "Response from `OffsetsConnectV1Api.GetConnectv1ConnectorOffsets`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetConnectv1ConnectorOffsetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ConnectV1ConnectorOffsets**](ConnectV1ConnectorOffsets.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectv1ConnectorOffsetsRequestStatus

> ConnectV1AlterOffsetStatus GetConnectv1ConnectorOffsetsRequestStatus(ctx, connectorName, environmentId, kafkaClusterId).Execute()

Get a Connector Offsets Request Status



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
    resp, r, err := api_client.OffsetsConnectV1Api.GetConnectv1ConnectorOffsetsRequestStatus(context.Background(), connectorName, environmentId, kafkaClusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OffsetsConnectV1Api.GetConnectv1ConnectorOffsetsRequestStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectv1ConnectorOffsetsRequestStatus`: ConnectV1AlterOffsetStatus
    fmt.Fprintf(os.Stdout, "Response from `OffsetsConnectV1Api.GetConnectv1ConnectorOffsetsRequestStatus`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetConnectv1ConnectorOffsetsRequestStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ConnectV1AlterOffsetStatus**](ConnectV1AlterOffsetStatus.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

