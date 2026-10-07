# \InferenceBudgetsBudgetV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBudgetV1InferenceBudget**](InferenceBudgetsBudgetV1Api.md#CreateBudgetV1InferenceBudget) | **Post** /budget/v1/inference-budgets | Create an Inference Budget
[**DeleteBudgetV1InferenceBudget**](InferenceBudgetsBudgetV1Api.md#DeleteBudgetV1InferenceBudget) | **Delete** /budget/v1/inference-budgets/{id} | Delete an Inference Budget
[**GetBudgetV1InferenceBudget**](InferenceBudgetsBudgetV1Api.md#GetBudgetV1InferenceBudget) | **Get** /budget/v1/inference-budgets/{id} | Read an Inference Budget
[**ListBudgetV1InferenceBudgets**](InferenceBudgetsBudgetV1Api.md#ListBudgetV1InferenceBudgets) | **Get** /budget/v1/inference-budgets | List of Inference Budgets
[**UpdateBudgetV1InferenceBudget**](InferenceBudgetsBudgetV1Api.md#UpdateBudgetV1InferenceBudget) | **Patch** /budget/v1/inference-budgets/{id} | Update an Inference Budget



## CreateBudgetV1InferenceBudget

> BudgetV1InferenceBudget CreateBudgetV1InferenceBudget(ctx).BudgetV1InferenceBudget(budgetV1InferenceBudget).Execute()

Create an Inference Budget



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
    budgetV1InferenceBudget := *openapiclient.NewBudgetV1InferenceBudget() // BudgetV1InferenceBudget |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InferenceBudgetsBudgetV1Api.CreateBudgetV1InferenceBudget(context.Background()).BudgetV1InferenceBudget(budgetV1InferenceBudget).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InferenceBudgetsBudgetV1Api.CreateBudgetV1InferenceBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBudgetV1InferenceBudget`: BudgetV1InferenceBudget
    fmt.Fprintf(os.Stdout, "Response from `InferenceBudgetsBudgetV1Api.CreateBudgetV1InferenceBudget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBudgetV1InferenceBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **budgetV1InferenceBudget** | [**BudgetV1InferenceBudget**](BudgetV1InferenceBudget.md) |  | 

### Return type

[**BudgetV1InferenceBudget**](budget.v1.InferenceBudget.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token), [confluent_auth](../README.md#confluent_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBudgetV1InferenceBudget

> DeleteBudgetV1InferenceBudget(ctx, id).Execute()

Delete an Inference Budget



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
    id := "id_example" // string | The unique identifier for the inference budget.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InferenceBudgetsBudgetV1Api.DeleteBudgetV1InferenceBudget(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InferenceBudgetsBudgetV1Api.DeleteBudgetV1InferenceBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the inference budget. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBudgetV1InferenceBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token), [confluent_auth](../README.md#confluent_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBudgetV1InferenceBudget

> BudgetV1InferenceBudget GetBudgetV1InferenceBudget(ctx, id).Execute()

Read an Inference Budget



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
    id := "id_example" // string | The unique identifier for the inference budget.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InferenceBudgetsBudgetV1Api.GetBudgetV1InferenceBudget(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InferenceBudgetsBudgetV1Api.GetBudgetV1InferenceBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBudgetV1InferenceBudget`: BudgetV1InferenceBudget
    fmt.Fprintf(os.Stdout, "Response from `InferenceBudgetsBudgetV1Api.GetBudgetV1InferenceBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the inference budget. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBudgetV1InferenceBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BudgetV1InferenceBudget**](budget.v1.InferenceBudget.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token), [confluent_auth](../README.md#confluent_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBudgetV1InferenceBudgets

> BudgetV1InferenceBudgetList ListBudgetV1InferenceBudgets(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Inference Budgets



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
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InferenceBudgetsBudgetV1Api.ListBudgetV1InferenceBudgets(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InferenceBudgetsBudgetV1Api.ListBudgetV1InferenceBudgets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBudgetV1InferenceBudgets`: BudgetV1InferenceBudgetList
    fmt.Fprintf(os.Stdout, "Response from `InferenceBudgetsBudgetV1Api.ListBudgetV1InferenceBudgets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBudgetV1InferenceBudgetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**BudgetV1InferenceBudgetList**](budget.v1.InferenceBudgetList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token), [confluent_auth](../README.md#confluent_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBudgetV1InferenceBudget

> BudgetV1InferenceBudget UpdateBudgetV1InferenceBudget(ctx, id).BudgetV1InferenceBudget(budgetV1InferenceBudget).Execute()

Update an Inference Budget



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
    id := "id_example" // string | The unique identifier for the inference budget.
    budgetV1InferenceBudget := *openapiclient.NewBudgetV1InferenceBudget() // BudgetV1InferenceBudget |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InferenceBudgetsBudgetV1Api.UpdateBudgetV1InferenceBudget(context.Background(), id).BudgetV1InferenceBudget(budgetV1InferenceBudget).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InferenceBudgetsBudgetV1Api.UpdateBudgetV1InferenceBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBudgetV1InferenceBudget`: BudgetV1InferenceBudget
    fmt.Fprintf(os.Stdout, "Response from `InferenceBudgetsBudgetV1Api.UpdateBudgetV1InferenceBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the inference budget. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBudgetV1InferenceBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **budgetV1InferenceBudget** | [**BudgetV1InferenceBudget**](BudgetV1InferenceBudget.md) |  | 

### Return type

[**BudgetV1InferenceBudget**](budget.v1.InferenceBudget.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token), [confluent_auth](../README.md#confluent_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

