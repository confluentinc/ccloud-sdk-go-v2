# \ConfigV1Api

All URIs are relative to *https://psrc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSubjectConfig**](ConfigV1Api.md#DeleteSubjectConfig) | **Delete** /config/{subject} | Delete subject compatibility level
[**DeleteTopLevelConfig**](ConfigV1Api.md#DeleteTopLevelConfig) | **Delete** /config | Delete global compatibility level
[**GetClusterConfig**](ConfigV1Api.md#GetClusterConfig) | **Get** /clusterconfig | Get cluster config
[**GetSubjectLevelConfig**](ConfigV1Api.md#GetSubjectLevelConfig) | **Get** /config/{subject} | Get subject compatibility level
[**GetTopLevelConfig**](ConfigV1Api.md#GetTopLevelConfig) | **Get** /config | Get global compatibility level
[**UpdateSubjectLevelConfig**](ConfigV1Api.md#UpdateSubjectLevelConfig) | **Put** /config/{subject} | Update subject compatibility level
[**UpdateTopLevelConfig**](ConfigV1Api.md#UpdateTopLevelConfig) | **Put** /config | Update global compatibility level



## DeleteSubjectConfig

> string DeleteSubjectConfig(ctx, subject).Execute()

Delete subject compatibility level



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
    subject := "subject_example" // string | Name of the subject

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigV1Api.DeleteSubjectConfig(context.Background(), subject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigV1Api.DeleteSubjectConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSubjectConfig`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigV1Api.DeleteSubjectConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Name of the subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubjectConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTopLevelConfig

> string DeleteTopLevelConfig(ctx).Execute()

Delete global compatibility level



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigV1Api.DeleteTopLevelConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigV1Api.DeleteTopLevelConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTopLevelConfig`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigV1Api.DeleteTopLevelConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTopLevelConfigRequest struct via the builder pattern


### Return type

**string**

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterConfig

> ClusterConfig GetClusterConfig(ctx).Execute()

Get cluster config



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigV1Api.GetClusterConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigV1Api.GetClusterConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterConfig`: ClusterConfig
    fmt.Fprintf(os.Stdout, "Response from `ConfigV1Api.GetClusterConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterConfigRequest struct via the builder pattern


### Return type

[**ClusterConfig**](ClusterConfig.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubjectLevelConfig

> Config GetSubjectLevelConfig(ctx, subject).DefaultToGlobal(defaultToGlobal).Execute()

Get subject compatibility level



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
    subject := "subject_example" // string | Name of the subject
    defaultToGlobal := true // bool | Whether to return the global compatibility level  if subject compatibility level not found (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigV1Api.GetSubjectLevelConfig(context.Background(), subject).DefaultToGlobal(defaultToGlobal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigV1Api.GetSubjectLevelConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubjectLevelConfig`: Config
    fmt.Fprintf(os.Stdout, "Response from `ConfigV1Api.GetSubjectLevelConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Name of the subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubjectLevelConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **defaultToGlobal** | **bool** | Whether to return the global compatibility level  if subject compatibility level not found | 

### Return type

[**Config**](Config.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopLevelConfig

> Config GetTopLevelConfig(ctx).Execute()

Get global compatibility level



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigV1Api.GetTopLevelConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigV1Api.GetTopLevelConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopLevelConfig`: Config
    fmt.Fprintf(os.Stdout, "Response from `ConfigV1Api.GetTopLevelConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopLevelConfigRequest struct via the builder pattern


### Return type

[**Config**](Config.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubjectLevelConfig

> ConfigUpdateRequest UpdateSubjectLevelConfig(ctx, subject).ConfigUpdateRequest(configUpdateRequest).Execute()

Update subject compatibility level



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
    subject := "subject_example" // string | Name of the subject
    configUpdateRequest := *openapiclient.NewConfigUpdateRequest() // ConfigUpdateRequest | Config Update Request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigV1Api.UpdateSubjectLevelConfig(context.Background(), subject).ConfigUpdateRequest(configUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigV1Api.UpdateSubjectLevelConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSubjectLevelConfig`: ConfigUpdateRequest
    fmt.Fprintf(os.Stdout, "Response from `ConfigV1Api.UpdateSubjectLevelConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Name of the subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubjectLevelConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configUpdateRequest** | [**ConfigUpdateRequest**](ConfigUpdateRequest.md) | Config Update Request | 

### Return type

[**ConfigUpdateRequest**](ConfigUpdateRequest.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTopLevelConfig

> ConfigUpdateRequest UpdateTopLevelConfig(ctx).ConfigUpdateRequest(configUpdateRequest).Execute()

Update global compatibility level



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
    configUpdateRequest := *openapiclient.NewConfigUpdateRequest() // ConfigUpdateRequest | Config Update Request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigV1Api.UpdateTopLevelConfig(context.Background()).ConfigUpdateRequest(configUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigV1Api.UpdateTopLevelConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTopLevelConfig`: ConfigUpdateRequest
    fmt.Fprintf(os.Stdout, "Response from `ConfigV1Api.UpdateTopLevelConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTopLevelConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configUpdateRequest** | [**ConfigUpdateRequest**](ConfigUpdateRequest.md) | Config Update Request | 

### Return type

[**ConfigUpdateRequest**](ConfigUpdateRequest.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

