# \AppliedQuotasServiceQuotaV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceQuotaV2AppliedQuota**](AppliedQuotasServiceQuotaV2Api.md#GetServiceQuotaV2AppliedQuota) | **Get** /service-quota/v2/applied-quotas/{id} | Read an Applied Quota
[**ListServiceQuotaV2AppliedQuotas**](AppliedQuotasServiceQuotaV2Api.md#ListServiceQuotaV2AppliedQuotas) | **Get** /service-quota/v2/applied-quotas | List of Applied Quotas



## GetServiceQuotaV2AppliedQuota

> ServiceQuotaV2AppliedQuota GetServiceQuotaV2AppliedQuota(ctx, id).Environment(environment).Network(network).KafkaCluster(kafkaCluster).Execute()

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
    resp, r, err := api_client.AppliedQuotasServiceQuotaV2Api.GetServiceQuotaV2AppliedQuota(context.Background(), id).Environment(environment).Network(network).KafkaCluster(kafkaCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppliedQuotasServiceQuotaV2Api.GetServiceQuotaV2AppliedQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceQuotaV2AppliedQuota`: ServiceQuotaV2AppliedQuota
    fmt.Fprintf(os.Stdout, "Response from `AppliedQuotasServiceQuotaV2Api.GetServiceQuotaV2AppliedQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the applied quota. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceQuotaV2AppliedQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environment** | **string** | The environment ID the quota is associated with. This field is only required when retrieving a single quota and the scope of quota is \&quot;ENVIRONMENT\&quot; or \&quot;NETWORK\&quot; or \&quot;KAFKA_CLUSTER\&quot;.  | 
 **network** | **string** | The network ID the quota is associated with. This field is only required when retrieving a single quota and the scope of quota is \&quot;NETWORK\&quot;.  | 
 **kafkaCluster** | **string** | The kafka cluster ID the quota is associated with. This field is required only when the scope of quota is \&quot;KAFKA_CLUSTER\&quot;.  | 

### Return type

[**ServiceQuotaV2AppliedQuota**](service-quota.v2.AppliedQuota.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceQuotaV2AppliedQuotas

> ServiceQuotaV2AppliedQuotaList ListServiceQuotaV2AppliedQuotas(ctx).Scope(scope).Environment(environment).Network(network).KafkaCluster(kafkaCluster).QuotaCode(quotaCode).PageSize(pageSize).PageToken(pageToken).Execute()

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
    quotaCode := "iam.max_environments.per_org" // string | The quota code that this quota belongs to.  (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppliedQuotasServiceQuotaV2Api.ListServiceQuotaV2AppliedQuotas(context.Background()).Scope(scope).Environment(environment).Network(network).KafkaCluster(kafkaCluster).QuotaCode(quotaCode).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppliedQuotasServiceQuotaV2Api.ListServiceQuotaV2AppliedQuotas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceQuotaV2AppliedQuotas`: ServiceQuotaV2AppliedQuotaList
    fmt.Fprintf(os.Stdout, "Response from `AppliedQuotasServiceQuotaV2Api.ListServiceQuotaV2AppliedQuotas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServiceQuotaV2AppliedQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string** | The applied scope the quota belongs to.  | 
 **environment** | **string** | The environment ID the quota is associated with.  | 
 **network** | **string** | The network ID the quota is associated with.  | 
 **kafkaCluster** | **string** | The kafka cluster ID the quota is associated with.  | 
 **quotaCode** | **string** | The quota code that this quota belongs to.  | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**ServiceQuotaV2AppliedQuotaList**](service-quota.v2.AppliedQuotaList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

