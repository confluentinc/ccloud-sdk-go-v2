# \AppliedQuotaQuotasV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQuotasV2AppliedQuota**](AppliedQuotaQuotasV2Api.md#GetQuotasV2AppliedQuota) | **Get** /quotas/v2/applied-quotas/{id} | Read an Applied Quota
[**ListQuotasV2AppliedQuota**](AppliedQuotaQuotasV2Api.md#ListQuotasV2AppliedQuota) | **Get** /quotas/v2/applied-quotas | List of Applied Quota



## GetQuotasV2AppliedQuota

> QuotasV2AppliedQuota GetQuotasV2AppliedQuota(ctx, id).Environment(environment).KafkaCluster(kafkaCluster).Execute()

Read an Applied Quota



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
    id := "id_example" // string | The unique identifier for the applied quota.
    environment := "env-00000" // string | A unique environment id to associate a specific environment to this quota. This field is required only if scope is set to \"ENVIRONMENT\" or \"KAFKA-CLUSTER\" and it is doing a single quota query.  (optional)
    kafkaCluster := "lkc-00000" // string | A unique Kafka cluster id to associate a specific kafka cluster to this quota. This field is required only if scope is set to \"KAFKA-CLUSTER\" and it is doing a single quota query.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppliedQuotaQuotasV2Api.GetQuotasV2AppliedQuota(context.Background(), id).Environment(environment).KafkaCluster(kafkaCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppliedQuotaQuotasV2Api.GetQuotasV2AppliedQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQuotasV2AppliedQuota`: QuotasV2AppliedQuota
    fmt.Fprintf(os.Stdout, "Response from `AppliedQuotaQuotasV2Api.GetQuotasV2AppliedQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the applied quota. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuotasV2AppliedQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environment** | **string** | A unique environment id to associate a specific environment to this quota. This field is required only if scope is set to \&quot;ENVIRONMENT\&quot; or \&quot;KAFKA-CLUSTER\&quot; and it is doing a single quota query.  | 
 **kafkaCluster** | **string** | A unique Kafka cluster id to associate a specific kafka cluster to this quota. This field is required only if scope is set to \&quot;KAFKA-CLUSTER\&quot; and it is doing a single quota query.  | 

### Return type

[**QuotasV2AppliedQuota**](quotas.v2.AppliedQuota.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQuotasV2AppliedQuota

> QuotasV2AppliedQuotaList ListQuotasV2AppliedQuota(ctx).Scope(scope).Environment(environment).KafkaCluster(kafkaCluster).PageSize(pageSize).PageToken(pageToken).Execute()

List of Applied Quota



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
    scope := "ORGANIZATION" // string | The applied scope that the quota belong to. 
    environment := "env-00000" // string | A unique environment id to associate a specific environment to this quota. This field is required only if scope is set to \"ENVIRONMENT\" or \"KAFKA-CLUSTER\" and it is doing a single quota query.  (optional)
    kafkaCluster := "lkc-00000" // string | A unique Kafka cluster id to associate a specific kafka cluster to this quota. This field is required only if scope is set to \"KAFKA-CLUSTER\" and it is doing a single quota query.  (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppliedQuotaQuotasV2Api.ListQuotasV2AppliedQuota(context.Background()).Scope(scope).Environment(environment).KafkaCluster(kafkaCluster).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppliedQuotaQuotasV2Api.ListQuotasV2AppliedQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListQuotasV2AppliedQuota`: QuotasV2AppliedQuotaList
    fmt.Fprintf(os.Stdout, "Response from `AppliedQuotaQuotasV2Api.ListQuotasV2AppliedQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListQuotasV2AppliedQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string** | The applied scope that the quota belong to.  | 
 **environment** | **string** | A unique environment id to associate a specific environment to this quota. This field is required only if scope is set to \&quot;ENVIRONMENT\&quot; or \&quot;KAFKA-CLUSTER\&quot; and it is doing a single quota query.  | 
 **kafkaCluster** | **string** | A unique Kafka cluster id to associate a specific kafka cluster to this quota. This field is required only if scope is set to \&quot;KAFKA-CLUSTER\&quot; and it is doing a single quota query.  | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**QuotasV2AppliedQuotaList**](QuotasV2AppliedQuotaList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

