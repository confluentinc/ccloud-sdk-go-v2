# \CIDRV1alpha1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CidrValidate**](CIDRV1alpha1Api.md#CidrValidate) | **Post** /cidr/validate | validates network CIDR



## CidrValidate

> CidrValidate(ctx).CidrValidate(cidrValidate).Execute()

validates network CIDR



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
    cidrValidate := *openapiclient.NewCidrValidate() // CidrValidate | request validation (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CIDRV1alpha1Api.CidrValidate(context.Background()).CidrValidate(cidrValidate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CIDRV1alpha1Api.CidrValidate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCidrValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cidrValidate** | [**CidrValidate**](CidrValidate.md) | request validation | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

