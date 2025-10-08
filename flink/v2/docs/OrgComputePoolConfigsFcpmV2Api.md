# \OrgComputePoolConfigsFcpmV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFcpmV2OrgComputePoolConfig**](OrgComputePoolConfigsFcpmV2Api.md#GetFcpmV2OrgComputePoolConfig) | **Get** /fcpm/v2/compute-pool-config | Read an Org Compute Pool Config
[**UpdateFcpmV2OrgComputePoolConfig**](OrgComputePoolConfigsFcpmV2Api.md#UpdateFcpmV2OrgComputePoolConfig) | **Patch** /fcpm/v2/compute-pool-config | Update an Org Compute Pool Config



## GetFcpmV2OrgComputePoolConfig

> FcpmV2OrgComputePoolConfig GetFcpmV2OrgComputePoolConfig(ctx).Execute()

Read an Org Compute Pool Config



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgComputePoolConfigsFcpmV2Api.GetFcpmV2OrgComputePoolConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgComputePoolConfigsFcpmV2Api.GetFcpmV2OrgComputePoolConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFcpmV2OrgComputePoolConfig`: FcpmV2OrgComputePoolConfig
    fmt.Fprintf(os.Stdout, "Response from `OrgComputePoolConfigsFcpmV2Api.GetFcpmV2OrgComputePoolConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFcpmV2OrgComputePoolConfigRequest struct via the builder pattern


### Return type

[**FcpmV2OrgComputePoolConfig**](fcpm.v2.OrgComputePoolConfig.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFcpmV2OrgComputePoolConfig

> FcpmV2OrgComputePoolConfig UpdateFcpmV2OrgComputePoolConfig(ctx).FcpmV2OrgComputePoolConfigUpdate(fcpmV2OrgComputePoolConfigUpdate).Execute()

Update an Org Compute Pool Config



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
    fcpmV2OrgComputePoolConfigUpdate := *openapiclient.NewFcpmV2OrgComputePoolConfigUpdate() // FcpmV2OrgComputePoolConfigUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgComputePoolConfigsFcpmV2Api.UpdateFcpmV2OrgComputePoolConfig(context.Background()).FcpmV2OrgComputePoolConfigUpdate(fcpmV2OrgComputePoolConfigUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgComputePoolConfigsFcpmV2Api.UpdateFcpmV2OrgComputePoolConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFcpmV2OrgComputePoolConfig`: FcpmV2OrgComputePoolConfig
    fmt.Fprintf(os.Stdout, "Response from `OrgComputePoolConfigsFcpmV2Api.UpdateFcpmV2OrgComputePoolConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFcpmV2OrgComputePoolConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fcpmV2OrgComputePoolConfigUpdate** | [**FcpmV2OrgComputePoolConfigUpdate**](FcpmV2OrgComputePoolConfigUpdate.md) |  | 

### Return type

[**FcpmV2OrgComputePoolConfig**](fcpm.v2.OrgComputePoolConfig.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

