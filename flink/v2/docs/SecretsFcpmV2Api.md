# \SecretsFcpmV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFcpmV2Secret**](SecretsFcpmV2Api.md#CreateFcpmV2Secret) | **Post** /fcpm/v2/secrets | Create a Secret
[**DeleteFcpmV2Secret**](SecretsFcpmV2Api.md#DeleteFcpmV2Secret) | **Delete** /fcpm/v2/secrets/{id} | Delete a Secret
[**GetFcpmV2Secret**](SecretsFcpmV2Api.md#GetFcpmV2Secret) | **Get** /fcpm/v2/secrets/{id} | Read a Secret
[**ListFcpmV2Secrets**](SecretsFcpmV2Api.md#ListFcpmV2Secrets) | **Get** /fcpm/v2/secrets | List of Secrets
[**UpdateFcpmV2Secret**](SecretsFcpmV2Api.md#UpdateFcpmV2Secret) | **Patch** /fcpm/v2/secrets/{id} | Update a Secret



## CreateFcpmV2Secret

> FcpmV2Secret CreateFcpmV2Secret(ctx).FcpmV2Secret(fcpmV2Secret).Execute()

Create a Secret



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
    fcpmV2Secret := *openapiclient.NewFcpmV2Secret() // FcpmV2Secret |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecretsFcpmV2Api.CreateFcpmV2Secret(context.Background()).FcpmV2Secret(fcpmV2Secret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsFcpmV2Api.CreateFcpmV2Secret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFcpmV2Secret`: FcpmV2Secret
    fmt.Fprintf(os.Stdout, "Response from `SecretsFcpmV2Api.CreateFcpmV2Secret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFcpmV2SecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fcpmV2Secret** | [**FcpmV2Secret**](FcpmV2Secret.md) |  | 

### Return type

[**FcpmV2Secret**](fcpm.v2.Secret.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFcpmV2Secret

> DeleteFcpmV2Secret(ctx, id).Execute()

Delete a Secret



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
    id := "id_example" // string | The unique identifier for the secret.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecretsFcpmV2Api.DeleteFcpmV2Secret(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsFcpmV2Api.DeleteFcpmV2Secret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFcpmV2SecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFcpmV2Secret

> FcpmV2Secret GetFcpmV2Secret(ctx, id).Execute()

Read a Secret



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
    id := "id_example" // string | The unique identifier for the secret.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecretsFcpmV2Api.GetFcpmV2Secret(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsFcpmV2Api.GetFcpmV2Secret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFcpmV2Secret`: FcpmV2Secret
    fmt.Fprintf(os.Stdout, "Response from `SecretsFcpmV2Api.GetFcpmV2Secret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFcpmV2SecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FcpmV2Secret**](fcpm.v2.Secret.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFcpmV2Secrets

> FcpmV2SecretList ListFcpmV2Secrets(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Secrets



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
    resp, r, err := api_client.SecretsFcpmV2Api.ListFcpmV2Secrets(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsFcpmV2Api.ListFcpmV2Secrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFcpmV2Secrets`: FcpmV2SecretList
    fmt.Fprintf(os.Stdout, "Response from `SecretsFcpmV2Api.ListFcpmV2Secrets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFcpmV2SecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**FcpmV2SecretList**](fcpm.v2.SecretList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFcpmV2Secret

> FcpmV2Secret UpdateFcpmV2Secret(ctx, id).FcpmV2SecretUpdate(fcpmV2SecretUpdate).Execute()

Update a Secret



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
    id := "id_example" // string | The unique identifier for the secret.
    fcpmV2SecretUpdate := *openapiclient.NewFcpmV2SecretUpdate() // FcpmV2SecretUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecretsFcpmV2Api.UpdateFcpmV2Secret(context.Background(), id).FcpmV2SecretUpdate(fcpmV2SecretUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsFcpmV2Api.UpdateFcpmV2Secret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFcpmV2Secret`: FcpmV2Secret
    fmt.Fprintf(os.Stdout, "Response from `SecretsFcpmV2Api.UpdateFcpmV2Secret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFcpmV2SecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fcpmV2SecretUpdate** | [**FcpmV2SecretUpdate**](FcpmV2SecretUpdate.md) |  | 

### Return type

[**FcpmV2Secret**](fcpm.v2.Secret.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

