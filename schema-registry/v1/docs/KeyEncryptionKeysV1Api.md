# \KeyEncryptionKeysV1Api

All URIs are relative to *https://psrc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKek**](KeyEncryptionKeysV1Api.md#CreateKek) | **Post** /dek-registry/v1/keks | Create a kek.
[**DeleteKek**](KeyEncryptionKeysV1Api.md#DeleteKek) | **Delete** /dek-registry/v1/keks/{name} | Delete a kek.
[**GetKek**](KeyEncryptionKeysV1Api.md#GetKek) | **Get** /dek-registry/v1/keks/{name} | Get a kek by name.
[**GetKekNames**](KeyEncryptionKeysV1Api.md#GetKekNames) | **Get** /dek-registry/v1/keks | Get a list of kek names.
[**PutKek**](KeyEncryptionKeysV1Api.md#PutKek) | **Put** /dek-registry/v1/keks/{name} | Alters a kek.
[**UndeleteKek**](KeyEncryptionKeysV1Api.md#UndeleteKek) | **Post** /dek-registry/v1/keks/{name}/undelete | Undelete a kek.



## CreateKek

> Kek CreateKek(ctx).CreateKekRequest(createKekRequest).Execute()

Create a kek.

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
    createKekRequest := *openapiclient.NewCreateKekRequest() // CreateKekRequest | The create request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeyEncryptionKeysV1Api.CreateKek(context.Background()).CreateKekRequest(createKekRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyEncryptionKeysV1Api.CreateKek``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKek`: Kek
    fmt.Fprintf(os.Stdout, "Response from `KeyEncryptionKeysV1Api.CreateKek`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKekRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createKekRequest** | [**CreateKekRequest**](CreateKekRequest.md) | The create request | 

### Return type

[**Kek**](Kek.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKek

> DeleteKek(ctx, name).Permanent(permanent).Execute()

Delete a kek.

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
    name := "name_example" // string | Name of the kek
    permanent := true // bool | Whether to perform a permanent delete (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeyEncryptionKeysV1Api.DeleteKek(context.Background(), name).Permanent(permanent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyEncryptionKeysV1Api.DeleteKek``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKekRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **permanent** | **bool** | Whether to perform a permanent delete | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKek

> Kek GetKek(ctx, name).Deleted(deleted).Execute()

Get a kek by name.

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
    name := "name_example" // string | Name of the kek
    deleted := true // bool | Whether to include deleted keys (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeyEncryptionKeysV1Api.GetKek(context.Background(), name).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyEncryptionKeysV1Api.GetKek``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKek`: Kek
    fmt.Fprintf(os.Stdout, "Response from `KeyEncryptionKeysV1Api.GetKek`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKekRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleted** | **bool** | Whether to include deleted keys | 

### Return type

[**Kek**](Kek.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKekNames

> []string GetKekNames(ctx).Deleted(deleted).Execute()

Get a list of kek names.

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
    deleted := true // bool | Whether to include deleted keys (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeyEncryptionKeysV1Api.GetKekNames(context.Background()).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyEncryptionKeysV1Api.GetKekNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKekNames`: []string
    fmt.Fprintf(os.Stdout, "Response from `KeyEncryptionKeysV1Api.GetKekNames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKekNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleted** | **bool** | Whether to include deleted keys | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutKek

> Kek PutKek(ctx, name).UpdateKekRequest(updateKekRequest).Execute()

Alters a kek.

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
    name := "name_example" // string | Name of the kek
    updateKekRequest := *openapiclient.NewUpdateKekRequest() // UpdateKekRequest | The update request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeyEncryptionKeysV1Api.PutKek(context.Background(), name).UpdateKekRequest(updateKekRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyEncryptionKeysV1Api.PutKek``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutKek`: Kek
    fmt.Fprintf(os.Stdout, "Response from `KeyEncryptionKeysV1Api.PutKek`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutKekRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateKekRequest** | [**UpdateKekRequest**](UpdateKekRequest.md) | The update request | 

### Return type

[**Kek**](Kek.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UndeleteKek

> UndeleteKek(ctx, name).Execute()

Undelete a kek.

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
    name := "name_example" // string | Name of the kek

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeyEncryptionKeysV1Api.UndeleteKek(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyEncryptionKeysV1Api.UndeleteKek``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 

### Other Parameters

Other parameters are passed through a pointer to a apiUndeleteKekRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

