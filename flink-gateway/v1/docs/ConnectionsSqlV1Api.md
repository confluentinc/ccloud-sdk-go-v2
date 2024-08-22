# \ConnectionsSqlV1Api

All URIs are relative to *https://flink.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSqlv1Connection**](ConnectionsSqlV1Api.md#CreateSqlv1Connection) | **Post** /sql/v1/organizations/{organization_id}/environments/{environment_id}/connections | Create a Connection
[**DeleteSqlv1Connection**](ConnectionsSqlV1Api.md#DeleteSqlv1Connection) | **Delete** /sql/v1/organizations/{organization_id}/environments/{environment_id}/connections/{connection_name} | Delete a Connection
[**GetSqlv1Connection**](ConnectionsSqlV1Api.md#GetSqlv1Connection) | **Get** /sql/v1/organizations/{organization_id}/environments/{environment_id}/connections/{connection_name} | Read a Connection
[**ListSqlv1Connections**](ConnectionsSqlV1Api.md#ListSqlv1Connections) | **Get** /sql/v1/organizations/{organization_id}/environments/{environment_id}/connections | List of Connections
[**UpdateSqlv1Connection**](ConnectionsSqlV1Api.md#UpdateSqlv1Connection) | **Put** /sql/v1/organizations/{organization_id}/environments/{environment_id}/connections/{connection_name} | Update a Connection



## CreateSqlv1Connection

> SqlV1Connection CreateSqlv1Connection(ctx, organizationId, environmentId).SqlV1Connection(sqlV1Connection).Execute()

Create a Connection



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
    sqlV1Connection := *openapiclient.NewSqlV1Connection() // SqlV1Connection |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsSqlV1Api.CreateSqlv1Connection(context.Background(), organizationId, environmentId).SqlV1Connection(sqlV1Connection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsSqlV1Api.CreateSqlv1Connection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSqlv1Connection`: SqlV1Connection
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsSqlV1Api.CreateSqlv1Connection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSqlv1ConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sqlV1Connection** | [**SqlV1Connection**](SqlV1Connection.md) |  | 

### Return type

[**SqlV1Connection**](sql.v1.Connection.md)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSqlv1Connection

> DeleteSqlv1Connection(ctx, organizationId, environmentId, connectionName).Execute()

Delete a Connection



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
    connectionName := "connectionName_example" // string | The unique identifier for the connection.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsSqlV1Api.DeleteSqlv1Connection(context.Background(), organizationId, environmentId, connectionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsSqlV1Api.DeleteSqlv1Connection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 
**connectionName** | **string** | The unique identifier for the connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSqlv1ConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSqlv1Connection

> SqlV1Connection GetSqlv1Connection(ctx, organizationId, environmentId, connectionName).Execute()

Read a Connection



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
    connectionName := "connectionName_example" // string | The user provided name of the Connection. Unique within a region within an org and env.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsSqlV1Api.GetSqlv1Connection(context.Background(), organizationId, environmentId, connectionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsSqlV1Api.GetSqlv1Connection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSqlv1Connection`: SqlV1Connection
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsSqlV1Api.GetSqlv1Connection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 
**connectionName** | **string** | The user provided name of the Connection. Unique within a region within an org and env. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSqlv1ConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SqlV1Connection**](sql.v1.Connection.md)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSqlv1Connections

> SqlV1ConnectionList ListSqlv1Connections(ctx, organizationId, environmentId).SpecConnectionType(specConnectionType).PageSize(pageSize).PageToken(pageToken).Execute()

List of Connections



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
    specConnectionType := "specConnectionType_example" // string | Filter the results by exact match for spec.connection_type (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsSqlV1Api.ListSqlv1Connections(context.Background(), organizationId, environmentId).SpecConnectionType(specConnectionType).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsSqlV1Api.ListSqlv1Connections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSqlv1Connections`: SqlV1ConnectionList
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsSqlV1Api.ListSqlv1Connections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSqlv1ConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **specConnectionType** | **string** | Filter the results by exact match for spec.connection_type | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**SqlV1ConnectionList**](SqlV1ConnectionList.md)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSqlv1Connection

> UpdateSqlv1Connection(ctx, organizationId, environmentId, connectionName).SqlV1Connection(sqlV1Connection).Execute()

Update a Connection



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
    connectionName := "connectionName_example" // string | The unique identifier for the connection.
    sqlV1Connection := *openapiclient.NewSqlV1Connection() // SqlV1Connection |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsSqlV1Api.UpdateSqlv1Connection(context.Background(), organizationId, environmentId, connectionName).SqlV1Connection(sqlV1Connection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsSqlV1Api.UpdateSqlv1Connection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 
**connectionName** | **string** | The unique identifier for the connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSqlv1ConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sqlV1Connection** | [**SqlV1Connection**](SqlV1Connection.md) |  | 

### Return type

 (empty response body)

### Authorization

[resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

