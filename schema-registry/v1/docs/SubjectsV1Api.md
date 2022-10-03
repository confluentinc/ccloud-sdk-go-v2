# \SubjectsV1Api

All URIs are relative to *https://psrc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSubject**](SubjectsV1Api.md#DeleteSubject) | **Delete** /subjects/{subject} | Delete subject
[**List**](SubjectsV1Api.md#List) | **Get** /subjects | List subjects
[**LookUpSchemaUnderSubject**](SubjectsV1Api.md#LookUpSchemaUnderSubject) | **Post** /subjects/{subject} | Lookup schema under subject



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

[api-key](../README.md#api-key)

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

[api-key](../README.md#api-key)

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

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

