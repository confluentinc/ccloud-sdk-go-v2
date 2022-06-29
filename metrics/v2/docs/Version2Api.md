# \Version2Api

All URIs are relative to *https://api.telemetry.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2MetricsDatasetAttributesPost**](Version2Api.md#V2MetricsDatasetAttributesPost) | **Post** /v2/metrics/{dataset}/attributes | Query label values
[**V2MetricsDatasetDescriptorsMetricsGet**](Version2Api.md#V2MetricsDatasetDescriptorsMetricsGet) | **Get** /v2/metrics/{dataset}/descriptors/metrics | List metric descriptors
[**V2MetricsDatasetDescriptorsResourcesGet**](Version2Api.md#V2MetricsDatasetDescriptorsResourcesGet) | **Get** /v2/metrics/{dataset}/descriptors/resources | List resource descriptors
[**V2MetricsDatasetExportGet**](Version2Api.md#V2MetricsDatasetExportGet) | **Get** /v2/metrics/{dataset}/export | Export metric values
[**V2MetricsDatasetQueryPost**](Version2Api.md#V2MetricsDatasetQueryPost) | **Post** /v2/metrics/{dataset}/query | Query metric values



## V2MetricsDatasetAttributesPost

> AttributesResponse V2MetricsDatasetAttributesPost(ctx, dataset).PageToken(pageToken).AttributesRequest(attributesRequest).Execute()

Query label values



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
    dataset := "dataset_example" // string | The dataset to query.
    pageToken := "pageToken_example" // string | The next page token. The token is returned by the previous request as part of `meta.pagination`. (optional)
    attributesRequest := *openapiclient.NewAttributesRequest([]string{"GroupBy_example"}) // AttributesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Version2Api.V2MetricsDatasetAttributesPost(context.Background(), dataset).PageToken(pageToken).AttributesRequest(attributesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Version2Api.V2MetricsDatasetAttributesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2MetricsDatasetAttributesPost`: AttributesResponse
    fmt.Fprintf(os.Stdout, "Response from `Version2Api.V2MetricsDatasetAttributesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataset** | **string** | The dataset to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2MetricsDatasetAttributesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | The next page token. The token is returned by the previous request as part of &#x60;meta.pagination&#x60;. | 
 **attributesRequest** | [**AttributesRequest**](AttributesRequest.md) |  | 

### Return type

[**AttributesResponse**](AttributesResponse.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2MetricsDatasetDescriptorsMetricsGet

> ListMetricDescriptorsResponse V2MetricsDatasetDescriptorsMetricsGet(ctx, dataset).PageSize(pageSize).PageToken(pageToken).ResourceType(resourceType).Execute()

List metric descriptors



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
    dataset := "dataset_example" // string | The dataset to list metric descriptors for. Currently the only supported dataset name is `cloud`. See [here](#section/Object-Model/Datasets).
    pageSize := int32(56) // int32 | The maximum number of results to return. The page size is an integer in the range from 1 through 1000. (optional) (default to 100)
    pageToken := "pageToken_example" // string | The next page token. The token is returned by the previous request as part of `meta.pagination`. (optional)
    resourceType := "resourceType_example" // string | The type of the resource to list metric descriptors for. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Version2Api.V2MetricsDatasetDescriptorsMetricsGet(context.Background(), dataset).PageSize(pageSize).PageToken(pageToken).ResourceType(resourceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Version2Api.V2MetricsDatasetDescriptorsMetricsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2MetricsDatasetDescriptorsMetricsGet`: ListMetricDescriptorsResponse
    fmt.Fprintf(os.Stdout, "Response from `Version2Api.V2MetricsDatasetDescriptorsMetricsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataset** | **string** | The dataset to list metric descriptors for. Currently the only supported dataset name is &#x60;cloud&#x60;. See [here](#section/Object-Model/Datasets). | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2MetricsDatasetDescriptorsMetricsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The maximum number of results to return. The page size is an integer in the range from 1 through 1000. | [default to 100]
 **pageToken** | **string** | The next page token. The token is returned by the previous request as part of &#x60;meta.pagination&#x60;. | 
 **resourceType** | **string** | The type of the resource to list metric descriptors for. | 

### Return type

[**ListMetricDescriptorsResponse**](ListMetricDescriptorsResponse.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2MetricsDatasetDescriptorsResourcesGet

> ListResourceDescriptorsResponse V2MetricsDatasetDescriptorsResourcesGet(ctx, dataset).PageSize(pageSize).PageToken(pageToken).Execute()

List resource descriptors



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
    dataset := "dataset_example" // string | The dataset to list resource descriptors for. Currently the only supported dataset name is `cloud`. See [here](#section/Object-Model/Datasets).
    pageSize := int32(56) // int32 | The maximum number of results to return. The page size is an integer in the range from 1 through 1000. (optional) (default to 100)
    pageToken := "pageToken_example" // string | The next page token. The token is returned by the previous request as part of `meta.pagination`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Version2Api.V2MetricsDatasetDescriptorsResourcesGet(context.Background(), dataset).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Version2Api.V2MetricsDatasetDescriptorsResourcesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2MetricsDatasetDescriptorsResourcesGet`: ListResourceDescriptorsResponse
    fmt.Fprintf(os.Stdout, "Response from `Version2Api.V2MetricsDatasetDescriptorsResourcesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataset** | **string** | The dataset to list resource descriptors for. Currently the only supported dataset name is &#x60;cloud&#x60;. See [here](#section/Object-Model/Datasets). | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2MetricsDatasetDescriptorsResourcesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The maximum number of results to return. The page size is an integer in the range from 1 through 1000. | [default to 100]
 **pageToken** | **string** | The next page token. The token is returned by the previous request as part of &#x60;meta.pagination&#x60;. | 

### Return type

[**ListResourceDescriptorsResponse**](ListResourceDescriptorsResponse.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2MetricsDatasetExportGet

> string V2MetricsDatasetExportGet(ctx, dataset).ResourceKafkaId(resourceKafkaId).ResourceConnectorId(resourceConnectorId).ResourceKsqlId(resourceKsqlId).ResourceSchemaRegistryId(resourceSchemaRegistryId).Metric(metric).Execute()

Export metric values



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
    dataset := "dataset_example" // string | The dataset to export metrics for. Currently the only supported dataset name is `cloud`. See [here](#section/Object-Model/Datasets).
    resourceKafkaId := []string{"Inner_example"} // []string | The ID of the Kafka cluster to export metrics for. This parameter can be specified multiple times (e.g. `?resource.kafka.id=lkc-1&resource.kafka.id=lkc-2`). (optional)
    resourceConnectorId := []string{"Inner_example"} // []string | The ID of the Connector to export metrics for. This parameter can be specified multiple times. (optional)
    resourceKsqlId := []string{"Inner_example"} // []string | The ID of the ksqlDB application to export metrics for. This parameter can be specified multiple times. (optional)
    resourceSchemaRegistryId := []string{"Inner_example"} // []string | The ID of the Schema Registry to export metrics for. This parameter can be specified multiple times. (optional)
    metric := []string{"Inner_example"} // []string | The metric to export. If this parameter is not specified, all metrics for the resource will be exported. This parameter can be specified multiple times. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Version2Api.V2MetricsDatasetExportGet(context.Background(), dataset).ResourceKafkaId(resourceKafkaId).ResourceConnectorId(resourceConnectorId).ResourceKsqlId(resourceKsqlId).ResourceSchemaRegistryId(resourceSchemaRegistryId).Metric(metric).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Version2Api.V2MetricsDatasetExportGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2MetricsDatasetExportGet`: string
    fmt.Fprintf(os.Stdout, "Response from `Version2Api.V2MetricsDatasetExportGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataset** | **string** | The dataset to export metrics for. Currently the only supported dataset name is &#x60;cloud&#x60;. See [here](#section/Object-Model/Datasets). | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2MetricsDatasetExportGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceKafkaId** | **[]string** | The ID of the Kafka cluster to export metrics for. This parameter can be specified multiple times (e.g. &#x60;?resource.kafka.id&#x3D;lkc-1&amp;resource.kafka.id&#x3D;lkc-2&#x60;). | 
 **resourceConnectorId** | **[]string** | The ID of the Connector to export metrics for. This parameter can be specified multiple times. | 
 **resourceKsqlId** | **[]string** | The ID of the ksqlDB application to export metrics for. This parameter can be specified multiple times. | 
 **resourceSchemaRegistryId** | **[]string** | The ID of the Schema Registry to export metrics for. This parameter can be specified multiple times. | 
 **metric** | **[]string** | The metric to export. If this parameter is not specified, all metrics for the resource will be exported. This parameter can be specified multiple times. | 

### Return type

**string**

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain;version=0.0.4, application/openmetrics-text;version=1.0.0;charset=utf-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2MetricsDatasetQueryPost

> QueryResponse V2MetricsDatasetQueryPost(ctx, dataset).PageToken(pageToken).QueryRequest(queryRequest).Execute()

Query metric values



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
    dataset := "dataset_example" // string | The dataset to query. Currently the only supported dataset name is `cloud`. See [here](#section/Object-Model/Datasets).
    pageToken := "pageToken_example" // string | The next page token. The token is returned by the previous request as part of `meta.pagination`. Pagination is only supported for requests containing a `group_by` element. (optional)
    queryRequest := *openapiclient.NewQueryRequest([]openapiclient.Aggregation{*openapiclient.NewAggregation("Metric_example")}, openapiclient.Granularity("PT1M"), []string{"Intervals_example"}) // QueryRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Version2Api.V2MetricsDatasetQueryPost(context.Background(), dataset).PageToken(pageToken).QueryRequest(queryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Version2Api.V2MetricsDatasetQueryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2MetricsDatasetQueryPost`: QueryResponse
    fmt.Fprintf(os.Stdout, "Response from `Version2Api.V2MetricsDatasetQueryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataset** | **string** | The dataset to query. Currently the only supported dataset name is &#x60;cloud&#x60;. See [here](#section/Object-Model/Datasets). | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2MetricsDatasetQueryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | The next page token. The token is returned by the previous request as part of &#x60;meta.pagination&#x60;. Pagination is only supported for requests containing a &#x60;group_by&#x60; element. | 
 **queryRequest** | [**QueryRequest**](QueryRequest.md) |  | 

### Return type

[**QueryResponse**](QueryResponse.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

