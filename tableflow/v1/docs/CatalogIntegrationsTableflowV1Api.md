# \CatalogIntegrationsTableflowV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTableflowV1CatalogIntegration**](CatalogIntegrationsTableflowV1Api.md#CreateTableflowV1CatalogIntegration) | **Post** /tableflow/v1/catalog-integrations | Create a Catalog Integration
[**DeleteTableflowV1CatalogIntegration**](CatalogIntegrationsTableflowV1Api.md#DeleteTableflowV1CatalogIntegration) | **Delete** /tableflow/v1/catalog-integrations/{id} | Delete a Catalog Integration
[**GetTableflowV1CatalogIntegration**](CatalogIntegrationsTableflowV1Api.md#GetTableflowV1CatalogIntegration) | **Get** /tableflow/v1/catalog-integrations/{id} | Read a Catalog Integration
[**ListTableflowV1CatalogIntegrations**](CatalogIntegrationsTableflowV1Api.md#ListTableflowV1CatalogIntegrations) | **Get** /tableflow/v1/catalog-integrations | List of Catalog Integrations
[**UpdateTableflowV1CatalogIntegration**](CatalogIntegrationsTableflowV1Api.md#UpdateTableflowV1CatalogIntegration) | **Patch** /tableflow/v1/catalog-integrations/{id} | Update a Catalog Integration



## CreateTableflowV1CatalogIntegration

> TableflowV1CatalogIntegration CreateTableflowV1CatalogIntegration(ctx).TableflowV1CatalogIntegration(tableflowV1CatalogIntegration).Execute()

Create a Catalog Integration



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
    tableflowV1CatalogIntegration := *openapiclient.NewTableflowV1CatalogIntegration() // TableflowV1CatalogIntegration |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogIntegrationsTableflowV1Api.CreateTableflowV1CatalogIntegration(context.Background()).TableflowV1CatalogIntegration(tableflowV1CatalogIntegration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogIntegrationsTableflowV1Api.CreateTableflowV1CatalogIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTableflowV1CatalogIntegration`: TableflowV1CatalogIntegration
    fmt.Fprintf(os.Stdout, "Response from `CatalogIntegrationsTableflowV1Api.CreateTableflowV1CatalogIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTableflowV1CatalogIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tableflowV1CatalogIntegration** | [**TableflowV1CatalogIntegration**](TableflowV1CatalogIntegration.md) |  | 

### Return type

[**TableflowV1CatalogIntegration**](tableflow.v1.CatalogIntegration.md)

### Authorization

[confluent-sts-access-token](../README.md#confluent-sts-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTableflowV1CatalogIntegration

> DeleteTableflowV1CatalogIntegration(ctx, id).Environment(environment).SpecKafkaCluster(specKafkaCluster).Execute()

Delete a Catalog Integration



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
    environment := "env-00000" // string | Scope the operation to the given environment.
    specKafkaCluster := "lkc-00000" // string | Scope the operation to the given spec.kafka_cluster.
    id := "id_example" // string | The unique identifier for the catalog integration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogIntegrationsTableflowV1Api.DeleteTableflowV1CatalogIntegration(context.Background(), id).Environment(environment).SpecKafkaCluster(specKafkaCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogIntegrationsTableflowV1Api.DeleteTableflowV1CatalogIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the catalog integration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTableflowV1CatalogIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 
 **specKafkaCluster** | **string** | Scope the operation to the given spec.kafka_cluster. | 


### Return type

 (empty response body)

### Authorization

[confluent-sts-access-token](../README.md#confluent-sts-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTableflowV1CatalogIntegration

> TableflowV1CatalogIntegration GetTableflowV1CatalogIntegration(ctx, id).Environment(environment).SpecKafkaCluster(specKafkaCluster).Execute()

Read a Catalog Integration



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
    environment := "env-00000" // string | Scope the operation to the given environment.
    specKafkaCluster := "lkc-00000" // string | Scope the operation to the given spec.kafka_cluster.
    id := "id_example" // string | The unique identifier for the catalog integration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogIntegrationsTableflowV1Api.GetTableflowV1CatalogIntegration(context.Background(), id).Environment(environment).SpecKafkaCluster(specKafkaCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogIntegrationsTableflowV1Api.GetTableflowV1CatalogIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTableflowV1CatalogIntegration`: TableflowV1CatalogIntegration
    fmt.Fprintf(os.Stdout, "Response from `CatalogIntegrationsTableflowV1Api.GetTableflowV1CatalogIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the catalog integration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTableflowV1CatalogIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 
 **specKafkaCluster** | **string** | Scope the operation to the given spec.kafka_cluster. | 


### Return type

[**TableflowV1CatalogIntegration**](tableflow.v1.CatalogIntegration.md)

### Authorization

[confluent-sts-access-token](../README.md#confluent-sts-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTableflowV1CatalogIntegrations

> TableflowV1CatalogIntegrationList ListTableflowV1CatalogIntegrations(ctx).Environment(environment).SpecKafkaCluster(specKafkaCluster).PageSize(pageSize).PageToken(pageToken).Execute()

List of Catalog Integrations



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
    environment := "env-00000" // string | Filter the results by exact match for environment.
    specKafkaCluster := "lkc-00000" // string | Filter the results by exact match for spec.kafka_cluster.
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogIntegrationsTableflowV1Api.ListTableflowV1CatalogIntegrations(context.Background()).Environment(environment).SpecKafkaCluster(specKafkaCluster).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogIntegrationsTableflowV1Api.ListTableflowV1CatalogIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTableflowV1CatalogIntegrations`: TableflowV1CatalogIntegrationList
    fmt.Fprintf(os.Stdout, "Response from `CatalogIntegrationsTableflowV1Api.ListTableflowV1CatalogIntegrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTableflowV1CatalogIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **specKafkaCluster** | **string** | Filter the results by exact match for spec.kafka_cluster. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**TableflowV1CatalogIntegrationList**](tableflow.v1.CatalogIntegrationList.md)

### Authorization

[confluent-sts-access-token](../README.md#confluent-sts-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTableflowV1CatalogIntegration

> TableflowV1CatalogIntegration UpdateTableflowV1CatalogIntegration(ctx, id).TableflowV1CatalogIntegrationUpdateRequest(tableflowV1CatalogIntegrationUpdateRequest).Execute()

Update a Catalog Integration



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
    id := "id_example" // string | The unique identifier for the catalog integration.
    tableflowV1CatalogIntegrationUpdateRequest := *openapiclient.NewTableflowV1CatalogIntegrationUpdateRequest() // TableflowV1CatalogIntegrationUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogIntegrationsTableflowV1Api.UpdateTableflowV1CatalogIntegration(context.Background(), id).TableflowV1CatalogIntegrationUpdateRequest(tableflowV1CatalogIntegrationUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogIntegrationsTableflowV1Api.UpdateTableflowV1CatalogIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTableflowV1CatalogIntegration`: TableflowV1CatalogIntegration
    fmt.Fprintf(os.Stdout, "Response from `CatalogIntegrationsTableflowV1Api.UpdateTableflowV1CatalogIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the catalog integration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTableflowV1CatalogIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tableflowV1CatalogIntegrationUpdateRequest** | [**TableflowV1CatalogIntegrationUpdateRequest**](TableflowV1CatalogIntegrationUpdateRequest.md) |  | 

### Return type

[**TableflowV1CatalogIntegration**](tableflow.v1.CatalogIntegration.md)

### Authorization

[confluent-sts-access-token](../README.md#confluent-sts-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

