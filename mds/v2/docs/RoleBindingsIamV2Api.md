# \RoleBindingsIamV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIamV2RoleBinding**](RoleBindingsIamV2Api.md#CreateIamV2RoleBinding) | **Post** /iam/v2/role-bindings | Create a Role Binding
[**DeleteIamV2RoleBinding**](RoleBindingsIamV2Api.md#DeleteIamV2RoleBinding) | **Delete** /iam/v2/role-bindings/{id} | Delete a Role Binding
[**GetIamV2RoleBinding**](RoleBindingsIamV2Api.md#GetIamV2RoleBinding) | **Get** /iam/v2/role-bindings/{id} | Read a Role Binding
[**ListIamV2RoleBindings**](RoleBindingsIamV2Api.md#ListIamV2RoleBindings) | **Get** /iam/v2/role-bindings | List of Role Bindings



## CreateIamV2RoleBinding

> IamV2RoleBinding CreateIamV2RoleBinding(ctx).IamV2RoleBinding(iamV2RoleBinding).Execute()

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
    iamV2RoleBinding := *openapiclient.NewIamV2RoleBinding() // IamV2RoleBinding |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleBindingsIamV2Api.CreateIamV2RoleBinding(context.Background()).IamV2RoleBinding(iamV2RoleBinding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsIamV2Api.CreateIamV2RoleBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIamV2RoleBinding`: IamV2RoleBinding
    fmt.Fprintf(os.Stdout, "Response from `RoleBindingsIamV2Api.CreateIamV2RoleBinding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIamV2RoleBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamV2RoleBinding** | [**IamV2RoleBinding**](IamV2RoleBinding.md) |  | 

### Return type

[**IamV2RoleBinding**](iam.v2.RoleBinding.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIamV2RoleBinding

> DeleteIamV2RoleBinding(ctx, id).Execute()

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
    resp, r, err := api_client.RoleBindingsIamV2Api.DeleteIamV2RoleBinding(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsIamV2Api.DeleteIamV2RoleBinding``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteIamV2RoleBindingRequest struct via the builder pattern


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


## GetIamV2RoleBinding

> IamV2RoleBinding GetIamV2RoleBinding(ctx, id).Execute()

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
    resp, r, err := api_client.RoleBindingsIamV2Api.GetIamV2RoleBinding(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsIamV2Api.GetIamV2RoleBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIamV2RoleBinding`: IamV2RoleBinding
    fmt.Fprintf(os.Stdout, "Response from `RoleBindingsIamV2Api.GetIamV2RoleBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the role binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamV2RoleBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamV2RoleBinding**](iam.v2.RoleBinding.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIamV2RoleBindings

> IamV2RoleBindingList ListIamV2RoleBindings(ctx).CrnPattern(crnPattern).Principal(principal).RoleName(roleName).PageSize(pageSize).PageToken(pageToken).Execute()

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
    resp, r, err := api_client.RoleBindingsIamV2Api.ListIamV2RoleBindings(context.Background()).CrnPattern(crnPattern).Principal(principal).RoleName(roleName).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsIamV2Api.ListIamV2RoleBindings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIamV2RoleBindings`: IamV2RoleBindingList
    fmt.Fprintf(os.Stdout, "Response from `RoleBindingsIamV2Api.ListIamV2RoleBindings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIamV2RoleBindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **crnPattern** | **string** | Filter the results by a partial search of crn_pattern. | 
 **principal** | **string** | Filter the results by exact match for principal. | 
 **roleName** | **string** | Filter the results by exact match for role_name. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**IamV2RoleBindingList**](IamV2RoleBindingList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

