# \ApiKeysIamV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIamV2ApiKey**](ApiKeysIamV2Api.md#CreateIamV2ApiKey) | **Post** /iam/v2/api-keys | Create an Api Key
[**DeleteIamV2ApiKey**](ApiKeysIamV2Api.md#DeleteIamV2ApiKey) | **Delete** /iam/v2/api-keys/{id} | Delete an Api Key
[**GetIamV2ApiKey**](ApiKeysIamV2Api.md#GetIamV2ApiKey) | **Get** /iam/v2/api-keys/{id} | Read an Api Key
[**ListIamV2ApiKeys**](ApiKeysIamV2Api.md#ListIamV2ApiKeys) | **Get** /iam/v2/api-keys | List of Api Keys
[**UpdateIamV2ApiKey**](ApiKeysIamV2Api.md#UpdateIamV2ApiKey) | **Patch** /iam/v2/api-keys/{id} | Update an Api Key



## CreateIamV2ApiKey

> IamV2ApiKey CreateIamV2ApiKey(ctx).IamV2ApiKey(iamV2ApiKey).Execute()

Create an Api Key



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
    iamV2ApiKey := *openapiclient.NewIamV2ApiKey() // IamV2ApiKey |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApiKeysIamV2Api.CreateIamV2ApiKey(context.Background()).IamV2ApiKey(iamV2ApiKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysIamV2Api.CreateIamV2ApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIamV2ApiKey`: IamV2ApiKey
    fmt.Fprintf(os.Stdout, "Response from `ApiKeysIamV2Api.CreateIamV2ApiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIamV2ApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamV2ApiKey** | [**IamV2ApiKey**](IamV2ApiKey.md) |  | 

### Return type

[**IamV2ApiKey**](iam.v2.ApiKey.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIamV2ApiKey

> DeleteIamV2ApiKey(ctx, id).Execute()

Delete an Api Key



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
    id := "id_example" // string | The unique identifier for the api key.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApiKeysIamV2Api.DeleteIamV2ApiKey(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysIamV2Api.DeleteIamV2ApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the api key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamV2ApiKeyRequest struct via the builder pattern


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


## GetIamV2ApiKey

> IamV2ApiKey GetIamV2ApiKey(ctx, id).Execute()

Read an Api Key



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
    id := "id_example" // string | The unique identifier for the api key.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApiKeysIamV2Api.GetIamV2ApiKey(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysIamV2Api.GetIamV2ApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIamV2ApiKey`: IamV2ApiKey
    fmt.Fprintf(os.Stdout, "Response from `ApiKeysIamV2Api.GetIamV2ApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the api key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamV2ApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamV2ApiKey**](iam.v2.ApiKey.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIamV2ApiKeys

> IamV2ApiKeyList ListIamV2ApiKeys(ctx).Owner(owner).Resource(resource).PageSize(pageSize).PageToken(pageToken).Execute()

List of Api Keys



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
    owner := "u-a83k9b" // string | Filter the results by exact match for owner. (optional)
    resource := "resource_example" // string | Filter the results by exact match for resource. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApiKeysIamV2Api.ListIamV2ApiKeys(context.Background()).Owner(owner).Resource(resource).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysIamV2Api.ListIamV2ApiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIamV2ApiKeys`: IamV2ApiKeyList
    fmt.Fprintf(os.Stdout, "Response from `ApiKeysIamV2Api.ListIamV2ApiKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIamV2ApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | Filter the results by exact match for owner. | 
 **resource** | **string** | Filter the results by exact match for resource. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**IamV2ApiKeyList**](IamV2ApiKeyList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIamV2ApiKey

> IamV2ApiKey UpdateIamV2ApiKey(ctx, id).IamV2ApiKeyUpdate(iamV2ApiKeyUpdate).Execute()

Update an Api Key



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
    id := "id_example" // string | The unique identifier for the api key.
    iamV2ApiKeyUpdate := *openapiclient.NewIamV2ApiKeyUpdate() // IamV2ApiKeyUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApiKeysIamV2Api.UpdateIamV2ApiKey(context.Background(), id).IamV2ApiKeyUpdate(iamV2ApiKeyUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysIamV2Api.UpdateIamV2ApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIamV2ApiKey`: IamV2ApiKey
    fmt.Fprintf(os.Stdout, "Response from `ApiKeysIamV2Api.UpdateIamV2ApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the api key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIamV2ApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamV2ApiKeyUpdate** | [**IamV2ApiKeyUpdate**](IamV2ApiKeyUpdate.md) |  | 

### Return type

[**IamV2ApiKey**](iam.v2.ApiKey.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

