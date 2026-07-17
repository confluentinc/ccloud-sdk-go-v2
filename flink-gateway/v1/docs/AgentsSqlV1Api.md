# \AgentsSqlV1Api

All URIs are relative to *https://flink.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSqlv1Agent**](AgentsSqlV1Api.md#CreateSqlv1Agent) | **Post** /sql/v1/organizations/{organization_id}/environments/{environment_id}/databases/{kafka_cluster_id}/agents | Create an Agent
[**DeleteSqlv1Agent**](AgentsSqlV1Api.md#DeleteSqlv1Agent) | **Delete** /sql/v1/organizations/{organization_id}/environments/{environment_id}/databases/{kafka_cluster_id}/agents/{agent_name} | Delete an Agent
[**GetSqlv1Agent**](AgentsSqlV1Api.md#GetSqlv1Agent) | **Get** /sql/v1/organizations/{organization_id}/environments/{environment_id}/databases/{kafka_cluster_id}/agents/{agent_name} | Read an Agent
[**ListSqlv1Agents**](AgentsSqlV1Api.md#ListSqlv1Agents) | **Get** /sql/v1/organizations/{organization_id}/environments/{environment_id}/agents | List all agents
[**UpdateSqlv1Agent**](AgentsSqlV1Api.md#UpdateSqlv1Agent) | **Put** /sql/v1/organizations/{organization_id}/environments/{environment_id}/databases/{kafka_cluster_id}/agents/{agent_name} | Alter an Agent



## CreateSqlv1Agent

> SqlV1Agent CreateSqlv1Agent(ctx, organizationId, environmentId, kafkaClusterId).SqlV1Agent(sqlV1Agent).Execute()

Create an Agent



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
    organizationId := TODO // string | The unique identifier for the organization
    environmentId := "environmentId_example" // string | The unique identifier for the environment.
    kafkaClusterId := "kafkaClusterId_example" // string | The unique identifier for the database.
    sqlV1Agent := *openapiclient.NewSqlV1Agent("sql/v1", "Kind_example", *openapiclient.NewObjectMeta("https://flink.us-west1.aws.confluent.cloud/sql/v1/environments/env-123/statements/my-statement"), "chat-listener-agent", "OrganizationId_example", "EnvironmentId_example", *openapiclient.NewSqlV1AgentSpec()) // SqlV1Agent |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgentsSqlV1Api.CreateSqlv1Agent(context.Background(), organizationId, environmentId, kafkaClusterId).SqlV1Agent(sqlV1Agent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentsSqlV1Api.CreateSqlv1Agent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSqlv1Agent`: SqlV1Agent
    fmt.Fprintf(os.Stdout, "Response from `AgentsSqlV1Api.CreateSqlv1Agent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization | 
**environmentId** | **string** | The unique identifier for the environment. | 
**kafkaClusterId** | **string** | The unique identifier for the database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSqlv1AgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sqlV1Agent** | [**SqlV1Agent**](SqlV1Agent.md) |  | 

### Return type

[**SqlV1Agent**](SqlV1Agent.md)

### Authorization

[external-access-token](../README.md#external-access-token), [global-api-key](../README.md#global-api-key), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSqlv1Agent

> DeleteSqlv1Agent(ctx, organizationId, environmentId, kafkaClusterId, agentName).Execute()

Delete an Agent



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
    organizationId := TODO // string | The unique identifier for the organization
    environmentId := "environmentId_example" // string | The unique identifier for the environment.
    kafkaClusterId := "kafkaClusterId_example" // string | The unique identifier for the database.
    agentName := "agentName_example" // string | The unique identifier for the Agent

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgentsSqlV1Api.DeleteSqlv1Agent(context.Background(), organizationId, environmentId, kafkaClusterId, agentName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentsSqlV1Api.DeleteSqlv1Agent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization | 
**environmentId** | **string** | The unique identifier for the environment. | 
**kafkaClusterId** | **string** | The unique identifier for the database. | 
**agentName** | **string** | The unique identifier for the Agent | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSqlv1AgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[external-access-token](../README.md#external-access-token), [global-api-key](../README.md#global-api-key), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSqlv1Agent

> SqlV1Agent GetSqlv1Agent(ctx, organizationId, environmentId, kafkaClusterId, agentName).Execute()

Read an Agent



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
    organizationId := TODO // string | The unique identifier for the organization
    environmentId := "environmentId_example" // string | The unique identifier for the environment.
    kafkaClusterId := "kafkaClusterId_example" // string | The unique identifier for the database.
    agentName := "agentName_example" // string | The unique identifier for the Agent

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgentsSqlV1Api.GetSqlv1Agent(context.Background(), organizationId, environmentId, kafkaClusterId, agentName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentsSqlV1Api.GetSqlv1Agent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSqlv1Agent`: SqlV1Agent
    fmt.Fprintf(os.Stdout, "Response from `AgentsSqlV1Api.GetSqlv1Agent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization | 
**environmentId** | **string** | The unique identifier for the environment. | 
**kafkaClusterId** | **string** | The unique identifier for the database. | 
**agentName** | **string** | The unique identifier for the Agent | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSqlv1AgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**SqlV1Agent**](SqlV1Agent.md)

### Authorization

[external-access-token](../README.md#external-access-token), [global-api-key](../README.md#global-api-key), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSqlv1Agents

> SqlV1AgentList ListSqlv1Agents(ctx, organizationId, environmentId).PageSize(pageSize).PageToken(pageToken).Execute()

List all agents



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
    organizationId := TODO // string | The unique identifier for the organization.
    environmentId := "environmentId_example" // string | The unique identifier for the environment.
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgentsSqlV1Api.ListSqlv1Agents(context.Background(), organizationId, environmentId).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentsSqlV1Api.ListSqlv1Agents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSqlv1Agents`: SqlV1AgentList
    fmt.Fprintf(os.Stdout, "Response from `AgentsSqlV1Api.ListSqlv1Agents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSqlv1AgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**SqlV1AgentList**](SqlV1AgentList.md)

### Authorization

[external-access-token](../README.md#external-access-token), [global-api-key](../README.md#global-api-key), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSqlv1Agent

> SqlV1Agent UpdateSqlv1Agent(ctx, organizationId, environmentId, kafkaClusterId, agentName).SqlV1Agent(sqlV1Agent).Execute()

Alter an Agent



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
    organizationId := TODO // string | The unique identifier for the organization
    environmentId := "environmentId_example" // string | The unique identifier for the environment.
    kafkaClusterId := "kafkaClusterId_example" // string | The unique identifier for the database.
    agentName := "agentName_example" // string | The unique identifier for the Agent
    sqlV1Agent := *openapiclient.NewSqlV1Agent("sql/v1", "Kind_example", *openapiclient.NewObjectMeta("https://flink.us-west1.aws.confluent.cloud/sql/v1/environments/env-123/statements/my-statement"), "chat-listener-agent", "OrganizationId_example", "EnvironmentId_example", *openapiclient.NewSqlV1AgentSpec()) // SqlV1Agent | The Agent resource with updated spec fields.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgentsSqlV1Api.UpdateSqlv1Agent(context.Background(), organizationId, environmentId, kafkaClusterId, agentName).SqlV1Agent(sqlV1Agent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentsSqlV1Api.UpdateSqlv1Agent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSqlv1Agent`: SqlV1Agent
    fmt.Fprintf(os.Stdout, "Response from `AgentsSqlV1Api.UpdateSqlv1Agent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization | 
**environmentId** | **string** | The unique identifier for the environment. | 
**kafkaClusterId** | **string** | The unique identifier for the database. | 
**agentName** | **string** | The unique identifier for the Agent | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSqlv1AgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **sqlV1Agent** | [**SqlV1Agent**](SqlV1Agent.md) | The Agent resource with updated spec fields. | 

### Return type

[**SqlV1Agent**](SqlV1Agent.md)

### Authorization

[external-access-token](../README.md#external-access-token), [global-api-key](../README.md#global-api-key), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

