# \TransitGatewayAttachmentsNetworkingV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkingV1TransitGatewayAttachment**](TransitGatewayAttachmentsNetworkingV1Api.md#CreateNetworkingV1TransitGatewayAttachment) | **Post** /networking/v1/transit-gateway-attachments | Create a Transit Gateway Attachment
[**DeleteNetworkingV1TransitGatewayAttachment**](TransitGatewayAttachmentsNetworkingV1Api.md#DeleteNetworkingV1TransitGatewayAttachment) | **Delete** /networking/v1/transit-gateway-attachments/{id} | Delete a Transit Gateway Attachment
[**GetNetworkingV1TransitGatewayAttachment**](TransitGatewayAttachmentsNetworkingV1Api.md#GetNetworkingV1TransitGatewayAttachment) | **Get** /networking/v1/transit-gateway-attachments/{id} | Read a Transit Gateway Attachment
[**ListNetworkingV1TransitGatewayAttachments**](TransitGatewayAttachmentsNetworkingV1Api.md#ListNetworkingV1TransitGatewayAttachments) | **Get** /networking/v1/transit-gateway-attachments | List of Transit Gateway Attachments
[**UpdateNetworkingV1TransitGatewayAttachment**](TransitGatewayAttachmentsNetworkingV1Api.md#UpdateNetworkingV1TransitGatewayAttachment) | **Patch** /networking/v1/transit-gateway-attachments/{id} | Update a Transit Gateway Attachment



## CreateNetworkingV1TransitGatewayAttachment

> NetworkingV1TransitGatewayAttachment CreateNetworkingV1TransitGatewayAttachment(ctx).NetworkingV1TransitGatewayAttachment(networkingV1TransitGatewayAttachment).Execute()

Create a Transit Gateway Attachment



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
    networkingV1TransitGatewayAttachment := *openapiclient.NewNetworkingV1TransitGatewayAttachment() // NetworkingV1TransitGatewayAttachment |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransitGatewayAttachmentsNetworkingV1Api.CreateNetworkingV1TransitGatewayAttachment(context.Background()).NetworkingV1TransitGatewayAttachment(networkingV1TransitGatewayAttachment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAttachmentsNetworkingV1Api.CreateNetworkingV1TransitGatewayAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkingV1TransitGatewayAttachment`: NetworkingV1TransitGatewayAttachment
    fmt.Fprintf(os.Stdout, "Response from `TransitGatewayAttachmentsNetworkingV1Api.CreateNetworkingV1TransitGatewayAttachment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkingV1TransitGatewayAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkingV1TransitGatewayAttachment** | [**NetworkingV1TransitGatewayAttachment**](NetworkingV1TransitGatewayAttachment.md) |  | 

### Return type

[**NetworkingV1TransitGatewayAttachment**](networking.v1.TransitGatewayAttachment.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkingV1TransitGatewayAttachment

> DeleteNetworkingV1TransitGatewayAttachment(ctx, id).Environment(environment).Execute()

Delete a Transit Gateway Attachment



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
    id := "id_example" // string | The unique identifier for the transit gateway attachment.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransitGatewayAttachmentsNetworkingV1Api.DeleteNetworkingV1TransitGatewayAttachment(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAttachmentsNetworkingV1Api.DeleteNetworkingV1TransitGatewayAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the transit gateway attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkingV1TransitGatewayAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkingV1TransitGatewayAttachment

> NetworkingV1TransitGatewayAttachment GetNetworkingV1TransitGatewayAttachment(ctx, id).Environment(environment).Execute()

Read a Transit Gateway Attachment



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
    id := "id_example" // string | The unique identifier for the transit gateway attachment.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransitGatewayAttachmentsNetworkingV1Api.GetNetworkingV1TransitGatewayAttachment(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAttachmentsNetworkingV1Api.GetNetworkingV1TransitGatewayAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkingV1TransitGatewayAttachment`: NetworkingV1TransitGatewayAttachment
    fmt.Fprintf(os.Stdout, "Response from `TransitGatewayAttachmentsNetworkingV1Api.GetNetworkingV1TransitGatewayAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the transit gateway attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkingV1TransitGatewayAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**NetworkingV1TransitGatewayAttachment**](networking.v1.TransitGatewayAttachment.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkingV1TransitGatewayAttachments

> NetworkingV1TransitGatewayAttachmentList ListNetworkingV1TransitGatewayAttachments(ctx).Environment(environment).SpecDisplayName(specDisplayName).StatusPhase(statusPhase).SpecNetwork(specNetwork).PageSize(pageSize).PageToken(pageToken).Execute()

List of Transit Gateway Attachments



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
    specDisplayName := *openapiclient.NewMultipleSearchFilter() // MultipleSearchFilter | Filter the results by exact match for spec.display_name. Pass multiple times to see results matching any of the values. (optional)
    statusPhase := *openapiclient.NewMultipleSearchFilter() // MultipleSearchFilter | Filter the results by exact match for status.phase. Pass multiple times to see results matching any of the values. (optional)
    specNetwork :=  // MultipleSearchFilter | Filter the results by exact match for spec.network. Pass multiple times to see results matching any of the values. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransitGatewayAttachmentsNetworkingV1Api.ListNetworkingV1TransitGatewayAttachments(context.Background()).Environment(environment).SpecDisplayName(specDisplayName).StatusPhase(statusPhase).SpecNetwork(specNetwork).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAttachmentsNetworkingV1Api.ListNetworkingV1TransitGatewayAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkingV1TransitGatewayAttachments`: NetworkingV1TransitGatewayAttachmentList
    fmt.Fprintf(os.Stdout, "Response from `TransitGatewayAttachmentsNetworkingV1Api.ListNetworkingV1TransitGatewayAttachments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkingV1TransitGatewayAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **specDisplayName** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter the results by exact match for spec.display_name. Pass multiple times to see results matching any of the values. | 
 **statusPhase** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter the results by exact match for status.phase. Pass multiple times to see results matching any of the values. | 
 **specNetwork** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter the results by exact match for spec.network. Pass multiple times to see results matching any of the values. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**NetworkingV1TransitGatewayAttachmentList**](networking.v1.TransitGatewayAttachmentList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkingV1TransitGatewayAttachment

> NetworkingV1TransitGatewayAttachment UpdateNetworkingV1TransitGatewayAttachment(ctx, id).NetworkingV1TransitGatewayAttachmentUpdate(networkingV1TransitGatewayAttachmentUpdate).Execute()

Update a Transit Gateway Attachment



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
    id := "id_example" // string | The unique identifier for the transit gateway attachment.
    networkingV1TransitGatewayAttachmentUpdate := *openapiclient.NewNetworkingV1TransitGatewayAttachmentUpdate() // NetworkingV1TransitGatewayAttachmentUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransitGatewayAttachmentsNetworkingV1Api.UpdateNetworkingV1TransitGatewayAttachment(context.Background(), id).NetworkingV1TransitGatewayAttachmentUpdate(networkingV1TransitGatewayAttachmentUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAttachmentsNetworkingV1Api.UpdateNetworkingV1TransitGatewayAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkingV1TransitGatewayAttachment`: NetworkingV1TransitGatewayAttachment
    fmt.Fprintf(os.Stdout, "Response from `TransitGatewayAttachmentsNetworkingV1Api.UpdateNetworkingV1TransitGatewayAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the transit gateway attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkingV1TransitGatewayAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkingV1TransitGatewayAttachmentUpdate** | [**NetworkingV1TransitGatewayAttachmentUpdate**](NetworkingV1TransitGatewayAttachmentUpdate.md) |  | 

### Return type

[**NetworkingV1TransitGatewayAttachment**](networking.v1.TransitGatewayAttachment.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

