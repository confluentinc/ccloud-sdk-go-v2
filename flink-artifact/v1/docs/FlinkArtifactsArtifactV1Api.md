# \FlinkArtifactsArtifactV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateArtifactV1FlinkArtifact**](FlinkArtifactsArtifactV1Api.md#CreateArtifactV1FlinkArtifact) | **Post** /artifact/v1/flink-artifacts | Create a new Flink Artifact.
[**DeleteArtifactV1FlinkArtifact**](FlinkArtifactsArtifactV1Api.md#DeleteArtifactV1FlinkArtifact) | **Delete** /artifact/v1/flink-artifacts/{id} | Delete a Flink Artifact
[**GetArtifactV1FlinkArtifact**](FlinkArtifactsArtifactV1Api.md#GetArtifactV1FlinkArtifact) | **Get** /artifact/v1/flink-artifacts/{id} | Read a Flink Artifact
[**ListArtifactV1FlinkArtifacts**](FlinkArtifactsArtifactV1Api.md#ListArtifactV1FlinkArtifacts) | **Get** /artifact/v1/flink-artifacts | List of Flink Artifacts
[**UpdateArtifactV1FlinkArtifact**](FlinkArtifactsArtifactV1Api.md#UpdateArtifactV1FlinkArtifact) | **Patch** /artifact/v1/flink-artifacts/{id} | Update a Flink Artifact



## CreateArtifactV1FlinkArtifact

> ArtifactV1FlinkArtifact CreateArtifactV1FlinkArtifact(ctx).Cloud(cloud).Region(region).InlineObject(inlineObject).Execute()

Create a new Flink Artifact.



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
    cloud := "AWS" // string | Scope the operation to the given cloud.
    region := "us-east-1" // string | Scope the operation to the given region.
    inlineObject := *openapiclient.NewInlineObject("AWS", "us-east-1", "env-00000", "DisplayName_example", "io.confluent.example.SumScalarFunction", openapiclient.InlineObjectUploadSourceOneOf{ArtifactV1UploadSourcePresignedUrl: openapiclient.NewArtifactV1UploadSourcePresignedUrl()}) // InlineObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlinkArtifactsArtifactV1Api.CreateArtifactV1FlinkArtifact(context.Background()).Cloud(cloud).Region(region).InlineObject(inlineObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlinkArtifactsArtifactV1Api.CreateArtifactV1FlinkArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateArtifactV1FlinkArtifact`: ArtifactV1FlinkArtifact
    fmt.Fprintf(os.Stdout, "Response from `FlinkArtifactsArtifactV1Api.CreateArtifactV1FlinkArtifact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateArtifactV1FlinkArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud** | **string** | Scope the operation to the given cloud. | 
 **region** | **string** | Scope the operation to the given region. | 
 **inlineObject** | [**InlineObject**](InlineObject.md) |  | 

### Return type

[**ArtifactV1FlinkArtifact**](artifact.v1.FlinkArtifact.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteArtifactV1FlinkArtifact

> DeleteArtifactV1FlinkArtifact(ctx, id).Cloud(cloud).Region(region).Execute()

Delete a Flink Artifact



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
    cloud := "AWS" // string | Scope the operation to the given cloud.
    region := "us-east-1" // string | Scope the operation to the given region.
    id := "id_example" // string | The unique identifier for the flink artifact.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlinkArtifactsArtifactV1Api.DeleteArtifactV1FlinkArtifact(context.Background(), id).Cloud(cloud).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlinkArtifactsArtifactV1Api.DeleteArtifactV1FlinkArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the flink artifact. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArtifactV1FlinkArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud** | **string** | Scope the operation to the given cloud. | 
 **region** | **string** | Scope the operation to the given region. | 


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


## GetArtifactV1FlinkArtifact

> ArtifactV1FlinkArtifact GetArtifactV1FlinkArtifact(ctx, id).Cloud(cloud).Region(region).Execute()

Read a Flink Artifact



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
    cloud := "AWS" // string | Scope the operation to the given cloud.
    region := "us-east-1" // string | Scope the operation to the given region.
    id := "id_example" // string | The unique identifier for the flink artifact.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlinkArtifactsArtifactV1Api.GetArtifactV1FlinkArtifact(context.Background(), id).Cloud(cloud).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlinkArtifactsArtifactV1Api.GetArtifactV1FlinkArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArtifactV1FlinkArtifact`: ArtifactV1FlinkArtifact
    fmt.Fprintf(os.Stdout, "Response from `FlinkArtifactsArtifactV1Api.GetArtifactV1FlinkArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the flink artifact. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactV1FlinkArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud** | **string** | Scope the operation to the given cloud. | 
 **region** | **string** | Scope the operation to the given region. | 


### Return type

[**ArtifactV1FlinkArtifact**](artifact.v1.FlinkArtifact.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListArtifactV1FlinkArtifacts

> ArtifactV1FlinkArtifactList ListArtifactV1FlinkArtifacts(ctx).Cloud(cloud).Region(region).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()

List of Flink Artifacts



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
    cloud := "AWS" // string | Filter the results by exact match for cloud.
    region := "us-east-1" // string | Filter the results by exact match for region.
    environment := "env-00000" // string | Filter the results by exact match for environment. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlinkArtifactsArtifactV1Api.ListArtifactV1FlinkArtifacts(context.Background()).Cloud(cloud).Region(region).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlinkArtifactsArtifactV1Api.ListArtifactV1FlinkArtifacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListArtifactV1FlinkArtifacts`: ArtifactV1FlinkArtifactList
    fmt.Fprintf(os.Stdout, "Response from `FlinkArtifactsArtifactV1Api.ListArtifactV1FlinkArtifacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListArtifactV1FlinkArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud** | **string** | Filter the results by exact match for cloud. | 
 **region** | **string** | Filter the results by exact match for region. | 
 **environment** | **string** | Filter the results by exact match for environment. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**ArtifactV1FlinkArtifactList**](artifact.v1.FlinkArtifactList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateArtifactV1FlinkArtifact

> ArtifactV1FlinkArtifact UpdateArtifactV1FlinkArtifact(ctx, id).Cloud(cloud).Region(region).ArtifactV1FlinkArtifactUpdate(artifactV1FlinkArtifactUpdate).Execute()

Update a Flink Artifact



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
    cloud := "AWS" // string | Scope the operation to the given cloud.
    region := "us-east-1" // string | Scope the operation to the given region.
    id := "id_example" // string | The unique identifier for the flink artifact.
    artifactV1FlinkArtifactUpdate := *openapiclient.NewArtifactV1FlinkArtifactUpdate() // ArtifactV1FlinkArtifactUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlinkArtifactsArtifactV1Api.UpdateArtifactV1FlinkArtifact(context.Background(), id).Cloud(cloud).Region(region).ArtifactV1FlinkArtifactUpdate(artifactV1FlinkArtifactUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlinkArtifactsArtifactV1Api.UpdateArtifactV1FlinkArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateArtifactV1FlinkArtifact`: ArtifactV1FlinkArtifact
    fmt.Fprintf(os.Stdout, "Response from `FlinkArtifactsArtifactV1Api.UpdateArtifactV1FlinkArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the flink artifact. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateArtifactV1FlinkArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud** | **string** | Scope the operation to the given cloud. | 
 **region** | **string** | Scope the operation to the given region. | 

 **artifactV1FlinkArtifactUpdate** | [**ArtifactV1FlinkArtifactUpdate**](ArtifactV1FlinkArtifactUpdate.md) |  | 

### Return type

[**ArtifactV1FlinkArtifact**](artifact.v1.FlinkArtifact.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

