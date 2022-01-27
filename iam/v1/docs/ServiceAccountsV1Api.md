# \ServiceAccountsV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListV1ServiceAccounts**](ServiceAccountsV1Api.md#ListV1ServiceAccounts) | **Get** /service_accounts | List of Service Accounts



## ListV1ServiceAccounts

> V1ServiceAccountList ListV1ServiceAccounts(ctx).Execute()

List of Service Accounts



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
    resp, r, err := api_client.ServiceAccountsV1Api.ListV1ServiceAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsV1Api.ListV1ServiceAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListV1ServiceAccounts`: V1ServiceAccountList
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsV1Api.ListV1ServiceAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListV1ServiceAccountsRequest struct via the builder pattern


### Return type

[**V1ServiceAccountList**](v1.ServiceAccountList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

