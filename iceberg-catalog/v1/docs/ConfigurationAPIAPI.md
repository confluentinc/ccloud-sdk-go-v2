# \ConfigurationAPIAPI

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfig**](ConfigurationAPIAPI.md#GetConfig) | **Get** /v1/config | List all catalog configuration settings



## GetConfig

> CatalogConfig GetConfig(ctx).Warehouse(warehouse).Execute()

List all catalog configuration settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/confluentinc/ccloud-sdk-go-v2"
)

func main() {
	warehouse := "warehouse_example" // string | Warehouse location or identifier to request from the service (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPIAPI.GetConfig(context.Background()).Warehouse(warehouse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPIAPI.GetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfig`: CatalogConfig
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPIAPI.GetConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **warehouse** | **string** | Warehouse location or identifier to request from the service | 

### Return type

[**CatalogConfig**](CatalogConfig.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

