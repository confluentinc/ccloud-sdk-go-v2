# \IdentityProvidersIamV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIamV2IdentityProvider**](IdentityProvidersIamV2Api.md#CreateIamV2IdentityProvider) | **Post** /iam/v2/identity-providers | Create an Identity Provider
[**DeleteIamV2IdentityProvider**](IdentityProvidersIamV2Api.md#DeleteIamV2IdentityProvider) | **Delete** /iam/v2/identity-providers/{id} | Delete an Identity Provider
[**GetIamV2IdentityProvider**](IdentityProvidersIamV2Api.md#GetIamV2IdentityProvider) | **Get** /iam/v2/identity-providers/{id} | Read an Identity Provider
[**ListIamV2IdentityProviders**](IdentityProvidersIamV2Api.md#ListIamV2IdentityProviders) | **Get** /iam/v2/identity-providers | List of Identity Providers
[**UpdateIamV2IdentityProvider**](IdentityProvidersIamV2Api.md#UpdateIamV2IdentityProvider) | **Patch** /iam/v2/identity-providers/{id} | Update an Identity Provider



## CreateIamV2IdentityProvider

> IamV2IdentityProvider CreateIamV2IdentityProvider(ctx).IamV2IdentityProvider(iamV2IdentityProvider).Execute()

Create an Identity Provider



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
    iamV2IdentityProvider := *openapiclient.NewIamV2IdentityProvider() // IamV2IdentityProvider |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityProvidersIamV2Api.CreateIamV2IdentityProvider(context.Background()).IamV2IdentityProvider(iamV2IdentityProvider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersIamV2Api.CreateIamV2IdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIamV2IdentityProvider`: IamV2IdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersIamV2Api.CreateIamV2IdentityProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIamV2IdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamV2IdentityProvider** | [**IamV2IdentityProvider**](IamV2IdentityProvider.md) |  | 

### Return type

[**IamV2IdentityProvider**](iam.v2.IdentityProvider.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIamV2IdentityProvider

> DeleteIamV2IdentityProvider(ctx, id).Execute()

Delete an Identity Provider



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
    id := "id_example" // string | The unique identifier for the identity provider.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityProvidersIamV2Api.DeleteIamV2IdentityProvider(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersIamV2Api.DeleteIamV2IdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the identity provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamV2IdentityProviderRequest struct via the builder pattern


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


## GetIamV2IdentityProvider

> IamV2IdentityProvider GetIamV2IdentityProvider(ctx, id).Execute()

Read an Identity Provider



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
    id := "id_example" // string | The unique identifier for the identity provider.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityProvidersIamV2Api.GetIamV2IdentityProvider(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersIamV2Api.GetIamV2IdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIamV2IdentityProvider`: IamV2IdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersIamV2Api.GetIamV2IdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the identity provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamV2IdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamV2IdentityProvider**](iam.v2.IdentityProvider.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIamV2IdentityProviders

> IamV2IdentityProviderList ListIamV2IdentityProviders(ctx).Deactivated(deactivated).PageSize(pageSize).PageToken(pageToken).Execute()

List of Identity Providers



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
    deactivated := "deactivated_example" // string | Filter the results by exact match for deactivated. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityProvidersIamV2Api.ListIamV2IdentityProviders(context.Background()).Deactivated(deactivated).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersIamV2Api.ListIamV2IdentityProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIamV2IdentityProviders`: IamV2IdentityProviderList
    fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersIamV2Api.ListIamV2IdentityProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIamV2IdentityProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deactivated** | **string** | Filter the results by exact match for deactivated. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**IamV2IdentityProviderList**](iam.v2.IdentityProviderList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIamV2IdentityProvider

> IamV2IdentityProvider UpdateIamV2IdentityProvider(ctx, id).IamV2IdentityProviderUpdate(iamV2IdentityProviderUpdate).Execute()

Update an Identity Provider



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
    id := "id_example" // string | The unique identifier for the identity provider.
    iamV2IdentityProviderUpdate := *openapiclient.NewIamV2IdentityProviderUpdate() // IamV2IdentityProviderUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityProvidersIamV2Api.UpdateIamV2IdentityProvider(context.Background(), id).IamV2IdentityProviderUpdate(iamV2IdentityProviderUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersIamV2Api.UpdateIamV2IdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIamV2IdentityProvider`: IamV2IdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersIamV2Api.UpdateIamV2IdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the identity provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIamV2IdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamV2IdentityProviderUpdate** | [**IamV2IdentityProviderUpdate**](IamV2IdentityProviderUpdate.md) |  | 

### Return type

[**IamV2IdentityProvider**](iam.v2.IdentityProvider.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

