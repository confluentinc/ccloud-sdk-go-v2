# \IPFilterSummariesIamV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIamV2IpFilterSummary**](IPFilterSummariesIamV2Api.md#GetIamV2IpFilterSummary) | **Get** /iam/v2/ip-filter-summary | Read an IP Filter Summary



## GetIamV2IpFilterSummary

> IamV2IpFilterSummary GetIamV2IpFilterSummary(ctx).Scope(scope).Execute()

Read an IP Filter Summary



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
    scope := "crn://confluent.cloud/organization=org-123/environment=env-abc" // string | Scope the operation to the given scope.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPFilterSummariesIamV2Api.GetIamV2IpFilterSummary(context.Background()).Scope(scope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPFilterSummariesIamV2Api.GetIamV2IpFilterSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIamV2IpFilterSummary`: IamV2IpFilterSummary
    fmt.Fprintf(os.Stdout, "Response from `IPFilterSummariesIamV2Api.GetIamV2IpFilterSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIamV2IpFilterSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string** | Scope the operation to the given scope. | 

### Return type

[**IamV2IpFilterSummary**](iam.v2.IpFilterSummary.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

