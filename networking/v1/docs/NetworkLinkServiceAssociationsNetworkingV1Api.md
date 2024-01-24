# \NetworkLinkServiceAssociationsNetworkingV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkingV1NetworkLinkServiceAssociation**](NetworkLinkServiceAssociationsNetworkingV1Api.md#GetNetworkingV1NetworkLinkServiceAssociation) | **Get** /networking/v1/network-link-service-associations/{id} | Read a Network Link Service Association
[**ListNetworkingV1NetworkLinkServiceAssociations**](NetworkLinkServiceAssociationsNetworkingV1Api.md#ListNetworkingV1NetworkLinkServiceAssociations) | **Get** /networking/v1/network-link-service-associations | List of Network Link Service Associations



## GetNetworkingV1NetworkLinkServiceAssociation

> NetworkingV1NetworkLinkServiceAssociation GetNetworkingV1NetworkLinkServiceAssociation(ctx, id).SpecNetworkLinkService(specNetworkLinkService).Environment(environment).Execute()

Read a Network Link Service Association



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
    specNetworkLinkService := "nls-abcde" // string | Scope the operation to the given spec.network_link_service.
    environment := "env-00000" // string | Scope the operation to the given environment.
    id := "id_example" // string | The unique identifier for the network link service association.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkLinkServiceAssociationsNetworkingV1Api.GetNetworkingV1NetworkLinkServiceAssociation(context.Background(), id).SpecNetworkLinkService(specNetworkLinkService).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLinkServiceAssociationsNetworkingV1Api.GetNetworkingV1NetworkLinkServiceAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkingV1NetworkLinkServiceAssociation`: NetworkingV1NetworkLinkServiceAssociation
    fmt.Fprintf(os.Stdout, "Response from `NetworkLinkServiceAssociationsNetworkingV1Api.GetNetworkingV1NetworkLinkServiceAssociation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the network link service association. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkingV1NetworkLinkServiceAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **specNetworkLinkService** | **string** | Scope the operation to the given spec.network_link_service. | 
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**NetworkingV1NetworkLinkServiceAssociation**](networking.v1.NetworkLinkServiceAssociation.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkingV1NetworkLinkServiceAssociations

> NetworkingV1NetworkLinkServiceAssociationList ListNetworkingV1NetworkLinkServiceAssociations(ctx).SpecNetworkLinkService(specNetworkLinkService).Environment(environment).StatusPhase(statusPhase).PageSize(pageSize).PageToken(pageToken).Execute()

List of Network Link Service Associations



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
    specNetworkLinkService := "nls-abcde" // string | Filter the results by exact match for spec.network_link_service.
    environment := "env-00000" // string | Filter the results by exact match for environment.
    statusPhase := []string{"Inner_example"} // []string | Filter the results by exact match for status.phase. Pass multiple times to see results matching any of the values. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkLinkServiceAssociationsNetworkingV1Api.ListNetworkingV1NetworkLinkServiceAssociations(context.Background()).SpecNetworkLinkService(specNetworkLinkService).Environment(environment).StatusPhase(statusPhase).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLinkServiceAssociationsNetworkingV1Api.ListNetworkingV1NetworkLinkServiceAssociations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkingV1NetworkLinkServiceAssociations`: NetworkingV1NetworkLinkServiceAssociationList
    fmt.Fprintf(os.Stdout, "Response from `NetworkLinkServiceAssociationsNetworkingV1Api.ListNetworkingV1NetworkLinkServiceAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkingV1NetworkLinkServiceAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **specNetworkLinkService** | **string** | Filter the results by exact match for spec.network_link_service. | 
 **environment** | **string** | Filter the results by exact match for environment. | 
 **statusPhase** | **[]string** | Filter the results by exact match for status.phase. Pass multiple times to see results matching any of the values. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**NetworkingV1NetworkLinkServiceAssociationList**](networking.v1.NetworkLinkServiceAssociationList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

