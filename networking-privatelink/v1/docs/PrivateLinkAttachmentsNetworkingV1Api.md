# \PrivateLinkAttachmentsNetworkingV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkingV1PrivateLinkAttachment**](PrivateLinkAttachmentsNetworkingV1Api.md#CreateNetworkingV1PrivateLinkAttachment) | **Post** /networking/v1/private-link-attachments | Create a Private Link Attachment
[**DeleteNetworkingV1PrivateLinkAttachment**](PrivateLinkAttachmentsNetworkingV1Api.md#DeleteNetworkingV1PrivateLinkAttachment) | **Delete** /networking/v1/private-link-attachments/{id} | Delete a Private Link Attachment
[**GetNetworkingV1PrivateLinkAttachment**](PrivateLinkAttachmentsNetworkingV1Api.md#GetNetworkingV1PrivateLinkAttachment) | **Get** /networking/v1/private-link-attachments/{id} | Read a Private Link Attachment
[**ListNetworkingV1PrivateLinkAttachments**](PrivateLinkAttachmentsNetworkingV1Api.md#ListNetworkingV1PrivateLinkAttachments) | **Get** /networking/v1/private-link-attachments | List of Private Link Attachments
[**UpdateNetworkingV1PrivateLinkAttachment**](PrivateLinkAttachmentsNetworkingV1Api.md#UpdateNetworkingV1PrivateLinkAttachment) | **Patch** /networking/v1/private-link-attachments/{id} | Update a Private Link Attachment



## CreateNetworkingV1PrivateLinkAttachment

> NetworkingV1PrivateLinkAttachment CreateNetworkingV1PrivateLinkAttachment(ctx).NetworkingV1PrivateLinkAttachment(networkingV1PrivateLinkAttachment).Execute()

Create a Private Link Attachment



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
    networkingV1PrivateLinkAttachment := *openapiclient.NewNetworkingV1PrivateLinkAttachment() // NetworkingV1PrivateLinkAttachment |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivateLinkAttachmentsNetworkingV1Api.CreateNetworkingV1PrivateLinkAttachment(context.Background()).NetworkingV1PrivateLinkAttachment(networkingV1PrivateLinkAttachment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateLinkAttachmentsNetworkingV1Api.CreateNetworkingV1PrivateLinkAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkingV1PrivateLinkAttachment`: NetworkingV1PrivateLinkAttachment
    fmt.Fprintf(os.Stdout, "Response from `PrivateLinkAttachmentsNetworkingV1Api.CreateNetworkingV1PrivateLinkAttachment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkingV1PrivateLinkAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkingV1PrivateLinkAttachment** | [**NetworkingV1PrivateLinkAttachment**](NetworkingV1PrivateLinkAttachment.md) |  | 

### Return type

[**NetworkingV1PrivateLinkAttachment**](networking.v1.PrivateLinkAttachment.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkingV1PrivateLinkAttachment

> DeleteNetworkingV1PrivateLinkAttachment(ctx, id).Environment(environment).Execute()

Delete a Private Link Attachment



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
    id := "id_example" // string | The unique identifier for the private link attachment.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivateLinkAttachmentsNetworkingV1Api.DeleteNetworkingV1PrivateLinkAttachment(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateLinkAttachmentsNetworkingV1Api.DeleteNetworkingV1PrivateLinkAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the private link attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkingV1PrivateLinkAttachmentRequest struct via the builder pattern


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


## GetNetworkingV1PrivateLinkAttachment

> NetworkingV1PrivateLinkAttachment GetNetworkingV1PrivateLinkAttachment(ctx, id).Environment(environment).Execute()

Read a Private Link Attachment



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
    id := "id_example" // string | The unique identifier for the private link attachment.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivateLinkAttachmentsNetworkingV1Api.GetNetworkingV1PrivateLinkAttachment(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateLinkAttachmentsNetworkingV1Api.GetNetworkingV1PrivateLinkAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkingV1PrivateLinkAttachment`: NetworkingV1PrivateLinkAttachment
    fmt.Fprintf(os.Stdout, "Response from `PrivateLinkAttachmentsNetworkingV1Api.GetNetworkingV1PrivateLinkAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the private link attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkingV1PrivateLinkAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**NetworkingV1PrivateLinkAttachment**](networking.v1.PrivateLinkAttachment.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkingV1PrivateLinkAttachments

> NetworkingV1PrivateLinkAttachmentList ListNetworkingV1PrivateLinkAttachments(ctx).Environment(environment).SpecDisplayName(specDisplayName).SpecCloud(specCloud).SpecRegion(specRegion).StatusPhase(statusPhase).PageSize(pageSize).PageToken(pageToken).Execute()

List of Private Link Attachments



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
    specDisplayName := []string{"Inner_example"} // []string | Filter the results by exact match for spec.display_name. Pass multiple times to see results matching any of the values. (optional)
    specCloud := []string{"Inner_example"} // []string | Filter the results by exact match for spec.cloud. Pass multiple times to see results matching any of the values. (optional)
    specRegion := []string{"Inner_example"} // []string | Filter the results by exact match for spec.region. Pass multiple times to see results matching any of the values. (optional)
    statusPhase := []string{"Inner_example"} // []string | Filter the results by exact match for status.phase. Pass multiple times to see results matching any of the values. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivateLinkAttachmentsNetworkingV1Api.ListNetworkingV1PrivateLinkAttachments(context.Background()).Environment(environment).SpecDisplayName(specDisplayName).SpecCloud(specCloud).SpecRegion(specRegion).StatusPhase(statusPhase).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateLinkAttachmentsNetworkingV1Api.ListNetworkingV1PrivateLinkAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkingV1PrivateLinkAttachments`: NetworkingV1PrivateLinkAttachmentList
    fmt.Fprintf(os.Stdout, "Response from `PrivateLinkAttachmentsNetworkingV1Api.ListNetworkingV1PrivateLinkAttachments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkingV1PrivateLinkAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **specDisplayName** | **[]string** | Filter the results by exact match for spec.display_name. Pass multiple times to see results matching any of the values. | 
 **specCloud** | **[]string** | Filter the results by exact match for spec.cloud. Pass multiple times to see results matching any of the values. | 
 **specRegion** | **[]string** | Filter the results by exact match for spec.region. Pass multiple times to see results matching any of the values. | 
 **statusPhase** | **[]string** | Filter the results by exact match for status.phase. Pass multiple times to see results matching any of the values. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**NetworkingV1PrivateLinkAttachmentList**](networking.v1.PrivateLinkAttachmentList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkingV1PrivateLinkAttachment

> NetworkingV1PrivateLinkAttachment UpdateNetworkingV1PrivateLinkAttachment(ctx, id).NetworkingV1PrivateLinkAttachmentUpdate(networkingV1PrivateLinkAttachmentUpdate).Execute()

Update a Private Link Attachment



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
    id := "id_example" // string | The unique identifier for the private link attachment.
    networkingV1PrivateLinkAttachmentUpdate := *openapiclient.NewNetworkingV1PrivateLinkAttachmentUpdate() // NetworkingV1PrivateLinkAttachmentUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivateLinkAttachmentsNetworkingV1Api.UpdateNetworkingV1PrivateLinkAttachment(context.Background(), id).NetworkingV1PrivateLinkAttachmentUpdate(networkingV1PrivateLinkAttachmentUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateLinkAttachmentsNetworkingV1Api.UpdateNetworkingV1PrivateLinkAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkingV1PrivateLinkAttachment`: NetworkingV1PrivateLinkAttachment
    fmt.Fprintf(os.Stdout, "Response from `PrivateLinkAttachmentsNetworkingV1Api.UpdateNetworkingV1PrivateLinkAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the private link attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkingV1PrivateLinkAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkingV1PrivateLinkAttachmentUpdate** | [**NetworkingV1PrivateLinkAttachmentUpdate**](NetworkingV1PrivateLinkAttachmentUpdate.md) |  | 

### Return type

[**NetworkingV1PrivateLinkAttachment**](networking.v1.PrivateLinkAttachment.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

