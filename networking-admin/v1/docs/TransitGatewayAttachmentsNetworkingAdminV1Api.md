# \TransitGatewayAttachmentsNetworkingAdminV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkingAdminV1TransitGatewayAttachment**](TransitGatewayAttachmentsNetworkingAdminV1Api.md#GetNetworkingAdminV1TransitGatewayAttachment) | **Get** /networking-admin/v1/transit-gateway-attachments/{id} | Read a Transit Gateway Attachment
[**ListNetworkingAdminV1TransitGatewayAttachments**](TransitGatewayAttachmentsNetworkingAdminV1Api.md#ListNetworkingAdminV1TransitGatewayAttachments) | **Get** /networking-admin/v1/transit-gateway-attachments | List of Transit Gateway Attachments



## GetNetworkingAdminV1TransitGatewayAttachment

> NetworkingAdminV1TransitGatewayAttachment GetNetworkingAdminV1TransitGatewayAttachment(ctx, id).Environment(environment).Execute()

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
    resp, r, err := api_client.TransitGatewayAttachmentsNetworkingAdminV1Api.GetNetworkingAdminV1TransitGatewayAttachment(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAttachmentsNetworkingAdminV1Api.GetNetworkingAdminV1TransitGatewayAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkingAdminV1TransitGatewayAttachment`: NetworkingAdminV1TransitGatewayAttachment
    fmt.Fprintf(os.Stdout, "Response from `TransitGatewayAttachmentsNetworkingAdminV1Api.GetNetworkingAdminV1TransitGatewayAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the transit gateway attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkingAdminV1TransitGatewayAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**NetworkingAdminV1TransitGatewayAttachment**](networking-admin.v1.TransitGatewayAttachment.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkingAdminV1TransitGatewayAttachments

> NetworkingAdminV1TransitGatewayAttachmentList ListNetworkingAdminV1TransitGatewayAttachments(ctx).Environment(environment).SpecDisplayName(specDisplayName).StatusPhase(statusPhase).SpecNetwork(specNetwork).PageSize(pageSize).PageToken(pageToken).Execute()

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
    resp, r, err := api_client.TransitGatewayAttachmentsNetworkingAdminV1Api.ListNetworkingAdminV1TransitGatewayAttachments(context.Background()).Environment(environment).SpecDisplayName(specDisplayName).StatusPhase(statusPhase).SpecNetwork(specNetwork).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAttachmentsNetworkingAdminV1Api.ListNetworkingAdminV1TransitGatewayAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkingAdminV1TransitGatewayAttachments`: NetworkingAdminV1TransitGatewayAttachmentList
    fmt.Fprintf(os.Stdout, "Response from `TransitGatewayAttachmentsNetworkingAdminV1Api.ListNetworkingAdminV1TransitGatewayAttachments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkingAdminV1TransitGatewayAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **specDisplayName** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter the results by exact match for spec.display_name. Pass multiple times to see results matching any of the values. | 
 **statusPhase** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter the results by exact match for status.phase. Pass multiple times to see results matching any of the values. | 
 **specNetwork** | [**MultipleSearchFilter**](MultipleSearchFilter.md) | Filter the results by exact match for spec.network. Pass multiple times to see results matching any of the values. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**NetworkingAdminV1TransitGatewayAttachmentList**](networking-admin.v1.TransitGatewayAttachmentList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

