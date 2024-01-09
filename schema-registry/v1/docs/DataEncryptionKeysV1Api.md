# \DataEncryptionKeysV1Api

All URIs are relative to *https://psrc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDek**](DataEncryptionKeysV1Api.md#CreateDek) | **Post** /dek-registry/v1/keks/{name}/deks | Create a dek.
[**DeleteDekVersion**](DataEncryptionKeysV1Api.md#DeleteDekVersion) | **Delete** /dek-registry/v1/keks/{name}/deks/{subject}/versions/{version} | Delete a dek version.
[**DeleteDekVersions**](DataEncryptionKeysV1Api.md#DeleteDekVersions) | **Delete** /dek-registry/v1/keks/{name}/deks/{subject} | Delete all versions of a dek.
[**GetDek**](DataEncryptionKeysV1Api.md#GetDek) | **Get** /dek-registry/v1/keks/{name}/deks/{subject} | Get a dek by subject.
[**GetDekByVersion**](DataEncryptionKeysV1Api.md#GetDekByVersion) | **Get** /dek-registry/v1/keks/{name}/deks/{subject}/versions/{version} | Get a dek by subject and version.
[**GetDekSubjects**](DataEncryptionKeysV1Api.md#GetDekSubjects) | **Get** /dek-registry/v1/keks/{name}/deks | Get a list of dek subjects.
[**GetDekVersions**](DataEncryptionKeysV1Api.md#GetDekVersions) | **Get** /dek-registry/v1/keks/{name}/deks/{subject}/versions | List versions of dek.
[**UndeleteDekVersion**](DataEncryptionKeysV1Api.md#UndeleteDekVersion) | **Post** /dek-registry/v1/keks/{name}/deks/{subject}/versions/{version}/undelete | Undelete a dek version.
[**UndeleteDekVersions**](DataEncryptionKeysV1Api.md#UndeleteDekVersions) | **Post** /dek-registry/v1/keks/{name}/deks/{subject}/undelete | Undelete all versions of a dek.



## CreateDek

> Dek CreateDek(ctx, name).CreateDekRequest(createDekRequest).Execute()

Create a dek.

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
    createDekRequest := *openapiclient.NewCreateDekRequest() // CreateDekRequest | The create request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataEncryptionKeysV1Api.CreateDek(context.Background(), name).CreateDekRequest(createDekRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataEncryptionKeysV1Api.CreateDek``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDek`: Dek
    fmt.Fprintf(os.Stdout, "Response from `DataEncryptionKeysV1Api.CreateDek`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDekRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDekRequest** | [**CreateDekRequest**](CreateDekRequest.md) | The create request | 

### Return type

[**Dek**](Dek.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDekVersion

> DeleteDekVersion(ctx, name, subject, version).Algorithm(algorithm).Permanent(permanent).Execute()

Delete a dek version.

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
    subject := "subject_example" // string | Subject of the dek
    version := "version_example" // string | Version of the dek
    algorithm := "algorithm_example" // string | Algorithm of the dek (optional)
    permanent := true // bool | Whether to perform a permanent delete (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataEncryptionKeysV1Api.DeleteDekVersion(context.Background(), name, subject, version).Algorithm(algorithm).Permanent(permanent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataEncryptionKeysV1Api.DeleteDekVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 
**subject** | **string** | Subject of the dek | 
**version** | **string** | Version of the dek | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDekVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **algorithm** | **string** | Algorithm of the dek | 
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


## DeleteDekVersions

> DeleteDekVersions(ctx, name, subject).Algorithm(algorithm).Permanent(permanent).Execute()

Delete all versions of a dek.

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
    subject := "subject_example" // string | Subject of the dek
    algorithm := "algorithm_example" // string | Algorithm of the dek (optional)
    permanent := true // bool | Whether to perform a permanent delete (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataEncryptionKeysV1Api.DeleteDekVersions(context.Background(), name, subject).Algorithm(algorithm).Permanent(permanent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataEncryptionKeysV1Api.DeleteDekVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 
**subject** | **string** | Subject of the dek | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDekVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **algorithm** | **string** | Algorithm of the dek | 
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


## GetDek

> Dek GetDek(ctx, name, subject).Algorithm(algorithm).Deleted(deleted).Execute()

Get a dek by subject.

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
    subject := "subject_example" // string | Subject of the dek
    algorithm := "algorithm_example" // string | Algorithm of the dek (optional)
    deleted := true // bool | Whether to include deleted keys (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataEncryptionKeysV1Api.GetDek(context.Background(), name, subject).Algorithm(algorithm).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataEncryptionKeysV1Api.GetDek``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDek`: Dek
    fmt.Fprintf(os.Stdout, "Response from `DataEncryptionKeysV1Api.GetDek`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 
**subject** | **string** | Subject of the dek | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDekRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **algorithm** | **string** | Algorithm of the dek | 
 **deleted** | **bool** | Whether to include deleted keys | 

### Return type

[**Dek**](Dek.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDekByVersion

> Dek GetDekByVersion(ctx, name, subject, version).Algorithm(algorithm).Deleted(deleted).Execute()

Get a dek by subject and version.

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
    subject := "subject_example" // string | Subject of the dek
    version := "version_example" // string | Version of the dek
    algorithm := "algorithm_example" // string | Algorithm of the dek (optional)
    deleted := true // bool | Whether to include deleted keys (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataEncryptionKeysV1Api.GetDekByVersion(context.Background(), name, subject, version).Algorithm(algorithm).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataEncryptionKeysV1Api.GetDekByVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDekByVersion`: Dek
    fmt.Fprintf(os.Stdout, "Response from `DataEncryptionKeysV1Api.GetDekByVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 
**subject** | **string** | Subject of the dek | 
**version** | **string** | Version of the dek | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDekByVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **algorithm** | **string** | Algorithm of the dek | 
 **deleted** | **bool** | Whether to include deleted keys | 

### Return type

[**Dek**](Dek.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDekSubjects

> []string GetDekSubjects(ctx, name).Deleted(deleted).Execute()

Get a list of dek subjects.

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
    resp, r, err := api_client.DataEncryptionKeysV1Api.GetDekSubjects(context.Background(), name).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataEncryptionKeysV1Api.GetDekSubjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDekSubjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `DataEncryptionKeysV1Api.GetDekSubjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDekSubjectsRequest struct via the builder pattern


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


## GetDekVersions

> []int32 GetDekVersions(ctx, name, subject).Algorithm(algorithm).Deleted(deleted).Execute()

List versions of dek.

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
    subject := "subject_example" // string | Subject of the dek
    algorithm := "algorithm_example" // string | Algorithm of the dek (optional)
    deleted := true // bool | Whether to include deleted keys (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataEncryptionKeysV1Api.GetDekVersions(context.Background(), name, subject).Algorithm(algorithm).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataEncryptionKeysV1Api.GetDekVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDekVersions`: []int32
    fmt.Fprintf(os.Stdout, "Response from `DataEncryptionKeysV1Api.GetDekVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 
**subject** | **string** | Subject of the dek | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDekVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **algorithm** | **string** | Algorithm of the dek | 
 **deleted** | **bool** | Whether to include deleted keys | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UndeleteDekVersion

> UndeleteDekVersion(ctx, name, subject, version).Algorithm(algorithm).Execute()

Undelete a dek version.

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
    subject := "subject_example" // string | Subject of the dek
    version := "version_example" // string | Version of the dek
    algorithm := "algorithm_example" // string | Algorithm of the dek (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataEncryptionKeysV1Api.UndeleteDekVersion(context.Background(), name, subject, version).Algorithm(algorithm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataEncryptionKeysV1Api.UndeleteDekVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 
**subject** | **string** | Subject of the dek | 
**version** | **string** | Version of the dek | 

### Other Parameters

Other parameters are passed through a pointer to a apiUndeleteDekVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **algorithm** | **string** | Algorithm of the dek | 

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


## UndeleteDekVersions

> UndeleteDekVersions(ctx, name, subject).Algorithm(algorithm).Execute()

Undelete all versions of a dek.

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
    subject := "subject_example" // string | Subject of the dek
    algorithm := "algorithm_example" // string | Algorithm of the dek (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataEncryptionKeysV1Api.UndeleteDekVersions(context.Background(), name, subject).Algorithm(algorithm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataEncryptionKeysV1Api.UndeleteDekVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 
**subject** | **string** | Subject of the dek | 

### Other Parameters

Other parameters are passed through a pointer to a apiUndeleteDekVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **algorithm** | **string** | Algorithm of the dek | 

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

