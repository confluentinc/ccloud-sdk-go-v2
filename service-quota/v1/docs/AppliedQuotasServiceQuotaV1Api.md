# \AppliedQuotasServiceQuotaV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceQuotaV1AppliedQuota**](AppliedQuotasServiceQuotaV1Api.md#GetServiceQuotaV1AppliedQuota) | **Get** /service-quota/v1/applied-quotas/{id} | Read an Applied Quota
[**ListServiceQuotaV1AppliedQuotas**](AppliedQuotasServiceQuotaV1Api.md#ListServiceQuotaV1AppliedQuotas) | **Get** /service-quota/v1/applied-quotas | List of Applied Quotas



## GetServiceQuotaV1AppliedQuota

> ServiceQuotaV1AppliedQuota GetServiceQuotaV1AppliedQuota(ctx, id).Environment(environment).Network(network).KafkaCluster(kafkaCluster).Execute()

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
    environment := "env-00000" // string | The environment ID the quota is associated with. This field is only required when retrieving a single quota and the scope of quota is \"ENVIRONMENT\" or \"NETWORK\" or \"KAFKA_CLUSTER\".  (optional)
    network := "n-12034" // string | The network ID the quota is associated with. This field is only required when retrieving a single quota and the scope of quota is \"NETWORK\".  (optional)
    kafkaCluster := "lkc-00000" // string | The kafka cluster ID the quota is associated with. This field is required only when the scope of quota is \"KAFKA_CLUSTER\".  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppliedQuotasServiceQuotaV1Api.GetServiceQuotaV1AppliedQuota(context.Background(), id).Environment(environment).Network(network).KafkaCluster(kafkaCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppliedQuotasServiceQuotaV1Api.GetServiceQuotaV1AppliedQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceQuotaV1AppliedQuota`: ServiceQuotaV1AppliedQuota
    fmt.Fprintf(os.Stdout, "Response from `AppliedQuotasServiceQuotaV1Api.GetServiceQuotaV1AppliedQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the applied quota. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceQuotaV1AppliedQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environment** | **string** | The environment ID the quota is associated with. This field is only required when retrieving a single quota and the scope of quota is \&quot;ENVIRONMENT\&quot; or \&quot;NETWORK\&quot; or \&quot;KAFKA_CLUSTER\&quot;.  | 
 **network** | **string** | The network ID the quota is associated with. This field is only required when retrieving a single quota and the scope of quota is \&quot;NETWORK\&quot;.  | 
 **kafkaCluster** | **string** | The kafka cluster ID the quota is associated with. This field is required only when the scope of quota is \&quot;KAFKA_CLUSTER\&quot;.  | 

### Return type

[**ServiceQuotaV1AppliedQuota**](service-quota.v1.AppliedQuota.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceQuotaV1AppliedQuotas

> ServiceQuotaV1AppliedQuotaList ListServiceQuotaV1AppliedQuotas(ctx).Scope(scope).Environment(environment).Network(network).KafkaCluster(kafkaCluster).Id(id).PageSize(pageSize).PageToken(pageToken).Execute()

List of Applied Quotas



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
    scope := "ORGANIZATION" // string | The applied scope the quota belongs to. 
    environment := "env-00000" // string | The environment ID the quota is associated with.  (optional)
    network := "n-12034" // string | The network ID the quota is associated with.  (optional)
    kafkaCluster := "lkc-00000" // string | The kafka cluster ID the quota is associated with.  (optional)
    id := "iam.max_environments.per_org" // string | The id (quota code) that this quota belongs to.  (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppliedQuotasServiceQuotaV1Api.ListServiceQuotaV1AppliedQuotas(context.Background()).Scope(scope).Environment(environment).Network(network).KafkaCluster(kafkaCluster).Id(id).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppliedQuotasServiceQuotaV1Api.ListServiceQuotaV1AppliedQuotas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceQuotaV1AppliedQuotas`: ServiceQuotaV1AppliedQuotaList
    fmt.Fprintf(os.Stdout, "Response from `AppliedQuotasServiceQuotaV1Api.ListServiceQuotaV1AppliedQuotas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServiceQuotaV1AppliedQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string** | The applied scope the quota belongs to.  | 
 **environment** | **string** | The environment ID the quota is associated with.  | 
 **network** | **string** | The network ID the quota is associated with.  | 
 **kafkaCluster** | **string** | The kafka cluster ID the quota is associated with.  | 
 **id** | **string** | The id (quota code) that this quota belongs to.  | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**ServiceQuotaV1AppliedQuotaList**](service-quota.v1.AppliedQuotaList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

