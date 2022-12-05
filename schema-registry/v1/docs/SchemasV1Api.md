# \SchemasV1Api

All URIs are relative to *https://psrc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSchema**](SchemasV1Api.md#GetSchema) | **Get** /schemas/ids/{id} | Get schema string by ID
[**GetSchemaOnly**](SchemasV1Api.md#GetSchemaOnly) | **Get** /schemas/ids/{id}/schema | Get schema by ID
[**GetSchemaTypes**](SchemasV1Api.md#GetSchemaTypes) | **Get** /schemas/types | List supported schema types
[**GetSchemas**](SchemasV1Api.md#GetSchemas) | **Get** /schemas | List schemas
[**GetSubjects**](SchemasV1Api.md#GetSubjects) | **Get** /schemas/ids/{id}/subjects | List subjects associated to schema ID
[**GetVersions**](SchemasV1Api.md#GetVersions) | **Get** /schemas/ids/{id}/versions | List subject-versions associated to schema ID



## GetSchema

> SchemaString GetSchema(ctx, id).Subject(subject).Format(format).FetchMaxId(fetchMaxId).Execute()

Get schema string by ID



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
    id := int32(56) // int32 | Globally unique identifier of the schema
    subject := "subject_example" // string | Name of the subject (optional)
    format := "format_example" // string | Desired output format, dependent on schema type (optional) (default to "")
    fetchMaxId := true // bool | Whether to fetch the maximum schema identifier that exists (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemasV1Api.GetSchema(context.Background(), id).Subject(subject).Format(format).FetchMaxId(fetchMaxId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasV1Api.GetSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchema`: SchemaString
    fmt.Fprintf(os.Stdout, "Response from `SchemasV1Api.GetSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Globally unique identifier of the schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subject** | **string** | Name of the subject | 
 **format** | **string** | Desired output format, dependent on schema type | [default to &quot;&quot;]
 **fetchMaxId** | **bool** | Whether to fetch the maximum schema identifier that exists | [default to false]

### Return type

[**SchemaString**](SchemaString.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchemaOnly

> string GetSchemaOnly(ctx, id).Subject(subject).Format(format).Execute()

Get schema by ID



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
    id := int32(56) // int32 | Globally unique identifier of the schema
    subject := "subject_example" // string | Name of the subject (optional)
    format := "format_example" // string | Desired output format, dependent on schema type (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemasV1Api.GetSchemaOnly(context.Background(), id).Subject(subject).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasV1Api.GetSchemaOnly``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemaOnly`: string
    fmt.Fprintf(os.Stdout, "Response from `SchemasV1Api.GetSchemaOnly`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Globally unique identifier of the schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaOnlyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subject** | **string** | Name of the subject | 
 **format** | **string** | Desired output format, dependent on schema type | [default to &quot;&quot;]

### Return type

**string**

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchemaTypes

> []string GetSchemaTypes(ctx).Execute()

List supported schema types



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
    resp, r, err := api_client.SchemasV1Api.GetSchemaTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasV1Api.GetSchemaTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemaTypes`: []string
    fmt.Fprintf(os.Stdout, "Response from `SchemasV1Api.GetSchemaTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaTypesRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchemas

> []Schema GetSchemas(ctx).SubjectPrefix(subjectPrefix).Deleted(deleted).LatestOnly(latestOnly).Offset(offset).Limit(limit).Execute()

List schemas



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
    subjectPrefix := "subjectPrefix_example" // string | Filters results by the respective subject prefix (optional) (default to "")
    deleted := true // bool | Whether to return soft deleted schemas (optional) (default to false)
    latestOnly := true // bool | Whether to return latest schema versions only for each matching subject (optional) (default to false)
    offset := int32(56) // int32 | Pagination offset for results (optional) (default to 0)
    limit := int32(56) // int32 | Pagination size for results. Ignored if negative (optional) (default to -1)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemasV1Api.GetSchemas(context.Background()).SubjectPrefix(subjectPrefix).Deleted(deleted).LatestOnly(latestOnly).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasV1Api.GetSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemas`: []Schema
    fmt.Fprintf(os.Stdout, "Response from `SchemasV1Api.GetSchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subjectPrefix** | **string** | Filters results by the respective subject prefix | [default to &quot;&quot;]
 **deleted** | **bool** | Whether to return soft deleted schemas | [default to false]
 **latestOnly** | **bool** | Whether to return latest schema versions only for each matching subject | [default to false]
 **offset** | **int32** | Pagination offset for results | [default to 0]
 **limit** | **int32** | Pagination size for results. Ignored if negative | [default to -1]

### Return type

[**[]Schema**](Schema.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubjects

> []string GetSubjects(ctx, id).Subject(subject).Deleted(deleted).Execute()

List subjects associated to schema ID



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
    id := int32(56) // int32 | Globally unique identifier of the schema
    subject := "subject_example" // string | Filters results by the respective subject (optional)
    deleted := true // bool | Whether to include subjects where the schema was deleted (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemasV1Api.GetSubjects(context.Background(), id).Subject(subject).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasV1Api.GetSubjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `SchemasV1Api.GetSubjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Globally unique identifier of the schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subject** | **string** | Filters results by the respective subject | 
 **deleted** | **bool** | Whether to include subjects where the schema was deleted | 

### Return type

**[]string**

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersions

> []SubjectVersion GetVersions(ctx, id).Subject(subject).Deleted(deleted).Execute()

List subject-versions associated to schema ID



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
    id := int32(56) // int32 | Globally unique identifier of the schema
    subject := "subject_example" // string | Filters results by the respective subject (optional)
    deleted := true // bool | Whether to include subject versions where the schema was deleted (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemasV1Api.GetVersions(context.Background(), id).Subject(subject).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasV1Api.GetVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVersions`: []SubjectVersion
    fmt.Fprintf(os.Stdout, "Response from `SchemasV1Api.GetVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Globally unique identifier of the schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subject** | **string** | Filters results by the respective subject | 
 **deleted** | **bool** | Whether to include subject versions where the schema was deleted | 

### Return type

[**[]SubjectVersion**](SubjectVersion.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

