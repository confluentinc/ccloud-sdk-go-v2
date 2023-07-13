# \IamBindingsFcpmV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFcpmV2IamBinding**](IamBindingsFcpmV2Api.md#CreateFcpmV2IamBinding) | **Post** /fcpm/v2/iam-bindings | Create an Iam Binding
[**DeleteFcpmV2IamBinding**](IamBindingsFcpmV2Api.md#DeleteFcpmV2IamBinding) | **Delete** /fcpm/v2/iam-bindings/{id} | Delete an Iam Binding
[**ListFcpmV2IamBindings**](IamBindingsFcpmV2Api.md#ListFcpmV2IamBindings) | **Get** /fcpm/v2/iam-bindings | List of Iam Bindings



## CreateFcpmV2IamBinding

> FcpmV2IamBinding CreateFcpmV2IamBinding(ctx).FcpmV2IamBinding(fcpmV2IamBinding).Execute()

Create an Iam Binding



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
    fcpmV2IamBinding := *openapiclient.NewFcpmV2IamBinding() // FcpmV2IamBinding |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamBindingsFcpmV2Api.CreateFcpmV2IamBinding(context.Background()).FcpmV2IamBinding(fcpmV2IamBinding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamBindingsFcpmV2Api.CreateFcpmV2IamBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFcpmV2IamBinding`: FcpmV2IamBinding
    fmt.Fprintf(os.Stdout, "Response from `IamBindingsFcpmV2Api.CreateFcpmV2IamBinding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFcpmV2IamBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fcpmV2IamBinding** | [**FcpmV2IamBinding**](FcpmV2IamBinding.md) |  | 

### Return type

[**FcpmV2IamBinding**](fcpm.v2.IamBinding.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFcpmV2IamBinding

> DeleteFcpmV2IamBinding(ctx, id).Environment(environment).IdentityPool(identityPool).Execute()

Delete an Iam Binding



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
    environment := "env-00000" // string | Scope the operation to the given environment.
    id := "id_example" // string | The unique identifier for the iam binding.
    identityPool := "iam-00000" // string | Scope the operation to the given identity_pool. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamBindingsFcpmV2Api.DeleteFcpmV2IamBinding(context.Background(), id).Environment(environment).IdentityPool(identityPool).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamBindingsFcpmV2Api.DeleteFcpmV2IamBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the iam binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFcpmV2IamBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 

 **identityPool** | **string** | Scope the operation to the given identity_pool. | 

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


## ListFcpmV2IamBindings

> FcpmV2IamBindingList ListFcpmV2IamBindings(ctx).Environment(environment).Region(region).Cloud(cloud).IdentityPool(identityPool).PageSize(pageSize).PageToken(pageToken).Execute()

List of Iam Bindings



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
    environment := "env-00000" // string | Filter the results by exact match for environment.
    region := "us-west-1" // string | Filter the results by exact match for region. (optional)
    cloud := "AWS" // string | Filter the results by exact match for cloud. (optional)
    identityPool := "iam-00000" // string | Filter the results by exact match for identity_pool. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamBindingsFcpmV2Api.ListFcpmV2IamBindings(context.Background()).Environment(environment).Region(region).Cloud(cloud).IdentityPool(identityPool).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamBindingsFcpmV2Api.ListFcpmV2IamBindings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFcpmV2IamBindings`: FcpmV2IamBindingList
    fmt.Fprintf(os.Stdout, "Response from `IamBindingsFcpmV2Api.ListFcpmV2IamBindings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFcpmV2IamBindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **region** | **string** | Filter the results by exact match for region. | 
 **cloud** | **string** | Filter the results by exact match for cloud. | 
 **identityPool** | **string** | Filter the results by exact match for identity_pool. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**FcpmV2IamBindingList**](fcpm.v2.IamBindingList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

