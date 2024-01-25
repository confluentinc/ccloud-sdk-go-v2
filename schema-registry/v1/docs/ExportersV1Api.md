# \ExportersV1Api

All URIs are relative to *https://psrc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteExporter**](ExportersV1Api.md#DeleteExporter) | **Delete** /exporters/{name} | Delete schema exporter by name.
[**GetExporterConfigByName**](ExportersV1Api.md#GetExporterConfigByName) | **Get** /exporters/{name}/config | Gets schema exporter config by name.
[**GetExporterInfoByName**](ExportersV1Api.md#GetExporterInfoByName) | **Get** /exporters/{name} | Gets schema exporter by name.
[**GetExporterStatusByName**](ExportersV1Api.md#GetExporterStatusByName) | **Get** /exporters/{name}/status | Gets schema exporter status by name.
[**ListExporters**](ExportersV1Api.md#ListExporters) | **Get** /exporters | Gets all schema exporters.
[**PauseExporterByName**](ExportersV1Api.md#PauseExporterByName) | **Put** /exporters/{name}/pause | Pause schema exporter by name.
[**RegisterExporter**](ExportersV1Api.md#RegisterExporter) | **Post** /exporters | Creates a new schema exporter.
[**ResetExporterByName**](ExportersV1Api.md#ResetExporterByName) | **Put** /exporters/{name}/reset | Reset schema exporter by name.
[**ResumeExporterByName**](ExportersV1Api.md#ResumeExporterByName) | **Put** /exporters/{name}/resume | Resume schema exporter by name.
[**UpdateExporterConfigByName**](ExportersV1Api.md#UpdateExporterConfigByName) | **Put** /exporters/{name}/config | Update schema exporter config by name.
[**UpdateExporterInfo**](ExportersV1Api.md#UpdateExporterInfo) | **Put** /exporters/{name} | Update schema exporter by name.



## DeleteExporter

> DeleteExporter(ctx, name).Execute()

Delete schema exporter by name.



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
    name := "name_example" // string | Name of the exporter

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersV1Api.DeleteExporter(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersV1Api.DeleteExporter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the exporter | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExporterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExporterConfigByName

> ExporterConfigResponse GetExporterConfigByName(ctx, name).Execute()

Gets schema exporter config by name.



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
    name := "name_example" // string | Name of the exporter

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersV1Api.GetExporterConfigByName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersV1Api.GetExporterConfigByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExporterConfigByName`: ExporterConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersV1Api.GetExporterConfigByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the exporter | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExporterConfigByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExporterConfigResponse**](ExporterConfigResponse.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExporterInfoByName

> ExporterReference GetExporterInfoByName(ctx, name).Execute()

Gets schema exporter by name.



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
    name := "name_example" // string | Name of the exporter

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersV1Api.GetExporterInfoByName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersV1Api.GetExporterInfoByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExporterInfoByName`: ExporterReference
    fmt.Fprintf(os.Stdout, "Response from `ExportersV1Api.GetExporterInfoByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the exporter | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExporterInfoByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExporterReference**](ExporterReference.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExporterStatusByName

> ExporterStatusResponse GetExporterStatusByName(ctx, name).Execute()

Gets schema exporter status by name.



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
    name := "name_example" // string | Name of the exporter

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersV1Api.GetExporterStatusByName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersV1Api.GetExporterStatusByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExporterStatusByName`: ExporterStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersV1Api.GetExporterStatusByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the exporter | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExporterStatusByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExporterStatusResponse**](ExporterStatusResponse.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExporters

> []string ListExporters(ctx).Execute()

Gets all schema exporters.



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
    resp, r, err := api_client.ExportersV1Api.ListExporters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersV1Api.ListExporters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExporters`: []string
    fmt.Fprintf(os.Stdout, "Response from `ExportersV1Api.ListExporters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListExportersRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseExporterByName

> ExporterResponse PauseExporterByName(ctx, name).Execute()

Pause schema exporter by name.



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
    name := "name_example" // string | Name of the exporter

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersV1Api.PauseExporterByName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersV1Api.PauseExporterByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PauseExporterByName`: ExporterResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersV1Api.PauseExporterByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the exporter | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseExporterByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExporterResponse**](ExporterResponse.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterExporter

> ExporterResponse RegisterExporter(ctx).ExporterReference(exporterReference).Execute()

Creates a new schema exporter.



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
    exporterReference := *openapiclient.NewExporterReference() // ExporterReference | Schema

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersV1Api.RegisterExporter(context.Background()).ExporterReference(exporterReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersV1Api.RegisterExporter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterExporter`: ExporterResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersV1Api.RegisterExporter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterExporterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exporterReference** | [**ExporterReference**](ExporterReference.md) | Schema | 

### Return type

[**ExporterResponse**](ExporterResponse.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json
- **Accept**: application/json; qs=0.5, application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetExporterByName

> ExporterResponse ResetExporterByName(ctx, name).Execute()

Reset schema exporter by name.



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
    name := "name_example" // string | Name of the exporter

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersV1Api.ResetExporterByName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersV1Api.ResetExporterByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetExporterByName`: ExporterResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersV1Api.ResetExporterByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the exporter | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetExporterByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExporterResponse**](ExporterResponse.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeExporterByName

> ExporterResponse ResumeExporterByName(ctx, name).Execute()

Resume schema exporter by name.



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
    name := "name_example" // string | Name of the exporter

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersV1Api.ResumeExporterByName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersV1Api.ResumeExporterByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeExporterByName`: ExporterResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersV1Api.ResumeExporterByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the exporter | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeExporterByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExporterResponse**](ExporterResponse.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExporterConfigByName

> ExporterResponse UpdateExporterConfigByName(ctx, name).ExporterConfigResponse(exporterConfigResponse).Execute()

Update schema exporter config by name.



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
    name := "name_example" // string | Name of the exporter
    exporterConfigResponse := *openapiclient.NewExporterConfigResponse() // ExporterConfigResponse | Exporter Update Request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersV1Api.UpdateExporterConfigByName(context.Background(), name).ExporterConfigResponse(exporterConfigResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersV1Api.UpdateExporterConfigByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExporterConfigByName`: ExporterResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersV1Api.UpdateExporterConfigByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the exporter | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExporterConfigByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exporterConfigResponse** | [**ExporterConfigResponse**](ExporterConfigResponse.md) | Exporter Update Request | 

### Return type

[**ExporterResponse**](ExporterResponse.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json
- **Accept**: application/vnd.schemaregistry.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExporterInfo

> ExporterResponse UpdateExporterInfo(ctx, name).ExporterUpdateRequest(exporterUpdateRequest).Execute()

Update schema exporter by name.



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
    name := "name_example" // string | Name of the exporter
    exporterUpdateRequest := *openapiclient.NewExporterUpdateRequest() // ExporterUpdateRequest | Exporter Update Request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersV1Api.UpdateExporterInfo(context.Background(), name).ExporterUpdateRequest(exporterUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersV1Api.UpdateExporterInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExporterInfo`: ExporterResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersV1Api.UpdateExporterInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the exporter | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExporterInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exporterUpdateRequest** | [**ExporterUpdateRequest**](ExporterUpdateRequest.md) | Exporter Update Request | 

### Return type

[**ExporterResponse**](ExporterResponse.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json
- **Accept**: application/vnd.schemaregistry.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

