# \Version1Api

All URIs are relative to *https://api.telemetry.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1MetricsDatasetAttributesPost**](Version1Api.md#V1MetricsDatasetAttributesPost) | **Post** /v1/metrics/{dataset}/attributes | Query label values
[**V1MetricsDatasetDescriptorsGet**](Version1Api.md#V1MetricsDatasetDescriptorsGet) | **Get** /v1/metrics/{dataset}/descriptors | List all metric descriptors
[**V1MetricsDatasetQueryPost**](Version1Api.md#V1MetricsDatasetQueryPost) | **Post** /v1/metrics/{dataset}/query | Query metric values



## V1MetricsDatasetAttributesPost

> AttributesResponse V1MetricsDatasetAttributesPost(ctx, dataset).PageToken(pageToken).V1MetricsDatasetAttributesPostRequest(v1MetricsDatasetAttributesPostRequest).Execute()

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
    v1MetricsDatasetAttributesPostRequest := *openapiclient.NewV1MetricsDatasetAttributesPostRequest("Metric_example", []string{"GroupBy_example"}) // V1MetricsDatasetAttributesPostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Version1Api.V1MetricsDatasetAttributesPost(context.Background(), dataset).PageToken(pageToken).V1MetricsDatasetAttributesPostRequest(v1MetricsDatasetAttributesPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Version1Api.V1MetricsDatasetAttributesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1MetricsDatasetAttributesPost`: AttributesResponse
    fmt.Fprintf(os.Stdout, "Response from `Version1Api.V1MetricsDatasetAttributesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataset** | **string** | The dataset to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MetricsDatasetAttributesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | The next page token. The token is returned by the previous request as part of &#x60;meta.pagination&#x60;. | 
 **v1MetricsDatasetAttributesPostRequest** | [**V1MetricsDatasetAttributesPostRequest**](V1MetricsDatasetAttributesPostRequest.md) |  | 

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


## V1MetricsDatasetDescriptorsGet

> ListMetricDescriptorsResponse V1MetricsDatasetDescriptorsGet(ctx, dataset).PageSize(pageSize).PageToken(pageToken).Execute()

List all metric descriptors



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Version1Api.V1MetricsDatasetDescriptorsGet(context.Background(), dataset).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Version1Api.V1MetricsDatasetDescriptorsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1MetricsDatasetDescriptorsGet`: ListMetricDescriptorsResponse
    fmt.Fprintf(os.Stdout, "Response from `Version1Api.V1MetricsDatasetDescriptorsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataset** | **string** | The dataset to list metric descriptors for. Currently the only supported dataset name is &#x60;cloud&#x60;. See [here](#section/Object-Model/Datasets). | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MetricsDatasetDescriptorsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The maximum number of results to return. The page size is an integer in the range from 1 through 1000. | [default to 100]
 **pageToken** | **string** | The next page token. The token is returned by the previous request as part of &#x60;meta.pagination&#x60;. | 

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


## V1MetricsDatasetQueryPost

> QueryResponse V1MetricsDatasetQueryPost(ctx, dataset).QueryRequest(queryRequest).Execute()

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
    queryRequest := *openapiclient.NewQueryRequest([]openapiclient.Aggregation{*openapiclient.NewAggregation("Metric_example")}, openapiclient.Granularity("PT1M"), []string{"Intervals_example"}) // QueryRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Version1Api.V1MetricsDatasetQueryPost(context.Background(), dataset).QueryRequest(queryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Version1Api.V1MetricsDatasetQueryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1MetricsDatasetQueryPost`: QueryResponse
    fmt.Fprintf(os.Stdout, "Response from `Version1Api.V1MetricsDatasetQueryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataset** | **string** | The dataset to query. Currently the only supported dataset name is &#x60;cloud&#x60;. See [here](#section/Object-Model/Datasets). | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MetricsDatasetQueryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

