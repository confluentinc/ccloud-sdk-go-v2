# \UsagesTerraformUsageV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTerraformUsageV1Usage**](UsagesTerraformUsageV1Api.md#CreateTerraformUsageV1Usage) | **Post** /terraform-usage/v1/usages | Create a Usage



## CreateTerraformUsageV1Usage

> CreateTerraformUsageV1Usage(ctx).TerraformUsageV1Usage(terraformUsageV1Usage).Execute()

Create a Usage



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
    terraformUsageV1Usage := *openapiclient.NewTerraformUsageV1Usage() // TerraformUsageV1Usage |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsagesTerraformUsageV1Api.CreateTerraformUsageV1Usage(context.Background()).TerraformUsageV1Usage(terraformUsageV1Usage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsagesTerraformUsageV1Api.CreateTerraformUsageV1Usage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTerraformUsageV1UsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **terraformUsageV1Usage** | [**TerraformUsageV1Usage**](TerraformUsageV1Usage.md) |  | 

### Return type

 (empty response body)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token), [confluent_auth](../README.md#confluent_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

