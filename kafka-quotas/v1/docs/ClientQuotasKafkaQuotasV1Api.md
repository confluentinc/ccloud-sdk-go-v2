# \ClientQuotasKafkaQuotasV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKafkaQuotasV1ClientQuota**](ClientQuotasKafkaQuotasV1Api.md#CreateKafkaQuotasV1ClientQuota) | **Post** /kafka-quotas/v1/client-quotas | Create a Client Quota
[**DeleteKafkaQuotasV1ClientQuota**](ClientQuotasKafkaQuotasV1Api.md#DeleteKafkaQuotasV1ClientQuota) | **Delete** /kafka-quotas/v1/client-quotas/{id} | Delete a Client Quota
[**GetKafkaQuotasV1ClientQuota**](ClientQuotasKafkaQuotasV1Api.md#GetKafkaQuotasV1ClientQuota) | **Get** /kafka-quotas/v1/client-quotas/{id} | Read a Client Quota
[**ListKafkaQuotasV1ClientQuotas**](ClientQuotasKafkaQuotasV1Api.md#ListKafkaQuotasV1ClientQuotas) | **Get** /kafka-quotas/v1/client-quotas | List of Client Quotas
[**UpdateKafkaQuotasV1ClientQuota**](ClientQuotasKafkaQuotasV1Api.md#UpdateKafkaQuotasV1ClientQuota) | **Patch** /kafka-quotas/v1/client-quotas/{id} | Update a Client Quota



## CreateKafkaQuotasV1ClientQuota

> KafkaQuotasV1ClientQuota CreateKafkaQuotasV1ClientQuota(ctx).KafkaQuotasV1ClientQuota(kafkaQuotasV1ClientQuota).Execute()

Create a Client Quota



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
    kafkaQuotasV1ClientQuota := *openapiclient.NewKafkaQuotasV1ClientQuota() // KafkaQuotasV1ClientQuota |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientQuotasKafkaQuotasV1Api.CreateKafkaQuotasV1ClientQuota(context.Background()).KafkaQuotasV1ClientQuota(kafkaQuotasV1ClientQuota).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientQuotasKafkaQuotasV1Api.CreateKafkaQuotasV1ClientQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKafkaQuotasV1ClientQuota`: KafkaQuotasV1ClientQuota
    fmt.Fprintf(os.Stdout, "Response from `ClientQuotasKafkaQuotasV1Api.CreateKafkaQuotasV1ClientQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKafkaQuotasV1ClientQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kafkaQuotasV1ClientQuota** | [**KafkaQuotasV1ClientQuota**](KafkaQuotasV1ClientQuota.md) |  | 

### Return type

[**KafkaQuotasV1ClientQuota**](kafka-quotas.v1.ClientQuota.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKafkaQuotasV1ClientQuota

> DeleteKafkaQuotasV1ClientQuota(ctx, id).Execute()

Delete a Client Quota



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
    id := "id_example" // string | The unique identifier for the client quota.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientQuotasKafkaQuotasV1Api.DeleteKafkaQuotasV1ClientQuota(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientQuotasKafkaQuotasV1Api.DeleteKafkaQuotasV1ClientQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the client quota. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKafkaQuotasV1ClientQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetKafkaQuotasV1ClientQuota

> KafkaQuotasV1ClientQuota GetKafkaQuotasV1ClientQuota(ctx, id).Execute()

Read a Client Quota



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
    id := "id_example" // string | The unique identifier for the client quota.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientQuotasKafkaQuotasV1Api.GetKafkaQuotasV1ClientQuota(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientQuotasKafkaQuotasV1Api.GetKafkaQuotasV1ClientQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaQuotasV1ClientQuota`: KafkaQuotasV1ClientQuota
    fmt.Fprintf(os.Stdout, "Response from `ClientQuotasKafkaQuotasV1Api.GetKafkaQuotasV1ClientQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the client quota. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaQuotasV1ClientQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KafkaQuotasV1ClientQuota**](kafka-quotas.v1.ClientQuota.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaQuotasV1ClientQuotas

> KafkaQuotasV1ClientQuotaList ListKafkaQuotasV1ClientQuotas(ctx).Cluster(cluster).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()

List of Client Quotas



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
    cluster := "lkc-xxxxx" // string | Filter the results by exact match for cluster.
    environment := "env-xxxxx" // string | Filter the results by exact match for environment.
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientQuotasKafkaQuotasV1Api.ListKafkaQuotasV1ClientQuotas(context.Background()).Cluster(cluster).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientQuotasKafkaQuotasV1Api.ListKafkaQuotasV1ClientQuotas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaQuotasV1ClientQuotas`: KafkaQuotasV1ClientQuotaList
    fmt.Fprintf(os.Stdout, "Response from `ClientQuotasKafkaQuotasV1Api.ListKafkaQuotasV1ClientQuotas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaQuotasV1ClientQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | Filter the results by exact match for cluster. | 
 **environment** | **string** | Filter the results by exact match for environment. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**KafkaQuotasV1ClientQuotaList**](kafka-quotas.v1.ClientQuotaList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaQuotasV1ClientQuota

> KafkaQuotasV1ClientQuota UpdateKafkaQuotasV1ClientQuota(ctx, id).KafkaQuotasV1ClientQuotaUpdate(kafkaQuotasV1ClientQuotaUpdate).Execute()

Update a Client Quota



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
    id := "id_example" // string | The unique identifier for the client quota.
    kafkaQuotasV1ClientQuotaUpdate := *openapiclient.NewKafkaQuotasV1ClientQuotaUpdate() // KafkaQuotasV1ClientQuotaUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientQuotasKafkaQuotasV1Api.UpdateKafkaQuotasV1ClientQuota(context.Background(), id).KafkaQuotasV1ClientQuotaUpdate(kafkaQuotasV1ClientQuotaUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientQuotasKafkaQuotasV1Api.UpdateKafkaQuotasV1ClientQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKafkaQuotasV1ClientQuota`: KafkaQuotasV1ClientQuota
    fmt.Fprintf(os.Stdout, "Response from `ClientQuotasKafkaQuotasV1Api.UpdateKafkaQuotasV1ClientQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the client quota. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKafkaQuotasV1ClientQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kafkaQuotasV1ClientQuotaUpdate** | [**KafkaQuotasV1ClientQuotaUpdate**](KafkaQuotasV1ClientQuotaUpdate.md) |  | 

### Return type

[**KafkaQuotasV1ClientQuota**](kafka-quotas.v1.ClientQuota.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

