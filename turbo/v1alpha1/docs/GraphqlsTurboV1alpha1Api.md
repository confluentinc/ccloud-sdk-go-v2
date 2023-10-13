# \GraphqlsTurboV1alpha1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryTurboV1alpha1Graphql**](GraphqlsTurboV1alpha1Api.md#QueryTurboV1alpha1Graphql) | **Post** /turbo/v1alpha1/graphql | Query a Graphql



## QueryTurboV1alpha1Graphql

> string QueryTurboV1alpha1Graphql(ctx).XPrincipal(xPrincipal).XPrincipalResourceId(xPrincipalResourceId).XOrganizationId(xOrganizationId).XOrganizationResourceId(xOrganizationResourceId).XExternalIdentity(xExternalIdentity).TurboV1alpha1GraphqlRequest(turboV1alpha1GraphqlRequest).Execute()

Query a Graphql



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
    xPrincipal := "xPrincipal_example" // string |  (optional)
    xPrincipalResourceId := "xPrincipalResourceId_example" // string |  (optional)
    xOrganizationId := "xOrganizationId_example" // string |  (optional)
    xOrganizationResourceId := "xOrganizationResourceId_example" // string |  (optional)
    xExternalIdentity := "xExternalIdentity_example" // string |  (optional)
    turboV1alpha1GraphqlRequest := *openapiclient.NewTurboV1alpha1GraphqlRequest() // TurboV1alpha1GraphqlRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GraphqlsTurboV1alpha1Api.QueryTurboV1alpha1Graphql(context.Background()).XPrincipal(xPrincipal).XPrincipalResourceId(xPrincipalResourceId).XOrganizationId(xOrganizationId).XOrganizationResourceId(xOrganizationResourceId).XExternalIdentity(xExternalIdentity).TurboV1alpha1GraphqlRequest(turboV1alpha1GraphqlRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphqlsTurboV1alpha1Api.QueryTurboV1alpha1Graphql``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryTurboV1alpha1Graphql`: string
    fmt.Fprintf(os.Stdout, "Response from `GraphqlsTurboV1alpha1Api.QueryTurboV1alpha1Graphql`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryTurboV1alpha1GraphqlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPrincipal** | **string** |  | 
 **xPrincipalResourceId** | **string** |  | 
 **xOrganizationId** | **string** |  | 
 **xOrganizationResourceId** | **string** |  | 
 **xExternalIdentity** | **string** |  | 
 **turboV1alpha1GraphqlRequest** | [**TurboV1alpha1GraphqlRequest**](TurboV1alpha1GraphqlRequest.md) |  | 

### Return type

**string**

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

