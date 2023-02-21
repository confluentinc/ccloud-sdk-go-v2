# \KeysByokV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateByokV1Key**](KeysByokV1Api.md#CreateByokV1Key) | **Post** /byok/v1/keys | Create a Key
[**DeleteByokV1Key**](KeysByokV1Api.md#DeleteByokV1Key) | **Delete** /byok/v1/keys/{id} | Delete a Key
[**GetByokV1Key**](KeysByokV1Api.md#GetByokV1Key) | **Get** /byok/v1/keys/{id} | Read a Key
[**ListByokV1Keys**](KeysByokV1Api.md#ListByokV1Keys) | **Get** /byok/v1/keys | List of Keys



## CreateByokV1Key

> ByokV1Key CreateByokV1Key(ctx).ByokV1Key(byokV1Key).Execute()

Create a Key



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
    byokV1Key := *openapiclient.NewByokV1Key() // ByokV1Key |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeysByokV1Api.CreateByokV1Key(context.Background()).ByokV1Key(byokV1Key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeysByokV1Api.CreateByokV1Key``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateByokV1Key`: ByokV1Key
    fmt.Fprintf(os.Stdout, "Response from `KeysByokV1Api.CreateByokV1Key`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateByokV1KeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **byokV1Key** | [**ByokV1Key**](ByokV1Key.md) |  | 

### Return type

[**ByokV1Key**](byok.v1.Key.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteByokV1Key

> DeleteByokV1Key(ctx, id).Execute()

Delete a Key



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
    id := "id_example" // string | The unique identifier for the key.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeysByokV1Api.DeleteByokV1Key(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeysByokV1Api.DeleteByokV1Key``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteByokV1KeyRequest struct via the builder pattern


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


## GetByokV1Key

> ByokV1Key GetByokV1Key(ctx, id).Execute()

Read a Key



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
    id := "id_example" // string | The unique identifier for the key.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeysByokV1Api.GetByokV1Key(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeysByokV1Api.GetByokV1Key``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetByokV1Key`: ByokV1Key
    fmt.Fprintf(os.Stdout, "Response from `KeysByokV1Api.GetByokV1Key`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByokV1KeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ByokV1Key**](byok.v1.Key.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListByokV1Keys

> ByokV1KeyList ListByokV1Keys(ctx).Provider(provider).State(state).PageSize(pageSize).PageToken(pageToken).Execute()

List of Keys



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
    provider := "AWS" // string | Filter the results by exact match for provider. (optional)
    state := "IN_USE" // string | Filter the results by exact match for state. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeysByokV1Api.ListByokV1Keys(context.Background()).Provider(provider).State(state).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeysByokV1Api.ListByokV1Keys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListByokV1Keys`: ByokV1KeyList
    fmt.Fprintf(os.Stdout, "Response from `KeysByokV1Api.ListByokV1Keys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListByokV1KeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** | Filter the results by exact match for provider. | 
 **state** | **string** | Filter the results by exact match for state. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**ByokV1KeyList**](byok.v1.KeyList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

