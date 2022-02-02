# \RoleBindingsV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateV2RoleBinding**](RoleBindingsV2Api.md#CreateV2RoleBinding) | **Post** /v2/role-bindings | Create a Role Binding
[**DeleteV2RoleBinding**](RoleBindingsV2Api.md#DeleteV2RoleBinding) | **Delete** /v2/role-bindings/{id} | Delete a Role Binding
[**GetV2RoleBinding**](RoleBindingsV2Api.md#GetV2RoleBinding) | **Get** /v2/role-bindings/{id} | Read a Role Binding
[**ListV2RoleBindings**](RoleBindingsV2Api.md#ListV2RoleBindings) | **Get** /v2/role-bindings | List of Role Bindings



## CreateV2RoleBinding

> V2RoleBinding CreateV2RoleBinding(ctx).V2RoleBinding(v2RoleBinding).Execute()

Create a Role Binding



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
    v2RoleBinding := *openapiclient.NewV2RoleBinding() // V2RoleBinding |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleBindingsV2Api.CreateV2RoleBinding(context.Background()).V2RoleBinding(v2RoleBinding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsV2Api.CreateV2RoleBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateV2RoleBinding`: V2RoleBinding
    fmt.Fprintf(os.Stdout, "Response from `RoleBindingsV2Api.CreateV2RoleBinding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateV2RoleBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2RoleBinding** | [**V2RoleBinding**](V2RoleBinding.md) |  | 

### Return type

[**V2RoleBinding**](v2.RoleBinding.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteV2RoleBinding

> DeleteV2RoleBinding(ctx, id).Execute()

Delete a Role Binding



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
    id := "id_example" // string | The unique identifier for the role binding.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleBindingsV2Api.DeleteV2RoleBinding(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsV2Api.DeleteV2RoleBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the role binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV2RoleBindingRequest struct via the builder pattern


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


## GetV2RoleBinding

> V2RoleBinding GetV2RoleBinding(ctx, id).Execute()

Read a Role Binding



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
    id := "id_example" // string | The unique identifier for the role binding.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleBindingsV2Api.GetV2RoleBinding(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsV2Api.GetV2RoleBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetV2RoleBinding`: V2RoleBinding
    fmt.Fprintf(os.Stdout, "Response from `RoleBindingsV2Api.GetV2RoleBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the role binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV2RoleBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2RoleBinding**](v2.RoleBinding.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListV2RoleBindings

> V2RoleBindingList ListV2RoleBindings(ctx).CrnPattern(crnPattern).Principal(principal).RoleName(roleName).PageSize(pageSize).PageToken(pageToken).Execute()

List of Role Bindings



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
    crnPattern := "crn://confluent.cloud/organization=1111aaaa-11aa-11aa-11aa-111111aaaaaa/environment=t-38nbaz/cloud-cluster=lkc-1111aaa" // string | Filter the results by a partial search of crn_pattern.
    principal := "User:u-111aaa" // string | Filter the results by exact match for principal. (optional)
    roleName := "CloudClusterAdmin" // string | Filter the results by exact match for role_name. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleBindingsV2Api.ListV2RoleBindings(context.Background()).CrnPattern(crnPattern).Principal(principal).RoleName(roleName).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsV2Api.ListV2RoleBindings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListV2RoleBindings`: V2RoleBindingList
    fmt.Fprintf(os.Stdout, "Response from `RoleBindingsV2Api.ListV2RoleBindings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListV2RoleBindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **crnPattern** | **string** | Filter the results by a partial search of crn_pattern. | 
 **principal** | **string** | Filter the results by exact match for principal. | 
 **roleName** | **string** | Filter the results by exact match for role_name. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**V2RoleBindingList**](V2RoleBindingList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

