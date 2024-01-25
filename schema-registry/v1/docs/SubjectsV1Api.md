# \SubjectsV1Api

All URIs are relative to *https://psrc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSchemaVersion**](SubjectsV1Api.md#DeleteSchemaVersion) | **Delete** /subjects/{subject}/versions/{version} | Delete schema version
[**DeleteSubject**](SubjectsV1Api.md#DeleteSubject) | **Delete** /subjects/{subject} | Delete subject
[**GetReferencedBy**](SubjectsV1Api.md#GetReferencedBy) | **Get** /subjects/{subject}/versions/{version}/referencedby | List schemas referencing a schema
[**GetSchemaByVersion**](SubjectsV1Api.md#GetSchemaByVersion) | **Get** /subjects/{subject}/versions/{version} | Get schema by version
[**GetSchemaOnly1**](SubjectsV1Api.md#GetSchemaOnly1) | **Get** /subjects/{subject}/versions/{version}/schema | Get schema string by version
[**List**](SubjectsV1Api.md#List) | **Get** /subjects | List subjects
[**ListVersions**](SubjectsV1Api.md#ListVersions) | **Get** /subjects/{subject}/versions | List versions under subject
[**LookUpSchemaUnderSubject**](SubjectsV1Api.md#LookUpSchemaUnderSubject) | **Post** /subjects/{subject} | Lookup schema under subject
[**Register**](SubjectsV1Api.md#Register) | **Post** /subjects/{subject}/versions | Register schema under a subject



## DeleteSchemaVersion

> int32 DeleteSchemaVersion(ctx, subject, version).Permanent(permanent).Execute()

Delete schema version



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
    version := "version_example" // string | Version of the schema to be returned. Valid values for versionId are between [1,2^31-1] or the string \"latest\". \"latest\" returns the last registered schema under the specified subject. Note that there may be a new latest schema that gets registered right after this request is served.
    permanent := true // bool | Whether to perform a permanent delete (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubjectsV1Api.DeleteSchemaVersion(context.Background(), subject, version).Permanent(permanent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubjectsV1Api.DeleteSchemaVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSchemaVersion`: int32
    fmt.Fprintf(os.Stdout, "Response from `SubjectsV1Api.DeleteSchemaVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Name of the subject | 
**version** | **string** | Version of the schema to be returned. Valid values for versionId are between [1,2^31-1] or the string \&quot;latest\&quot;. \&quot;latest\&quot; returns the last registered schema under the specified subject. Note that there may be a new latest schema that gets registered right after this request is served. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSchemaVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **permanent** | **bool** | Whether to perform a permanent delete | 

### Return type

**int32**

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubject

> []int32 DeleteSubject(ctx, subject).Permanent(permanent).Execute()

Delete subject



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
    permanent := true // bool | Whether to perform a permanent delete (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubjectsV1Api.DeleteSubject(context.Background(), subject).Permanent(permanent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubjectsV1Api.DeleteSubject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSubject`: []int32
    fmt.Fprintf(os.Stdout, "Response from `SubjectsV1Api.DeleteSubject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Name of the subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **permanent** | **bool** | Whether to perform a permanent delete | 

### Return type

**[]int32**

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReferencedBy

> []int32 GetReferencedBy(ctx, subject, version).Execute()

List schemas referencing a schema



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
    version := "version_example" // string | Version of the schema to be returned. Valid values for versionId are between [1,2^31-1] or the string \"latest\". \"latest\" returns the last registered schema under the specified subject. Note that there may be a new latest schema that gets registered right after this request is served.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubjectsV1Api.GetReferencedBy(context.Background(), subject, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubjectsV1Api.GetReferencedBy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReferencedBy`: []int32
    fmt.Fprintf(os.Stdout, "Response from `SubjectsV1Api.GetReferencedBy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Name of the subject | 
**version** | **string** | Version of the schema to be returned. Valid values for versionId are between [1,2^31-1] or the string \&quot;latest\&quot;. \&quot;latest\&quot; returns the last registered schema under the specified subject. Note that there may be a new latest schema that gets registered right after this request is served. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReferencedByRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**[]int32**

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchemaByVersion

> Schema GetSchemaByVersion(ctx, subject, version).Deleted(deleted).Execute()

Get schema by version



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
    version := "version_example" // string | Version of the schema to be returned. Valid values for versionId are between [1,2^31-1] or the string \"latest\". \"latest\" returns the last registered schema under the specified subject. Note that there may be a new latest schema that gets registered right after this request is served.
    deleted := true // bool | Whether to include deleted schema (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubjectsV1Api.GetSchemaByVersion(context.Background(), subject, version).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubjectsV1Api.GetSchemaByVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemaByVersion`: Schema
    fmt.Fprintf(os.Stdout, "Response from `SubjectsV1Api.GetSchemaByVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Name of the subject | 
**version** | **string** | Version of the schema to be returned. Valid values for versionId are between [1,2^31-1] or the string \&quot;latest\&quot;. \&quot;latest\&quot; returns the last registered schema under the specified subject. Note that there may be a new latest schema that gets registered right after this request is served. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaByVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleted** | **bool** | Whether to include deleted schema | 

### Return type

[**Schema**](Schema.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchemaOnly1

> string GetSchemaOnly1(ctx, subject, version).Deleted(deleted).Execute()

Get schema string by version



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
    version := "version_example" // string | Version of the schema to be returned. Valid values for versionId are between [1,2^31-1] or the string \"latest\". \"latest\" returns the last registered schema under the specified subject. Note that there may be a new latest schema that gets registered right after this request is served.
    deleted := true // bool | Whether to include deleted schema (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubjectsV1Api.GetSchemaOnly1(context.Background(), subject, version).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubjectsV1Api.GetSchemaOnly1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemaOnly1`: string
    fmt.Fprintf(os.Stdout, "Response from `SubjectsV1Api.GetSchemaOnly1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Name of the subject | 
**version** | **string** | Version of the schema to be returned. Valid values for versionId are between [1,2^31-1] or the string \&quot;latest\&quot;. \&quot;latest\&quot; returns the last registered schema under the specified subject. Note that there may be a new latest schema that gets registered right after this request is served. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaOnly1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleted** | **bool** | Whether to include deleted schema | 

### Return type

**string**

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []string List(ctx).SubjectPrefix(subjectPrefix).Deleted(deleted).Execute()

List subjects



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
    subjectPrefix := "subjectPrefix_example" // string | Subject name prefix (optional) (default to ":*:")
    deleted := true // bool | Whether to look up deleted subjects (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubjectsV1Api.List(context.Background()).SubjectPrefix(subjectPrefix).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubjectsV1Api.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: []string
    fmt.Fprintf(os.Stdout, "Response from `SubjectsV1Api.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subjectPrefix** | **string** | Subject name prefix | [default to &quot;:*:&quot;]
 **deleted** | **bool** | Whether to look up deleted subjects | 

### Return type

**[]string**

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVersions

> []int32 ListVersions(ctx, subject).Deleted(deleted).Execute()

List versions under subject



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
    deleted := true // bool | Whether to include deleted schemas (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubjectsV1Api.ListVersions(context.Background(), subject).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubjectsV1Api.ListVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVersions`: []int32
    fmt.Fprintf(os.Stdout, "Response from `SubjectsV1Api.ListVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Name of the subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleted** | **bool** | Whether to include deleted schemas | 

### Return type

**[]int32**

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookUpSchemaUnderSubject

> Schema LookUpSchemaUnderSubject(ctx, subject).RegisterSchemaRequest(registerSchemaRequest).Normalize(normalize).Deleted(deleted).Execute()

Lookup schema under subject



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
    subject := "subject_example" // string | Subject under which the schema will be registered
    registerSchemaRequest := *openapiclient.NewRegisterSchemaRequest() // RegisterSchemaRequest | Schema
    normalize := true // bool | Whether to lookup the normalized schema (optional)
    deleted := true // bool | Whether to lookup deleted schemas (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubjectsV1Api.LookUpSchemaUnderSubject(context.Background(), subject).RegisterSchemaRequest(registerSchemaRequest).Normalize(normalize).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubjectsV1Api.LookUpSchemaUnderSubject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LookUpSchemaUnderSubject`: Schema
    fmt.Fprintf(os.Stdout, "Response from `SubjectsV1Api.LookUpSchemaUnderSubject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Subject under which the schema will be registered | 

### Other Parameters

Other parameters are passed through a pointer to a apiLookUpSchemaUnderSubjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registerSchemaRequest** | [**RegisterSchemaRequest**](RegisterSchemaRequest.md) | Schema | 
 **normalize** | **bool** | Whether to lookup the normalized schema | 
 **deleted** | **bool** | Whether to lookup deleted schemas | 

### Return type

[**Schema**](Schema.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Register

> RegisterSchemaResponse Register(ctx, subject).RegisterSchemaRequest(registerSchemaRequest).Normalize(normalize).Execute()

Register schema under a subject



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
    registerSchemaRequest := *openapiclient.NewRegisterSchemaRequest() // RegisterSchemaRequest | Schema
    normalize := true // bool | Whether to register the normalized schema (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubjectsV1Api.Register(context.Background(), subject).RegisterSchemaRequest(registerSchemaRequest).Normalize(normalize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubjectsV1Api.Register``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Register`: RegisterSchemaResponse
    fmt.Fprintf(os.Stdout, "Response from `SubjectsV1Api.Register`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Name of the subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registerSchemaRequest** | [**RegisterSchemaRequest**](RegisterSchemaRequest.md) | Schema | 
 **normalize** | **bool** | Whether to register the normalized schema | 

### Return type

[**RegisterSchemaResponse**](RegisterSchemaResponse.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

