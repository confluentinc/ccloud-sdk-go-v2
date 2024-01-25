# \CompatibilityV1Api

All URIs are relative to *https://psrc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TestCompatibilityBySubjectName**](CompatibilityV1Api.md#TestCompatibilityBySubjectName) | **Post** /compatibility/subjects/{subject}/versions/{version} | Test schema compatibility against a particular schema subject-version
[**TestCompatibilityForSubject**](CompatibilityV1Api.md#TestCompatibilityForSubject) | **Post** /compatibility/subjects/{subject}/versions | Test schema compatibility against all schemas under a subject



## TestCompatibilityBySubjectName

> CompatibilityCheckResponse TestCompatibilityBySubjectName(ctx, subject, version).RegisterSchemaRequest(registerSchemaRequest).Verbose(verbose).Execute()

Test schema compatibility against a particular schema subject-version



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
    subject := "subject_example" // string | Subject of the schema version against which compatibility is to be tested
    version := "version_example" // string | Version of the subject's schema against which compatibility is to be tested. Valid values for versionId are between [1,2^31-1] or the string \"latest\".\"latest\" checks compatibility of the input schema with the last registered schema under the specified subject
    registerSchemaRequest := *openapiclient.NewRegisterSchemaRequest() // RegisterSchemaRequest | Schema
    verbose := true // bool | Whether to return detailed error messages (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompatibilityV1Api.TestCompatibilityBySubjectName(context.Background(), subject, version).RegisterSchemaRequest(registerSchemaRequest).Verbose(verbose).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompatibilityV1Api.TestCompatibilityBySubjectName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestCompatibilityBySubjectName`: CompatibilityCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `CompatibilityV1Api.TestCompatibilityBySubjectName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Subject of the schema version against which compatibility is to be tested | 
**version** | **string** | Version of the subject&#39;s schema against which compatibility is to be tested. Valid values for versionId are between [1,2^31-1] or the string \&quot;latest\&quot;.\&quot;latest\&quot; checks compatibility of the input schema with the last registered schema under the specified subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestCompatibilityBySubjectNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **registerSchemaRequest** | [**RegisterSchemaRequest**](RegisterSchemaRequest.md) | Schema | 
 **verbose** | **bool** | Whether to return detailed error messages | 

### Return type

[**CompatibilityCheckResponse**](CompatibilityCheckResponse.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestCompatibilityForSubject

> CompatibilityCheckResponse TestCompatibilityForSubject(ctx, subject).RegisterSchemaRequest(registerSchemaRequest).Verbose(verbose).Execute()

Test schema compatibility against all schemas under a subject



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
    subject := "subject_example" // string | Subject of the schema version against which compatibility is to be tested
    registerSchemaRequest := *openapiclient.NewRegisterSchemaRequest() // RegisterSchemaRequest | Schema
    verbose := true // bool | Whether to return detailed error messages (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompatibilityV1Api.TestCompatibilityForSubject(context.Background(), subject).RegisterSchemaRequest(registerSchemaRequest).Verbose(verbose).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompatibilityV1Api.TestCompatibilityForSubject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestCompatibilityForSubject`: CompatibilityCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `CompatibilityV1Api.TestCompatibilityForSubject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Subject of the schema version against which compatibility is to be tested | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestCompatibilityForSubjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registerSchemaRequest** | [**RegisterSchemaRequest**](RegisterSchemaRequest.md) | Schema | 
 **verbose** | **bool** | Whether to return detailed error messages | 

### Return type

[**CompatibilityCheckResponse**](CompatibilityCheckResponse.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

