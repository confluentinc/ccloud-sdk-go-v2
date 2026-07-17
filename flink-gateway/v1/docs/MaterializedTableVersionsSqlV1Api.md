# \MaterializedTableVersionsSqlV1Api

All URIs are relative to *https://flink.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSqlv1MaterializedTableVersion**](MaterializedTableVersionsSqlV1Api.md#GetSqlv1MaterializedTableVersion) | **Get** /sql/v1/organizations/{organization_id}/environments/{environment_id}/databases/{kafka_cluster_id}/materialized-tables/{table_name}/versions/{version} | Read a materialized table version
[**ListSqlv1MaterializedTableVersions**](MaterializedTableVersionsSqlV1Api.md#ListSqlv1MaterializedTableVersions) | **Get** /sql/v1/organizations/{organization_id}/environments/{environment_id}/databases/{kafka_cluster_id}/materialized-tables/{table_name}/versions | List all the versions of a materialized table



## GetSqlv1MaterializedTableVersion

> SqlV1MaterializedTableVersion GetSqlv1MaterializedTableVersion(ctx, organizationId, environmentId, kafkaClusterId, tableName, version).Execute()

Read a materialized table version



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
    kafkaClusterId := "kafkaClusterId_example" // string | The unique identifier for the database.
    tableName := "tableName_example" // string | The unique identifier for the Materialized Table.
    version := int32(56) // int32 | The version number of the Materialized Table.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MaterializedTableVersionsSqlV1Api.GetSqlv1MaterializedTableVersion(context.Background(), organizationId, environmentId, kafkaClusterId, tableName, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaterializedTableVersionsSqlV1Api.GetSqlv1MaterializedTableVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSqlv1MaterializedTableVersion`: SqlV1MaterializedTableVersion
    fmt.Fprintf(os.Stdout, "Response from `MaterializedTableVersionsSqlV1Api.GetSqlv1MaterializedTableVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 
**kafkaClusterId** | **string** | The unique identifier for the database. | 
**tableName** | **string** | The unique identifier for the Materialized Table. | 
**version** | **int32** | The version number of the Materialized Table. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSqlv1MaterializedTableVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**SqlV1MaterializedTableVersion**](SqlV1MaterializedTableVersion.md)

### Authorization

[external-access-token](../README.md#external-access-token), [global-api-key](../README.md#global-api-key), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSqlv1MaterializedTableVersions

> SqlV1MaterializedTableVersionList ListSqlv1MaterializedTableVersions(ctx, organizationId, environmentId, kafkaClusterId, tableName).PageSize(pageSize).PageToken(pageToken).Execute()

List all the versions of a materialized table



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
    kafkaClusterId := "kafkaClusterId_example" // string | The unique identifier for the database.
    tableName := "tableName_example" // string | The unique identifier for the Materialized Table.
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MaterializedTableVersionsSqlV1Api.ListSqlv1MaterializedTableVersions(context.Background(), organizationId, environmentId, kafkaClusterId, tableName).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaterializedTableVersionsSqlV1Api.ListSqlv1MaterializedTableVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSqlv1MaterializedTableVersions`: SqlV1MaterializedTableVersionList
    fmt.Fprintf(os.Stdout, "Response from `MaterializedTableVersionsSqlV1Api.ListSqlv1MaterializedTableVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | The unique identifier for the organization. | 
**environmentId** | **string** | The unique identifier for the environment. | 
**kafkaClusterId** | **string** | The unique identifier for the database. | 
**tableName** | **string** | The unique identifier for the Materialized Table. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSqlv1MaterializedTableVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**SqlV1MaterializedTableVersionList**](SqlV1MaterializedTableVersionList.md)

### Authorization

[external-access-token](../README.md#external-access-token), [global-api-key](../README.md#global-api-key), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

