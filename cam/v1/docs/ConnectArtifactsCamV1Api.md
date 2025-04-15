# \ConnectArtifactsCamV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCamV1ConnectArtifact**](ConnectArtifactsCamV1Api.md#CreateCamV1ConnectArtifact) | **Post** /cam/v1/connect-artifacts | Create a new Connect Artifact.
[**DeleteCamV1ConnectArtifact**](ConnectArtifactsCamV1Api.md#DeleteCamV1ConnectArtifact) | **Delete** /cam/v1/connect-artifacts/{id} | Delete a Connect Artifact
[**GetCamV1ConnectArtifact**](ConnectArtifactsCamV1Api.md#GetCamV1ConnectArtifact) | **Get** /cam/v1/connect-artifacts/{id} | Read a Connect Artifact
[**ListCamV1ConnectArtifacts**](ConnectArtifactsCamV1Api.md#ListCamV1ConnectArtifacts) | **Get** /cam/v1/connect-artifacts | List of Connect Artifacts



## CreateCamV1ConnectArtifact

> CamV1ConnectArtifact CreateCamV1ConnectArtifact(ctx).SpecCloud(specCloud).Environment(environment).CamV1ConnectArtifact(camV1ConnectArtifact).Execute()

Create a new Connect Artifact.



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
    specCloud := "AWS" // string | Scope the operation to the given spec.cloud.
    environment := "env-00000" // string | Scope the operation to the given environment.
    camV1ConnectArtifact := *openapiclient.NewCamV1ConnectArtifact() // CamV1ConnectArtifact |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectArtifactsCamV1Api.CreateCamV1ConnectArtifact(context.Background()).SpecCloud(specCloud).Environment(environment).CamV1ConnectArtifact(camV1ConnectArtifact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectArtifactsCamV1Api.CreateCamV1ConnectArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCamV1ConnectArtifact`: CamV1ConnectArtifact
    fmt.Fprintf(os.Stdout, "Response from `ConnectArtifactsCamV1Api.CreateCamV1ConnectArtifact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCamV1ConnectArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **specCloud** | **string** | Scope the operation to the given spec.cloud. | 
 **environment** | **string** | Scope the operation to the given environment. | 
 **camV1ConnectArtifact** | [**CamV1ConnectArtifact**](CamV1ConnectArtifact.md) |  | 

### Return type

[**CamV1ConnectArtifact**](cam.v1.ConnectArtifact.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCamV1ConnectArtifact

> DeleteCamV1ConnectArtifact(ctx, id).SpecCloud(specCloud).Environment(environment).Execute()

Delete a Connect Artifact



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
    specCloud := "AWS" // string | Scope the operation to the given spec.cloud.
    environment := "env-00000" // string | Scope the operation to the given environment.
    id := "id_example" // string | The unique identifier for the connect artifact.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectArtifactsCamV1Api.DeleteCamV1ConnectArtifact(context.Background(), id).SpecCloud(specCloud).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectArtifactsCamV1Api.DeleteCamV1ConnectArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the connect artifact. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCamV1ConnectArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **specCloud** | **string** | Scope the operation to the given spec.cloud. | 
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


## GetCamV1ConnectArtifact

> CamV1ConnectArtifact GetCamV1ConnectArtifact(ctx, id).SpecCloud(specCloud).Environment(environment).Execute()

Read a Connect Artifact



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
    specCloud := "AWS" // string | Scope the operation to the given spec.cloud.
    environment := "env-00000" // string | Scope the operation to the given environment.
    id := "id_example" // string | The unique identifier for the connect artifact.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectArtifactsCamV1Api.GetCamV1ConnectArtifact(context.Background(), id).SpecCloud(specCloud).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectArtifactsCamV1Api.GetCamV1ConnectArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCamV1ConnectArtifact`: CamV1ConnectArtifact
    fmt.Fprintf(os.Stdout, "Response from `ConnectArtifactsCamV1Api.GetCamV1ConnectArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the connect artifact. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCamV1ConnectArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **specCloud** | **string** | Scope the operation to the given spec.cloud. | 
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**CamV1ConnectArtifact**](cam.v1.ConnectArtifact.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCamV1ConnectArtifacts

> CamV1ConnectArtifactList ListCamV1ConnectArtifacts(ctx).SpecCloud(specCloud).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()

List of Connect Artifacts



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
    specCloud := "AWS" // string | Filter the results by exact match for spec.cloud.
    environment := "env-00000" // string | Filter the results by exact match for environment.
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectArtifactsCamV1Api.ListCamV1ConnectArtifacts(context.Background()).SpecCloud(specCloud).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectArtifactsCamV1Api.ListCamV1ConnectArtifacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCamV1ConnectArtifacts`: CamV1ConnectArtifactList
    fmt.Fprintf(os.Stdout, "Response from `ConnectArtifactsCamV1Api.ListCamV1ConnectArtifacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCamV1ConnectArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **specCloud** | **string** | Filter the results by exact match for spec.cloud. | 
 **environment** | **string** | Filter the results by exact match for environment. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**CamV1ConnectArtifactList**](cam.v1.ConnectArtifactList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

