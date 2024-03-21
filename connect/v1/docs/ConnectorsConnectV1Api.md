# \ConnectorsConnectV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnectv1Connector**](ConnectorsConnectV1Api.md#CreateConnectv1Connector) | **Post** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors | Create a Connector
[**CreateOrUpdateConnectv1ConnectorConfig**](ConnectorsConnectV1Api.md#CreateOrUpdateConnectv1ConnectorConfig) | **Put** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors/{connector_name}/config | Create or Update a Connector Configuration
[**DeleteConnectv1Connector**](ConnectorsConnectV1Api.md#DeleteConnectv1Connector) | **Delete** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors/{connector_name} | Delete a Connector
[**GetConnectv1ConnectorConfig**](ConnectorsConnectV1Api.md#GetConnectv1ConnectorConfig) | **Get** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors/{connector_name}/config | Read a Connector Configuration
[**ListConnectv1Connectors**](ConnectorsConnectV1Api.md#ListConnectv1Connectors) | **Get** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors | List of Connectors
[**ListConnectv1ConnectorsWithExpansions**](ConnectorsConnectV1Api.md#ListConnectv1ConnectorsWithExpansions) | **Get** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors?expand&#x3D;info,status,id | List of Connectors with Expansions
[**ReadConnectv1Connector**](ConnectorsConnectV1Api.md#ReadConnectv1Connector) | **Get** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors/{connector_name} | Read a Connector



## CreateConnectv1Connector

> ConnectV1Connector CreateConnectv1Connector(ctx, environmentId, kafkaClusterId).InlineObject(inlineObject).Execute()

Create a Connector



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
    inlineObject := *openapiclient.NewInlineObject() // InlineObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorsConnectV1Api.CreateConnectv1Connector(context.Background(), environmentId, kafkaClusterId).InlineObject(inlineObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsConnectV1Api.CreateConnectv1Connector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConnectv1Connector`: ConnectV1Connector
    fmt.Fprintf(os.Stdout, "Response from `ConnectorsConnectV1Api.CreateConnectv1Connector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The unique identifier of the environment this resource belongs to. | 
**kafkaClusterId** | **string** | The unique identifier for the Kafka cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectv1ConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject** | [**InlineObject**](InlineObject.md) |  | 

### Return type

[**ConnectV1Connector**](ConnectV1Connector.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateConnectv1ConnectorConfig

> ConnectV1Connector CreateOrUpdateConnectv1ConnectorConfig(ctx, connectorName, environmentId, kafkaClusterId).ModelMap(modelMap).Execute()

Create or Update a Connector Configuration



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
    modelMap := TODO // ModelMap | Configuration parameters for the connector. All values should be strings. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorsConnectV1Api.CreateOrUpdateConnectv1ConnectorConfig(context.Background(), connectorName, environmentId, kafkaClusterId).ModelMap(modelMap).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsConnectV1Api.CreateOrUpdateConnectv1ConnectorConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateConnectv1ConnectorConfig`: ConnectV1Connector
    fmt.Fprintf(os.Stdout, "Response from `ConnectorsConnectV1Api.CreateOrUpdateConnectv1ConnectorConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCreateOrUpdateConnectv1ConnectorConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **modelMap** | [**ModelMap**](ModelMap.md) | Configuration parameters for the connector. All values should be strings. | 

### Return type

[**ConnectV1Connector**](ConnectV1Connector.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnectv1Connector

> InlineResponse200 DeleteConnectv1Connector(ctx, connectorName, environmentId, kafkaClusterId).Execute()

Delete a Connector



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
    resp, r, err := api_client.ConnectorsConnectV1Api.DeleteConnectv1Connector(context.Background(), connectorName, environmentId, kafkaClusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsConnectV1Api.DeleteConnectv1Connector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteConnectv1Connector`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `ConnectorsConnectV1Api.DeleteConnectv1Connector`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDeleteConnectv1ConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectv1ConnectorConfig

> map[string]string GetConnectv1ConnectorConfig(ctx, connectorName, environmentId, kafkaClusterId).Execute()

Read a Connector Configuration



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
    resp, r, err := api_client.ConnectorsConnectV1Api.GetConnectv1ConnectorConfig(context.Background(), connectorName, environmentId, kafkaClusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsConnectV1Api.GetConnectv1ConnectorConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectv1ConnectorConfig`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `ConnectorsConnectV1Api.GetConnectv1ConnectorConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetConnectv1ConnectorConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]string**

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectv1Connectors

> []string ListConnectv1Connectors(ctx, environmentId, kafkaClusterId).Execute()

List of Connectors



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
    resp, r, err := api_client.ConnectorsConnectV1Api.ListConnectv1Connectors(context.Background(), environmentId, kafkaClusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsConnectV1Api.ListConnectv1Connectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectv1Connectors`: []string
    fmt.Fprintf(os.Stdout, "Response from `ConnectorsConnectV1Api.ListConnectv1Connectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The unique identifier of the environment this resource belongs to. | 
**kafkaClusterId** | **string** | The unique identifier for the Kafka cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectv1ConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**[]string**

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectv1ConnectorsWithExpansions

> ConnectV1ConnectorExpansionMap ListConnectv1ConnectorsWithExpansions(ctx, environmentId, kafkaClusterId).Expand(expand).Execute()

List of Connectors with Expansions



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
    expand := "expand_example" // string | - id : Returns metadata of each connector such as id and id type. - info : Returns metadata of each connector such as the configuration, task information, and type of connector. - status : Returns additional state information of each connector including their status and tasks. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorsConnectV1Api.ListConnectv1ConnectorsWithExpansions(context.Background(), environmentId, kafkaClusterId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsConnectV1Api.ListConnectv1ConnectorsWithExpansions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectv1ConnectorsWithExpansions`: ConnectV1ConnectorExpansionMap
    fmt.Fprintf(os.Stdout, "Response from `ConnectorsConnectV1Api.ListConnectv1ConnectorsWithExpansions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The unique identifier of the environment this resource belongs to. | 
**kafkaClusterId** | **string** | The unique identifier for the Kafka cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectv1ConnectorsWithExpansionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** | - id : Returns metadata of each connector such as id and id type. - info : Returns metadata of each connector such as the configuration, task information, and type of connector. - status : Returns additional state information of each connector including their status and tasks. | 

### Return type

[**ConnectV1ConnectorExpansionMap**](ConnectV1ConnectorExpansionMap.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadConnectv1Connector

> ConnectV1Connector ReadConnectv1Connector(ctx, connectorName, environmentId, kafkaClusterId).Execute()

Read a Connector



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
    resp, r, err := api_client.ConnectorsConnectV1Api.ReadConnectv1Connector(context.Background(), connectorName, environmentId, kafkaClusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsConnectV1Api.ReadConnectv1Connector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadConnectv1Connector`: ConnectV1Connector
    fmt.Fprintf(os.Stdout, "Response from `ConnectorsConnectV1Api.ReadConnectv1Connector`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiReadConnectv1ConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ConnectV1Connector**](ConnectV1Connector.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

