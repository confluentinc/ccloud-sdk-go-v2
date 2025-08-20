# \IntegrationsPimV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePimV2Integration**](IntegrationsPimV2Api.md#CreatePimV2Integration) | **Post** /pim/v2/integrations | Create an Integration
[**DeletePimV2Integration**](IntegrationsPimV2Api.md#DeletePimV2Integration) | **Delete** /pim/v2/integrations/{id} | Delete an Integration
[**GetPimV2Integration**](IntegrationsPimV2Api.md#GetPimV2Integration) | **Get** /pim/v2/integrations/{id} | Read an Integration
[**ListPimV2Integrations**](IntegrationsPimV2Api.md#ListPimV2Integrations) | **Get** /pim/v2/integrations | List of Integrations
[**UpdatePimV2Integration**](IntegrationsPimV2Api.md#UpdatePimV2Integration) | **Patch** /pim/v2/integrations/{id} | Update an Integration
[**ValidatePimV2Integration**](IntegrationsPimV2Api.md#ValidatePimV2Integration) | **Post** /pim/v2/integrations:validate | Validate an Integration



## CreatePimV2Integration

> PimV2Integration CreatePimV2Integration(ctx).PimV2Integration(pimV2Integration).Execute()

Create an Integration



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
    pimV2Integration := *openapiclient.NewPimV2Integration() // PimV2Integration |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsPimV2Api.CreatePimV2Integration(context.Background()).PimV2Integration(pimV2Integration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsPimV2Api.CreatePimV2Integration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePimV2Integration`: PimV2Integration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsPimV2Api.CreatePimV2Integration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePimV2IntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pimV2Integration** | [**PimV2Integration**](PimV2Integration.md) |  | 

### Return type

[**PimV2Integration**](pim.v2.Integration.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePimV2Integration

> DeletePimV2Integration(ctx, id).Environment(environment).Execute()

Delete an Integration



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
    environment := "env-00000" // string | Scope the operation to the given environment.
    id := "id_example" // string | The unique identifier for the integration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsPimV2Api.DeletePimV2Integration(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsPimV2Api.DeletePimV2Integration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the integration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePimV2IntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

 (empty response body)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPimV2Integration

> PimV2Integration GetPimV2Integration(ctx, id).Environment(environment).Execute()

Read an Integration



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
    environment := "env-00000" // string | Scope the operation to the given environment.
    id := "id_example" // string | The unique identifier for the integration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsPimV2Api.GetPimV2Integration(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsPimV2Api.GetPimV2Integration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPimV2Integration`: PimV2Integration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsPimV2Api.GetPimV2Integration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the integration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPimV2IntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**PimV2Integration**](pim.v2.Integration.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPimV2Integrations

> PimV2IntegrationList ListPimV2Integrations(ctx).Environment(environment).DisplayName(displayName).Provider(provider).Status(status).PageSize(pageSize).PageToken(pageToken).Execute()

List of Integrations



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
    environment := "env-00000" // string | Filter the results by exact match for environment.
    displayName := "bigquery_provider_integration" // string | Filter the results by a partial search of display_name. (optional)
    provider := "GCP" // string | Filter the results by exact match for provider. (optional)
    status := "CREATED" // string | Filter the results by exact match for status. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsPimV2Api.ListPimV2Integrations(context.Background()).Environment(environment).DisplayName(displayName).Provider(provider).Status(status).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsPimV2Api.ListPimV2Integrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPimV2Integrations`: PimV2IntegrationList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsPimV2Api.ListPimV2Integrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPimV2IntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **displayName** | **string** | Filter the results by a partial search of display_name. | 
 **provider** | **string** | Filter the results by exact match for provider. | 
 **status** | **string** | Filter the results by exact match for status. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**PimV2IntegrationList**](pim.v2.IntegrationList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePimV2Integration

> PimV2Integration UpdatePimV2Integration(ctx, id).PimV2IntegrationUpdate(pimV2IntegrationUpdate).Execute()

Update an Integration



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
    id := "id_example" // string | The unique identifier for the integration.
    pimV2IntegrationUpdate := *openapiclient.NewPimV2IntegrationUpdate() // PimV2IntegrationUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsPimV2Api.UpdatePimV2Integration(context.Background(), id).PimV2IntegrationUpdate(pimV2IntegrationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsPimV2Api.UpdatePimV2Integration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePimV2Integration`: PimV2Integration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsPimV2Api.UpdatePimV2Integration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the integration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePimV2IntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pimV2IntegrationUpdate** | [**PimV2IntegrationUpdate**](PimV2IntegrationUpdate.md) |  | 

### Return type

[**PimV2Integration**](pim.v2.Integration.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidatePimV2Integration

> ValidatePimV2Integration(ctx).PimV2IntegrationValidateRequest(pimV2IntegrationValidateRequest).Execute()

Validate an Integration



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
    pimV2IntegrationValidateRequest := *openapiclient.NewPimV2IntegrationValidateRequest() // PimV2IntegrationValidateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsPimV2Api.ValidatePimV2Integration(context.Background()).PimV2IntegrationValidateRequest(pimV2IntegrationValidateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsPimV2Api.ValidatePimV2Integration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidatePimV2IntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pimV2IntegrationValidateRequest** | [**PimV2IntegrationValidateRequest**](PimV2IntegrationValidateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

