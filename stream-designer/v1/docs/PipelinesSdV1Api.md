# \PipelinesSdV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSdV1Pipeline**](PipelinesSdV1Api.md#CreateSdV1Pipeline) | **Post** /sd/v1/pipelines | Create a Pipeline
[**DeleteSdV1Pipeline**](PipelinesSdV1Api.md#DeleteSdV1Pipeline) | **Delete** /sd/v1/pipelines/{id} | Delete a Pipeline
[**GetSdV1Pipeline**](PipelinesSdV1Api.md#GetSdV1Pipeline) | **Get** /sd/v1/pipelines/{id} | Read a Pipeline
[**ListSdV1Pipelines**](PipelinesSdV1Api.md#ListSdV1Pipelines) | **Get** /sd/v1/pipelines | List of Pipelines
[**UpdateSdV1Pipeline**](PipelinesSdV1Api.md#UpdateSdV1Pipeline) | **Patch** /sd/v1/pipelines/{id} | Update a Pipeline



## CreateSdV1Pipeline

> SdV1Pipeline CreateSdV1Pipeline(ctx).SdV1Pipeline(sdV1Pipeline).Execute()

Create a Pipeline



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
    sdV1Pipeline := *openapiclient.NewSdV1Pipeline() // SdV1Pipeline |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesSdV1Api.CreateSdV1Pipeline(context.Background()).SdV1Pipeline(sdV1Pipeline).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesSdV1Api.CreateSdV1Pipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSdV1Pipeline`: SdV1Pipeline
    fmt.Fprintf(os.Stdout, "Response from `PipelinesSdV1Api.CreateSdV1Pipeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSdV1PipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sdV1Pipeline** | [**SdV1Pipeline**](SdV1Pipeline.md) |  | 

### Return type

[**SdV1Pipeline**](sd.v1.Pipeline.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSdV1Pipeline

> DeleteSdV1Pipeline(ctx, id).Environment(environment).SpecKafkaCluster(specKafkaCluster).Execute()

Delete a Pipeline



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
    id := "id_example" // string | The unique identifier for the pipeline.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesSdV1Api.DeleteSdV1Pipeline(context.Background(), id).Environment(environment).SpecKafkaCluster(specKafkaCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesSdV1Api.DeleteSdV1Pipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSdV1PipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 
 **specKafkaCluster** | **string** | Scope the operation to the given spec.kafka_cluster. | 


### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdV1Pipeline

> SdV1Pipeline GetSdV1Pipeline(ctx, id).Environment(environment).SpecKafkaCluster(specKafkaCluster).IncludeSql(includeSql).Execute()

Read a Pipeline



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
    id := "id_example" // string | The unique identifier for the pipeline.
    includeSql := "includeSql_example" // string | Include source code or not, values can be true or false (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesSdV1Api.GetSdV1Pipeline(context.Background(), id).Environment(environment).SpecKafkaCluster(specKafkaCluster).IncludeSql(includeSql).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesSdV1Api.GetSdV1Pipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdV1Pipeline`: SdV1Pipeline
    fmt.Fprintf(os.Stdout, "Response from `PipelinesSdV1Api.GetSdV1Pipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSdV1PipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 
 **specKafkaCluster** | **string** | Scope the operation to the given spec.kafka_cluster. | 

 **includeSql** | **string** | Include source code or not, values can be true or false | 

### Return type

[**SdV1Pipeline**](sd.v1.Pipeline.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSdV1Pipelines

> SdV1PipelineList ListSdV1Pipelines(ctx).Environment(environment).SpecKafkaCluster(specKafkaCluster).PageSize(pageSize).PageToken(pageToken).Execute()

List of Pipelines



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
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 100)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesSdV1Api.ListSdV1Pipelines(context.Background()).Environment(environment).SpecKafkaCluster(specKafkaCluster).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesSdV1Api.ListSdV1Pipelines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSdV1Pipelines`: SdV1PipelineList
    fmt.Fprintf(os.Stdout, "Response from `PipelinesSdV1Api.ListSdV1Pipelines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSdV1PipelinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **specKafkaCluster** | **string** | Filter the results by exact match for spec.kafka_cluster. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 100]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**SdV1PipelineList**](sd.v1.PipelineList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSdV1Pipeline

> SdV1Pipeline UpdateSdV1Pipeline(ctx, id).Environment(environment).SpecKafkaCluster(specKafkaCluster).SdV1PipelineUpdate(sdV1PipelineUpdate).Execute()

Update a Pipeline



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
    id := "id_example" // string | The unique identifier for the pipeline.
    sdV1PipelineUpdate := *openapiclient.NewSdV1PipelineUpdate() // SdV1PipelineUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesSdV1Api.UpdateSdV1Pipeline(context.Background(), id).Environment(environment).SpecKafkaCluster(specKafkaCluster).SdV1PipelineUpdate(sdV1PipelineUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesSdV1Api.UpdateSdV1Pipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSdV1Pipeline`: SdV1Pipeline
    fmt.Fprintf(os.Stdout, "Response from `PipelinesSdV1Api.UpdateSdV1Pipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSdV1PipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 
 **specKafkaCluster** | **string** | Scope the operation to the given spec.kafka_cluster. | 

 **sdV1PipelineUpdate** | [**SdV1PipelineUpdate**](SdV1PipelineUpdate.md) |  | 

### Return type

[**SdV1Pipeline**](sd.v1.Pipeline.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

