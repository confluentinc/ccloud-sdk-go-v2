# \EnvironmentsOrgV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgV2Environment**](EnvironmentsOrgV2Api.md#CreateOrgV2Environment) | **Post** /org/v2/environments | Create an Environment
[**DeleteOrgV2Environment**](EnvironmentsOrgV2Api.md#DeleteOrgV2Environment) | **Delete** /org/v2/environments/{id} | Delete an Environment
[**GetOrgV2Environment**](EnvironmentsOrgV2Api.md#GetOrgV2Environment) | **Get** /org/v2/environments/{id} | Read an Environment
[**ListOrgV2Environments**](EnvironmentsOrgV2Api.md#ListOrgV2Environments) | **Get** /org/v2/environments | List of Environments
[**UpdateOrgV2Environment**](EnvironmentsOrgV2Api.md#UpdateOrgV2Environment) | **Patch** /org/v2/environments/{id} | Update an Environment



## CreateOrgV2Environment

> OrgV2Environment CreateOrgV2Environment(ctx).OrgV2Environment(orgV2Environment).Execute()

Create an Environment



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
    orgV2Environment := *openapiclient.NewOrgV2Environment() // OrgV2Environment |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentsOrgV2Api.CreateOrgV2Environment(context.Background()).OrgV2Environment(orgV2Environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsOrgV2Api.CreateOrgV2Environment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrgV2Environment`: OrgV2Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsOrgV2Api.CreateOrgV2Environment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgV2EnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgV2Environment** | [**OrgV2Environment**](OrgV2Environment.md) |  | 

### Return type

[**OrgV2Environment**](org.v2.Environment.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgV2Environment

> DeleteOrgV2Environment(ctx, id).Execute()

Delete an Environment



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
    id := "id_example" // string | The unique identifier for the environment.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentsOrgV2Api.DeleteOrgV2Environment(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsOrgV2Api.DeleteOrgV2Environment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgV2EnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgV2Environment

> OrgV2Environment GetOrgV2Environment(ctx, id).Execute()

Read an Environment



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
    id := "id_example" // string | The unique identifier for the environment.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentsOrgV2Api.GetOrgV2Environment(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsOrgV2Api.GetOrgV2Environment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgV2Environment`: OrgV2Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsOrgV2Api.GetOrgV2Environment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgV2EnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrgV2Environment**](org.v2.Environment.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgV2Environments

> OrgV2EnvironmentList ListOrgV2Environments(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Environments



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
    resp, r, err := api_client.EnvironmentsOrgV2Api.ListOrgV2Environments(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsOrgV2Api.ListOrgV2Environments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgV2Environments`: OrgV2EnvironmentList
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsOrgV2Api.ListOrgV2Environments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrgV2EnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**OrgV2EnvironmentList**](org.v2.EnvironmentList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgV2Environment

> OrgV2Environment UpdateOrgV2Environment(ctx, id).OrgV2Environment(orgV2Environment).Execute()

Update an Environment



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
    id := "id_example" // string | The unique identifier for the environment.
    orgV2Environment := *openapiclient.NewOrgV2Environment() // OrgV2Environment |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentsOrgV2Api.UpdateOrgV2Environment(context.Background(), id).OrgV2Environment(orgV2Environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsOrgV2Api.UpdateOrgV2Environment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrgV2Environment`: OrgV2Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsOrgV2Api.UpdateOrgV2Environment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgV2EnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgV2Environment** | [**OrgV2Environment**](OrgV2Environment.md) |  | 

### Return type

[**OrgV2Environment**](org.v2.Environment.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

