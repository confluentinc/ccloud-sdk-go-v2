# \AppliedQuotaQuotasV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQuotasV2AppliedQuota**](AppliedQuotaQuotasV2Api.md#GetQuotasV2AppliedQuota) | **Get** /quotas/v2/applied-quotas/{id} | Read an Applied Quota
[**ListQuotasV2AppliedQuota**](AppliedQuotaQuotasV2Api.md#ListQuotasV2AppliedQuota) | **Get** /quotas/v2/applied-quotas | List of Applied Quota



## GetQuotasV2AppliedQuota

> QuotasV2AppliedQuota GetQuotasV2AppliedQuota(ctx, id).Scope(scope).User(user).Organization(organization).Environment(environment).Cluster(cluster).Execute()

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
    scope := "cluster" // string | Scope the operation to the given scope.
    id := "id_example" // string | The unique identifier for the applied quota.
    user := "u-4r6rk7" // string | Scope the operation to the given user. (optional)
    organization := "b3a17773-05cc-4431-9560-433fb4613da8" // string | Scope the operation to the given organization. (optional)
    environment := "env-00000" // string | Scope the operation to the given environment. (optional)
    cluster := "lkc-00000" // string | Scope the operation to the given cluster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppliedQuotaQuotasV2Api.GetQuotasV2AppliedQuota(context.Background(), id).Scope(scope).User(user).Organization(organization).Environment(environment).Cluster(cluster).Execute()
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
 **scope** | **string** | Scope the operation to the given scope. | 

 **user** | **string** | Scope the operation to the given user. | 
 **organization** | **string** | Scope the operation to the given organization. | 
 **environment** | **string** | Scope the operation to the given environment. | 
 **cluster** | **string** | Scope the operation to the given cluster. | 

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

> QuotasV2AppliedQuotaList ListQuotasV2AppliedQuota(ctx).Scope(scope).User(user).Organization(organization).Environment(environment).Cluster(cluster).Execute()

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
    scope := "cluster" // string | Filter the results by exact match for scope.
    user := "u-4r6rk7" // string | Filter the results by exact match for user. (optional)
    organization := "b3a17773-05cc-4431-9560-433fb4613da8" // string | Filter the results by exact match for organization. (optional)
    environment := "env-00000" // string | Filter the results by exact match for environment. (optional)
    cluster := "lkc-00000" // string | Filter the results by exact match for cluster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppliedQuotaQuotasV2Api.ListQuotasV2AppliedQuota(context.Background()).Scope(scope).User(user).Organization(organization).Environment(environment).Cluster(cluster).Execute()
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
 **scope** | **string** | Filter the results by exact match for scope. | 
 **user** | **string** | Filter the results by exact match for user. | 
 **organization** | **string** | Filter the results by exact match for organization. | 
 **environment** | **string** | Filter the results by exact match for environment. | 
 **cluster** | **string** | Filter the results by exact match for cluster. | 

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

