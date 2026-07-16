# \MaterializedTablesSqlV1Api

All URIs are relative to *https://flink.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSqlv1MaterializedTable**](MaterializedTablesSqlV1Api.md#CreateSqlv1MaterializedTable) | **Post** /sql/v1/organizations/{organization_id}/environments/{environment_id}/databases/{kafka_cluster_id}/materialized-tables | Create a materialized table
[**DeleteSqlv1MaterializedTable**](MaterializedTablesSqlV1Api.md#DeleteSqlv1MaterializedTable) | **Delete** /sql/v1/organizations/{organization_id}/environments/{environment_id}/databases/{kafka_cluster_id}/materialized-tables/{table_name} | Delete a materialized table
[**GetSqlv1MaterializedTable**](MaterializedTablesSqlV1Api.md#GetSqlv1MaterializedTable) | **Get** /sql/v1/organizations/{organization_id}/environments/{environment_id}/databases/{kafka_cluster_id}/materialized-tables/{table_name} | Read a materialized table
[**ListSqlv1MaterializedTables**](MaterializedTablesSqlV1Api.md#ListSqlv1MaterializedTables) | **Get** /sql/v1/organizations/{organization_id}/environments/{environment_id}/materialized-tables | List all materialized tables
[**UpdateSqlv1MaterializedTable**](MaterializedTablesSqlV1Api.md#UpdateSqlv1MaterializedTable) | **Put** /sql/v1/organizations/{organization_id}/environments/{environment_id}/databases/{kafka_cluster_id}/materialized-tables/{table_name} | Update/Evolve a materialized table



## CreateSqlv1MaterializedTable

> SqlV1MaterializedTable CreateSqlv1MaterializedTable(ctx, organizationId, environmentId, kafkaClusterId).SqlV1MaterializedTable(sqlV1MaterializedTable).Execute()

Create a materialized table



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
    sqlV1MaterializedTable := *openapiclient.NewSqlV1MaterializedTable("sql/v1", "Kind_example", *openapiclient.NewObjectMeta("https://flink.us-west1.aws.confluent.cloud/sql/v1/environments/env-123/statements/my-statement"), "high-value-orders", "OrganizationId_example", "EnvironmentId_example", *openapiclient.NewSqlV1MaterializedTableSpec()) // SqlV1MaterializedTable | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MaterializedTablesSqlV1Api.CreateSqlv1MaterializedTable(context.Background(), organizationId, environmentId, kafkaClusterId).SqlV1MaterializedTable(sqlV1MaterializedTable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaterializedTablesSqlV1Api.CreateSqlv1MaterializedTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSqlv1MaterializedTable`: SqlV1MaterializedTable
    fmt.Fprintf(os.Stdout, "Response from `MaterializedTablesSqlV1Api.CreateSqlv1MaterializedTable`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCreateSqlv1MaterializedTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sqlV1MaterializedTable** | [**SqlV1MaterializedTable**](SqlV1MaterializedTable.md) |  | 

### Return type

[**SqlV1MaterializedTable**](SqlV1MaterializedTable.md)

### Authorization

[external-access-token](../README.md#external-access-token), [global-api-key](../README.md#global-api-key), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSqlv1MaterializedTable

> DeleteSqlv1MaterializedTable(ctx, organizationId, environmentId, kafkaClusterId, tableName).Execute()

Delete a materialized table



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
    tableName := "tableName_example" // string | The unique identifier for the Materialized Table

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MaterializedTablesSqlV1Api.DeleteSqlv1MaterializedTable(context.Background(), organizationId, environmentId, kafkaClusterId, tableName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaterializedTablesSqlV1Api.DeleteSqlv1MaterializedTable``: %v\n", err)
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
**tableName** | **string** | The unique identifier for the Materialized Table | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSqlv1MaterializedTableRequest struct via the builder pattern


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


## GetSqlv1MaterializedTable

> SqlV1MaterializedTable GetSqlv1MaterializedTable(ctx, organizationId, environmentId, kafkaClusterId, tableName).Execute()

Read a materialized table



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
    tableName := "tableName_example" // string | The unique identifier for the Materialized Table

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MaterializedTablesSqlV1Api.GetSqlv1MaterializedTable(context.Background(), organizationId, environmentId, kafkaClusterId, tableName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaterializedTablesSqlV1Api.GetSqlv1MaterializedTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSqlv1MaterializedTable`: SqlV1MaterializedTable
    fmt.Fprintf(os.Stdout, "Response from `MaterializedTablesSqlV1Api.GetSqlv1MaterializedTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization | 
**environmentId** | **string** | The unique identifier for the environment. | 
**kafkaClusterId** | **string** | The unique identifier for the database. | 
**tableName** | **string** | The unique identifier for the Materialized Table | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSqlv1MaterializedTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**SqlV1MaterializedTable**](SqlV1MaterializedTable.md)

### Authorization

[external-access-token](../README.md#external-access-token), [global-api-key](../README.md#global-api-key), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSqlv1MaterializedTables

> SqlV1MaterializedTableList ListSqlv1MaterializedTables(ctx, organizationId, environmentId).PageSize(pageSize).PageToken(pageToken).Execute()

List all materialized tables



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
    resp, r, err := api_client.MaterializedTablesSqlV1Api.ListSqlv1MaterializedTables(context.Background(), organizationId, environmentId).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaterializedTablesSqlV1Api.ListSqlv1MaterializedTables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSqlv1MaterializedTables`: SqlV1MaterializedTableList
    fmt.Fprintf(os.Stdout, "Response from `MaterializedTablesSqlV1Api.ListSqlv1MaterializedTables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSqlv1MaterializedTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**SqlV1MaterializedTableList**](SqlV1MaterializedTableList.md)

### Authorization

[external-access-token](../README.md#external-access-token), [global-api-key](../README.md#global-api-key), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSqlv1MaterializedTable

> SqlV1MaterializedTable UpdateSqlv1MaterializedTable(ctx, organizationId, environmentId, kafkaClusterId, tableName).SqlV1MaterializedTable(sqlV1MaterializedTable).Execute()

Update/Evolve a materialized table



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
    tableName := "tableName_example" // string | The unique identifier for the Materialized Table
    sqlV1MaterializedTable := *openapiclient.NewSqlV1MaterializedTable("sql/v1", "Kind_example", *openapiclient.NewObjectMeta("https://flink.us-west1.aws.confluent.cloud/sql/v1/environments/env-123/statements/my-statement"), "high-value-orders", "OrganizationId_example", "EnvironmentId_example", *openapiclient.NewSqlV1MaterializedTableSpec()) // SqlV1MaterializedTable | The Materialized Table resource with updated spec fields.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MaterializedTablesSqlV1Api.UpdateSqlv1MaterializedTable(context.Background(), organizationId, environmentId, kafkaClusterId, tableName).SqlV1MaterializedTable(sqlV1MaterializedTable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaterializedTablesSqlV1Api.UpdateSqlv1MaterializedTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSqlv1MaterializedTable`: SqlV1MaterializedTable
    fmt.Fprintf(os.Stdout, "Response from `MaterializedTablesSqlV1Api.UpdateSqlv1MaterializedTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization | 
**environmentId** | **string** | The unique identifier for the environment. | 
**kafkaClusterId** | **string** | The unique identifier for the database. | 
**tableName** | **string** | The unique identifier for the Materialized Table | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSqlv1MaterializedTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **sqlV1MaterializedTable** | [**SqlV1MaterializedTable**](SqlV1MaterializedTable.md) | The Materialized Table resource with updated spec fields. | 

### Return type

[**SqlV1MaterializedTable**](SqlV1MaterializedTable.md)

### Authorization

[external-access-token](../README.md#external-access-token), [global-api-key](../README.md#global-api-key), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

