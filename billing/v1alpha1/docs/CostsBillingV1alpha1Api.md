# \CostsBillingV1alpha1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListBillingV1alpha1Costs**](CostsBillingV1alpha1Api.md#ListBillingV1alpha1Costs) | **Get** /billing/v1alpha1/costs | List of Costs



## ListBillingV1alpha1Costs

> BillingV1alpha1CostList ListBillingV1alpha1Costs(ctx).StartDate(startDate).EndDate(endDate).PageSize(pageSize).PageToken(pageToken).Execute()

List of Costs



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
    startDate := "2022-10-12" // string | Filter the results by exact match for start_date.
    endDate := "2022-10-15" // string | Filter the results by exact match for end_date.
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 5000)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CostsBillingV1alpha1Api.ListBillingV1alpha1Costs(context.Background()).StartDate(startDate).EndDate(endDate).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CostsBillingV1alpha1Api.ListBillingV1alpha1Costs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBillingV1alpha1Costs`: BillingV1alpha1CostList
    fmt.Fprintf(os.Stdout, "Response from `CostsBillingV1alpha1Api.ListBillingV1alpha1Costs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBillingV1alpha1CostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Filter the results by exact match for start_date. | 
 **endDate** | **string** | Filter the results by exact match for end_date. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 5000]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**BillingV1alpha1CostList**](billing.v1alpha1.CostList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

